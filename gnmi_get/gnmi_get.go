/* Copyright 2017 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Binary gnmi_get performs a get request against a gNMI target.
package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
	"regexp"

	log "github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/google/gnxi/utils"
	"github.com/google/gnxi/utils/credentials"
	"github.com/google/gnxi/utils/xpath"

	pb "github.com/openconfig/gnmi/proto/gnmi"
	gnmi_sonic "github.com/google/gnxi/proto"
	proto "github.com/golang/protobuf/proto"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func parseModelData(s *string) (*pb.ModelData, error) {
	pbModelData := new(pb.ModelData)
	modelDataVars := strings.Split(*s, ",")
	if len(modelDataVars) != 3 {
		return pbModelData, fmt.Errorf("Unable to parse string")
	}
	pbModelData = &pb.ModelData{
		Name:         modelDataVars[0],
		Organization: modelDataVars[1],
		Version:      modelDataVars[2],
	}
	return pbModelData, nil
}

func parseTestcaseRib(val *pb.TypedValue) error {
	var ribTableEntries gnmi_sonic.RibTableEntries
	err := proto.Unmarshal(val.GetProtoBytes(), &ribTableEntries)
	if err != nil {
		return fmt.Errorf("Unmarshal error")
	}

	for i, entry := range(ribTableEntries.GetEntry()) {
		log.Infof("Route Entry: %d, route=%s", i, entry.GetEntryDetail())
	}

	return nil
}

func parseTestcaseBgpNas(val *pb.TypedValue) error {
	fmt.Println("parseTestcaseBgpNas")

	var nas gnmi_sonic.NeighborAddresses
	err := proto.Unmarshal(val.GetProtoBytes(), &nas)
	if err != nil {
		return fmt.Errorf("Unmarshal error")
	}

	for i, na := range(nas.GetNa()) {
		fmt.Println("BGP Entry: ", i, ", GetRemoteAddr=" + na.GetRemoteAddr() +
			", GetLocalAddr=" + na.GetLocalAddr() + ", GetState=" + na.GetState() +
			", GetAsn=" + na.GetAsn() + "")
	}

	return nil
}

func parseTestcaseConfigDbBgpNas(val *pb.TypedValue) error {
	return nil
}

type testCaseParseFunc func(val *pb.TypedValue) error

type testCase struct {
	path string
	parseFunc testCaseParseFunc
}

var (
	xPathFlags       arrayFlags
	pbPathFlags      arrayFlags
	pbModelDataFlags arrayFlags
	targetAddr       = flag.String("target_addr", "localhost:10161", "The target address in the format of host:port")
	targetName       = flag.String("target_name", "hostname.com", "The target name use to verify the hostname returned by TLS handshake")
	timeOut          = flag.Duration("time_out", 10*time.Second, "Timeout for the Get request, 10 seconds by default")
	encodingName     = flag.String("encoding", "PROTO", "value encoding format to be used")
	pathTarget       = flag.String("xpath_target", "NONE", "name of the target for which the path is a member")

	testCases = []testCase {
		testCase {
			path: "/rib/table/entries",
			parseFunc: parseTestcaseRib,
		},
		testCase {
			path: "/rib/table/entries/entry",
			parseFunc: parseTestcaseRib,
		},
		testCase {
			path: "/bgp/neighbors/neighbor/state/neighbor-addresses",
			parseFunc: parseTestcaseBgpNas,
		},
		testCase {
			path: "/BGP_NEIGHBOR",
			parseFunc: parseTestcaseConfigDbBgpNas,
		},
	}
)

func getFullPath(path *pb.Path) string {
	fullPath := ""

	for _, pathElem := range(path.GetElem()) {
		fullPath += "/" + pathElem.GetName()
	}

    return fullPath
}

func checkYangPath(xpath string) int {
	xpath = "/" + xpath

	for i, tc := range(testCases) {
		match, _ := regexp.MatchString(tc.path, xpath)
        if match {
			return i
		}
	}

	return -1
}

func main() {
	flag.Var(&xPathFlags, "xpath", "xpath of the config node to be fetched")
	flag.Var(&pbPathFlags, "pbpath", "protobuf format path of the config node to be fetched")
	flag.Var(&pbModelDataFlags, "model_data", "Data models to be used by the target in the format of 'name,organization,version'")
	flag.Parse()

	// 1. 创建GNMI client
	opts := credentials.ClientCredentials(*targetName)
	conn, err := grpc.Dial(*targetAddr, opts...)
	if err != nil {
		log.Exitf("Dialing to %q failed: %v", *targetAddr, err)
	}
	defer conn.Close()

	cli := pb.NewGNMIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), *timeOut)
	defer cancel()

	// 2. 通过命令行参数，获取编码格式encoding
	encoding, ok := pb.Encoding_value[*encodingName]
	if !ok {
		var gnmiEncodingList []string
		for _, name := range pb.Encoding_name {
			gnmiEncodingList = append(gnmiEncodingList, name)
		}
		log.Exitf("Supported encodings: %s", strings.Join(gnmiEncodingList, ", "))
	}

	// 3. 封装getRequest
	var pbPathList []*pb.Path
	var pbModelDataList []*pb.ModelData
	for _, xPath := range xPathFlags {
		// 检查xpath是否是proto文件支持的path
		if checkYangPath(xPath) == -1 {
			log.Exitf("error in checkYangPath, xPath = %q", xPath)
		}

		pbPath, err := xpath.ToGNMIPath(xPath)
		if err != nil {
			log.Exitf("error in parsing xpath %q to gnmi path", xPath)
		}
		pbPathList = append(pbPathList, pbPath)
	}
	for _, textPbPath := range pbPathFlags {
		var pbPath pb.Path
		if err := proto.UnmarshalText(textPbPath, &pbPath); err != nil {
			log.Exitf("error in unmarshaling %q to gnmi Path", textPbPath)
		}
		pbPathList = append(pbPathList, &pbPath)
	}
	for _, textPbModelData := range pbModelDataFlags {
		pbModelData, err := parseModelData(&textPbModelData)
		if err == nil {
			pbModelDataList = append(pbModelDataList, pbModelData)
		}
	}

	// 3.2 封装Target
	var prefix pb.Path
	prefix.Target = *pathTarget

	// 3.3 封装getRequest
	getRequest := &pb.GetRequest{
		Prefix:	&prefix,
		Encoding:  pb.Encoding(encoding),
		Path:      pbPathList,
		UseModels: pbModelDataList,
	}

	fmt.Println("== getRequest:")
	utils.PrintProto(getRequest)

	// 4. 发送getRequest并获得getResponse
	getResponse, err := cli.Get(ctx, getRequest)
	if err != nil {
		log.Exitf("Get failed: %v", err)
	}

	// 5. 打印、解析getResponse
	fmt.Println("== getResponse:")
	//utils.PrintProto(getResponse)

	for _, notify := range(getResponse.GetNotification()) {
		//if notify.GetPrefix().GetTarget() == "MTNOS" {
			for _, update := range(notify.GetUpdate()) {
				res := checkYangPath(getFullPath(update.GetPath()))

				if res != -1 {
					fmt.Println("== parse getResponse:")
					testCases[res].parseFunc(update.GetVal())
				}
			}
		//}
	}
}

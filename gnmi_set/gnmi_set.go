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

/*
gnmi_set \
  -update local-routes/static-routes/static:"*1" \
  -target_addr 172.18.8.210:8080 -alsologtostderr -insecure true
*/

// Binary gnmi_set performs a set request against a gNMI target with the specified config file.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	log "github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/google/gnxi/utils"
	"github.com/google/gnxi/utils/credentials"
	"github.com/google/gnxi/utils/xpath"

	pb "github.com/openconfig/gnmi/proto/gnmi"
	lr "github.com/google/gnxi/proto"
	proto "github.com/golang/protobuf/proto"
	mt_proto "github.com/Azure/sonic-telemetry/open-proto/oc"
	yw "github.com/openconfig/ygot/proto/ywrapper"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

type TestCaseStaticRoute struct {
	path string
	value lr.StaticRoutesStatic
}

var (
	deleteOpt  arrayFlags
	replaceOpt arrayFlags
	updateOpt  arrayFlags
	targetAddr = flag.String("target_addr", "localhost:10161", "The target address in the format of host:port")
	targetName = flag.String("target_name", "hostname.com", "The target name use to verify the hostname returned by TLS handshake")
	timeOut    = flag.Duration("time_out", 10*time.Second, "Timeout for the Get request, 10 seconds by default")
	pathTarget = flag.String("xpath_target", "NONE", "name of the target for which the path is a member")
)

func buildPbUpdateList(pathValuePairs []string) []*pb.Update {
	var pbUpdateList []*pb.Update
	for _, item := range pathValuePairs {
		pathValuePair := strings.SplitN(item, ":", 2)
		// TODO (leguo): check if any path attribute contains ':'
		if len(pathValuePair) != 2 || len(pathValuePair[1]) == 0 {
			log.Exitf("invalid path-value pair: %v", item)
		}
		pbPath, err := xpath.ToGNMIPath(pathValuePair[0])
		if err != nil {
			log.Exitf("error in parsing xpath %q to gnmi path", pathValuePair[0])
		}
		var pbVal *pb.TypedValue
		if pathValuePair[1] == "*1" {
			testcase := TestCaseStaticRoute{
				path: "local-routes/static-routes/static",
				value: lr.StaticRoutesStatic{
					Address: "111.111.111.111/32",
					NextHops: &lr.StaticNextHops {
						NextHop: []*lr.NextHopsNextHop {
							&lr.NextHopsNextHop {
								Config: &lr.NextHopConfig{
									NextHop: "222.222.222.222/32",
								},
							},
						},
					},
				},
			}
			out, err := proto.Marshal(&testcase.value)
			if err != nil {
				log.Exitf("cannot read data from testcase")
			}
			pbVal = &pb.TypedValue{
				Value: &pb.TypedValue_ProtoBytes{
					ProtoBytes: out,
				},
			}
		} else if pathValuePair[1] == "*2" {
			value := mt_proto.Bgp_Neighbor{
				Enabled: &yw.BoolValue{
					Value: true,
				},
			}
			out, err := proto.Marshal(&value)
			if err != nil {
				log.Exitf("cannot read data from testcase")
			}
			pbVal = &pb.TypedValue{
				Value: &pb.TypedValue_ProtoBytes{
					ProtoBytes: out,
				},
			}
		} else if pathValuePair[1] == "*3" {
			value := mt_proto.Bgp_Neighbor{
				Enabled: &yw.BoolValue{
					Value: false,
				},
			}
			out, err := proto.Marshal(&value)
			if err != nil {
				log.Exitf("cannot read data from testcase")
			}
			pbVal = &pb.TypedValue{
				Value: &pb.TypedValue_ProtoBytes{
					ProtoBytes: out,
				},
			}
		} else if pathValuePair[1] == "*4" {
			value := mt_proto.Interface{
				Enabled: &yw.BoolValue{
					Value: true,
				},
			}
			out, err := proto.Marshal(&value)
			if err != nil {
				log.Exitf("cannot read data from testcase")
			}
			pbVal = &pb.TypedValue{
				Value: &pb.TypedValue_ProtoBytes{
					ProtoBytes: out,
				},
			}
		} else if pathValuePair[1] == "*5" {
			value := mt_proto.Interface{
				Enabled: &yw.BoolValue{
					Value: false,
				},
			}
			out, err := proto.Marshal(&value)
			if err != nil {
				log.Exitf("cannot read data from testcase")
			}
			pbVal = &pb.TypedValue{
				Value: &pb.TypedValue_ProtoBytes{
					ProtoBytes: out,
				},
			}
		} else if pathValuePair[1][0] == '@' {
			jsonFile := pathValuePair[1][1:]
			jsonConfig, err := ioutil.ReadFile(jsonFile)
			if err != nil {
				log.Exitf("cannot read data from file %v", jsonFile)
			}
			jsonConfig = bytes.Trim(jsonConfig, " \r\n\t")
			pbVal = &pb.TypedValue{
				Value: &pb.TypedValue_JsonIetfVal{
					JsonIetfVal: jsonConfig,
				},
			}
		} else {
			if strVal, err := strconv.Unquote(pathValuePair[1]); err == nil {
				pbVal = &pb.TypedValue{
					Value: &pb.TypedValue_StringVal{
						StringVal: strVal,
					},
				}
			} else {
				if intVal, err := strconv.ParseInt(pathValuePair[1], 10, 64); err == nil {
					pbVal = &pb.TypedValue{
						Value: &pb.TypedValue_IntVal{
							IntVal: intVal,
						},
					}
				} else if floatVal, err := strconv.ParseFloat(pathValuePair[1], 32); err == nil {
					pbVal = &pb.TypedValue{
						Value: &pb.TypedValue_FloatVal{
							FloatVal: float32(floatVal),
						},
					}
				} else if boolVal, err := strconv.ParseBool(pathValuePair[1]); err == nil {
					pbVal = &pb.TypedValue{
						Value: &pb.TypedValue_BoolVal{
							BoolVal: boolVal,
						},
					}
				} else {
					pbVal = &pb.TypedValue{
						Value: &pb.TypedValue_StringVal{
							StringVal: pathValuePair[1],
						},
					}
				}
			}
		}
		pbUpdateList = append(pbUpdateList, &pb.Update{Path: pbPath, Val: pbVal})
	}
	return pbUpdateList
}

func buildPbPrefix (pathTarget *string) *pb.Path {
	var prefix pb.Path
	prefix.Target = *pathTarget
	return &prefix
}

func main() {
	flag.Var(&deleteOpt, "delete", "xpath to be deleted.")
	flag.Var(&replaceOpt, "replace", "xpath:value pair to be replaced. Value can be numeric, boolean, string, or IETF JSON file (. starts with '@').")
	flag.Var(&updateOpt, "update", "xpath:value pair to be updated. Value can be numeric, boolean, string, or IETF JSON file (. starts with '@').")
	flag.Parse()

	opts := credentials.ClientCredentials(*targetName)
	conn, err := grpc.Dial(*targetAddr, opts...)
	if err != nil {
		log.Exitf("Dialing to %q failed: %v", *targetAddr, err)
	}
	defer conn.Close()

	var deleteList []*pb.Path
	for _, xPath := range deleteOpt {
		pbPath, err := xpath.ToGNMIPath(xPath)
		if err != nil {
			log.Exitf("error in parsing xpath %q to gnmi path", xPath)
		}
		deleteList = append(deleteList, pbPath)
	}

	replaceList := buildPbUpdateList(replaceOpt)
	updateList := buildPbUpdateList(updateOpt)
	prefix := buildPbPrefix(pathTarget)

	setRequest := &pb.SetRequest{
		Prefix:	prefix,
		Delete:  deleteList,
		Replace: replaceList,
		Update:  updateList,
	}

	fmt.Println("== setRequest:")
	utils.PrintProto(setRequest)

	cli := pb.NewGNMIClient(conn)
	setResponse, err := cli.Set(context.Background(), setRequest)
	if err != nil {
		log.Exitf("Set failed: %v", err)
	}

	fmt.Println("== getResponse:")
	utils.PrintProto(setResponse)
}

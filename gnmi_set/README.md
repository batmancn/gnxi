# gNMI Set

A simple shell binary that performs a SET against a gNMI Target.

## Install

```
go get github.com/google/gnxi/gnmi_set
go install github.com/google/gnxi/gnmi_set
```

## Run

Run gnmi\_set -help to see usage. For example:

```
gnmi_set \
  -delete /system/openflow/agent/config/max-backoff \
  -replace /system/clock:@clock-config.json \
  -replace /system/openflow/agent/config/max-backoff:12 \
  -update /system/clock/config/timezone-name:"US/New York" \
  -target_addr localhost:10161 \
  -target_name hostname.com \
  -key client.key \
  -cert client.crt \
  -ca ca.crt \
  -username foo \
  -password bar
  -alsologtostderr
```

```
./gnmi_set -xpath_target "MTNOS" -replace local-routes/static-routes/static:"*1" -target_addr 172.18.8.210:8080 -alsologtostderr -insecure true

./gnmi_set -xpath_target "MTNOS" -update /interfaces/interface[name=Ethernet16]/subinterfaces/subinterface/ipv4/addresses/address/config/ip:@clock-config.json -target_addr 172.18.8.241:8080 -alsologtostderr -insecure true
```

## bug

```
gyw@sonic107:~/go/src/github.com/google/gnxi/gnmi_set$ go build
# github.com/google/gnxi/gnmi_set
./gnmi_set.go:115:5: cannot use &"github.com/openconfig/ygot/proto/ywrapper".BoolValue literal (type *"github.com/openconfig/ygot/proto/ywrapper".BoolValue) as type *"github.com/Azure/sonic-telemetry/vendor/github.com/openconfig/ygot/proto/ywrapper".BoolValue in field value
./gnmi_set.go:130:5: cannot use &"github.com/openconfig/ygot/proto/ywrapper".BoolValue literal (type *"github.com/openconfig/ygot/proto/ywrapper".BoolValue) as type *"github.com/Azure/sonic-telemetry/vendor/github.com/openconfig/ygot/proto/ywrapper".BoolValue in field value
./gnmi_set.go:145:5: cannot use &"github.com/openconfig/ygot/proto/ywrapper".BoolValue literal (type *"github.com/openconfig/ygot/proto/ywrapper".BoolValue) as type *"github.com/Azure/sonic-telemetry/vendor/github.com/openconfig/ygot/proto/ywrapper".BoolValue in field value
./gnmi_set.go:160:5: cannot use &"github.com/openconfig/ygot/proto/ywrapper".BoolValue literal (type *"github.com/openconfig/ygot/proto/ywrapper".BoolValue) as type *"github.com/Azure/sonic-telemetry/vendor/github.com/openconfig/ygot/proto/ywrapper".BoolValue in field value
```

gyw@sonic107:~/go/src/github.com/Azure/sonic-telemetry$ pwd
/opt/gyw/go/src/github.com/Azure/sonic-telemetry
gyw@sonic107:~/go/src/github.com/Azure/sonic-telemetry$ mv vendor/ vendor-bak
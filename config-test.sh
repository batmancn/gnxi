#!/bin/bash

# ./mt_telemetry_bin --port 8080 --insecure --allow_no_client_auth --logtostderr
# gyw@sonic107:~/mos-code/mtnos-buildimage/src/sonic-telemetry/proto$ cd ~/go/src/github.com/google/gnxi/gnmi_get
# scp blogs/mos_code/mos-mtsonic-utilities/ansible/mtnos-telemetry/config-test.sh gyw@sonic107:/opt/gyw/go/src/github.com/google/gnxi

RES=""

run_cmd () {
    RES=""
    echo
    echo "$1"
    RES=$(eval "$1")
    echo "$RES"
}

main_test() {
    echo "--- /interfaces/interface[name=Ethernet1] ---"
    echo
    run_cmd "./gnmi_get/gnmi_get -xpath_target "MTNOS" -xpath /interfaces/interface[name=Ethernet1] -target_addr 172.18.8.214:8080 -alsologtostderr -insecure true"
    echo
    echo "--- check result ---"
    re=".*aa.Enabled: true.*"
    if [[ $RES =~ $re ]]; then
        echo "PASS"
    else
        echo "FAILD, $re not match"
    fi

    echo "--- /interfaces/interface[name=Ethernet1]/config/enabled ---"
    echo
    run_cmd "./gnmi_set/gnmi_set -xpath_target "MTNOS" -replace /interfaces/interface[name=Ethernet1]/config/enabled:"*5" -target_addr 172.18.8.214:8080 -alsologtostderr -insecure true"
    echo

    run_cmd "./gnmi_get/gnmi_get -xpath_target "MTNOS" -xpath /interfaces/interface[name=Ethernet1] -target_addr 172.18.8.214:8080 -alsologtostderr -insecure true"
    echo
    echo "--- check result ---"
    re=".*aa.Enabled: false.*"
    if [[ $RES =~ $re ]]; then
        echo "PASS"
    else
        echo "FAILD, $re not match"
    fi

    run_cmd "./gnmi_set/gnmi_set -xpath_target "MTNOS" -replace /interfaces/interface[name=Ethernet1]/config/enabled:"*4" -target_addr 172.18.8.214:8080 -alsologtostderr -insecure true"
    echo

    run_cmd "./gnmi_get/gnmi_get -xpath_target "MTNOS" -xpath /interfaces/interface[name=Ethernet1] -target_addr 172.18.8.214:8080 -alsologtostderr -insecure true"
    echo
    echo "--- check result ---"
    re=".*aa.Enabled: true.*"
    if [[ $RES =~ $re ]]; then
        echo "PASS"
    else
        echo "FAILD, $re not match"
    fi

    echo "--- /bgp ---"
    echo
    run_cmd "./gnmi_get/gnmi_get -xpath_target "MTNOS" -xpath /bgp -target_addr 172.18.8.214:8080 -alsologtostderr -insecure true"
    echo
    echo "--- check result ---"
    RES=$(echo "$RES" | grep 192.168.40.10)
    re="^BGP Entry:.*192.168.40.10.*neighbor.SessionState=6, neighbor.Enabled=true$"
    if [[ $RES =~ $re ]]; then
        echo "PASS"
    else
        echo "FAILD, $re not match"
    fi

    echo "--- /bgp/neighbors/neighbor[neighbor-address=192.168.40.10]/config ---"
    echo

    run_cmd "./gnmi_set/gnmi_set -xpath_target MTNOS -replace /bgp/neighbors/neighbor[neighbor-address=192.168.40.10]/config:"*3"  -target_addr 172.18.8.214:8080 -alsologtostderr -insecure true"
    echo
    sleep 5

    run_cmd "./gnmi_get/gnmi_get -xpath_target "MTNOS" -xpath /bgp -target_addr 172.18.8.214:8080 -alsologtostderr -insecure true"
    echo
    echo "--- check result ---"
    RES=$(echo "$RES" | grep 192.168.40.10)
    re="^BGP Entry:.*192.168.40.10.*neighbor.SessionState=1, neighbor.Enabled=false$"
    if [[ $RES =~ $re ]]; then
        echo "PASS"
    else
        echo "FAILD, $re not match"
    fi

    run_cmd "./gnmi_set/gnmi_set -xpath_target MTNOS -replace /bgp/neighbors/neighbor[neighbor-address=192.168.40.10]/config:"*2"  -target_addr 172.18.8.214:8080 -alsologtostderr -insecure true"
    echo
    sleep 15

    run_cmd "./gnmi_get/gnmi_get -xpath_target "MTNOS" -xpath /bgp -target_addr 172.18.8.214:8080 -alsologtostderr -insecure true"
    echo
    echo "--- check result ---"
    RES=$(echo "$RES" | grep 192.168.40.10)
    re="^BGP Entry:.*192.168.40.10.*neighbor.SessionState=6, neighbor.Enabled=true$"
    if [[ $RES =~ $re ]]; then
        echo "PASS"
    else
        echo "FAILD, $re not match"
    fi

    echo "--- /acl ---"
    echo
    run_cmd "./gnmi_get/gnmi_get -xpath_target "MTNOS" -xpath /acl -target_addr 172.18.8.214:8080 -alsologtostderr -insecure true"
}

for (( i = 0; i < 1; i++ )); do
    echo "Start: `date "+%Y-%m-%d_%H_%M_%S"`"
    main_test
    sleep 5
    echo "Stop: `date "+%Y-%m-%d_%H_%M_%S"`"
done

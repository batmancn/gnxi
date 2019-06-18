# /bin/bash

# 遇到error退出
set -o nounset

# 封装函数，例如log
log() {
    local prefix="[$(date +%Y/%m/%d\ %H:%M:%S)]: "
    echo "${prefix} $@" >&2
}

protoc --go_out=./ ./*.proto

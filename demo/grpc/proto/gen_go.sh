#!/usr/bin/env bash

TARGET="../"

if [ -n "$1" ]; then
    TARGET=$1
fi

# 排除掉 extra/src 目录
for file in `find . -path ./extra/src -prune -o -name '*.proto' -print`;
do
	echo $file
	protoc -I=extra/src:. --grpc-gateway_out=$TARGET --go_out=plugins=grpc,paths=import:$TARGET  $file
done
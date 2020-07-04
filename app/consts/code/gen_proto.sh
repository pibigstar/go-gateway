#!/usr/bin/env bash

TARGET="./"

if [ -n "$1" ]; then
    TARGET=$1
fi

for file in `find . -path ./dmeo -prune -o -name '*.proto' -print`;
do
	echo $file
	protoc --go_out=plugins=grpc,paths=import:$TARGET $file
done

cat >> code.pb.go <<EOF
// Code impl
func (e Error) Code() int {
	return int(e)
}
EOF
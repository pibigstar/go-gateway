// Code generated by protoc-gen-go. DO NOT EDIT.
// source: code.proto

package code

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Error int32

const (
	// 公共Code
	Error_Unknown       Error = 0
	Error_Success       Error = 200
	Error_NoPermission  Error = 403
	Error_Server_Error  Error = 500
	Error_ParamError    Error = 900
	Error_Token_Expired Error = 910
	// admin
	Error_User_Not_Exist Error = 10001
	Error_Password_Error Error = 10002
	Error_Not_Login      Error = 10003
)

var Error_name = map[int32]string{
	0:     "Unknown",
	200:   "Success",
	403:   "NoPermission",
	500:   "Server_Error",
	900:   "ParamError",
	910:   "Token_Expired",
	10001: "User_Not_Exist",
	10002: "Password_Error",
	10003: "Not_Login",
}

var Error_value = map[string]int32{
	"Unknown":        0,
	"Success":        200,
	"NoPermission":   403,
	"Server_Error":   500,
	"ParamError":     900,
	"Token_Expired":  910,
	"User_Not_Exist": 10001,
	"Password_Error": 10002,
	"Not_Login":      10003,
}

func (x Error) String() string {
	return proto.EnumName(Error_name, int32(x))
}

func (Error) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6e9b0151640170c3, []int{0}
}

func init() {
	proto.RegisterEnum("code.Error", Error_name, Error_value)
}

func init() {
	proto.RegisterFile("code.proto", fileDescriptor_6e9b0151640170c3)
}

var fileDescriptor_6e9b0151640170c3 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0xce, 0x41, 0x4e, 0x86, 0x40,
	0x0c, 0x05, 0x60, 0xff, 0xa0, 0x4e, 0xac, 0x88, 0x63, 0xbd, 0x85, 0x0b, 0x37, 0x9e, 0x81, 0x9d,
	0x99, 0x90, 0x20, 0xeb, 0x09, 0x42, 0x63, 0x26, 0x84, 0x29, 0x69, 0x51, 0x38, 0x80, 0x6b, 0x13,
	0xe5, 0x10, 0x5e, 0xc5, 0x03, 0x79, 0x00, 0x33, 0xf8, 0xef, 0x9a, 0x2f, 0x7d, 0x2f, 0x0f, 0xa0,
	0xe3, 0x9e, 0xee, 0x27, 0xe1, 0x99, 0xf1, 0x34, 0xdd, 0x77, 0xdf, 0x07, 0x38, 0x2b, 0x45, 0x58,
	0xf0, 0x12, 0x4c, 0x13, 0x87, 0xc8, 0x4b, 0xb4, 0x27, 0x98, 0x83, 0xa9, 0x5f, 0xbb, 0x8e, 0x54,
	0xed, 0xcf, 0x01, 0x6f, 0x20, 0x77, 0x5c, 0x91, 0x8c, 0x41, 0x35, 0x70, 0xb4, 0x5b, 0x96, 0xa8,
	0x26, 0x79, 0x23, 0xf1, 0x7b, 0xda, 0xfe, 0x66, 0x78, 0x0d, 0x50, 0xb5, 0xd2, 0x8e, 0xff, 0xf0,
	0x6e, 0x10, 0xe1, 0xea, 0x89, 0x07, 0x8a, 0xbe, 0x5c, 0xa7, 0x20, 0xd4, 0xdb, 0x0f, 0x83, 0xb7,
	0x50, 0x34, 0x4a, 0xe2, 0x1d, 0xcf, 0xbe, 0x5c, 0x83, 0xce, 0xf6, 0xd3, 0x25, 0xac, 0x5a, 0xd5,
	0x85, 0xa5, 0x3f, 0xd6, 0x7d, 0x39, 0x2c, 0xe0, 0x22, 0x3d, 0x3d, 0xf2, 0x4b, 0x88, 0x76, 0x73,
	0xcf, 0xe7, 0xfb, 0xec, 0x87, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x75, 0xf1, 0x6c, 0xc4,
	0x00, 0x00, 0x00,
}
// Code impl
func (e Error) Code() int {
	return int(e)
}

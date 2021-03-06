// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/keyword_plan_ad_group_error.proto

package errors

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

// Enum describing possible errors from applying a keyword plan ad group.
type KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError int32

const (
	// Enum unspecified.
	KeywordPlanAdGroupErrorEnum_UNSPECIFIED KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError = 0
	// The received error code is not known in this version.
	KeywordPlanAdGroupErrorEnum_UNKNOWN KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError = 1
	// The keyword plan ad group name is missing, empty, longer than allowed
	// limit or contains invalid chars.
	KeywordPlanAdGroupErrorEnum_INVALID_NAME KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError = 2
	// The keyword plan ad group name is duplicate to an existing keyword plan
	// AdGroup name or other keyword plan AdGroup name in the request.
	KeywordPlanAdGroupErrorEnum_DUPLICATE_NAME KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError = 3
)

var KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "INVALID_NAME",
	3: "DUPLICATE_NAME",
}

var KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError_value = map[string]int32{
	"UNSPECIFIED":    0,
	"UNKNOWN":        1,
	"INVALID_NAME":   2,
	"DUPLICATE_NAME": 3,
}

func (x KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError) String() string {
	return proto.EnumName(KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError_name, int32(x))
}

func (KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2bdfaf63ed9ec36a, []int{0, 0}
}

// Container for enum describing possible errors from applying a keyword plan
// ad group.
type KeywordPlanAdGroupErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeywordPlanAdGroupErrorEnum) Reset()         { *m = KeywordPlanAdGroupErrorEnum{} }
func (m *KeywordPlanAdGroupErrorEnum) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanAdGroupErrorEnum) ProtoMessage()    {}
func (*KeywordPlanAdGroupErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bdfaf63ed9ec36a, []int{0}
}

func (m *KeywordPlanAdGroupErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanAdGroupErrorEnum.Unmarshal(m, b)
}
func (m *KeywordPlanAdGroupErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanAdGroupErrorEnum.Marshal(b, m, deterministic)
}
func (m *KeywordPlanAdGroupErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanAdGroupErrorEnum.Merge(m, src)
}
func (m *KeywordPlanAdGroupErrorEnum) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanAdGroupErrorEnum.Size(m)
}
func (m *KeywordPlanAdGroupErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanAdGroupErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanAdGroupErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.errors.KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError", KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError_name, KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError_value)
	proto.RegisterType((*KeywordPlanAdGroupErrorEnum)(nil), "google.ads.googleads.v0.errors.KeywordPlanAdGroupErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/keyword_plan_ad_group_error.proto", fileDescriptor_2bdfaf63ed9ec36a)
}

var fileDescriptor_2bdfaf63ed9ec36a = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xc1, 0x4a, 0xc4, 0x30,
	0x14, 0x74, 0xbb, 0xa0, 0x90, 0x15, 0x2d, 0xb9, 0x78, 0x50, 0xf6, 0xd0, 0x0f, 0x48, 0x0b, 0xde,
	0xe2, 0xc5, 0xec, 0x36, 0x96, 0xb2, 0x6b, 0x2c, 0x68, 0x2b, 0x48, 0xa5, 0x44, 0x53, 0x82, 0xd8,
	0x6d, 0x4a, 0xe2, 0xae, 0x08, 0x7e, 0x8d, 0x47, 0x3f, 0xc5, 0x4f, 0xf1, 0x07, 0xbc, 0x4a, 0x1b,
	0xb7, 0xb7, 0x7a, 0xca, 0xf0, 0x66, 0xde, 0x64, 0xde, 0x80, 0x73, 0xa9, 0x94, 0xac, 0x4a, 0x9f,
	0x0b, 0xe3, 0x5b, 0xd8, 0xa2, 0x4d, 0xe0, 0x97, 0x5a, 0x2b, 0x6d, 0xfc, 0xe7, 0xf2, 0xed, 0x55,
	0x69, 0x51, 0x34, 0x15, 0xaf, 0x0b, 0x2e, 0x0a, 0xa9, 0xd5, 0xba, 0x29, 0x3a, 0x12, 0x35, 0x5a,
	0xbd, 0x28, 0x38, 0xb5, 0x6b, 0x88, 0x0b, 0x83, 0x7a, 0x07, 0xb4, 0x09, 0x90, 0x75, 0xf0, 0xde,
	0xc1, 0xf1, 0xc2, 0x9a, 0x24, 0x15, 0xaf, 0x89, 0x88, 0x5a, 0x07, 0xda, 0x72, 0xb4, 0x5e, 0xaf,
	0xbc, 0x7b, 0x70, 0x34, 0x40, 0xc3, 0x43, 0x30, 0x49, 0xd9, 0x75, 0x42, 0xe7, 0xf1, 0x45, 0x4c,
	0x43, 0x77, 0x07, 0x4e, 0xc0, 0x5e, 0xca, 0x16, 0xec, 0xea, 0x96, 0xb9, 0x23, 0xe8, 0x82, 0xfd,
	0x98, 0x65, 0x64, 0x19, 0x87, 0x05, 0x23, 0x97, 0xd4, 0x75, 0x20, 0x04, 0x07, 0x61, 0x9a, 0x2c,
	0xe3, 0x39, 0xb9, 0xa1, 0x76, 0x36, 0x9e, 0xfd, 0x8c, 0x80, 0xf7, 0xa8, 0x56, 0xe8, 0xff, 0x90,
	0xb3, 0x93, 0x81, 0x0c, 0x49, 0x7b, 0x62, 0x32, 0xba, 0x0b, 0xff, 0xf6, 0xa5, 0xaa, 0x78, 0x2d,
	0x91, 0xd2, 0xd2, 0x97, 0x65, 0xdd, 0x15, 0xb0, 0xad, 0xad, 0x79, 0x32, 0x43, 0x2d, 0x9e, 0xd9,
	0xe7, 0xc3, 0x19, 0x47, 0x84, 0x7c, 0x3a, 0xd3, 0xc8, 0x9a, 0x11, 0x61, 0x90, 0x85, 0x2d, 0xca,
	0x02, 0xd4, 0x7d, 0x69, 0xbe, 0xb6, 0x82, 0x9c, 0x08, 0x93, 0xf7, 0x82, 0x3c, 0x0b, 0x72, 0x2b,
	0xf8, 0x76, 0x3c, 0x3b, 0xc5, 0x98, 0x08, 0x83, 0x71, 0x2f, 0xc1, 0x38, 0x0b, 0x30, 0xb6, 0xa2,
	0x87, 0xdd, 0x2e, 0xdd, 0xe9, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb2, 0xc6, 0xaf, 0x56, 0xe2,
	0x01, 0x00, 0x00,
}

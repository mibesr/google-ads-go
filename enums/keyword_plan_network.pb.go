// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/keyword_plan_network.proto

package enums

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

// Enumerates keyword plan forecastable network types.
type KeywordPlanNetworkEnum_KeywordPlanNetwork int32

const (
	// Not specified.
	KeywordPlanNetworkEnum_UNSPECIFIED KeywordPlanNetworkEnum_KeywordPlanNetwork = 0
	// The value is unknown in this version.
	KeywordPlanNetworkEnum_UNKNOWN KeywordPlanNetworkEnum_KeywordPlanNetwork = 1
	// Google Search.
	KeywordPlanNetworkEnum_GOOGLE_SEARCH KeywordPlanNetworkEnum_KeywordPlanNetwork = 2
	// Google Search + Search partners.
	KeywordPlanNetworkEnum_GOOGLE_SEARCH_AND_PARTNERS KeywordPlanNetworkEnum_KeywordPlanNetwork = 3
)

var KeywordPlanNetworkEnum_KeywordPlanNetwork_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "GOOGLE_SEARCH",
	3: "GOOGLE_SEARCH_AND_PARTNERS",
}

var KeywordPlanNetworkEnum_KeywordPlanNetwork_value = map[string]int32{
	"UNSPECIFIED":                0,
	"UNKNOWN":                    1,
	"GOOGLE_SEARCH":              2,
	"GOOGLE_SEARCH_AND_PARTNERS": 3,
}

func (x KeywordPlanNetworkEnum_KeywordPlanNetwork) String() string {
	return proto.EnumName(KeywordPlanNetworkEnum_KeywordPlanNetwork_name, int32(x))
}

func (KeywordPlanNetworkEnum_KeywordPlanNetwork) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9df278a694673a9e, []int{0, 0}
}

// Container for enumeration of keyword plan forecastable network types.
type KeywordPlanNetworkEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeywordPlanNetworkEnum) Reset()         { *m = KeywordPlanNetworkEnum{} }
func (m *KeywordPlanNetworkEnum) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanNetworkEnum) ProtoMessage()    {}
func (*KeywordPlanNetworkEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_9df278a694673a9e, []int{0}
}

func (m *KeywordPlanNetworkEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanNetworkEnum.Unmarshal(m, b)
}
func (m *KeywordPlanNetworkEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanNetworkEnum.Marshal(b, m, deterministic)
}
func (m *KeywordPlanNetworkEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanNetworkEnum.Merge(m, src)
}
func (m *KeywordPlanNetworkEnum) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanNetworkEnum.Size(m)
}
func (m *KeywordPlanNetworkEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanNetworkEnum.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanNetworkEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.KeywordPlanNetworkEnum_KeywordPlanNetwork", KeywordPlanNetworkEnum_KeywordPlanNetwork_name, KeywordPlanNetworkEnum_KeywordPlanNetwork_value)
	proto.RegisterType((*KeywordPlanNetworkEnum)(nil), "google.ads.googleads.v0.enums.KeywordPlanNetworkEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/keyword_plan_network.proto", fileDescriptor_9df278a694673a9e)
}

var fileDescriptor_9df278a694673a9e = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x48, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0xec, 0xd4, 0xca, 0xf2, 0xfc, 0xa2, 0x94, 0xf8, 0x82, 0x9c, 0xc4, 0xbc,
	0xf8, 0xbc, 0xd4, 0x92, 0xf2, 0xfc, 0xa2, 0x6c, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x59,
	0x88, 0x72, 0xbd, 0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x4e, 0xbd, 0x32, 0x03, 0x3d, 0xb0, 0x4e, 0xa5,
	0x7a, 0x2e, 0x31, 0x6f, 0x88, 0xe6, 0x80, 0x9c, 0xc4, 0x3c, 0x3f, 0x88, 0x56, 0xd7, 0xbc, 0xd2,
	0x5c, 0xa5, 0x54, 0x2e, 0x21, 0x4c, 0x19, 0x21, 0x7e, 0x2e, 0xee, 0x50, 0xbf, 0xe0, 0x00, 0x57,
	0x67, 0x4f, 0x37, 0x4f, 0x57, 0x17, 0x01, 0x06, 0x21, 0x6e, 0x2e, 0xf6, 0x50, 0x3f, 0x6f, 0x3f,
	0xff, 0x70, 0x3f, 0x01, 0x46, 0x21, 0x41, 0x2e, 0x5e, 0x77, 0x7f, 0x7f, 0x77, 0x1f, 0xd7, 0xf8,
	0x60, 0x57, 0xc7, 0x20, 0x67, 0x0f, 0x01, 0x26, 0x21, 0x39, 0x2e, 0x29, 0x14, 0xa1, 0x78, 0x47,
	0x3f, 0x97, 0xf8, 0x00, 0xc7, 0xa0, 0x10, 0x3f, 0xd7, 0xa0, 0x60, 0x01, 0x66, 0xa7, 0x37, 0x8c,
	0x5c, 0x8a, 0xc9, 0xf9, 0xb9, 0x7a, 0x78, 0x9d, 0xe9, 0x24, 0x8e, 0xe9, 0x94, 0x00, 0x90, 0xf7,
	0x02, 0x18, 0xa3, 0x9c, 0xa0, 0x3a, 0xd3, 0xf3, 0x73, 0x12, 0xf3, 0xd2, 0xf5, 0xf2, 0x8b, 0xd2,
	0xf5, 0xd3, 0x53, 0xf3, 0xc0, 0x9e, 0x87, 0x05, 0x55, 0x41, 0x66, 0x31, 0x8e, 0x90, 0xb3, 0x06,
	0x93, 0x8b, 0x98, 0x98, 0xdd, 0x1d, 0x1d, 0x57, 0x31, 0xc9, 0xba, 0x43, 0x8c, 0x72, 0x4c, 0x29,
	0xd6, 0x83, 0x30, 0x41, 0xac, 0x30, 0x03, 0x3d, 0x50, 0x80, 0x14, 0x9f, 0x82, 0xc9, 0xc7, 0x38,
	0xa6, 0x14, 0xc7, 0xc0, 0xe5, 0x63, 0xc2, 0x0c, 0x62, 0xc0, 0xf2, 0xaf, 0x98, 0x14, 0x21, 0x82,
	0x56, 0x56, 0x8e, 0x29, 0xc5, 0x56, 0x56, 0x70, 0x15, 0x56, 0x56, 0x61, 0x06, 0x56, 0x56, 0x60,
	0x35, 0x49, 0x6c, 0x60, 0x87, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2c, 0xdf, 0xdb, 0x28,
	0xd1, 0x01, 0x00, 0x00,
}

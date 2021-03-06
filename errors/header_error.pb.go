// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/header_error.proto

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

// Enum describing possible header errors.
type HeaderErrorEnum_HeaderError int32

const (
	// Enum unspecified.
	HeaderErrorEnum_UNSPECIFIED HeaderErrorEnum_HeaderError = 0
	// The received error code is not known in this version.
	HeaderErrorEnum_UNKNOWN HeaderErrorEnum_HeaderError = 1
	// The login customer id could not be validated.
	HeaderErrorEnum_INVALID_LOGIN_CUSTOMER_ID HeaderErrorEnum_HeaderError = 3
)

var HeaderErrorEnum_HeaderError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	3: "INVALID_LOGIN_CUSTOMER_ID",
}

var HeaderErrorEnum_HeaderError_value = map[string]int32{
	"UNSPECIFIED":               0,
	"UNKNOWN":                   1,
	"INVALID_LOGIN_CUSTOMER_ID": 3,
}

func (x HeaderErrorEnum_HeaderError) String() string {
	return proto.EnumName(HeaderErrorEnum_HeaderError_name, int32(x))
}

func (HeaderErrorEnum_HeaderError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_418ddd4782dda673, []int{0, 0}
}

// Container for enum describing possible header errors.
type HeaderErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeaderErrorEnum) Reset()         { *m = HeaderErrorEnum{} }
func (m *HeaderErrorEnum) String() string { return proto.CompactTextString(m) }
func (*HeaderErrorEnum) ProtoMessage()    {}
func (*HeaderErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_418ddd4782dda673, []int{0}
}

func (m *HeaderErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderErrorEnum.Unmarshal(m, b)
}
func (m *HeaderErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderErrorEnum.Marshal(b, m, deterministic)
}
func (m *HeaderErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderErrorEnum.Merge(m, src)
}
func (m *HeaderErrorEnum) XXX_Size() int {
	return xxx_messageInfo_HeaderErrorEnum.Size(m)
}
func (m *HeaderErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.errors.HeaderErrorEnum_HeaderError", HeaderErrorEnum_HeaderError_name, HeaderErrorEnum_HeaderError_value)
	proto.RegisterType((*HeaderErrorEnum)(nil), "google.ads.googleads.v0.errors.HeaderErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/header_error.proto", fileDescriptor_418ddd4782dda673)
}

var fileDescriptor_418ddd4782dda673 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4c, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xa2,
	0xa2, 0xfc, 0xa2, 0x62, 0xfd, 0x8c, 0xd4, 0xc4, 0x94, 0xd4, 0xa2, 0x78, 0x30, 0x4f, 0xaf, 0xa0,
	0x28, 0xbf, 0x24, 0x5f, 0x48, 0x0e, 0xa2, 0x4e, 0x2f, 0x31, 0xa5, 0x58, 0x0f, 0xae, 0x45, 0xaf,
	0xcc, 0x40, 0x0f, 0xa2, 0x45, 0x29, 0x96, 0x8b, 0xdf, 0x03, 0xac, 0xcb, 0x15, 0xc4, 0x77, 0xcd,
	0x2b, 0xcd, 0x55, 0xf2, 0xe2, 0xe2, 0x46, 0x12, 0x12, 0xe2, 0xe7, 0xe2, 0x0e, 0xf5, 0x0b, 0x0e,
	0x70, 0x75, 0xf6, 0x74, 0xf3, 0x74, 0x75, 0x11, 0x60, 0x10, 0xe2, 0xe6, 0x62, 0x0f, 0xf5, 0xf3,
	0xf6, 0xf3, 0x0f, 0xf7, 0x13, 0x60, 0x14, 0x92, 0xe5, 0x92, 0xf4, 0xf4, 0x0b, 0x73, 0xf4, 0xf1,
	0x74, 0x89, 0xf7, 0xf1, 0x77, 0xf7, 0xf4, 0x8b, 0x77, 0x0e, 0x0d, 0x0e, 0xf1, 0xf7, 0x75, 0x0d,
	0x8a, 0xf7, 0x74, 0x11, 0x60, 0x76, 0x7a, 0xcd, 0xc8, 0xa5, 0x94, 0x9c, 0x9f, 0xab, 0x87, 0xdf,
	0x15, 0x4e, 0x02, 0x48, 0x16, 0x06, 0x80, 0xdc, 0x1d, 0xc0, 0x18, 0xe5, 0x02, 0xd5, 0x93, 0x9e,
	0x9f, 0x93, 0x98, 0x97, 0xae, 0x97, 0x5f, 0x94, 0xae, 0x9f, 0x9e, 0x9a, 0x07, 0xf6, 0x15, 0xcc,
	0xf3, 0x05, 0x99, 0xc5, 0xb8, 0xc2, 0xc2, 0x1a, 0x42, 0x2d, 0x62, 0x62, 0x76, 0x77, 0x74, 0x5c,
	0xc5, 0x24, 0xe7, 0x0e, 0x31, 0xcc, 0x31, 0xa5, 0x58, 0x0f, 0xc2, 0x04, 0xb1, 0xc2, 0x0c, 0xf4,
	0xc0, 0x56, 0x16, 0x9f, 0x82, 0x29, 0x88, 0x71, 0x4c, 0x29, 0x8e, 0x81, 0x2b, 0x88, 0x09, 0x33,
	0x88, 0x81, 0x28, 0x78, 0xc5, 0xa4, 0x04, 0x11, 0xb5, 0xb2, 0x72, 0x4c, 0x29, 0xb6, 0xb2, 0x82,
	0x2b, 0xb1, 0xb2, 0x0a, 0x33, 0xb0, 0xb2, 0x82, 0x28, 0x4a, 0x62, 0x03, 0xbb, 0xce, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0xe8, 0xc1, 0x33, 0x33, 0xa8, 0x01, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/shared_set_status.proto

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

// Enum listing the possible shared set statuses.
type SharedSetStatusEnum_SharedSetStatus int32

const (
	// Not specified.
	SharedSetStatusEnum_UNSPECIFIED SharedSetStatusEnum_SharedSetStatus = 0
	// Used for return value only. Represents value unknown in this version.
	SharedSetStatusEnum_UNKNOWN SharedSetStatusEnum_SharedSetStatus = 1
	// The shared set is enabled.
	SharedSetStatusEnum_ENABLED SharedSetStatusEnum_SharedSetStatus = 2
	// The shared set is removed and can no longer be used.
	SharedSetStatusEnum_REMOVED SharedSetStatusEnum_SharedSetStatus = 3
)

var SharedSetStatusEnum_SharedSetStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "REMOVED",
}

var SharedSetStatusEnum_SharedSetStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"REMOVED":     3,
}

func (x SharedSetStatusEnum_SharedSetStatus) String() string {
	return proto.EnumName(SharedSetStatusEnum_SharedSetStatus_name, int32(x))
}

func (SharedSetStatusEnum_SharedSetStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3337f4109e961f02, []int{0, 0}
}

// Container for enum describing types of shared set statuses.
type SharedSetStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SharedSetStatusEnum) Reset()         { *m = SharedSetStatusEnum{} }
func (m *SharedSetStatusEnum) String() string { return proto.CompactTextString(m) }
func (*SharedSetStatusEnum) ProtoMessage()    {}
func (*SharedSetStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_3337f4109e961f02, []int{0}
}

func (m *SharedSetStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharedSetStatusEnum.Unmarshal(m, b)
}
func (m *SharedSetStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharedSetStatusEnum.Marshal(b, m, deterministic)
}
func (m *SharedSetStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharedSetStatusEnum.Merge(m, src)
}
func (m *SharedSetStatusEnum) XXX_Size() int {
	return xxx_messageInfo_SharedSetStatusEnum.Size(m)
}
func (m *SharedSetStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_SharedSetStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_SharedSetStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.SharedSetStatusEnum_SharedSetStatus", SharedSetStatusEnum_SharedSetStatus_name, SharedSetStatusEnum_SharedSetStatus_value)
	proto.RegisterType((*SharedSetStatusEnum)(nil), "google.ads.googleads.v0.enums.SharedSetStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/shared_set_status.proto", fileDescriptor_3337f4109e961f02)
}

var fileDescriptor_3337f4109e961f02 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4d, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0xe2, 0x8c, 0xc4, 0xa2, 0xd4, 0x94, 0xf8, 0xe2, 0xd4, 0x92, 0xf8, 0xe2,
	0x92, 0xc4, 0x92, 0xd2, 0x62, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x59, 0x88, 0x5a, 0xbd,
	0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x36, 0xbd, 0x32, 0x03, 0x3d, 0xb0, 0x36, 0xa5, 0x04, 0x2e, 0xe1,
	0x60, 0xb0, 0xce, 0xe0, 0xd4, 0x92, 0x60, 0xb0, 0x3e, 0xd7, 0xbc, 0xd2, 0x5c, 0x25, 0x4f, 0x2e,
	0x7e, 0x34, 0x61, 0x21, 0x7e, 0x2e, 0xee, 0x50, 0xbf, 0xe0, 0x00, 0x57, 0x67, 0x4f, 0x37, 0x4f,
	0x57, 0x17, 0x01, 0x06, 0x21, 0x6e, 0x2e, 0xf6, 0x50, 0x3f, 0x6f, 0x3f, 0xff, 0x70, 0x3f, 0x01,
	0x46, 0x10, 0xc7, 0xd5, 0xcf, 0xd1, 0xc9, 0xc7, 0xd5, 0x45, 0x80, 0x09, 0xc4, 0x09, 0x72, 0xf5,
	0xf5, 0x0f, 0x73, 0x75, 0x11, 0x60, 0x76, 0x3a, 0xca, 0xc8, 0xa5, 0x98, 0x9c, 0x9f, 0xab, 0x87,
	0xd7, 0x1d, 0x4e, 0x22, 0x68, 0xd6, 0x05, 0x80, 0x1c, 0x1f, 0xc0, 0x18, 0xe5, 0x04, 0xd5, 0x96,
	0x9e, 0x9f, 0x93, 0x98, 0x97, 0xae, 0x97, 0x5f, 0x94, 0xae, 0x9f, 0x9e, 0x9a, 0x07, 0xf6, 0x1a,
	0x2c, 0x14, 0x0a, 0x32, 0x8b, 0x71, 0x04, 0x8a, 0x35, 0x98, 0x5c, 0xc4, 0xc4, 0xec, 0xee, 0xe8,
	0xb8, 0x8a, 0x49, 0xd6, 0x1d, 0x62, 0x94, 0x63, 0x4a, 0xb1, 0x1e, 0x84, 0x09, 0x62, 0x85, 0x19,
	0xe8, 0x81, 0x7c, 0x5c, 0x7c, 0x0a, 0x26, 0x1f, 0xe3, 0x98, 0x52, 0x1c, 0x03, 0x97, 0x8f, 0x09,
	0x33, 0x88, 0x01, 0xcb, 0x27, 0xb1, 0x81, 0x2d, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x74,
	0x72, 0x4c, 0x3e, 0x88, 0x01, 0x00, 0x00,
}
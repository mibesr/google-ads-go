// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/policy_violation_error.proto

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

// Enum describing possible policy violation errors.
type PolicyViolationErrorEnum_PolicyViolationError int32

const (
	// Enum unspecified.
	PolicyViolationErrorEnum_UNSPECIFIED PolicyViolationErrorEnum_PolicyViolationError = 0
	// The received error code is not known in this version.
	PolicyViolationErrorEnum_UNKNOWN PolicyViolationErrorEnum_PolicyViolationError = 1
	// A policy was violated. See PolicyViolationDetails for more detail.
	PolicyViolationErrorEnum_POLICY_ERROR PolicyViolationErrorEnum_PolicyViolationError = 2
)

var PolicyViolationErrorEnum_PolicyViolationError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "POLICY_ERROR",
}

var PolicyViolationErrorEnum_PolicyViolationError_value = map[string]int32{
	"UNSPECIFIED":  0,
	"UNKNOWN":      1,
	"POLICY_ERROR": 2,
}

func (x PolicyViolationErrorEnum_PolicyViolationError) String() string {
	return proto.EnumName(PolicyViolationErrorEnum_PolicyViolationError_name, int32(x))
}

func (PolicyViolationErrorEnum_PolicyViolationError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a4b5e7134ee70800, []int{0, 0}
}

// Container for enum describing possible policy violation errors.
type PolicyViolationErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolicyViolationErrorEnum) Reset()         { *m = PolicyViolationErrorEnum{} }
func (m *PolicyViolationErrorEnum) String() string { return proto.CompactTextString(m) }
func (*PolicyViolationErrorEnum) ProtoMessage()    {}
func (*PolicyViolationErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b5e7134ee70800, []int{0}
}

func (m *PolicyViolationErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyViolationErrorEnum.Unmarshal(m, b)
}
func (m *PolicyViolationErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyViolationErrorEnum.Marshal(b, m, deterministic)
}
func (m *PolicyViolationErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyViolationErrorEnum.Merge(m, src)
}
func (m *PolicyViolationErrorEnum) XXX_Size() int {
	return xxx_messageInfo_PolicyViolationErrorEnum.Size(m)
}
func (m *PolicyViolationErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyViolationErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyViolationErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.errors.PolicyViolationErrorEnum_PolicyViolationError", PolicyViolationErrorEnum_PolicyViolationError_name, PolicyViolationErrorEnum_PolicyViolationError_value)
	proto.RegisterType((*PolicyViolationErrorEnum)(nil), "google.ads.googleads.v0.errors.PolicyViolationErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/policy_violation_error.proto", fileDescriptor_a4b5e7134ee70800)
}

var fileDescriptor_a4b5e7134ee70800 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4e, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xa2,
	0xa2, 0xfc, 0xa2, 0x62, 0xfd, 0x82, 0xfc, 0x9c, 0xcc, 0xe4, 0xca, 0xf8, 0xb2, 0xcc, 0xfc, 0x9c,
	0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x78, 0xb0, 0xb8, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x1c,
	0x44, 0x87, 0x5e, 0x62, 0x4a, 0xb1, 0x1e, 0x5c, 0xb3, 0x5e, 0x99, 0x81, 0x1e, 0x44, 0xb3, 0x52,
	0x12, 0x97, 0x44, 0x00, 0x58, 0x7f, 0x18, 0x4c, 0xbb, 0x2b, 0x48, 0xc2, 0x35, 0xaf, 0x34, 0x57,
	0xc9, 0x8d, 0x4b, 0x04, 0x9b, 0x9c, 0x10, 0x3f, 0x17, 0x77, 0xa8, 0x5f, 0x70, 0x80, 0xab, 0xb3,
	0xa7, 0x9b, 0xa7, 0xab, 0x8b, 0x00, 0x83, 0x10, 0x37, 0x17, 0x7b, 0xa8, 0x9f, 0xb7, 0x9f, 0x7f,
	0xb8, 0x9f, 0x00, 0xa3, 0x90, 0x00, 0x17, 0x4f, 0x80, 0xbf, 0x8f, 0xa7, 0x73, 0x64, 0xbc, 0x6b,
	0x50, 0x90, 0x7f, 0x90, 0x00, 0x93, 0xd3, 0x17, 0x46, 0x2e, 0xa5, 0xe4, 0xfc, 0x5c, 0x3d, 0xfc,
	0x4e, 0x71, 0x92, 0xc4, 0x66, 0x59, 0x00, 0xc8, 0x17, 0x01, 0x8c, 0x51, 0x2e, 0x50, 0xcd, 0xe9,
	0xf9, 0x39, 0x89, 0x79, 0xe9, 0x7a, 0xf9, 0x45, 0xe9, 0xfa, 0xe9, 0xa9, 0x79, 0x60, 0x3f, 0xc2,
	0x02, 0xa5, 0x20, 0xb3, 0x18, 0x57, 0x18, 0x59, 0x43, 0xa8, 0x45, 0x4c, 0xcc, 0xee, 0x8e, 0x8e,
	0xab, 0x98, 0xe4, 0xdc, 0x21, 0x86, 0x39, 0xa6, 0x14, 0xeb, 0x41, 0x98, 0x20, 0x56, 0x98, 0x81,
	0x1e, 0xd8, 0xca, 0xe2, 0x53, 0x30, 0x05, 0x31, 0x8e, 0x29, 0xc5, 0x31, 0x70, 0x05, 0x31, 0x61,
	0x06, 0x31, 0x10, 0x05, 0xaf, 0x98, 0x94, 0x20, 0xa2, 0x56, 0x56, 0x8e, 0x29, 0xc5, 0x56, 0x56,
	0x70, 0x25, 0x56, 0x56, 0x61, 0x06, 0x56, 0x56, 0x10, 0x45, 0x49, 0x6c, 0x60, 0xd7, 0x19, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x11, 0x3e, 0x4e, 0x3d, 0xc0, 0x01, 0x00, 0x00,
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/billing_setup_status.proto

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

// The possible statuses of a BillingSetup.
type BillingSetupStatusEnum_BillingSetupStatus int32

const (
	// Not specified.
	BillingSetupStatusEnum_UNSPECIFIED BillingSetupStatusEnum_BillingSetupStatus = 0
	// Used for return value only. Represents value unknown in this version.
	BillingSetupStatusEnum_UNKNOWN BillingSetupStatusEnum_BillingSetupStatus = 1
	// The billing setup is pending approval.
	BillingSetupStatusEnum_PENDING BillingSetupStatusEnum_BillingSetupStatus = 2
	// The billing setup has been approved but the corresponding first budget
	// has not.  This can only occur for billing setups configured for monthly
	// invoicing.
	BillingSetupStatusEnum_APPROVED_HELD BillingSetupStatusEnum_BillingSetupStatus = 3
	// The billing setup has been approved.
	BillingSetupStatusEnum_APPROVED BillingSetupStatusEnum_BillingSetupStatus = 4
	// The billing setup was cancelled by the user prior to approval.
	BillingSetupStatusEnum_CANCELLED BillingSetupStatusEnum_BillingSetupStatus = 5
)

var BillingSetupStatusEnum_BillingSetupStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "PENDING",
	3: "APPROVED_HELD",
	4: "APPROVED",
	5: "CANCELLED",
}

var BillingSetupStatusEnum_BillingSetupStatus_value = map[string]int32{
	"UNSPECIFIED":   0,
	"UNKNOWN":       1,
	"PENDING":       2,
	"APPROVED_HELD": 3,
	"APPROVED":      4,
	"CANCELLED":     5,
}

func (x BillingSetupStatusEnum_BillingSetupStatus) String() string {
	return proto.EnumName(BillingSetupStatusEnum_BillingSetupStatus_name, int32(x))
}

func (BillingSetupStatusEnum_BillingSetupStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f14fb11741a90b4, []int{0, 0}
}

// Message describing BillingSetup statuses.
type BillingSetupStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BillingSetupStatusEnum) Reset()         { *m = BillingSetupStatusEnum{} }
func (m *BillingSetupStatusEnum) String() string { return proto.CompactTextString(m) }
func (*BillingSetupStatusEnum) ProtoMessage()    {}
func (*BillingSetupStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f14fb11741a90b4, []int{0}
}

func (m *BillingSetupStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BillingSetupStatusEnum.Unmarshal(m, b)
}
func (m *BillingSetupStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BillingSetupStatusEnum.Marshal(b, m, deterministic)
}
func (m *BillingSetupStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BillingSetupStatusEnum.Merge(m, src)
}
func (m *BillingSetupStatusEnum) XXX_Size() int {
	return xxx_messageInfo_BillingSetupStatusEnum.Size(m)
}
func (m *BillingSetupStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_BillingSetupStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_BillingSetupStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.BillingSetupStatusEnum_BillingSetupStatus", BillingSetupStatusEnum_BillingSetupStatus_name, BillingSetupStatusEnum_BillingSetupStatus_value)
	proto.RegisterType((*BillingSetupStatusEnum)(nil), "google.ads.googleads.v0.enums.BillingSetupStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/billing_setup_status.proto", fileDescriptor_8f14fb11741a90b4)
}

var fileDescriptor_8f14fb11741a90b4 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xdf, 0x4a, 0xc3, 0x30,
	0x18, 0xc5, 0x6d, 0xe7, 0xdf, 0xcc, 0x61, 0xcd, 0x85, 0x5e, 0xed, 0x62, 0x7b, 0x80, 0xb4, 0xe0,
	0x8d, 0xc4, 0xab, 0x74, 0x8d, 0x73, 0x38, 0xb2, 0xe0, 0x58, 0x05, 0x29, 0x8c, 0xce, 0x96, 0x30,
	0xe8, 0x9a, 0xb1, 0xb4, 0x7b, 0x07, 0x5f, 0xc3, 0x4b, 0x1f, 0xc5, 0x47, 0x11, 0x1f, 0x42, 0x92,
	0xb8, 0xde, 0x0c, 0xbd, 0x09, 0x27, 0xdf, 0xf9, 0x7e, 0xe1, 0xe4, 0x80, 0x5b, 0x21, 0xa5, 0x28,
	0x72, 0x3f, 0xcd, 0x94, 0x6f, 0xa5, 0x56, 0xdb, 0xc0, 0xcf, 0xcb, 0x7a, 0xa5, 0xfc, 0xc5, 0xb2,
	0x28, 0x96, 0xa5, 0x98, 0xab, 0xbc, 0xaa, 0xd7, 0x73, 0x55, 0xa5, 0x55, 0xad, 0xd0, 0x7a, 0x23,
	0x2b, 0x09, 0xbb, 0x76, 0x1d, 0xa5, 0x99, 0x42, 0x0d, 0x89, 0xb6, 0x01, 0x32, 0x64, 0xff, 0xcd,
	0x01, 0x57, 0xa1, 0xa5, 0xa7, 0x1a, 0x9e, 0x1a, 0x96, 0x96, 0xf5, 0xaa, 0x2f, 0x01, 0xdc, 0x77,
	0xe0, 0x05, 0x68, 0xcf, 0xd8, 0x94, 0xd3, 0xc1, 0xe8, 0x7e, 0x44, 0x23, 0xef, 0x00, 0xb6, 0xc1,
	0xc9, 0x8c, 0x3d, 0xb2, 0xc9, 0x33, 0xf3, 0x1c, 0x7d, 0xe1, 0x94, 0x45, 0x23, 0x36, 0xf4, 0x5c,
	0x78, 0x09, 0x3a, 0x84, 0xf3, 0xa7, 0x49, 0x4c, 0xa3, 0xf9, 0x03, 0x1d, 0x47, 0x5e, 0x0b, 0x9e,
	0x83, 0xd3, 0xdd, 0xc8, 0x3b, 0x84, 0x1d, 0x70, 0x36, 0x20, 0x6c, 0x40, 0xc7, 0x63, 0x1a, 0x79,
	0x47, 0xe1, 0xb7, 0x03, 0x7a, 0xaf, 0x72, 0x85, 0xfe, 0x4d, 0x1c, 0x5e, 0xef, 0x87, 0xe2, 0xfa,
	0xa7, 0xdc, 0x79, 0x09, 0x7f, 0x49, 0x21, 0x8b, 0xb4, 0x14, 0x48, 0x6e, 0x84, 0x2f, 0xf2, 0xd2,
	0xf4, 0xb0, 0x6b, 0x6d, 0xbd, 0x54, 0x7f, 0x94, 0x78, 0x67, 0xce, 0x77, 0xb7, 0x35, 0x24, 0xe4,
	0xc3, 0xed, 0x0e, 0xed, 0x53, 0x24, 0x53, 0xc8, 0x4a, 0xad, 0xe2, 0x00, 0xe9, 0x6a, 0xd4, 0xe7,
	0xce, 0x4f, 0x48, 0xa6, 0x92, 0xc6, 0x4f, 0xe2, 0x20, 0x31, 0xfe, 0x97, 0xdb, 0xb3, 0x43, 0x8c,
	0x49, 0xa6, 0x30, 0x6e, 0x36, 0x30, 0x8e, 0x03, 0x8c, 0xcd, 0xce, 0xe2, 0xd8, 0x04, 0xbb, 0xf9,
	0x09, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x6a, 0xf6, 0xe9, 0xdc, 0x01, 0x00, 0x00,
}

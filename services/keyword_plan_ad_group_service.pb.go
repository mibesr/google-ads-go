// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/keyword_plan_ad_group_service.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	resources "github.com/kritzware/google-ads-go/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
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

// Request message for [KeywordPlanAdGroupService.GetKeywordPlanAdGroup][google.ads.googleads.v0.services.KeywordPlanAdGroupService.GetKeywordPlanAdGroup].
type GetKeywordPlanAdGroupRequest struct {
	// The resource name of the Keyword Plan ad group to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetKeywordPlanAdGroupRequest) Reset()         { *m = GetKeywordPlanAdGroupRequest{} }
func (m *GetKeywordPlanAdGroupRequest) String() string { return proto.CompactTextString(m) }
func (*GetKeywordPlanAdGroupRequest) ProtoMessage()    {}
func (*GetKeywordPlanAdGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_06d429e8eeffcde4, []int{0}
}

func (m *GetKeywordPlanAdGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetKeywordPlanAdGroupRequest.Unmarshal(m, b)
}
func (m *GetKeywordPlanAdGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetKeywordPlanAdGroupRequest.Marshal(b, m, deterministic)
}
func (m *GetKeywordPlanAdGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetKeywordPlanAdGroupRequest.Merge(m, src)
}
func (m *GetKeywordPlanAdGroupRequest) XXX_Size() int {
	return xxx_messageInfo_GetKeywordPlanAdGroupRequest.Size(m)
}
func (m *GetKeywordPlanAdGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetKeywordPlanAdGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetKeywordPlanAdGroupRequest proto.InternalMessageInfo

func (m *GetKeywordPlanAdGroupRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [KeywordPlanAdGroupService.MutateKeywordPlanAdGroups][google.ads.googleads.v0.services.KeywordPlanAdGroupService.MutateKeywordPlanAdGroups].
type MutateKeywordPlanAdGroupsRequest struct {
	// The ID of the customer whose Keyword Plan ad groups are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual Keyword Plan ad groups.
	Operations           []*KeywordPlanAdGroupOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *MutateKeywordPlanAdGroupsRequest) Reset()         { *m = MutateKeywordPlanAdGroupsRequest{} }
func (m *MutateKeywordPlanAdGroupsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateKeywordPlanAdGroupsRequest) ProtoMessage()    {}
func (*MutateKeywordPlanAdGroupsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_06d429e8eeffcde4, []int{1}
}

func (m *MutateKeywordPlanAdGroupsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateKeywordPlanAdGroupsRequest.Unmarshal(m, b)
}
func (m *MutateKeywordPlanAdGroupsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateKeywordPlanAdGroupsRequest.Marshal(b, m, deterministic)
}
func (m *MutateKeywordPlanAdGroupsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateKeywordPlanAdGroupsRequest.Merge(m, src)
}
func (m *MutateKeywordPlanAdGroupsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateKeywordPlanAdGroupsRequest.Size(m)
}
func (m *MutateKeywordPlanAdGroupsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateKeywordPlanAdGroupsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateKeywordPlanAdGroupsRequest proto.InternalMessageInfo

func (m *MutateKeywordPlanAdGroupsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateKeywordPlanAdGroupsRequest) GetOperations() []*KeywordPlanAdGroupOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

// A single operation (create, update, remove) on a Keyword Plan ad group.
type KeywordPlanAdGroupOperation struct {
	// The FieldMask that determines which resource fields are modified in an
	// update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*KeywordPlanAdGroupOperation_Create
	//	*KeywordPlanAdGroupOperation_Update
	//	*KeywordPlanAdGroupOperation_Remove
	Operation            isKeywordPlanAdGroupOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *KeywordPlanAdGroupOperation) Reset()         { *m = KeywordPlanAdGroupOperation{} }
func (m *KeywordPlanAdGroupOperation) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanAdGroupOperation) ProtoMessage()    {}
func (*KeywordPlanAdGroupOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_06d429e8eeffcde4, []int{2}
}

func (m *KeywordPlanAdGroupOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanAdGroupOperation.Unmarshal(m, b)
}
func (m *KeywordPlanAdGroupOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanAdGroupOperation.Marshal(b, m, deterministic)
}
func (m *KeywordPlanAdGroupOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanAdGroupOperation.Merge(m, src)
}
func (m *KeywordPlanAdGroupOperation) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanAdGroupOperation.Size(m)
}
func (m *KeywordPlanAdGroupOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanAdGroupOperation.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanAdGroupOperation proto.InternalMessageInfo

func (m *KeywordPlanAdGroupOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isKeywordPlanAdGroupOperation_Operation interface {
	isKeywordPlanAdGroupOperation_Operation()
}

type KeywordPlanAdGroupOperation_Create struct {
	Create *resources.KeywordPlanAdGroup `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type KeywordPlanAdGroupOperation_Update struct {
	Update *resources.KeywordPlanAdGroup `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type KeywordPlanAdGroupOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*KeywordPlanAdGroupOperation_Create) isKeywordPlanAdGroupOperation_Operation() {}

func (*KeywordPlanAdGroupOperation_Update) isKeywordPlanAdGroupOperation_Operation() {}

func (*KeywordPlanAdGroupOperation_Remove) isKeywordPlanAdGroupOperation_Operation() {}

func (m *KeywordPlanAdGroupOperation) GetOperation() isKeywordPlanAdGroupOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *KeywordPlanAdGroupOperation) GetCreate() *resources.KeywordPlanAdGroup {
	if x, ok := m.GetOperation().(*KeywordPlanAdGroupOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *KeywordPlanAdGroupOperation) GetUpdate() *resources.KeywordPlanAdGroup {
	if x, ok := m.GetOperation().(*KeywordPlanAdGroupOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *KeywordPlanAdGroupOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*KeywordPlanAdGroupOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*KeywordPlanAdGroupOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*KeywordPlanAdGroupOperation_Create)(nil),
		(*KeywordPlanAdGroupOperation_Update)(nil),
		(*KeywordPlanAdGroupOperation_Remove)(nil),
	}
}

// Response message for a Keyword Plan ad group mutate.
type MutateKeywordPlanAdGroupsResponse struct {
	// All results for the mutate.
	Results              []*MutateKeywordPlanAdGroupResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *MutateKeywordPlanAdGroupsResponse) Reset()         { *m = MutateKeywordPlanAdGroupsResponse{} }
func (m *MutateKeywordPlanAdGroupsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateKeywordPlanAdGroupsResponse) ProtoMessage()    {}
func (*MutateKeywordPlanAdGroupsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_06d429e8eeffcde4, []int{3}
}

func (m *MutateKeywordPlanAdGroupsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateKeywordPlanAdGroupsResponse.Unmarshal(m, b)
}
func (m *MutateKeywordPlanAdGroupsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateKeywordPlanAdGroupsResponse.Marshal(b, m, deterministic)
}
func (m *MutateKeywordPlanAdGroupsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateKeywordPlanAdGroupsResponse.Merge(m, src)
}
func (m *MutateKeywordPlanAdGroupsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateKeywordPlanAdGroupsResponse.Size(m)
}
func (m *MutateKeywordPlanAdGroupsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateKeywordPlanAdGroupsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateKeywordPlanAdGroupsResponse proto.InternalMessageInfo

func (m *MutateKeywordPlanAdGroupsResponse) GetResults() []*MutateKeywordPlanAdGroupResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the Keyword Plan ad group mutate.
type MutateKeywordPlanAdGroupResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateKeywordPlanAdGroupResult) Reset()         { *m = MutateKeywordPlanAdGroupResult{} }
func (m *MutateKeywordPlanAdGroupResult) String() string { return proto.CompactTextString(m) }
func (*MutateKeywordPlanAdGroupResult) ProtoMessage()    {}
func (*MutateKeywordPlanAdGroupResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_06d429e8eeffcde4, []int{4}
}

func (m *MutateKeywordPlanAdGroupResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateKeywordPlanAdGroupResult.Unmarshal(m, b)
}
func (m *MutateKeywordPlanAdGroupResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateKeywordPlanAdGroupResult.Marshal(b, m, deterministic)
}
func (m *MutateKeywordPlanAdGroupResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateKeywordPlanAdGroupResult.Merge(m, src)
}
func (m *MutateKeywordPlanAdGroupResult) XXX_Size() int {
	return xxx_messageInfo_MutateKeywordPlanAdGroupResult.Size(m)
}
func (m *MutateKeywordPlanAdGroupResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateKeywordPlanAdGroupResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateKeywordPlanAdGroupResult proto.InternalMessageInfo

func (m *MutateKeywordPlanAdGroupResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetKeywordPlanAdGroupRequest)(nil), "google.ads.googleads.v0.services.GetKeywordPlanAdGroupRequest")
	proto.RegisterType((*MutateKeywordPlanAdGroupsRequest)(nil), "google.ads.googleads.v0.services.MutateKeywordPlanAdGroupsRequest")
	proto.RegisterType((*KeywordPlanAdGroupOperation)(nil), "google.ads.googleads.v0.services.KeywordPlanAdGroupOperation")
	proto.RegisterType((*MutateKeywordPlanAdGroupsResponse)(nil), "google.ads.googleads.v0.services.MutateKeywordPlanAdGroupsResponse")
	proto.RegisterType((*MutateKeywordPlanAdGroupResult)(nil), "google.ads.googleads.v0.services.MutateKeywordPlanAdGroupResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/keyword_plan_ad_group_service.proto", fileDescriptor_06d429e8eeffcde4)
}

var fileDescriptor_06d429e8eeffcde4 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x4f, 0x8b, 0xd3, 0x4e,
	0x18, 0xc7, 0x7f, 0x49, 0x7f, 0x54, 0x76, 0xa2, 0x97, 0x01, 0x21, 0x5b, 0x97, 0x35, 0x46, 0x0f,
	0xa5, 0x87, 0xa4, 0x54, 0x16, 0x65, 0x97, 0x88, 0x6d, 0xd5, 0xae, 0xc8, 0xba, 0x25, 0xc2, 0x1e,
	0x96, 0x4a, 0x98, 0x6d, 0x66, 0x43, 0x68, 0x92, 0x89, 0x33, 0x49, 0x45, 0x96, 0x45, 0xf0, 0x2d,
	0xf8, 0x0e, 0xd6, 0x9b, 0x2f, 0x45, 0xf0, 0xe4, 0xc1, 0x9b, 0x27, 0x2f, 0xbe, 0x0b, 0x99, 0x4c,
	0xa6, 0x56, 0xda, 0xb4, 0xb2, 0x7b, 0x7b, 0x32, 0xf3, 0xcc, 0xe7, 0xf9, 0xf3, 0x9d, 0x67, 0x02,
	0x9e, 0x04, 0x84, 0x04, 0x11, 0xb6, 0x91, 0xcf, 0x6c, 0x61, 0x72, 0x6b, 0xda, 0xb6, 0x19, 0xa6,
	0xd3, 0x70, 0x8c, 0x99, 0x3d, 0xc1, 0xef, 0xde, 0x12, 0xea, 0x7b, 0x69, 0x84, 0x12, 0x0f, 0xf9,
	0x5e, 0x40, 0x49, 0x9e, 0x7a, 0xe5, 0xb6, 0x95, 0x52, 0x92, 0x11, 0x68, 0x88, 0xa3, 0x16, 0xf2,
	0x99, 0x35, 0xa3, 0x58, 0xd3, 0xb6, 0x25, 0x29, 0x0d, 0xa7, 0x2a, 0x0e, 0xc5, 0x8c, 0xe4, 0xb4,
	0x32, 0x90, 0x08, 0xd0, 0xd8, 0x92, 0xc7, 0xd3, 0xd0, 0x46, 0x49, 0x42, 0x32, 0x94, 0x85, 0x24,
	0x61, 0xe5, 0x6e, 0x19, 0xde, 0x2e, 0xbe, 0x4e, 0xf2, 0x53, 0xfb, 0x34, 0xc4, 0x91, 0xef, 0xc5,
	0x88, 0x4d, 0x84, 0x87, 0xd9, 0x07, 0x5b, 0x03, 0x9c, 0xbd, 0x10, 0x11, 0x86, 0x11, 0x4a, 0xba,
	0xfe, 0x80, 0xe3, 0x5d, 0xfc, 0x26, 0xc7, 0x2c, 0x83, 0x77, 0xc1, 0x0d, 0x99, 0x88, 0x97, 0xa0,
	0x18, 0xeb, 0x8a, 0xa1, 0x34, 0x37, 0xdc, 0xeb, 0x72, 0xf1, 0x25, 0x8a, 0xb1, 0x79, 0xa1, 0x00,
	0xe3, 0x20, 0xcf, 0x50, 0x86, 0x17, 0x41, 0x4c, 0x92, 0x6e, 0x03, 0x6d, 0x9c, 0xb3, 0x8c, 0xc4,
	0x98, 0x7a, 0xa1, 0x5f, 0x72, 0x80, 0x5c, 0x7a, 0xee, 0xc3, 0xd7, 0x00, 0x90, 0x14, 0x53, 0x51,
	0x80, 0xae, 0x1a, 0xb5, 0xa6, 0xd6, 0x71, 0xac, 0x75, 0x0d, 0xb4, 0x16, 0x43, 0x1e, 0x4a, 0x8a,
	0x3b, 0x07, 0x34, 0x3f, 0xa9, 0xe0, 0xd6, 0x0a, 0x5f, 0xb8, 0x07, 0xb4, 0x3c, 0xf5, 0x51, 0x86,
	0x8b, 0xf6, 0xe8, 0xff, 0x1b, 0x4a, 0x53, 0xeb, 0x34, 0x64, 0x7c, 0xd9, 0x41, 0xeb, 0x19, 0xef,
	0xe0, 0x01, 0x62, 0x13, 0x17, 0x08, 0x77, 0x6e, 0xc3, 0x43, 0x50, 0x1f, 0x53, 0x8c, 0x32, 0xd1,
	0x1f, 0xad, 0xb3, 0x53, 0x99, 0xf7, 0x4c, 0xd6, 0x25, 0x89, 0xef, 0xff, 0xe7, 0x96, 0x18, 0x0e,
	0x14, 0x78, 0x5d, 0xbd, 0x22, 0x50, 0x60, 0xa0, 0x0e, 0xea, 0x14, 0xc7, 0x64, 0x8a, 0xf5, 0x1a,
	0xef, 0x3c, 0xdf, 0x11, 0xdf, 0x3d, 0x0d, 0x6c, 0xcc, 0xda, 0x64, 0xbe, 0x07, 0x77, 0x56, 0x28,
	0xc9, 0x52, 0x92, 0x30, 0x0c, 0x8f, 0xc1, 0x35, 0x8a, 0x59, 0x1e, 0x65, 0x52, 0xa6, 0xc7, 0xeb,
	0x65, 0xaa, 0xa2, 0xba, 0x05, 0xc8, 0x95, 0x40, 0xf3, 0x29, 0xd8, 0x5e, 0xed, 0xfa, 0x4f, 0x57,
	0xb2, 0xf3, 0xbd, 0x06, 0x36, 0x17, 0x09, 0xaf, 0x44, 0x36, 0xf0, 0xab, 0x02, 0x6e, 0x2e, 0xbd,
	0xf6, 0xf0, 0xd1, 0xfa, 0x4a, 0x56, 0xcd, 0x4b, 0xe3, 0x72, 0x3a, 0x99, 0xce, 0x87, 0x6f, 0x3f,
	0x3f, 0xaa, 0x0f, 0xe0, 0x0e, 0x9f, 0xfc, 0xb3, 0xbf, 0xca, 0x73, 0xe4, 0x88, 0x30, 0xbb, 0x25,
	0x9f, 0x82, 0x79, 0x55, 0xec, 0xd6, 0x39, 0xfc, 0xa5, 0x80, 0xcd, 0x4a, 0xd9, 0x60, 0xef, 0xf2,
	0xea, 0xc8, 0xe9, 0x6d, 0xf4, 0xaf, 0xc4, 0x10, 0xf7, 0xc6, 0xec, 0x17, 0x55, 0x3a, 0xe6, 0x43,
	0x5e, 0xe5, 0x9f, 0xb2, 0xce, 0xe6, 0xde, 0x05, 0xa7, 0x75, 0xbe, 0xac, 0xc8, 0xdd, 0xb8, 0x80,
	0xef, 0x2a, 0xad, 0xde, 0x0f, 0x05, 0xdc, 0x1b, 0x93, 0x78, 0x6d, 0x3e, 0xbd, 0xed, 0x4a, 0xfd,
	0x87, 0x7c, 0x98, 0x87, 0xca, 0xf1, 0x7e, 0xc9, 0x08, 0x48, 0x84, 0x92, 0xc0, 0x22, 0x34, 0xb0,
	0x03, 0x9c, 0x14, 0xa3, 0x2e, 0xdf, 0xe2, 0x34, 0x64, 0xd5, 0xbf, 0x80, 0x3d, 0x69, 0x5c, 0xa8,
	0xb5, 0x41, 0xb7, 0xfb, 0x59, 0x35, 0x06, 0x02, 0xd8, 0xf5, 0x99, 0x25, 0x4c, 0x6e, 0x1d, 0xb5,
	0xad, 0x32, 0x30, 0xfb, 0x22, 0x5d, 0x46, 0x5d, 0x9f, 0x8d, 0x66, 0x2e, 0xa3, 0xa3, 0xf6, 0x48,
	0xba, 0x9c, 0xd4, 0x8b, 0x04, 0xee, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x19, 0x0f, 0xc5,
	0x82, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeywordPlanAdGroupServiceClient is the client API for KeywordPlanAdGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeywordPlanAdGroupServiceClient interface {
	// Returns the requested Keyword Plan ad group in full detail.
	GetKeywordPlanAdGroup(ctx context.Context, in *GetKeywordPlanAdGroupRequest, opts ...grpc.CallOption) (*resources.KeywordPlanAdGroup, error)
	// Creates, updates, or removes Keyword Plan ad groups. Operation statuses are
	// returned.
	MutateKeywordPlanAdGroups(ctx context.Context, in *MutateKeywordPlanAdGroupsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanAdGroupsResponse, error)
}

type keywordPlanAdGroupServiceClient struct {
	cc *grpc.ClientConn
}

func NewKeywordPlanAdGroupServiceClient(cc *grpc.ClientConn) KeywordPlanAdGroupServiceClient {
	return &keywordPlanAdGroupServiceClient{cc}
}

func (c *keywordPlanAdGroupServiceClient) GetKeywordPlanAdGroup(ctx context.Context, in *GetKeywordPlanAdGroupRequest, opts ...grpc.CallOption) (*resources.KeywordPlanAdGroup, error) {
	out := new(resources.KeywordPlanAdGroup)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.KeywordPlanAdGroupService/GetKeywordPlanAdGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keywordPlanAdGroupServiceClient) MutateKeywordPlanAdGroups(ctx context.Context, in *MutateKeywordPlanAdGroupsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanAdGroupsResponse, error) {
	out := new(MutateKeywordPlanAdGroupsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.KeywordPlanAdGroupService/MutateKeywordPlanAdGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordPlanAdGroupServiceServer is the server API for KeywordPlanAdGroupService service.
type KeywordPlanAdGroupServiceServer interface {
	// Returns the requested Keyword Plan ad group in full detail.
	GetKeywordPlanAdGroup(context.Context, *GetKeywordPlanAdGroupRequest) (*resources.KeywordPlanAdGroup, error)
	// Creates, updates, or removes Keyword Plan ad groups. Operation statuses are
	// returned.
	MutateKeywordPlanAdGroups(context.Context, *MutateKeywordPlanAdGroupsRequest) (*MutateKeywordPlanAdGroupsResponse, error)
}

func RegisterKeywordPlanAdGroupServiceServer(s *grpc.Server, srv KeywordPlanAdGroupServiceServer) {
	s.RegisterService(&_KeywordPlanAdGroupService_serviceDesc, srv)
}

func _KeywordPlanAdGroupService_GetKeywordPlanAdGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeywordPlanAdGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanAdGroupServiceServer).GetKeywordPlanAdGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.KeywordPlanAdGroupService/GetKeywordPlanAdGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanAdGroupServiceServer).GetKeywordPlanAdGroup(ctx, req.(*GetKeywordPlanAdGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeywordPlanAdGroupService_MutateKeywordPlanAdGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateKeywordPlanAdGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanAdGroupServiceServer).MutateKeywordPlanAdGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.KeywordPlanAdGroupService/MutateKeywordPlanAdGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanAdGroupServiceServer).MutateKeywordPlanAdGroups(ctx, req.(*MutateKeywordPlanAdGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeywordPlanAdGroupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.KeywordPlanAdGroupService",
	HandlerType: (*KeywordPlanAdGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKeywordPlanAdGroup",
			Handler:    _KeywordPlanAdGroupService_GetKeywordPlanAdGroup_Handler,
		},
		{
			MethodName: "MutateKeywordPlanAdGroups",
			Handler:    _KeywordPlanAdGroupService_MutateKeywordPlanAdGroups_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/keyword_plan_ad_group_service.proto",
}
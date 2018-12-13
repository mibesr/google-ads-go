// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/video.proto

package resources

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A video.
type Video struct {
	// The resource name of the video.
	// Video resource names have the form:
	//
	// `customers/{customer_id}/videos/{video_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the video.
	Id *wrappers.StringValue `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The owner channel id of the video.
	ChannelId *wrappers.StringValue `protobuf:"bytes,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// The duration of the video in milliseconds.
	DurationMillis *wrappers.Int64Value `protobuf:"bytes,4,opt,name=duration_millis,json=durationMillis,proto3" json:"duration_millis,omitempty"`
	// The title of the video.
	Title                *wrappers.StringValue `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Video) Reset()         { *m = Video{} }
func (m *Video) String() string { return proto.CompactTextString(m) }
func (*Video) ProtoMessage()    {}
func (*Video) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbd1369129a1a79e, []int{0}
}

func (m *Video) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Video.Unmarshal(m, b)
}
func (m *Video) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Video.Marshal(b, m, deterministic)
}
func (m *Video) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Video.Merge(m, src)
}
func (m *Video) XXX_Size() int {
	return xxx_messageInfo_Video.Size(m)
}
func (m *Video) XXX_DiscardUnknown() {
	xxx_messageInfo_Video.DiscardUnknown(m)
}

var xxx_messageInfo_Video proto.InternalMessageInfo

func (m *Video) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Video) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Video) GetChannelId() *wrappers.StringValue {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

func (m *Video) GetDurationMillis() *wrappers.Int64Value {
	if m != nil {
		return m.DurationMillis
	}
	return nil
}

func (m *Video) GetTitle() *wrappers.StringValue {
	if m != nil {
		return m.Title
	}
	return nil
}

func init() {
	proto.RegisterType((*Video)(nil), "google.ads.googleads.v0.resources.Video")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/video.proto", fileDescriptor_cbd1369129a1a79e)
}

var fileDescriptor_cbd1369129a1a79e = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0xc7, 0x69, 0xf7, 0xdb, 0x0f, 0x16, 0xff, 0x41, 0x4f, 0x45, 0x45, 0x36, 0x45, 0xd8, 0x41,
	0xd3, 0x32, 0xc5, 0xcb, 0x4e, 0x1d, 0xc2, 0x98, 0xa0, 0x8c, 0x09, 0x3d, 0x48, 0x61, 0x64, 0x4b,
	0x8c, 0x81, 0x34, 0x29, 0x49, 0x3b, 0x5f, 0x87, 0x6f, 0xc1, 0xa3, 0x2f, 0xc5, 0x8b, 0x6f, 0x49,
	0x9a, 0x34, 0xbd, 0x08, 0xba, 0xdb, 0x97, 0xf0, 0xf9, 0x7e, 0x9e, 0xf0, 0x3c, 0xe0, 0x92, 0x4a,
	0x49, 0x39, 0x89, 0x10, 0xd6, 0x91, 0x8d, 0x75, 0xda, 0xc4, 0x91, 0x22, 0x5a, 0x56, 0x6a, 0x4d,
	0x74, 0xb4, 0x61, 0x98, 0x48, 0x58, 0x28, 0x59, 0xca, 0x60, 0x60, 0x19, 0x88, 0xb0, 0x86, 0x2d,
	0x0e, 0x37, 0x31, 0x6c, 0xf1, 0xc3, 0x93, 0xc6, 0x68, 0x0a, 0xab, 0xea, 0x39, 0x7a, 0x55, 0xa8,
	0x28, 0x88, 0xd2, 0x56, 0x71, 0xfa, 0xe6, 0x83, 0x6e, 0x5a, 0x2b, 0x83, 0x33, 0xb0, 0xe7, 0x6a,
	0x4b, 0x81, 0x72, 0x12, 0x7a, 0x7d, 0x6f, 0xd8, 0x5b, 0xec, 0xba, 0xc7, 0x07, 0x94, 0x93, 0xe0,
	0x02, 0xf8, 0x0c, 0x87, 0x7e, 0xdf, 0x1b, 0xee, 0x8c, 0x8e, 0x9b, 0x99, 0xd0, 0xb9, 0xe1, 0x63,
	0xa9, 0x98, 0xa0, 0x29, 0xe2, 0x15, 0x59, 0xf8, 0x0c, 0x07, 0x63, 0x00, 0xd6, 0x2f, 0x48, 0x08,
	0xc2, 0x97, 0x0c, 0x87, 0x9d, 0x2d, 0x5a, 0xbd, 0x86, 0x9f, 0xe1, 0xe0, 0x16, 0x1c, 0xe0, 0x4a,
	0xa1, 0x92, 0x49, 0xb1, 0xcc, 0x19, 0xe7, 0x4c, 0x87, 0xff, 0x8c, 0xe1, 0xe8, 0x87, 0x61, 0x26,
	0xca, 0x9b, 0x6b, 0x2b, 0xd8, 0x77, 0x9d, 0x7b, 0x53, 0x09, 0x46, 0xa0, 0x5b, 0xb2, 0x92, 0x93,
	0xb0, 0xbb, 0xc5, 0x74, 0x8b, 0x4e, 0xbe, 0x3c, 0x70, 0xbe, 0x96, 0x39, 0xfc, 0x73, 0xbb, 0x13,
	0x60, 0x56, 0x37, 0xaf, 0x5d, 0x73, 0xef, 0xe9, 0xae, 0x29, 0x50, 0xc9, 0x91, 0xa0, 0x50, 0x2a,
	0x1a, 0x51, 0x22, 0xcc, 0x24, 0x77, 0xcd, 0x82, 0xe9, 0x5f, 0x8e, 0x3b, 0x6e, 0xd3, 0xbb, 0xdf,
	0x99, 0x26, 0xc9, 0x87, 0x3f, 0x98, 0x5a, 0x65, 0x82, 0x35, 0xb4, 0xb1, 0x4e, 0x69, 0x0c, 0x17,
	0x8e, 0xfc, 0x74, 0x4c, 0x96, 0x60, 0x9d, 0xb5, 0x4c, 0x96, 0xc6, 0x59, 0xcb, 0xac, 0xfe, 0x9b,
	0x4f, 0x5c, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x47, 0x06, 0x9e, 0x37, 0x60, 0x02, 0x00, 0x00,
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/kritzware/google-ads-go/resources/bidding_strategy.proto

package resources

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/kritzware/google-ads-go/common"
	enums "github.com/kritzware/google-ads-go/enums"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A bidding strategy.
type BiddingStrategy struct {
	// The resource name of the bidding strategy.
	// Bidding strategy resource names have the form:
	//
	// `customers/{customer_id}/biddingStrategies/{bidding_strategy_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the bidding strategy.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the bidding strategy.
	// All bidding strategies within an account must be named distinctly.
	//
	// The length of this string should be between 1 and 255, inclusive,
	// in UTF-8 bytes, (trimmed).
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The type of the bidding strategy.
	// Create a bidding strategy by setting the bidding scheme.
	//
	// This field is read-only.
	Type enums.BiddingStrategyTypeEnum_BiddingStrategyType `protobuf:"varint,5,opt,name=type,proto3,enum=google.ads.googleads.v0.enums.BiddingStrategyTypeEnum_BiddingStrategyType" json:"type,omitempty"`
	// The bidding scheme.
	//
	// Only one can be set.
	//
	// Types that are valid to be assigned to Scheme:
	//	*BiddingStrategy_EnhancedCpc
	//	*BiddingStrategy_PageOnePromoted
	//	*BiddingStrategy_TargetCpa
	//	*BiddingStrategy_TargetOutrankShare
	//	*BiddingStrategy_TargetRoas
	//	*BiddingStrategy_TargetSpend
	Scheme               isBiddingStrategy_Scheme `protobuf_oneof:"scheme"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *BiddingStrategy) Reset()         { *m = BiddingStrategy{} }
func (m *BiddingStrategy) String() string { return proto.CompactTextString(m) }
func (*BiddingStrategy) ProtoMessage()    {}
func (*BiddingStrategy) Descriptor() ([]byte, []int) {
	return fileDescriptor_954492da6a94ea83, []int{0}
}
func (m *BiddingStrategy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BiddingStrategy.Unmarshal(m, b)
}
func (m *BiddingStrategy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BiddingStrategy.Marshal(b, m, deterministic)
}
func (m *BiddingStrategy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BiddingStrategy.Merge(m, src)
}
func (m *BiddingStrategy) XXX_Size() int {
	return xxx_messageInfo_BiddingStrategy.Size(m)
}
func (m *BiddingStrategy) XXX_DiscardUnknown() {
	xxx_messageInfo_BiddingStrategy.DiscardUnknown(m)
}

var xxx_messageInfo_BiddingStrategy proto.InternalMessageInfo

func (m *BiddingStrategy) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *BiddingStrategy) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *BiddingStrategy) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *BiddingStrategy) GetType() enums.BiddingStrategyTypeEnum_BiddingStrategyType {
	if m != nil {
		return m.Type
	}
	return enums.BiddingStrategyTypeEnum_UNSPECIFIED
}

type isBiddingStrategy_Scheme interface {
	isBiddingStrategy_Scheme()
}

type BiddingStrategy_EnhancedCpc struct {
	EnhancedCpc *common.EnhancedCpc `protobuf:"bytes,7,opt,name=enhanced_cpc,json=enhancedCpc,proto3,oneof"`
}

type BiddingStrategy_PageOnePromoted struct {
	PageOnePromoted *common.PageOnePromoted `protobuf:"bytes,8,opt,name=page_one_promoted,json=pageOnePromoted,proto3,oneof"`
}

type BiddingStrategy_TargetCpa struct {
	TargetCpa *common.TargetCpa `protobuf:"bytes,9,opt,name=target_cpa,json=targetCpa,proto3,oneof"`
}

type BiddingStrategy_TargetOutrankShare struct {
	TargetOutrankShare *common.TargetOutrankShare `protobuf:"bytes,10,opt,name=target_outrank_share,json=targetOutrankShare,proto3,oneof"`
}

type BiddingStrategy_TargetRoas struct {
	TargetRoas *common.TargetRoas `protobuf:"bytes,11,opt,name=target_roas,json=targetRoas,proto3,oneof"`
}

type BiddingStrategy_TargetSpend struct {
	TargetSpend *common.TargetSpend `protobuf:"bytes,12,opt,name=target_spend,json=targetSpend,proto3,oneof"`
}

func (*BiddingStrategy_EnhancedCpc) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_PageOnePromoted) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetCpa) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetOutrankShare) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetRoas) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetSpend) isBiddingStrategy_Scheme() {}

func (m *BiddingStrategy) GetScheme() isBiddingStrategy_Scheme {
	if m != nil {
		return m.Scheme
	}
	return nil
}

func (m *BiddingStrategy) GetEnhancedCpc() *common.EnhancedCpc {
	if x, ok := m.GetScheme().(*BiddingStrategy_EnhancedCpc); ok {
		return x.EnhancedCpc
	}
	return nil
}

func (m *BiddingStrategy) GetPageOnePromoted() *common.PageOnePromoted {
	if x, ok := m.GetScheme().(*BiddingStrategy_PageOnePromoted); ok {
		return x.PageOnePromoted
	}
	return nil
}

func (m *BiddingStrategy) GetTargetCpa() *common.TargetCpa {
	if x, ok := m.GetScheme().(*BiddingStrategy_TargetCpa); ok {
		return x.TargetCpa
	}
	return nil
}

func (m *BiddingStrategy) GetTargetOutrankShare() *common.TargetOutrankShare {
	if x, ok := m.GetScheme().(*BiddingStrategy_TargetOutrankShare); ok {
		return x.TargetOutrankShare
	}
	return nil
}

func (m *BiddingStrategy) GetTargetRoas() *common.TargetRoas {
	if x, ok := m.GetScheme().(*BiddingStrategy_TargetRoas); ok {
		return x.TargetRoas
	}
	return nil
}

func (m *BiddingStrategy) GetTargetSpend() *common.TargetSpend {
	if x, ok := m.GetScheme().(*BiddingStrategy_TargetSpend); ok {
		return x.TargetSpend
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*BiddingStrategy) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _BiddingStrategy_OneofMarshaler, _BiddingStrategy_OneofUnmarshaler, _BiddingStrategy_OneofSizer, []interface{}{
		(*BiddingStrategy_EnhancedCpc)(nil),
		(*BiddingStrategy_PageOnePromoted)(nil),
		(*BiddingStrategy_TargetCpa)(nil),
		(*BiddingStrategy_TargetOutrankShare)(nil),
		(*BiddingStrategy_TargetRoas)(nil),
		(*BiddingStrategy_TargetSpend)(nil),
	}
}

func _BiddingStrategy_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*BiddingStrategy)
	// scheme
	switch x := m.Scheme.(type) {
	case *BiddingStrategy_EnhancedCpc:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EnhancedCpc); err != nil {
			return err
		}
	case *BiddingStrategy_PageOnePromoted:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PageOnePromoted); err != nil {
			return err
		}
	case *BiddingStrategy_TargetCpa:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TargetCpa); err != nil {
			return err
		}
	case *BiddingStrategy_TargetOutrankShare:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TargetOutrankShare); err != nil {
			return err
		}
	case *BiddingStrategy_TargetRoas:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TargetRoas); err != nil {
			return err
		}
	case *BiddingStrategy_TargetSpend:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TargetSpend); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("BiddingStrategy.Scheme has unexpected type %T", x)
	}
	return nil
}

func _BiddingStrategy_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*BiddingStrategy)
	switch tag {
	case 7: // scheme.enhanced_cpc
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.EnhancedCpc)
		err := b.DecodeMessage(msg)
		m.Scheme = &BiddingStrategy_EnhancedCpc{msg}
		return true, err
	case 8: // scheme.page_one_promoted
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.PageOnePromoted)
		err := b.DecodeMessage(msg)
		m.Scheme = &BiddingStrategy_PageOnePromoted{msg}
		return true, err
	case 9: // scheme.target_cpa
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.TargetCpa)
		err := b.DecodeMessage(msg)
		m.Scheme = &BiddingStrategy_TargetCpa{msg}
		return true, err
	case 10: // scheme.target_outrank_share
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.TargetOutrankShare)
		err := b.DecodeMessage(msg)
		m.Scheme = &BiddingStrategy_TargetOutrankShare{msg}
		return true, err
	case 11: // scheme.target_roas
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.TargetRoas)
		err := b.DecodeMessage(msg)
		m.Scheme = &BiddingStrategy_TargetRoas{msg}
		return true, err
	case 12: // scheme.target_spend
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.TargetSpend)
		err := b.DecodeMessage(msg)
		m.Scheme = &BiddingStrategy_TargetSpend{msg}
		return true, err
	default:
		return false, nil
	}
}

func _BiddingStrategy_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*BiddingStrategy)
	// scheme
	switch x := m.Scheme.(type) {
	case *BiddingStrategy_EnhancedCpc:
		s := proto.Size(x.EnhancedCpc)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BiddingStrategy_PageOnePromoted:
		s := proto.Size(x.PageOnePromoted)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BiddingStrategy_TargetCpa:
		s := proto.Size(x.TargetCpa)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BiddingStrategy_TargetOutrankShare:
		s := proto.Size(x.TargetOutrankShare)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BiddingStrategy_TargetRoas:
		s := proto.Size(x.TargetRoas)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BiddingStrategy_TargetSpend:
		s := proto.Size(x.TargetSpend)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*BiddingStrategy)(nil), "google.ads.googleads.v0.resources.BiddingStrategy")
}

func init() {
	proto.RegisterFile("github.com/kritzware/google-ads-go/resources/bidding_strategy.proto", fileDescriptor_954492da6a94ea83)
}

var fileDescriptor_954492da6a94ea83 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xd1, 0x6e, 0xd3, 0x3c,
	0x14, 0xc7, 0x9b, 0x6e, 0x5f, 0xbf, 0xd5, 0x2d, 0x4c, 0x58, 0xbb, 0x88, 0x06, 0x42, 0x1d, 0x08,
	0xa9, 0x30, 0xe4, 0x54, 0x05, 0x21, 0xb8, 0x6c, 0xa7, 0x69, 0x65, 0x12, 0xac, 0x4a, 0xa7, 0x5e,
	0x15, 0x22, 0x37, 0x3e, 0x4b, 0x2b, 0x16, 0xdb, 0xb2, 0x9d, 0xa1, 0x3e, 0x00, 0x2f, 0xc2, 0x25,
	0x77, 0xbc, 0x06, 0x8f, 0xc1, 0x93, 0xa0, 0xd8, 0x49, 0x27, 0x56, 0xaa, 0xf6, 0xee, 0x9c, 0x93,
	0xff, 0xff, 0xe7, 0x73, 0x72, 0x6c, 0xf4, 0x36, 0x11, 0x22, 0xb9, 0x86, 0x80, 0x32, 0x1d, 0xb8,
	0x30, 0x8f, 0x6e, 0x3a, 0x81, 0x02, 0x2d, 0x32, 0x15, 0x83, 0x0e, 0xa6, 0x73, 0xc6, 0xe6, 0x3c,
	0x89, 0xb4, 0x51, 0xd4, 0x40, 0xb2, 0x20, 0x52, 0x09, 0x23, 0xf0, 0x91, 0x93, 0x13, 0xca, 0x34,
	0x59, 0x3a, 0xc9, 0x4d, 0x87, 0x2c, 0x9d, 0x87, 0x2f, 0xd7, 0xc1, 0x63, 0x91, 0xa6, 0x82, 0x97,
	0x64, 0x07, 0x3c, 0x7c, 0xb7, 0x4e, 0x0d, 0x3c, 0x4b, 0x57, 0xdb, 0x88, 0xcc, 0x42, 0x42, 0x61,
	0x7d, 0x5c, 0x58, 0x6d, 0x36, 0xcd, 0xae, 0x82, 0xaf, 0x8a, 0x4a, 0x09, 0x4a, 0xbb, 0xef, 0x4f,
	0xbe, 0xd5, 0xd0, 0x7e, 0xdf, 0xf9, 0x47, 0x85, 0x1d, 0x3f, 0x45, 0xf7, 0xca, 0x4e, 0x23, 0x4e,
	0x53, 0xf0, 0xbd, 0x96, 0xd7, 0xae, 0x87, 0xcd, 0xb2, 0xf8, 0x91, 0xa6, 0x80, 0x8f, 0x51, 0x75,
	0xce, 0xfc, 0x9d, 0x96, 0xd7, 0x6e, 0x74, 0x1f, 0x16, 0x63, 0x92, 0xf2, 0x14, 0xf2, 0x9e, 0x9b,
	0x37, 0xaf, 0xc7, 0xf4, 0x3a, 0x83, 0xb0, 0x3a, 0x67, 0xb8, 0x83, 0x76, 0x2d, 0x68, 0xd7, 0xca,
	0x1f, 0xad, 0xc8, 0x47, 0x46, 0xcd, 0x79, 0xe2, 0xf4, 0x56, 0x89, 0x3f, 0xa3, 0xdd, 0x7c, 0x0a,
	0xff, 0xbf, 0x96, 0xd7, 0xbe, 0xdf, 0x3d, 0x27, 0xeb, 0x7e, 0xa9, 0xfd, 0x03, 0xe4, 0xce, 0x04,
	0x97, 0x0b, 0x09, 0xa7, 0x3c, 0x4b, 0xff, 0x55, 0x0f, 0x2d, 0x17, 0x0f, 0x51, 0x13, 0xf8, 0x8c,
	0xf2, 0x18, 0x58, 0x14, 0xcb, 0xd8, 0xff, 0xdf, 0x76, 0x76, 0xbc, 0xf6, 0x1c, 0xb7, 0x17, 0x72,
	0x5a, 0x78, 0x4e, 0x64, 0x3c, 0xa8, 0x84, 0x0d, 0xb8, 0x4d, 0xf1, 0x27, 0xf4, 0x40, 0xd2, 0x04,
	0x22, 0xc1, 0x21, 0x92, 0x4a, 0xa4, 0xc2, 0x00, 0xf3, 0xf7, 0x2c, 0x36, 0xd8, 0x84, 0x1d, 0xd2,
	0x04, 0x2e, 0x38, 0x0c, 0x0b, 0xdb, 0xa0, 0x12, 0xee, 0xcb, 0xbf, 0x4b, 0xf8, 0x1c, 0x21, 0x43,
	0x55, 0x02, 0x26, 0x8a, 0x25, 0xf5, 0xeb, 0x96, 0xfb, 0x7c, 0x13, 0xf7, 0xd2, 0x3a, 0x4e, 0x24,
	0x1d, 0x54, 0xc2, 0xba, 0x29, 0x13, 0x7c, 0x85, 0x0e, 0x0a, 0x96, 0xc8, 0x8c, 0xa2, 0xfc, 0x4b,
	0xa4, 0x67, 0x54, 0x81, 0x8f, 0x2c, 0xb5, 0xbb, 0x1d, 0xf5, 0xc2, 0x59, 0x47, 0xb9, 0x73, 0x50,
	0x09, 0xb1, 0x59, 0xa9, 0xe2, 0x0f, 0xa8, 0x51, 0x9c, 0xa3, 0x04, 0xd5, 0x7e, 0xc3, 0xe2, 0x5f,
	0x6c, 0x87, 0x0f, 0x05, 0xd5, 0x83, 0x4a, 0x58, 0x0c, 0x9d, 0x67, 0xf9, 0xce, 0x0a, 0x9c, 0x96,
	0xc0, 0x99, 0xdf, 0xdc, 0x6e, 0x67, 0x8e, 0x37, 0xca, 0x2d, 0xf9, 0xce, 0xcc, 0x6d, 0xda, 0xdf,
	0x43, 0x35, 0x1d, 0xcf, 0x20, 0x85, 0xfe, 0x4f, 0x0f, 0x3d, 0x8b, 0x45, 0x4a, 0x36, 0x3e, 0xdd,
	0xfe, 0xc1, 0x9d, 0x4b, 0x35, 0xcc, 0x2f, 0xf1, 0xd0, 0xfb, 0x5e, 0xdd, 0x39, 0xeb, 0xf5, 0x7e,
	0x54, 0x8f, 0xce, 0x1c, 0xa1, 0xc7, 0x34, 0x71, 0x61, 0x1e, 0x8d, 0x3b, 0x24, 0x2c, 0x09, 0xbf,
	0x4a, 0xcd, 0xa4, 0xc7, 0xf4, 0x64, 0xa9, 0x99, 0x8c, 0x3b, 0x93, 0xa5, 0xe6, 0xf7, 0x16, 0x9a,
	0x69, 0xcd, 0xbe, 0x9f, 0x57, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x71, 0xb5, 0x7b, 0x25, 0xaa,
	0x04, 0x00, 0x00,
}
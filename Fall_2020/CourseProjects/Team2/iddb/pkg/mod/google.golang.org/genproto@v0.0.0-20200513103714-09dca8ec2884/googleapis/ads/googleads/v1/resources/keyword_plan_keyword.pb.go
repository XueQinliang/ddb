// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/keyword_plan_keyword.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A Keyword Plan ad group keyword.
// Max number of keyword plan keywords per plan: 2500.
type KeywordPlanKeyword struct {
	// Immutable. The resource name of the Keyword Plan ad group keyword.
	// KeywordPlanKeyword resource names have the form:
	//
	// `customers/{customer_id}/keywordPlanKeywords/{kp_ad_group_keyword_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The Keyword Plan ad group to which this keyword belongs.
	KeywordPlanAdGroup *wrappers.StringValue `protobuf:"bytes,2,opt,name=keyword_plan_ad_group,json=keywordPlanAdGroup,proto3" json:"keyword_plan_ad_group,omitempty"`
	// Output only. The ID of the Keyword Plan keyword.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The keyword text.
	Text *wrappers.StringValue `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	// The keyword match type.
	MatchType enums.KeywordMatchTypeEnum_KeywordMatchType `protobuf:"varint,5,opt,name=match_type,json=matchType,proto3,enum=google.ads.googleads.v1.enums.KeywordMatchTypeEnum_KeywordMatchType" json:"match_type,omitempty"`
	// A keyword level max cpc bid in micros, in the account currency, that
	// overrides the keyword plan ad group cpc bid.
	CpcBidMicros         *wrappers.Int64Value `protobuf:"bytes,6,opt,name=cpc_bid_micros,json=cpcBidMicros,proto3" json:"cpc_bid_micros,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *KeywordPlanKeyword) Reset()         { *m = KeywordPlanKeyword{} }
func (m *KeywordPlanKeyword) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanKeyword) ProtoMessage()    {}
func (*KeywordPlanKeyword) Descriptor() ([]byte, []int) {
	return fileDescriptor_50ba98432990bcb2, []int{0}
}

func (m *KeywordPlanKeyword) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanKeyword.Unmarshal(m, b)
}
func (m *KeywordPlanKeyword) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanKeyword.Marshal(b, m, deterministic)
}
func (m *KeywordPlanKeyword) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanKeyword.Merge(m, src)
}
func (m *KeywordPlanKeyword) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanKeyword.Size(m)
}
func (m *KeywordPlanKeyword) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanKeyword.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanKeyword proto.InternalMessageInfo

func (m *KeywordPlanKeyword) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *KeywordPlanKeyword) GetKeywordPlanAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.KeywordPlanAdGroup
	}
	return nil
}

func (m *KeywordPlanKeyword) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *KeywordPlanKeyword) GetText() *wrappers.StringValue {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *KeywordPlanKeyword) GetMatchType() enums.KeywordMatchTypeEnum_KeywordMatchType {
	if m != nil {
		return m.MatchType
	}
	return enums.KeywordMatchTypeEnum_UNSPECIFIED
}

func (m *KeywordPlanKeyword) GetCpcBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidMicros
	}
	return nil
}

func init() {
	proto.RegisterType((*KeywordPlanKeyword)(nil), "google.ads.googleads.v1.resources.KeywordPlanKeyword")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/keyword_plan_keyword.proto", fileDescriptor_50ba98432990bcb2)
}

var fileDescriptor_50ba98432990bcb2 = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4f, 0x6b, 0xd4, 0x40,
	0x14, 0x27, 0xd9, 0xb6, 0xd0, 0x58, 0x7b, 0x08, 0x88, 0xb1, 0x16, 0x6d, 0x85, 0xc2, 0x82, 0x38,
	0xb3, 0xdb, 0x4a, 0x0f, 0x51, 0x90, 0x2c, 0xca, 0xa2, 0x52, 0x59, 0xa2, 0x2c, 0x22, 0x0b, 0x61,
	0x76, 0x66, 0x9a, 0x0e, 0x9b, 0xcc, 0x8c, 0x33, 0xc9, 0xd6, 0xa5, 0xf4, 0x20, 0xf8, 0x49, 0x3c,
	0xfa, 0x51, 0xfc, 0x14, 0x3d, 0xf7, 0xe6, 0xd5, 0x93, 0x6c, 0x32, 0xc9, 0x56, 0x62, 0x5d, 0xbd,
	0xfd, 0x32, 0xef, 0xf7, 0xe7, 0xbd, 0xcc, 0x3c, 0xe7, 0x69, 0x2c, 0x44, 0x9c, 0x50, 0x88, 0x88,
	0x86, 0x25, 0x9c, 0xa3, 0x69, 0x17, 0x2a, 0xaa, 0x45, 0xae, 0x30, 0xd5, 0x70, 0x42, 0x67, 0xa7,
	0x42, 0x91, 0x48, 0x26, 0x88, 0x47, 0xe6, 0x03, 0x48, 0x25, 0x32, 0xe1, 0xee, 0x96, 0x12, 0x80,
	0x88, 0x06, 0xb5, 0x1a, 0x4c, 0xbb, 0xa0, 0x56, 0x6f, 0x1d, 0x5e, 0x17, 0x40, 0x79, 0x9e, 0x2e,
	0xcc, 0x53, 0x94, 0xe1, 0x93, 0x28, 0x9b, 0x49, 0x5a, 0x5a, 0x6f, 0xdd, 0xaf, 0x74, 0x92, 0xc1,
	0x63, 0x46, 0x13, 0x12, 0x8d, 0xe9, 0x09, 0x9a, 0x32, 0xa1, 0x0c, 0xe1, 0xce, 0x15, 0x42, 0x15,
	0x67, 0x4a, 0xf7, 0x4c, 0xa9, 0xf8, 0x1a, 0xe7, 0xc7, 0xf0, 0x54, 0x21, 0x29, 0xa9, 0xd2, 0xa6,
	0xbe, 0x7d, 0x45, 0x8a, 0x38, 0x17, 0x19, 0xca, 0x98, 0xe0, 0xa6, 0xfa, 0xe0, 0xc7, 0x8a, 0xe3,
	0xbe, 0x2e, 0xdb, 0x1a, 0x24, 0x88, 0x1b, 0xe8, 0xbe, 0x77, 0x6e, 0x56, 0x31, 0x11, 0x47, 0x29,
	0xf5, 0xac, 0x1d, 0xab, 0xbd, 0xde, 0x3b, 0xb8, 0x08, 0x56, 0x7f, 0x06, 0x8f, 0x9c, 0x87, 0x8b,
	0xf9, 0x0d, 0x92, 0x4c, 0x03, 0x2c, 0x52, 0xd8, 0xf4, 0x0a, 0x37, 0x2a, 0xa7, 0x37, 0x28, 0xa5,
	0xee, 0x67, 0xcb, 0xb9, 0xf5, 0xdb, 0x4f, 0x46, 0x24, 0x8a, 0x95, 0xc8, 0xa5, 0x67, 0xef, 0x58,
	0xed, 0x1b, 0xfb, 0xdb, 0xc6, 0x11, 0x54, 0xf3, 0x80, 0xb7, 0x99, 0x62, 0x3c, 0x1e, 0xa2, 0x24,
	0xa7, 0xbd, 0xce, 0x3f, 0xa7, 0x07, 0xa4, 0x3f, 0x77, 0x0d, 0xdd, 0x49, 0xe3, 0xcc, 0xed, 0x38,
	0x36, 0x23, 0x5e, 0xab, 0xc8, 0xbb, 0xdb, 0xc8, 0x7b, 0xc9, 0xb3, 0xc3, 0xc7, 0x65, 0x5c, 0xeb,
	0x22, 0x68, 0x85, 0x36, 0x23, 0x6e, 0xc7, 0x59, 0xc9, 0xe8, 0xa7, 0xcc, 0x5b, 0x59, 0xde, 0x63,
	0x58, 0x30, 0x5d, 0xec, 0x38, 0x8b, 0x6b, 0xf6, 0x56, 0x77, 0xac, 0xf6, 0xe6, 0xfe, 0x73, 0x70,
	0xdd, 0x13, 0x2a, 0xde, 0x07, 0x30, 0xed, 0x1f, 0xcd, 0x75, 0xef, 0x66, 0x92, 0xbe, 0xe0, 0x79,
	0xda, 0x38, 0x0c, 0xd7, 0xd3, 0x0a, 0xba, 0x81, 0xb3, 0x89, 0x25, 0x8e, 0xc6, 0x8c, 0x44, 0x29,
	0xc3, 0x4a, 0x68, 0x6f, 0x6d, 0xe9, 0x50, 0xe1, 0x06, 0x96, 0xb8, 0xc7, 0xc8, 0x51, 0x21, 0xf0,
	0x3f, 0x5e, 0x06, 0xfc, 0xbf, 0xee, 0xd3, 0x7d, 0x86, 0x73, 0x9d, 0x89, 0x94, 0x2a, 0x0d, 0xcf,
	0x2a, 0x78, 0x0e, 0x27, 0x0d, 0xa2, 0x86, 0x67, 0x7f, 0xda, 0xa6, 0xf3, 0xde, 0x17, 0xdb, 0xd9,
	0xc3, 0x22, 0x05, 0x4b, 0xf7, 0xa9, 0x77, 0xbb, 0x19, 0x3f, 0x98, 0x4f, 0x34, 0xb0, 0x3e, 0xbc,
	0x32, 0xea, 0x58, 0x24, 0x88, 0xc7, 0x40, 0xa8, 0x18, 0xc6, 0x94, 0x17, 0xf3, 0xc2, 0xc5, 0x0c,
	0x7f, 0x59, 0xf5, 0x27, 0x35, 0xfa, 0x6a, 0xb7, 0xfa, 0x41, 0xf0, 0xcd, 0xde, 0xed, 0x97, 0x96,
	0x01, 0xd1, 0xa0, 0x84, 0x73, 0x34, 0xec, 0x82, 0xb0, 0x62, 0x7e, 0xaf, 0x38, 0xa3, 0x80, 0xe8,
	0x51, 0xcd, 0x19, 0x0d, 0xbb, 0xa3, 0x9a, 0x73, 0x69, 0xef, 0x95, 0x05, 0xdf, 0x0f, 0x88, 0xf6,
	0xfd, 0x9a, 0xe5, 0xfb, 0xc3, 0xae, 0xef, 0xd7, 0xbc, 0xf1, 0x5a, 0xd1, 0xec, 0xc1, 0xaf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x98, 0xdb, 0xba, 0xd6, 0x96, 0x04, 0x00, 0x00,
}
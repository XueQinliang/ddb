// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/mutate_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Enum describing possible mutate errors.
type MutateErrorEnum_MutateError int32

const (
	// Enum unspecified.
	MutateErrorEnum_UNSPECIFIED MutateErrorEnum_MutateError = 0
	// The received error code is not known in this version.
	MutateErrorEnum_UNKNOWN MutateErrorEnum_MutateError = 1
	// Requested resource was not found.
	MutateErrorEnum_RESOURCE_NOT_FOUND MutateErrorEnum_MutateError = 3
	// Cannot mutate the same resource twice in one request.
	MutateErrorEnum_ID_EXISTS_IN_MULTIPLE_MUTATES MutateErrorEnum_MutateError = 7
	// The field's contents don't match another field that represents the same
	// data.
	MutateErrorEnum_INCONSISTENT_FIELD_VALUES MutateErrorEnum_MutateError = 8
	// Mutates are not allowed for the requested resource.
	MutateErrorEnum_MUTATE_NOT_ALLOWED MutateErrorEnum_MutateError = 9
	// The resource isn't in Google Ads. It belongs to another ads system.
	MutateErrorEnum_RESOURCE_NOT_IN_GOOGLE_ADS MutateErrorEnum_MutateError = 10
	// The resource being created already exists.
	MutateErrorEnum_RESOURCE_ALREADY_EXISTS MutateErrorEnum_MutateError = 11
	// This resource cannot be used with "validate_only".
	MutateErrorEnum_RESOURCE_DOES_NOT_SUPPORT_VALIDATE_ONLY MutateErrorEnum_MutateError = 12
	// Attempt to write to read-only fields.
	MutateErrorEnum_RESOURCE_READ_ONLY MutateErrorEnum_MutateError = 13
)

var MutateErrorEnum_MutateError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	3:  "RESOURCE_NOT_FOUND",
	7:  "ID_EXISTS_IN_MULTIPLE_MUTATES",
	8:  "INCONSISTENT_FIELD_VALUES",
	9:  "MUTATE_NOT_ALLOWED",
	10: "RESOURCE_NOT_IN_GOOGLE_ADS",
	11: "RESOURCE_ALREADY_EXISTS",
	12: "RESOURCE_DOES_NOT_SUPPORT_VALIDATE_ONLY",
	13: "RESOURCE_READ_ONLY",
}

var MutateErrorEnum_MutateError_value = map[string]int32{
	"UNSPECIFIED":                             0,
	"UNKNOWN":                                 1,
	"RESOURCE_NOT_FOUND":                      3,
	"ID_EXISTS_IN_MULTIPLE_MUTATES":           7,
	"INCONSISTENT_FIELD_VALUES":               8,
	"MUTATE_NOT_ALLOWED":                      9,
	"RESOURCE_NOT_IN_GOOGLE_ADS":              10,
	"RESOURCE_ALREADY_EXISTS":                 11,
	"RESOURCE_DOES_NOT_SUPPORT_VALIDATE_ONLY": 12,
	"RESOURCE_READ_ONLY":                      13,
}

func (x MutateErrorEnum_MutateError) String() string {
	return proto.EnumName(MutateErrorEnum_MutateError_name, int32(x))
}

func (MutateErrorEnum_MutateError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_29e86d4daa207c30, []int{0, 0}
}

// Container for enum describing possible mutate errors.
type MutateErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateErrorEnum) Reset()         { *m = MutateErrorEnum{} }
func (m *MutateErrorEnum) String() string { return proto.CompactTextString(m) }
func (*MutateErrorEnum) ProtoMessage()    {}
func (*MutateErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_29e86d4daa207c30, []int{0}
}

func (m *MutateErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateErrorEnum.Unmarshal(m, b)
}
func (m *MutateErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateErrorEnum.Marshal(b, m, deterministic)
}
func (m *MutateErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateErrorEnum.Merge(m, src)
}
func (m *MutateErrorEnum) XXX_Size() int {
	return xxx_messageInfo_MutateErrorEnum.Size(m)
}
func (m *MutateErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MutateErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.MutateErrorEnum_MutateError", MutateErrorEnum_MutateError_name, MutateErrorEnum_MutateError_value)
	proto.RegisterType((*MutateErrorEnum)(nil), "google.ads.googleads.v3.errors.MutateErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/mutate_error.proto", fileDescriptor_29e86d4daa207c30)
}

var fileDescriptor_29e86d4daa207c30 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0x69, 0x26, 0x31, 0x70, 0x41, 0xb3, 0x7c, 0x00, 0x31, 0x58, 0x25, 0x7a, 0xe1, 0x80,
	0xe4, 0x08, 0xe5, 0x16, 0x4e, 0x5e, 0xed, 0x56, 0x16, 0xae, 0x1d, 0xd5, 0x49, 0xc7, 0x50, 0x25,
	0x2b, 0x90, 0x28, 0xaa, 0xb4, 0xc6, 0x55, 0x9c, 0xed, 0x75, 0x90, 0x38, 0x72, 0xe5, 0x2d, 0x78,
	0x14, 0x78, 0x09, 0x94, 0x78, 0x8d, 0xb6, 0x03, 0x3b, 0xe5, 0xcb, 0xe7, 0xdf, 0xff, 0xff, 0x77,
	0xbe, 0x2f, 0xe0, 0x43, 0x65, 0x6d, 0x75, 0x55, 0x86, 0x79, 0xe1, 0x42, 0x5f, 0x76, 0xd5, 0x4d,
	0x14, 0x96, 0x4d, 0x63, 0x1b, 0x17, 0xee, 0xae, 0xdb, 0xbc, 0x2d, 0x4d, 0xff, 0x86, 0xf7, 0x8d,
	0x6d, 0x2d, 0x9a, 0x78, 0x0e, 0xe7, 0x85, 0xc3, 0x83, 0x04, 0xdf, 0x44, 0xd8, 0x4b, 0x4e, 0xdf,
	0x1c, 0x2c, 0xf7, 0xdb, 0x30, 0xaf, 0x6b, 0xdb, 0xe6, 0xed, 0xd6, 0xd6, 0xce, 0xab, 0xa7, 0xbf,
	0x02, 0x70, 0xb2, 0xec, 0x4d, 0x59, 0x87, 0xb3, 0xfa, 0x7a, 0x37, 0xfd, 0x1e, 0x80, 0xf1, 0x9d,
	0x1e, 0x3a, 0x01, 0xe3, 0x4c, 0xea, 0x84, 0xcd, 0xf8, 0x9c, 0x33, 0x0a, 0x1f, 0xa1, 0x31, 0x38,
	0xce, 0xe4, 0x27, 0xa9, 0x2e, 0x24, 0x1c, 0xa1, 0x17, 0x00, 0xad, 0x98, 0x56, 0xd9, 0x6a, 0xc6,
	0x8c, 0x54, 0xa9, 0x99, 0xab, 0x4c, 0x52, 0x78, 0x84, 0xde, 0x82, 0x33, 0x4e, 0x0d, 0xfb, 0xcc,
	0x75, 0xaa, 0x0d, 0x97, 0x66, 0x99, 0x89, 0x94, 0x27, 0x82, 0x99, 0x65, 0x96, 0x92, 0x94, 0x69,
	0x78, 0x8c, 0xce, 0xc0, 0x2b, 0x2e, 0x67, 0x4a, 0x6a, 0xae, 0x53, 0x26, 0x53, 0x33, 0xe7, 0x4c,
	0x50, 0xb3, 0x26, 0x22, 0x63, 0x1a, 0x3e, 0xe9, 0x9c, 0x3d, 0xdb, 0xfb, 0x12, 0x21, 0xd4, 0x05,
	0xa3, 0xf0, 0x29, 0x9a, 0x80, 0xd3, 0x7b, 0x89, 0x5c, 0x9a, 0x85, 0x52, 0x0b, 0xc1, 0x0c, 0xa1,
	0x1a, 0x02, 0xf4, 0x1a, 0xbc, 0x1c, 0xce, 0x89, 0x58, 0x31, 0x42, 0x2f, 0x6f, 0xef, 0x01, 0xc7,
	0xe8, 0x3d, 0x78, 0x37, 0x1c, 0x52, 0xc5, 0x74, 0xef, 0xa0, 0xb3, 0x24, 0x51, 0xab, 0xb4, 0x8b,
	0xe6, 0xb4, 0x0b, 0x54, 0x52, 0x5c, 0xc2, 0x67, 0xf7, 0xbe, 0xad, 0xf3, 0xf1, 0xfd, 0xe7, 0xe7,
	0x7f, 0x47, 0x60, 0xfa, 0xcd, 0xee, 0xf0, 0xc3, 0xa3, 0x3f, 0x87, 0x77, 0xa6, 0x98, 0x74, 0xe3,
	0x4e, 0x46, 0x5f, 0xe8, 0xad, 0xa6, 0xb2, 0x57, 0x79, 0x5d, 0x61, 0xdb, 0x54, 0x61, 0x55, 0xd6,
	0xfd, 0x32, 0x0e, 0x1b, 0xdf, 0x6f, 0xdd, 0xff, 0x7e, 0x80, 0x8f, 0xfe, 0xf1, 0x23, 0x38, 0x5a,
	0x10, 0xf2, 0x33, 0x98, 0x2c, 0xbc, 0x19, 0x29, 0x1c, 0xf6, 0x65, 0x57, 0xad, 0x23, 0xdc, 0x47,
	0xba, 0xdf, 0x07, 0x60, 0x43, 0x0a, 0xb7, 0x19, 0x80, 0xcd, 0x3a, 0xda, 0x78, 0xe0, 0x4f, 0x30,
	0xf5, 0xdd, 0x38, 0x26, 0x85, 0x8b, 0xe3, 0x01, 0x89, 0xe3, 0x75, 0x14, 0xc7, 0x1e, 0xfa, 0xfa,
	0xb8, 0xbf, 0x5d, 0xf4, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x4c, 0xdd, 0x71, 0x9d, 0x02, 0x00,
	0x00,
}
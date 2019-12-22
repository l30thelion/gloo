// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: projects/gloo/api/v1/options/azure/azure.proto

package azure

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/ext/gogoproto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type UpstreamSpec_FunctionSpec_AuthLevel int32

const (
	UpstreamSpec_FunctionSpec_Anonymous UpstreamSpec_FunctionSpec_AuthLevel = 0
	UpstreamSpec_FunctionSpec_Function  UpstreamSpec_FunctionSpec_AuthLevel = 1
	UpstreamSpec_FunctionSpec_Admin     UpstreamSpec_FunctionSpec_AuthLevel = 2
)

var UpstreamSpec_FunctionSpec_AuthLevel_name = map[int32]string{
	0: "Anonymous",
	1: "Function",
	2: "Admin",
}

var UpstreamSpec_FunctionSpec_AuthLevel_value = map[string]int32{
	"Anonymous": 0,
	"Function":  1,
	"Admin":     2,
}

func (x UpstreamSpec_FunctionSpec_AuthLevel) String() string {
	return proto.EnumName(UpstreamSpec_FunctionSpec_AuthLevel_name, int32(x))
}

func (UpstreamSpec_FunctionSpec_AuthLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6af3639d20cf7966, []int{0, 0, 0}
}

// Upstream Spec for Azure Functions Upstreams
// Azure Upstreams represent a collection of Azure Functions for a particular Azure Account
// within a particular Function App
type UpstreamSpec struct {
	// The Name of the Azure Function App where the functions are grouped
	FunctionAppName string `protobuf:"bytes,1,opt,name=function_app_name,json=functionAppName,proto3" json:"function_app_name,omitempty"`
	// A [Gloo Secret Ref](https://gloo.solo.io/introduction/concepts/#Secrets) to an [Azure Publish Profile JSON file](https://azure.microsoft.com/en-us/downloads/publishing-profile-overview/).
	// {{ hide_not_implemented "Azure Secrets can be created with `glooctl secret create azure ...`" }}
	// Note that this secret is not required unless Function Discovery is enabled
	SecretRef            core.ResourceRef             `protobuf:"bytes,2,opt,name=secret_ref,json=secretRef,proto3" json:"secret_ref"`
	Functions            []*UpstreamSpec_FunctionSpec `protobuf:"bytes,3,rep,name=functions,proto3" json:"functions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6af3639d20cf7966, []int{0}
}
func (m *UpstreamSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSpec.Unmarshal(m, b)
}
func (m *UpstreamSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSpec.Marshal(b, m, deterministic)
}
func (m *UpstreamSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSpec.Merge(m, src)
}
func (m *UpstreamSpec) XXX_Size() int {
	return xxx_messageInfo_UpstreamSpec.Size(m)
}
func (m *UpstreamSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamSpec.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamSpec proto.InternalMessageInfo

func (m *UpstreamSpec) GetFunctionAppName() string {
	if m != nil {
		return m.FunctionAppName
	}
	return ""
}

func (m *UpstreamSpec) GetSecretRef() core.ResourceRef {
	if m != nil {
		return m.SecretRef
	}
	return core.ResourceRef{}
}

func (m *UpstreamSpec) GetFunctions() []*UpstreamSpec_FunctionSpec {
	if m != nil {
		return m.Functions
	}
	return nil
}

// Function Spec for Functions on Azure Functions Upstreams
// The Function Spec contains data necessary for Gloo to invoke Azure functions
type UpstreamSpec_FunctionSpec struct {
	// The Name of the Azure Function as it appears in the Azure Functions Portal
	FunctionName string `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	// Auth Level can bve either "anonymous" "function" or "admin"
	// See https://vincentlauzon.com/2017/12/04/azure-functions-http-authorization-levels/ for more details
	AuthLevel            UpstreamSpec_FunctionSpec_AuthLevel `protobuf:"varint,2,opt,name=auth_level,json=authLevel,proto3,enum=azure.options.gloo.solo.io.UpstreamSpec_FunctionSpec_AuthLevel" json:"auth_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *UpstreamSpec_FunctionSpec) Reset()         { *m = UpstreamSpec_FunctionSpec{} }
func (m *UpstreamSpec_FunctionSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec_FunctionSpec) ProtoMessage()    {}
func (*UpstreamSpec_FunctionSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6af3639d20cf7966, []int{0, 0}
}
func (m *UpstreamSpec_FunctionSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSpec_FunctionSpec.Unmarshal(m, b)
}
func (m *UpstreamSpec_FunctionSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSpec_FunctionSpec.Marshal(b, m, deterministic)
}
func (m *UpstreamSpec_FunctionSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSpec_FunctionSpec.Merge(m, src)
}
func (m *UpstreamSpec_FunctionSpec) XXX_Size() int {
	return xxx_messageInfo_UpstreamSpec_FunctionSpec.Size(m)
}
func (m *UpstreamSpec_FunctionSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamSpec_FunctionSpec.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamSpec_FunctionSpec proto.InternalMessageInfo

func (m *UpstreamSpec_FunctionSpec) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func (m *UpstreamSpec_FunctionSpec) GetAuthLevel() UpstreamSpec_FunctionSpec_AuthLevel {
	if m != nil {
		return m.AuthLevel
	}
	return UpstreamSpec_FunctionSpec_Anonymous
}

type DestinationSpec struct {
	// The Function Name of the FunctionSpec to be invoked.
	FunctionName         string   `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DestinationSpec) Reset()         { *m = DestinationSpec{} }
func (m *DestinationSpec) String() string { return proto.CompactTextString(m) }
func (*DestinationSpec) ProtoMessage()    {}
func (*DestinationSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6af3639d20cf7966, []int{1}
}
func (m *DestinationSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestinationSpec.Unmarshal(m, b)
}
func (m *DestinationSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestinationSpec.Marshal(b, m, deterministic)
}
func (m *DestinationSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestinationSpec.Merge(m, src)
}
func (m *DestinationSpec) XXX_Size() int {
	return xxx_messageInfo_DestinationSpec.Size(m)
}
func (m *DestinationSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_DestinationSpec.DiscardUnknown(m)
}

var xxx_messageInfo_DestinationSpec proto.InternalMessageInfo

func (m *DestinationSpec) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func init() {
	proto.RegisterEnum("azure.options.gloo.solo.io.UpstreamSpec_FunctionSpec_AuthLevel", UpstreamSpec_FunctionSpec_AuthLevel_name, UpstreamSpec_FunctionSpec_AuthLevel_value)
	proto.RegisterType((*UpstreamSpec)(nil), "azure.options.gloo.solo.io.UpstreamSpec")
	proto.RegisterType((*UpstreamSpec_FunctionSpec)(nil), "azure.options.gloo.solo.io.UpstreamSpec.FunctionSpec")
	proto.RegisterType((*DestinationSpec)(nil), "azure.options.gloo.solo.io.DestinationSpec")
}

func init() {
	proto.RegisterFile("projects/gloo/api/v1/options/azure/azure.proto", fileDescriptor_6af3639d20cf7966)
}

var fileDescriptor_6af3639d20cf7966 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xc1, 0xae, 0xd2, 0x40,
	0x14, 0xa5, 0x80, 0xc6, 0x0e, 0x45, 0x70, 0xe2, 0x02, 0xba, 0x50, 0x82, 0x1b, 0x62, 0xe2, 0x34,
	0x42, 0x74, 0x89, 0x29, 0x31, 0xc6, 0x85, 0x71, 0x51, 0xe2, 0xc6, 0x85, 0xcd, 0x50, 0x6f, 0xcb,
	0x48, 0xdb, 0x3b, 0x99, 0x99, 0x12, 0xf4, 0x53, 0xfc, 0x02, 0x3f, 0xc1, 0x4f, 0xe0, 0x2b, 0x5c,
	0xf8, 0x0f, 0xee, 0x4d, 0x5b, 0x8a, 0x98, 0xbc, 0xb7, 0x78, 0x6f, 0x33, 0x99, 0x7b, 0xee, 0x3d,
	0xe7, 0xde, 0x93, 0x1c, 0xc2, 0xa4, 0xc2, 0x2f, 0x10, 0x19, 0xed, 0x25, 0x29, 0xa2, 0xc7, 0xa5,
	0xf0, 0xf6, 0xcf, 0x3d, 0x94, 0x46, 0x60, 0xae, 0x3d, 0xfe, 0xad, 0x50, 0x50, 0xbf, 0xe5, 0xa0,
	0x41, 0xea, 0xd6, 0xc5, 0x69, 0x80, 0x95, 0x24, 0xa6, 0x31, 0x45, 0x26, 0xd0, 0x7d, 0x98, 0x60,
	0x82, 0xd5, 0x98, 0x57, 0xfe, 0x6a, 0x86, 0x4b, 0xe1, 0x60, 0x6a, 0x10, 0x0e, 0xe6, 0x84, 0x8d,
	0x4b, 0xca, 0xb3, 0x9d, 0x30, 0xcd, 0x42, 0x05, 0x71, 0xdd, 0x9a, 0x7e, 0xef, 0x10, 0xe7, 0x83,
	0xd4, 0x46, 0x01, 0xcf, 0xd6, 0x12, 0x22, 0xfa, 0x94, 0x3c, 0x88, 0x8b, 0x3c, 0x2a, 0xf7, 0x85,
	0x5c, 0xca, 0x30, 0xe7, 0x19, 0x8c, 0xac, 0x89, 0x35, 0xb3, 0x83, 0x41, 0xd3, 0xf0, 0xa5, 0x7c,
	0xcf, 0x33, 0xa0, 0x4b, 0x42, 0x34, 0x44, 0x0a, 0x4c, 0xa8, 0x20, 0x1e, 0xb5, 0x27, 0xd6, 0xac,
	0x37, 0x1f, 0xb3, 0x08, 0x15, 0x34, 0x47, 0xb2, 0x00, 0x34, 0x16, 0x2a, 0x82, 0x00, 0xe2, 0x55,
	0xf7, 0xf8, 0xeb, 0x71, 0x2b, 0xb0, 0x6b, 0x4a, 0x00, 0x31, 0x5d, 0x13, 0xbb, 0x91, 0xd4, 0xa3,
	0xce, 0xa4, 0x33, 0xeb, 0xcd, 0x5f, 0xb0, 0xeb, 0x1d, 0xb3, 0xcb, 0x43, 0xd9, 0x9b, 0x13, 0xb3,
	0x2c, 0x82, 0x7f, 0x3a, 0xee, 0xd1, 0x22, 0xce, 0x65, 0x8f, 0x3e, 0x21, 0xfd, 0xb3, 0xa3, 0x0b,
	0x37, 0x4e, 0x03, 0x56, 0x56, 0x3e, 0x11, 0xc2, 0x0b, 0xb3, 0x0d, 0x53, 0xd8, 0x43, 0x5a, 0x59,
	0xb9, 0x3f, 0x7f, 0x75, 0xab, 0x5b, 0x98, 0x5f, 0x98, 0xed, 0xbb, 0x52, 0x26, 0xb0, 0x79, 0xf3,
	0x9d, 0x2e, 0x88, 0x7d, 0xc6, 0x69, 0x9f, 0xd8, 0x7e, 0x8e, 0xf9, 0xd7, 0x0c, 0x0b, 0x3d, 0x6c,
	0x51, 0x87, 0xdc, 0x6b, 0x04, 0x86, 0x16, 0xb5, 0xc9, 0x1d, 0xff, 0x73, 0x26, 0xf2, 0x61, 0x7b,
	0xfa, 0x92, 0x0c, 0x5e, 0x83, 0x36, 0x22, 0xe7, 0x37, 0x32, 0xb3, 0x7a, 0xfb, 0xf3, 0x4f, 0xd7,
	0xfa, 0xf1, 0xfb, 0x91, 0xf5, 0x71, 0x99, 0x08, 0xb3, 0x2d, 0x36, 0x2c, 0xc2, 0xcc, 0xab, 0x42,
	0x20, 0xb0, 0x4e, 0xde, 0xff, 0x39, 0x94, 0xbb, 0xe4, 0xca, 0x2c, 0x6e, 0xee, 0x56, 0x29, 0x59,
	0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xaa, 0xcb, 0xa5, 0xb8, 0x02, 0x00, 0x00,
}

func (this *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.FunctionAppName != that1.FunctionAppName {
		return false
	}
	if !this.SecretRef.Equal(&that1.SecretRef) {
		return false
	}
	if len(this.Functions) != len(that1.Functions) {
		return false
	}
	for i := range this.Functions {
		if !this.Functions[i].Equal(that1.Functions[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *UpstreamSpec_FunctionSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec_FunctionSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec_FunctionSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.FunctionName != that1.FunctionName {
		return false
	}
	if this.AuthLevel != that1.AuthLevel {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DestinationSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DestinationSpec)
	if !ok {
		that2, ok := that.(DestinationSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.FunctionName != that1.FunctionName {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

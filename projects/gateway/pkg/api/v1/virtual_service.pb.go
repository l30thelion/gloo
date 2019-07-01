// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway/api/v1/virtual_service.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//A virtual service describes the set of routes to match for a set of domains.
//Domains must be unique across all virtual services within a gateway (i.e. no overlap between sets).
type VirtualService struct {
	VirtualHost *v1.VirtualHost `protobuf:"bytes,1,opt,name=virtual_host,json=virtualHost,proto3" json:"virtual_host,omitempty"`
	// If provided, the Gateway will serve TLS/SSL traffic for this set of routes
	SslConfig *v1.SslConfig `protobuf:"bytes,2,opt,name=ssl_config,json=sslConfig,proto3" json:"ssl_config,omitempty"`
	// Display only, optional descriptive name.
	// Unlike metadata.name, DisplayName can be changed without deleting the resource.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,6,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *VirtualService) Reset()         { *m = VirtualService{} }
func (m *VirtualService) String() string { return proto.CompactTextString(m) }
func (*VirtualService) ProtoMessage()    {}
func (*VirtualService) Descriptor() ([]byte, []int) {
	return fileDescriptor_93fa9472926a2049, []int{0}
}
func (m *VirtualService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualService.Unmarshal(m, b)
}
func (m *VirtualService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualService.Marshal(b, m, deterministic)
}
func (m *VirtualService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualService.Merge(m, src)
}
func (m *VirtualService) XXX_Size() int {
	return xxx_messageInfo_VirtualService.Size(m)
}
func (m *VirtualService) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualService.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualService proto.InternalMessageInfo

func (m *VirtualService) GetVirtualHost() *v1.VirtualHost {
	if m != nil {
		return m.VirtualHost
	}
	return nil
}

func (m *VirtualService) GetSslConfig() *v1.SslConfig {
	if m != nil {
		return m.SslConfig
	}
	return nil
}

func (m *VirtualService) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *VirtualService) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *VirtualService) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func init() {
	proto.RegisterType((*VirtualService)(nil), "gateway.solo.io.VirtualService")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gateway/api/v1/virtual_service.proto", fileDescriptor_93fa9472926a2049)
}

var fileDescriptor_93fa9472926a2049 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0xae, 0xd2, 0x40,
	0x14, 0xc6, 0x2d, 0x36, 0xe8, 0x1d, 0x6e, 0xfc, 0x33, 0xb9, 0xd1, 0xde, 0xbb, 0x00, 0xc2, 0x46,
	0x36, 0x76, 0x82, 0x24, 0x04, 0x89, 0xab, 0x1a, 0xa3, 0x1b, 0x5d, 0x94, 0xc4, 0x85, 0x1b, 0x32,
	0x94, 0x61, 0x18, 0x69, 0x39, 0x4d, 0xcf, 0x50, 0x65, 0x67, 0x78, 0x1a, 0x1f, 0xc5, 0xb5, 0x0f,
	0xc0, 0xc2, 0x37, 0xc0, 0x27, 0x30, 0x9d, 0x4e, 0x25, 0x18, 0x93, 0x0b, 0xab, 0x99, 0x93, 0xef,
	0xfc, 0xbe, 0x93, 0x7c, 0xe7, 0x90, 0x37, 0x52, 0xe9, 0xc5, 0x7a, 0xea, 0x47, 0x90, 0x30, 0x84,
	0x18, 0x9e, 0x2b, 0x60, 0x32, 0x06, 0x60, 0x69, 0x06, 0x9f, 0x45, 0xa4, 0x91, 0x49, 0xae, 0xc5,
	0x17, 0xbe, 0x61, 0x3c, 0x55, 0x2c, 0xef, 0xb1, 0x5c, 0x65, 0x7a, 0xcd, 0xe3, 0x09, 0x8a, 0x2c,
	0x57, 0x91, 0xf0, 0xd3, 0x0c, 0x34, 0xd0, 0x87, 0xb6, 0xcb, 0x2f, 0x3c, 0x7c, 0x05, 0x37, 0x57,
	0x12, 0x24, 0x18, 0x8d, 0x15, 0xbf, 0xb2, 0xed, 0xa6, 0xf7, 0x9f, 0x69, 0xe6, 0x5d, 0x2a, 0x5d,
	0x0d, 0x48, 0x84, 0xe6, 0x33, 0xae, 0xb9, 0x45, 0xd8, 0x09, 0x08, 0x6a, 0xae, 0xd7, 0x78, 0xc6,
	0x8c, 0xaa, 0xb6, 0xc8, 0xe0, 0xf6, 0x10, 0x8a, 0xaa, 0x82, 0x31, 0xb6, 0xdc, 0xf0, 0x2c, 0x2e,
	0xcd, 0xe0, 0xeb, 0xa6, 0x24, 0x3b, 0x3f, 0x6b, 0xe4, 0xc1, 0xc7, 0x32, 0xc9, 0x71, 0x19, 0x24,
	0x7d, 0x45, 0x2e, 0xab, 0x6c, 0x17, 0x80, 0xda, 0x73, 0xda, 0x4e, 0xb7, 0xf1, 0xe2, 0xda, 0x2f,
	0x2c, 0xaa, 0x58, 0x7d, 0xcb, 0xbc, 0x03, 0xd4, 0x61, 0x23, 0x3f, 0x14, 0x74, 0x40, 0x08, 0x62,
	0x3c, 0x89, 0x60, 0x35, 0x57, 0xd2, 0xab, 0x19, 0xf6, 0xe9, 0x31, 0x3b, 0xc6, 0xf8, 0xb5, 0x91,
	0xc3, 0x0b, 0xac, 0xbe, 0xf4, 0x19, 0xb9, 0x9c, 0x29, 0x4c, 0x63, 0xbe, 0x99, 0xac, 0x78, 0x22,
	0xbc, 0xbb, 0x6d, 0xa7, 0x7b, 0x11, 0xb8, 0xdf, 0xf6, 0xae, 0x13, 0x36, 0xac, 0xf2, 0x81, 0x27,
	0x82, 0xbe, 0x25, 0xf5, 0x32, 0x66, 0xaf, 0x6e, 0xcc, 0xaf, 0xfc, 0x08, 0x32, 0x71, 0x30, 0x37,
	0x5a, 0x70, 0xfd, 0x63, 0xd7, 0xba, 0xf3, 0x7b, 0xd7, 0x7a, 0xac, 0x05, 0xea, 0x99, 0x9a, 0xcf,
	0x47, 0x1d, 0x25, 0x57, 0x90, 0x89, 0x4e, 0x68, 0x71, 0x3a, 0x24, 0xf7, 0xab, 0x15, 0x7b, 0xf7,
	0x8c, 0xd5, 0x93, 0x63, 0xab, 0xf7, 0x56, 0x0d, 0xdc, 0xc2, 0x2c, 0xfc, 0xdb, 0x3d, 0x6a, 0x6e,
	0xf7, 0xae, 0x4b, 0x6a, 0x39, 0x6e, 0xf7, 0x2e, 0xa5, 0x8f, 0xfe, 0xb9, 0x44, 0x0c, 0x5e, 0x7e,
	0xff, 0xd5, 0x74, 0x3e, 0xf5, 0x4f, 0xbe, 0xe8, 0x74, 0x29, 0xed, 0x6e, 0xa6, 0x75, 0xb3, 0x96,
	0xfe, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x98, 0xcc, 0x82, 0x70, 0x0f, 0x03, 0x00, 0x00,
}

func (this *VirtualService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualService)
	if !ok {
		that2, ok := that.(VirtualService)
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
	if !this.VirtualHost.Equal(that1.VirtualHost) {
		return false
	}
	if !this.SslConfig.Equal(that1.SslConfig) {
		return false
	}
	if this.DisplayName != that1.DisplayName {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

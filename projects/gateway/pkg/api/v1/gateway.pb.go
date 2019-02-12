// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway/api/v1/gateway.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//@solo-kit:resource.short_name=gw
//@solo-kit:resource.plural_name=gateways
//@solo-kit:resource.resource_groups=api.gateway.solo.io
//
//A gateway describes the routes to upstreams that are reachable via a specific port on the Gateway Proxy itself.
type Gateway struct {
	// names of the the virtual services, which contain the actual routes for the gateway
	// if the list is empty, the gateway will apply all virtual services to this gateway
	VirtualServices []core.ResourceRef `protobuf:"bytes,2,rep,name=virtual_services,json=virtualServices,proto3" json:"virtual_services"`
	// the bind address the gateway should serve traffic on
	BindAddress string `protobuf:"bytes,3,opt,name=bind_address,json=bindAddress,proto3" json:"bind_address,omitempty"`
	// bind ports must not conflict across gateways in a namespace
	BindPort uint32 `protobuf:"varint,4,opt,name=bind_port,json=bindPort,proto3" json:"bind_port,omitempty"`
	// top level plugin configuration for all routes on the gateway
	Plugins *v1.ListenerPlugins `protobuf:"bytes,5,opt,name=plugins,proto3" json:"plugins,omitempty"`
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,6,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Gateway) Reset()         { *m = Gateway{} }
func (m *Gateway) String() string { return proto.CompactTextString(m) }
func (*Gateway) ProtoMessage()    {}
func (*Gateway) Descriptor() ([]byte, []int) {
	return fileDescriptor_30f7529f6633771c, []int{0}
}
func (m *Gateway) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gateway.Unmarshal(m, b)
}
func (m *Gateway) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gateway.Marshal(b, m, deterministic)
}
func (m *Gateway) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gateway.Merge(m, src)
}
func (m *Gateway) XXX_Size() int {
	return xxx_messageInfo_Gateway.Size(m)
}
func (m *Gateway) XXX_DiscardUnknown() {
	xxx_messageInfo_Gateway.DiscardUnknown(m)
}

var xxx_messageInfo_Gateway proto.InternalMessageInfo

func (m *Gateway) GetVirtualServices() []core.ResourceRef {
	if m != nil {
		return m.VirtualServices
	}
	return nil
}

func (m *Gateway) GetBindAddress() string {
	if m != nil {
		return m.BindAddress
	}
	return ""
}

func (m *Gateway) GetBindPort() uint32 {
	if m != nil {
		return m.BindPort
	}
	return 0
}

func (m *Gateway) GetPlugins() *v1.ListenerPlugins {
	if m != nil {
		return m.Plugins
	}
	return nil
}

func (m *Gateway) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *Gateway) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func init() {
	proto.RegisterType((*Gateway)(nil), "gateway.solo.io.Gateway")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gateway/api/v1/gateway.proto", fileDescriptor_30f7529f6633771c)
}

var fileDescriptor_30f7529f6633771c = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x6e, 0xda, 0x40,
	0x10, 0xc6, 0x6b, 0xa0, 0xfc, 0x59, 0x5a, 0xd1, 0x5a, 0xa8, 0x32, 0x54, 0x2d, 0x2e, 0x27, 0x1f,
	0x5a, 0xaf, 0x80, 0x43, 0x29, 0x52, 0x0f, 0xe5, 0x82, 0x54, 0xb5, 0x12, 0x32, 0xb7, 0x5e, 0xd0,
	0x62, 0xaf, 0xb7, 0x1b, 0x8c, 0xc7, 0xda, 0x5d, 0x93, 0xe4, 0x8d, 0xf2, 0x28, 0xb9, 0xe6, 0x05,
	0x38, 0xe4, 0x11, 0xf2, 0x04, 0x91, 0xed, 0x35, 0x12, 0x52, 0xa4, 0x90, 0x13, 0xcc, 0xcc, 0xf7,
	0x1b, 0x7f, 0x33, 0xb3, 0xe8, 0x27, 0xe3, 0xea, 0x7f, 0xba, 0x71, 0x7d, 0xd8, 0x61, 0x09, 0x11,
	0x7c, 0xe3, 0x80, 0x59, 0x04, 0x80, 0x13, 0x01, 0x17, 0xd4, 0x57, 0x12, 0x33, 0xa2, 0xe8, 0x25,
	0xb9, 0xc6, 0x24, 0xe1, 0x78, 0x3f, 0x2a, 0x43, 0x37, 0x11, 0xa0, 0xc0, 0xec, 0x94, 0x61, 0xc6,
	0xba, 0x1c, 0xfa, 0x5d, 0x06, 0x0c, 0xf2, 0x1a, 0xce, 0xfe, 0x15, 0xb2, 0xfe, 0xe8, 0x89, 0xaf,
	0xe4, 0xbf, 0x5b, 0xae, 0xca, 0xc6, 0x3b, 0xaa, 0x48, 0x40, 0x14, 0xd1, 0x08, 0x3e, 0x03, 0x91,
	0x8a, 0xa8, 0x54, 0x6a, 0xe0, 0xeb, 0x19, 0x80, 0xa0, 0xa1, 0x56, 0x4f, 0x9f, 0x9f, 0x3b, 0x8b,
	0x34, 0x97, 0x08, 0xb8, 0xd2, 0x23, 0xf7, 0x67, 0x2f, 0x23, 0xa3, 0x94, 0xf1, 0x58, 0x7b, 0x1c,
	0xde, 0x55, 0x50, 0x63, 0x51, 0x6c, 0xcc, 0xfc, 0x8d, 0xde, 0xed, 0xb9, 0x50, 0x29, 0x89, 0xd6,
	0x92, 0x8a, 0x3d, 0xf7, 0xa9, 0xb4, 0x2a, 0x76, 0xd5, 0x69, 0x8f, 0x7b, 0xae, 0x0f, 0x82, 0x96,
	0x2b, 0x75, 0x3d, 0x2a, 0x21, 0x15, 0x3e, 0xf5, 0x68, 0x38, 0xaf, 0xdd, 0x1e, 0x06, 0xaf, 0xbc,
	0x8e, 0x06, 0x57, 0x9a, 0x33, 0xbf, 0xa0, 0x37, 0x1b, 0x1e, 0x07, 0x6b, 0x12, 0x04, 0x82, 0x4a,
	0x69, 0x55, 0x6d, 0xc3, 0x69, 0x79, 0xed, 0x2c, 0xf7, 0xab, 0x48, 0x99, 0x1f, 0x51, 0x2b, 0x97,
	0x24, 0x20, 0x94, 0x55, 0xb3, 0x0d, 0xe7, 0xad, 0xd7, 0xcc, 0x12, 0x4b, 0x10, 0xca, 0xfc, 0x8e,
	0x1a, 0xda, 0xa8, 0xf5, 0xda, 0x36, 0x9c, 0xf6, 0xf8, 0x93, 0x9b, 0x0d, 0x71, 0xb4, 0xf0, 0x87,
	0x4b, 0x45, 0x63, 0x2a, 0x96, 0x85, 0xc8, 0x2b, 0xd5, 0xe6, 0x02, 0xd5, 0x8b, 0x23, 0x58, 0xf5,
	0x9c, 0xeb, 0x9e, 0x5a, 0x5f, 0xe5, 0xb5, 0x79, 0x2f, 0x73, 0xfd, 0x70, 0x18, 0xbc, 0x57, 0x54,
	0xaa, 0x80, 0x87, 0xe1, 0x6c, 0xc8, 0x59, 0x0c, 0x82, 0x0e, 0x3d, 0x8d, 0x9b, 0x53, 0xd4, 0x2c,
	0x1f, 0x80, 0xd5, 0xc8, 0x5b, 0x7d, 0x38, 0x6d, 0xf5, 0x57, 0x57, 0xf5, 0x0a, 0x8e, 0xea, 0xf9,
	0x8f, 0x9b, 0xfb, 0xcf, 0xc6, 0xbf, 0xc9, 0xd9, 0xef, 0x38, 0xd9, 0x32, 0x7d, 0x9c, 0x4d, 0x3d,
	0xbf, 0xca, 0xe4, 0x31, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x6d, 0xb4, 0xe2, 0x05, 0x03, 0x00, 0x00,
}

func (this *Gateway) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Gateway)
	if !ok {
		that2, ok := that.(Gateway)
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
	if len(this.VirtualServices) != len(that1.VirtualServices) {
		return false
	}
	for i := range this.VirtualServices {
		if !this.VirtualServices[i].Equal(&that1.VirtualServices[i]) {
			return false
		}
	}
	if this.BindAddress != that1.BindAddress {
		return false
	}
	if this.BindPort != that1.BindPort {
		return false
	}
	if !this.Plugins.Equal(that1.Plugins) {
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

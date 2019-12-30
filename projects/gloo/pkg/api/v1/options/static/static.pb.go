// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/static/static.proto

package static

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	options "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

// Static upstreams are used to route request to services listening at fixed IP/Host & Port pairs.
// Static upstreams can be used to proxy any kind of service, and therefore contain a ServiceSpec
// for additional service-specific configuration.
// Unlike upstreams created by service discovery, Static Upstreams must be created manually by users
type UpstreamSpec struct {
	// A list of addresses and ports
	// at least one must be specified
	Hosts []*Host `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	// Attempt to use outbound TLS
	// Gloo will automatically set this to true for port 443
	UseTls bool `protobuf:"varint,3,opt,name=use_tls,json=useTls,proto3" json:"use_tls,omitempty"`
	// An optional Service Spec describing the service listening at this address
	ServiceSpec          *options.ServiceSpec `protobuf:"bytes,5,opt,name=service_spec,json=serviceSpec,proto3" json:"service_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_c08b3c87c0f36512, []int{0}
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

func (m *UpstreamSpec) GetHosts() []*Host {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *UpstreamSpec) GetUseTls() bool {
	if m != nil {
		return m.UseTls
	}
	return false
}

func (m *UpstreamSpec) GetServiceSpec() *options.ServiceSpec {
	if m != nil {
		return m.ServiceSpec
	}
	return nil
}

// Represents a single instance of an upstream
type Host struct {
	// Address (hostname or IP)
	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	// Port the instance is listening on
	Port                 uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Host) Reset()         { *m = Host{} }
func (m *Host) String() string { return proto.CompactTextString(m) }
func (*Host) ProtoMessage()    {}
func (*Host) Descriptor() ([]byte, []int) {
	return fileDescriptor_c08b3c87c0f36512, []int{1}
}
func (m *Host) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Host.Unmarshal(m, b)
}
func (m *Host) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Host.Marshal(b, m, deterministic)
}
func (m *Host) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Host.Merge(m, src)
}
func (m *Host) XXX_Size() int {
	return xxx_messageInfo_Host.Size(m)
}
func (m *Host) XXX_DiscardUnknown() {
	xxx_messageInfo_Host.DiscardUnknown(m)
}

var xxx_messageInfo_Host proto.InternalMessageInfo

func (m *Host) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Host) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "static.options.gloo.solo.io.UpstreamSpec")
	proto.RegisterType((*Host)(nil), "static.options.gloo.solo.io.Host")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/static/static.proto", fileDescriptor_c08b3c87c0f36512)
}

var fileDescriptor_c08b3c87c0f36512 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x51, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0x25, 0xf6, 0x43, 0x4d, 0xeb, 0x25, 0x08, 0x2e, 0x15, 0xca, 0xda, 0xd3, 0x5e, 0x4c, 0x50,
	0x0f, 0x1e, 0x05, 0x11, 0xac, 0xd7, 0xad, 0x5e, 0xbc, 0x94, 0xed, 0x36, 0xa6, 0xd1, 0xad, 0x13,
	0x32, 0xb3, 0xb5, 0xbf, 0x48, 0xfc, 0x09, 0xfe, 0x1e, 0xff, 0x83, 0x77, 0xd9, 0xdd, 0x80, 0x1e,
	0x8a, 0x88, 0xa7, 0xcc, 0x7b, 0x79, 0xf3, 0xe6, 0x65, 0xc2, 0xc7, 0xc6, 0xd2, 0xa2, 0x9c, 0xc9,
	0x1c, 0x96, 0x0a, 0xa1, 0x80, 0x63, 0x0b, 0xca, 0x14, 0x00, 0xca, 0x79, 0x78, 0xd4, 0x39, 0x61,
	0x83, 0x32, 0x67, 0xd5, 0xea, 0x44, 0x81, 0x23, 0x0b, 0xcf, 0xa8, 0x90, 0x32, 0xb2, 0x79, 0x38,
	0xa4, 0xf3, 0x40, 0x20, 0x0e, 0x03, 0x0a, 0x1a, 0x59, 0xf5, 0xc9, 0xca, 0x52, 0x5a, 0x18, 0xec,
	0x1b, 0x30, 0x50, 0xeb, 0x54, 0x55, 0x35, 0x2d, 0x03, 0xa1, 0xd7, 0xd4, 0x90, 0x7a, 0x4d, 0x81,
	0x1b, 0x1a, 0x00, 0x53, 0x68, 0x55, 0xa3, 0x59, 0xf9, 0xa0, 0x5e, 0x7c, 0xe6, 0x9c, 0xf6, 0x18,
	0xee, 0xaf, 0xff, 0x17, 0x58, 0xfb, 0x95, 0xcd, 0xf5, 0x14, 0x9d, 0x0e, 0x79, 0x47, 0xaf, 0x8c,
	0xf7, 0xef, 0x1c, 0x92, 0xd7, 0xd9, 0x72, 0xe2, 0x74, 0x2e, 0xce, 0x79, 0x67, 0x01, 0x48, 0x18,
	0xb1, 0xb8, 0x95, 0xf4, 0x4e, 0x8f, 0xe4, 0x2f, 0x0f, 0x92, 0x63, 0x40, 0x4a, 0x1b, 0xbd, 0x38,
	0xe0, 0xdb, 0x25, 0xea, 0x29, 0x15, 0x18, 0xb5, 0x62, 0x96, 0xec, 0xa4, 0xdd, 0x12, 0xf5, 0x6d,
	0x81, 0xe2, 0x8a, 0xf7, 0x7f, 0x0e, 0x8e, 0x3a, 0x31, 0xab, 0x8d, 0x37, 0x3a, 0x4e, 0x1a, 0x65,
	0x15, 0x25, 0xed, 0xe1, 0x37, 0x18, 0x49, 0xde, 0xae, 0xa6, 0x09, 0xc1, 0xdb, 0xd9, 0x7c, 0xee,
	0x23, 0x16, 0xb3, 0x64, 0x37, 0xad, 0xeb, 0x8a, 0x73, 0xe0, 0x29, 0xda, 0x8a, 0x59, 0xb2, 0x97,
	0xd6, 0xf5, 0xe5, 0xcd, 0xfb, 0x67, 0x9b, 0xbd, 0x7d, 0x0c, 0xd9, 0xfd, 0xc5, 0xdf, 0x76, 0xe5,
	0x9e, 0xcc, 0xe6, 0x0f, 0x9e, 0x75, 0xeb, 0x55, 0x9d, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0x68,
	0xd3, 0xed, 0x4f, 0x26, 0x02, 0x00, 0x00,
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
	if len(this.Hosts) != len(that1.Hosts) {
		return false
	}
	for i := range this.Hosts {
		if !this.Hosts[i].Equal(that1.Hosts[i]) {
			return false
		}
	}
	if this.UseTls != that1.UseTls {
		return false
	}
	if !this.ServiceSpec.Equal(that1.ServiceSpec) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Host) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Host)
	if !ok {
		that2, ok := that.(Host)
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
	if this.Addr != that1.Addr {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

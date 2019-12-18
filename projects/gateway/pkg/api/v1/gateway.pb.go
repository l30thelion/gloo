// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway/api/v1/gateway.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//A Gateway describes a single Listener (bind address:port)
//and the routing configuration to upstreams that are reachable via a specific port on the Gateway Proxy itself.
type Gateway struct {
	// if set to false, only use virtual services without ssl configured.
	// if set to true, only use virtual services with ssl configured.
	Ssl bool `protobuf:"varint,1,opt,name=ssl,proto3" json:"ssl,omitempty"`
	// the bind address the gateway should serve traffic on
	BindAddress string `protobuf:"bytes,3,opt,name=bind_address,json=bindAddress,proto3" json:"bind_address,omitempty"`
	// bind ports must not conflict across gateways for a single proxy
	BindPort uint32 `protobuf:"varint,4,opt,name=bind_port,json=bindPort,proto3" json:"bind_port,omitempty"`
	// top level optional configuration for all routes on the gateway
	Options *v1.ListenerOptions `protobuf:"bytes,5,opt,name=options,proto3" json:"options,omitempty"`
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,6,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	// Enable ProxyProtocol support for this listener
	UseProxyProto *types.BoolValue `protobuf:"bytes,8,opt,name=use_proxy_proto,json=useProxyProto,proto3" json:"use_proxy_proto,omitempty"`
	// The type of gateway being created
	// HttpGateway creates a listener with an http_connection_manager
	// TcpGateway creates a listener with a tcp proxy filter
	//
	// Types that are valid to be assigned to GatewayType:
	//	*Gateway_HttpGateway
	//	*Gateway_TcpGateway
	GatewayType isGateway_GatewayType `protobuf_oneof:"GatewayType"`
	//
	// Names of the [`Proxy`](https://gloo.solo.io/api/github.com/solo-io/gloo/projects/gloo/api/v1/proxy.proto.sk/)
	// resources to generate from this gateway. If other gateways exist which point to the same proxy,
	// Gloo will join them together.
	//
	// Proxies have a one-to-many relationship with Envoy bootstrap configuration.
	// In order to connect to Gloo, the Envoy bootstrap configuration sets a `role` in
	// the [node metadata](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/core/base.proto#envoy-api-msg-core-node)
	// Envoy instances announce their `role` to Gloo, which maps to the  `{{ .Namespace }}~{{ .Name }}`
	// of the Proxy resource.
	//
	// The template for this value can be seen in the [Gloo Helm chart](https://github.com/solo-io/gloo/blob/master/install/helm/gloo/templates/9-gateway-proxy-configmap.yaml#L22)
	//
	// Note: this field also accepts fields written in camel-case. They will be converted
	// to kebab-case in the Proxy name. This allows use of the [Gateway Name Helm value](https://github.com/solo-io/gloo/blob/master/install/helm/gloo/values-gateway-template.yaml#L47)
	// for this field
	//
	// Defaults to `["gateway-proxy"]`
	ProxyNames           []string `protobuf:"bytes,12,rep,name=proxy_names,json=proxyNames,proto3" json:"proxy_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
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

type isGateway_GatewayType interface {
	isGateway_GatewayType()
	Equal(interface{}) bool
}

type Gateway_HttpGateway struct {
	HttpGateway *HttpGateway `protobuf:"bytes,9,opt,name=http_gateway,json=httpGateway,proto3,oneof" json:"http_gateway,omitempty"`
}
type Gateway_TcpGateway struct {
	TcpGateway *TcpGateway `protobuf:"bytes,10,opt,name=tcp_gateway,json=tcpGateway,proto3,oneof" json:"tcp_gateway,omitempty"`
}

func (*Gateway_HttpGateway) isGateway_GatewayType() {}
func (*Gateway_TcpGateway) isGateway_GatewayType()  {}

func (m *Gateway) GetGatewayType() isGateway_GatewayType {
	if m != nil {
		return m.GatewayType
	}
	return nil
}

func (m *Gateway) GetSsl() bool {
	if m != nil {
		return m.Ssl
	}
	return false
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

func (m *Gateway) GetOptions() *v1.ListenerOptions {
	if m != nil {
		return m.Options
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

func (m *Gateway) GetUseProxyProto() *types.BoolValue {
	if m != nil {
		return m.UseProxyProto
	}
	return nil
}

func (m *Gateway) GetHttpGateway() *HttpGateway {
	if x, ok := m.GetGatewayType().(*Gateway_HttpGateway); ok {
		return x.HttpGateway
	}
	return nil
}

func (m *Gateway) GetTcpGateway() *TcpGateway {
	if x, ok := m.GetGatewayType().(*Gateway_TcpGateway); ok {
		return x.TcpGateway
	}
	return nil
}

func (m *Gateway) GetProxyNames() []string {
	if m != nil {
		return m.ProxyNames
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Gateway) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Gateway_HttpGateway)(nil),
		(*Gateway_TcpGateway)(nil),
	}
}

type HttpGateway struct {
	// Names & namespace refs of the the virtual services which contain the actual routes for the gateway.
	// If the list is empty, all virtual services in the same namespace as this gateway will apply,
	// with accordance to `ssl` flag on `Gateway` above.
	// The default namespace matching behavior can be overridden via `virtual_service_namespaces` flag below.
	// Only one of `virtualServices` or `virtualServiceSelector` should be provided.
	VirtualServices []core.ResourceRef `protobuf:"bytes,1,rep,name=virtual_services,json=virtualServices,proto3" json:"virtual_services"`
	// Select virtual services by their label. This will apply only to virtual services in the same namespace
	// as the gateway resource, unless `virtual_service_namespaces` is provided below.
	// Only one of `virtualServices` or `virtualServiceSelector` should be provided.
	VirtualServiceSelector map[string]string `protobuf:"bytes,2,rep,name=virtual_service_selector,json=virtualServiceSelector,proto3" json:"virtual_service_selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Expand the search by providing a list of valid search namespaces here, or search all namespaces by setting '*'
	VirtualServiceNamespaces []string `protobuf:"bytes,3,rep,name=virtual_service_namespaces,json=virtualServiceNamespaces,proto3" json:"virtual_service_namespaces,omitempty"`
	// HTTP Gateway configuration
	Options              *v1.HttpListenerOptions `protobuf:"bytes,8,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *HttpGateway) Reset()         { *m = HttpGateway{} }
func (m *HttpGateway) String() string { return proto.CompactTextString(m) }
func (*HttpGateway) ProtoMessage()    {}
func (*HttpGateway) Descriptor() ([]byte, []int) {
	return fileDescriptor_30f7529f6633771c, []int{1}
}
func (m *HttpGateway) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpGateway.Unmarshal(m, b)
}
func (m *HttpGateway) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpGateway.Marshal(b, m, deterministic)
}
func (m *HttpGateway) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpGateway.Merge(m, src)
}
func (m *HttpGateway) XXX_Size() int {
	return xxx_messageInfo_HttpGateway.Size(m)
}
func (m *HttpGateway) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpGateway.DiscardUnknown(m)
}

var xxx_messageInfo_HttpGateway proto.InternalMessageInfo

func (m *HttpGateway) GetVirtualServices() []core.ResourceRef {
	if m != nil {
		return m.VirtualServices
	}
	return nil
}

func (m *HttpGateway) GetVirtualServiceSelector() map[string]string {
	if m != nil {
		return m.VirtualServiceSelector
	}
	return nil
}

func (m *HttpGateway) GetVirtualServiceNamespaces() []string {
	if m != nil {
		return m.VirtualServiceNamespaces
	}
	return nil
}

func (m *HttpGateway) GetOptions() *v1.HttpListenerOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type TcpGateway struct {
	// TCP hosts that the gateway can route to
	TcpHosts []*v1.TcpHost `protobuf:"bytes,1,rep,name=tcp_hosts,json=tcpHosts,proto3" json:"tcp_hosts,omitempty"`
	// TCP Gateway configuration
	Options              *v1.TcpListenerOptions `protobuf:"bytes,8,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *TcpGateway) Reset()         { *m = TcpGateway{} }
func (m *TcpGateway) String() string { return proto.CompactTextString(m) }
func (*TcpGateway) ProtoMessage()    {}
func (*TcpGateway) Descriptor() ([]byte, []int) {
	return fileDescriptor_30f7529f6633771c, []int{2}
}
func (m *TcpGateway) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpGateway.Unmarshal(m, b)
}
func (m *TcpGateway) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpGateway.Marshal(b, m, deterministic)
}
func (m *TcpGateway) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpGateway.Merge(m, src)
}
func (m *TcpGateway) XXX_Size() int {
	return xxx_messageInfo_TcpGateway.Size(m)
}
func (m *TcpGateway) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpGateway.DiscardUnknown(m)
}

var xxx_messageInfo_TcpGateway proto.InternalMessageInfo

func (m *TcpGateway) GetTcpHosts() []*v1.TcpHost {
	if m != nil {
		return m.TcpHosts
	}
	return nil
}

func (m *TcpGateway) GetOptions() *v1.TcpListenerOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func init() {
	proto.RegisterType((*Gateway)(nil), "gateway.solo.io.Gateway")
	proto.RegisterType((*HttpGateway)(nil), "gateway.solo.io.HttpGateway")
	proto.RegisterMapType((map[string]string)(nil), "gateway.solo.io.HttpGateway.VirtualServiceSelectorEntry")
	proto.RegisterType((*TcpGateway)(nil), "gateway.solo.io.TcpGateway")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gateway/api/v1/gateway.proto", fileDescriptor_30f7529f6633771c)
}

var fileDescriptor_30f7529f6633771c = []byte{
	// 732 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xcd, 0x4e, 0xdb, 0x4a,
	0x14, 0xc6, 0x49, 0x80, 0x64, 0x0c, 0x82, 0x3b, 0xe2, 0x22, 0xdf, 0xc0, 0x85, 0x90, 0x55, 0x16,
	0xf7, 0xda, 0x02, 0x16, 0x8d, 0x52, 0x5a, 0x09, 0x4b, 0x55, 0x29, 0x6a, 0x29, 0x32, 0x88, 0x45,
	0x37, 0x91, 0xe3, 0x4c, 0x1c, 0x17, 0x93, 0x19, 0xcd, 0x1c, 0x07, 0x22, 0x75, 0xc5, 0x2b, 0xf4,
	0x25, 0xda, 0x37, 0xe9, 0x53, 0xb0, 0xe8, 0x1b, 0xd0, 0x15, 0xcb, 0x6a, 0xc6, 0xe3, 0x84, 0x84,
	0x9f, 0xc2, 0x2a, 0x73, 0xce, 0x99, 0xef, 0x9b, 0x33, 0xdf, 0xf9, 0x26, 0x46, 0xaf, 0xc2, 0x08,
	0xba, 0x49, 0xcb, 0x0e, 0xe8, 0x99, 0x23, 0x68, 0x4c, 0xff, 0x8f, 0xa8, 0x13, 0xc6, 0x94, 0x3a,
	0x8c, 0xd3, 0xcf, 0x24, 0x00, 0xe1, 0x84, 0x3e, 0x90, 0x73, 0x7f, 0xe0, 0xf8, 0x2c, 0x72, 0xfa,
	0x9b, 0x59, 0x68, 0x33, 0x4e, 0x81, 0xe2, 0x85, 0x2c, 0x94, 0x58, 0x3b, 0xa2, 0xe5, 0xa5, 0x90,
	0x86, 0x54, 0xd5, 0x1c, 0xb9, 0x4a, 0xb7, 0x95, 0x31, 0xb9, 0x80, 0x34, 0x49, 0x2e, 0x40, 0xe7,
	0xd6, 0x42, 0x4a, 0xc3, 0x98, 0x38, 0x2a, 0x6a, 0x25, 0x1d, 0xe7, 0x9c, 0xfb, 0x8c, 0x11, 0x2e,
	0x74, 0x7d, 0xf3, 0x9e, 0xce, 0xd4, 0xef, 0x69, 0x04, 0x59, 0x33, 0x67, 0x04, 0xfc, 0xb6, 0x0f,
	0xbe, 0x86, 0x38, 0x4f, 0x80, 0x08, 0xf0, 0x21, 0xc9, 0xce, 0xf8, 0xef, 0x09, 0x00, 0x4e, 0x3a,
	0xcf, 0xe8, 0x28, 0x8b, 0x35, 0xa4, 0xfe, 0x67, 0x79, 0x65, 0xa4, 0xc1, 0x8c, 0xd3, 0x0b, 0xad,
	0x6c, 0xb9, 0xf1, 0x2c, 0x24, 0x65, 0x10, 0xd1, 0x9e, 0xbe, 0x56, 0xf5, 0x7b, 0x01, 0xcd, 0xbe,
	0x4d, 0x07, 0x83, 0x17, 0x51, 0x5e, 0x88, 0xd8, 0x32, 0x2a, 0x46, 0xad, 0xe8, 0xc9, 0x25, 0xde,
	0x40, 0x73, 0xad, 0xa8, 0xd7, 0x6e, 0xfa, 0xed, 0x36, 0x27, 0x42, 0x58, 0xf9, 0x8a, 0x51, 0x2b,
	0x79, 0xa6, 0xcc, 0xed, 0xa6, 0x29, 0xbc, 0x82, 0x4a, 0x6a, 0x0b, 0xa3, 0x1c, 0xac, 0x42, 0xc5,
	0xa8, 0xcd, 0x7b, 0x45, 0x99, 0x38, 0xa4, 0x1c, 0xf0, 0x0b, 0x34, 0xab, 0x8f, 0xb3, 0xa6, 0x2b,
	0x46, 0xcd, 0xdc, 0xfa, 0xd7, 0x96, 0xad, 0x64, 0x16, 0xb0, 0xdf, 0x47, 0x02, 0x48, 0x8f, 0xf0,
	0x8f, 0xe9, 0x26, 0x2f, 0xdb, 0x8d, 0xf7, 0xd1, 0x4c, 0xaa, 0xbe, 0x35, 0xa3, 0x70, 0x4b, 0x76,
	0x40, 0x39, 0x19, 0xe2, 0x8e, 0x54, 0xcd, 0x5d, 0xbd, 0x71, 0x8d, 0x1f, 0x57, 0xeb, 0x53, 0xbf,
	0xae, 0xd6, 0xff, 0x02, 0x22, 0xa0, 0x1d, 0x75, 0x3a, 0x8d, 0x6a, 0x14, 0xf6, 0x28, 0x27, 0x55,
	0x4f, 0x33, 0xe0, 0x3a, 0x2a, 0x66, 0xc3, 0xb7, 0x66, 0x15, 0xdb, 0xf2, 0x38, 0xdb, 0x07, 0x5d,
	0x75, 0x0b, 0x92, 0xcc, 0x1b, 0xee, 0xc6, 0x2e, 0x5a, 0x48, 0x04, 0x69, 0x2a, 0xad, 0x9b, 0x4a,
	0x2f, 0xab, 0xa8, 0x08, 0xca, 0x76, 0xea, 0x48, 0x3b, 0x73, 0xa4, 0xed, 0x52, 0x1a, 0x9f, 0xf8,
	0x71, 0x42, 0xbc, 0xf9, 0x44, 0x90, 0x43, 0x89, 0x38, 0x54, 0xb6, 0xdf, 0x45, 0x73, 0x5d, 0x00,
	0xd6, 0xd4, 0xee, 0xb7, 0x4a, 0x8a, 0x60, 0xd5, 0x9e, 0x78, 0x0d, 0xf6, 0x1e, 0x00, 0xd3, 0x83,
	0xd8, 0x9b, 0xf2, 0xcc, 0xee, 0x28, 0xc4, 0xaf, 0x91, 0x09, 0xc1, 0x88, 0x01, 0x29, 0x86, 0x95,
	0x3b, 0x0c, 0xc7, 0xc1, 0x2d, 0x02, 0x04, 0xc3, 0x08, 0xaf, 0x23, 0x33, 0xbd, 0x42, 0xcf, 0x3f,
	0x23, 0xc2, 0x9a, 0xab, 0xe4, 0x6b, 0x25, 0x0f, 0xa9, 0xd4, 0x81, 0xcc, 0x34, 0x96, 0x2f, 0xaf,
	0x0b, 0x05, 0x94, 0x0b, 0xcf, 0x2f, 0xaf, 0x0b, 0x08, 0x17, 0x35, 0xb1, 0x70, 0xe7, 0x91, 0xa9,
	0x39, 0x8e, 0x07, 0x8c, 0x54, 0xbf, 0xe6, 0x91, 0x79, 0xab, 0x4d, 0xbc, 0x8f, 0x16, 0xfb, 0x11,
	0x87, 0xc4, 0x8f, 0x9b, 0x82, 0xf0, 0x7e, 0x14, 0x10, 0x61, 0x19, 0x95, 0x7c, 0xcd, 0xdc, 0xfa,
	0x67, 0x5c, 0x60, 0x8f, 0x08, 0x9a, 0xf0, 0x80, 0x78, 0xa4, 0xa3, 0x35, 0x5e, 0xd0, 0xc0, 0x23,
	0x8d, 0xc3, 0x1c, 0x59, 0x13, 0x5c, 0x4d, 0x41, 0x62, 0x12, 0x00, 0xe5, 0x56, 0x4e, 0x71, 0xd6,
	0x1f, 0x93, 0xcc, 0x3e, 0x19, 0xe3, 0x3b, 0xd2, 0xd0, 0x37, 0x3d, 0xe0, 0x03, 0x6f, 0xb9, 0x7f,
	0x6f, 0x11, 0xef, 0xa0, 0xf2, 0xe4, 0x99, 0x4a, 0x21, 0xe6, 0xcb, 0x9b, 0xe4, 0x95, 0x4c, 0xd6,
	0x38, 0xf6, 0x60, 0x58, 0xc7, 0x2f, 0x47, 0xde, 0x4e, 0x4d, 0xb1, 0x31, 0xee, 0x6d, 0xd9, 0xdd,
	0x43, 0xfe, 0x2e, 0xbf, 0x43, 0x2b, 0x8f, 0x74, 0x2c, 0x5f, 0xe2, 0x29, 0x19, 0xa8, 0x97, 0x58,
	0xf2, 0xe4, 0x12, 0x2f, 0xa1, 0xe9, 0xbe, 0xb4, 0x97, 0x95, 0x53, 0xb9, 0x34, 0x68, 0xe4, 0xea,
	0x46, 0xf5, 0x0b, 0x42, 0xa3, 0xc9, 0xe3, 0x2d, 0x54, 0x92, 0x5e, 0xe9, 0x52, 0x01, 0xd9, 0x30,
	0xfe, 0x1e, 0xef, 0xeb, 0x38, 0x60, 0x7b, 0x54, 0x80, 0x57, 0x84, 0x74, 0x21, 0x70, 0x63, 0xf2,
	0x26, 0x95, 0x3b, 0x88, 0x87, 0x2e, 0xe2, 0xee, 0xdc, 0xb8, 0xc6, 0xb7, 0x9f, 0x6b, 0xc6, 0xa7,
	0xed, 0x27, 0x7f, 0x1d, 0xd8, 0x69, 0xa8, 0xff, 0x8b, 0x5a, 0x33, 0xea, 0xfd, 0x6c, 0xff, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0xea, 0x65, 0xf6, 0x84, 0x5b, 0x06, 0x00, 0x00,
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
	if this.Ssl != that1.Ssl {
		return false
	}
	if this.BindAddress != that1.BindAddress {
		return false
	}
	if this.BindPort != that1.BindPort {
		return false
	}
	if !this.Options.Equal(that1.Options) {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !this.UseProxyProto.Equal(that1.UseProxyProto) {
		return false
	}
	if that1.GatewayType == nil {
		if this.GatewayType != nil {
			return false
		}
	} else if this.GatewayType == nil {
		return false
	} else if !this.GatewayType.Equal(that1.GatewayType) {
		return false
	}
	if len(this.ProxyNames) != len(that1.ProxyNames) {
		return false
	}
	for i := range this.ProxyNames {
		if this.ProxyNames[i] != that1.ProxyNames[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Gateway_HttpGateway) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Gateway_HttpGateway)
	if !ok {
		that2, ok := that.(Gateway_HttpGateway)
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
	if !this.HttpGateway.Equal(that1.HttpGateway) {
		return false
	}
	return true
}
func (this *Gateway_TcpGateway) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Gateway_TcpGateway)
	if !ok {
		that2, ok := that.(Gateway_TcpGateway)
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
	if !this.TcpGateway.Equal(that1.TcpGateway) {
		return false
	}
	return true
}
func (this *HttpGateway) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HttpGateway)
	if !ok {
		that2, ok := that.(HttpGateway)
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
	if len(this.VirtualServiceSelector) != len(that1.VirtualServiceSelector) {
		return false
	}
	for i := range this.VirtualServiceSelector {
		if this.VirtualServiceSelector[i] != that1.VirtualServiceSelector[i] {
			return false
		}
	}
	if len(this.VirtualServiceNamespaces) != len(that1.VirtualServiceNamespaces) {
		return false
	}
	for i := range this.VirtualServiceNamespaces {
		if this.VirtualServiceNamespaces[i] != that1.VirtualServiceNamespaces[i] {
			return false
		}
	}
	if !this.Options.Equal(that1.Options) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TcpGateway) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TcpGateway)
	if !ok {
		that2, ok := that.(TcpGateway)
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
	if len(this.TcpHosts) != len(that1.TcpHosts) {
		return false
	}
	for i := range this.TcpHosts {
		if !this.TcpHosts[i].Equal(that1.TcpHosts[i]) {
			return false
		}
	}
	if !this.Options.Equal(that1.Options) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

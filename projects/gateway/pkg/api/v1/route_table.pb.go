// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway/api/v1/route_table.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
//
// The **RouteTable** is a child Routing object for the Gloo Gateway.
//
// A **RouteTable** gets built into the complete routing configuration
// when it is referenced by a `delegateAction`, either
// in a parent VirtualService or another RouteTable.
//
// Routes specified in a RouteTable must have their paths start with the prefix provided in the
// parent's matcher.
//
// For example, the following configuration:
//
//
// ```
// virtualService: mydomain.com
// match: /a
// delegate: a-routes
// ---
// routeTable: a-routes
// match: /1
//
// ```
//
// would *not be valid*, while
//
// ```
// virtualService: mydomain.com
// match: /a
// delegate: a-routes
// ---
// routeTable: a-routes
// match: /a/1
//
// ```
//
// *would* be valid.
//
//
// A complete configuration might look as follows:
//
// ```yaml
// apiVersion: gateway.solo.io/v1
// kind: VirtualService
// metadata:
//   name: 'any'
//   namespace: 'any'
// spec:
//   virtualHost:
//     domains:
//     - 'any.com'
//     routes:
//     - matcher:
//         prefix: '/a' # delegate ownership of routes for `any.com/a`
//       delegateAction:
//         name: 'a-routes'
//         namespace: 'a'
//     - matcher:
//         prefix: '/b' # delegate ownership of routes for `any.com/b`
//       delegateAction:
//         name: 'b-routes'
//         namespace: 'b'
// ```
//
// * A root-level **VirtualService** which delegates routing to to the `a-routes` and `b-routes` **RouteTables**.
// * Routes with `delegateActions` can only use a `prefix` matcher.
//
// ```yaml
// apiVersion: gateway.solo.io/v1
// kind: RouteTable
// metadata:
//   name: 'a-routes'
//   namespace: 'a'
// spec:
//   routes:
//     - matcher:
//         # the path matchers in this RouteTable must begin with the prefix `/a/`
//         prefix: '/a/1'
//       routeAction:
//         single:
//           upstream:
//             name: 'foo-upstream'
//
//     - matcher:
//         prefix: '/a/2'
//       routeAction:
//         single:
//           upstream:
//             name: 'bar-upstream'
// ```
//
// * A **RouteTable** which defines two routes.
//
// ```yaml
// apiVersion: gateway.solo.io/v1
// kind: RouteTable
// metadata:
//   name: 'b-routes'
//   namespace: 'b'
// spec:
//   routes:
//     - matcher:
//         # the path matchers in this RouteTable must begin with the prefix `/b/`
//         regex: '/b/3'
//       routeAction:
//         single:
//           upstream:
//             name: 'bar-upstream'
//     - matcher:
//         prefix: '/b/c/'
//       # routes in the RouteTable can perform any action, including a delegateAction
//       delegateAction:
//         name: 'c-routes'
//         namespace: 'c'
//
// ```
//
// * A **RouteTable** which both *defines a route* and *delegates to* another **RouteTable**.
//
//
// ```yaml
// apiVersion: gateway.solo.io/v1
// kind: RouteTable
// metadata:
//   name: 'c-routes'
//   namespace: 'c'
// spec:
//   routes:
//     - matcher:
//         exact: '/b/c/4'
//       routeAction:
//         single:
//           upstream:
//             name: 'qux-upstream'
// ```
//
// * A RouteTable which is a child of another route table.
//
//
// Would produce the following route config for `mydomain.com`:
//
// ```
// /a/1 -> foo-upstream
// /a/2 -> bar-upstream
// /b/3 -> baz-upstream
// /b/c/4 -> qux-upstream
// ```
//
type RouteTable struct {
	// the list of routes for the route table
	Routes []*Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,6,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RouteTable) Reset()         { *m = RouteTable{} }
func (m *RouteTable) String() string { return proto.CompactTextString(m) }
func (*RouteTable) ProtoMessage()    {}
func (*RouteTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d1ea5a66e7f9a13, []int{0}
}
func (m *RouteTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteTable.Unmarshal(m, b)
}
func (m *RouteTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteTable.Marshal(b, m, deterministic)
}
func (m *RouteTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteTable.Merge(m, src)
}
func (m *RouteTable) XXX_Size() int {
	return xxx_messageInfo_RouteTable.Size(m)
}
func (m *RouteTable) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteTable.DiscardUnknown(m)
}

var xxx_messageInfo_RouteTable proto.InternalMessageInfo

func (m *RouteTable) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *RouteTable) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *RouteTable) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func init() {
	proto.RegisterType((*RouteTable)(nil), "gateway.solo.io.RouteTable")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gateway/api/v1/route_table.proto", fileDescriptor_4d1ea5a66e7f9a13)
}

var fileDescriptor_4d1ea5a66e7f9a13 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xbd, 0x4e, 0x2a, 0x41,
	0x14, 0xc7, 0xef, 0xde, 0xbb, 0xd9, 0x7b, 0x33, 0xdc, 0x68, 0xdc, 0x10, 0x42, 0x88, 0x11, 0x42,
	0x45, 0xe3, 0x4c, 0x80, 0xc6, 0x10, 0x1b, 0x37, 0xb1, 0x31, 0xb1, 0x59, 0xad, 0x6c, 0xc8, 0xb0,
	0x0c, 0xe3, 0xc8, 0xe2, 0xd9, 0xcc, 0x9c, 0x45, 0x6c, 0x79, 0x1a, 0x1f, 0xc5, 0xa7, 0xa0, 0xf0,
	0x0d, 0xb0, 0xa2, 0x34, 0x3b, 0x0c, 0x44, 0x8d, 0x05, 0x54, 0xbb, 0x73, 0xce, 0xff, 0x23, 0xf3,
	0x1b, 0x72, 0x21, 0x15, 0xde, 0xe7, 0x03, 0x9a, 0xc0, 0x84, 0x19, 0x48, 0xe1, 0x54, 0x01, 0x93,
	0x29, 0x00, 0xcb, 0x34, 0x3c, 0x88, 0x04, 0x0d, 0x93, 0x1c, 0xc5, 0x13, 0x7f, 0x66, 0x3c, 0x53,
	0x6c, 0xda, 0x66, 0x1a, 0x72, 0x14, 0x7d, 0xe4, 0x83, 0x54, 0xd0, 0x4c, 0x03, 0x42, 0x78, 0xe8,
	0x14, 0xb4, 0xf0, 0x53, 0x05, 0xb5, 0xb2, 0x04, 0x09, 0x76, 0xc7, 0x8a, 0xbf, 0xb5, 0xac, 0x16,
	0x8a, 0x19, 0xae, 0x87, 0x62, 0x86, 0x6e, 0xd6, 0xfe, 0xa1, 0xdd, 0x7e, 0xc7, 0x0a, 0x37, 0x85,
	0x13, 0x81, 0x7c, 0xc8, 0x91, 0x3b, 0x0b, 0xdb, 0xc1, 0x62, 0x90, 0x63, 0x6e, 0xf6, 0xe8, 0xd8,
	0x9c, 0x9d, 0xe5, 0x72, 0x5f, 0x28, 0x53, 0xa5, 0x31, 0xe7, 0x69, 0xdf, 0x08, 0x3d, 0x55, 0x89,
	0x03, 0xd3, 0x5c, 0x78, 0x84, 0xc4, 0x05, 0xae, 0xdb, 0x82, 0x56, 0x48, 0x49, 0x60, 0xe1, 0x99,
	0xaa, 0xd7, 0xf8, 0xd3, 0x2a, 0x75, 0x2a, 0xf4, 0x1b, 0x38, 0x6a, 0xc5, 0xb1, 0x53, 0x85, 0x57,
	0x24, 0x58, 0x5f, 0xa4, 0x1a, 0x34, 0xbc, 0x56, 0xa9, 0x53, 0xa6, 0x09, 0x68, 0xb1, 0x15, 0xdf,
	0xd8, 0x5d, 0x74, 0xbc, 0x8a, 0xbc, 0xd7, 0x45, 0xfd, 0xd7, 0xfb, 0xa2, 0x7e, 0x84, 0xc2, 0xe0,
	0x50, 0x8d, 0x46, 0xbd, 0xa6, 0x92, 0x8f, 0xa0, 0x45, 0x33, 0x76, 0x09, 0xe1, 0x19, 0xf9, 0xb7,
	0xe1, 0x58, 0xfd, 0x6b, 0xd3, 0x2a, 0x5f, 0xd3, 0xae, 0xdd, 0x36, 0xf2, 0x8b, 0xb0, 0x78, 0xab,
	0xee, 0xd5, 0xe6, 0x4b, 0xdf, 0x27, 0xbf, 0x35, 0xce, 0x97, 0xfe, 0x41, 0xf8, 0xff, 0xd3, 0xf3,
	0x9b, 0xe8, 0x7c, 0x15, 0x79, 0x2f, 0x6f, 0x27, 0xde, 0x5d, 0x77, 0x67, 0x60, 0xd9, 0x58, 0x3a,
	0x68, 0x83, 0xc0, 0x52, 0xea, 0x7e, 0x04, 0x00, 0x00, 0xff, 0xff, 0x65, 0xd7, 0xc3, 0x0a, 0x83,
	0x02, 0x00, 0x00,
}

func (this *RouteTable) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteTable)
	if !ok {
		that2, ok := that.(RouteTable)
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
	if len(this.Routes) != len(that1.Routes) {
		return false
	}
	for i := range this.Routes {
		if !this.Routes[i].Equal(that1.Routes[i]) {
			return false
		}
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

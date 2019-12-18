// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/faultinjection/fault.proto

package faultinjection

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type RouteAbort struct {
	// Percentage of requests that should be aborted, defaulting to 0.
	// This should be a value between 0.0 and 100.0, with up to 6 significant digits.
	Percentage float32 `protobuf:"fixed32,1,opt,name=percentage,proto3" json:"percentage,omitempty"`
	// This should be a standard HTTP status, i.e. 503. Defaults to 0.
	HttpStatus           uint32   `protobuf:"varint,2,opt,name=http_status,json=httpStatus,proto3" json:"http_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteAbort) Reset()         { *m = RouteAbort{} }
func (m *RouteAbort) String() string { return proto.CompactTextString(m) }
func (*RouteAbort) ProtoMessage()    {}
func (*RouteAbort) Descriptor() ([]byte, []int) {
	return fileDescriptor_2594ab72fbfb9a06, []int{0}
}
func (m *RouteAbort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteAbort.Unmarshal(m, b)
}
func (m *RouteAbort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteAbort.Marshal(b, m, deterministic)
}
func (m *RouteAbort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteAbort.Merge(m, src)
}
func (m *RouteAbort) XXX_Size() int {
	return xxx_messageInfo_RouteAbort.Size(m)
}
func (m *RouteAbort) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteAbort.DiscardUnknown(m)
}

var xxx_messageInfo_RouteAbort proto.InternalMessageInfo

func (m *RouteAbort) GetPercentage() float32 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

func (m *RouteAbort) GetHttpStatus() uint32 {
	if m != nil {
		return m.HttpStatus
	}
	return 0
}

type RouteDelay struct {
	// Percentage of requests that should be delayed, defaulting to 0.
	// This should be a value between 0.0 and 100.0, with up to 6 significant digits.
	Percentage float32 `protobuf:"fixed32,1,opt,name=percentage,proto3" json:"percentage,omitempty"`
	// Fixed delay, defaulting to 0.
	FixedDelay           *time.Duration `protobuf:"bytes,2,opt,name=fixed_delay,json=fixedDelay,proto3,stdduration" json:"fixed_delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RouteDelay) Reset()         { *m = RouteDelay{} }
func (m *RouteDelay) String() string { return proto.CompactTextString(m) }
func (*RouteDelay) ProtoMessage()    {}
func (*RouteDelay) Descriptor() ([]byte, []int) {
	return fileDescriptor_2594ab72fbfb9a06, []int{1}
}
func (m *RouteDelay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteDelay.Unmarshal(m, b)
}
func (m *RouteDelay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteDelay.Marshal(b, m, deterministic)
}
func (m *RouteDelay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteDelay.Merge(m, src)
}
func (m *RouteDelay) XXX_Size() int {
	return xxx_messageInfo_RouteDelay.Size(m)
}
func (m *RouteDelay) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteDelay.DiscardUnknown(m)
}

var xxx_messageInfo_RouteDelay proto.InternalMessageInfo

func (m *RouteDelay) GetPercentage() float32 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

func (m *RouteDelay) GetFixedDelay() *time.Duration {
	if m != nil {
		return m.FixedDelay
	}
	return nil
}

type RouteFaults struct {
	Abort                *RouteAbort `protobuf:"bytes,1,opt,name=abort,proto3" json:"abort,omitempty"`
	Delay                *RouteDelay `protobuf:"bytes,2,opt,name=delay,proto3" json:"delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RouteFaults) Reset()         { *m = RouteFaults{} }
func (m *RouteFaults) String() string { return proto.CompactTextString(m) }
func (*RouteFaults) ProtoMessage()    {}
func (*RouteFaults) Descriptor() ([]byte, []int) {
	return fileDescriptor_2594ab72fbfb9a06, []int{2}
}
func (m *RouteFaults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteFaults.Unmarshal(m, b)
}
func (m *RouteFaults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteFaults.Marshal(b, m, deterministic)
}
func (m *RouteFaults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteFaults.Merge(m, src)
}
func (m *RouteFaults) XXX_Size() int {
	return xxx_messageInfo_RouteFaults.Size(m)
}
func (m *RouteFaults) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteFaults.DiscardUnknown(m)
}

var xxx_messageInfo_RouteFaults proto.InternalMessageInfo

func (m *RouteFaults) GetAbort() *RouteAbort {
	if m != nil {
		return m.Abort
	}
	return nil
}

func (m *RouteFaults) GetDelay() *RouteDelay {
	if m != nil {
		return m.Delay
	}
	return nil
}

func init() {
	proto.RegisterType((*RouteAbort)(nil), "fault.options.gloo.solo.io.RouteAbort")
	proto.RegisterType((*RouteDelay)(nil), "fault.options.gloo.solo.io.RouteDelay")
	proto.RegisterType((*RouteFaults)(nil), "fault.options.gloo.solo.io.RouteFaults")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/faultinjection/fault.proto", fileDescriptor_2594ab72fbfb9a06)
}

var fileDescriptor_2594ab72fbfb9a06 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x4a, 0x2b, 0x31,
	0x14, 0x7d, 0xe9, 0xa3, 0x6f, 0x91, 0xe1, 0xb9, 0x18, 0x04, 0x6b, 0x17, 0x6d, 0xe9, 0x42, 0x8a,
	0x60, 0x82, 0x75, 0xeb, 0xc6, 0x52, 0x14, 0x04, 0x51, 0xc6, 0x9d, 0x9b, 0x92, 0xe9, 0xa4, 0x69,
	0x74, 0x9c, 0x1b, 0x66, 0xee, 0x94, 0xf1, 0x13, 0xfc, 0x03, 0x3f, 0x41, 0xfc, 0x04, 0x57, 0xfe,
	0x89, 0xe0, 0xce, 0x3f, 0x70, 0x29, 0x49, 0xa6, 0xd8, 0x8d, 0xd6, 0xdd, 0xbd, 0x27, 0xe7, 0xdc,
	0x9b, 0x73, 0xb8, 0xf4, 0x5c, 0x69, 0x9c, 0x97, 0x31, 0x9b, 0xc2, 0x2d, 0x2f, 0x20, 0x85, 0x3d,
	0x0d, 0x5c, 0xa5, 0x00, 0xdc, 0xe4, 0x70, 0x2d, 0xa7, 0x58, 0xf8, 0x4e, 0x18, 0xcd, 0x17, 0xfb,
	0x1c, 0x0c, 0x6a, 0xc8, 0x0a, 0x3e, 0x13, 0x65, 0x8a, 0x3a, 0xb3, 0x04, 0x0d, 0x99, 0x6f, 0x99,
	0xc9, 0x01, 0x21, 0x6c, 0xfb, 0xa6, 0x66, 0x32, 0xab, 0x66, 0x76, 0x30, 0xd3, 0xd0, 0xee, 0x28,
	0x00, 0x95, 0x4a, 0xee, 0x98, 0x71, 0x39, 0xe3, 0x49, 0x99, 0x0b, 0xcb, 0xf3, 0xda, 0xf6, 0xd6,
	0x42, 0xa4, 0x3a, 0x11, 0x28, 0xf9, 0xb2, 0xa8, 0x1f, 0x36, 0x15, 0x28, 0x70, 0x25, 0xb7, 0x55,
	0x8d, 0x86, 0xb2, 0x42, 0x0f, 0xca, 0xaa, 0x5e, 0xdf, 0x3f, 0xa3, 0x34, 0x82, 0x12, 0xe5, 0x51,
	0x0c, 0x39, 0x86, 0x1d, 0x4a, 0x8d, 0xcc, 0xa7, 0x32, 0x43, 0xa1, 0x64, 0x8b, 0xf4, 0xc8, 0xa0,
	0x11, 0xad, 0x20, 0x61, 0x97, 0x06, 0x73, 0x44, 0x33, 0x29, 0x50, 0x60, 0x59, 0xb4, 0x1a, 0x3d,
	0x32, 0xf8, 0x1f, 0x51, 0x0b, 0x5d, 0x3a, 0xa4, 0x5f, 0xd5, 0xe3, 0xc6, 0x32, 0x15, 0x77, 0x6b,
	0xc7, 0x9d, 0xd2, 0x60, 0xa6, 0x2b, 0x99, 0x4c, 0x12, 0x4b, 0x77, 0xe3, 0x82, 0xe1, 0x36, 0xf3,
	0xae, 0xd9, 0xd2, 0x35, 0x1b, 0xd7, 0xae, 0x47, 0x1b, 0x0f, 0xaf, 0x5d, 0xf2, 0xfc, 0xfe, 0xf2,
	0xb7, 0xf9, 0x44, 0x1a, 0xbb, 0x7f, 0x22, 0xea, 0xd4, 0x6e, 0x57, 0xff, 0x9e, 0xd0, 0xc0, 0xad,
	0x3e, 0xb6, 0x79, 0x16, 0xe1, 0x21, 0x6d, 0x0a, 0xeb, 0xc9, 0xad, 0x0d, 0x86, 0x3b, 0xec, 0xfb,
	0x9c, 0xd9, 0x57, 0x02, 0x91, 0x17, 0x59, 0xf5, 0xea, 0x9f, 0xd6, 0xab, 0xdd, 0x27, 0x22, 0x2f,
	0x1a, 0x5d, 0x7c, 0x8c, 0xc8, 0xe3, 0x5b, 0x87, 0x5c, 0x9d, 0xfc, 0xee, 0x5a, 0xcc, 0x8d, 0xfa,
	0xf9, 0x62, 0xe2, 0x7f, 0x2e, 0x8c, 0x83, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x2a, 0x35,
	0xbf, 0x7f, 0x02, 0x00, 0x00,
}

func (this *RouteAbort) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteAbort)
	if !ok {
		that2, ok := that.(RouteAbort)
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
	if this.Percentage != that1.Percentage {
		return false
	}
	if this.HttpStatus != that1.HttpStatus {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RouteDelay) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteDelay)
	if !ok {
		that2, ok := that.(RouteDelay)
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
	if this.Percentage != that1.Percentage {
		return false
	}
	if this.FixedDelay != nil && that1.FixedDelay != nil {
		if *this.FixedDelay != *that1.FixedDelay {
			return false
		}
	} else if this.FixedDelay != nil {
		return false
	} else if that1.FixedDelay != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RouteFaults) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteFaults)
	if !ok {
		that2, ok := that.(RouteFaults)
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
	if !this.Abort.Equal(that1.Abort) {
		return false
	}
	if !this.Delay.Equal(that1.Delay) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

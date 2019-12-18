// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/retries/retries.proto

package retries

import (
	bytes "bytes"
	fmt "fmt"
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

// Retry Policy applied at the Route and/or Virtual Hosts levels.
type RetryPolicy struct {
	// Specifies the conditions under which retry takes place. These are the same
	// conditions [documented for Envoy](https://www.envoyproxy.io/docs/envoy/latest/configuration/http_filters/router_filter#config-http-filters-router-x-envoy-retry-on)
	RetryOn string `protobuf:"bytes,1,opt,name=retry_on,json=retryOn,proto3" json:"retry_on,omitempty"`
	// Specifies the allowed number of retries. This parameter is optional and
	// defaults to 1. These are the same conditions [documented for Envoy](https://www.envoyproxy.io/docs/envoy/latest/configuration/http_filters/router_filter#config-http-filters-router-x-envoy-retry-on)
	NumRetries uint32 `protobuf:"varint,2,opt,name=num_retries,json=numRetries,proto3" json:"num_retries,omitempty"`
	// Specifies a non-zero upstream timeout per retry attempt. This parameter is optional.
	PerTryTimeout        *time.Duration `protobuf:"bytes,3,opt,name=per_try_timeout,json=perTryTimeout,proto3,stdduration" json:"per_try_timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RetryPolicy) Reset()         { *m = RetryPolicy{} }
func (m *RetryPolicy) String() string { return proto.CompactTextString(m) }
func (*RetryPolicy) ProtoMessage()    {}
func (*RetryPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c06018876f3ed3e, []int{0}
}
func (m *RetryPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetryPolicy.Unmarshal(m, b)
}
func (m *RetryPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetryPolicy.Marshal(b, m, deterministic)
}
func (m *RetryPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetryPolicy.Merge(m, src)
}
func (m *RetryPolicy) XXX_Size() int {
	return xxx_messageInfo_RetryPolicy.Size(m)
}
func (m *RetryPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_RetryPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_RetryPolicy proto.InternalMessageInfo

func (m *RetryPolicy) GetRetryOn() string {
	if m != nil {
		return m.RetryOn
	}
	return ""
}

func (m *RetryPolicy) GetNumRetries() uint32 {
	if m != nil {
		return m.NumRetries
	}
	return 0
}

func (m *RetryPolicy) GetPerTryTimeout() *time.Duration {
	if m != nil {
		return m.PerTryTimeout
	}
	return nil
}

func init() {
	proto.RegisterType((*RetryPolicy)(nil), "retries.options.gloo.solo.io.RetryPolicy")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/retries/retries.proto", fileDescriptor_3c06018876f3ed3e)
}

var fileDescriptor_3c06018876f3ed3e = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcf, 0x4a, 0xf4, 0x30,
	0x14, 0xc5, 0xc9, 0xf7, 0x89, 0x7f, 0x52, 0x06, 0xa1, 0xb8, 0xe8, 0x0c, 0x32, 0x53, 0x5c, 0x75,
	0x63, 0x82, 0xfa, 0x02, 0x52, 0x04, 0xd1, 0x8d, 0x52, 0x66, 0xe5, 0xa6, 0x4c, 0x6b, 0x8c, 0xd1,
	0x36, 0x37, 0xa4, 0x37, 0x32, 0x7d, 0x0d, 0x57, 0x3e, 0x82, 0x6f, 0x25, 0xf8, 0x14, 0x2e, 0x25,
	0x4d, 0xc7, 0x9d, 0xe0, 0x2a, 0xf7, 0x1c, 0xce, 0xc9, 0xfd, 0x71, 0xe9, 0xb5, 0x54, 0xf8, 0xe8,
	0x2a, 0x56, 0x43, 0xcb, 0x3b, 0x68, 0xe0, 0x58, 0x01, 0x97, 0x0d, 0x00, 0x37, 0x16, 0x9e, 0x44,
	0x8d, 0x5d, 0x50, 0x2b, 0xa3, 0xf8, 0xcb, 0x09, 0x07, 0x83, 0x0a, 0x74, 0xc7, 0xad, 0x40, 0xab,
	0xc4, 0xcf, 0xcb, 0x8c, 0x05, 0x84, 0xf8, 0x70, 0x23, 0xc7, 0x18, 0xf3, 0x55, 0xe6, 0x7f, 0x65,
	0x0a, 0x66, 0x73, 0x09, 0x20, 0x1b, 0xc1, 0x87, 0x6c, 0xe5, 0x1e, 0xf8, 0xbd, 0xb3, 0x2b, 0x9f,
	0x0b, 0xed, 0xd9, 0x81, 0x04, 0x09, 0xc3, 0xc8, 0xfd, 0x34, 0xba, 0xb1, 0x58, 0x63, 0x30, 0xc5,
	0x1a, 0x83, 0x77, 0xf4, 0x4a, 0x68, 0x54, 0x08, 0xb4, 0xfd, 0x2d, 0x34, 0xaa, 0xee, 0xe3, 0x29,
	0xdd, 0xf5, 0x9b, 0xfb, 0x12, 0x74, 0x42, 0x52, 0x92, 0xed, 0x15, 0x3b, 0x83, 0xbe, 0xd1, 0xf1,
	0x82, 0x46, 0xda, 0xb5, 0xe5, 0x08, 0x96, 0xfc, 0x4b, 0x49, 0x36, 0x29, 0xa8, 0x76, 0x6d, 0x11,
	0x9c, 0xf8, 0x92, 0xee, 0x1b, 0x61, 0x4b, 0xdf, 0x46, 0xd5, 0x0a, 0x70, 0x98, 0xfc, 0x4f, 0x49,
	0x16, 0x9d, 0x4e, 0x59, 0xe0, 0x65, 0x1b, 0x5e, 0x76, 0x31, 0xf2, 0xe6, 0x5b, 0x6f, 0x1f, 0x0b,
	0x52, 0x4c, 0x8c, 0xb0, 0x4b, 0xdb, 0x2f, 0x43, 0x2b, 0xbf, 0xfa, 0xca, 0xc9, 0xfb, 0xe7, 0x9c,
	0xdc, 0x9d, 0xff, 0xed, 0xa2, 0xe6, 0x59, 0xfe, 0x72, 0xd5, 0x6a, 0x7b, 0x58, 0x79, 0xf6, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x17, 0x83, 0xb8, 0x88, 0x9c, 0x01, 0x00, 0x00,
}

func (this *RetryPolicy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RetryPolicy)
	if !ok {
		that2, ok := that.(RetryPolicy)
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
	if this.RetryOn != that1.RetryOn {
		return false
	}
	if this.NumRetries != that1.NumRetries {
		return false
	}
	if this.PerTryTimeout != nil && that1.PerTryTimeout != nil {
		if *this.PerTryTimeout != *that1.PerTryTimeout {
			return false
		}
	} else if this.PerTryTimeout != nil {
		return false
	} else if that1.PerTryTimeout != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: projects/gloo/api/v1/artifact.proto

package v1

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

//
//@solo-kit:resource.short_name=art
//@solo-kit:resource.plural_name=artifacts
//
//Gloo Artifacts are used by Gloo to store small bits of binary or file data.
//
//Certain options such as the gRPC option read and write artifacts to one of Gloo's configured
//storage layer.
//
//Artifacts can be backed by files on disk, Kubernetes ConfigMaps, and Consul Key/Value pairs.
//
//Supported artifact backends can be selected in Gloo's boostrap options.
type Artifact struct {
	// Raw data data being stored
	Data map[string]string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Artifact) Reset()         { *m = Artifact{} }
func (m *Artifact) String() string { return proto.CompactTextString(m) }
func (*Artifact) ProtoMessage()    {}
func (*Artifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae18cb266eb4aa4b, []int{0}
}
func (m *Artifact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artifact.Unmarshal(m, b)
}
func (m *Artifact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artifact.Marshal(b, m, deterministic)
}
func (m *Artifact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artifact.Merge(m, src)
}
func (m *Artifact) XXX_Size() int {
	return xxx_messageInfo_Artifact.Size(m)
}
func (m *Artifact) XXX_DiscardUnknown() {
	xxx_messageInfo_Artifact.DiscardUnknown(m)
}

var xxx_messageInfo_Artifact proto.InternalMessageInfo

func (m *Artifact) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Artifact) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func init() {
	proto.RegisterType((*Artifact)(nil), "gloo.solo.io.Artifact")
	proto.RegisterMapType((map[string]string)(nil), "gloo.solo.io.Artifact.DataEntry")
}

func init() {
	proto.RegisterFile("projects/gloo/api/v1/artifact.proto", fileDescriptor_ae18cb266eb4aa4b)
}

var fileDescriptor_ae18cb266eb4aa4b = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0x28, 0xca, 0xcf,
	0x4a, 0x4d, 0x2e, 0x29, 0xd6, 0x4f, 0xcf, 0xc9, 0xcf, 0xd7, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0x33,
	0xd4, 0x4f, 0x2c, 0x2a, 0xc9, 0x4c, 0x4b, 0x4c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x01, 0xc9, 0xe9, 0x15, 0xe7, 0xe7, 0xe4, 0xeb, 0x65, 0xe6, 0x4b, 0x89, 0xa4, 0xe7, 0xa7,
	0xe7, 0x83, 0x25, 0xf4, 0x41, 0x2c, 0x88, 0x1a, 0x29, 0xa1, 0xd4, 0x8a, 0x12, 0x88, 0x60, 0x6a,
	0x05, 0x54, 0x9f, 0x94, 0x1c, 0x48, 0x8b, 0x6e, 0x76, 0x66, 0x09, 0xcc, 0xdc, 0xdc, 0xd4, 0x92,
	0xc4, 0x94, 0xc4, 0x92, 0x44, 0x5c, 0xf2, 0x30, 0x3e, 0x44, 0x5e, 0xe9, 0x32, 0x23, 0x17, 0x87,
	0x23, 0xd4, 0x29, 0x42, 0x26, 0x5c, 0x2c, 0x20, 0xad, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46,
	0x0a, 0x7a, 0xc8, 0x6e, 0xd2, 0x83, 0xa9, 0xd2, 0x73, 0x49, 0x2c, 0x49, 0x74, 0xcd, 0x2b, 0x29,
	0xaa, 0x0c, 0x02, 0xab, 0x16, 0xb2, 0xe0, 0xe2, 0x80, 0x59, 0x2a, 0xc1, 0xae, 0xc0, 0xa8, 0xc1,
	0x6d, 0x24, 0xa6, 0x97, 0x9c, 0x5f, 0x94, 0x0a, 0xd7, 0xe9, 0x0b, 0x95, 0x75, 0x62, 0x39, 0x71,
	0x4f, 0x9e, 0x21, 0x08, 0xae, 0x5a, 0xca, 0x9c, 0x8b, 0x13, 0x6e, 0x98, 0x90, 0x00, 0x17, 0x73,
	0x76, 0x6a, 0xa5, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x88, 0x29, 0x24, 0xc2, 0xc5, 0x5a,
	0x96, 0x98, 0x53, 0x9a, 0x2a, 0xc1, 0x04, 0x16, 0x83, 0x70, 0xac, 0x98, 0x2c, 0x18, 0xad, 0xe4,
	0x9a, 0x3e, 0xb2, 0xb0, 0x72, 0x31, 0x27, 0x16, 0x95, 0x34, 0x7d, 0x64, 0xe1, 0x16, 0xe2, 0x84,
	0x05, 0x66, 0x71, 0xd3, 0x47, 0x16, 0x26, 0x0d, 0x46, 0x27, 0xab, 0x1d, 0x5f, 0x59, 0x18, 0x57,
	0x3c, 0x92, 0x63, 0x8c, 0x32, 0x48, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0x85,
	0x78, 0x3d, 0x33, 0x1f, 0x12, 0x0d, 0xa8, 0x91, 0x52, 0x90, 0x9d, 0x0e, 0x0d, 0xa0, 0x24, 0x36,
	0x70, 0xc0, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x40, 0xa8, 0x0e, 0x7f, 0xb7, 0x01, 0x00,
	0x00,
}

func (this *Artifact) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Artifact)
	if !ok {
		that2, ok := that.(Artifact)
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
	if len(this.Data) != len(that1.Data) {
		return false
	}
	for i := range this.Data {
		if this.Data[i] != that1.Data[i] {
			return false
		}
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

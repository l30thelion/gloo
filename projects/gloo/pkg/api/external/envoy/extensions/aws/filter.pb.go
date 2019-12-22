// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: projects/gloo/api/external/envoy/extensions/aws/filter.proto

package aws

import (
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// AWS Lambda contains the configuration necessary to perform transform regular
// http calls to AWS Lambda invocations.
type AWSLambdaPerRoute struct {
	// The name of the function
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The qualifier of the function (defaults to $LATEST if not specified)
	Qualifier string `protobuf:"bytes,2,opt,name=qualifier,proto3" json:"qualifier,omitempty"`
	// Invocation type - async or regular.
	Async bool `protobuf:"varint,3,opt,name=async,proto3" json:"async,omitempty"`
	// Optional default body if the body is empty. By default on default
	// body is used if the body empty, and an empty body will be sent upstream.
	EmptyBodyOverride    *types.StringValue `protobuf:"bytes,4,opt,name=empty_body_override,json=emptyBodyOverride,proto3" json:"empty_body_override,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AWSLambdaPerRoute) Reset()         { *m = AWSLambdaPerRoute{} }
func (m *AWSLambdaPerRoute) String() string { return proto.CompactTextString(m) }
func (*AWSLambdaPerRoute) ProtoMessage()    {}
func (*AWSLambdaPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3e66a9e332271a5, []int{0}
}
func (m *AWSLambdaPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AWSLambdaPerRoute.Unmarshal(m, b)
}
func (m *AWSLambdaPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AWSLambdaPerRoute.Marshal(b, m, deterministic)
}
func (m *AWSLambdaPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AWSLambdaPerRoute.Merge(m, src)
}
func (m *AWSLambdaPerRoute) XXX_Size() int {
	return xxx_messageInfo_AWSLambdaPerRoute.Size(m)
}
func (m *AWSLambdaPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_AWSLambdaPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_AWSLambdaPerRoute proto.InternalMessageInfo

func (m *AWSLambdaPerRoute) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AWSLambdaPerRoute) GetQualifier() string {
	if m != nil {
		return m.Qualifier
	}
	return ""
}

func (m *AWSLambdaPerRoute) GetAsync() bool {
	if m != nil {
		return m.Async
	}
	return false
}

func (m *AWSLambdaPerRoute) GetEmptyBodyOverride() *types.StringValue {
	if m != nil {
		return m.EmptyBodyOverride
	}
	return nil
}

type AWSLambdaProtocolExtension struct {
	// The host header for AWS this cluster
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// The region for this cluster
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	// The access_key for AWS this cluster
	AccessKey string `protobuf:"bytes,3,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	// The secret_key for AWS this cluster
	SecretKey            string   `protobuf:"bytes,4,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AWSLambdaProtocolExtension) Reset()         { *m = AWSLambdaProtocolExtension{} }
func (m *AWSLambdaProtocolExtension) String() string { return proto.CompactTextString(m) }
func (*AWSLambdaProtocolExtension) ProtoMessage()    {}
func (*AWSLambdaProtocolExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3e66a9e332271a5, []int{1}
}
func (m *AWSLambdaProtocolExtension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AWSLambdaProtocolExtension.Unmarshal(m, b)
}
func (m *AWSLambdaProtocolExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AWSLambdaProtocolExtension.Marshal(b, m, deterministic)
}
func (m *AWSLambdaProtocolExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AWSLambdaProtocolExtension.Merge(m, src)
}
func (m *AWSLambdaProtocolExtension) XXX_Size() int {
	return xxx_messageInfo_AWSLambdaProtocolExtension.Size(m)
}
func (m *AWSLambdaProtocolExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_AWSLambdaProtocolExtension.DiscardUnknown(m)
}

var xxx_messageInfo_AWSLambdaProtocolExtension proto.InternalMessageInfo

func (m *AWSLambdaProtocolExtension) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *AWSLambdaProtocolExtension) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *AWSLambdaProtocolExtension) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *AWSLambdaProtocolExtension) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

type AWSLambdaConfig struct {
	// Use AWS default credentials chain to get credentials.
	// This will search environment variables, ECS metadata and instance metadata
	// to get the credentials. credentials will be rotated automatically.
	//
	// If credentials are provided on the cluster (using the
	// AWSLambdaProtocolExtension), it will override these credentials. This
	// defaults to false, but may change in the future to true.
	UseDefaultCredentials *types.BoolValue `protobuf:"bytes,1,opt,name=use_default_credentials,json=useDefaultCredentials,proto3" json:"use_default_credentials,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}         `json:"-"`
	XXX_unrecognized      []byte           `json:"-"`
	XXX_sizecache         int32            `json:"-"`
}

func (m *AWSLambdaConfig) Reset()         { *m = AWSLambdaConfig{} }
func (m *AWSLambdaConfig) String() string { return proto.CompactTextString(m) }
func (*AWSLambdaConfig) ProtoMessage()    {}
func (*AWSLambdaConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3e66a9e332271a5, []int{2}
}
func (m *AWSLambdaConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AWSLambdaConfig.Unmarshal(m, b)
}
func (m *AWSLambdaConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AWSLambdaConfig.Marshal(b, m, deterministic)
}
func (m *AWSLambdaConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AWSLambdaConfig.Merge(m, src)
}
func (m *AWSLambdaConfig) XXX_Size() int {
	return xxx_messageInfo_AWSLambdaConfig.Size(m)
}
func (m *AWSLambdaConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AWSLambdaConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AWSLambdaConfig proto.InternalMessageInfo

func (m *AWSLambdaConfig) GetUseDefaultCredentials() *types.BoolValue {
	if m != nil {
		return m.UseDefaultCredentials
	}
	return nil
}

func init() {
	proto.RegisterType((*AWSLambdaPerRoute)(nil), "envoy.config.filter.http.aws_lambda.v2.AWSLambdaPerRoute")
	proto.RegisterType((*AWSLambdaProtocolExtension)(nil), "envoy.config.filter.http.aws_lambda.v2.AWSLambdaProtocolExtension")
	proto.RegisterType((*AWSLambdaConfig)(nil), "envoy.config.filter.http.aws_lambda.v2.AWSLambdaConfig")
}

func init() {
	proto.RegisterFile("projects/gloo/api/external/envoy/extensions/aws/filter.proto", fileDescriptor_a3e66a9e332271a5)
}

var fileDescriptor_a3e66a9e332271a5 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x95, 0x51, 0x06, 0xf5, 0x24, 0xd0, 0x02, 0x68, 0x55, 0x19, 0x50, 0xf5, 0x80, 0x7a,
	0xc1, 0x96, 0x0a, 0x47, 0x2e, 0xcb, 0xe0, 0xd4, 0x49, 0x4c, 0x99, 0x04, 0x12, 0x97, 0xc8, 0x49,
	0x5e, 0x52, 0x33, 0xd7, 0xcf, 0xd8, 0x4e, 0x3b, 0xff, 0x21, 0x48, 0xfc, 0x1d, 0xfc, 0x79, 0x9c,
	0x50, 0xec, 0x96, 0x82, 0x76, 0x80, 0x5b, 0xfc, 0x7d, 0xfe, 0xac, 0xdf, 0x7b, 0xf9, 0xc8, 0x5b,
	0x6d, 0xf0, 0x0b, 0x54, 0xce, 0xb2, 0x56, 0x22, 0x32, 0xae, 0x05, 0x83, 0x1b, 0x07, 0x46, 0x71,
	0xc9, 0x40, 0xad, 0xd1, 0x87, 0xa3, 0xb2, 0x02, 0x95, 0x65, 0x7c, 0x63, 0x59, 0x23, 0xa4, 0x03,
	0x43, 0xb5, 0x41, 0x87, 0xe9, 0xcb, 0x70, 0x85, 0x56, 0xa8, 0x1a, 0xd1, 0xd2, 0xad, 0xb5, 0x74,
	0x4e, 0x53, 0xbe, 0xb1, 0x85, 0xe4, 0xab, 0xb2, 0xe6, 0x74, 0x3d, 0x1f, 0x3f, 0x6f, 0x11, 0x5b,
	0x09, 0x2c, 0xa4, 0xca, 0xae, 0x61, 0x1b, 0xc3, 0xb5, 0x06, 0x63, 0xe3, 0x3b, 0xe3, 0x93, 0x35,
	0x97, 0xa2, 0xe6, 0x0e, 0xd8, 0xee, 0x23, 0x1a, 0xd3, 0x1f, 0x09, 0x39, 0x3e, 0xfb, 0x74, 0x75,
	0x11, 0x5e, 0xba, 0x04, 0x93, 0x63, 0xe7, 0x20, 0x7d, 0x4a, 0x06, 0x8a, 0xaf, 0x60, 0x94, 0x4c,
	0x92, 0xd9, 0x30, 0xbb, 0xf7, 0x33, 0x1b, 0x98, 0x83, 0x49, 0x92, 0x07, 0x31, 0x3d, 0x25, 0xc3,
	0xaf, 0x1d, 0x97, 0xa2, 0x11, 0x60, 0x46, 0x07, 0xfd, 0x8d, 0x7c, 0x2f, 0xa4, 0x8f, 0xc9, 0x5d,
	0x6e, 0xbd, 0xaa, 0x46, 0x77, 0x26, 0xc9, 0xec, 0x7e, 0x1e, 0x0f, 0xe9, 0x05, 0x79, 0x04, 0x2b,
	0xed, 0x7c, 0x51, 0x62, 0xed, 0x0b, 0x5c, 0x83, 0x31, 0xa2, 0x86, 0xd1, 0x60, 0x92, 0xcc, 0x8e,
	0xe6, 0xa7, 0x34, 0xd2, 0xd3, 0x1d, 0x3d, 0xbd, 0x72, 0x46, 0xa8, 0xf6, 0x23, 0x97, 0x1d, 0xe4,
	0xc7, 0x21, 0x98, 0x61, 0xed, 0x3f, 0x6c, 0x63, 0xd3, 0xef, 0x09, 0x19, 0xef, 0xa1, 0xfb, 0x50,
	0x85, 0xf2, 0xfd, 0x6e, 0x8d, 0x3d, 0xfd, 0x12, 0xad, 0xbb, 0x45, 0xdf, 0x8b, 0xe9, 0x0b, 0x72,
	0x68, 0xa0, 0x15, 0xa8, 0x22, 0xfa, 0xde, 0xde, 0xca, 0xe9, 0x33, 0x42, 0x78, 0x55, 0x81, 0xb5,
	0xc5, 0x35, 0xf8, 0x30, 0xc5, 0x30, 0x1f, 0x46, 0x65, 0x01, 0xbe, 0xb7, 0x2d, 0x54, 0x06, 0x5c,
	0xb0, 0x07, 0xd1, 0x8e, 0xca, 0x02, 0xfc, 0x14, 0xc8, 0xc3, 0xdf, 0x64, 0xe7, 0xe1, 0xaf, 0xa5,
	0x39, 0x39, 0xe9, 0x2c, 0x14, 0x35, 0x34, 0xbc, 0x93, 0xae, 0xa8, 0x0c, 0xd4, 0xa0, 0x9c, 0xe0,
	0xd2, 0x06, 0xc2, 0xa3, 0xf9, 0xf8, 0xd6, 0xfc, 0x19, 0xa2, 0x8c, 0xd3, 0x3f, 0xe9, 0x2c, 0xbc,
	0x8b, 0xc9, 0xf3, 0x7d, 0x30, 0xfb, 0x96, 0x90, 0x37, 0x02, 0x69, 0x68, 0x87, 0x36, 0x78, 0xe3,
	0xe9, 0xff, 0x15, 0x25, 0x7b, 0x70, 0xb6, 0xb1, 0x7f, 0xec, 0xed, 0x32, 0xf9, 0xbc, 0x68, 0x85,
	0x5b, 0x76, 0x25, 0xad, 0x70, 0xc5, 0x2c, 0x4a, 0x7c, 0x25, 0x30, 0x56, 0xf5, 0xef, 0xe2, 0xea,
	0xeb, 0xf6, 0xdf, 0xe5, 0x2d, 0x0f, 0xc3, 0x08, 0xaf, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x53,
	0xb2, 0x60, 0xe9, 0xf6, 0x02, 0x00, 0x00,
}

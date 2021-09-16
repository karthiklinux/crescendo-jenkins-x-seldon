// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/named_tensor.proto

package protobuf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	framework "github.com/tensorflow/tensorflow/tensorflow/go/core/framework"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A pair of tensor name and tensor values.
type NamedTensorProto struct {
	// Name of the tensor.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The client can populate a TensorProto using a tensorflow::Tensor`, or
	// directly using the protobuf field accessors.
	//
	// The client specifies whether the returned tensor values should be
	// filled tensor fields (float_val, int_val, etc.) or encoded in a
	// compact form in tensor.tensor_content.
	Tensor               *framework.TensorProto `protobuf:"bytes,2,opt,name=tensor,proto3" json:"tensor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *NamedTensorProto) Reset()         { *m = NamedTensorProto{} }
func (m *NamedTensorProto) String() string { return proto.CompactTextString(m) }
func (*NamedTensorProto) ProtoMessage()    {}
func (*NamedTensorProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c12ce8841c69dcd, []int{0}
}

func (m *NamedTensorProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamedTensorProto.Unmarshal(m, b)
}
func (m *NamedTensorProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamedTensorProto.Marshal(b, m, deterministic)
}
func (m *NamedTensorProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamedTensorProto.Merge(m, src)
}
func (m *NamedTensorProto) XXX_Size() int {
	return xxx_messageInfo_NamedTensorProto.Size(m)
}
func (m *NamedTensorProto) XXX_DiscardUnknown() {
	xxx_messageInfo_NamedTensorProto.DiscardUnknown(m)
}

var xxx_messageInfo_NamedTensorProto proto.InternalMessageInfo

func (m *NamedTensorProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NamedTensorProto) GetTensor() *framework.TensorProto {
	if m != nil {
		return m.Tensor
	}
	return nil
}

func init() {
	proto.RegisterType((*NamedTensorProto)(nil), "tensorflow.NamedTensorProto")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/named_tensor.proto", fileDescriptor_5c12ce8841c69dcd)
}

var fileDescriptor_5c12ce8841c69dcd = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2e, 0x49, 0xcd, 0x2b,
	0xce, 0x2f, 0x4a, 0xcb, 0xc9, 0x2f, 0xd7, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0xcf, 0x4b, 0xcc, 0x4d, 0x4d, 0x89, 0x87, 0x48, 0xeb, 0x81, 0x45,
	0x85, 0xb8, 0x10, 0x8a, 0xa5, 0xd4, 0xd0, 0x35, 0xa6, 0x15, 0x25, 0xe6, 0xa6, 0x96, 0xe7, 0x17,
	0x65, 0xeb, 0x23, 0xeb, 0x51, 0x0a, 0xe7, 0x12, 0xf0, 0x03, 0x99, 0x14, 0x02, 0x16, 0x0c, 0x00,
	0x9b, 0x23, 0xc4, 0xc5, 0x02, 0x32, 0x5d, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x16,
	0xd2, 0xe7, 0x62, 0x83, 0xe8, 0x93, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x36, 0x12, 0xd7, 0x43, 0x58,
	0xa0, 0x87, 0xa4, 0x39, 0x08, 0xaa, 0xcc, 0xa9, 0x80, 0x4b, 0x22, 0xbf, 0x28, 0x1d, 0x59, 0x15,
	0xdc, 0x05, 0x4e, 0x82, 0xe8, 0x56, 0x16, 0x07, 0x30, 0x46, 0xd9, 0xa4, 0x67, 0x96, 0x64, 0x94,
	0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x23, 0x39, 0x1e, 0x3b, 0x33, 0x3d, 0x1f, 0x35, 0x38, 0x7e,
	0x30, 0x32, 0x26, 0xb1, 0x81, 0x39, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x66, 0x75, 0xd5,
	0x40, 0x34, 0x01, 0x00, 0x00,
}

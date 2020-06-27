// Code generated by protoc-gen-go. DO NOT EDIT.
// source: processor_message.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type CPU struct {
	// Brand of the CPU
	Brand string `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
	//
	// Name of the CPU
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	NumberCores          uint32   `protobuf:"varint,3,opt,name=number_cores,json=numberCores,proto3" json:"number_cores,omitempty"`
	NumberThreads        uint32   `protobuf:"varint,4,opt,name=number_threads,json=numberThreads,proto3" json:"number_threads,omitempty"`
	MinGhz               float32  `protobuf:"fixed32,5,opt,name=min_ghz,json=minGhz,proto3" json:"min_ghz,omitempty"`
	MaxGhz               float64  `protobuf:"fixed64,6,opt,name=max_ghz,json=maxGhz,proto3" json:"max_ghz,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPU) Reset()         { *m = CPU{} }
func (m *CPU) String() string { return proto.CompactTextString(m) }
func (*CPU) ProtoMessage()    {}
func (*CPU) Descriptor() ([]byte, []int) {
	return fileDescriptor_466578cecc6db379, []int{0}
}

func (m *CPU) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPU.Unmarshal(m, b)
}
func (m *CPU) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPU.Marshal(b, m, deterministic)
}
func (m *CPU) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPU.Merge(m, src)
}
func (m *CPU) XXX_Size() int {
	return xxx_messageInfo_CPU.Size(m)
}
func (m *CPU) XXX_DiscardUnknown() {
	xxx_messageInfo_CPU.DiscardUnknown(m)
}

var xxx_messageInfo_CPU proto.InternalMessageInfo

func (m *CPU) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *CPU) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CPU) GetNumberCores() uint32 {
	if m != nil {
		return m.NumberCores
	}
	return 0
}

func (m *CPU) GetNumberThreads() uint32 {
	if m != nil {
		return m.NumberThreads
	}
	return 0
}

func (m *CPU) GetMinGhz() float32 {
	if m != nil {
		return m.MinGhz
	}
	return 0
}

func (m *CPU) GetMaxGhz() float64 {
	if m != nil {
		return m.MaxGhz
	}
	return 0
}

type GPU struct {
	Brand                string   `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MinGhz               float64  `protobuf:"fixed64,3,opt,name=min_ghz,json=minGhz,proto3" json:"min_ghz,omitempty"`
	MaxGhz               float64  `protobuf:"fixed64,4,opt,name=max_ghz,json=maxGhz,proto3" json:"max_ghz,omitempty"`
	Memory               *Memory  `protobuf:"bytes,5,opt,name=memory,proto3" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GPU) Reset()         { *m = GPU{} }
func (m *GPU) String() string { return proto.CompactTextString(m) }
func (*GPU) ProtoMessage()    {}
func (*GPU) Descriptor() ([]byte, []int) {
	return fileDescriptor_466578cecc6db379, []int{1}
}

func (m *GPU) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GPU.Unmarshal(m, b)
}
func (m *GPU) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GPU.Marshal(b, m, deterministic)
}
func (m *GPU) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GPU.Merge(m, src)
}
func (m *GPU) XXX_Size() int {
	return xxx_messageInfo_GPU.Size(m)
}
func (m *GPU) XXX_DiscardUnknown() {
	xxx_messageInfo_GPU.DiscardUnknown(m)
}

var xxx_messageInfo_GPU proto.InternalMessageInfo

func (m *GPU) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *GPU) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GPU) GetMinGhz() float64 {
	if m != nil {
		return m.MinGhz
	}
	return 0
}

func (m *GPU) GetMaxGhz() float64 {
	if m != nil {
		return m.MaxGhz
	}
	return 0
}

func (m *GPU) GetMemory() *Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func init() {
	proto.RegisterType((*CPU)(nil), "techschool.pcbook.CPU")
	proto.RegisterType((*GPU)(nil), "techschool.pcbook.GPU")
}

func init() { proto.RegisterFile("processor_message.proto", fileDescriptor_466578cecc6db379) }

var fileDescriptor_466578cecc6db379 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0x49, 0xdb, 0x45, 0xcc, 0x9c, 0x60, 0x18, 0x2c, 0x7a, 0xaa, 0x03, 0xa1, 0xa7, 0x82,
	0xfa, 0x06, 0xee, 0xd0, 0x93, 0x20, 0x41, 0x2f, 0x5e, 0x4a, 0x92, 0x85, 0x65, 0x68, 0xf2, 0x95,
	0xa4, 0xc2, 0xec, 0x73, 0xf8, 0x16, 0xbe, 0xa4, 0x98, 0xf4, 0x30, 0xe9, 0xc9, 0x5b, 0xf2, 0xfb,
	0x7e, 0x7c, 0xff, 0x8f, 0x3f, 0x59, 0x75, 0x1e, 0x94, 0x0e, 0x01, 0x7c, 0x6b, 0x75, 0x08, 0x62,
	0xa7, 0xeb, 0xce, 0x43, 0x0f, 0xf4, 0xa2, 0xd7, 0xca, 0x04, 0x65, 0x00, 0xde, 0xeb, 0x4e, 0x49,
	0x80, 0xb7, 0xab, 0xa5, 0xd5, 0x16, 0xfc, 0xe7, 0x5f, 0x71, 0xfd, 0x8d, 0x48, 0xbe, 0x79, 0x7a,
	0xa1, 0x4b, 0x32, 0x93, 0x5e, 0xb8, 0x2d, 0x43, 0x25, 0xaa, 0x4e, 0x79, 0xfa, 0x50, 0x4a, 0x0a,
	0x27, 0xac, 0x66, 0x59, 0x84, 0xf1, 0x4d, 0xaf, 0xc9, 0x99, 0xfb, 0xb0, 0x52, 0xfb, 0x56, 0x81,
	0xd7, 0x81, 0xe5, 0x25, 0xaa, 0x16, 0x7c, 0x9e, 0xd8, 0xe6, 0x17, 0xd1, 0x1b, 0x72, 0x3e, 0x2a,
	0xbd, 0xf1, 0x5a, 0x6c, 0x03, 0x2b, 0xa2, 0xb4, 0x48, 0xf4, 0x39, 0x41, 0xba, 0x22, 0x27, 0x76,
	0xef, 0xda, 0x9d, 0x19, 0xd8, 0xac, 0x44, 0x55, 0xc6, 0xb1, 0xdd, 0xbb, 0xc6, 0x0c, 0x71, 0x20,
	0x0e, 0x71, 0x80, 0x4b, 0x54, 0x21, 0x8e, 0xad, 0x38, 0x34, 0x66, 0x58, 0x7f, 0x21, 0x92, 0x37,
	0xff, 0xba, 0xf6, 0x28, 0x23, 0x1f, 0x57, 0x4d, 0x32, 0x8a, 0xe3, 0x0c, 0x7a, 0x4b, 0x70, 0x6a,
	0x2a, 0x1e, 0x35, 0xbf, 0xbb, 0xac, 0x27, 0x5d, 0xd6, 0x8f, 0x51, 0xe0, 0xa3, 0xf8, 0x50, 0xbc,
	0x66, 0x9d, 0x94, 0x38, 0x36, 0x7a, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x32, 0x02, 0x57, 0x7d,
	0x95, 0x01, 0x00, 0x00,
}
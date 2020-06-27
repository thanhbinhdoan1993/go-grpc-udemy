// Code generated by protoc-gen-go. DO NOT EDIT.
// source: memory_message.proto

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

type Memory_Unit int32

const (
	Memory_UNKNOWN  Memory_Unit = 0
	Memory_BIT      Memory_Unit = 1
	Memory_BYTE     Memory_Unit = 2
	Memory_KILOBYTE Memory_Unit = 3
	Memory_MEGABYTE Memory_Unit = 4
	Memory_GIGABYTE Memory_Unit = 5
	Memory_TERABYTE Memory_Unit = 6
)

var Memory_Unit_name = map[int32]string{
	0: "UNKNOWN",
	1: "BIT",
	2: "BYTE",
	3: "KILOBYTE",
	4: "MEGABYTE",
	5: "GIGABYTE",
	6: "TERABYTE",
}

var Memory_Unit_value = map[string]int32{
	"UNKNOWN":  0,
	"BIT":      1,
	"BYTE":     2,
	"KILOBYTE": 3,
	"MEGABYTE": 4,
	"GIGABYTE": 5,
	"TERABYTE": 6,
}

func (x Memory_Unit) String() string {
	return proto.EnumName(Memory_Unit_name, int32(x))
}

func (Memory_Unit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c0c7f919ccc765da, []int{0, 0}
}

type Memory struct {
	Value                uint64      `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Unit                 Memory_Unit `protobuf:"varint,2,opt,name=unit,proto3,enum=techschool.pcbook.Memory_Unit" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Memory) Reset()         { *m = Memory{} }
func (m *Memory) String() string { return proto.CompactTextString(m) }
func (*Memory) ProtoMessage()    {}
func (*Memory) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0c7f919ccc765da, []int{0}
}

func (m *Memory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Memory.Unmarshal(m, b)
}
func (m *Memory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Memory.Marshal(b, m, deterministic)
}
func (m *Memory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Memory.Merge(m, src)
}
func (m *Memory) XXX_Size() int {
	return xxx_messageInfo_Memory.Size(m)
}
func (m *Memory) XXX_DiscardUnknown() {
	xxx_messageInfo_Memory.DiscardUnknown(m)
}

var xxx_messageInfo_Memory proto.InternalMessageInfo

func (m *Memory) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Memory) GetUnit() Memory_Unit {
	if m != nil {
		return m.Unit
	}
	return Memory_UNKNOWN
}

func init() {
	proto.RegisterEnum("techschool.pcbook.Memory_Unit", Memory_Unit_name, Memory_Unit_value)
	proto.RegisterType((*Memory)(nil), "techschool.pcbook.Memory")
}

func init() { proto.RegisterFile("memory_message.proto", fileDescriptor_c0c7f919ccc765da) }

var fileDescriptor_c0c7f919ccc765da = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0x4d, 0xcd, 0xcd,
	0x2f, 0xaa, 0x8c, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0x2c, 0x49, 0x4d, 0xce, 0x28, 0x4e, 0xce, 0xc8, 0xcf, 0xcf, 0xd1, 0x2b, 0x48, 0x4e,
	0xca, 0xcf, 0xcf, 0x56, 0xda, 0xc4, 0xc8, 0xc5, 0xe6, 0x0b, 0x56, 0x2b, 0x24, 0xc2, 0xc5, 0x5a,
	0x96, 0x98, 0x53, 0x9a, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x12, 0x04, 0xe1, 0x08, 0x19, 0x71,
	0xb1, 0x94, 0xe6, 0x65, 0x96, 0x48, 0x30, 0x29, 0x30, 0x6a, 0xf0, 0x19, 0xc9, 0xe9, 0x61, 0x18,
	0xa1, 0x07, 0xd1, 0xae, 0x17, 0x9a, 0x97, 0x59, 0x12, 0x04, 0x56, 0xab, 0x14, 0xc7, 0xc5, 0x02,
	0xe2, 0x09, 0x71, 0x73, 0xb1, 0x87, 0xfa, 0x79, 0xfb, 0xf9, 0x87, 0xfb, 0x09, 0x30, 0x08, 0xb1,
	0x73, 0x31, 0x3b, 0x79, 0x86, 0x08, 0x30, 0x0a, 0x71, 0x70, 0xb1, 0x38, 0x45, 0x86, 0xb8, 0x0a,
	0x30, 0x09, 0xf1, 0x70, 0x71, 0x78, 0x7b, 0xfa, 0xf8, 0x83, 0x79, 0xcc, 0x20, 0x9e, 0xaf, 0xab,
	0xbb, 0x23, 0x98, 0xc7, 0x02, 0xe2, 0xb9, 0x7b, 0x42, 0x79, 0xac, 0x20, 0x5e, 0x88, 0x6b, 0x10,
	0x84, 0xc7, 0xe6, 0xc4, 0x12, 0xc5, 0x54, 0x90, 0x94, 0xc4, 0x06, 0xf6, 0x94, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0xb7, 0x16, 0xe2, 0xe4, 0xec, 0x00, 0x00, 0x00,
}

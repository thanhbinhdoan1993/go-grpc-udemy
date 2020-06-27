// Code generated by protoc-gen-go. DO NOT EDIT.
// source: keyboard_message.proto

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

type Keyboard_Layout int32

const (
	Keyboard_UNKNOWN Keyboard_Layout = 0
	Keyboard_QWERTY  Keyboard_Layout = 1
	Keyboard_QWERTZ  Keyboard_Layout = 2
	Keyboard_AZERTY  Keyboard_Layout = 3
)

var Keyboard_Layout_name = map[int32]string{
	0: "UNKNOWN",
	1: "QWERTY",
	2: "QWERTZ",
	3: "AZERTY",
}

var Keyboard_Layout_value = map[string]int32{
	"UNKNOWN": 0,
	"QWERTY":  1,
	"QWERTZ":  2,
	"AZERTY":  3,
}

func (x Keyboard_Layout) String() string {
	return proto.EnumName(Keyboard_Layout_name, int32(x))
}

func (Keyboard_Layout) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8fd4226b8635fe5d, []int{0, 0}
}

type Keyboard struct {
	Layout               Keyboard_Layout `protobuf:"varint,1,opt,name=layout,proto3,enum=techschool.pcbook.Keyboard_Layout" json:"layout,omitempty"`
	Backlit              bool            `protobuf:"varint,2,opt,name=backlit,proto3" json:"backlit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Keyboard) Reset()         { *m = Keyboard{} }
func (m *Keyboard) String() string { return proto.CompactTextString(m) }
func (*Keyboard) ProtoMessage()    {}
func (*Keyboard) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fd4226b8635fe5d, []int{0}
}

func (m *Keyboard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Keyboard.Unmarshal(m, b)
}
func (m *Keyboard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Keyboard.Marshal(b, m, deterministic)
}
func (m *Keyboard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keyboard.Merge(m, src)
}
func (m *Keyboard) XXX_Size() int {
	return xxx_messageInfo_Keyboard.Size(m)
}
func (m *Keyboard) XXX_DiscardUnknown() {
	xxx_messageInfo_Keyboard.DiscardUnknown(m)
}

var xxx_messageInfo_Keyboard proto.InternalMessageInfo

func (m *Keyboard) GetLayout() Keyboard_Layout {
	if m != nil {
		return m.Layout
	}
	return Keyboard_UNKNOWN
}

func (m *Keyboard) GetBacklit() bool {
	if m != nil {
		return m.Backlit
	}
	return false
}

func init() {
	proto.RegisterEnum("techschool.pcbook.Keyboard_Layout", Keyboard_Layout_name, Keyboard_Layout_value)
	proto.RegisterType((*Keyboard)(nil), "techschool.pcbook.Keyboard")
}

func init() { proto.RegisterFile("keyboard_message.proto", fileDescriptor_8fd4226b8635fe5d) }

var fileDescriptor_8fd4226b8635fe5d = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcb, 0x4e, 0xad, 0x4c,
	0xca, 0x4f, 0x2c, 0x4a, 0x89, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x2c, 0x49, 0x4d, 0xce, 0x28, 0x4e, 0xce, 0xc8, 0xcf, 0xcf, 0xd1, 0x2b,
	0x48, 0x4e, 0xca, 0xcf, 0xcf, 0x56, 0x9a, 0xcd, 0xc8, 0xc5, 0xe1, 0x0d, 0x55, 0x2d, 0x64, 0xc5,
	0xc5, 0x96, 0x93, 0x58, 0x99, 0x5f, 0x5a, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x67, 0xa4, 0xa4,
	0x87, 0xa1, 0x41, 0x0f, 0xa6, 0x58, 0xcf, 0x07, 0xac, 0x32, 0x08, 0xaa, 0x43, 0x48, 0x82, 0x8b,
	0x3d, 0x29, 0x31, 0x39, 0x3b, 0x27, 0xb3, 0x44, 0x82, 0x49, 0x81, 0x51, 0x83, 0x23, 0x08, 0xc6,
	0x55, 0xb2, 0xe4, 0x62, 0x83, 0xa8, 0x15, 0xe2, 0xe6, 0x62, 0x0f, 0xf5, 0xf3, 0xf6, 0xf3, 0x0f,
	0xf7, 0x13, 0x60, 0x10, 0xe2, 0xe2, 0x62, 0x0b, 0x0c, 0x77, 0x0d, 0x0a, 0x89, 0x14, 0x60, 0x84,
	0xb3, 0xa3, 0x04, 0x98, 0x40, 0x6c, 0xc7, 0x28, 0xb0, 0x38, 0xb3, 0x13, 0x4b, 0x14, 0x53, 0x41,
	0x52, 0x12, 0x1b, 0xd8, 0xf5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0x9d, 0x94, 0x5f,
	0xd7, 0x00, 0x00, 0x00,
}

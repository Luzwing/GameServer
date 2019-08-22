// Code generated by protoc-gen-go. DO NOT EDIT.
// source: operation.proto

package operation

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

type C2GS_Operation struct {
	Opcodes              []int32   `protobuf:"varint,1,rep,packed,name=opcodes,proto3" json:"opcodes,omitempty"`
	CurrentFrame         int32     `protobuf:"varint,2,opt,name=currentFrame,proto3" json:"currentFrame,omitempty"`
	PlayerId             int32     `protobuf:"varint,3,opt,name=playerId,proto3" json:"playerId,omitempty"`
	Movement             []float32 `protobuf:"fixed32,4,rep,packed,name=movement,proto3" json:"movement,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *C2GS_Operation) Reset()         { *m = C2GS_Operation{} }
func (m *C2GS_Operation) String() string { return proto.CompactTextString(m) }
func (*C2GS_Operation) ProtoMessage()    {}
func (*C2GS_Operation) Descriptor() ([]byte, []int) {
	return fileDescriptor_619dee0fded31cb3, []int{0}
}

func (m *C2GS_Operation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2GS_Operation.Unmarshal(m, b)
}
func (m *C2GS_Operation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2GS_Operation.Marshal(b, m, deterministic)
}
func (m *C2GS_Operation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2GS_Operation.Merge(m, src)
}
func (m *C2GS_Operation) XXX_Size() int {
	return xxx_messageInfo_C2GS_Operation.Size(m)
}
func (m *C2GS_Operation) XXX_DiscardUnknown() {
	xxx_messageInfo_C2GS_Operation.DiscardUnknown(m)
}

var xxx_messageInfo_C2GS_Operation proto.InternalMessageInfo

func (m *C2GS_Operation) GetOpcodes() []int32 {
	if m != nil {
		return m.Opcodes
	}
	return nil
}

func (m *C2GS_Operation) GetCurrentFrame() int32 {
	if m != nil {
		return m.CurrentFrame
	}
	return 0
}

func (m *C2GS_Operation) GetPlayerId() int32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *C2GS_Operation) GetMovement() []float32 {
	if m != nil {
		return m.Movement
	}
	return nil
}

type GS2C_Operation struct {
	Operations           []*C2GS_Operation `protobuf:"bytes,1,rep,name=operations,proto3" json:"operations,omitempty"`
	ServerFrame          int32             `protobuf:"varint,2,opt,name=serverFrame,proto3" json:"serverFrame,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GS2C_Operation) Reset()         { *m = GS2C_Operation{} }
func (m *GS2C_Operation) String() string { return proto.CompactTextString(m) }
func (*GS2C_Operation) ProtoMessage()    {}
func (*GS2C_Operation) Descriptor() ([]byte, []int) {
	return fileDescriptor_619dee0fded31cb3, []int{1}
}

func (m *GS2C_Operation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GS2C_Operation.Unmarshal(m, b)
}
func (m *GS2C_Operation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GS2C_Operation.Marshal(b, m, deterministic)
}
func (m *GS2C_Operation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GS2C_Operation.Merge(m, src)
}
func (m *GS2C_Operation) XXX_Size() int {
	return xxx_messageInfo_GS2C_Operation.Size(m)
}
func (m *GS2C_Operation) XXX_DiscardUnknown() {
	xxx_messageInfo_GS2C_Operation.DiscardUnknown(m)
}

var xxx_messageInfo_GS2C_Operation proto.InternalMessageInfo

func (m *GS2C_Operation) GetOperations() []*C2GS_Operation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *GS2C_Operation) GetServerFrame() int32 {
	if m != nil {
		return m.ServerFrame
	}
	return 0
}

//七分钟通知
type GS2C_SevenNote struct {
	RoomId               int32    `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GS2C_SevenNote) Reset()         { *m = GS2C_SevenNote{} }
func (m *GS2C_SevenNote) String() string { return proto.CompactTextString(m) }
func (*GS2C_SevenNote) ProtoMessage()    {}
func (*GS2C_SevenNote) Descriptor() ([]byte, []int) {
	return fileDescriptor_619dee0fded31cb3, []int{2}
}

func (m *GS2C_SevenNote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GS2C_SevenNote.Unmarshal(m, b)
}
func (m *GS2C_SevenNote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GS2C_SevenNote.Marshal(b, m, deterministic)
}
func (m *GS2C_SevenNote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GS2C_SevenNote.Merge(m, src)
}
func (m *GS2C_SevenNote) XXX_Size() int {
	return xxx_messageInfo_GS2C_SevenNote.Size(m)
}
func (m *GS2C_SevenNote) XXX_DiscardUnknown() {
	xxx_messageInfo_GS2C_SevenNote.DiscardUnknown(m)
}

var xxx_messageInfo_GS2C_SevenNote proto.InternalMessageInfo

func (m *GS2C_SevenNote) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

//游戏中死亡退出
type C2GS_PlayerQuit struct {
	RoomId               int32    `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	PlayerId             int32    `protobuf:"varint,2,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2GS_PlayerQuit) Reset()         { *m = C2GS_PlayerQuit{} }
func (m *C2GS_PlayerQuit) String() string { return proto.CompactTextString(m) }
func (*C2GS_PlayerQuit) ProtoMessage()    {}
func (*C2GS_PlayerQuit) Descriptor() ([]byte, []int) {
	return fileDescriptor_619dee0fded31cb3, []int{3}
}

func (m *C2GS_PlayerQuit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2GS_PlayerQuit.Unmarshal(m, b)
}
func (m *C2GS_PlayerQuit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2GS_PlayerQuit.Marshal(b, m, deterministic)
}
func (m *C2GS_PlayerQuit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2GS_PlayerQuit.Merge(m, src)
}
func (m *C2GS_PlayerQuit) XXX_Size() int {
	return xxx_messageInfo_C2GS_PlayerQuit.Size(m)
}
func (m *C2GS_PlayerQuit) XXX_DiscardUnknown() {
	xxx_messageInfo_C2GS_PlayerQuit.DiscardUnknown(m)
}

var xxx_messageInfo_C2GS_PlayerQuit proto.InternalMessageInfo

func (m *C2GS_PlayerQuit) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *C2GS_PlayerQuit) GetPlayerId() int32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

//游戏结束
type C2GS_GameEnd struct {
	RoomId               int32    `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2GS_GameEnd) Reset()         { *m = C2GS_GameEnd{} }
func (m *C2GS_GameEnd) String() string { return proto.CompactTextString(m) }
func (*C2GS_GameEnd) ProtoMessage()    {}
func (*C2GS_GameEnd) Descriptor() ([]byte, []int) {
	return fileDescriptor_619dee0fded31cb3, []int{4}
}

func (m *C2GS_GameEnd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2GS_GameEnd.Unmarshal(m, b)
}
func (m *C2GS_GameEnd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2GS_GameEnd.Marshal(b, m, deterministic)
}
func (m *C2GS_GameEnd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2GS_GameEnd.Merge(m, src)
}
func (m *C2GS_GameEnd) XXX_Size() int {
	return xxx_messageInfo_C2GS_GameEnd.Size(m)
}
func (m *C2GS_GameEnd) XXX_DiscardUnknown() {
	xxx_messageInfo_C2GS_GameEnd.DiscardUnknown(m)
}

var xxx_messageInfo_C2GS_GameEnd proto.InternalMessageInfo

func (m *C2GS_GameEnd) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func init() {
	proto.RegisterType((*C2GS_Operation)(nil), "C2GS_Operation")
	proto.RegisterType((*GS2C_Operation)(nil), "GS2C_Operation")
	proto.RegisterType((*GS2C_SevenNote)(nil), "GS2C_SevenNote")
	proto.RegisterType((*C2GS_PlayerQuit)(nil), "C2GS_PlayerQuit")
	proto.RegisterType((*C2GS_GameEnd)(nil), "C2GS_GameEnd")
}

func init() { proto.RegisterFile("operation.proto", fileDescriptor_619dee0fded31cb3) }

var fileDescriptor_619dee0fded31cb3 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4d, 0x6a, 0xc3, 0x30,
	0x10, 0x85, 0xb1, 0xdd, 0xa4, 0x65, 0x12, 0x62, 0xd0, 0xa2, 0x88, 0xae, 0x8c, 0x16, 0xc5, 0xab,
	0x14, 0xdc, 0x23, 0x84, 0xd4, 0x64, 0xd3, 0x1f, 0xfb, 0x00, 0xc5, 0xb5, 0x67, 0x11, 0xa8, 0x34,
	0x62, 0xa2, 0x18, 0x7a, 0x81, 0x9e, 0xbb, 0x54, 0x49, 0x8c, 0xb4, 0xc8, 0xf2, 0xcd, 0x7b, 0xcc,
	0x7c, 0x8f, 0x81, 0x9c, 0x2c, 0x72, 0xe7, 0xf6, 0x64, 0xd6, 0x96, 0xc9, 0x91, 0xfa, 0x4d, 0x60,
	0xb5, 0xa9, 0xea, 0xf6, 0xf3, 0xed, 0x62, 0x08, 0x09, 0xb7, 0x64, 0x7b, 0x1a, 0xf0, 0x20, 0x93,
	0x22, 0x2b, 0x67, 0xcd, 0x45, 0x0a, 0x05, 0xcb, 0xfe, 0xc8, 0x8c, 0xc6, 0xbd, 0x70, 0xa7, 0x51,
	0xa6, 0x45, 0x52, 0xce, 0x9a, 0x68, 0x26, 0x1e, 0xe0, 0xce, 0x7e, 0x77, 0x3f, 0xc8, 0xbb, 0x41,
	0x66, 0xde, 0x9f, 0xf4, 0xbf, 0xa7, 0x69, 0x44, 0x8d, 0xc6, 0xc9, 0x9b, 0x22, 0x2b, 0xd3, 0x66,
	0xd2, 0xaa, 0x87, 0x55, 0xdd, 0x56, 0x9b, 0x80, 0xe3, 0x09, 0x60, 0xa2, 0x3d, 0xa1, 0x2c, 0xaa,
	0x7c, 0x1d, 0xc3, 0x36, 0x41, 0x44, 0x14, 0xb0, 0x38, 0x20, 0x8f, 0xc8, 0x21, 0x5d, 0x38, 0x52,
	0xe5, 0xf9, 0x48, 0x8b, 0x23, 0x9a, 0x57, 0x72, 0x28, 0xee, 0x61, 0xce, 0x44, 0x7a, 0x37, 0xc8,
	0xc4, 0xc7, 0xcf, 0x4a, 0x6d, 0x21, 0xf7, 0x97, 0xde, 0x3d, 0xfb, 0xc7, 0x71, 0xef, 0xae, 0x45,
	0xa3, 0xc6, 0x69, 0xdc, 0x58, 0x3d, 0xc2, 0xd2, 0xaf, 0xa9, 0x3b, 0x8d, 0x5b, 0x33, 0x5c, 0xdb,
	0xf1, 0x35, 0xf7, 0xdf, 0x78, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x88, 0xd9, 0x93, 0xc2, 0xa0,
	0x01, 0x00, 0x00,
}
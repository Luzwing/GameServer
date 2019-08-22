// Code generated by protoc-gen-go. DO NOT EDIT.
// source: roomMessage.proto

package roomMessage

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

//status 0：未在房间中
//1：未就绪
//2：已准备
//3：选好英雄
//4：游戏进行中
type Player struct {
	PlayerId             int32    `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	PlayName             string   `protobuf:"bytes,2,opt,name=playName,proto3" json:"playName,omitempty"`
	PlayerStatus         int32    `protobuf:"varint,3,opt,name=playerStatus,proto3" json:"playerStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2d5dd4590280ee, []int{0}
}

func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetPlayerId() int32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *Player) GetPlayName() string {
	if m != nil {
		return m.PlayName
	}
	return ""
}

func (m *Player) GetPlayerStatus() int32 {
	if m != nil {
		return m.PlayerStatus
	}
	return 0
}

type Room struct {
	RoomId               int32     `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	Players              []*Player `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Room) Reset()         { *m = Room{} }
func (m *Room) String() string { return proto.CompactTextString(m) }
func (*Room) ProtoMessage()    {}
func (*Room) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2d5dd4590280ee, []int{1}
}

func (m *Room) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Room.Unmarshal(m, b)
}
func (m *Room) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Room.Marshal(b, m, deterministic)
}
func (m *Room) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Room.Merge(m, src)
}
func (m *Room) XXX_Size() int {
	return xxx_messageInfo_Room.Size(m)
}
func (m *Room) XXX_DiscardUnknown() {
	xxx_messageInfo_Room.DiscardUnknown(m)
}

var xxx_messageInfo_Room proto.InternalMessageInfo

func (m *Room) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *Room) GetPlayers() []*Player {
	if m != nil {
		return m.Players
	}
	return nil
}

//房间列表获取,只发送给请求客户端
type C2GS_RoomGet struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2GS_RoomGet) Reset()         { *m = C2GS_RoomGet{} }
func (m *C2GS_RoomGet) String() string { return proto.CompactTextString(m) }
func (*C2GS_RoomGet) ProtoMessage()    {}
func (*C2GS_RoomGet) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2d5dd4590280ee, []int{2}
}

func (m *C2GS_RoomGet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2GS_RoomGet.Unmarshal(m, b)
}
func (m *C2GS_RoomGet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2GS_RoomGet.Marshal(b, m, deterministic)
}
func (m *C2GS_RoomGet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2GS_RoomGet.Merge(m, src)
}
func (m *C2GS_RoomGet) XXX_Size() int {
	return xxx_messageInfo_C2GS_RoomGet.Size(m)
}
func (m *C2GS_RoomGet) XXX_DiscardUnknown() {
	xxx_messageInfo_C2GS_RoomGet.DiscardUnknown(m)
}

var xxx_messageInfo_C2GS_RoomGet proto.InternalMessageInfo

type GS2C_RoomGet struct {
	Rooms                []*Room  `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GS2C_RoomGet) Reset()         { *m = GS2C_RoomGet{} }
func (m *GS2C_RoomGet) String() string { return proto.CompactTextString(m) }
func (*GS2C_RoomGet) ProtoMessage()    {}
func (*GS2C_RoomGet) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2d5dd4590280ee, []int{3}
}

func (m *GS2C_RoomGet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GS2C_RoomGet.Unmarshal(m, b)
}
func (m *GS2C_RoomGet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GS2C_RoomGet.Marshal(b, m, deterministic)
}
func (m *GS2C_RoomGet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GS2C_RoomGet.Merge(m, src)
}
func (m *GS2C_RoomGet) XXX_Size() int {
	return xxx_messageInfo_GS2C_RoomGet.Size(m)
}
func (m *GS2C_RoomGet) XXX_DiscardUnknown() {
	xxx_messageInfo_GS2C_RoomGet.DiscardUnknown(m)
}

var xxx_messageInfo_GS2C_RoomGet proto.InternalMessageInfo

func (m *GS2C_RoomGet) GetRooms() []*Room {
	if m != nil {
		return m.Rooms
	}
	return nil
}

//房间创建，只发送给请求客户端
type C2GS_RoomCrt struct {
	RoomOwner            *Player  `protobuf:"bytes,1,opt,name=roomOwner,proto3" json:"roomOwner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2GS_RoomCrt) Reset()         { *m = C2GS_RoomCrt{} }
func (m *C2GS_RoomCrt) String() string { return proto.CompactTextString(m) }
func (*C2GS_RoomCrt) ProtoMessage()    {}
func (*C2GS_RoomCrt) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2d5dd4590280ee, []int{4}
}

func (m *C2GS_RoomCrt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2GS_RoomCrt.Unmarshal(m, b)
}
func (m *C2GS_RoomCrt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2GS_RoomCrt.Marshal(b, m, deterministic)
}
func (m *C2GS_RoomCrt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2GS_RoomCrt.Merge(m, src)
}
func (m *C2GS_RoomCrt) XXX_Size() int {
	return xxx_messageInfo_C2GS_RoomCrt.Size(m)
}
func (m *C2GS_RoomCrt) XXX_DiscardUnknown() {
	xxx_messageInfo_C2GS_RoomCrt.DiscardUnknown(m)
}

var xxx_messageInfo_C2GS_RoomCrt proto.InternalMessageInfo

func (m *C2GS_RoomCrt) GetRoomOwner() *Player {
	if m != nil {
		return m.RoomOwner
	}
	return nil
}

type GS2C_RoomCrt struct {
	ResStatus            int32    `protobuf:"varint,1,opt,name=resStatus,proto3" json:"resStatus,omitempty"`
	RoomId               int32    `protobuf:"varint,2,opt,name=roomId,proto3" json:"roomId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GS2C_RoomCrt) Reset()         { *m = GS2C_RoomCrt{} }
func (m *GS2C_RoomCrt) String() string { return proto.CompactTextString(m) }
func (*GS2C_RoomCrt) ProtoMessage()    {}
func (*GS2C_RoomCrt) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2d5dd4590280ee, []int{5}
}

func (m *GS2C_RoomCrt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GS2C_RoomCrt.Unmarshal(m, b)
}
func (m *GS2C_RoomCrt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GS2C_RoomCrt.Marshal(b, m, deterministic)
}
func (m *GS2C_RoomCrt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GS2C_RoomCrt.Merge(m, src)
}
func (m *GS2C_RoomCrt) XXX_Size() int {
	return xxx_messageInfo_GS2C_RoomCrt.Size(m)
}
func (m *GS2C_RoomCrt) XXX_DiscardUnknown() {
	xxx_messageInfo_GS2C_RoomCrt.DiscardUnknown(m)
}

var xxx_messageInfo_GS2C_RoomCrt proto.InternalMessageInfo

func (m *GS2C_RoomCrt) GetResStatus() int32 {
	if m != nil {
		return m.ResStatus
	}
	return 0
}

func (m *GS2C_RoomCrt) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

//加入房间
//加入成功给房间内所有客户端发送消息，消息包括
//加入失败则给请求客户端发送？？
type C2GS_RoomJoin struct {
	RoomId               int32    `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	PlayerId             int32    `protobuf:"varint,2,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2GS_RoomJoin) Reset()         { *m = C2GS_RoomJoin{} }
func (m *C2GS_RoomJoin) String() string { return proto.CompactTextString(m) }
func (*C2GS_RoomJoin) ProtoMessage()    {}
func (*C2GS_RoomJoin) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2d5dd4590280ee, []int{6}
}

func (m *C2GS_RoomJoin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2GS_RoomJoin.Unmarshal(m, b)
}
func (m *C2GS_RoomJoin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2GS_RoomJoin.Marshal(b, m, deterministic)
}
func (m *C2GS_RoomJoin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2GS_RoomJoin.Merge(m, src)
}
func (m *C2GS_RoomJoin) XXX_Size() int {
	return xxx_messageInfo_C2GS_RoomJoin.Size(m)
}
func (m *C2GS_RoomJoin) XXX_DiscardUnknown() {
	xxx_messageInfo_C2GS_RoomJoin.DiscardUnknown(m)
}

var xxx_messageInfo_C2GS_RoomJoin proto.InternalMessageInfo

func (m *C2GS_RoomJoin) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *C2GS_RoomJoin) GetPlayerId() int32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

type GS2C_RoomJoin struct {
	ResStatus            int32    `protobuf:"varint,1,opt,name=resStatus,proto3" json:"resStatus,omitempty"`
	RoomInfo             *Room    `protobuf:"bytes,2,opt,name=roomInfo,proto3" json:"roomInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GS2C_RoomJoin) Reset()         { *m = GS2C_RoomJoin{} }
func (m *GS2C_RoomJoin) String() string { return proto.CompactTextString(m) }
func (*GS2C_RoomJoin) ProtoMessage()    {}
func (*GS2C_RoomJoin) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2d5dd4590280ee, []int{7}
}

func (m *GS2C_RoomJoin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GS2C_RoomJoin.Unmarshal(m, b)
}
func (m *GS2C_RoomJoin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GS2C_RoomJoin.Marshal(b, m, deterministic)
}
func (m *GS2C_RoomJoin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GS2C_RoomJoin.Merge(m, src)
}
func (m *GS2C_RoomJoin) XXX_Size() int {
	return xxx_messageInfo_GS2C_RoomJoin.Size(m)
}
func (m *GS2C_RoomJoin) XXX_DiscardUnknown() {
	xxx_messageInfo_GS2C_RoomJoin.DiscardUnknown(m)
}

var xxx_messageInfo_GS2C_RoomJoin proto.InternalMessageInfo

func (m *GS2C_RoomJoin) GetResStatus() int32 {
	if m != nil {
		return m.ResStatus
	}
	return 0
}

func (m *GS2C_RoomJoin) GetRoomInfo() *Room {
	if m != nil {
		return m.RoomInfo
	}
	return nil
}

//退出房间
type C2GS_RoomQuit struct {
	RoomId               int32    `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	PlayerId             int32    `protobuf:"varint,2,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2GS_RoomQuit) Reset()         { *m = C2GS_RoomQuit{} }
func (m *C2GS_RoomQuit) String() string { return proto.CompactTextString(m) }
func (*C2GS_RoomQuit) ProtoMessage()    {}
func (*C2GS_RoomQuit) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2d5dd4590280ee, []int{8}
}

func (m *C2GS_RoomQuit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2GS_RoomQuit.Unmarshal(m, b)
}
func (m *C2GS_RoomQuit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2GS_RoomQuit.Marshal(b, m, deterministic)
}
func (m *C2GS_RoomQuit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2GS_RoomQuit.Merge(m, src)
}
func (m *C2GS_RoomQuit) XXX_Size() int {
	return xxx_messageInfo_C2GS_RoomQuit.Size(m)
}
func (m *C2GS_RoomQuit) XXX_DiscardUnknown() {
	xxx_messageInfo_C2GS_RoomQuit.DiscardUnknown(m)
}

var xxx_messageInfo_C2GS_RoomQuit proto.InternalMessageInfo

func (m *C2GS_RoomQuit) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *C2GS_RoomQuit) GetPlayerId() int32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

type GS2C_RoomQuit struct {
	ResStatus            int32    `protobuf:"varint,1,opt,name=resStatus,proto3" json:"resStatus,omitempty"`
	PlayerId             int32    `protobuf:"varint,2,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GS2C_RoomQuit) Reset()         { *m = GS2C_RoomQuit{} }
func (m *GS2C_RoomQuit) String() string { return proto.CompactTextString(m) }
func (*GS2C_RoomQuit) ProtoMessage()    {}
func (*GS2C_RoomQuit) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2d5dd4590280ee, []int{9}
}

func (m *GS2C_RoomQuit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GS2C_RoomQuit.Unmarshal(m, b)
}
func (m *GS2C_RoomQuit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GS2C_RoomQuit.Marshal(b, m, deterministic)
}
func (m *GS2C_RoomQuit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GS2C_RoomQuit.Merge(m, src)
}
func (m *GS2C_RoomQuit) XXX_Size() int {
	return xxx_messageInfo_GS2C_RoomQuit.Size(m)
}
func (m *GS2C_RoomQuit) XXX_DiscardUnknown() {
	xxx_messageInfo_GS2C_RoomQuit.DiscardUnknown(m)
}

var xxx_messageInfo_GS2C_RoomQuit proto.InternalMessageInfo

func (m *GS2C_RoomQuit) GetResStatus() int32 {
	if m != nil {
		return m.ResStatus
	}
	return 0
}

func (m *GS2C_RoomQuit) GetPlayerId() int32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

//玩家准备
type C2GS_PlayerReady struct {
	PlayerId             int32    `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	RoomId               int32    `protobuf:"varint,2,opt,name=roomId,proto3" json:"roomId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2GS_PlayerReady) Reset()         { *m = C2GS_PlayerReady{} }
func (m *C2GS_PlayerReady) String() string { return proto.CompactTextString(m) }
func (*C2GS_PlayerReady) ProtoMessage()    {}
func (*C2GS_PlayerReady) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2d5dd4590280ee, []int{10}
}

func (m *C2GS_PlayerReady) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2GS_PlayerReady.Unmarshal(m, b)
}
func (m *C2GS_PlayerReady) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2GS_PlayerReady.Marshal(b, m, deterministic)
}
func (m *C2GS_PlayerReady) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2GS_PlayerReady.Merge(m, src)
}
func (m *C2GS_PlayerReady) XXX_Size() int {
	return xxx_messageInfo_C2GS_PlayerReady.Size(m)
}
func (m *C2GS_PlayerReady) XXX_DiscardUnknown() {
	xxx_messageInfo_C2GS_PlayerReady.DiscardUnknown(m)
}

var xxx_messageInfo_C2GS_PlayerReady proto.InternalMessageInfo

func (m *C2GS_PlayerReady) GetPlayerId() int32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *C2GS_PlayerReady) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

type GS2C_PlayerReady struct {
	ResStatus            int32    `protobuf:"varint,1,opt,name=resStatus,proto3" json:"resStatus,omitempty"`
	PlayerId             int32    `protobuf:"varint,2,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GS2C_PlayerReady) Reset()         { *m = GS2C_PlayerReady{} }
func (m *GS2C_PlayerReady) String() string { return proto.CompactTextString(m) }
func (*GS2C_PlayerReady) ProtoMessage()    {}
func (*GS2C_PlayerReady) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2d5dd4590280ee, []int{11}
}

func (m *GS2C_PlayerReady) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GS2C_PlayerReady.Unmarshal(m, b)
}
func (m *GS2C_PlayerReady) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GS2C_PlayerReady.Marshal(b, m, deterministic)
}
func (m *GS2C_PlayerReady) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GS2C_PlayerReady.Merge(m, src)
}
func (m *GS2C_PlayerReady) XXX_Size() int {
	return xxx_messageInfo_GS2C_PlayerReady.Size(m)
}
func (m *GS2C_PlayerReady) XXX_DiscardUnknown() {
	xxx_messageInfo_GS2C_PlayerReady.DiscardUnknown(m)
}

var xxx_messageInfo_GS2C_PlayerReady proto.InternalMessageInfo

func (m *GS2C_PlayerReady) GetResStatus() int32 {
	if m != nil {
		return m.ResStatus
	}
	return 0
}

func (m *GS2C_PlayerReady) GetPlayerId() int32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

//进入选择
type C2GS_StartChoose struct {
	RoomId               int32    `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	PlayerId             int32    `protobuf:"varint,2,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2GS_StartChoose) Reset()         { *m = C2GS_StartChoose{} }
func (m *C2GS_StartChoose) String() string { return proto.CompactTextString(m) }
func (*C2GS_StartChoose) ProtoMessage()    {}
func (*C2GS_StartChoose) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2d5dd4590280ee, []int{12}
}

func (m *C2GS_StartChoose) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2GS_StartChoose.Unmarshal(m, b)
}
func (m *C2GS_StartChoose) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2GS_StartChoose.Marshal(b, m, deterministic)
}
func (m *C2GS_StartChoose) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2GS_StartChoose.Merge(m, src)
}
func (m *C2GS_StartChoose) XXX_Size() int {
	return xxx_messageInfo_C2GS_StartChoose.Size(m)
}
func (m *C2GS_StartChoose) XXX_DiscardUnknown() {
	xxx_messageInfo_C2GS_StartChoose.DiscardUnknown(m)
}

var xxx_messageInfo_C2GS_StartChoose proto.InternalMessageInfo

func (m *C2GS_StartChoose) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *C2GS_StartChoose) GetPlayerId() int32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

type GS2C_StartChoose struct {
	Ltype                int32    `protobuf:"varint,1,opt,name=ltype,proto3" json:"ltype,omitempty"`
	RoomId               int32    `protobuf:"varint,2,opt,name=roomId,proto3" json:"roomId,omitempty"`
	PlayerId             int32    `protobuf:"varint,3,opt,name=playerId,proto3" json:"playerId,omitempty"`
	ResStatus            int32    `protobuf:"varint,4,opt,name=resStatus,proto3" json:"resStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GS2C_StartChoose) Reset()         { *m = GS2C_StartChoose{} }
func (m *GS2C_StartChoose) String() string { return proto.CompactTextString(m) }
func (*GS2C_StartChoose) ProtoMessage()    {}
func (*GS2C_StartChoose) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2d5dd4590280ee, []int{13}
}

func (m *GS2C_StartChoose) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GS2C_StartChoose.Unmarshal(m, b)
}
func (m *GS2C_StartChoose) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GS2C_StartChoose.Marshal(b, m, deterministic)
}
func (m *GS2C_StartChoose) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GS2C_StartChoose.Merge(m, src)
}
func (m *GS2C_StartChoose) XXX_Size() int {
	return xxx_messageInfo_GS2C_StartChoose.Size(m)
}
func (m *GS2C_StartChoose) XXX_DiscardUnknown() {
	xxx_messageInfo_GS2C_StartChoose.DiscardUnknown(m)
}

var xxx_messageInfo_GS2C_StartChoose proto.InternalMessageInfo

func (m *GS2C_StartChoose) GetLtype() int32 {
	if m != nil {
		return m.Ltype
	}
	return 0
}

func (m *GS2C_StartChoose) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *GS2C_StartChoose) GetPlayerId() int32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *GS2C_StartChoose) GetResStatus() int32 {
	if m != nil {
		return m.ResStatus
	}
	return 0
}

//英雄选择
type C2GS_ChooseLegend struct {
	Ltype                int32    `protobuf:"varint,1,opt,name=ltype,proto3" json:"ltype,omitempty"`
	RoomId               int32    `protobuf:"varint,2,opt,name=roomId,proto3" json:"roomId,omitempty"`
	PlayerId             int32    `protobuf:"varint,3,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2GS_ChooseLegend) Reset()         { *m = C2GS_ChooseLegend{} }
func (m *C2GS_ChooseLegend) String() string { return proto.CompactTextString(m) }
func (*C2GS_ChooseLegend) ProtoMessage()    {}
func (*C2GS_ChooseLegend) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2d5dd4590280ee, []int{14}
}

func (m *C2GS_ChooseLegend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2GS_ChooseLegend.Unmarshal(m, b)
}
func (m *C2GS_ChooseLegend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2GS_ChooseLegend.Marshal(b, m, deterministic)
}
func (m *C2GS_ChooseLegend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2GS_ChooseLegend.Merge(m, src)
}
func (m *C2GS_ChooseLegend) XXX_Size() int {
	return xxx_messageInfo_C2GS_ChooseLegend.Size(m)
}
func (m *C2GS_ChooseLegend) XXX_DiscardUnknown() {
	xxx_messageInfo_C2GS_ChooseLegend.DiscardUnknown(m)
}

var xxx_messageInfo_C2GS_ChooseLegend proto.InternalMessageInfo

func (m *C2GS_ChooseLegend) GetLtype() int32 {
	if m != nil {
		return m.Ltype
	}
	return 0
}

func (m *C2GS_ChooseLegend) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *C2GS_ChooseLegend) GetPlayerId() int32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

type GS2C_ChooseLegend struct {
	ResStatus            int32    `protobuf:"varint,1,opt,name=resStatus,proto3" json:"resStatus,omitempty"`
	Ltype                int32    `protobuf:"varint,2,opt,name=ltype,proto3" json:"ltype,omitempty"`
	PlayerId             int32    `protobuf:"varint,3,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GS2C_ChooseLegend) Reset()         { *m = GS2C_ChooseLegend{} }
func (m *GS2C_ChooseLegend) String() string { return proto.CompactTextString(m) }
func (*GS2C_ChooseLegend) ProtoMessage()    {}
func (*GS2C_ChooseLegend) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2d5dd4590280ee, []int{15}
}

func (m *GS2C_ChooseLegend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GS2C_ChooseLegend.Unmarshal(m, b)
}
func (m *GS2C_ChooseLegend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GS2C_ChooseLegend.Marshal(b, m, deterministic)
}
func (m *GS2C_ChooseLegend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GS2C_ChooseLegend.Merge(m, src)
}
func (m *GS2C_ChooseLegend) XXX_Size() int {
	return xxx_messageInfo_GS2C_ChooseLegend.Size(m)
}
func (m *GS2C_ChooseLegend) XXX_DiscardUnknown() {
	xxx_messageInfo_GS2C_ChooseLegend.DiscardUnknown(m)
}

var xxx_messageInfo_GS2C_ChooseLegend proto.InternalMessageInfo

func (m *GS2C_ChooseLegend) GetResStatus() int32 {
	if m != nil {
		return m.ResStatus
	}
	return 0
}

func (m *GS2C_ChooseLegend) GetLtype() int32 {
	if m != nil {
		return m.Ltype
	}
	return 0
}

func (m *GS2C_ChooseLegend) GetPlayerId() int32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

//开始游戏
type GS2C_GameStart struct {
	RoomId               int32    `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GS2C_GameStart) Reset()         { *m = GS2C_GameStart{} }
func (m *GS2C_GameStart) String() string { return proto.CompactTextString(m) }
func (*GS2C_GameStart) ProtoMessage()    {}
func (*GS2C_GameStart) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2d5dd4590280ee, []int{16}
}

func (m *GS2C_GameStart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GS2C_GameStart.Unmarshal(m, b)
}
func (m *GS2C_GameStart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GS2C_GameStart.Marshal(b, m, deterministic)
}
func (m *GS2C_GameStart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GS2C_GameStart.Merge(m, src)
}
func (m *GS2C_GameStart) XXX_Size() int {
	return xxx_messageInfo_GS2C_GameStart.Size(m)
}
func (m *GS2C_GameStart) XXX_DiscardUnknown() {
	xxx_messageInfo_GS2C_GameStart.DiscardUnknown(m)
}

var xxx_messageInfo_GS2C_GameStart proto.InternalMessageInfo

func (m *GS2C_GameStart) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

//状态同步
type C2GS_PlayerStatusSync struct {
	RoomId               int32    `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	PlayerId             int32    `protobuf:"varint,2,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2GS_PlayerStatusSync) Reset()         { *m = C2GS_PlayerStatusSync{} }
func (m *C2GS_PlayerStatusSync) String() string { return proto.CompactTextString(m) }
func (*C2GS_PlayerStatusSync) ProtoMessage()    {}
func (*C2GS_PlayerStatusSync) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2d5dd4590280ee, []int{17}
}

func (m *C2GS_PlayerStatusSync) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2GS_PlayerStatusSync.Unmarshal(m, b)
}
func (m *C2GS_PlayerStatusSync) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2GS_PlayerStatusSync.Marshal(b, m, deterministic)
}
func (m *C2GS_PlayerStatusSync) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2GS_PlayerStatusSync.Merge(m, src)
}
func (m *C2GS_PlayerStatusSync) XXX_Size() int {
	return xxx_messageInfo_C2GS_PlayerStatusSync.Size(m)
}
func (m *C2GS_PlayerStatusSync) XXX_DiscardUnknown() {
	xxx_messageInfo_C2GS_PlayerStatusSync.DiscardUnknown(m)
}

var xxx_messageInfo_C2GS_PlayerStatusSync proto.InternalMessageInfo

func (m *C2GS_PlayerStatusSync) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *C2GS_PlayerStatusSync) GetPlayerId() int32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

type GS2C_PlayerStatusSync struct {
	ResStatus            int32    `protobuf:"varint,1,opt,name=resStatus,proto3" json:"resStatus,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GS2C_PlayerStatusSync) Reset()         { *m = GS2C_PlayerStatusSync{} }
func (m *GS2C_PlayerStatusSync) String() string { return proto.CompactTextString(m) }
func (*GS2C_PlayerStatusSync) ProtoMessage()    {}
func (*GS2C_PlayerStatusSync) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2d5dd4590280ee, []int{18}
}

func (m *GS2C_PlayerStatusSync) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GS2C_PlayerStatusSync.Unmarshal(m, b)
}
func (m *GS2C_PlayerStatusSync) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GS2C_PlayerStatusSync.Marshal(b, m, deterministic)
}
func (m *GS2C_PlayerStatusSync) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GS2C_PlayerStatusSync.Merge(m, src)
}
func (m *GS2C_PlayerStatusSync) XXX_Size() int {
	return xxx_messageInfo_GS2C_PlayerStatusSync.Size(m)
}
func (m *GS2C_PlayerStatusSync) XXX_DiscardUnknown() {
	xxx_messageInfo_GS2C_PlayerStatusSync.DiscardUnknown(m)
}

var xxx_messageInfo_GS2C_PlayerStatusSync proto.InternalMessageInfo

func (m *GS2C_PlayerStatusSync) GetResStatus() int32 {
	if m != nil {
		return m.ResStatus
	}
	return 0
}

func (m *GS2C_PlayerStatusSync) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

//快速匹配
type C2GS_QuickJoin struct {
	PlayerId             int32    `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2GS_QuickJoin) Reset()         { *m = C2GS_QuickJoin{} }
func (m *C2GS_QuickJoin) String() string { return proto.CompactTextString(m) }
func (*C2GS_QuickJoin) ProtoMessage()    {}
func (*C2GS_QuickJoin) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2d5dd4590280ee, []int{19}
}

func (m *C2GS_QuickJoin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2GS_QuickJoin.Unmarshal(m, b)
}
func (m *C2GS_QuickJoin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2GS_QuickJoin.Marshal(b, m, deterministic)
}
func (m *C2GS_QuickJoin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2GS_QuickJoin.Merge(m, src)
}
func (m *C2GS_QuickJoin) XXX_Size() int {
	return xxx_messageInfo_C2GS_QuickJoin.Size(m)
}
func (m *C2GS_QuickJoin) XXX_DiscardUnknown() {
	xxx_messageInfo_C2GS_QuickJoin.DiscardUnknown(m)
}

var xxx_messageInfo_C2GS_QuickJoin proto.InternalMessageInfo

func (m *C2GS_QuickJoin) GetPlayerId() int32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

type GS2C_QuickJoin struct {
	ResStatus            int32    `protobuf:"varint,1,opt,name=resStatus,proto3" json:"resStatus,omitempty"`
	RoomInfo             *Room    `protobuf:"bytes,2,opt,name=roomInfo,proto3" json:"roomInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GS2C_QuickJoin) Reset()         { *m = GS2C_QuickJoin{} }
func (m *GS2C_QuickJoin) String() string { return proto.CompactTextString(m) }
func (*GS2C_QuickJoin) ProtoMessage()    {}
func (*GS2C_QuickJoin) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee2d5dd4590280ee, []int{20}
}

func (m *GS2C_QuickJoin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GS2C_QuickJoin.Unmarshal(m, b)
}
func (m *GS2C_QuickJoin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GS2C_QuickJoin.Marshal(b, m, deterministic)
}
func (m *GS2C_QuickJoin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GS2C_QuickJoin.Merge(m, src)
}
func (m *GS2C_QuickJoin) XXX_Size() int {
	return xxx_messageInfo_GS2C_QuickJoin.Size(m)
}
func (m *GS2C_QuickJoin) XXX_DiscardUnknown() {
	xxx_messageInfo_GS2C_QuickJoin.DiscardUnknown(m)
}

var xxx_messageInfo_GS2C_QuickJoin proto.InternalMessageInfo

func (m *GS2C_QuickJoin) GetResStatus() int32 {
	if m != nil {
		return m.ResStatus
	}
	return 0
}

func (m *GS2C_QuickJoin) GetRoomInfo() *Room {
	if m != nil {
		return m.RoomInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*Player)(nil), "Player")
	proto.RegisterType((*Room)(nil), "Room")
	proto.RegisterType((*C2GS_RoomGet)(nil), "C2GS_RoomGet")
	proto.RegisterType((*GS2C_RoomGet)(nil), "GS2C_RoomGet")
	proto.RegisterType((*C2GS_RoomCrt)(nil), "C2GS_RoomCrt")
	proto.RegisterType((*GS2C_RoomCrt)(nil), "GS2C_RoomCrt")
	proto.RegisterType((*C2GS_RoomJoin)(nil), "C2GS_RoomJoin")
	proto.RegisterType((*GS2C_RoomJoin)(nil), "GS2C_RoomJoin")
	proto.RegisterType((*C2GS_RoomQuit)(nil), "C2GS_RoomQuit")
	proto.RegisterType((*GS2C_RoomQuit)(nil), "GS2C_RoomQuit")
	proto.RegisterType((*C2GS_PlayerReady)(nil), "C2GS_PlayerReady")
	proto.RegisterType((*GS2C_PlayerReady)(nil), "GS2C_PlayerReady")
	proto.RegisterType((*C2GS_StartChoose)(nil), "C2GS_StartChoose")
	proto.RegisterType((*GS2C_StartChoose)(nil), "GS2C_StartChoose")
	proto.RegisterType((*C2GS_ChooseLegend)(nil), "C2GS_ChooseLegend")
	proto.RegisterType((*GS2C_ChooseLegend)(nil), "GS2C_ChooseLegend")
	proto.RegisterType((*GS2C_GameStart)(nil), "GS2C_GameStart")
	proto.RegisterType((*C2GS_PlayerStatusSync)(nil), "C2GS_PlayerStatusSync")
	proto.RegisterType((*GS2C_PlayerStatusSync)(nil), "GS2C_PlayerStatusSync")
	proto.RegisterType((*C2GS_QuickJoin)(nil), "C2GS_QuickJoin")
	proto.RegisterType((*GS2C_QuickJoin)(nil), "GS2C_QuickJoin")
}

func init() { proto.RegisterFile("roomMessage.proto", fileDescriptor_ee2d5dd4590280ee) }

var fileDescriptor_ee2d5dd4590280ee = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x5f, 0x6f, 0xd3, 0x30,
	0x10, 0x57, 0xd3, 0xb5, 0x5b, 0x8e, 0xae, 0x5a, 0x23, 0x86, 0x22, 0xd8, 0x43, 0x67, 0x09, 0x29,
	0x12, 0xa8, 0x0f, 0x45, 0x7c, 0x00, 0x14, 0x44, 0x34, 0x18, 0xb0, 0x26, 0xcf, 0x08, 0x99, 0xc6,
	0x8c, 0x88, 0x25, 0x8e, 0x62, 0x57, 0x28, 0x7c, 0x7a, 0x94, 0x73, 0xea, 0x38, 0x93, 0xe2, 0x87,
	0x96, 0xb7, 0xdc, 0x9d, 0xfd, 0xfb, 0x73, 0x77, 0x0e, 0x2c, 0x2a, 0xce, 0xf3, 0xcf, 0x4c, 0x08,
	0x7a, 0xcf, 0x56, 0x65, 0xc5, 0x25, 0x27, 0x29, 0x4c, 0xef, 0x1e, 0x68, 0xcd, 0x2a, 0xef, 0x39,
	0x9c, 0x95, 0xf8, 0x75, 0x93, 0xfa, 0xa3, 0xe5, 0x28, 0x98, 0xc4, 0x3a, 0xde, 0xd7, 0xbe, 0xd0,
	0x9c, 0xf9, 0xce, 0x72, 0x14, 0xb8, 0xb1, 0x8e, 0x3d, 0x02, 0x33, 0x75, 0x2e, 0x91, 0x54, 0xee,
	0x84, 0x3f, 0xc6, 0xbb, 0xbd, 0x1c, 0x79, 0x07, 0x27, 0x31, 0xe7, 0xb9, 0xf7, 0x0c, 0xa6, 0x8d,
	0x04, 0xcd, 0xd0, 0x46, 0xde, 0x35, 0x9c, 0xaa, 0xf3, 0xc2, 0x77, 0x96, 0xe3, 0xe0, 0xc9, 0xfa,
	0x74, 0xa5, 0x54, 0xc5, 0xfb, 0x3c, 0x99, 0xc3, 0x2c, 0x5c, 0x47, 0xc9, 0xf7, 0x06, 0x27, 0x62,
	0x92, 0xbc, 0x82, 0x59, 0x94, 0xac, 0xc3, 0x7d, 0xec, 0xbd, 0x80, 0x49, 0x03, 0x26, 0xfc, 0x11,
	0x02, 0x4c, 0x56, 0x4d, 0x21, 0x56, 0x39, 0xf2, 0xd6, 0xb8, 0x1c, 0x56, 0xd2, 0x7b, 0x09, 0x6e,
	0x53, 0xf8, 0xfa, 0xa7, 0x60, 0x15, 0x4a, 0x31, 0x18, 0xbb, 0x0a, 0x79, 0x6f, 0x70, 0x34, 0xd7,
	0xae, 0xc0, 0xad, 0x98, 0x68, 0x7d, 0x2a, 0x07, 0x5d, 0xc2, 0x30, 0xe7, 0x98, 0xe6, 0x48, 0x08,
	0xe7, 0x9a, 0xfc, 0x23, 0xcf, 0x8a, 0xc1, 0x2e, 0x98, 0x13, 0x70, 0xfa, 0x13, 0x20, 0x77, 0x70,
	0xae, 0xa5, 0x20, 0x88, 0x5d, 0xcb, 0x35, 0x9c, 0x21, 0x68, 0xf1, 0x93, 0x23, 0x94, 0x6e, 0x88,
	0x4e, 0xf7, 0x64, 0x6d, 0x76, 0x99, 0x3c, 0x48, 0xd6, 0x8d, 0x21, 0x0b, 0x41, 0xec, 0xb2, 0x6c,
	0x50, 0x1f, 0xe0, 0x02, 0xf5, 0xb4, 0x63, 0x60, 0x34, 0xad, 0xad, 0x3b, 0x39, 0xd4, 0xee, 0x5b,
	0xb8, 0x40, 0x49, 0x26, 0xce, 0xf1, 0xaa, 0x12, 0x49, 0x2b, 0x19, 0xfe, 0xe2, 0x5c, 0xb0, 0x83,
	0x1a, 0xf5, 0xb7, 0x55, 0x65, 0xe2, 0x3c, 0x85, 0xc9, 0x83, 0xac, 0x4b, 0xd6, 0xc2, 0xa8, 0x60,
	0xc8, 0x57, 0x0f, 0x7d, 0xfc, 0xa8, 0x17, 0x3d, 0x7f, 0x27, 0x8f, 0xfc, 0x91, 0x6f, 0xb0, 0x40,
	0x0f, 0x8a, 0xf6, 0x96, 0xdd, 0xb3, 0x22, 0xfd, 0x7f, 0xe4, 0x64, 0x0b, 0x0b, 0xb4, 0xd6, 0x83,
	0xb7, 0x77, 0x5c, 0x93, 0x3b, 0x26, 0xb9, 0x8d, 0x24, 0x80, 0x39, 0x92, 0x44, 0x34, 0x67, 0xd8,
	0xc3, 0xa1, 0x29, 0x90, 0x4f, 0x70, 0x69, 0xec, 0x91, 0x22, 0x4c, 0xea, 0x62, 0x7b, 0xd0, 0xd8,
	0x12, 0xb8, 0x34, 0x96, 0xc9, 0x00, 0xb3, 0xfb, 0xbb, 0x02, 0x57, 0x66, 0x39, 0x13, 0x92, 0xe6,
	0x25, 0x62, 0x8e, 0xe3, 0x2e, 0x41, 0x5e, 0xc3, 0x1c, 0x15, 0x6e, 0x76, 0xd9, 0xf6, 0x37, 0x3e,
	0x66, 0xcb, 0x9e, 0x93, 0x4d, 0xeb, 0xbc, 0x3b, 0x7d, 0xec, 0xd3, 0xff, 0x31, 0xc5, 0x7f, 0xff,
	0x9b, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x94, 0x24, 0xc3, 0xc7, 0x10, 0x06, 0x00, 0x00,
}
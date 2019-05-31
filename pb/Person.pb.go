// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Person.proto

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

//返回给玩家上线的ID信息
type SyncPid struct {
	Pid                  int32    `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncPid) Reset()         { *m = SyncPid{} }
func (m *SyncPid) String() string { return proto.CompactTextString(m) }
func (*SyncPid) ProtoMessage()    {}
func (*SyncPid) Descriptor() ([]byte, []int) {
	return fileDescriptor_841ab6396175eaf3, []int{0}
}

func (m *SyncPid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncPid.Unmarshal(m, b)
}
func (m *SyncPid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncPid.Marshal(b, m, deterministic)
}
func (m *SyncPid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncPid.Merge(m, src)
}
func (m *SyncPid) XXX_Size() int {
	return xxx_messageInfo_SyncPid.Size(m)
}
func (m *SyncPid) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncPid.DiscardUnknown(m)
}

var xxx_messageInfo_SyncPid proto.InternalMessageInfo

func (m *SyncPid) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

//返回给上线玩家初始的坐标
type BroadCast struct {
	Pid int32 `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"`
	Tp  int32 `protobuf:"varint,2,opt,name=Tp,proto3" json:"Tp,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*BroadCast_Content
	//	*BroadCast_P
	//	*BroadCast_ActionData
	Data                 isBroadCast_Data `protobuf_oneof:"Data"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BroadCast) Reset()         { *m = BroadCast{} }
func (m *BroadCast) String() string { return proto.CompactTextString(m) }
func (*BroadCast) ProtoMessage()    {}
func (*BroadCast) Descriptor() ([]byte, []int) {
	return fileDescriptor_841ab6396175eaf3, []int{1}
}

func (m *BroadCast) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BroadCast.Unmarshal(m, b)
}
func (m *BroadCast) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BroadCast.Marshal(b, m, deterministic)
}
func (m *BroadCast) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadCast.Merge(m, src)
}
func (m *BroadCast) XXX_Size() int {
	return xxx_messageInfo_BroadCast.Size(m)
}
func (m *BroadCast) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadCast.DiscardUnknown(m)
}

var xxx_messageInfo_BroadCast proto.InternalMessageInfo

func (m *BroadCast) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *BroadCast) GetTp() int32 {
	if m != nil {
		return m.Tp
	}
	return 0
}

type isBroadCast_Data interface {
	isBroadCast_Data()
}

type BroadCast_Content struct {
	Content string `protobuf:"bytes,3,opt,name=Content,proto3,oneof"`
}

type BroadCast_P struct {
	P *Position `protobuf:"bytes,4,opt,name=P,proto3,oneof"`
}

type BroadCast_ActionData struct {
	ActionData int32 `protobuf:"varint,5,opt,name=ActionData,proto3,oneof"`
}

func (*BroadCast_Content) isBroadCast_Data() {}

func (*BroadCast_P) isBroadCast_Data() {}

func (*BroadCast_ActionData) isBroadCast_Data() {}

func (m *BroadCast) GetData() isBroadCast_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *BroadCast) GetContent() string {
	if x, ok := m.GetData().(*BroadCast_Content); ok {
		return x.Content
	}
	return ""
}

func (m *BroadCast) GetP() *Position {
	if x, ok := m.GetData().(*BroadCast_P); ok {
		return x.P
	}
	return nil
}

func (m *BroadCast) GetActionData() int32 {
	if x, ok := m.GetData().(*BroadCast_ActionData); ok {
		return x.ActionData
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*BroadCast) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _BroadCast_OneofMarshaler, _BroadCast_OneofUnmarshaler, _BroadCast_OneofSizer, []interface{}{
		(*BroadCast_Content)(nil),
		(*BroadCast_P)(nil),
		(*BroadCast_ActionData)(nil),
	}
}

func _BroadCast_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*BroadCast)
	// Data
	switch x := m.Data.(type) {
	case *BroadCast_Content:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Content)
	case *BroadCast_P:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.P); err != nil {
			return err
		}
	case *BroadCast_ActionData:
		b.EncodeVarint(5<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.ActionData))
	case nil:
	default:
		return fmt.Errorf("BroadCast.Data has unexpected type %T", x)
	}
	return nil
}

func _BroadCast_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*BroadCast)
	switch tag {
	case 3: // Data.Content
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Data = &BroadCast_Content{x}
		return true, err
	case 4: // Data.P
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Position)
		err := b.DecodeMessage(msg)
		m.Data = &BroadCast_P{msg}
		return true, err
	case 5: // Data.ActionData
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Data = &BroadCast_ActionData{int32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _BroadCast_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*BroadCast)
	// Data
	switch x := m.Data.(type) {
	case *BroadCast_Content:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Content)))
		n += len(x.Content)
	case *BroadCast_P:
		s := proto.Size(x.P)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BroadCast_ActionData:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.ActionData))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

//位置信息
type Position struct {
	X                    float32  `protobuf:"fixed32,1,opt,name=X,proto3" json:"X,omitempty"`
	Y                    float32  `protobuf:"fixed32,2,opt,name=Y,proto3" json:"Y,omitempty"`
	Z                    float32  `protobuf:"fixed32,3,opt,name=Z,proto3" json:"Z,omitempty"`
	V                    float32  `protobuf:"fixed32,4,opt,name=V,proto3" json:"V,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_841ab6396175eaf3, []int{2}
}

func (m *Position) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Position.Unmarshal(m, b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Position.Marshal(b, m, deterministic)
}
func (m *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(m, src)
}
func (m *Position) XXX_Size() int {
	return xxx_messageInfo_Position.Size(m)
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

func (m *Position) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Position) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Position) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

func (m *Position) GetV() float32 {
	if m != nil {
		return m.V
	}
	return 0
}

//聊天信息
type Talk struct {
	Content              string   `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Talk) Reset()         { *m = Talk{} }
func (m *Talk) String() string { return proto.CompactTextString(m) }
func (*Talk) ProtoMessage()    {}
func (*Talk) Descriptor() ([]byte, []int) {
	return fileDescriptor_841ab6396175eaf3, []int{3}
}

func (m *Talk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Talk.Unmarshal(m, b)
}
func (m *Talk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Talk.Marshal(b, m, deterministic)
}
func (m *Talk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Talk.Merge(m, src)
}
func (m *Talk) XXX_Size() int {
	return xxx_messageInfo_Talk.Size(m)
}
func (m *Talk) XXX_DiscardUnknown() {
	xxx_messageInfo_Talk.DiscardUnknown(m)
}

var xxx_messageInfo_Talk proto.InternalMessageInfo

func (m *Talk) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

//告知当前玩家 周边都有哪些玩家的位置信息
type SyncPlayers struct {
	Ps                   []*Player `protobuf:"bytes,1,rep,name=ps,proto3" json:"ps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SyncPlayers) Reset()         { *m = SyncPlayers{} }
func (m *SyncPlayers) String() string { return proto.CompactTextString(m) }
func (*SyncPlayers) ProtoMessage()    {}
func (*SyncPlayers) Descriptor() ([]byte, []int) {
	return fileDescriptor_841ab6396175eaf3, []int{4}
}

func (m *SyncPlayers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncPlayers.Unmarshal(m, b)
}
func (m *SyncPlayers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncPlayers.Marshal(b, m, deterministic)
}
func (m *SyncPlayers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncPlayers.Merge(m, src)
}
func (m *SyncPlayers) XXX_Size() int {
	return xxx_messageInfo_SyncPlayers.Size(m)
}
func (m *SyncPlayers) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncPlayers.DiscardUnknown(m)
}

var xxx_messageInfo_SyncPlayers proto.InternalMessageInfo

func (m *SyncPlayers) GetPs() []*Player {
	if m != nil {
		return m.Ps
	}
	return nil
}

//其中一个玩家的信息
type Player struct {
	Pid                  int32     `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"`
	P                    *Position `protobuf:"bytes,2,opt,name=P,proto3" json:"P,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_841ab6396175eaf3, []int{5}
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

func (m *Player) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *Player) GetP() *Position {
	if m != nil {
		return m.P
	}
	return nil
}

func init() {
	proto.RegisterType((*SyncPid)(nil), "pb.SyncPid")
	proto.RegisterType((*BroadCast)(nil), "pb.BroadCast")
	proto.RegisterType((*Position)(nil), "pb.Position")
	proto.RegisterType((*Talk)(nil), "pb.Talk")
	proto.RegisterType((*SyncPlayers)(nil), "pb.SyncPlayers")
	proto.RegisterType((*Player)(nil), "pb.Player")
}

func init() { proto.RegisterFile("Person.proto", fileDescriptor_841ab6396175eaf3) }

var fileDescriptor_841ab6396175eaf3 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x3f, 0x6b, 0xf3, 0x30,
	0x10, 0x87, 0x73, 0xca, 0xbf, 0xd7, 0x67, 0xf3, 0x52, 0x34, 0x09, 0xb7, 0x83, 0xd0, 0xe4, 0x2e,
	0x1e, 0x52, 0xe8, 0x5e, 0xa7, 0x83, 0x47, 0xa1, 0x9a, 0x90, 0x64, 0x93, 0x63, 0x0f, 0xa6, 0xc1,
	0x12, 0x96, 0x96, 0x7c, 0x8c, 0x7e, 0xe3, 0x22, 0x95, 0xb4, 0x85, 0x74, 0xb2, 0x9f, 0xdf, 0xa1,
	0xd3, 0x73, 0x27, 0xcc, 0x64, 0x3f, 0x39, 0x33, 0x96, 0x76, 0x32, 0xde, 0x50, 0x62, 0x5b, 0x71,
	0x8f, 0xeb, 0xb7, 0xcb, 0x78, 0x92, 0x43, 0x47, 0xef, 0x70, 0x2e, 0x87, 0x8e, 0x01, 0x87, 0x62,
	0xa9, 0xc2, 0xaf, 0xf8, 0x00, 0x4c, 0xaa, 0xc9, 0xe8, 0x6e, 0xab, 0x9d, 0xbf, 0xad, 0xd3, 0xff,
	0x48, 0x1a, 0xcb, 0x48, 0x0c, 0x48, 0x63, 0x69, 0x8e, 0xeb, 0xad, 0x19, 0x7d, 0x3f, 0x7a, 0x36,
	0xe7, 0x50, 0x24, 0xf5, 0x4c, 0x5d, 0x03, 0xfa, 0x80, 0x20, 0xd9, 0x82, 0x43, 0x91, 0x6e, 0xb2,
	0xd2, 0xb6, 0xa5, 0x34, 0x6e, 0xf0, 0x83, 0x19, 0xeb, 0x99, 0x02, 0x49, 0x39, 0xe2, 0xcb, 0x29,
	0xe0, 0xab, 0xf6, 0x9a, 0x2d, 0x43, 0xc7, 0x7a, 0xa6, 0x7e, 0x65, 0xd5, 0x0a, 0x17, 0xe1, 0x2b,
	0x2a, 0xfc, 0x77, 0x3d, 0x4a, 0x33, 0x84, 0x7d, 0xf4, 0x21, 0x0a, 0xf6, 0x81, 0x0e, 0x51, 0x86,
	0x28, 0x38, 0x04, 0x3a, 0x46, 0x0b, 0xa2, 0xe0, 0x18, 0x68, 0x17, 0x6f, 0x27, 0x0a, 0x76, 0x82,
	0xe3, 0xa2, 0xd1, 0xe7, 0x77, 0xca, 0x7e, 0x7c, 0x43, 0x97, 0xe4, 0xdb, 0x56, 0x3c, 0x62, 0x1a,
	0xd7, 0x72, 0xd6, 0x97, 0x7e, 0x72, 0x34, 0x47, 0x62, 0x1d, 0x03, 0x3e, 0x2f, 0xd2, 0x0d, 0x46,
	0xfb, 0x58, 0x50, 0xc4, 0x3a, 0xf1, 0x8c, 0xab, 0x2f, 0xfa, 0x63, 0x41, 0x79, 0x18, 0x9a, 0xdc,
	0x0e, 0xad, 0x40, 0xb6, 0xab, 0xf8, 0x08, 0x4f, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x42, 0x14,
	0x27, 0x21, 0x94, 0x01, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat/chat.proto

package pbChat

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type WSToMsgSvrChatMsg struct {
	SendID               string   `protobuf:"bytes,1,opt,name=SendID,proto3" json:"SendID,omitempty"`
	RecvID               string   `protobuf:"bytes,2,opt,name=RecvID,proto3" json:"RecvID,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
	SendTime             int64    `protobuf:"varint,4,opt,name=SendTime,proto3" json:"SendTime,omitempty"`
	MsgFrom              int32    `protobuf:"varint,5,opt,name=MsgFrom,proto3" json:"MsgFrom,omitempty"`
	ContentType          int32    `protobuf:"varint,6,opt,name=ContentType,proto3" json:"ContentType,omitempty"`
	SessionType          int32    `protobuf:"varint,7,opt,name=SessionType,proto3" json:"SessionType,omitempty"`
	OperationID          string   `protobuf:"bytes,8,opt,name=OperationID,proto3" json:"OperationID,omitempty"`
	MsgID                string   `protobuf:"bytes,9,opt,name=MsgID,proto3" json:"MsgID,omitempty"`
	Token                string   `protobuf:"bytes,10,opt,name=Token,proto3" json:"Token,omitempty"`
	OfflineInfo          string   `protobuf:"bytes,11,opt,name=OfflineInfo,proto3" json:"OfflineInfo,omitempty"`
	Options              string   `protobuf:"bytes,12,opt,name=Options,proto3" json:"Options,omitempty"`
	PlatformID           int32    `protobuf:"varint,13,opt,name=PlatformID,proto3" json:"PlatformID,omitempty"`
	ForceList            []string `protobuf:"bytes,14,rep,name=ForceList,proto3" json:"ForceList,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WSToMsgSvrChatMsg) Reset()         { *m = WSToMsgSvrChatMsg{} }
func (m *WSToMsgSvrChatMsg) String() string { return proto.CompactTextString(m) }
func (*WSToMsgSvrChatMsg) ProtoMessage()    {}
func (*WSToMsgSvrChatMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a49518b351716, []int{0}
}

func (m *WSToMsgSvrChatMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WSToMsgSvrChatMsg.Unmarshal(m, b)
}
func (m *WSToMsgSvrChatMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WSToMsgSvrChatMsg.Marshal(b, m, deterministic)
}
func (m *WSToMsgSvrChatMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WSToMsgSvrChatMsg.Merge(m, src)
}
func (m *WSToMsgSvrChatMsg) XXX_Size() int {
	return xxx_messageInfo_WSToMsgSvrChatMsg.Size(m)
}
func (m *WSToMsgSvrChatMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_WSToMsgSvrChatMsg.DiscardUnknown(m)
}

var xxx_messageInfo_WSToMsgSvrChatMsg proto.InternalMessageInfo

func (m *WSToMsgSvrChatMsg) GetSendID() string {
	if m != nil {
		return m.SendID
	}
	return ""
}

func (m *WSToMsgSvrChatMsg) GetRecvID() string {
	if m != nil {
		return m.RecvID
	}
	return ""
}

func (m *WSToMsgSvrChatMsg) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *WSToMsgSvrChatMsg) GetSendTime() int64 {
	if m != nil {
		return m.SendTime
	}
	return 0
}

func (m *WSToMsgSvrChatMsg) GetMsgFrom() int32 {
	if m != nil {
		return m.MsgFrom
	}
	return 0
}

func (m *WSToMsgSvrChatMsg) GetContentType() int32 {
	if m != nil {
		return m.ContentType
	}
	return 0
}

func (m *WSToMsgSvrChatMsg) GetSessionType() int32 {
	if m != nil {
		return m.SessionType
	}
	return 0
}

func (m *WSToMsgSvrChatMsg) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

func (m *WSToMsgSvrChatMsg) GetMsgID() string {
	if m != nil {
		return m.MsgID
	}
	return ""
}

func (m *WSToMsgSvrChatMsg) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *WSToMsgSvrChatMsg) GetOfflineInfo() string {
	if m != nil {
		return m.OfflineInfo
	}
	return ""
}

func (m *WSToMsgSvrChatMsg) GetOptions() string {
	if m != nil {
		return m.Options
	}
	return ""
}

func (m *WSToMsgSvrChatMsg) GetPlatformID() int32 {
	if m != nil {
		return m.PlatformID
	}
	return 0
}

func (m *WSToMsgSvrChatMsg) GetForceList() []string {
	if m != nil {
		return m.ForceList
	}
	return nil
}

type MsgSvrToPushSvrChatMsg struct {
	SendID               string   `protobuf:"bytes,1,opt,name=SendID,proto3" json:"SendID,omitempty"`
	RecvID               string   `protobuf:"bytes,2,opt,name=RecvID,proto3" json:"RecvID,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
	RecvSeq              int64    `protobuf:"varint,4,opt,name=RecvSeq,proto3" json:"RecvSeq,omitempty"`
	SendTime             int64    `protobuf:"varint,5,opt,name=SendTime,proto3" json:"SendTime,omitempty"`
	MsgFrom              int32    `protobuf:"varint,6,opt,name=MsgFrom,proto3" json:"MsgFrom,omitempty"`
	ContentType          int32    `protobuf:"varint,7,opt,name=ContentType,proto3" json:"ContentType,omitempty"`
	SessionType          int32    `protobuf:"varint,8,opt,name=SessionType,proto3" json:"SessionType,omitempty"`
	OperationID          string   `protobuf:"bytes,9,opt,name=OperationID,proto3" json:"OperationID,omitempty"`
	MsgID                string   `protobuf:"bytes,10,opt,name=MsgID,proto3" json:"MsgID,omitempty"`
	OfflineInfo          string   `protobuf:"bytes,11,opt,name=OfflineInfo,proto3" json:"OfflineInfo,omitempty"`
	Options              string   `protobuf:"bytes,12,opt,name=Options,proto3" json:"Options,omitempty"`
	PlatformID           int32    `protobuf:"varint,13,opt,name=PlatformID,proto3" json:"PlatformID,omitempty"`
	IsEmphasize          bool     `protobuf:"varint,14,opt,name=IsEmphasize,proto3" json:"IsEmphasize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgSvrToPushSvrChatMsg) Reset()         { *m = MsgSvrToPushSvrChatMsg{} }
func (m *MsgSvrToPushSvrChatMsg) String() string { return proto.CompactTextString(m) }
func (*MsgSvrToPushSvrChatMsg) ProtoMessage()    {}
func (*MsgSvrToPushSvrChatMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a49518b351716, []int{1}
}

func (m *MsgSvrToPushSvrChatMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgSvrToPushSvrChatMsg.Unmarshal(m, b)
}
func (m *MsgSvrToPushSvrChatMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgSvrToPushSvrChatMsg.Marshal(b, m, deterministic)
}
func (m *MsgSvrToPushSvrChatMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSvrToPushSvrChatMsg.Merge(m, src)
}
func (m *MsgSvrToPushSvrChatMsg) XXX_Size() int {
	return xxx_messageInfo_MsgSvrToPushSvrChatMsg.Size(m)
}
func (m *MsgSvrToPushSvrChatMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSvrToPushSvrChatMsg.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSvrToPushSvrChatMsg proto.InternalMessageInfo

func (m *MsgSvrToPushSvrChatMsg) GetSendID() string {
	if m != nil {
		return m.SendID
	}
	return ""
}

func (m *MsgSvrToPushSvrChatMsg) GetRecvID() string {
	if m != nil {
		return m.RecvID
	}
	return ""
}

func (m *MsgSvrToPushSvrChatMsg) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *MsgSvrToPushSvrChatMsg) GetRecvSeq() int64 {
	if m != nil {
		return m.RecvSeq
	}
	return 0
}

func (m *MsgSvrToPushSvrChatMsg) GetSendTime() int64 {
	if m != nil {
		return m.SendTime
	}
	return 0
}

func (m *MsgSvrToPushSvrChatMsg) GetMsgFrom() int32 {
	if m != nil {
		return m.MsgFrom
	}
	return 0
}

func (m *MsgSvrToPushSvrChatMsg) GetContentType() int32 {
	if m != nil {
		return m.ContentType
	}
	return 0
}

func (m *MsgSvrToPushSvrChatMsg) GetSessionType() int32 {
	if m != nil {
		return m.SessionType
	}
	return 0
}

func (m *MsgSvrToPushSvrChatMsg) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

func (m *MsgSvrToPushSvrChatMsg) GetMsgID() string {
	if m != nil {
		return m.MsgID
	}
	return ""
}

func (m *MsgSvrToPushSvrChatMsg) GetOfflineInfo() string {
	if m != nil {
		return m.OfflineInfo
	}
	return ""
}

func (m *MsgSvrToPushSvrChatMsg) GetOptions() string {
	if m != nil {
		return m.Options
	}
	return ""
}

func (m *MsgSvrToPushSvrChatMsg) GetPlatformID() int32 {
	if m != nil {
		return m.PlatformID
	}
	return 0
}

func (m *MsgSvrToPushSvrChatMsg) GetIsEmphasize() bool {
	if m != nil {
		return m.IsEmphasize
	}
	return false
}

type PullMessageReq struct {
	UserID               string   `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	SeqBegin             int64    `protobuf:"varint,2,opt,name=SeqBegin,proto3" json:"SeqBegin,omitempty"`
	SeqEnd               int64    `protobuf:"varint,3,opt,name=SeqEnd,proto3" json:"SeqEnd,omitempty"`
	OperationID          string   `protobuf:"bytes,4,opt,name=OperationID,proto3" json:"OperationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PullMessageReq) Reset()         { *m = PullMessageReq{} }
func (m *PullMessageReq) String() string { return proto.CompactTextString(m) }
func (*PullMessageReq) ProtoMessage()    {}
func (*PullMessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a49518b351716, []int{2}
}

func (m *PullMessageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullMessageReq.Unmarshal(m, b)
}
func (m *PullMessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullMessageReq.Marshal(b, m, deterministic)
}
func (m *PullMessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullMessageReq.Merge(m, src)
}
func (m *PullMessageReq) XXX_Size() int {
	return xxx_messageInfo_PullMessageReq.Size(m)
}
func (m *PullMessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PullMessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_PullMessageReq proto.InternalMessageInfo

func (m *PullMessageReq) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *PullMessageReq) GetSeqBegin() int64 {
	if m != nil {
		return m.SeqBegin
	}
	return 0
}

func (m *PullMessageReq) GetSeqEnd() int64 {
	if m != nil {
		return m.SeqEnd
	}
	return 0
}

func (m *PullMessageReq) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

type PullMessageResp struct {
	ErrCode              int32           `protobuf:"varint,1,opt,name=ErrCode,proto3" json:"ErrCode,omitempty"`
	ErrMsg               string          `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	MaxSeq               int64           `protobuf:"varint,3,opt,name=MaxSeq,proto3" json:"MaxSeq,omitempty"`
	MinSeq               int64           `protobuf:"varint,4,opt,name=MinSeq,proto3" json:"MinSeq,omitempty"`
	SingleUserMsg        []*GatherFormat `protobuf:"bytes,5,rep,name=SingleUserMsg,proto3" json:"SingleUserMsg,omitempty"`
	GroupUserMsg         []*GatherFormat `protobuf:"bytes,6,rep,name=GroupUserMsg,proto3" json:"GroupUserMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PullMessageResp) Reset()         { *m = PullMessageResp{} }
func (m *PullMessageResp) String() string { return proto.CompactTextString(m) }
func (*PullMessageResp) ProtoMessage()    {}
func (*PullMessageResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a49518b351716, []int{3}
}

func (m *PullMessageResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullMessageResp.Unmarshal(m, b)
}
func (m *PullMessageResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullMessageResp.Marshal(b, m, deterministic)
}
func (m *PullMessageResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullMessageResp.Merge(m, src)
}
func (m *PullMessageResp) XXX_Size() int {
	return xxx_messageInfo_PullMessageResp.Size(m)
}
func (m *PullMessageResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PullMessageResp.DiscardUnknown(m)
}

var xxx_messageInfo_PullMessageResp proto.InternalMessageInfo

func (m *PullMessageResp) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *PullMessageResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *PullMessageResp) GetMaxSeq() int64 {
	if m != nil {
		return m.MaxSeq
	}
	return 0
}

func (m *PullMessageResp) GetMinSeq() int64 {
	if m != nil {
		return m.MinSeq
	}
	return 0
}

func (m *PullMessageResp) GetSingleUserMsg() []*GatherFormat {
	if m != nil {
		return m.SingleUserMsg
	}
	return nil
}

func (m *PullMessageResp) GetGroupUserMsg() []*GatherFormat {
	if m != nil {
		return m.GroupUserMsg
	}
	return nil
}

type GetNewSeqReq struct {
	UserID               string   `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	OperationID          string   `protobuf:"bytes,2,opt,name=OperationID,proto3" json:"OperationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNewSeqReq) Reset()         { *m = GetNewSeqReq{} }
func (m *GetNewSeqReq) String() string { return proto.CompactTextString(m) }
func (*GetNewSeqReq) ProtoMessage()    {}
func (*GetNewSeqReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a49518b351716, []int{4}
}

func (m *GetNewSeqReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNewSeqReq.Unmarshal(m, b)
}
func (m *GetNewSeqReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNewSeqReq.Marshal(b, m, deterministic)
}
func (m *GetNewSeqReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNewSeqReq.Merge(m, src)
}
func (m *GetNewSeqReq) XXX_Size() int {
	return xxx_messageInfo_GetNewSeqReq.Size(m)
}
func (m *GetNewSeqReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNewSeqReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetNewSeqReq proto.InternalMessageInfo

func (m *GetNewSeqReq) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *GetNewSeqReq) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

type GetNewSeqResp struct {
	ErrCode              int32    `protobuf:"varint,1,opt,name=ErrCode,proto3" json:"ErrCode,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	Seq                  int64    `protobuf:"varint,3,opt,name=Seq,proto3" json:"Seq,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNewSeqResp) Reset()         { *m = GetNewSeqResp{} }
func (m *GetNewSeqResp) String() string { return proto.CompactTextString(m) }
func (*GetNewSeqResp) ProtoMessage()    {}
func (*GetNewSeqResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a49518b351716, []int{5}
}

func (m *GetNewSeqResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNewSeqResp.Unmarshal(m, b)
}
func (m *GetNewSeqResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNewSeqResp.Marshal(b, m, deterministic)
}
func (m *GetNewSeqResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNewSeqResp.Merge(m, src)
}
func (m *GetNewSeqResp) XXX_Size() int {
	return xxx_messageInfo_GetNewSeqResp.Size(m)
}
func (m *GetNewSeqResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNewSeqResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetNewSeqResp proto.InternalMessageInfo

func (m *GetNewSeqResp) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *GetNewSeqResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *GetNewSeqResp) GetSeq() int64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

type GatherFormat struct {
	// @inject_tag: json:"id"
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"id"`
	// @inject_tag: json:"list"
	List                 []*MsgFormat `protobuf:"bytes,2,rep,name=List,proto3" json:"list"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GatherFormat) Reset()         { *m = GatherFormat{} }
func (m *GatherFormat) String() string { return proto.CompactTextString(m) }
func (*GatherFormat) ProtoMessage()    {}
func (*GatherFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a49518b351716, []int{6}
}

func (m *GatherFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatherFormat.Unmarshal(m, b)
}
func (m *GatherFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatherFormat.Marshal(b, m, deterministic)
}
func (m *GatherFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatherFormat.Merge(m, src)
}
func (m *GatherFormat) XXX_Size() int {
	return xxx_messageInfo_GatherFormat.Size(m)
}
func (m *GatherFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_GatherFormat.DiscardUnknown(m)
}

var xxx_messageInfo_GatherFormat proto.InternalMessageInfo

func (m *GatherFormat) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *GatherFormat) GetList() []*MsgFormat {
	if m != nil {
		return m.List
	}
	return nil
}

type MsgFormat struct {
	// @inject_tag: json:"sendID"
	SendID string `protobuf:"bytes,1,opt,name=SendID,proto3" json:"sendID"`
	// @inject_tag: json:"recvID"
	RecvID string `protobuf:"bytes,2,opt,name=RecvID,proto3" json:"recvID"`
	// @inject_tag: json:"msgFrom"
	MsgFrom int32 `protobuf:"varint,3,opt,name=MsgFrom,proto3" json:"msgFrom"`
	// @inject_tag: json:"contentType"
	ContentType int32 `protobuf:"varint,4,opt,name=ContentType,proto3" json:"contentType"`
	// @inject_tag: json:"serverMsgID"
	ServerMsgID string `protobuf:"bytes,5,opt,name=ServerMsgID,proto3" json:"serverMsgID"`
	// @inject_tag: json:"content"
	Content string `protobuf:"bytes,6,opt,name=Content,proto3" json:"content"`
	// @inject_tag: json:"seq"
	Seq int64 `protobuf:"varint,7,opt,name=Seq,proto3" json:"seq"`
	// @inject_tag: json:"sendTime"
	SendTime int64 `protobuf:"varint,8,opt,name=SendTime,proto3" json:"sendTime"`
	// @inject_tag: json:"senderPlatformID"
	SenderPlatformID     int32    `protobuf:"varint,9,opt,name=SenderPlatformID,proto3" json:"senderPlatformID"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgFormat) Reset()         { *m = MsgFormat{} }
func (m *MsgFormat) String() string { return proto.CompactTextString(m) }
func (*MsgFormat) ProtoMessage()    {}
func (*MsgFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a49518b351716, []int{7}
}

func (m *MsgFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgFormat.Unmarshal(m, b)
}
func (m *MsgFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgFormat.Marshal(b, m, deterministic)
}
func (m *MsgFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFormat.Merge(m, src)
}
func (m *MsgFormat) XXX_Size() int {
	return xxx_messageInfo_MsgFormat.Size(m)
}
func (m *MsgFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFormat.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFormat proto.InternalMessageInfo

func (m *MsgFormat) GetSendID() string {
	if m != nil {
		return m.SendID
	}
	return ""
}

func (m *MsgFormat) GetRecvID() string {
	if m != nil {
		return m.RecvID
	}
	return ""
}

func (m *MsgFormat) GetMsgFrom() int32 {
	if m != nil {
		return m.MsgFrom
	}
	return 0
}

func (m *MsgFormat) GetContentType() int32 {
	if m != nil {
		return m.ContentType
	}
	return 0
}

func (m *MsgFormat) GetServerMsgID() string {
	if m != nil {
		return m.ServerMsgID
	}
	return ""
}

func (m *MsgFormat) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *MsgFormat) GetSeq() int64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *MsgFormat) GetSendTime() int64 {
	if m != nil {
		return m.SendTime
	}
	return 0
}

func (m *MsgFormat) GetSenderPlatformID() int32 {
	if m != nil {
		return m.SenderPlatformID
	}
	return 0
}

type UserSendMsgReq struct {
	ReqIdentifier        int32    `protobuf:"varint,1,opt,name=ReqIdentifier,proto3" json:"ReqIdentifier,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	SendID               string   `protobuf:"bytes,3,opt,name=SendID,proto3" json:"SendID,omitempty"`
	OperationID          string   `protobuf:"bytes,4,opt,name=OperationID,proto3" json:"OperationID,omitempty"`
	MsgIncr              int32    `protobuf:"varint,5,opt,name=MsgIncr,proto3" json:"MsgIncr,omitempty"`
	PlatformID           int32    `protobuf:"varint,6,opt,name=PlatformID,proto3" json:"PlatformID,omitempty"`
	SessionType          int32    `protobuf:"varint,7,opt,name=SessionType,proto3" json:"SessionType,omitempty"`
	MsgFrom              int32    `protobuf:"varint,8,opt,name=MsgFrom,proto3" json:"MsgFrom,omitempty"`
	ContentType          int32    `protobuf:"varint,9,opt,name=ContentType,proto3" json:"ContentType,omitempty"`
	RecvID               string   `protobuf:"bytes,10,opt,name=RecvID,proto3" json:"RecvID,omitempty"`
	ForceList            []string `protobuf:"bytes,11,rep,name=ForceList,proto3" json:"ForceList,omitempty"`
	Content              string   `protobuf:"bytes,12,opt,name=Content,proto3" json:"Content,omitempty"`
	Options              string   `protobuf:"bytes,13,opt,name=Options,proto3" json:"Options,omitempty"`
	ClientMsgID          string   `protobuf:"bytes,14,opt,name=ClientMsgID,proto3" json:"ClientMsgID,omitempty"`
	OffLineInfo          string   `protobuf:"bytes,15,opt,name=OffLineInfo,proto3" json:"OffLineInfo,omitempty"`
	Ex                   string   `protobuf:"bytes,16,opt,name=Ex,proto3" json:"Ex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSendMsgReq) Reset()         { *m = UserSendMsgReq{} }
func (m *UserSendMsgReq) String() string { return proto.CompactTextString(m) }
func (*UserSendMsgReq) ProtoMessage()    {}
func (*UserSendMsgReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a49518b351716, []int{8}
}

func (m *UserSendMsgReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSendMsgReq.Unmarshal(m, b)
}
func (m *UserSendMsgReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSendMsgReq.Marshal(b, m, deterministic)
}
func (m *UserSendMsgReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSendMsgReq.Merge(m, src)
}
func (m *UserSendMsgReq) XXX_Size() int {
	return xxx_messageInfo_UserSendMsgReq.Size(m)
}
func (m *UserSendMsgReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSendMsgReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserSendMsgReq proto.InternalMessageInfo

func (m *UserSendMsgReq) GetReqIdentifier() int32 {
	if m != nil {
		return m.ReqIdentifier
	}
	return 0
}

func (m *UserSendMsgReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UserSendMsgReq) GetSendID() string {
	if m != nil {
		return m.SendID
	}
	return ""
}

func (m *UserSendMsgReq) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

func (m *UserSendMsgReq) GetMsgIncr() int32 {
	if m != nil {
		return m.MsgIncr
	}
	return 0
}

func (m *UserSendMsgReq) GetPlatformID() int32 {
	if m != nil {
		return m.PlatformID
	}
	return 0
}

func (m *UserSendMsgReq) GetSessionType() int32 {
	if m != nil {
		return m.SessionType
	}
	return 0
}

func (m *UserSendMsgReq) GetMsgFrom() int32 {
	if m != nil {
		return m.MsgFrom
	}
	return 0
}

func (m *UserSendMsgReq) GetContentType() int32 {
	if m != nil {
		return m.ContentType
	}
	return 0
}

func (m *UserSendMsgReq) GetRecvID() string {
	if m != nil {
		return m.RecvID
	}
	return ""
}

func (m *UserSendMsgReq) GetForceList() []string {
	if m != nil {
		return m.ForceList
	}
	return nil
}

func (m *UserSendMsgReq) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *UserSendMsgReq) GetOptions() string {
	if m != nil {
		return m.Options
	}
	return ""
}

func (m *UserSendMsgReq) GetClientMsgID() string {
	if m != nil {
		return m.ClientMsgID
	}
	return ""
}

func (m *UserSendMsgReq) GetOffLineInfo() string {
	if m != nil {
		return m.OffLineInfo
	}
	return ""
}

func (m *UserSendMsgReq) GetEx() string {
	if m != nil {
		return m.Ex
	}
	return ""
}

type UserSendMsgResp struct {
	ErrCode              int32    `protobuf:"varint,1,opt,name=ErrCode,proto3" json:"ErrCode,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	ReqIdentifier        int32    `protobuf:"varint,3,opt,name=ReqIdentifier,proto3" json:"ReqIdentifier,omitempty"`
	MsgIncr              int32    `protobuf:"varint,4,opt,name=MsgIncr,proto3" json:"MsgIncr,omitempty"`
	SendTime             int64    `protobuf:"varint,5,opt,name=SendTime,proto3" json:"SendTime,omitempty"`
	ServerMsgID          string   `protobuf:"bytes,6,opt,name=ServerMsgID,proto3" json:"ServerMsgID,omitempty"`
	ClientMsgID          string   `protobuf:"bytes,7,opt,name=ClientMsgID,proto3" json:"ClientMsgID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSendMsgResp) Reset()         { *m = UserSendMsgResp{} }
func (m *UserSendMsgResp) String() string { return proto.CompactTextString(m) }
func (*UserSendMsgResp) ProtoMessage()    {}
func (*UserSendMsgResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a49518b351716, []int{9}
}

func (m *UserSendMsgResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSendMsgResp.Unmarshal(m, b)
}
func (m *UserSendMsgResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSendMsgResp.Marshal(b, m, deterministic)
}
func (m *UserSendMsgResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSendMsgResp.Merge(m, src)
}
func (m *UserSendMsgResp) XXX_Size() int {
	return xxx_messageInfo_UserSendMsgResp.Size(m)
}
func (m *UserSendMsgResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSendMsgResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserSendMsgResp proto.InternalMessageInfo

func (m *UserSendMsgResp) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *UserSendMsgResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *UserSendMsgResp) GetReqIdentifier() int32 {
	if m != nil {
		return m.ReqIdentifier
	}
	return 0
}

func (m *UserSendMsgResp) GetMsgIncr() int32 {
	if m != nil {
		return m.MsgIncr
	}
	return 0
}

func (m *UserSendMsgResp) GetSendTime() int64 {
	if m != nil {
		return m.SendTime
	}
	return 0
}

func (m *UserSendMsgResp) GetServerMsgID() string {
	if m != nil {
		return m.ServerMsgID
	}
	return ""
}

func (m *UserSendMsgResp) GetClientMsgID() string {
	if m != nil {
		return m.ClientMsgID
	}
	return ""
}

func init() {
	proto.RegisterType((*WSToMsgSvrChatMsg)(nil), "pbChat.WSToMsgSvrChatMsg")
	proto.RegisterType((*MsgSvrToPushSvrChatMsg)(nil), "pbChat.MsgSvrToPushSvrChatMsg")
	proto.RegisterType((*PullMessageReq)(nil), "pbChat.PullMessageReq")
	proto.RegisterType((*PullMessageResp)(nil), "pbChat.PullMessageResp")
	proto.RegisterType((*GetNewSeqReq)(nil), "pbChat.GetNewSeqReq")
	proto.RegisterType((*GetNewSeqResp)(nil), "pbChat.GetNewSeqResp")
	proto.RegisterType((*GatherFormat)(nil), "pbChat.GatherFormat")
	proto.RegisterType((*MsgFormat)(nil), "pbChat.MsgFormat")
	proto.RegisterType((*UserSendMsgReq)(nil), "pbChat.UserSendMsgReq")
	proto.RegisterType((*UserSendMsgResp)(nil), "pbChat.UserSendMsgResp")
}

func init() {
	proto.RegisterFile("chat/chat.proto", fileDescriptor_767a49518b351716)
}

var fileDescriptor_767a49518b351716 = []byte{
	// 874 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xcd, 0x6e, 0xeb, 0x44,
	0x18, 0x95, 0xe3, 0xc4, 0x89, 0xbf, 0xfc, 0xf5, 0x8e, 0x2e, 0x65, 0x54, 0x21, 0x64, 0x59, 0x20,
	0x45, 0x2c, 0x8a, 0x74, 0xd9, 0x5c, 0x81, 0xc4, 0xe2, 0x36, 0x69, 0xb1, 0x74, 0xcd, 0xad, 0xec,
	0x20, 0x24, 0x76, 0x6e, 0x32, 0x71, 0x2c, 0x12, 0xdb, 0x99, 0x71, 0x4b, 0x61, 0xc1, 0x1b, 0xf0,
	0x0c, 0xbc, 0x00, 0x0f, 0xc1, 0x83, 0xb0, 0x63, 0xc5, 0x5b, 0xa0, 0x99, 0xb1, 0xe3, 0xb1, 0x9d,
	0x26, 0x52, 0xa5, 0x6e, 0xa2, 0x9c, 0x33, 0x3f, 0x99, 0x39, 0xe7, 0xfb, 0xce, 0x04, 0xc6, 0x8b,
	0x75, 0x90, 0x7d, 0xc9, 0x3f, 0x2e, 0x53, 0x9a, 0x64, 0x09, 0x32, 0xd2, 0xbb, 0xab, 0x75, 0x90,
	0xd9, 0x7f, 0xea, 0xf0, 0xea, 0x47, 0x7f, 0x9e, 0xb8, 0x2c, 0xf4, 0x1f, 0x28, 0xa7, 0x5c, 0x16,
	0xa2, 0x73, 0x30, 0x7c, 0x12, 0x2f, 0x9d, 0x29, 0xd6, 0x2c, 0x6d, 0x62, 0x7a, 0x39, 0xe2, 0xbc,
	0x47, 0x16, 0x0f, 0xce, 0x14, 0xb7, 0x24, 0x2f, 0x11, 0xc2, 0xd0, 0xbd, 0x4a, 0xe2, 0x8c, 0xc4,
	0x19, 0xd6, 0xc5, 0x40, 0x01, 0xd1, 0x05, 0xf4, 0xf8, 0xda, 0x79, 0xb4, 0x25, 0xb8, 0x6d, 0x69,
	0x13, 0xdd, 0xdb, 0x63, 0xbe, 0xca, 0x65, 0xe1, 0x35, 0x4d, 0xb6, 0xb8, 0x63, 0x69, 0x93, 0x8e,
	0x57, 0x40, 0x64, 0x41, 0x3f, 0xdf, 0x60, 0xfe, 0x6b, 0x4a, 0xb0, 0x21, 0x46, 0x55, 0x8a, 0xcf,
	0xf0, 0x09, 0x63, 0x51, 0x12, 0x8b, 0x19, 0x5d, 0x39, 0x43, 0xa1, 0xf8, 0x8c, 0x0f, 0x29, 0xa1,
	0x41, 0x16, 0x25, 0xb1, 0x33, 0xc5, 0x3d, 0x71, 0x2e, 0x95, 0x42, 0xaf, 0xa1, 0xe3, 0xb2, 0xd0,
	0x99, 0x62, 0x53, 0x8c, 0x49, 0xc0, 0xd9, 0x79, 0xf2, 0x33, 0x89, 0x31, 0x48, 0x56, 0x00, 0xb1,
	0xdb, 0x6a, 0xb5, 0x89, 0x62, 0xe2, 0xc4, 0xab, 0x04, 0xf7, 0xf3, 0xdd, 0x4a, 0x8a, 0xdf, 0xe6,
	0x43, 0xca, 0x77, 0x66, 0x78, 0x20, 0x35, 0xc8, 0x21, 0xfa, 0x14, 0xe0, 0x76, 0x13, 0x64, 0xab,
	0x84, 0x6e, 0x9d, 0x29, 0x1e, 0x8a, 0xa3, 0x2a, 0x0c, 0xfa, 0x04, 0xcc, 0xeb, 0x84, 0x2e, 0xc8,
	0xfb, 0x88, 0x65, 0x78, 0x64, 0xe9, 0x13, 0xd3, 0x2b, 0x09, 0xfb, 0x2f, 0x1d, 0xce, 0xa5, 0x3b,
	0xf3, 0xe4, 0xf6, 0x9e, 0xad, 0x5f, 0xc4, 0x26, 0x0c, 0x5d, 0x3e, 0xc7, 0x27, 0xbb, 0xdc, 0xa5,
	0x02, 0x56, 0x0c, 0xec, 0x3c, 0x6d, 0xa0, 0x71, 0xd4, 0xc0, 0xee, 0x49, 0x03, 0x7b, 0x27, 0x0d,
	0x34, 0x8f, 0x18, 0x08, 0xaa, 0x81, 0x2f, 0x69, 0x95, 0x05, 0x7d, 0x87, 0xcd, 0xb6, 0xe9, 0x3a,
	0x60, 0xd1, 0x6f, 0x04, 0x8f, 0x2c, 0x6d, 0xd2, 0xf3, 0x54, 0xca, 0xfe, 0x1d, 0x46, 0xb7, 0xf7,
	0x9b, 0x8d, 0x4b, 0x18, 0x0b, 0x42, 0xe2, 0x91, 0x1d, 0x77, 0xe3, 0x07, 0x46, 0x68, 0xe9, 0x92,
	0x44, 0x52, 0xd9, 0xdd, 0x3b, 0x12, 0x46, 0xb1, 0xf0, 0x49, 0x28, 0x2b, 0xb1, 0x74, 0x76, 0x37,
	0x8b, 0x97, 0xc2, 0x28, 0xdd, 0xcb, 0x51, 0x5d, 0x93, 0x76, 0x43, 0x13, 0xfb, 0x3f, 0x0d, 0xc6,
	0x95, 0x03, 0xb0, 0x94, 0xdf, 0x77, 0x46, 0xe9, 0x55, 0xb2, 0x24, 0xe2, 0x08, 0x1d, 0xaf, 0x80,
	0xfc, 0x77, 0x66, 0x94, 0xba, 0x2c, 0x2c, 0x2a, 0x45, 0x22, 0xce, 0xbb, 0xc1, 0x23, 0x2f, 0x87,
	0xfc, 0xf7, 0x25, 0x12, 0x7c, 0x14, 0x97, 0x65, 0x92, 0x23, 0xf4, 0x35, 0x0c, 0xfd, 0x28, 0x0e,
	0x37, 0x84, 0xdf, 0x8d, 0x6f, 0xd7, 0xb1, 0xf4, 0x49, 0xff, 0xcd, 0xeb, 0x4b, 0x19, 0x33, 0x97,
	0x37, 0x41, 0xb6, 0x26, 0xf4, 0x3a, 0xa1, 0xdb, 0x20, 0xf3, 0xaa, 0x53, 0xd1, 0x5b, 0x18, 0xdc,
	0xd0, 0xe4, 0x3e, 0x2d, 0x96, 0x1a, 0x47, 0x96, 0x56, 0x66, 0xda, 0xdf, 0xc1, 0xe0, 0x86, 0x64,
	0xdf, 0x93, 0x5f, 0x7c, 0xb2, 0x3b, 0xa6, 0x74, 0x4d, 0xb5, 0x56, 0x53, 0x35, 0x1f, 0x86, 0xca,
	0x4e, 0xcf, 0x92, 0xec, 0x0c, 0xf4, 0x52, 0x2f, 0xfe, 0xd5, 0x9e, 0xc1, 0x40, 0x3d, 0x3c, 0x1a,
	0x41, 0x6b, 0x7f, 0xb4, 0x96, 0x33, 0x45, 0x9f, 0x43, 0x5b, 0xb4, 0x7c, 0x4b, 0x5c, 0xf8, 0x55,
	0x71, 0x61, 0xde, 0x43, 0xf2, 0xb6, 0x62, 0xd8, 0xfe, 0xa3, 0x05, 0xe6, 0x9e, 0x7b, 0x4e, 0xcf,
	0x17, 0x3d, 0xaa, 0x1f, 0xed, 0xd1, 0xf6, 0x13, 0x3d, 0x4a, 0x1f, 0x84, 0xd8, 0xce, 0x54, 0xb4,
	0xbf, 0xe9, 0xa9, 0x94, 0x9a, 0x28, 0x46, 0x35, 0x51, 0x72, 0x39, 0xba, 0x7b, 0x39, 0x2a, 0x49,
	0xd2, 0xab, 0x25, 0xc9, 0x17, 0x70, 0xc6, 0xbf, 0x13, 0xaa, 0x74, 0x9f, 0x29, 0x0e, 0xd4, 0xe0,
	0xed, 0x7f, 0x74, 0x18, 0x71, 0x63, 0xf9, 0x80, 0xcb, 0x42, 0x6e, 0xfc, 0x67, 0x30, 0xf4, 0xc8,
	0xce, 0x59, 0x92, 0x38, 0x8b, 0x56, 0x11, 0xa1, 0xb9, 0x67, 0x55, 0xb2, 0x4c, 0xf6, 0x96, 0x9a,
	0xec, 0xa5, 0xa0, 0x7a, 0x45, 0xd0, 0x93, 0xad, 0x96, 0x4b, 0xeb, 0xc4, 0x0b, 0xaa, 0xbc, 0x5f,
	0x1c, 0xd6, 0x62, 0xc4, 0x38, 0x14, 0x23, 0x27, 0x5e, 0x2f, 0xc5, 0xb6, 0xde, 0x51, 0xdb, 0xcc,
	0xa6, 0x6d, 0x65, 0x29, 0x40, 0xa5, 0x14, 0x2a, 0xef, 0x4c, 0xbf, 0xf6, 0xce, 0xa8, 0x56, 0x0e,
	0x1a, 0x8f, 0x43, 0x11, 0x97, 0xc3, 0x6a, 0x5c, 0xf2, 0xb3, 0x6c, 0x22, 0x12, 0x67, 0xb2, 0x40,
	0x46, 0x52, 0x23, 0x85, 0xca, 0xc3, 0xf8, 0x7d, 0x11, 0xc6, 0xe3, 0x7d, 0x18, 0x17, 0x14, 0xef,
	0x8a, 0xd9, 0x23, 0x3e, 0x93, 0x5d, 0x31, 0x7b, 0xb4, 0xff, 0xd5, 0x60, 0x5c, 0xb1, 0xf7, 0x59,
	0xdd, 0xd8, 0xa8, 0x08, 0xfd, 0x50, 0x45, 0x28, 0x0e, 0xb6, 0xab, 0x0e, 0x1e, 0x7b, 0xf6, 0x6a,
	0x6d, 0x61, 0x34, 0xdb, 0xa2, 0xa6, 0x4b, 0xb7, 0xa1, 0xcb, 0x9b, 0xbf, 0x35, 0x68, 0xf3, 0x6e,
	0x47, 0x6f, 0xc1, 0xdc, 0x27, 0x0f, 0x2a, 0x43, 0x4f, 0x89, 0xb5, 0x8b, 0x8f, 0x0e, 0xb0, 0x2c,
	0x45, 0xdf, 0x42, 0x5f, 0x09, 0x7a, 0x74, 0x5e, 0xcc, 0xaa, 0x3e, 0x3f, 0x17, 0x1f, 0x1f, 0xe4,
	0xe5, 0x7a, 0x45, 0xe7, 0x72, 0x7d, 0xb5, 0xb7, 0xca, 0xf5, 0x35, 0x53, 0xde, 0x0d, 0x7f, 0xea,
	0xf3, 0x3f, 0x94, 0xdf, 0xc8, 0xe1, 0x3b, 0x43, 0xfc, 0xb1, 0xfc, 0xea, 0xff, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x1f, 0x48, 0xc5, 0x48, 0x6b, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatClient interface {
	GetNewSeq(ctx context.Context, in *GetNewSeqReq, opts ...grpc.CallOption) (*GetNewSeqResp, error)
	PullMessage(ctx context.Context, in *PullMessageReq, opts ...grpc.CallOption) (*PullMessageResp, error)
	UserSendMsg(ctx context.Context, in *UserSendMsgReq, opts ...grpc.CallOption) (*UserSendMsgResp, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) GetNewSeq(ctx context.Context, in *GetNewSeqReq, opts ...grpc.CallOption) (*GetNewSeqResp, error) {
	out := new(GetNewSeqResp)
	err := c.cc.Invoke(ctx, "/pbChat.Chat/GetNewSeq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) PullMessage(ctx context.Context, in *PullMessageReq, opts ...grpc.CallOption) (*PullMessageResp, error) {
	out := new(PullMessageResp)
	err := c.cc.Invoke(ctx, "/pbChat.Chat/PullMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) UserSendMsg(ctx context.Context, in *UserSendMsgReq, opts ...grpc.CallOption) (*UserSendMsgResp, error) {
	out := new(UserSendMsgResp)
	err := c.cc.Invoke(ctx, "/pbChat.Chat/UserSendMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
type ChatServer interface {
	GetNewSeq(context.Context, *GetNewSeqReq) (*GetNewSeqResp, error)
	PullMessage(context.Context, *PullMessageReq) (*PullMessageResp, error)
	UserSendMsg(context.Context, *UserSendMsgReq) (*UserSendMsgResp, error)
}

// UnimplementedChatServer can be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (*UnimplementedChatServer) GetNewSeq(ctx context.Context, req *GetNewSeqReq) (*GetNewSeqResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewSeq not implemented")
}
func (*UnimplementedChatServer) PullMessage(ctx context.Context, req *PullMessageReq) (*PullMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullMessage not implemented")
}
func (*UnimplementedChatServer) UserSendMsg(ctx context.Context, req *UserSendMsgReq) (*UserSendMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSendMsg not implemented")
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_GetNewSeq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewSeqReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetNewSeq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbChat.Chat/GetNewSeq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetNewSeq(ctx, req.(*GetNewSeqReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_PullMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).PullMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbChat.Chat/PullMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).PullMessage(ctx, req.(*PullMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_UserSendMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSendMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).UserSendMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbChat.Chat/UserSendMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).UserSendMsg(ctx, req.(*UserSendMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbChat.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNewSeq",
			Handler:    _Chat_GetNewSeq_Handler,
		},
		{
			MethodName: "PullMessage",
			Handler:    _Chat_PullMessage_Handler,
		},
		{
			MethodName: "UserSendMsg",
			Handler:    _Chat_UserSendMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat/chat.proto",
}
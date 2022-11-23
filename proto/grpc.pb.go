// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.6
// source: proto/grpc.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Acknowledgement_OUTCOME int32

const (
	Acknowledgement_SUCCESS Acknowledgement_OUTCOME = 0
	Acknowledgement_FAIL    Acknowledgement_OUTCOME = 1
)

// Enum value maps for Acknowledgement_OUTCOME.
var (
	Acknowledgement_OUTCOME_name = map[int32]string{
		0: "SUCCESS",
		1: "FAIL",
	}
	Acknowledgement_OUTCOME_value = map[string]int32{
		"SUCCESS": 0,
		"FAIL":    1,
	}
)

func (x Acknowledgement_OUTCOME) Enum() *Acknowledgement_OUTCOME {
	p := new(Acknowledgement_OUTCOME)
	*p = x
	return p
}

func (x Acknowledgement_OUTCOME) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Acknowledgement_OUTCOME) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_grpc_proto_enumTypes[0].Descriptor()
}

func (Acknowledgement_OUTCOME) Type() protoreflect.EnumType {
	return &file_proto_grpc_proto_enumTypes[0]
}

func (x Acknowledgement_OUTCOME) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Acknowledgement_OUTCOME.Descriptor instead.
func (Acknowledgement_OUTCOME) EnumDescriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{0, 0}
}

type Ack_RESULT int32

const (
	Ack_SUCCESS   Ack_RESULT = 0
	Ack_FAIL      Ack_RESULT = 1
	Ack_EXCEPTION Ack_RESULT = 2
)

// Enum value maps for Ack_RESULT.
var (
	Ack_RESULT_name = map[int32]string{
		0: "SUCCESS",
		1: "FAIL",
		2: "EXCEPTION",
	}
	Ack_RESULT_value = map[string]int32{
		"SUCCESS":   0,
		"FAIL":      1,
		"EXCEPTION": 2,
	}
)

func (x Ack_RESULT) Enum() *Ack_RESULT {
	p := new(Ack_RESULT)
	*p = x
	return p
}

func (x Ack_RESULT) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ack_RESULT) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_grpc_proto_enumTypes[1].Descriptor()
}

func (Ack_RESULT) Type() protoreflect.EnumType {
	return &file_proto_grpc_proto_enumTypes[1]
}

func (x Ack_RESULT) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ack_RESULT.Descriptor instead.
func (Ack_RESULT) EnumDescriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{3, 0}
}

type Outcome_STATE int32

const (
	Outcome_ONGOING    Outcome_STATE = 0
	Outcome_FINISHED   Outcome_STATE = 1
	Outcome_NOTSTARTED Outcome_STATE = 2
)

// Enum value maps for Outcome_STATE.
var (
	Outcome_STATE_name = map[int32]string{
		0: "ONGOING",
		1: "FINISHED",
		2: "NOTSTARTED",
	}
	Outcome_STATE_value = map[string]int32{
		"ONGOING":    0,
		"FINISHED":   1,
		"NOTSTARTED": 2,
	}
)

func (x Outcome_STATE) Enum() *Outcome_STATE {
	p := new(Outcome_STATE)
	*p = x
	return p
}

func (x Outcome_STATE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Outcome_STATE) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_grpc_proto_enumTypes[2].Descriptor()
}

func (Outcome_STATE) Type() protoreflect.EnumType {
	return &file_proto_grpc_proto_enumTypes[2]
}

func (x Outcome_STATE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Outcome_STATE.Descriptor instead.
func (Outcome_STATE) EnumDescriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{4, 0}
}

type Acknowledgement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Outcome Acknowledgement_OUTCOME `protobuf:"varint,1,opt,name=outcome,proto3,enum=proto.Acknowledgement_OUTCOME" json:"outcome,omitempty"`
}

func (x *Acknowledgement) Reset() {
	*x = Acknowledgement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Acknowledgement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Acknowledgement) ProtoMessage() {}

func (x *Acknowledgement) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Acknowledgement.ProtoReflect.Descriptor instead.
func (*Acknowledgement) Descriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *Acknowledgement) GetOutcome() Acknowledgement_OUTCOME {
	if x != nil {
		return x.Outcome
	}
	return Acknowledgement_SUCCESS
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{1}
}

type Amount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LamportTime int32 `protobuf:"varint,1,opt,name=lamportTime,proto3" json:"lamportTime,omitempty"`
	ClientId    int32 `protobuf:"varint,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	BidAmount   int32 `protobuf:"varint,3,opt,name=bidAmount,proto3" json:"bidAmount,omitempty"`
}

func (x *Amount) Reset() {
	*x = Amount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Amount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Amount) ProtoMessage() {}

func (x *Amount) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Amount.ProtoReflect.Descriptor instead.
func (*Amount) Descriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *Amount) GetLamportTime() int32 {
	if x != nil {
		return x.LamportTime
	}
	return 0
}

func (x *Amount) GetClientId() int32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *Amount) GetBidAmount() int32 {
	if x != nil {
		return x.BidAmount
	}
	return 0
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LamportTime int32      `protobuf:"varint,1,opt,name=lamportTime,proto3" json:"lamportTime,omitempty"`
	Result      Ack_RESULT `protobuf:"varint,2,opt,name=result,proto3,enum=proto.Ack_RESULT" json:"result,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *Ack) GetLamportTime() int32 {
	if x != nil {
		return x.LamportTime
	}
	return 0
}

func (x *Ack) GetResult() Ack_RESULT {
	if x != nil {
		return x.Result
	}
	return Ack_SUCCESS
}

type Outcome struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State      Outcome_STATE `protobuf:"varint,1,opt,name=state,proto3,enum=proto.Outcome_STATE" json:"state,omitempty"`
	WinnerId   int32         `protobuf:"varint,2,opt,name=winnerId,proto3" json:"winnerId,omitempty"`
	WinningBid int32         `protobuf:"varint,3,opt,name=winningBid,proto3" json:"winningBid,omitempty"`
}

func (x *Outcome) Reset() {
	*x = Outcome{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outcome) ProtoMessage() {}

func (x *Outcome) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outcome.ProtoReflect.Descriptor instead.
func (*Outcome) Descriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *Outcome) GetState() Outcome_STATE {
	if x != nil {
		return x.State
	}
	return Outcome_ONGOING
}

func (x *Outcome) GetWinnerId() int32 {
	if x != nil {
		return x.WinnerId
	}
	return 0
}

func (x *Outcome) GetWinningBid() int32 {
	if x != nil {
		return x.WinningBid
	}
	return 0
}

var File_proto_grpc_proto protoreflect.FileDescriptor

var file_proto_grpc_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x0f, 0x41, 0x63, 0x6b,
	0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x07,
	0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4f, 0x55, 0x54, 0x43, 0x4f, 0x4d, 0x45, 0x52, 0x07, 0x6f,
	0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x22, 0x20, 0x0a, 0x07, 0x4f, 0x55, 0x54, 0x43, 0x4f, 0x4d,
	0x45, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x64, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6c,
	0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x69, 0x64,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x69,
	0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x82, 0x01, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12,
	0x20, 0x0a, 0x0b, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x29, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x6b, 0x2e, 0x52, 0x45,
	0x53, 0x55, 0x4c, 0x54, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2e, 0x0a, 0x06,
	0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x12, 0x0d, 0x0a,
	0x09, 0x45, 0x58, 0x43, 0x45, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x22, 0xa5, 0x01, 0x0a,
	0x07, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x2e, 0x53, 0x54, 0x41, 0x54, 0x45, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x42, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x77, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x42, 0x69, 0x64,
	0x22, 0x32, 0x0a, 0x05, 0x53, 0x54, 0x41, 0x54, 0x45, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x4e, 0x47,
	0x4f, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x4f, 0x54, 0x53, 0x54, 0x41, 0x52, 0x54,
	0x45, 0x44, 0x10, 0x02, 0x32, 0x53, 0x0a, 0x07, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x20, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63,
	0x6b, 0x12, 0x26, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x32, 0x9f, 0x01, 0x0a, 0x0b, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x09, 0x42, 0x69, 0x64,
	0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63,
	0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a,
	0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63,
	0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a,
	0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x6b, 0x6e,
	0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x27, 0x5a, 0x25, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x69, 0x63, 0x6b, 0x72, 0x6f,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2f, 0x44, 0x49, 0x53, 0x59, 0x53, 0x2d, 0x35, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_grpc_proto_rawDescOnce sync.Once
	file_proto_grpc_proto_rawDescData = file_proto_grpc_proto_rawDesc
)

func file_proto_grpc_proto_rawDescGZIP() []byte {
	file_proto_grpc_proto_rawDescOnce.Do(func() {
		file_proto_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_grpc_proto_rawDescData)
	})
	return file_proto_grpc_proto_rawDescData
}

var file_proto_grpc_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_proto_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_grpc_proto_goTypes = []interface{}{
	(Acknowledgement_OUTCOME)(0), // 0: proto.Acknowledgement.OUTCOME
	(Ack_RESULT)(0),              // 1: proto.Ack.RESULT
	(Outcome_STATE)(0),           // 2: proto.Outcome.STATE
	(*Acknowledgement)(nil),      // 3: proto.Acknowledgement
	(*Empty)(nil),                // 4: proto.Empty
	(*Amount)(nil),               // 5: proto.Amount
	(*Ack)(nil),                  // 6: proto.Ack
	(*Outcome)(nil),              // 7: proto.Outcome
}
var file_proto_grpc_proto_depIdxs = []int32{
	0, // 0: proto.Acknowledgement.outcome:type_name -> proto.Acknowledgement.OUTCOME
	1, // 1: proto.Ack.result:type_name -> proto.Ack.RESULT
	2, // 2: proto.Outcome.state:type_name -> proto.Outcome.STATE
	5, // 3: proto.Auction.Bid:input_type -> proto.Amount
	4, // 4: proto.Auction.Result:input_type -> proto.Empty
	5, // 5: proto.Replication.BidBackup:input_type -> proto.Amount
	4, // 6: proto.Replication.Result:input_type -> proto.Empty
	4, // 7: proto.Replication.Ping:input_type -> proto.Empty
	6, // 8: proto.Auction.Bid:output_type -> proto.Ack
	7, // 9: proto.Auction.Result:output_type -> proto.Outcome
	3, // 10: proto.Replication.BidBackup:output_type -> proto.Acknowledgement
	3, // 11: proto.Replication.Result:output_type -> proto.Acknowledgement
	3, // 12: proto.Replication.Ping:output_type -> proto.Acknowledgement
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_grpc_proto_init() }
func file_proto_grpc_proto_init() {
	if File_proto_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Acknowledgement); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Amount); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Outcome); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_grpc_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_grpc_proto_goTypes,
		DependencyIndexes: file_proto_grpc_proto_depIdxs,
		EnumInfos:         file_proto_grpc_proto_enumTypes,
		MessageInfos:      file_proto_grpc_proto_msgTypes,
	}.Build()
	File_proto_grpc_proto = out.File
	file_proto_grpc_proto_rawDesc = nil
	file_proto_grpc_proto_goTypes = nil
	file_proto_grpc_proto_depIdxs = nil
}

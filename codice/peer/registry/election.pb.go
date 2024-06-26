// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: registry/election.proto

package registry

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

type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr           string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	RecoveryString string `protobuf:"bytes,2,opt,name=recoveryString,proto3" json:"recoveryString,omitempty"`
	RecoveryId     int32  `protobuf:"varint,3,opt,name=recoveryId,proto3" json:"recoveryId,omitempty"`
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_election_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_registry_election_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRequest.ProtoReflect.Descriptor instead.
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return file_registry_election_proto_rawDescGZIP(), []int{0}
}

func (x *JoinRequest) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *JoinRequest) GetRecoveryString() string {
	if x != nil {
		return x.RecoveryString
	}
	return ""
}

func (x *JoinRequest) GetRecoveryId() int32 {
	if x != nil {
		return x.RecoveryId
	}
	return 0
}

type JoinReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PeerList []*PeerInfo `protobuf:"bytes,2,rep,name=peer_list,json=peerList,proto3" json:"peer_list,omitempty"`
}

func (x *JoinReply) Reset() {
	*x = JoinReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_election_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinReply) ProtoMessage() {}

func (x *JoinReply) ProtoReflect() protoreflect.Message {
	mi := &file_registry_election_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinReply.ProtoReflect.Descriptor instead.
func (*JoinReply) Descriptor() ([]byte, []int) {
	return file_registry_election_proto_rawDescGZIP(), []int{1}
}

func (x *JoinReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *JoinReply) GetPeerList() []*PeerInfo {
	if x != nil {
		return x.PeerList
	}
	return nil
}

type UpdateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateString string    `protobuf:"bytes,1,opt,name=updateString,proto3" json:"updateString,omitempty"`
	PeerList     *PeerInfo `protobuf:"bytes,2,opt,name=peer_list,json=peerList,proto3" json:"peer_list,omitempty"`
}

func (x *UpdateMessage) Reset() {
	*x = UpdateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_election_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMessage) ProtoMessage() {}

func (x *UpdateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_registry_election_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMessage.ProtoReflect.Descriptor instead.
func (*UpdateMessage) Descriptor() ([]byte, []int) {
	return file_registry_election_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateMessage) GetUpdateString() string {
	if x != nil {
		return x.UpdateString
	}
	return ""
}

func (x *UpdateMessage) GetPeerList() *PeerInfo {
	if x != nil {
		return x.PeerList
	}
	return nil
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateResponse string `protobuf:"bytes,1,opt,name=updateResponse,proto3" json:"updateResponse,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_election_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_registry_election_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_registry_election_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateResponse) GetUpdateResponse() string {
	if x != nil {
		return x.UpdateResponse
	}
	return ""
}

type PeerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Addr   string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Leader bool   `protobuf:"varint,3,opt,name=leader,proto3" json:"leader,omitempty"`
}

func (x *PeerInfo) Reset() {
	*x = PeerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_election_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerInfo) ProtoMessage() {}

func (x *PeerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_registry_election_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerInfo.ProtoReflect.Descriptor instead.
func (*PeerInfo) Descriptor() ([]byte, []int) {
	return file_registry_election_proto_rawDescGZIP(), []int{4}
}

func (x *PeerInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PeerInfo) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *PeerInfo) GetLeader() bool {
	if x != nil {
		return x.Leader
	}
	return false
}

// The request message containing string TIME
type HeartBeatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HeartBeatMessage) Reset() {
	*x = HeartBeatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_election_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartBeatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBeatMessage) ProtoMessage() {}

func (x *HeartBeatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_registry_election_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBeatMessage.ProtoReflect.Descriptor instead.
func (*HeartBeatMessage) Descriptor() ([]byte, []int) {
	return file_registry_election_proto_rawDescGZIP(), []int{5}
}

func (x *HeartBeatMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Nil struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Nil) Reset() {
	*x = Nil{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_election_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nil) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nil) ProtoMessage() {}

func (x *Nil) ProtoReflect() protoreflect.Message {
	mi := &file_registry_election_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nil.ProtoReflect.Descriptor instead.
func (*Nil) Descriptor() ([]byte, []int) {
	return file_registry_election_proto_rawDescGZIP(), []int{6}
}

type IdLeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdLeader int32 `protobuf:"varint,1,opt,name=idLeader,proto3" json:"idLeader,omitempty"`
}

func (x *IdLeader) Reset() {
	*x = IdLeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_election_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdLeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdLeader) ProtoMessage() {}

func (x *IdLeader) ProtoReflect() protoreflect.Message {
	mi := &file_registry_election_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdLeader.ProtoReflect.Descriptor instead.
func (*IdLeader) Descriptor() ([]byte, []int) {
	return file_registry_election_proto_rawDescGZIP(), []int{7}
}

func (x *IdLeader) GetIdLeader() int32 {
	if x != nil {
		return x.IdLeader
	}
	return 0
}

type ElectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Election   string `protobuf:"bytes,1,opt,name=election,proto3" json:"election,omitempty"`
	ElectionId int32  `protobuf:"varint,2,opt,name=electionId,proto3" json:"electionId,omitempty"`
	TokenId    int32  `protobuf:"varint,3,opt,name=TokenId,proto3" json:"TokenId,omitempty"`
}

func (x *ElectionRequest) Reset() {
	*x = ElectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_election_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElectionRequest) ProtoMessage() {}

func (x *ElectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_registry_election_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElectionRequest.ProtoReflect.Descriptor instead.
func (*ElectionRequest) Descriptor() ([]byte, []int) {
	return file_registry_election_proto_rawDescGZIP(), []int{8}
}

func (x *ElectionRequest) GetElection() string {
	if x != nil {
		return x.Election
	}
	return ""
}

func (x *ElectionRequest) GetElectionId() int32 {
	if x != nil {
		return x.ElectionId
	}
	return 0
}

func (x *ElectionRequest) GetTokenId() int32 {
	if x != nil {
		return x.TokenId
	}
	return 0
}

type ElectionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ElectionReply string `protobuf:"bytes,1,opt,name=electionReply,proto3" json:"electionReply,omitempty"`
	Id            int32  `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *ElectionReply) Reset() {
	*x = ElectionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_election_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElectionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElectionReply) ProtoMessage() {}

func (x *ElectionReply) ProtoReflect() protoreflect.Message {
	mi := &file_registry_election_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElectionReply.ProtoReflect.Descriptor instead.
func (*ElectionReply) Descriptor() ([]byte, []int) {
	return file_registry_election_proto_rawDescGZIP(), []int{9}
}

func (x *ElectionReply) GetElectionReply() string {
	if x != nil {
		return x.ElectionReply
	}
	return ""
}

func (x *ElectionReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_registry_election_proto protoreflect.FileDescriptor

var file_registry_election_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x22, 0x69, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1e,
	0x0a, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x49, 0x64, 0x22, 0x4c,
	0x0a, 0x09, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x09, 0x70,
	0x65, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x08, 0x70, 0x65, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x64, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x2f, 0x0a, 0x09, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x65, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x38, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x46, 0x0a, 0x08,
	0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x22, 0x2c, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x05, 0x0a, 0x03, 0x6e, 0x69, 0x6c, 0x22, 0x26, 0x0a, 0x08, 0x49, 0x64, 0x4c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x4c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x64, 0x4c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x22, 0x67, 0x0a, 0x0f, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x0d, 0x45, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49,
	0x64, 0x32, 0x47, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x3b, 0x0a,
	0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x15, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x4a,
	0x6f, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x32, 0x4e, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x17, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x18,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x50, 0x0a, 0x07, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65,
	0x61, 0x74, 0x12, 0x1a, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1a,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42,
	0x65, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x32, 0x8c, 0x02, 0x0a,
	0x08, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x0d, 0x42, 0x75, 0x6c,
	0x6c, 0x79, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x33,
	0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x12, 0x12, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x49, 0x64, 0x4c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x1a, 0x0d, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x6e, 0x69, 0x6c, 0x12, 0x43, 0x0a, 0x0d, 0x44, 0x6f, 0x6c, 0x65, 0x76, 0x45, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x41, 0x0a, 0x0b, 0x44, 0x4b, 0x52, 0x45,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x45, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x11, 0x5a, 0x0f, 0x63,
	0x6f, 0x64, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_registry_election_proto_rawDescOnce sync.Once
	file_registry_election_proto_rawDescData = file_registry_election_proto_rawDesc
)

func file_registry_election_proto_rawDescGZIP() []byte {
	file_registry_election_proto_rawDescOnce.Do(func() {
		file_registry_election_proto_rawDescData = protoimpl.X.CompressGZIP(file_registry_election_proto_rawDescData)
	})
	return file_registry_election_proto_rawDescData
}

var file_registry_election_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_registry_election_proto_goTypes = []interface{}{
	(*JoinRequest)(nil),      // 0: registry.JoinRequest
	(*JoinReply)(nil),        // 1: registry.JoinReply
	(*UpdateMessage)(nil),    // 2: registry.UpdateMessage
	(*UpdateResponse)(nil),   // 3: registry.UpdateResponse
	(*PeerInfo)(nil),         // 4: registry.PeerInfo
	(*HeartBeatMessage)(nil), // 5: registry.HeartBeatMessage
	(*Nil)(nil),              // 6: registry.nil
	(*IdLeader)(nil),         // 7: registry.IdLeader
	(*ElectionRequest)(nil),  // 8: registry.ElectionRequest
	(*ElectionReply)(nil),    // 9: registry.ElectionReply
}
var file_registry_election_proto_depIdxs = []int32{
	4, // 0: registry.JoinReply.peer_list:type_name -> registry.PeerInfo
	4, // 1: registry.UpdateMessage.peer_list:type_name -> registry.PeerInfo
	0, // 2: registry.Registry.JoinNetwork:input_type -> registry.JoinRequest
	2, // 3: registry.Update.UpdateNetwork:input_type -> registry.UpdateMessage
	5, // 4: registry.Service.HeartBeat:input_type -> registry.HeartBeatMessage
	8, // 5: registry.election.BullyElection:input_type -> registry.ElectionRequest
	7, // 6: registry.election.UpdateRegistry:input_type -> registry.IdLeader
	8, // 7: registry.election.DolevElection:input_type -> registry.ElectionRequest
	8, // 8: registry.election.DKRElection:input_type -> registry.ElectionRequest
	1, // 9: registry.Registry.JoinNetwork:output_type -> registry.JoinReply
	3, // 10: registry.Update.UpdateNetwork:output_type -> registry.UpdateResponse
	5, // 11: registry.Service.HeartBeat:output_type -> registry.HeartBeatMessage
	9, // 12: registry.election.BullyElection:output_type -> registry.ElectionReply
	6, // 13: registry.election.UpdateRegistry:output_type -> registry.nil
	9, // 14: registry.election.DolevElection:output_type -> registry.ElectionReply
	9, // 15: registry.election.DKRElection:output_type -> registry.ElectionReply
	9, // [9:16] is the sub-list for method output_type
	2, // [2:9] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_registry_election_proto_init() }
func file_registry_election_proto_init() {
	if File_registry_election_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_registry_election_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRequest); i {
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
		file_registry_election_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinReply); i {
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
		file_registry_election_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMessage); i {
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
		file_registry_election_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
		file_registry_election_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerInfo); i {
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
		file_registry_election_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartBeatMessage); i {
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
		file_registry_election_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nil); i {
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
		file_registry_election_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdLeader); i {
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
		file_registry_election_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElectionRequest); i {
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
		file_registry_election_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElectionReply); i {
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
			RawDescriptor: file_registry_election_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_registry_election_proto_goTypes,
		DependencyIndexes: file_registry_election_proto_depIdxs,
		MessageInfos:      file_registry_election_proto_msgTypes,
	}.Build()
	File_registry_election_proto = out.File
	file_registry_election_proto_rawDesc = nil
	file_registry_election_proto_goTypes = nil
	file_registry_election_proto_depIdxs = nil
}

// Copyright © 2022 Durudex
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: durudex/v1/user_session.proto

package durudexv1

import (
	pbtype "github.com/durudex/go-protobuf-type/pbtype"
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

// User session message.
type UserSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Session id.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Session user id.
	UserId []byte `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3,oneof" json:"user_id,omitempty"`
	// Session ip address.
	Ip string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	// Session expires in.
	ExpiresIn *pbtype.Timestamp `protobuf:"bytes,4,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
}

func (x *UserSession) Reset() {
	*x = UserSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durudex_v1_user_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSession) ProtoMessage() {}

func (x *UserSession) ProtoReflect() protoreflect.Message {
	mi := &file_durudex_v1_user_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSession.ProtoReflect.Descriptor instead.
func (*UserSession) Descriptor() ([]byte, []int) {
	return file_durudex_v1_user_session_proto_rawDescGZIP(), []int{0}
}

func (x *UserSession) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *UserSession) GetUserId() []byte {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *UserSession) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *UserSession) GetExpiresIn() *pbtype.Timestamp {
	if x != nil {
		return x.ExpiresIn
	}
	return nil
}

// Getting a user session request.
type GetUserSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Session id.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Session user id.
	UserId []byte `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserSessionRequest) Reset() {
	*x = GetUserSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durudex_v1_user_session_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserSessionRequest) ProtoMessage() {}

func (x *GetUserSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_durudex_v1_user_session_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserSessionRequest.ProtoReflect.Descriptor instead.
func (*GetUserSessionRequest) Descriptor() ([]byte, []int) {
	return file_durudex_v1_user_session_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserSessionRequest) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *GetUserSessionRequest) GetUserId() []byte {
	if x != nil {
		return x.UserId
	}
	return nil
}

// Getting a user session response.
type GetUserSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Session ip address.
	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	// Session expires in.
	ExpiresIn *pbtype.Timestamp `protobuf:"bytes,2,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
}

func (x *GetUserSessionResponse) Reset() {
	*x = GetUserSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durudex_v1_user_session_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserSessionResponse) ProtoMessage() {}

func (x *GetUserSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_durudex_v1_user_session_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserSessionResponse.ProtoReflect.Descriptor instead.
func (*GetUserSessionResponse) Descriptor() ([]byte, []int) {
	return file_durudex_v1_user_session_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserSessionResponse) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *GetUserSessionResponse) GetExpiresIn() *pbtype.Timestamp {
	if x != nil {
		return x.ExpiresIn
	}
	return nil
}

// Getting a user sessions request.
type GetUserSessionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Session user id.
	UserId []byte `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Query sort options.
	SortOptions *pbtype.SortOptions `protobuf:"bytes,2,opt,name=sort_options,json=sortOptions,proto3" json:"sort_options,omitempty"`
}

func (x *GetUserSessionsRequest) Reset() {
	*x = GetUserSessionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durudex_v1_user_session_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserSessionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserSessionsRequest) ProtoMessage() {}

func (x *GetUserSessionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_durudex_v1_user_session_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserSessionsRequest.ProtoReflect.Descriptor instead.
func (*GetUserSessionsRequest) Descriptor() ([]byte, []int) {
	return file_durudex_v1_user_session_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserSessionsRequest) GetUserId() []byte {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *GetUserSessionsRequest) GetSortOptions() *pbtype.SortOptions {
	if x != nil {
		return x.SortOptions
	}
	return nil
}

// Getting a user sessions response.
type GetUserSessionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User sessions.
	Sessions []*UserSession `protobuf:"bytes,1,rep,name=sessions,proto3" json:"sessions,omitempty"`
}

func (x *GetUserSessionsResponse) Reset() {
	*x = GetUserSessionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durudex_v1_user_session_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserSessionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserSessionsResponse) ProtoMessage() {}

func (x *GetUserSessionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_durudex_v1_user_session_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserSessionsResponse.ProtoReflect.Descriptor instead.
func (*GetUserSessionsResponse) Descriptor() ([]byte, []int) {
	return file_durudex_v1_user_session_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserSessionsResponse) GetSessions() []*UserSession {
	if x != nil {
		return x.Sessions
	}
	return nil
}

// Deleting a user session request.
type DeleteUserSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Session id.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Session user id.
	UserId []byte `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteUserSessionRequest) Reset() {
	*x = DeleteUserSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durudex_v1_user_session_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserSessionRequest) ProtoMessage() {}

func (x *DeleteUserSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_durudex_v1_user_session_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserSessionRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserSessionRequest) Descriptor() ([]byte, []int) {
	return file_durudex_v1_user_session_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteUserSessionRequest) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *DeleteUserSessionRequest) GetUserId() []byte {
	if x != nil {
		return x.UserId
	}
	return nil
}

// Deleting a user session response.
type DeleteUserSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteUserSessionResponse) Reset() {
	*x = DeleteUserSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durudex_v1_user_session_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserSessionResponse) ProtoMessage() {}

func (x *DeleteUserSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_durudex_v1_user_session_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserSessionResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserSessionResponse) Descriptor() ([]byte, []int) {
	return file_durudex_v1_user_session_proto_rawDescGZIP(), []int{6}
}

// Getting total session count request.
type GetTotalUserSessionCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Session user id.
	UserId []byte `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetTotalUserSessionCountRequest) Reset() {
	*x = GetTotalUserSessionCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durudex_v1_user_session_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTotalUserSessionCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTotalUserSessionCountRequest) ProtoMessage() {}

func (x *GetTotalUserSessionCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_durudex_v1_user_session_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTotalUserSessionCountRequest.ProtoReflect.Descriptor instead.
func (*GetTotalUserSessionCountRequest) Descriptor() ([]byte, []int) {
	return file_durudex_v1_user_session_proto_rawDescGZIP(), []int{7}
}

func (x *GetTotalUserSessionCountRequest) GetUserId() []byte {
	if x != nil {
		return x.UserId
	}
	return nil
}

// Getting total session count response.
type GetTotalUserSessionCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User session count.
	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetTotalUserSessionCountResponse) Reset() {
	*x = GetTotalUserSessionCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durudex_v1_user_session_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTotalUserSessionCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTotalUserSessionCountResponse) ProtoMessage() {}

func (x *GetTotalUserSessionCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_durudex_v1_user_session_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTotalUserSessionCountResponse.ProtoReflect.Descriptor instead.
func (*GetTotalUserSessionCountResponse) Descriptor() ([]byte, []int) {
	return file_durudex_v1_user_session_proto_rawDescGZIP(), []int{8}
}

func (x *GetTotalUserSessionCountResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_durudex_v1_user_session_proto protoreflect.FileDescriptor

var file_durudex_v1_user_session_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x64, 0x75, 0x72,
	0x75, 0x64, 0x65, 0x78, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x64, 0x75, 0x72, 0x75, 0x64,
	0x65, 0x78, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x0b, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x36, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x75,
	0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x60, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x36, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x75,
	0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x22,
	0x6f, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x3c, 0x0a, 0x0c, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64,
	0x65, 0x78, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x73, 0x6f, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x4e, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x43, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x1b, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x3a, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x38,
	0x0a, 0x20, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xa2, 0x03, 0x0a, 0x12, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x57, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x21, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x22, 0x2e, 0x64, 0x75,
	0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x64, 0x75, 0x72, 0x75,
	0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x75, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x2b, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2c, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xb3, 0x01,
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x76, 0x31,
	0x42, 0x10, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2f, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78,
	0x2d, 0x61, 0x75, 0x74, 0x68, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2f,
	0x76, 0x31, 0x3b, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x44,
	0x58, 0x58, 0xaa, 0x02, 0x0a, 0x44, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x0a, 0x44, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x44,
	0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x44, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_durudex_v1_user_session_proto_rawDescOnce sync.Once
	file_durudex_v1_user_session_proto_rawDescData = file_durudex_v1_user_session_proto_rawDesc
)

func file_durudex_v1_user_session_proto_rawDescGZIP() []byte {
	file_durudex_v1_user_session_proto_rawDescOnce.Do(func() {
		file_durudex_v1_user_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_durudex_v1_user_session_proto_rawDescData)
	})
	return file_durudex_v1_user_session_proto_rawDescData
}

var file_durudex_v1_user_session_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_durudex_v1_user_session_proto_goTypes = []interface{}{
	(*UserSession)(nil),                      // 0: durudex.v1.UserSession
	(*GetUserSessionRequest)(nil),            // 1: durudex.v1.GetUserSessionRequest
	(*GetUserSessionResponse)(nil),           // 2: durudex.v1.GetUserSessionResponse
	(*GetUserSessionsRequest)(nil),           // 3: durudex.v1.GetUserSessionsRequest
	(*GetUserSessionsResponse)(nil),          // 4: durudex.v1.GetUserSessionsResponse
	(*DeleteUserSessionRequest)(nil),         // 5: durudex.v1.DeleteUserSessionRequest
	(*DeleteUserSessionResponse)(nil),        // 6: durudex.v1.DeleteUserSessionResponse
	(*GetTotalUserSessionCountRequest)(nil),  // 7: durudex.v1.GetTotalUserSessionCountRequest
	(*GetTotalUserSessionCountResponse)(nil), // 8: durudex.v1.GetTotalUserSessionCountResponse
	(*pbtype.Timestamp)(nil),                 // 9: durudex.type.Timestamp
	(*pbtype.SortOptions)(nil),               // 10: durudex.type.SortOptions
}
var file_durudex_v1_user_session_proto_depIdxs = []int32{
	9,  // 0: durudex.v1.UserSession.expires_in:type_name -> durudex.type.Timestamp
	9,  // 1: durudex.v1.GetUserSessionResponse.expires_in:type_name -> durudex.type.Timestamp
	10, // 2: durudex.v1.GetUserSessionsRequest.sort_options:type_name -> durudex.type.SortOptions
	0,  // 3: durudex.v1.GetUserSessionsResponse.sessions:type_name -> durudex.v1.UserSession
	1,  // 4: durudex.v1.UserSessionService.GetUserSession:input_type -> durudex.v1.GetUserSessionRequest
	3,  // 5: durudex.v1.UserSessionService.GetUserSessions:input_type -> durudex.v1.GetUserSessionsRequest
	5,  // 6: durudex.v1.UserSessionService.DeleteUserSession:input_type -> durudex.v1.DeleteUserSessionRequest
	7,  // 7: durudex.v1.UserSessionService.GetTotalUserSessionCount:input_type -> durudex.v1.GetTotalUserSessionCountRequest
	2,  // 8: durudex.v1.UserSessionService.GetUserSession:output_type -> durudex.v1.GetUserSessionResponse
	4,  // 9: durudex.v1.UserSessionService.GetUserSessions:output_type -> durudex.v1.GetUserSessionsResponse
	6,  // 10: durudex.v1.UserSessionService.DeleteUserSession:output_type -> durudex.v1.DeleteUserSessionResponse
	8,  // 11: durudex.v1.UserSessionService.GetTotalUserSessionCount:output_type -> durudex.v1.GetTotalUserSessionCountResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_durudex_v1_user_session_proto_init() }
func file_durudex_v1_user_session_proto_init() {
	if File_durudex_v1_user_session_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_durudex_v1_user_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSession); i {
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
		file_durudex_v1_user_session_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserSessionRequest); i {
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
		file_durudex_v1_user_session_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserSessionResponse); i {
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
		file_durudex_v1_user_session_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserSessionsRequest); i {
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
		file_durudex_v1_user_session_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserSessionsResponse); i {
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
		file_durudex_v1_user_session_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserSessionRequest); i {
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
		file_durudex_v1_user_session_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserSessionResponse); i {
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
		file_durudex_v1_user_session_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTotalUserSessionCountRequest); i {
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
		file_durudex_v1_user_session_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTotalUserSessionCountResponse); i {
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
	file_durudex_v1_user_session_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_durudex_v1_user_session_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_durudex_v1_user_session_proto_goTypes,
		DependencyIndexes: file_durudex_v1_user_session_proto_depIdxs,
		MessageInfos:      file_durudex_v1_user_session_proto_msgTypes,
	}.Build()
	File_durudex_v1_user_session_proto = out.File
	file_durudex_v1_user_session_proto_rawDesc = nil
	file_durudex_v1_user_session_proto_goTypes = nil
	file_durudex_v1_user_session_proto_depIdxs = nil
}

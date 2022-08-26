// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/rpc/user/user.proto

package userv3

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	v31 "github.com/paralus/paralus/proto/types/commonpb/v3"
	v3 "github.com/paralus/paralus/proto/types/userpb/v3"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ApiKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Id       string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ApiKeyRequest) Reset() {
	*x = ApiKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiKeyRequest) ProtoMessage() {}

func (x *ApiKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiKeyRequest.ProtoReflect.Descriptor instead.
func (*ApiKeyRequest) Descriptor() ([]byte, []int) {
	return file_proto_rpc_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *ApiKeyRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ApiKeyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ApiKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModifiedAt *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=modifiedAt,proto3" json:"modifiedAt,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Id         string                 `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Key        string                 `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	Name       string                 `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ApiKeyResponse) Reset() {
	*x = ApiKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiKeyResponse) ProtoMessage() {}

func (x *ApiKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiKeyResponse.ProtoReflect.Descriptor instead.
func (*ApiKeyResponse) Descriptor() ([]byte, []int) {
	return file_proto_rpc_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *ApiKeyResponse) GetModifiedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ModifiedAt
	}
	return nil
}

func (x *ApiKeyResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ApiKeyResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ApiKeyResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ApiKeyResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UserListApiKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ApiKeyResponse `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *UserListApiKeysResponse) Reset() {
	*x = UserListApiKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListApiKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListApiKeysResponse) ProtoMessage() {}

func (x *UserListApiKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListApiKeysResponse.ProtoReflect.Descriptor instead.
func (*UserListApiKeysResponse) Descriptor() ([]byte, []int) {
	return file_proto_rpc_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserListApiKeysResponse) GetItems() []*ApiKeyResponse {
	if x != nil {
		return x.Items
	}
	return nil
}

type UserForgotPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *UserForgotPasswordRequest) Reset() {
	*x = UserForgotPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserForgotPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserForgotPasswordRequest) ProtoMessage() {}

func (x *UserForgotPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserForgotPasswordRequest.ProtoReflect.Descriptor instead.
func (*UserForgotPasswordRequest) Descriptor() ([]byte, []int) {
	return file_proto_rpc_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserForgotPasswordRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type UserForgotPasswordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecoveryLink string `protobuf:"bytes,1,opt,name=recoveryLink,proto3" json:"recoveryLink,omitempty"`
}

func (x *UserForgotPasswordResponse) Reset() {
	*x = UserForgotPasswordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_user_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserForgotPasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserForgotPasswordResponse) ProtoMessage() {}

func (x *UserForgotPasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_user_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserForgotPasswordResponse.ProtoReflect.Descriptor instead.
func (*UserForgotPasswordResponse) Descriptor() ([]byte, []int) {
	return file_proto_rpc_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *UserForgotPasswordResponse) GetRecoveryLink() string {
	if x != nil {
		return x.RecoveryLink
	}
	return ""
}

type UserDeleteApiKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserDeleteApiKeysResponse) Reset() {
	*x = UserDeleteApiKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_user_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDeleteApiKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDeleteApiKeysResponse) ProtoMessage() {}

func (x *UserDeleteApiKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_user_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDeleteApiKeysResponse.ProtoReflect.Descriptor instead.
func (*UserDeleteApiKeysResponse) Descriptor() ([]byte, []int) {
	return file_proto_rpc_user_user_proto_rawDescGZIP(), []int{5}
}

type CliConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CliConfigRequest) Reset() {
	*x = CliConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_user_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CliConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CliConfigRequest) ProtoMessage() {}

func (x *CliConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_user_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CliConfigRequest.ProtoReflect.Descriptor instead.
func (*CliConfigRequest) Descriptor() ([]byte, []int) {
	return file_proto_rpc_user_user_proto_rawDescGZIP(), []int{6}
}

var File_proto_rpc_user_user_proto protoreflect.FileDescriptor

var file_proto_rpc_user_user_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x70, 0x61, 0x72,
	0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2f, 0x76, 0x33, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x0d, 0x41,
	0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xbc, 0x01, 0x0a, 0x0e, 0x41, 0x70, 0x69,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x58, 0x0a, 0x17, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x70, 0x69, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0x37, 0x0a, 0x19, 0x55, 0x73, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x1a, 0x55, 0x73,
	0x65, 0x72, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x1b, 0x0a, 0x19,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x6c, 0x69,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xb7, 0x0b,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa2, 0x01,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x70,
	0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x1f, 0x2e,
	0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x52,
	0x92, 0x41, 0x36, 0x4a, 0x34, 0x0a, 0x03, 0x32, 0x30, 0x31, 0x12, 0x2d, 0x0a, 0x2b, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x75, 0x73, 0x65, 0x72,
	0x20, 0x69, 0x73, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x20, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22,
	0x0e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a,
	0x01, 0x2a, 0x12, 0x72, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x29,
	0x2e, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x23, 0x2e, 0x70, 0x61, 0x72, 0x61,
	0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x16,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x72, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x1f, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x1a, 0x1f, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x6e, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x2e, 0x70, 0x61, 0x72, 0x61,
	0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x23, 0x2e, 0x70, 0x61, 0x72,
	0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76,
	0x33, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x78, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6c,
	0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x1f, 0x2e, 0x70, 0x61, 0x72, 0x61,
	0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x22, 0x1a, 0x1d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65,
	0x7d, 0x3a, 0x01, 0x2a, 0x12, 0xc1, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x1a, 0x32, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64,
	0x65, 0x76, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5e, 0x92, 0x41, 0x36, 0x4a, 0x34, 0x0a,
	0x03, 0x32, 0x30, 0x34, 0x12, 0x2d, 0x0a, 0x2b, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64,
	0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x69, 0x73, 0x20, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c,
	0x6c, 0x79, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x2a, 0x1d, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x76, 0x33, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x82, 0x01, 0x0a, 0x11, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6c, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x29,
	0x2e, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6c, 0x69, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x61, 0x72, 0x61,
	0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79,
	0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x76, 0x33, 0x2f, 0x63, 0x6c, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x95, 0x01,
	0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79,
	0x73, 0x12, 0x26, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x70, 0x69, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x70, 0x61, 0x72, 0x61,
	0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x69, 0x4b,
	0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x61, 0x70,
	0x69, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x9e, 0x01, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x26, 0x2e, 0x70, 0x61,
	0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x2a,
	0x25, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xae, 0x01, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x46,
	0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x32, 0x2e,
	0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x67,
	0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x33, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x12, 0x27,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x66, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0xc5, 0x04, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x42, 0x09, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3b,
	0x75, 0x73, 0x65, 0x72, 0x76, 0x33, 0xa2, 0x02, 0x04, 0x50, 0x44, 0x52, 0x55, 0xaa, 0x02, 0x17,
	0x50, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x44, 0x65, 0x76, 0x2e, 0x52, 0x70, 0x63, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x2e, 0x56, 0x33, 0xca, 0x02, 0x17, 0x50, 0x61, 0x72, 0x61, 0x6c, 0x75,
	0x73, 0x5c, 0x44, 0x65, 0x76, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x56,
	0x33, 0xe2, 0x02, 0x23, 0x50, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x5c, 0x44, 0x65, 0x76, 0x5c,
	0x52, 0x70, 0x63, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x56, 0x33, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x50, 0x61, 0x72, 0x61, 0x6c, 0x75,
	0x73, 0x3a, 0x3a, 0x44, 0x65, 0x76, 0x3a, 0x3a, 0x52, 0x70, 0x63, 0x3a, 0x3a, 0x55, 0x73, 0x65,
	0x72, 0x3a, 0x3a, 0x56, 0x33, 0x92, 0x41, 0xe6, 0x02, 0x12, 0x2d, 0x0a, 0x17, 0x55, 0x73, 0x65,
	0x72, 0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x20,
	0x44, 0x65, 0x76, 0x32, 0x03, 0x32, 0x2e, 0x30, 0x2a, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x32, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x79, 0x61, 0x6d, 0x6c, 0x3a,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x79,
	0x61, 0x6d, 0x6c, 0x52, 0x50, 0x0a, 0x03, 0x34, 0x30, 0x33, 0x12, 0x49, 0x0a, 0x47, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x75, 0x73, 0x65, 0x72, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x68, 0x61,
	0x76, 0x65, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f,
	0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x3b, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x34, 0x0a, 0x2a,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20,
	0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x2e, 0x12, 0x06, 0x0a, 0x04, 0x9a, 0x02,
	0x01, 0x07, 0x5a, 0x3a, 0x0a, 0x27, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75,
	0x74, 0x68, 0x12, 0x19, 0x08, 0x02, 0x1a, 0x13, 0x58, 0x2d, 0x50, 0x41, 0x52, 0x41, 0x4c, 0x55,
	0x53, 0x2d, 0x41, 0x50, 0x49, 0x2d, 0x4b, 0x45, 0x59, 0x49, 0x44, 0x20, 0x02, 0x0a, 0x0f, 0x0a,
	0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x02, 0x08, 0x01, 0x62, 0x1f,
	0x0a, 0x0e, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00,
	0x0a, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rpc_user_user_proto_rawDescOnce sync.Once
	file_proto_rpc_user_user_proto_rawDescData = file_proto_rpc_user_user_proto_rawDesc
)

func file_proto_rpc_user_user_proto_rawDescGZIP() []byte {
	file_proto_rpc_user_user_proto_rawDescOnce.Do(func() {
		file_proto_rpc_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rpc_user_user_proto_rawDescData)
	})
	return file_proto_rpc_user_user_proto_rawDescData
}

var file_proto_rpc_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_rpc_user_user_proto_goTypes = []interface{}{
	(*ApiKeyRequest)(nil),              // 0: paralus.dev.rpc.user.v3.ApiKeyRequest
	(*ApiKeyResponse)(nil),             // 1: paralus.dev.rpc.user.v3.ApiKeyResponse
	(*UserListApiKeysResponse)(nil),    // 2: paralus.dev.rpc.user.v3.UserListApiKeysResponse
	(*UserForgotPasswordRequest)(nil),  // 3: paralus.dev.rpc.user.v3.UserForgotPasswordRequest
	(*UserForgotPasswordResponse)(nil), // 4: paralus.dev.rpc.user.v3.UserForgotPasswordResponse
	(*UserDeleteApiKeysResponse)(nil),  // 5: paralus.dev.rpc.user.v3.UserDeleteApiKeysResponse
	(*CliConfigRequest)(nil),           // 6: paralus.dev.rpc.user.v3.CliConfigRequest
	(*timestamppb.Timestamp)(nil),      // 7: google.protobuf.Timestamp
	(*v3.User)(nil),                    // 8: paralus.dev.types.user.v3.User
	(*v31.QueryOptions)(nil),           // 9: paralus.dev.types.common.v3.QueryOptions
	(*v3.UserList)(nil),                // 10: paralus.dev.types.user.v3.UserList
	(*v3.UserInfo)(nil),                // 11: paralus.dev.types.user.v3.UserInfo
	(*v31.HttpBody)(nil),               // 12: paralus.dev.types.common.v3.HttpBody
}
var file_proto_rpc_user_user_proto_depIdxs = []int32{
	7,  // 0: paralus.dev.rpc.user.v3.ApiKeyResponse.modifiedAt:type_name -> google.protobuf.Timestamp
	7,  // 1: paralus.dev.rpc.user.v3.ApiKeyResponse.createdAt:type_name -> google.protobuf.Timestamp
	1,  // 2: paralus.dev.rpc.user.v3.UserListApiKeysResponse.items:type_name -> paralus.dev.rpc.user.v3.ApiKeyResponse
	8,  // 3: paralus.dev.rpc.user.v3.UserService.CreateUser:input_type -> paralus.dev.types.user.v3.User
	9,  // 4: paralus.dev.rpc.user.v3.UserService.GetUsers:input_type -> paralus.dev.types.common.v3.QueryOptions
	8,  // 5: paralus.dev.rpc.user.v3.UserService.GetUser:input_type -> paralus.dev.types.user.v3.User
	8,  // 6: paralus.dev.rpc.user.v3.UserService.GetUserInfo:input_type -> paralus.dev.types.user.v3.User
	8,  // 7: paralus.dev.rpc.user.v3.UserService.UpdateUser:input_type -> paralus.dev.types.user.v3.User
	8,  // 8: paralus.dev.rpc.user.v3.UserService.DeleteUser:input_type -> paralus.dev.types.user.v3.User
	6,  // 9: paralus.dev.rpc.user.v3.UserService.DownloadCliConfig:input_type -> paralus.dev.rpc.user.v3.CliConfigRequest
	0,  // 10: paralus.dev.rpc.user.v3.UserService.UserListApiKeys:input_type -> paralus.dev.rpc.user.v3.ApiKeyRequest
	0,  // 11: paralus.dev.rpc.user.v3.UserService.UserDeleteApiKeys:input_type -> paralus.dev.rpc.user.v3.ApiKeyRequest
	3,  // 12: paralus.dev.rpc.user.v3.UserService.UserForgotPassword:input_type -> paralus.dev.rpc.user.v3.UserForgotPasswordRequest
	8,  // 13: paralus.dev.rpc.user.v3.UserService.CreateUser:output_type -> paralus.dev.types.user.v3.User
	10, // 14: paralus.dev.rpc.user.v3.UserService.GetUsers:output_type -> paralus.dev.types.user.v3.UserList
	8,  // 15: paralus.dev.rpc.user.v3.UserService.GetUser:output_type -> paralus.dev.types.user.v3.User
	11, // 16: paralus.dev.rpc.user.v3.UserService.GetUserInfo:output_type -> paralus.dev.types.user.v3.UserInfo
	8,  // 17: paralus.dev.rpc.user.v3.UserService.UpdateUser:output_type -> paralus.dev.types.user.v3.User
	5,  // 18: paralus.dev.rpc.user.v3.UserService.DeleteUser:output_type -> paralus.dev.rpc.user.v3.UserDeleteApiKeysResponse
	12, // 19: paralus.dev.rpc.user.v3.UserService.DownloadCliConfig:output_type -> paralus.dev.types.common.v3.HttpBody
	2,  // 20: paralus.dev.rpc.user.v3.UserService.UserListApiKeys:output_type -> paralus.dev.rpc.user.v3.UserListApiKeysResponse
	5,  // 21: paralus.dev.rpc.user.v3.UserService.UserDeleteApiKeys:output_type -> paralus.dev.rpc.user.v3.UserDeleteApiKeysResponse
	4,  // 22: paralus.dev.rpc.user.v3.UserService.UserForgotPassword:output_type -> paralus.dev.rpc.user.v3.UserForgotPasswordResponse
	13, // [13:23] is the sub-list for method output_type
	3,  // [3:13] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_proto_rpc_user_user_proto_init() }
func file_proto_rpc_user_user_proto_init() {
	if File_proto_rpc_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_rpc_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiKeyRequest); i {
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
		file_proto_rpc_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiKeyResponse); i {
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
		file_proto_rpc_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListApiKeysResponse); i {
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
		file_proto_rpc_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserForgotPasswordRequest); i {
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
		file_proto_rpc_user_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserForgotPasswordResponse); i {
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
		file_proto_rpc_user_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDeleteApiKeysResponse); i {
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
		file_proto_rpc_user_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CliConfigRequest); i {
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
			RawDescriptor: file_proto_rpc_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_rpc_user_user_proto_goTypes,
		DependencyIndexes: file_proto_rpc_user_user_proto_depIdxs,
		MessageInfos:      file_proto_rpc_user_user_proto_msgTypes,
	}.Build()
	File_proto_rpc_user_user_proto = out.File
	file_proto_rpc_user_user_proto_rawDesc = nil
	file_proto_rpc_user_user_proto_goTypes = nil
	file_proto_rpc_user_user_proto_depIdxs = nil
}

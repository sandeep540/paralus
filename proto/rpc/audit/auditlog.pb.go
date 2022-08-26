// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/rpc/audit/auditlog.proto

package eventv1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	v3 "github.com/paralus/paralus/proto/types/commonpb/v3"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuditLogQueryFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type          string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	User          string   `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Client        string   `protobuf:"bytes,3,opt,name=client,proto3" json:"client,omitempty"`
	Timefrom      string   `protobuf:"bytes,4,opt,name=timefrom,proto3" json:"timefrom,omitempty"`
	Portal        string   `protobuf:"bytes,5,opt,name=portal,proto3" json:"portal,omitempty"`
	Cluster       string   `protobuf:"bytes,6,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Namespace     string   `protobuf:"bytes,7,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Kind          string   `protobuf:"bytes,8,opt,name=kind,proto3" json:"kind,omitempty"`
	Method        string   `protobuf:"bytes,9,opt,name=method,proto3" json:"method,omitempty"`
	Projects      []string `protobuf:"bytes,10,rep,name=projects,proto3" json:"projects,omitempty"`
	QueryString   string   `protobuf:"bytes,11,opt,name=queryString,proto3" json:"queryString,omitempty"`
	DashboardData bool     `protobuf:"varint,12,opt,name=dashboardData,proto3" json:"dashboardData,omitempty"`
}

func (x *AuditLogQueryFilter) Reset() {
	*x = AuditLogQueryFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_audit_auditlog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditLogQueryFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditLogQueryFilter) ProtoMessage() {}

func (x *AuditLogQueryFilter) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_audit_auditlog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditLogQueryFilter.ProtoReflect.Descriptor instead.
func (*AuditLogQueryFilter) Descriptor() ([]byte, []int) {
	return file_proto_rpc_audit_auditlog_proto_rawDescGZIP(), []int{0}
}

func (x *AuditLogQueryFilter) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AuditLogQueryFilter) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *AuditLogQueryFilter) GetClient() string {
	if x != nil {
		return x.Client
	}
	return ""
}

func (x *AuditLogQueryFilter) GetTimefrom() string {
	if x != nil {
		return x.Timefrom
	}
	return ""
}

func (x *AuditLogQueryFilter) GetPortal() string {
	if x != nil {
		return x.Portal
	}
	return ""
}

func (x *AuditLogQueryFilter) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *AuditLogQueryFilter) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *AuditLogQueryFilter) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *AuditLogQueryFilter) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *AuditLogQueryFilter) GetProjects() []string {
	if x != nil {
		return x.Projects
	}
	return nil
}

func (x *AuditLogQueryFilter) GetQueryString() string {
	if x != nil {
		return x.QueryString
	}
	return ""
}

func (x *AuditLogQueryFilter) GetDashboardData() bool {
	if x != nil {
		return x.DashboardData
	}
	return false
}

type GetAuditLogSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *v3.Metadata         `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Filter   *AuditLogQueryFilter `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *GetAuditLogSearchRequest) Reset() {
	*x = GetAuditLogSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_audit_auditlog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuditLogSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuditLogSearchRequest) ProtoMessage() {}

func (x *GetAuditLogSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_audit_auditlog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuditLogSearchRequest.ProtoReflect.Descriptor instead.
func (*GetAuditLogSearchRequest) Descriptor() ([]byte, []int) {
	return file_proto_rpc_audit_auditlog_proto_rawDescGZIP(), []int{1}
}

func (x *GetAuditLogSearchRequest) GetMetadata() *v3.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *GetAuditLogSearchRequest) GetFilter() *AuditLogQueryFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type GetAuditLogSearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *structpb.Struct `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GetAuditLogSearchResponse) Reset() {
	*x = GetAuditLogSearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_audit_auditlog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuditLogSearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuditLogSearchResponse) ProtoMessage() {}

func (x *GetAuditLogSearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_audit_auditlog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuditLogSearchResponse.ProtoReflect.Descriptor instead.
func (*GetAuditLogSearchResponse) Descriptor() ([]byte, []int) {
	return file_proto_rpc_audit_auditlog_proto_rawDescGZIP(), []int{2}
}

func (x *GetAuditLogSearchResponse) GetResult() *structpb.Struct {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_proto_rpc_audit_auditlog_proto protoreflect.FileDescriptor

var file_proto_rpc_audit_auditlog_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x16, 0x72, 0x65, 0x70, 0x2e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x02, 0x0a, 0x13, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x22, 0xa2,
	0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x43,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x72, 0x65, 0x70, 0x2e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x22, 0x4c, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c,
	0x6f, 0x67, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x32, 0xdb, 0x02, 0x0a, 0x0f, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xac, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x30, 0x2e, 0x72, 0x65, 0x70, 0x2e, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x72, 0x65, 0x70, 0x2e, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x32, 0x12, 0x30, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x72, 0x6c, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x2a, 0x7d, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x6c, 0x6f, 0x67, 0x12, 0x98, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x4c, 0x6f, 0x67, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x30,
	0x2e, 0x72, 0x65, 0x70, 0x2e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74,
	0x4c, 0x6f, 0x67, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x31, 0x2e, 0x72, 0x65, 0x70, 0x2e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6c, 0x6f, 0x67, 0x42,
	0xbd, 0x04, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x70, 0x2e, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0d,
	0x41, 0x75, 0x64, 0x69, 0x74, 0x6c, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x72, 0x61,
	0x6c, 0x75, 0x73, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x3b, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x52, 0x46, 0x45, 0xaa, 0x02, 0x16, 0x52, 0x65, 0x70, 0x2e,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x16, 0x52, 0x65, 0x70, 0x5c, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f,
	0x72, 0x6b, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x22, 0x52, 0x65,
	0x70, 0x5c, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x5c, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x19, 0x52, 0x65, 0x70, 0x3a, 0x3a, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x3a, 0x3a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x92, 0x41, 0xdf, 0x02,
	0x12, 0x26, 0x0a, 0x10, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x20, 0x53, 0x65, 0x72,
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
	file_proto_rpc_audit_auditlog_proto_rawDescOnce sync.Once
	file_proto_rpc_audit_auditlog_proto_rawDescData = file_proto_rpc_audit_auditlog_proto_rawDesc
)

func file_proto_rpc_audit_auditlog_proto_rawDescGZIP() []byte {
	file_proto_rpc_audit_auditlog_proto_rawDescOnce.Do(func() {
		file_proto_rpc_audit_auditlog_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rpc_audit_auditlog_proto_rawDescData)
	})
	return file_proto_rpc_audit_auditlog_proto_rawDescData
}

var file_proto_rpc_audit_auditlog_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_rpc_audit_auditlog_proto_goTypes = []interface{}{
	(*AuditLogQueryFilter)(nil),       // 0: rep.framework.event.v1.AuditLogQueryFilter
	(*GetAuditLogSearchRequest)(nil),  // 1: rep.framework.event.v1.GetAuditLogSearchRequest
	(*GetAuditLogSearchResponse)(nil), // 2: rep.framework.event.v1.GetAuditLogSearchResponse
	(*v3.Metadata)(nil),               // 3: paralus.dev.types.common.v3.Metadata
	(*structpb.Struct)(nil),           // 4: google.protobuf.Struct
}
var file_proto_rpc_audit_auditlog_proto_depIdxs = []int32{
	3, // 0: rep.framework.event.v1.GetAuditLogSearchRequest.metadata:type_name -> paralus.dev.types.common.v3.Metadata
	0, // 1: rep.framework.event.v1.GetAuditLogSearchRequest.filter:type_name -> rep.framework.event.v1.AuditLogQueryFilter
	4, // 2: rep.framework.event.v1.GetAuditLogSearchResponse.result:type_name -> google.protobuf.Struct
	1, // 3: rep.framework.event.v1.AuditLogService.GetAuditLog:input_type -> rep.framework.event.v1.GetAuditLogSearchRequest
	1, // 4: rep.framework.event.v1.AuditLogService.GetAuditLogByProjects:input_type -> rep.framework.event.v1.GetAuditLogSearchRequest
	2, // 5: rep.framework.event.v1.AuditLogService.GetAuditLog:output_type -> rep.framework.event.v1.GetAuditLogSearchResponse
	2, // 6: rep.framework.event.v1.AuditLogService.GetAuditLogByProjects:output_type -> rep.framework.event.v1.GetAuditLogSearchResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_rpc_audit_auditlog_proto_init() }
func file_proto_rpc_audit_auditlog_proto_init() {
	if File_proto_rpc_audit_auditlog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_rpc_audit_auditlog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditLogQueryFilter); i {
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
		file_proto_rpc_audit_auditlog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuditLogSearchRequest); i {
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
		file_proto_rpc_audit_auditlog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuditLogSearchResponse); i {
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
			RawDescriptor: file_proto_rpc_audit_auditlog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_rpc_audit_auditlog_proto_goTypes,
		DependencyIndexes: file_proto_rpc_audit_auditlog_proto_depIdxs,
		MessageInfos:      file_proto_rpc_audit_auditlog_proto_msgTypes,
	}.Build()
	File_proto_rpc_audit_auditlog_proto = out.File
	file_proto_rpc_audit_auditlog_proto_rawDesc = nil
	file_proto_rpc_audit_auditlog_proto_goTypes = nil
	file_proto_rpc_audit_auditlog_proto_depIdxs = nil
}

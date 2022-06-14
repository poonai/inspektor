// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: api.proto

package apiproto

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

type AuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *AuthRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AuthRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
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
	return file_api_proto_rawDescGZIP(), []int{1}
}

type AuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups      []string `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
	ExpiresAt   int64    `protobuf:"varint,2,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	Passthrough bool     `protobuf:"varint,3,opt,name=passthrough,proto3" json:"passthrough,omitempty"`
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *AuthResponse) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *AuthResponse) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

func (x *AuthResponse) GetPassthrough() bool {
	if x != nil {
		return x.Passthrough
	}
	return false
}

type DataSourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataSourceName string `protobuf:"bytes,1,opt,name=data_source_name,json=dataSourceName,proto3" json:"data_source_name,omitempty"`
}

func (x *DataSourceResponse) Reset() {
	*x = DataSourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSourceResponse) ProtoMessage() {}

func (x *DataSourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSourceResponse.ProtoReflect.Descriptor instead.
func (*DataSourceResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *DataSourceResponse) GetDataSourceName() string {
	if x != nil {
		return x.DataSourceName
	}
	return ""
}

type InspektorPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WasmByteCode []byte `protobuf:"bytes,1,opt,name=wasm_byte_code,json=wasmByteCode,proto3" json:"wasm_byte_code,omitempty"`
}

func (x *InspektorPolicy) Reset() {
	*x = InspektorPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InspektorPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InspektorPolicy) ProtoMessage() {}

func (x *InspektorPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InspektorPolicy.ProtoReflect.Descriptor instead.
func (*InspektorPolicy) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *InspektorPolicy) GetWasmByteCode() []byte {
	if x != nil {
		return x.WasmByteCode
	}
	return nil
}

type MetricsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []*Metric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
	Groups  []string  `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *MetricsRequest) Reset() {
	*x = MetricsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsRequest) ProtoMessage() {}

func (x *MetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsRequest.ProtoReflect.Descriptor instead.
func (*MetricsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *MetricsRequest) GetMetrics() []*Metric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *MetricsRequest) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

type Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollectionName string   `protobuf:"bytes,1,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	PropertyName   []string `protobuf:"bytes,2,rep,name=property_name,json=propertyName,proto3" json:"property_name,omitempty"`
}

func (x *Metric) Reset() {
	*x = Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric) ProtoMessage() {}

func (x *Metric) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metric.ProtoReflect.Descriptor instead.
func (*Metric) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *Metric) GetCollectionName() string {
	if x != nil {
		return x.CollectionName
	}
	return ""
}

func (x *Metric) GetPropertyName() []string {
	if x != nil {
		return x.PropertyName
	}
	return nil
}

type IntegrationConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CloudWatchConfig *CloudWatchConfig `protobuf:"bytes,1,opt,name=cloud_watch_config,json=cloudWatchConfig,proto3" json:"cloud_watch_config,omitempty"`
	AuditLogConfig   *AuditLogConfig   `protobuf:"bytes,2,opt,name=audit_log_config,json=auditLogConfig,proto3" json:"audit_log_config,omitempty"`
}

func (x *IntegrationConfigResponse) Reset() {
	*x = IntegrationConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegrationConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegrationConfigResponse) ProtoMessage() {}

func (x *IntegrationConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegrationConfigResponse.ProtoReflect.Descriptor instead.
func (*IntegrationConfigResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

func (x *IntegrationConfigResponse) GetCloudWatchConfig() *CloudWatchConfig {
	if x != nil {
		return x.CloudWatchConfig
	}
	return nil
}

func (x *IntegrationConfigResponse) GetAuditLogConfig() *AuditLogConfig {
	if x != nil {
		return x.AuditLogConfig
	}
	return nil
}

type CloudWatchConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CredType      string `protobuf:"bytes,1,opt,name=cred_type,json=credType,proto3" json:"cred_type,omitempty"`
	RegionName    string `protobuf:"bytes,2,opt,name=region_name,json=regionName,proto3" json:"region_name,omitempty"`
	AccessKey     string `protobuf:"bytes,3,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	SecretKey     string `protobuf:"bytes,4,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	LogGroupName  string `protobuf:"bytes,5,opt,name=log_group_name,json=logGroupName,proto3" json:"log_group_name,omitempty"`
	LogStreamName string `protobuf:"bytes,6,opt,name=log_stream_name,json=logStreamName,proto3" json:"log_stream_name,omitempty"`
}

func (x *CloudWatchConfig) Reset() {
	*x = CloudWatchConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudWatchConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudWatchConfig) ProtoMessage() {}

func (x *CloudWatchConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudWatchConfig.ProtoReflect.Descriptor instead.
func (*CloudWatchConfig) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{8}
}

func (x *CloudWatchConfig) GetCredType() string {
	if x != nil {
		return x.CredType
	}
	return ""
}

func (x *CloudWatchConfig) GetRegionName() string {
	if x != nil {
		return x.RegionName
	}
	return ""
}

func (x *CloudWatchConfig) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *CloudWatchConfig) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *CloudWatchConfig) GetLogGroupName() string {
	if x != nil {
		return x.LogGroupName
	}
	return ""
}

func (x *CloudWatchConfig) GetLogStreamName() string {
	if x != nil {
		return x.LogStreamName
	}
	return ""
}

type AuditLogConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogPrefix string `protobuf:"bytes,1,opt,name=log_prefix,json=logPrefix,proto3" json:"log_prefix,omitempty"`
}

func (x *AuditLogConfig) Reset() {
	*x = AuditLogConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditLogConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditLogConfig) ProtoMessage() {}

func (x *AuditLogConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditLogConfig.ProtoReflect.Descriptor instead.
func (*AuditLogConfig) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{9}
}

func (x *AuditLogConfig) GetLogPrefix() string {
	if x != nil {
		return x.LogPrefix
	}
	return ""
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x22, 0x46, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x67, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x73, 0x73,
	0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x70,
	0x61, 0x73, 0x73, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x22, 0x3e, 0x0a, 0x12, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x28, 0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x0f, 0x49, 0x6e,
	0x73, 0x70, 0x65, 0x6b, 0x74, 0x6f, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x24, 0x0a,
	0x0e, 0x77, 0x61, 0x73, 0x6d, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x77, 0x61, 0x73, 0x6d, 0x42, 0x79, 0x74, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x22, 0x4f, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x22, 0x56, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x27,
	0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x9f, 0x01, 0x0a,
	0x19, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x12, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x5f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x57, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x10, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x57, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x3d, 0x0a, 0x10, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xdc,
	0x01, 0x0a, 0x10, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x57, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12,
	0x24, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x6f, 0x67, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2f, 0x0a,
	0x0e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x32, 0x9a,
	0x02, 0x0a, 0x09, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x6b, 0x74, 0x6f, 0x72, 0x12, 0x2d, 0x0a, 0x04,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x06, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x6b, 0x74, 0x6f,
	0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x00, 0x30, 0x01, 0x12, 0x36, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x0a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_proto_goTypes = []interface{}{
	(*AuthRequest)(nil),               // 0: api.AuthRequest
	(*Empty)(nil),                     // 1: api.Empty
	(*AuthResponse)(nil),              // 2: api.AuthResponse
	(*DataSourceResponse)(nil),        // 3: api.DataSourceResponse
	(*InspektorPolicy)(nil),           // 4: api.InspektorPolicy
	(*MetricsRequest)(nil),            // 5: api.MetricsRequest
	(*Metric)(nil),                    // 6: api.Metric
	(*IntegrationConfigResponse)(nil), // 7: api.IntegrationConfigResponse
	(*CloudWatchConfig)(nil),          // 8: api.CloudWatchConfig
	(*AuditLogConfig)(nil),            // 9: api.AuditLogConfig
}
var file_api_proto_depIdxs = []int32{
	6, // 0: api.MetricsRequest.metrics:type_name -> api.Metric
	8, // 1: api.IntegrationConfigResponse.cloud_watch_config:type_name -> api.CloudWatchConfig
	9, // 2: api.IntegrationConfigResponse.audit_log_config:type_name -> api.AuditLogConfig
	0, // 3: api.Inspektor.Auth:input_type -> api.AuthRequest
	1, // 4: api.Inspektor.Policy:input_type -> api.Empty
	1, // 5: api.Inspektor.GetDataSource:input_type -> api.Empty
	5, // 6: api.Inspektor.SendMetrics:input_type -> api.MetricsRequest
	1, // 7: api.Inspektor.GetIntegrationConfig:input_type -> api.Empty
	2, // 8: api.Inspektor.Auth:output_type -> api.AuthResponse
	4, // 9: api.Inspektor.Policy:output_type -> api.InspektorPolicy
	3, // 10: api.Inspektor.GetDataSource:output_type -> api.DataSourceResponse
	1, // 11: api.Inspektor.SendMetrics:output_type -> api.Empty
	7, // 12: api.Inspektor.GetIntegrationConfig:output_type -> api.IntegrationConfigResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRequest); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResponse); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSourceResponse); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InspektorPolicy); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricsRequest); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metric); i {
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
		file_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegrationConfigResponse); i {
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
		file_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudWatchConfig); i {
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
		file_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditLogConfig); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}

// Copied from https://github.com/pytorch/serve/blob/8c23585d2453f230c411721028ad4b07e58cc7dd/frontend/server/src/main/resources/proto/management.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: management.proto

package torchserve

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

type ManagementResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Response string of different management API calls.
	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ManagementResponse) Reset() {
	*x = ManagementResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagementResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagementResponse) ProtoMessage() {}

func (x *ManagementResponse) ProtoReflect() protoreflect.Message {
	mi := &file_management_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagementResponse.ProtoReflect.Descriptor instead.
func (*ManagementResponse) Descriptor() ([]byte, []int) {
	return file_management_proto_rawDescGZIP(), []int{0}
}

func (x *ManagementResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type DescribeModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of model to describe.
	ModelName string `protobuf:"bytes,1,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"` //required
	// Version of model to describe.
	ModelVersion string `protobuf:"bytes,2,opt,name=model_version,json=modelVersion,proto3" json:"model_version,omitempty"` //optional
	// Customized metadata
	Customized bool `protobuf:"varint,3,opt,name=customized,proto3" json:"customized,omitempty"` //optional
}

func (x *DescribeModelRequest) Reset() {
	*x = DescribeModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeModelRequest) ProtoMessage() {}

func (x *DescribeModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_management_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeModelRequest.ProtoReflect.Descriptor instead.
func (*DescribeModelRequest) Descriptor() ([]byte, []int) {
	return file_management_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeModelRequest) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *DescribeModelRequest) GetModelVersion() string {
	if x != nil {
		return x.ModelVersion
	}
	return ""
}

func (x *DescribeModelRequest) GetCustomized() bool {
	if x != nil {
		return x.Customized
	}
	return false
}

type ListModelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Use this parameter to specify the maximum number of items to return. When this value is present, TorchServe does not return more than the specified number of items, but it might return fewer. This value is optional. If you include a value, it must be between 1 and 1000, inclusive. If you do not include a value, it defaults to 100.
	Limit int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"` //optional
	// The token to retrieve the next set of results. TorchServe provides the token when the response from a previous call has more results than the maximum page size.
	NextPageToken int32 `protobuf:"varint,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"` //optional
}

func (x *ListModelsRequest) Reset() {
	*x = ListModelsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListModelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListModelsRequest) ProtoMessage() {}

func (x *ListModelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_management_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListModelsRequest.ProtoReflect.Descriptor instead.
func (*ListModelsRequest) Descriptor() ([]byte, []int) {
	return file_management_proto_rawDescGZIP(), []int{2}
}

func (x *ListModelsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListModelsRequest) GetNextPageToken() int32 {
	if x != nil {
		return x.NextPageToken
	}
	return 0
}

type RegisterModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Inference batch size, default: 1.
	BatchSize int32 `protobuf:"varint,1,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"` //optional
	// Inference handler entry-point. This value will override handler in MANIFEST.json if present.
	Handler string `protobuf:"bytes,2,opt,name=handler,proto3" json:"handler,omitempty"` //optional
	// Number of initial workers, default: 0.
	InitialWorkers int32 `protobuf:"varint,3,opt,name=initial_workers,json=initialWorkers,proto3" json:"initial_workers,omitempty"` //optional
	// Maximum delay for batch aggregation, default: 100.
	MaxBatchDelay int32 `protobuf:"varint,4,opt,name=max_batch_delay,json=maxBatchDelay,proto3" json:"max_batch_delay,omitempty"` //optional
	// Name of model. This value will override modelName in MANIFEST.json if present.
	ModelName string `protobuf:"bytes,5,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"` //optional
	// Maximum time, in seconds, the TorchServe waits for a response from the model inference code, default: 120.
	ResponseTimeout int32 `protobuf:"varint,6,opt,name=response_timeout,json=responseTimeout,proto3" json:"response_timeout,omitempty"` //optional
	// Runtime for the model custom service code. This value will override runtime in MANIFEST.json if present.
	Runtime string `protobuf:"bytes,7,opt,name=runtime,proto3" json:"runtime,omitempty"` //optional
	// Decides whether creation of worker synchronous or not, default: false.
	Synchronous bool `protobuf:"varint,8,opt,name=synchronous,proto3" json:"synchronous,omitempty"` //optional
	// Model archive download url, support local file or HTTP(s) protocol.
	Url string `protobuf:"bytes,9,opt,name=url,proto3" json:"url,omitempty"` //required
	// Decides whether S3 SSE KMS enabled or not, default: false.
	S3SseKms bool `protobuf:"varint,10,opt,name=s3_sse_kms,json=s3SseKms,proto3" json:"s3_sse_kms,omitempty"` //optional
}

func (x *RegisterModelRequest) Reset() {
	*x = RegisterModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterModelRequest) ProtoMessage() {}

func (x *RegisterModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_management_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterModelRequest.ProtoReflect.Descriptor instead.
func (*RegisterModelRequest) Descriptor() ([]byte, []int) {
	return file_management_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterModelRequest) GetBatchSize() int32 {
	if x != nil {
		return x.BatchSize
	}
	return 0
}

func (x *RegisterModelRequest) GetHandler() string {
	if x != nil {
		return x.Handler
	}
	return ""
}

func (x *RegisterModelRequest) GetInitialWorkers() int32 {
	if x != nil {
		return x.InitialWorkers
	}
	return 0
}

func (x *RegisterModelRequest) GetMaxBatchDelay() int32 {
	if x != nil {
		return x.MaxBatchDelay
	}
	return 0
}

func (x *RegisterModelRequest) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *RegisterModelRequest) GetResponseTimeout() int32 {
	if x != nil {
		return x.ResponseTimeout
	}
	return 0
}

func (x *RegisterModelRequest) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *RegisterModelRequest) GetSynchronous() bool {
	if x != nil {
		return x.Synchronous
	}
	return false
}

func (x *RegisterModelRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *RegisterModelRequest) GetS3SseKms() bool {
	if x != nil {
		return x.S3SseKms
	}
	return false
}

type ScaleWorkerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of model to scale workers.
	ModelName string `protobuf:"bytes,1,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"` //required
	// Model version.
	ModelVersion string `protobuf:"bytes,2,opt,name=model_version,json=modelVersion,proto3" json:"model_version,omitempty"` //optional
	// Maximum number of worker processes.
	MaxWorker int32 `protobuf:"varint,3,opt,name=max_worker,json=maxWorker,proto3" json:"max_worker,omitempty"` //optional
	// Minimum number of worker processes.
	MinWorker int32 `protobuf:"varint,4,opt,name=min_worker,json=minWorker,proto3" json:"min_worker,omitempty"` //optional
	// Number of GPU worker processes to create.
	NumberGpu int32 `protobuf:"varint,5,opt,name=number_gpu,json=numberGpu,proto3" json:"number_gpu,omitempty"` //optional
	// Decides whether the call is synchronous or not, default: false.
	Synchronous bool `protobuf:"varint,6,opt,name=synchronous,proto3" json:"synchronous,omitempty"` //optional
	// Waiting up to the specified wait time if necessary for a worker to complete all pending requests. Use 0 to terminate backend worker process immediately. Use -1 for wait infinitely.
	Timeout int32 `protobuf:"varint,7,opt,name=timeout,proto3" json:"timeout,omitempty"` //optional
}

func (x *ScaleWorkerRequest) Reset() {
	*x = ScaleWorkerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScaleWorkerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScaleWorkerRequest) ProtoMessage() {}

func (x *ScaleWorkerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_management_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScaleWorkerRequest.ProtoReflect.Descriptor instead.
func (*ScaleWorkerRequest) Descriptor() ([]byte, []int) {
	return file_management_proto_rawDescGZIP(), []int{4}
}

func (x *ScaleWorkerRequest) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *ScaleWorkerRequest) GetModelVersion() string {
	if x != nil {
		return x.ModelVersion
	}
	return ""
}

func (x *ScaleWorkerRequest) GetMaxWorker() int32 {
	if x != nil {
		return x.MaxWorker
	}
	return 0
}

func (x *ScaleWorkerRequest) GetMinWorker() int32 {
	if x != nil {
		return x.MinWorker
	}
	return 0
}

func (x *ScaleWorkerRequest) GetNumberGpu() int32 {
	if x != nil {
		return x.NumberGpu
	}
	return 0
}

func (x *ScaleWorkerRequest) GetSynchronous() bool {
	if x != nil {
		return x.Synchronous
	}
	return false
}

func (x *ScaleWorkerRequest) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type SetDefaultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of model whose default version needs to be updated.
	ModelName string `protobuf:"bytes,1,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"` //required
	// Version of model to be set as default version for the model
	ModelVersion string `protobuf:"bytes,2,opt,name=model_version,json=modelVersion,proto3" json:"model_version,omitempty"` //required
}

func (x *SetDefaultRequest) Reset() {
	*x = SetDefaultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDefaultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDefaultRequest) ProtoMessage() {}

func (x *SetDefaultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_management_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDefaultRequest.ProtoReflect.Descriptor instead.
func (*SetDefaultRequest) Descriptor() ([]byte, []int) {
	return file_management_proto_rawDescGZIP(), []int{5}
}

func (x *SetDefaultRequest) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *SetDefaultRequest) GetModelVersion() string {
	if x != nil {
		return x.ModelVersion
	}
	return ""
}

type UnregisterModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of model to unregister.
	ModelName string `protobuf:"bytes,1,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"` //required
	// Name of model to unregister.
	ModelVersion string `protobuf:"bytes,2,opt,name=model_version,json=modelVersion,proto3" json:"model_version,omitempty"` //optional
}

func (x *UnregisterModelRequest) Reset() {
	*x = UnregisterModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_management_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnregisterModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnregisterModelRequest) ProtoMessage() {}

func (x *UnregisterModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_management_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnregisterModelRequest.ProtoReflect.Descriptor instead.
func (*UnregisterModelRequest) Descriptor() ([]byte, []int) {
	return file_management_proto_rawDescGZIP(), []int{6}
}

func (x *UnregisterModelRequest) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *UnregisterModelRequest) GetModelVersion() string {
	if x != nil {
		return x.ModelVersion
	}
	return ""
}

var File_management_proto protoreflect.FileDescriptor

var file_management_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x21, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x79, 0x74, 0x6f, 0x72, 0x63, 0x68, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x26, 0x0a, 0x12, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x7a, 0x0a,
	0x14, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x22, 0x51, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6e,
	0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xd6, 0x02, 0x0a,
	0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x27,
	0x0a, 0x0f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x6d, 0x61, 0x78, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x12,
	0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29,
	0x0a, 0x10, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x6f,
	0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72,
	0x6f, 0x6e, 0x6f, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x0a, 0x73, 0x33, 0x5f, 0x73, 0x73,
	0x65, 0x5f, 0x6b, 0x6d, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x33, 0x53,
	0x73, 0x65, 0x4b, 0x6d, 0x73, 0x22, 0xf1, 0x01, 0x0a, 0x12, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12,
	0x1d, 0x0a, 0x0a, 0x6d, 0x69, 0x6e, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x69, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x1d,
	0x0a, 0x0a, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x67, 0x70, 0x75, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x47, 0x70, 0x75, 0x12, 0x20, 0x0a,
	0x0b, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x6f, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x6f, 0x75, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x57, 0x0a, 0x11, 0x53, 0x65, 0x74,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x5c, 0x0a, 0x16, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x32, 0xa0, 0x06, 0x0a, 0x15, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x41,
	0x50, 0x49, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x0d, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x37, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x70, 0x79, 0x74, 0x6f, 0x72, 0x63, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x79, 0x74, 0x6f,
	0x72, 0x63, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7b,
	0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x34, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x70, 0x79, 0x74, 0x6f, 0x72, 0x63, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x35, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x79, 0x74, 0x6f, 0x72, 0x63, 0x68,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x81, 0x01, 0x0a, 0x0d,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x37, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x70, 0x79, 0x74, 0x6f, 0x72, 0x63, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x79, 0x74,
	0x6f, 0x72, 0x63, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x7d, 0x0a, 0x0b, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x35,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x79, 0x74, 0x6f, 0x72, 0x63, 0x68, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x79, 0x74, 0x6f,
	0x72, 0x63, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7b,
	0x0a, 0x0a, 0x53, 0x65, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x34, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x70, 0x79, 0x74, 0x6f, 0x72, 0x63, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x53, 0x65, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x35, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x79, 0x74, 0x6f, 0x72, 0x63, 0x68,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x85, 0x01, 0x0a, 0x0f,
	0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x39, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x79, 0x74, 0x6f, 0x72, 0x63, 0x68, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x70, 0x79, 0x74, 0x6f, 0x72, 0x63, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x02, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_management_proto_rawDescOnce sync.Once
	file_management_proto_rawDescData = file_management_proto_rawDesc
)

func file_management_proto_rawDescGZIP() []byte {
	file_management_proto_rawDescOnce.Do(func() {
		file_management_proto_rawDescData = protoimpl.X.CompressGZIP(file_management_proto_rawDescData)
	})
	return file_management_proto_rawDescData
}

var file_management_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_management_proto_goTypes = []interface{}{
	(*ManagementResponse)(nil),     // 0: org.pytorch.serve.grpc.management.ManagementResponse
	(*DescribeModelRequest)(nil),   // 1: org.pytorch.serve.grpc.management.DescribeModelRequest
	(*ListModelsRequest)(nil),      // 2: org.pytorch.serve.grpc.management.ListModelsRequest
	(*RegisterModelRequest)(nil),   // 3: org.pytorch.serve.grpc.management.RegisterModelRequest
	(*ScaleWorkerRequest)(nil),     // 4: org.pytorch.serve.grpc.management.ScaleWorkerRequest
	(*SetDefaultRequest)(nil),      // 5: org.pytorch.serve.grpc.management.SetDefaultRequest
	(*UnregisterModelRequest)(nil), // 6: org.pytorch.serve.grpc.management.UnregisterModelRequest
}
var file_management_proto_depIdxs = []int32{
	1, // 0: org.pytorch.serve.grpc.management.ManagementAPIsService.DescribeModel:input_type -> org.pytorch.serve.grpc.management.DescribeModelRequest
	2, // 1: org.pytorch.serve.grpc.management.ManagementAPIsService.ListModels:input_type -> org.pytorch.serve.grpc.management.ListModelsRequest
	3, // 2: org.pytorch.serve.grpc.management.ManagementAPIsService.RegisterModel:input_type -> org.pytorch.serve.grpc.management.RegisterModelRequest
	4, // 3: org.pytorch.serve.grpc.management.ManagementAPIsService.ScaleWorker:input_type -> org.pytorch.serve.grpc.management.ScaleWorkerRequest
	5, // 4: org.pytorch.serve.grpc.management.ManagementAPIsService.SetDefault:input_type -> org.pytorch.serve.grpc.management.SetDefaultRequest
	6, // 5: org.pytorch.serve.grpc.management.ManagementAPIsService.UnregisterModel:input_type -> org.pytorch.serve.grpc.management.UnregisterModelRequest
	0, // 6: org.pytorch.serve.grpc.management.ManagementAPIsService.DescribeModel:output_type -> org.pytorch.serve.grpc.management.ManagementResponse
	0, // 7: org.pytorch.serve.grpc.management.ManagementAPIsService.ListModels:output_type -> org.pytorch.serve.grpc.management.ManagementResponse
	0, // 8: org.pytorch.serve.grpc.management.ManagementAPIsService.RegisterModel:output_type -> org.pytorch.serve.grpc.management.ManagementResponse
	0, // 9: org.pytorch.serve.grpc.management.ManagementAPIsService.ScaleWorker:output_type -> org.pytorch.serve.grpc.management.ManagementResponse
	0, // 10: org.pytorch.serve.grpc.management.ManagementAPIsService.SetDefault:output_type -> org.pytorch.serve.grpc.management.ManagementResponse
	0, // 11: org.pytorch.serve.grpc.management.ManagementAPIsService.UnregisterModel:output_type -> org.pytorch.serve.grpc.management.ManagementResponse
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_management_proto_init() }
func file_management_proto_init() {
	if File_management_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_management_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagementResponse); i {
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
		file_management_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeModelRequest); i {
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
		file_management_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListModelsRequest); i {
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
		file_management_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterModelRequest); i {
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
		file_management_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScaleWorkerRequest); i {
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
		file_management_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetDefaultRequest); i {
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
		file_management_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnregisterModelRequest); i {
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
			RawDescriptor: file_management_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_management_proto_goTypes,
		DependencyIndexes: file_management_proto_depIdxs,
		MessageInfos:      file_management_proto_msgTypes,
	}.Build()
	File_management_proto = out.File
	file_management_proto_rawDesc = nil
	file_management_proto_goTypes = nil
	file_management_proto_depIdxs = nil
}

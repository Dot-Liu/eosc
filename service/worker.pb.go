// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: worker.proto

package service

import (
	eosc "github.com/eolinker/eosc"
	config "github.com/eolinker/eosc/config"
	traffic "github.com/eolinker/eosc/traffic"
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

type WorkerStatusCode int32

const (
	WorkerStatusCode_SUCCESS WorkerStatusCode = 0
	WorkerStatusCode_FAIL    WorkerStatusCode = 1
	WorkerStatusCode_SKILL   WorkerStatusCode = 2
)

// Enum value maps for WorkerStatusCode.
var (
	WorkerStatusCode_name = map[int32]string{
		0: "SUCCESS",
		1: "FAIL",
		2: "SKILL",
	}
	WorkerStatusCode_value = map[string]int32{
		"SUCCESS": 0,
		"FAIL":    1,
		"SKILL":   2,
	}
)

func (x WorkerStatusCode) Enum() *WorkerStatusCode {
	p := new(WorkerStatusCode)
	*p = x
	return p
}

func (x WorkerStatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkerStatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_worker_proto_enumTypes[0].Descriptor()
}

func (WorkerStatusCode) Type() protoreflect.EnumType {
	return &file_worker_proto_enumTypes[0]
}

func (x WorkerStatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkerStatusCode.Descriptor instead.
func (WorkerStatusCode) EnumDescriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{0}
}

type WorkerHelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hello string `protobuf:"bytes,1,opt,name=hello,proto3" json:"hello,omitempty"`
}

func (x *WorkerHelloRequest) Reset() {
	*x = WorkerHelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerHelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerHelloRequest) ProtoMessage() {}

func (x *WorkerHelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerHelloRequest.ProtoReflect.Descriptor instead.
func (*WorkerHelloRequest) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{0}
}

func (x *WorkerHelloRequest) GetHello() string {
	if x != nil {
		return x.Hello
	}
	return ""
}

type WorkerDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *WorkerDeleteRequest) Reset() {
	*x = WorkerDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerDeleteRequest) ProtoMessage() {}

func (x *WorkerDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerDeleteRequest.ProtoReflect.Descriptor instead.
func (*WorkerDeleteRequest) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{1}
}

func (x *WorkerDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type WorkerSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Profession string `protobuf:"bytes,2,opt,name=profession,proto3" json:"profession,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Driver     string `protobuf:"bytes,4,opt,name=driver,proto3" json:"driver,omitempty"`
	Body       []byte `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *WorkerSetRequest) Reset() {
	*x = WorkerSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerSetRequest) ProtoMessage() {}

func (x *WorkerSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerSetRequest.ProtoReflect.Descriptor instead.
func (*WorkerSetRequest) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{2}
}

func (x *WorkerSetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WorkerSetRequest) GetProfession() string {
	if x != nil {
		return x.Profession
	}
	return ""
}

func (x *WorkerSetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WorkerSetRequest) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *WorkerSetRequest) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type WorkerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  WorkerStatusCode `protobuf:"varint,1,opt,name=status,proto3,enum=service.WorkerStatusCode" json:"status,omitempty"`
	Message string           `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *WorkerResponse) Reset() {
	*x = WorkerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerResponse) ProtoMessage() {}

func (x *WorkerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerResponse.ProtoReflect.Descriptor instead.
func (*WorkerResponse) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{3}
}

func (x *WorkerResponse) GetStatus() WorkerStatusCode {
	if x != nil {
		return x.Status
	}
	return WorkerStatusCode_SUCCESS
}

func (x *WorkerResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type WorkerLoadArg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Traffic         []*traffic.PbTraffic     `protobuf:"bytes,1,rep,name=traffic,proto3" json:"traffic,omitempty"`
	ListensMsg      *config.ListensMsg       `protobuf:"bytes,2,opt,name=listensMsg,proto3" json:"listensMsg,omitempty"`
	ExtenderSetting map[string]string        `protobuf:"bytes,3,rep,name=extenderSetting,proto3" json:"extenderSetting,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Professions     []*eosc.ProfessionConfig `protobuf:"bytes,4,rep,name=professions,proto3" json:"professions,omitempty"`
	Workers         []*eosc.WorkerConfig     `protobuf:"bytes,5,rep,name=workers,proto3" json:"workers,omitempty"`
}

func (x *WorkerLoadArg) Reset() {
	*x = WorkerLoadArg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerLoadArg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerLoadArg) ProtoMessage() {}

func (x *WorkerLoadArg) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerLoadArg.ProtoReflect.Descriptor instead.
func (*WorkerLoadArg) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{4}
}

func (x *WorkerLoadArg) GetTraffic() []*traffic.PbTraffic {
	if x != nil {
		return x.Traffic
	}
	return nil
}

func (x *WorkerLoadArg) GetListensMsg() *config.ListensMsg {
	if x != nil {
		return x.ListensMsg
	}
	return nil
}

func (x *WorkerLoadArg) GetExtenderSetting() map[string]string {
	if x != nil {
		return x.ExtenderSetting
	}
	return nil
}

func (x *WorkerLoadArg) GetProfessions() []*eosc.ProfessionConfig {
	if x != nil {
		return x.Professions
	}
	return nil
}

func (x *WorkerLoadArg) GetWorkers() []*eosc.WorkerConfig {
	if x != nil {
		return x.Workers
	}
	return nil
}

type ResetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Professions []*eosc.ProfessionConfig `protobuf:"bytes,2,rep,name=professions,proto3" json:"professions,omitempty"`
	Workers     []*eosc.WorkerConfig     `protobuf:"bytes,3,rep,name=workers,proto3" json:"workers,omitempty"`
}

func (x *ResetRequest) Reset() {
	*x = ResetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetRequest) ProtoMessage() {}

func (x *ResetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetRequest.ProtoReflect.Descriptor instead.
func (*ResetRequest) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{5}
}

func (x *ResetRequest) GetProfessions() []*eosc.ProfessionConfig {
	if x != nil {
		return x.Professions
	}
	return nil
}

func (x *ResetRequest) GetWorkers() []*eosc.WorkerConfig {
	if x != nil {
		return x.Workers
	}
	return nil
}

type StatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRequest) ProtoMessage() {}

func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRequest.ProtoReflect.Descriptor instead.
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{6}
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{7}
}

type WorkerAddExtender struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Extenders map[string]string `protobuf:"bytes,1,rep,name=Extenders,proto3" json:"Extenders,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *WorkerAddExtender) Reset() {
	*x = WorkerAddExtender{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerAddExtender) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerAddExtender) ProtoMessage() {}

func (x *WorkerAddExtender) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerAddExtender.ProtoReflect.Descriptor instead.
func (*WorkerAddExtender) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{8}
}

func (x *WorkerAddExtender) GetExtenders() map[string]string {
	if x != nil {
		return x.Extenders
	}
	return nil
}

type WorkerDelExtender struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Extenders []string `protobuf:"bytes,1,rep,name=extenders,proto3" json:"extenders,omitempty"`
}

func (x *WorkerDelExtender) Reset() {
	*x = WorkerDelExtender{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerDelExtender) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerDelExtender) ProtoMessage() {}

func (x *WorkerDelExtender) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerDelExtender.ProtoReflect.Descriptor instead.
func (*WorkerDelExtender) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{9}
}

func (x *WorkerDelExtender) GetExtenders() []string {
	if x != nil {
		return x.Extenders
	}
	return nil
}

var File_worker_proto protoreflect.FileDescriptor

var file_worker_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x12, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x22,
	0x25, 0x0a, 0x13, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x82, 0x01, 0x0a, 0x10, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x5d, 0x0a, 0x0e, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xfb, 0x02, 0x0a, 0x0d, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x72, 0x67, 0x12, 0x2c, 0x0a, 0x07,
	0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x62, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x52, 0x07, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x12, 0x33, 0x0a, 0x0a, 0x6c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x73, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x73,
	0x4d, 0x73, 0x67, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x73, 0x4d, 0x73, 0x67, 0x12,
	0x55, 0x0a, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x72, 0x67,
	0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x3b, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x73, 0x1a, 0x42, 0x0a, 0x14, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x7c, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x66,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9a, 0x01, 0x0a, 0x11, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x41, 0x64, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x47, 0x0a, 0x09, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x41, 0x64, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x3c, 0x0a, 0x0e, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x31, 0x0a, 0x11, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x44, 0x65, 0x6c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x2a, 0x34, 0x0a, 0x10, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41,
	0x49, 0x4c, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x10, 0x02, 0x32,
	0xe2, 0x04, 0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x46, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x08, 0x73, 0x65, 0x74,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b,
	0x0a, 0x03, 0x73, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x04, 0x70,
	0x69, 0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x05, 0x72,
	0x65, 0x73, 0x65, 0x74, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0b, 0x61, 0x64, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x41, 0x64, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x1a, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x10, 0x64, 0x65, 0x6c,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x1a, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x44, 0x65,
	0x6c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6f, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x2f, 0x65, 0x6f, 0x73, 0x63,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_worker_proto_rawDescOnce sync.Once
	file_worker_proto_rawDescData = file_worker_proto_rawDesc
)

func file_worker_proto_rawDescGZIP() []byte {
	file_worker_proto_rawDescOnce.Do(func() {
		file_worker_proto_rawDescData = protoimpl.X.CompressGZIP(file_worker_proto_rawDescData)
	})
	return file_worker_proto_rawDescData
}

var file_worker_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_worker_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_worker_proto_goTypes = []interface{}{
	(WorkerStatusCode)(0),         // 0: service.WorkerStatusCode
	(*WorkerHelloRequest)(nil),    // 1: service.WorkerHelloRequest
	(*WorkerDeleteRequest)(nil),   // 2: service.WorkerDeleteRequest
	(*WorkerSetRequest)(nil),      // 3: service.WorkerSetRequest
	(*WorkerResponse)(nil),        // 4: service.WorkerResponse
	(*WorkerLoadArg)(nil),         // 5: service.WorkerLoadArg
	(*ResetRequest)(nil),          // 6: service.ResetRequest
	(*StatusRequest)(nil),         // 7: service.StatusRequest
	(*StatusResponse)(nil),        // 8: service.StatusResponse
	(*WorkerAddExtender)(nil),     // 9: service.WorkerAddExtender
	(*WorkerDelExtender)(nil),     // 10: service.WorkerDelExtender
	nil,                           // 11: service.WorkerLoadArg.ExtenderSettingEntry
	nil,                           // 12: service.WorkerAddExtender.ExtendersEntry
	(*traffic.PbTraffic)(nil),     // 13: service.PbTraffic
	(*config.ListensMsg)(nil),     // 14: service.ListensMsg
	(*eosc.ProfessionConfig)(nil), // 15: service.ProfessionConfig
	(*eosc.WorkerConfig)(nil),     // 16: service.WorkerConfig
}
var file_worker_proto_depIdxs = []int32{
	0,  // 0: service.WorkerResponse.status:type_name -> service.WorkerStatusCode
	13, // 1: service.WorkerLoadArg.traffic:type_name -> service.PbTraffic
	14, // 2: service.WorkerLoadArg.listensMsg:type_name -> service.ListensMsg
	11, // 3: service.WorkerLoadArg.extenderSetting:type_name -> service.WorkerLoadArg.ExtenderSettingEntry
	15, // 4: service.WorkerLoadArg.professions:type_name -> service.ProfessionConfig
	16, // 5: service.WorkerLoadArg.workers:type_name -> service.WorkerConfig
	15, // 6: service.ResetRequest.professions:type_name -> service.ProfessionConfig
	16, // 7: service.ResetRequest.workers:type_name -> service.WorkerConfig
	12, // 8: service.WorkerAddExtender.Extenders:type_name -> service.WorkerAddExtender.ExtendersEntry
	2,  // 9: service.WorkerService.deleteCheck:input_type -> service.WorkerDeleteRequest
	3,  // 10: service.WorkerService.setCheck:input_type -> service.WorkerSetRequest
	2,  // 11: service.WorkerService.delete:input_type -> service.WorkerDeleteRequest
	3,  // 12: service.WorkerService.set:input_type -> service.WorkerSetRequest
	1,  // 13: service.WorkerService.ping:input_type -> service.WorkerHelloRequest
	6,  // 14: service.WorkerService.reset:input_type -> service.ResetRequest
	7,  // 15: service.WorkerService.status:input_type -> service.StatusRequest
	9,  // 16: service.WorkerService.addExtender:input_type -> service.WorkerAddExtender
	10, // 17: service.WorkerService.delExtenderCheck:input_type -> service.WorkerDelExtender
	4,  // 18: service.WorkerService.deleteCheck:output_type -> service.WorkerResponse
	4,  // 19: service.WorkerService.setCheck:output_type -> service.WorkerResponse
	4,  // 20: service.WorkerService.delete:output_type -> service.WorkerResponse
	4,  // 21: service.WorkerService.set:output_type -> service.WorkerResponse
	4,  // 22: service.WorkerService.ping:output_type -> service.WorkerResponse
	4,  // 23: service.WorkerService.reset:output_type -> service.WorkerResponse
	8,  // 24: service.WorkerService.status:output_type -> service.StatusResponse
	4,  // 25: service.WorkerService.addExtender:output_type -> service.WorkerResponse
	4,  // 26: service.WorkerService.delExtenderCheck:output_type -> service.WorkerResponse
	18, // [18:27] is the sub-list for method output_type
	9,  // [9:18] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_worker_proto_init() }
func file_worker_proto_init() {
	if File_worker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_worker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerHelloRequest); i {
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
		file_worker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerDeleteRequest); i {
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
		file_worker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerSetRequest); i {
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
		file_worker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerResponse); i {
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
		file_worker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerLoadArg); i {
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
		file_worker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetRequest); i {
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
		file_worker_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusRequest); i {
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
		file_worker_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_worker_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerAddExtender); i {
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
		file_worker_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerDelExtender); i {
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
			RawDescriptor: file_worker_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_worker_proto_goTypes,
		DependencyIndexes: file_worker_proto_depIdxs,
		EnumInfos:         file_worker_proto_enumTypes,
		MessageInfos:      file_worker_proto_msgTypes,
	}.Build()
	File_worker_proto = out.File
	file_worker_proto_rawDesc = nil
	file_worker_proto_goTypes = nil
	file_worker_proto_depIdxs = nil
}

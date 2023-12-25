// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: ekto/ekto.proto

package ekto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type QuerierMethod int32

const (
	QuerierMethod_FIND   QuerierMethod = 0
	QuerierMethod_LIST   QuerierMethod = 1
	QuerierMethod_CREATE QuerierMethod = 2
	QuerierMethod_UPDATE QuerierMethod = 3
	QuerierMethod_DELETE QuerierMethod = 4
)

// Enum value maps for QuerierMethod.
var (
	QuerierMethod_name = map[int32]string{
		0: "FIND",
		1: "LIST",
		2: "CREATE",
		3: "UPDATE",
		4: "DELETE",
	}
	QuerierMethod_value = map[string]int32{
		"FIND":   0,
		"LIST":   1,
		"CREATE": 2,
		"UPDATE": 3,
		"DELETE": 4,
	}
)

func (x QuerierMethod) Enum() *QuerierMethod {
	p := new(QuerierMethod)
	*p = x
	return p
}

func (x QuerierMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QuerierMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_ekto_ekto_proto_enumTypes[0].Descriptor()
}

func (QuerierMethod) Type() protoreflect.EnumType {
	return &file_ekto_ekto_proto_enumTypes[0]
}

func (x QuerierMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QuerierMethod.Descriptor instead.
func (QuerierMethod) EnumDescriptor() ([]byte, []int) {
	return file_ekto_ekto_proto_rawDescGZIP(), []int{0}
}

type MQOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Handles   string `protobuf:"bytes,1,opt,name=handles,proto3" json:"handles,omitempty"`
	EventName string `protobuf:"bytes,2,opt,name=event_name,json=eventName,proto3" json:"event_name,omitempty"`
}

func (x *MQOptions) Reset() {
	*x = MQOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ekto_ekto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MQOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MQOptions) ProtoMessage() {}

func (x *MQOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ekto_ekto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MQOptions.ProtoReflect.Descriptor instead.
func (*MQOptions) Descriptor() ([]byte, []int) {
	return file_ekto_ekto_proto_rawDescGZIP(), []int{0}
}

func (x *MQOptions) GetHandles() string {
	if x != nil {
		return x.Handles
	}
	return ""
}

func (x *MQOptions) GetEventName() string {
	if x != nil {
		return x.EventName
	}
	return ""
}

type DBOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DBOptions) Reset() {
	*x = DBOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ekto_ekto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBOptions) ProtoMessage() {}

func (x *DBOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ekto_ekto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBOptions.ProtoReflect.Descriptor instead.
func (*DBOptions) Descriptor() ([]byte, []int) {
	return file_ekto_ekto_proto_rawDescGZIP(), []int{1}
}

func (x *DBOptions) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type QuerierOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method QuerierMethod `protobuf:"varint,1,opt,name=method,proto3,enum=ekto.QuerierMethod" json:"method,omitempty"`
}

func (x *QuerierOptions) Reset() {
	*x = QuerierOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ekto_ekto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerierOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerierOptions) ProtoMessage() {}

func (x *QuerierOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ekto_ekto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuerierOptions.ProtoReflect.Descriptor instead.
func (*QuerierOptions) Descriptor() ([]byte, []int) {
	return file_ekto_ekto_proto_rawDescGZIP(), []int{2}
}

func (x *QuerierOptions) GetMethod() QuerierMethod {
	if x != nil {
		return x.Method
	}
	return QuerierMethod_FIND
}

type MessageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queryable bool `protobuf:"varint,1,opt,name=queryable,proto3" json:"queryable,omitempty"`
}

func (x *MessageOptions) Reset() {
	*x = MessageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ekto_ekto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageOptions) ProtoMessage() {}

func (x *MessageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ekto_ekto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageOptions.ProtoReflect.Descriptor instead.
func (*MessageOptions) Descriptor() ([]byte, []int) {
	return file_ekto_ekto_proto_rawDescGZIP(), []int{3}
}

func (x *MessageOptions) GetQueryable() bool {
	if x != nil {
		return x.Queryable
	}
	return false
}

type Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mq      *MQOptions      `protobuf:"bytes,1,opt,name=mq,proto3" json:"mq,omitempty"`
	Querier *QuerierOptions `protobuf:"bytes,2,opt,name=querier,proto3" json:"querier,omitempty"`
}

func (x *Options) Reset() {
	*x = Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ekto_ekto_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options) ProtoMessage() {}

func (x *Options) ProtoReflect() protoreflect.Message {
	mi := &file_ekto_ekto_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options.ProtoReflect.Descriptor instead.
func (*Options) Descriptor() ([]byte, []int) {
	return file_ekto_ekto_proto_rawDescGZIP(), []int{4}
}

func (x *Options) GetMq() *MQOptions {
	if x != nil {
		return x.Mq
	}
	return nil
}

func (x *Options) GetQuerier() *QuerierOptions {
	if x != nil {
		return x.Querier
	}
	return nil
}

type SvcOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Db *DBOptions `protobuf:"bytes,1,opt,name=db,proto3" json:"db,omitempty"`
}

func (x *SvcOptions) Reset() {
	*x = SvcOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ekto_ekto_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SvcOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SvcOptions) ProtoMessage() {}

func (x *SvcOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ekto_ekto_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SvcOptions.ProtoReflect.Descriptor instead.
func (*SvcOptions) Descriptor() ([]byte, []int) {
	return file_ekto_ekto_proto_rawDescGZIP(), []int{5}
}

func (x *SvcOptions) GetDb() *DBOptions {
	if x != nil {
		return x.Db
	}
	return nil
}

var file_ekto_ekto_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*SvcOptions)(nil),
		Field:         50386,
		Name:          "ekto.svc",
		Tag:           "bytes,50386,opt,name=svc",
		Filename:      "ekto/ekto.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*Options)(nil),
		Field:         50386,
		Name:          "ekto.dev",
		Tag:           "bytes,50386,opt,name=dev",
		Filename:      "ekto/ekto.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*MessageOptions)(nil),
		Field:         50386,
		Name:          "ekto.msg",
		Tag:           "bytes,50386,opt,name=msg",
		Filename:      "ekto/ekto.proto",
	},
}

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional ekto.SvcOptions svc = 50386;
	E_Svc = &file_ekto_ekto_proto_extTypes[0]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional ekto.Options dev = 50386;
	E_Dev = &file_ekto_ekto_proto_extTypes[1]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional ekto.MessageOptions msg = 50386;
	E_Msg = &file_ekto_ekto_proto_extTypes[2]
)

var File_ekto_ekto_proto protoreflect.FileDescriptor

var file_ekto_ekto_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x6b, 0x74, 0x6f, 0x2f, 0x65, 0x6b, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x65, 0x6b, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x09, 0x4d, 0x51, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x1f, 0x0a, 0x09, 0x44, 0x42, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x3d, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x13, 0x2e, 0x65, 0x6b, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65,
	0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22,
	0x2e, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x22,
	0x5a, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x02, 0x6d, 0x71,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x65, 0x6b, 0x74, 0x6f, 0x2e, 0x4d, 0x51,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x02, 0x6d, 0x71, 0x12, 0x2e, 0x0a, 0x07, 0x71,
	0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65,
	0x6b, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x22, 0x2d, 0x0a, 0x0a, 0x53,
	0x76, 0x63, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x02, 0x64, 0x62, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x65, 0x6b, 0x74, 0x6f, 0x2e, 0x44, 0x42, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x02, 0x64, 0x62, 0x2a, 0x47, 0x0a, 0x0d, 0x51, 0x75,
	0x65, 0x72, 0x69, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x46,
	0x49, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54,
	0x45, 0x10, 0x04, 0x3a, 0x45, 0x0a, 0x03, 0x73, 0x76, 0x63, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd2, 0x89, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x6b, 0x74, 0x6f, 0x2e, 0x53, 0x76, 0x63, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x03, 0x73, 0x76, 0x63, 0x3a, 0x41, 0x0a, 0x03, 0x64, 0x65,
	0x76, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xd2, 0x89, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x65, 0x6b, 0x74, 0x6f,
	0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x03, 0x64, 0x65, 0x76, 0x3a, 0x49, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd2, 0x89, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x65, 0x6b, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6b, 0x74, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f,
	0x65, 0x6b, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x65, 0x6b, 0x74, 0x6f, 0x2f, 0x65, 0x6b, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ekto_ekto_proto_rawDescOnce sync.Once
	file_ekto_ekto_proto_rawDescData = file_ekto_ekto_proto_rawDesc
)

func file_ekto_ekto_proto_rawDescGZIP() []byte {
	file_ekto_ekto_proto_rawDescOnce.Do(func() {
		file_ekto_ekto_proto_rawDescData = protoimpl.X.CompressGZIP(file_ekto_ekto_proto_rawDescData)
	})
	return file_ekto_ekto_proto_rawDescData
}

var file_ekto_ekto_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ekto_ekto_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ekto_ekto_proto_goTypes = []interface{}{
	(QuerierMethod)(0),                  // 0: ekto.QuerierMethod
	(*MQOptions)(nil),                   // 1: ekto.MQOptions
	(*DBOptions)(nil),                   // 2: ekto.DBOptions
	(*QuerierOptions)(nil),              // 3: ekto.QuerierOptions
	(*MessageOptions)(nil),              // 4: ekto.MessageOptions
	(*Options)(nil),                     // 5: ekto.Options
	(*SvcOptions)(nil),                  // 6: ekto.SvcOptions
	(*descriptorpb.ServiceOptions)(nil), // 7: google.protobuf.ServiceOptions
	(*descriptorpb.MethodOptions)(nil),  // 8: google.protobuf.MethodOptions
	(*descriptorpb.MessageOptions)(nil), // 9: google.protobuf.MessageOptions
}
var file_ekto_ekto_proto_depIdxs = []int32{
	0,  // 0: ekto.QuerierOptions.method:type_name -> ekto.QuerierMethod
	1,  // 1: ekto.Options.mq:type_name -> ekto.MQOptions
	3,  // 2: ekto.Options.querier:type_name -> ekto.QuerierOptions
	2,  // 3: ekto.SvcOptions.db:type_name -> ekto.DBOptions
	7,  // 4: ekto.svc:extendee -> google.protobuf.ServiceOptions
	8,  // 5: ekto.dev:extendee -> google.protobuf.MethodOptions
	9,  // 6: ekto.msg:extendee -> google.protobuf.MessageOptions
	6,  // 7: ekto.svc:type_name -> ekto.SvcOptions
	5,  // 8: ekto.dev:type_name -> ekto.Options
	4,  // 9: ekto.msg:type_name -> ekto.MessageOptions
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	7,  // [7:10] is the sub-list for extension type_name
	4,  // [4:7] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_ekto_ekto_proto_init() }
func file_ekto_ekto_proto_init() {
	if File_ekto_ekto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ekto_ekto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MQOptions); i {
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
		file_ekto_ekto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBOptions); i {
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
		file_ekto_ekto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuerierOptions); i {
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
		file_ekto_ekto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageOptions); i {
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
		file_ekto_ekto_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options); i {
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
		file_ekto_ekto_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SvcOptions); i {
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
			RawDescriptor: file_ekto_ekto_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_ekto_ekto_proto_goTypes,
		DependencyIndexes: file_ekto_ekto_proto_depIdxs,
		EnumInfos:         file_ekto_ekto_proto_enumTypes,
		MessageInfos:      file_ekto_ekto_proto_msgTypes,
		ExtensionInfos:    file_ekto_ekto_proto_extTypes,
	}.Build()
	File_ekto_ekto_proto = out.File
	file_ekto_ekto_proto_rawDesc = nil
	file_ekto_ekto_proto_goTypes = nil
	file_ekto_ekto_proto_depIdxs = nil
}

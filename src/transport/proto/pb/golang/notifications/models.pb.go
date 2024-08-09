// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: src/transport/proto/src/notifications/models.proto

package __

import (
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

type UserNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID               uint64                 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Type             string                 `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	SenderID         uint64                 `protobuf:"varint,3,opt,name=SenderID,proto3" json:"SenderID,omitempty"`
	RecipientID      uint64                 `protobuf:"varint,4,opt,name=RecipientID,proto3" json:"RecipientID,omitempty"`
	Title            string                 `protobuf:"bytes,5,opt,name=Title,proto3" json:"Title,omitempty"`
	TitleI18N        string                 `protobuf:"bytes,6,opt,name=TitleI18n,proto3" json:"TitleI18n,omitempty"`
	Text             string                 `protobuf:"bytes,7,opt,name=Text,proto3" json:"Text,omitempty"`
	TextI18N         string                 `protobuf:"bytes,8,opt,name=TextI18n,proto3" json:"TextI18n,omitempty"`
	CreatedTimestamp *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=CreatedTimestamp,proto3" json:"CreatedTimestamp,omitempty"`
	ReadTimestamp    *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=ReadTimestamp,proto3" json:"ReadTimestamp,omitempty"`
	RemovedTimestamp *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=RemovedTimestamp,proto3" json:"RemovedTimestamp,omitempty"`
}

func (x *UserNotification) Reset() {
	*x = UserNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_transport_proto_src_notifications_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserNotification) ProtoMessage() {}

func (x *UserNotification) ProtoReflect() protoreflect.Message {
	mi := &file_src_transport_proto_src_notifications_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserNotification.ProtoReflect.Descriptor instead.
func (*UserNotification) Descriptor() ([]byte, []int) {
	return file_src_transport_proto_src_notifications_models_proto_rawDescGZIP(), []int{0}
}

func (x *UserNotification) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UserNotification) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UserNotification) GetSenderID() uint64 {
	if x != nil {
		return x.SenderID
	}
	return 0
}

func (x *UserNotification) GetRecipientID() uint64 {
	if x != nil {
		return x.RecipientID
	}
	return 0
}

func (x *UserNotification) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UserNotification) GetTitleI18N() string {
	if x != nil {
		return x.TitleI18N
	}
	return ""
}

func (x *UserNotification) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *UserNotification) GetTextI18N() string {
	if x != nil {
		return x.TextI18N
	}
	return ""
}

func (x *UserNotification) GetCreatedTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTimestamp
	}
	return nil
}

func (x *UserNotification) GetReadTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.ReadTimestamp
	}
	return nil
}

func (x *UserNotification) GetRemovedTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.RemovedTimestamp
	}
	return nil
}

type PopupNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID               uint64                 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Type             string                 `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	SenderID         uint64                 `protobuf:"varint,3,opt,name=SenderID,proto3" json:"SenderID,omitempty"`
	RecipientID      string                 `protobuf:"bytes,4,opt,name=RecipientID,proto3" json:"RecipientID,omitempty"`
	Title            string                 `protobuf:"bytes,5,opt,name=Title,proto3" json:"Title,omitempty"`
	TitleI18N        string                 `protobuf:"bytes,6,opt,name=TitleI18n,proto3" json:"TitleI18n,omitempty"`
	Text             string                 `protobuf:"bytes,7,opt,name=Text,proto3" json:"Text,omitempty"`
	TextI18N         string                 `protobuf:"bytes,8,opt,name=TextI18n,proto3" json:"TextI18n,omitempty"`
	CreatedTimestamp *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=CreatedTimestamp,proto3" json:"CreatedTimestamp,omitempty"`
}

func (x *PopupNotification) Reset() {
	*x = PopupNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_transport_proto_src_notifications_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PopupNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PopupNotification) ProtoMessage() {}

func (x *PopupNotification) ProtoReflect() protoreflect.Message {
	mi := &file_src_transport_proto_src_notifications_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PopupNotification.ProtoReflect.Descriptor instead.
func (*PopupNotification) Descriptor() ([]byte, []int) {
	return file_src_transport_proto_src_notifications_models_proto_rawDescGZIP(), []int{1}
}

func (x *PopupNotification) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *PopupNotification) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PopupNotification) GetSenderID() uint64 {
	if x != nil {
		return x.SenderID
	}
	return 0
}

func (x *PopupNotification) GetRecipientID() string {
	if x != nil {
		return x.RecipientID
	}
	return ""
}

func (x *PopupNotification) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PopupNotification) GetTitleI18N() string {
	if x != nil {
		return x.TitleI18N
	}
	return ""
}

func (x *PopupNotification) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *PopupNotification) GetTextI18N() string {
	if x != nil {
		return x.TextI18N
	}
	return ""
}

func (x *PopupNotification) GetCreatedTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTimestamp
	}
	return nil
}

type UserNotificationConstructor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        string `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	SenderID    uint64 `protobuf:"varint,2,opt,name=SenderID,proto3" json:"SenderID,omitempty"`
	RecipientID uint64 `protobuf:"varint,3,opt,name=RecipientID,proto3" json:"RecipientID,omitempty"`
	Title       string `protobuf:"bytes,4,opt,name=Title,proto3" json:"Title,omitempty"`
	TitleI18N   string `protobuf:"bytes,5,opt,name=TitleI18n,proto3" json:"TitleI18n,omitempty"`
	Text        string `protobuf:"bytes,6,opt,name=Text,proto3" json:"Text,omitempty"`
	TextI18N    string `protobuf:"bytes,7,opt,name=TextI18n,proto3" json:"TextI18n,omitempty"`
}

func (x *UserNotificationConstructor) Reset() {
	*x = UserNotificationConstructor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_transport_proto_src_notifications_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserNotificationConstructor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserNotificationConstructor) ProtoMessage() {}

func (x *UserNotificationConstructor) ProtoReflect() protoreflect.Message {
	mi := &file_src_transport_proto_src_notifications_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserNotificationConstructor.ProtoReflect.Descriptor instead.
func (*UserNotificationConstructor) Descriptor() ([]byte, []int) {
	return file_src_transport_proto_src_notifications_models_proto_rawDescGZIP(), []int{2}
}

func (x *UserNotificationConstructor) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UserNotificationConstructor) GetSenderID() uint64 {
	if x != nil {
		return x.SenderID
	}
	return 0
}

func (x *UserNotificationConstructor) GetRecipientID() uint64 {
	if x != nil {
		return x.RecipientID
	}
	return 0
}

func (x *UserNotificationConstructor) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UserNotificationConstructor) GetTitleI18N() string {
	if x != nil {
		return x.TitleI18N
	}
	return ""
}

func (x *UserNotificationConstructor) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *UserNotificationConstructor) GetTextI18N() string {
	if x != nil {
		return x.TextI18N
	}
	return ""
}

type PopupNotificationConstructor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        string `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	SenderID    uint64 `protobuf:"varint,2,opt,name=SenderID,proto3" json:"SenderID,omitempty"`
	RecipientID string `protobuf:"bytes,3,opt,name=RecipientID,proto3" json:"RecipientID,omitempty"`
	Title       string `protobuf:"bytes,4,opt,name=Title,proto3" json:"Title,omitempty"`
	TitleI18N   string `protobuf:"bytes,5,opt,name=TitleI18n,proto3" json:"TitleI18n,omitempty"`
	Text        string `protobuf:"bytes,6,opt,name=Text,proto3" json:"Text,omitempty"`
	TextI18N    string `protobuf:"bytes,7,opt,name=TextI18n,proto3" json:"TextI18n,omitempty"`
}

func (x *PopupNotificationConstructor) Reset() {
	*x = PopupNotificationConstructor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_transport_proto_src_notifications_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PopupNotificationConstructor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PopupNotificationConstructor) ProtoMessage() {}

func (x *PopupNotificationConstructor) ProtoReflect() protoreflect.Message {
	mi := &file_src_transport_proto_src_notifications_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PopupNotificationConstructor.ProtoReflect.Descriptor instead.
func (*PopupNotificationConstructor) Descriptor() ([]byte, []int) {
	return file_src_transport_proto_src_notifications_models_proto_rawDescGZIP(), []int{3}
}

func (x *PopupNotificationConstructor) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PopupNotificationConstructor) GetSenderID() uint64 {
	if x != nil {
		return x.SenderID
	}
	return 0
}

func (x *PopupNotificationConstructor) GetRecipientID() string {
	if x != nil {
		return x.RecipientID
	}
	return ""
}

func (x *PopupNotificationConstructor) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PopupNotificationConstructor) GetTitleI18N() string {
	if x != nil {
		return x.TitleI18N
	}
	return ""
}

func (x *PopupNotificationConstructor) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *PopupNotificationConstructor) GetTextI18N() string {
	if x != nil {
		return x.TextI18N
	}
	return ""
}

var File_src_transport_proto_src_notifications_models_proto protoreflect.FileDescriptor

var file_src_transport_proto_src_notifications_models_proto_rawDesc = []byte{
	0x0a, 0x32, 0x73, 0x72, 0x63, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x03, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x49, 0x31, 0x38, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x49, 0x31, 0x38, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65,
	0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x65, 0x78, 0x74, 0x49, 0x31, 0x38, 0x6e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x65, 0x78, 0x74, 0x49, 0x31, 0x38, 0x6e, 0x12, 0x46,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x40, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x52, 0x65, 0x61, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x46, 0x0a, 0x10, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x22, 0xa1, 0x02, 0x0a, 0x11, 0x50, 0x6f, 0x70, 0x75, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x53, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x49, 0x31, 0x38, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x49, 0x31, 0x38, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x65, 0x78, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x54, 0x65, 0x78, 0x74, 0x49, 0x31, 0x38, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x54, 0x65, 0x78, 0x74, 0x49, 0x31, 0x38, 0x6e, 0x12, 0x46, 0x0a, 0x10,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0xd3, 0x01, 0x0a, 0x1b, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x53, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x52, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x49, 0x31, 0x38, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x49, 0x31, 0x38, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65,
	0x78, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x54, 0x65, 0x78, 0x74, 0x49, 0x31, 0x38, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x54, 0x65, 0x78, 0x74, 0x49, 0x31, 0x38, 0x6e, 0x22, 0xd4, 0x01, 0x0a, 0x1c, 0x50,
	0x6f, 0x70, 0x75, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x49, 0x31, 0x38, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x49, 0x31, 0x38,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x65, 0x78, 0x74, 0x49, 0x31, 0x38,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x65, 0x78, 0x74, 0x49, 0x31, 0x38,
	0x6e, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_transport_proto_src_notifications_models_proto_rawDescOnce sync.Once
	file_src_transport_proto_src_notifications_models_proto_rawDescData = file_src_transport_proto_src_notifications_models_proto_rawDesc
)

func file_src_transport_proto_src_notifications_models_proto_rawDescGZIP() []byte {
	file_src_transport_proto_src_notifications_models_proto_rawDescOnce.Do(func() {
		file_src_transport_proto_src_notifications_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_transport_proto_src_notifications_models_proto_rawDescData)
	})
	return file_src_transport_proto_src_notifications_models_proto_rawDescData
}

var file_src_transport_proto_src_notifications_models_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_src_transport_proto_src_notifications_models_proto_goTypes = []interface{}{
	(*UserNotification)(nil),             // 0: grpc_service.UserNotification
	(*PopupNotification)(nil),            // 1: grpc_service.PopupNotification
	(*UserNotificationConstructor)(nil),  // 2: grpc_service.UserNotificationConstructor
	(*PopupNotificationConstructor)(nil), // 3: grpc_service.PopupNotificationConstructor
	(*timestamppb.Timestamp)(nil),        // 4: google.protobuf.Timestamp
}
var file_src_transport_proto_src_notifications_models_proto_depIdxs = []int32{
	4, // 0: grpc_service.UserNotification.CreatedTimestamp:type_name -> google.protobuf.Timestamp
	4, // 1: grpc_service.UserNotification.ReadTimestamp:type_name -> google.protobuf.Timestamp
	4, // 2: grpc_service.UserNotification.RemovedTimestamp:type_name -> google.protobuf.Timestamp
	4, // 3: grpc_service.PopupNotification.CreatedTimestamp:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_src_transport_proto_src_notifications_models_proto_init() }
func file_src_transport_proto_src_notifications_models_proto_init() {
	if File_src_transport_proto_src_notifications_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_transport_proto_src_notifications_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserNotification); i {
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
		file_src_transport_proto_src_notifications_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PopupNotification); i {
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
		file_src_transport_proto_src_notifications_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserNotificationConstructor); i {
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
		file_src_transport_proto_src_notifications_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PopupNotificationConstructor); i {
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
			RawDescriptor: file_src_transport_proto_src_notifications_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_transport_proto_src_notifications_models_proto_goTypes,
		DependencyIndexes: file_src_transport_proto_src_notifications_models_proto_depIdxs,
		MessageInfos:      file_src_transport_proto_src_notifications_models_proto_msgTypes,
	}.Build()
	File_src_transport_proto_src_notifications_models_proto = out.File
	file_src_transport_proto_src_notifications_models_proto_rawDesc = nil
	file_src_transport_proto_src_notifications_models_proto_goTypes = nil
	file_src_transport_proto_src_notifications_models_proto_depIdxs = nil
}

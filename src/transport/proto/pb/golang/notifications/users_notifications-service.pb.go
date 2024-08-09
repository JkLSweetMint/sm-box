// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: src/transport/proto/src/notifications/users_notifications-service.proto

package __

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

type UserNotificationsCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Constructors []*UserNotificationConstructor `protobuf:"bytes,1,rep,name=Constructors,proto3" json:"Constructors,omitempty"`
}

func (x *UserNotificationsCreateRequest) Reset() {
	*x = UserNotificationsCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_transport_proto_src_notifications_users_notifications_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserNotificationsCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserNotificationsCreateRequest) ProtoMessage() {}

func (x *UserNotificationsCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_transport_proto_src_notifications_users_notifications_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserNotificationsCreateRequest.ProtoReflect.Descriptor instead.
func (*UserNotificationsCreateRequest) Descriptor() ([]byte, []int) {
	return file_src_transport_proto_src_notifications_users_notifications_service_proto_rawDescGZIP(), []int{0}
}

func (x *UserNotificationsCreateRequest) GetConstructors() []*UserNotificationConstructor {
	if x != nil {
		return x.Constructors
	}
	return nil
}

type UserNotificationsCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notifications []*UserNotification `protobuf:"bytes,1,rep,name=Notifications,proto3" json:"Notifications,omitempty"`
}

func (x *UserNotificationsCreateResponse) Reset() {
	*x = UserNotificationsCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_transport_proto_src_notifications_users_notifications_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserNotificationsCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserNotificationsCreateResponse) ProtoMessage() {}

func (x *UserNotificationsCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_transport_proto_src_notifications_users_notifications_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserNotificationsCreateResponse.ProtoReflect.Descriptor instead.
func (*UserNotificationsCreateResponse) Descriptor() ([]byte, []int) {
	return file_src_transport_proto_src_notifications_users_notifications_service_proto_rawDescGZIP(), []int{1}
}

func (x *UserNotificationsCreateResponse) GetNotifications() []*UserNotification {
	if x != nil {
		return x.Notifications
	}
	return nil
}

type UserNotificationsCreateOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Constructor *UserNotificationConstructor `protobuf:"bytes,1,opt,name=Constructor,proto3" json:"Constructor,omitempty"`
}

func (x *UserNotificationsCreateOneRequest) Reset() {
	*x = UserNotificationsCreateOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_transport_proto_src_notifications_users_notifications_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserNotificationsCreateOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserNotificationsCreateOneRequest) ProtoMessage() {}

func (x *UserNotificationsCreateOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_transport_proto_src_notifications_users_notifications_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserNotificationsCreateOneRequest.ProtoReflect.Descriptor instead.
func (*UserNotificationsCreateOneRequest) Descriptor() ([]byte, []int) {
	return file_src_transport_proto_src_notifications_users_notifications_service_proto_rawDescGZIP(), []int{2}
}

func (x *UserNotificationsCreateOneRequest) GetConstructor() *UserNotificationConstructor {
	if x != nil {
		return x.Constructor
	}
	return nil
}

type UserNotificationsCreateOneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notification *UserNotification `protobuf:"bytes,1,opt,name=Notification,proto3" json:"Notification,omitempty"`
}

func (x *UserNotificationsCreateOneResponse) Reset() {
	*x = UserNotificationsCreateOneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_transport_proto_src_notifications_users_notifications_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserNotificationsCreateOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserNotificationsCreateOneResponse) ProtoMessage() {}

func (x *UserNotificationsCreateOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_transport_proto_src_notifications_users_notifications_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserNotificationsCreateOneResponse.ProtoReflect.Descriptor instead.
func (*UserNotificationsCreateOneResponse) Descriptor() ([]byte, []int) {
	return file_src_transport_proto_src_notifications_users_notifications_service_proto_rawDescGZIP(), []int{3}
}

func (x *UserNotificationsCreateOneResponse) GetNotification() *UserNotification {
	if x != nil {
		return x.Notification
	}
	return nil
}

var File_src_transport_proto_src_notifications_users_notifications_service_proto protoreflect.FileDescriptor

var file_src_transport_proto_src_notifications_users_notifications_service_proto_rawDesc = []byte{
	0x0a, 0x47, 0x73, 0x72, 0x63, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x32, 0x73, 0x72, 0x63, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x1e, 0x55,
	0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4d, 0x0a,
	0x0c, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0c,
	0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x67, 0x0a, 0x1f,
	0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x44, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x70, 0x0a, 0x21, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x0b, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0b, 0x43, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x68, 0x0a, 0x22, 0x55, 0x73, 0x65, 0x72, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a,
	0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x32, 0xf1, 0x01, 0x0a, 0x18, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6e, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x6e, 0x65, 0x12, 0x2f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_src_transport_proto_src_notifications_users_notifications_service_proto_rawDescOnce sync.Once
	file_src_transport_proto_src_notifications_users_notifications_service_proto_rawDescData = file_src_transport_proto_src_notifications_users_notifications_service_proto_rawDesc
)

func file_src_transport_proto_src_notifications_users_notifications_service_proto_rawDescGZIP() []byte {
	file_src_transport_proto_src_notifications_users_notifications_service_proto_rawDescOnce.Do(func() {
		file_src_transport_proto_src_notifications_users_notifications_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_transport_proto_src_notifications_users_notifications_service_proto_rawDescData)
	})
	return file_src_transport_proto_src_notifications_users_notifications_service_proto_rawDescData
}

var file_src_transport_proto_src_notifications_users_notifications_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_src_transport_proto_src_notifications_users_notifications_service_proto_goTypes = []interface{}{
	(*UserNotificationsCreateRequest)(nil),     // 0: grpc_service.UserNotificationsCreateRequest
	(*UserNotificationsCreateResponse)(nil),    // 1: grpc_service.UserNotificationsCreateResponse
	(*UserNotificationsCreateOneRequest)(nil),  // 2: grpc_service.UserNotificationsCreateOneRequest
	(*UserNotificationsCreateOneResponse)(nil), // 3: grpc_service.UserNotificationsCreateOneResponse
	(*UserNotificationConstructor)(nil),        // 4: grpc_service.UserNotificationConstructor
	(*UserNotification)(nil),                   // 5: grpc_service.UserNotification
}
var file_src_transport_proto_src_notifications_users_notifications_service_proto_depIdxs = []int32{
	4, // 0: grpc_service.UserNotificationsCreateRequest.Constructors:type_name -> grpc_service.UserNotificationConstructor
	5, // 1: grpc_service.UserNotificationsCreateResponse.Notifications:type_name -> grpc_service.UserNotification
	4, // 2: grpc_service.UserNotificationsCreateOneRequest.Constructor:type_name -> grpc_service.UserNotificationConstructor
	5, // 3: grpc_service.UserNotificationsCreateOneResponse.Notification:type_name -> grpc_service.UserNotification
	0, // 4: grpc_service.UserNotificationsService.Create:input_type -> grpc_service.UserNotificationsCreateRequest
	2, // 5: grpc_service.UserNotificationsService.CreateOne:input_type -> grpc_service.UserNotificationsCreateOneRequest
	1, // 6: grpc_service.UserNotificationsService.Create:output_type -> grpc_service.UserNotificationsCreateResponse
	3, // 7: grpc_service.UserNotificationsService.CreateOne:output_type -> grpc_service.UserNotificationsCreateOneResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_src_transport_proto_src_notifications_users_notifications_service_proto_init() }
func file_src_transport_proto_src_notifications_users_notifications_service_proto_init() {
	if File_src_transport_proto_src_notifications_users_notifications_service_proto != nil {
		return
	}
	file_src_transport_proto_src_notifications_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_src_transport_proto_src_notifications_users_notifications_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserNotificationsCreateRequest); i {
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
		file_src_transport_proto_src_notifications_users_notifications_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserNotificationsCreateResponse); i {
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
		file_src_transport_proto_src_notifications_users_notifications_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserNotificationsCreateOneRequest); i {
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
		file_src_transport_proto_src_notifications_users_notifications_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserNotificationsCreateOneResponse); i {
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
			RawDescriptor: file_src_transport_proto_src_notifications_users_notifications_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_transport_proto_src_notifications_users_notifications_service_proto_goTypes,
		DependencyIndexes: file_src_transport_proto_src_notifications_users_notifications_service_proto_depIdxs,
		MessageInfos:      file_src_transport_proto_src_notifications_users_notifications_service_proto_msgTypes,
	}.Build()
	File_src_transport_proto_src_notifications_users_notifications_service_proto = out.File
	file_src_transport_proto_src_notifications_users_notifications_service_proto_rawDesc = nil
	file_src_transport_proto_src_notifications_users_notifications_service_proto_goTypes = nil
	file_src_transport_proto_src_notifications_users_notifications_service_proto_depIdxs = nil
}

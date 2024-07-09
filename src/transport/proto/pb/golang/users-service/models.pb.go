// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: src/transport/proto/src/users-service/models.proto

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        uint64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ProjectID uint64  `protobuf:"varint,2,opt,name=ProjectID,proto3" json:"ProjectID,omitempty"`
	Email     string  `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	Username  string  `protobuf:"bytes,4,opt,name=Username,proto3" json:"Username,omitempty"`
	Accesses  []*Role `protobuf:"bytes,6,rep,name=Accesses,proto3" json:"Accesses,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_transport_proto_src_users_service_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_src_transport_proto_src_users_service_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_src_transport_proto_src_users_service_models_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *User) GetProjectID() uint64 {
	if x != nil {
		return x.ProjectID
	}
	return 0
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetAccesses() []*Role {
	if x != nil {
		return x.Accesses
	}
	return nil
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           uint64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ProjectID    uint64  `protobuf:"varint,2,opt,name=ProjectID,proto3" json:"ProjectID,omitempty"`
	Name         string  `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	IsSystem     bool    `protobuf:"varint,4,opt,name=IsSystem,proto3" json:"IsSystem,omitempty"`
	Inheritances []*Role `protobuf:"bytes,5,rep,name=Inheritances,proto3" json:"Inheritances,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_transport_proto_src_users_service_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_src_transport_proto_src_users_service_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_src_transport_proto_src_users_service_models_proto_rawDescGZIP(), []int{1}
}

func (x *Role) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Role) GetProjectID() uint64 {
	if x != nil {
		return x.ProjectID
	}
	return 0
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Role) GetIsSystem() bool {
	if x != nil {
		return x.IsSystem
	}
	return false
}

func (x *Role) GetInheritances() []*Role {
	if x != nil {
		return x.Inheritances
	}
	return nil
}

var File_src_transport_proto_src_users_service_models_proto protoreflect.FileDescriptor

var file_src_transport_proto_src_users_service_models_proto_rawDesc = []byte{
	0x0a, 0x32, 0x73, 0x72, 0x63, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x22, 0x96, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x08, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x9c, 0x01, 0x0a, 0x04,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x73, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x49, 0x73, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x12, 0x36, 0x0a, 0x0c, 0x49, 0x6e, 0x68, 0x65, 0x72, 0x69, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x0c, 0x49, 0x6e,
	0x68, 0x65, 0x72, 0x69, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_transport_proto_src_users_service_models_proto_rawDescOnce sync.Once
	file_src_transport_proto_src_users_service_models_proto_rawDescData = file_src_transport_proto_src_users_service_models_proto_rawDesc
)

func file_src_transport_proto_src_users_service_models_proto_rawDescGZIP() []byte {
	file_src_transport_proto_src_users_service_models_proto_rawDescOnce.Do(func() {
		file_src_transport_proto_src_users_service_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_transport_proto_src_users_service_models_proto_rawDescData)
	})
	return file_src_transport_proto_src_users_service_models_proto_rawDescData
}

var file_src_transport_proto_src_users_service_models_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_src_transport_proto_src_users_service_models_proto_goTypes = []interface{}{
	(*User)(nil), // 0: grpc_service.User
	(*Role)(nil), // 1: grpc_service.Role
}
var file_src_transport_proto_src_users_service_models_proto_depIdxs = []int32{
	1, // 0: grpc_service.User.Accesses:type_name -> grpc_service.Role
	1, // 1: grpc_service.Role.Inheritances:type_name -> grpc_service.Role
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_src_transport_proto_src_users_service_models_proto_init() }
func file_src_transport_proto_src_users_service_models_proto_init() {
	if File_src_transport_proto_src_users_service_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_transport_proto_src_users_service_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_src_transport_proto_src_users_service_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
			RawDescriptor: file_src_transport_proto_src_users_service_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_transport_proto_src_users_service_models_proto_goTypes,
		DependencyIndexes: file_src_transport_proto_src_users_service_models_proto_depIdxs,
		MessageInfos:      file_src_transport_proto_src_users_service_models_proto_msgTypes,
	}.Build()
	File_src_transport_proto_src_users_service_models_proto = out.File
	file_src_transport_proto_src_users_service_models_proto_rawDesc = nil
	file_src_transport_proto_src_users_service_models_proto_goTypes = nil
	file_src_transport_proto_src_users_service_models_proto_depIdxs = nil
}
//*
// Сервис CRUD операций с клиентами

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.5
// source: api/grpc/protobuf/clients_v1/clients.proto

package clients_v1

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

// *
// Запрос на регистрацию
type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// * Логин
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// * Пароль
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// * Имя
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// * Фамилия
	Surname string `protobuf:"bytes,4,opt,name=surname,proto3" json:"surname,omitempty"`
	// * Отчество
	Patronymic string `protobuf:"bytes,5,opt,name=patronymic,proto3" json:"patronymic,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protobuf_clients_v1_clients_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protobuf_clients_v1_clients_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_protobuf_clients_v1_clients_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterRequest) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *RegisterRequest) GetPatronymic() string {
	if x != nil {
		return x.Patronymic
	}
	return ""
}

// *
// Ответ на регистрацию
type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protobuf_clients_v1_clients_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protobuf_clients_v1_clients_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_protobuf_clients_v1_clients_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_api_grpc_protobuf_clients_v1_clients_proto protoreflect.FileDescriptor

var file_api_grpc_protobuf_clients_v1_clients_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x76, 0x31, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x97, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61,
	0x74, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x61, 0x74, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x69, 0x63, 0x22, 0x22, 0x0a, 0x10, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0x69,
	0x0a, 0x0e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x57, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x6f, 0x6b, 0x74, 0x6f, 0x72, 0x47, 0x68,
	0x6f, 0x73, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2d, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x72, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_grpc_protobuf_clients_v1_clients_proto_rawDescOnce sync.Once
	file_api_grpc_protobuf_clients_v1_clients_proto_rawDescData = file_api_grpc_protobuf_clients_v1_clients_proto_rawDesc
)

func file_api_grpc_protobuf_clients_v1_clients_proto_rawDescGZIP() []byte {
	file_api_grpc_protobuf_clients_v1_clients_proto_rawDescOnce.Do(func() {
		file_api_grpc_protobuf_clients_v1_clients_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_protobuf_clients_v1_clients_proto_rawDescData)
	})
	return file_api_grpc_protobuf_clients_v1_clients_proto_rawDescData
}

var file_api_grpc_protobuf_clients_v1_clients_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_grpc_protobuf_clients_v1_clients_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),  // 0: clients_v1.service.RegisterRequest
	(*RegisterResponse)(nil), // 1: clients_v1.service.RegisterResponse
}
var file_api_grpc_protobuf_clients_v1_clients_proto_depIdxs = []int32{
	0, // 0: clients_v1.service.ClientsService.Register:input_type -> clients_v1.service.RegisterRequest
	1, // 1: clients_v1.service.ClientsService.Register:output_type -> clients_v1.service.RegisterResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_grpc_protobuf_clients_v1_clients_proto_init() }
func file_api_grpc_protobuf_clients_v1_clients_proto_init() {
	if File_api_grpc_protobuf_clients_v1_clients_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_protobuf_clients_v1_clients_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_api_grpc_protobuf_clients_v1_clients_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
			RawDescriptor: file_api_grpc_protobuf_clients_v1_clients_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_protobuf_clients_v1_clients_proto_goTypes,
		DependencyIndexes: file_api_grpc_protobuf_clients_v1_clients_proto_depIdxs,
		MessageInfos:      file_api_grpc_protobuf_clients_v1_clients_proto_msgTypes,
	}.Build()
	File_api_grpc_protobuf_clients_v1_clients_proto = out.File
	file_api_grpc_protobuf_clients_v1_clients_proto_rawDesc = nil
	file_api_grpc_protobuf_clients_v1_clients_proto_goTypes = nil
	file_api_grpc_protobuf_clients_v1_clients_proto_depIdxs = nil
}

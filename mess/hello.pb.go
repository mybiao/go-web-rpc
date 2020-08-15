// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: mess/hello.proto

package mess

import (
	context "context"
	reflect "reflect"
	"strings"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Str string `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mess_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_mess_hello_proto_msgTypes[0]
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
	return file_mess_hello_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

var File_mess_hello_proto protoreflect.FileDescriptor

var file_mess_hello_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x6d, 0x65, 0x73, 0x73, 0x22, 0x18, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73,
	0x74, 0x72, 0x32, 0x28, 0x0a, 0x05, 0x55, 0x70, 0x70, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x05, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0a, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x1a, 0x0a, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x06, 0x5a, 0x04,
	0x6d, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mess_hello_proto_rawDescOnce sync.Once
	file_mess_hello_proto_rawDescData = file_mess_hello_proto_rawDesc
)

func file_mess_hello_proto_rawDescGZIP() []byte {
	file_mess_hello_proto_rawDescOnce.Do(func() {
		file_mess_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_mess_hello_proto_rawDescData)
	})
	return file_mess_hello_proto_rawDescData
}

var file_mess_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mess_hello_proto_goTypes = []interface{}{
	(*User)(nil), // 0: mess.User
}
var file_mess_hello_proto_depIdxs = []int32{
	0, // 0: mess.Upper.Hello:input_type -> mess.User
	0, // 1: mess.Upper.Hello:output_type -> mess.User
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mess_hello_proto_init() }
func file_mess_hello_proto_init() {
	if File_mess_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mess_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mess_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mess_hello_proto_goTypes,
		DependencyIndexes: file_mess_hello_proto_depIdxs,
		MessageInfos:      file_mess_hello_proto_msgTypes,
	}.Build()
	File_mess_hello_proto = out.File
	file_mess_hello_proto_rawDesc = nil
	file_mess_hello_proto_goTypes = nil
	file_mess_hello_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UpperClient is the client API for Upper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UpperClient interface {
	Hello(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
}

type upperClient struct {
	cc grpc.ClientConnInterface
}

func NewUpperClient(cc grpc.ClientConnInterface) UpperClient {
	return &upperClient{cc}
}

func (c *upperClient) Hello(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/mess.Upper/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpperServer is the server API for Upper service.
type UpperServer interface {
	Hello(context.Context, *User) (*User, error)
}

// UnimplementedUpperServer can be embedded to have forward compatible implementations.
type ImplementedUpperServer struct {
}

func (*ImplementedUpperServer) Hello(c context.Context, u *User) (*User, error) {
	return &User{
		Str: strings.ToUpper(u.Str),
	}, nil
}

func RegisterUpperServer(s *grpc.Server, srv UpperServer) {
	s.RegisterService(&_Upper_serviceDesc, srv)
}

func _Upper_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpperServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mess.Upper/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpperServer).Hello(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _Upper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mess.Upper",
	HandlerType: (*UpperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Upper_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mess/hello.proto",
}

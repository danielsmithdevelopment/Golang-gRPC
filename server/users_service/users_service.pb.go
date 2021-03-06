// Code generated by protoc-gen-go. DO NOT EDIT.
// source: users_service.proto

package users_service

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	FirstName            string   `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_service_2a7ce690d201a496, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type AddUserRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddUserRequest) Reset()         { *m = AddUserRequest{} }
func (m *AddUserRequest) String() string { return proto.CompactTextString(m) }
func (*AddUserRequest) ProtoMessage()    {}
func (*AddUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_service_2a7ce690d201a496, []int{1}
}
func (m *AddUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUserRequest.Unmarshal(m, b)
}
func (m *AddUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUserRequest.Marshal(b, m, deterministic)
}
func (dst *AddUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUserRequest.Merge(dst, src)
}
func (m *AddUserRequest) XXX_Size() int {
	return xxx_messageInfo_AddUserRequest.Size(m)
}
func (m *AddUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddUserRequest proto.InternalMessageInfo

func (m *AddUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AddUserRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *AddUserRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type GetUsersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUsersRequest) Reset()         { *m = GetUsersRequest{} }
func (m *GetUsersRequest) String() string { return proto.CompactTextString(m) }
func (*GetUsersRequest) ProtoMessage()    {}
func (*GetUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_service_2a7ce690d201a496, []int{2}
}
func (m *GetUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUsersRequest.Unmarshal(m, b)
}
func (m *GetUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUsersRequest.Marshal(b, m, deterministic)
}
func (dst *GetUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUsersRequest.Merge(dst, src)
}
func (m *GetUsersRequest) XXX_Size() int {
	return xxx_messageInfo_GetUsersRequest.Size(m)
}
func (m *GetUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUsersRequest proto.InternalMessageInfo

type GetUsersResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUsersResponse) Reset()         { *m = GetUsersResponse{} }
func (m *GetUsersResponse) String() string { return proto.CompactTextString(m) }
func (*GetUsersResponse) ProtoMessage()    {}
func (*GetUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_service_2a7ce690d201a496, []int{3}
}
func (m *GetUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUsersResponse.Unmarshal(m, b)
}
func (m *GetUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUsersResponse.Marshal(b, m, deterministic)
}
func (dst *GetUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUsersResponse.Merge(dst, src)
}
func (m *GetUsersResponse) XXX_Size() int {
	return xxx_messageInfo_GetUsersResponse.Size(m)
}
func (m *GetUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUsersResponse proto.InternalMessageInfo

func (m *GetUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type DeleteUserRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_service_2a7ce690d201a496, []int{4}
}
func (m *DeleteUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserRequest.Unmarshal(m, b)
}
func (m *DeleteUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserRequest.Merge(dst, src)
}
func (m *DeleteUserRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUserRequest.Size(m)
}
func (m *DeleteUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserRequest proto.InternalMessageInfo

func (m *DeleteUserRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "users.User")
	proto.RegisterType((*AddUserRequest)(nil), "users.AddUserRequest")
	proto.RegisterType((*GetUsersRequest)(nil), "users.GetUsersRequest")
	proto.RegisterType((*GetUsersResponse)(nil), "users.GetUsersResponse")
	proto.RegisterType((*DeleteUserRequest)(nil), "users.DeleteUserRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*User, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*User, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/users.UserService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/users.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	AddUser(context.Context, *AddUserRequest) (*User, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*User, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "users.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _UserService_AddUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _UserService_GetUsers_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users_service.proto",
}

func init() { proto.RegisterFile("users_service.proto", fileDescriptor_users_service_2a7ce690d201a496) }

var fileDescriptor_users_service_2a7ce690d201a496 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x3b, 0x49, 0xfb, 0x7d, 0xed, 0x0d, 0x54, 0x7b, 0xfd, 0x37, 0x04, 0x17, 0x71, 0xdc,
	0x64, 0x15, 0xa1, 0xd2, 0xa5, 0x0b, 0x41, 0x70, 0xe7, 0x22, 0xe2, 0x5a, 0xa3, 0xb9, 0x85, 0x81,
	0xd4, 0xd4, 0xb9, 0x53, 0x9f, 0xcc, 0x07, 0x94, 0x4c, 0x46, 0x53, 0x53, 0x74, 0x79, 0xcf, 0x09,
	0xf9, 0x9d, 0x73, 0x06, 0x0e, 0x36, 0x4c, 0x86, 0x1f, 0x99, 0xcc, 0xbb, 0x7e, 0xa1, 0x6c, 0x6d,
	0x6a, 0x5b, 0xe3, 0xc8, 0x89, 0x6a, 0x09, 0xc3, 0x07, 0x26, 0x83, 0x53, 0x08, 0x74, 0x29, 0x45,
	0x22, 0xd2, 0x30, 0x0f, 0x74, 0x89, 0x87, 0x30, 0xa2, 0x55, 0xa1, 0x2b, 0x19, 0x24, 0x22, 0x9d,
	0xe4, 0xed, 0x81, 0xa7, 0x30, 0x59, 0x6a, 0xc3, 0xf6, 0xae, 0x58, 0x91, 0x0c, 0x9d, 0xd3, 0x09,
	0x18, 0xc3, 0xb8, 0x2a, 0xbc, 0x39, 0x74, 0xe6, 0xf7, 0xad, 0x9e, 0x60, 0x7a, 0x5d, 0x96, 0x0d,
	0x2a, 0xa7, 0xb7, 0x0d, 0xb1, 0xed, 0x08, 0xe2, 0x57, 0x42, 0xf0, 0x17, 0x21, 0xec, 0x11, 0x66,
	0xb0, 0x77, 0x4b, 0xb6, 0x21, 0xb0, 0x47, 0xa8, 0x05, 0xec, 0x77, 0x12, 0xaf, 0xeb, 0x57, 0x26,
	0x3c, 0x83, 0xb6, 0xb9, 0x14, 0x49, 0x98, 0x46, 0xf3, 0x28, 0x73, 0x57, 0xe6, 0x92, 0xf9, 0x4d,
	0xce, 0x61, 0x76, 0x43, 0x15, 0x59, 0xda, 0x8e, 0xdb, 0x1b, 0x68, 0xfe, 0x21, 0x20, 0x6a, 0xfc,
	0xfb, 0x76, 0x55, 0xbc, 0x80, 0xff, 0xbe, 0x20, 0x1e, 0xf9, 0x7f, 0xfe, 0x2c, 0x1c, 0x6f, 0xa3,
	0xd4, 0x00, 0xaf, 0x60, 0xfc, 0x15, 0x0e, 0x8f, 0xbd, 0xd5, 0x2b, 0x10, 0x9f, 0xec, 0xe8, 0x6d,
	0x0b, 0x35, 0xc0, 0x05, 0x40, 0x17, 0x12, 0xa5, 0xff, 0x70, 0x27, 0x77, 0x8f, 0xfa, 0xfc, 0xcf,
	0xbd, 0xfe, 0xe5, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x46, 0x27, 0x1b, 0xb2, 0x14, 0x02, 0x00,
	0x00,
}

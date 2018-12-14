// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/user.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/tcncloud/protoc-gen-persist/persist"

import (
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
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Friends              *Friends             `protobuf:"bytes,3,opt,name=friends,proto3" json:"friends,omitempty"`
	CreatedOn            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_on,json=createdOn,proto3" json:"created_on,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_f3305310eebd2b9a, []int{0}
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

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetFriends() *Friends {
	if m != nil {
		return m.Friends
	}
	return nil
}

func (m *User) GetCreatedOn() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedOn
	}
	return nil
}

type InsertUserReq struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Friends              *Friends             `protobuf:"bytes,3,opt,name=friends,proto3" json:"friends,omitempty"`
	CreatedOn            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_on,json=createdOn,proto3" json:"created_on,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *InsertUserReq) Reset()         { *m = InsertUserReq{} }
func (m *InsertUserReq) String() string { return proto.CompactTextString(m) }
func (*InsertUserReq) ProtoMessage()    {}
func (*InsertUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_f3305310eebd2b9a, []int{1}
}
func (m *InsertUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertUserReq.Unmarshal(m, b)
}
func (m *InsertUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertUserReq.Marshal(b, m, deterministic)
}
func (dst *InsertUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertUserReq.Merge(dst, src)
}
func (m *InsertUserReq) XXX_Size() int {
	return xxx_messageInfo_InsertUserReq.Size(m)
}
func (m *InsertUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_InsertUserReq proto.InternalMessageInfo

func (m *InsertUserReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *InsertUserReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InsertUserReq) GetFriends() *Friends {
	if m != nil {
		return m.Friends
	}
	return nil
}

func (m *InsertUserReq) GetCreatedOn() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedOn
	}
	return nil
}

type Friends struct {
	Names                []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Friends) Reset()         { *m = Friends{} }
func (m *Friends) String() string { return proto.CompactTextString(m) }
func (*Friends) ProtoMessage()    {}
func (*Friends) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_f3305310eebd2b9a, []int{2}
}
func (m *Friends) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Friends.Unmarshal(m, b)
}
func (m *Friends) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Friends.Marshal(b, m, deterministic)
}
func (dst *Friends) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Friends.Merge(dst, src)
}
func (m *Friends) XXX_Size() int {
	return xxx_messageInfo_Friends.Size(m)
}
func (m *Friends) XXX_DiscardUnknown() {
	xxx_messageInfo_Friends.DiscardUnknown(m)
}

var xxx_messageInfo_Friends proto.InternalMessageInfo

func (m *Friends) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

type SliceStringParam struct {
	Slice                []string `protobuf:"bytes,1,rep,name=slice,proto3" json:"slice,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SliceStringParam) Reset()         { *m = SliceStringParam{} }
func (m *SliceStringParam) String() string { return proto.CompactTextString(m) }
func (*SliceStringParam) ProtoMessage()    {}
func (*SliceStringParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_f3305310eebd2b9a, []int{3}
}
func (m *SliceStringParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SliceStringParam.Unmarshal(m, b)
}
func (m *SliceStringParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SliceStringParam.Marshal(b, m, deterministic)
}
func (dst *SliceStringParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SliceStringParam.Merge(dst, src)
}
func (m *SliceStringParam) XXX_Size() int {
	return xxx_messageInfo_SliceStringParam.Size(m)
}
func (m *SliceStringParam) XXX_DiscardUnknown() {
	xxx_messageInfo_SliceStringParam.DiscardUnknown(m)
}

var xxx_messageInfo_SliceStringParam proto.InternalMessageInfo

func (m *SliceStringParam) GetSlice() []string {
	if m != nil {
		return m.Slice
	}
	return nil
}

type FriendsReq struct {
	Names                *SliceStringParam `protobuf:"bytes,2,opt,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FriendsReq) Reset()         { *m = FriendsReq{} }
func (m *FriendsReq) String() string { return proto.CompactTextString(m) }
func (*FriendsReq) ProtoMessage()    {}
func (*FriendsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_f3305310eebd2b9a, []int{4}
}
func (m *FriendsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FriendsReq.Unmarshal(m, b)
}
func (m *FriendsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FriendsReq.Marshal(b, m, deterministic)
}
func (dst *FriendsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FriendsReq.Merge(dst, src)
}
func (m *FriendsReq) XXX_Size() int {
	return xxx_messageInfo_FriendsReq.Size(m)
}
func (m *FriendsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FriendsReq.DiscardUnknown(m)
}

var xxx_messageInfo_FriendsReq proto.InternalMessageInfo

func (m *FriendsReq) GetNames() *SliceStringParam {
	if m != nil {
		return m.Names
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_f3305310eebd2b9a, []int{5}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*InsertUserReq)(nil), "pb.InsertUserReq")
	proto.RegisterType((*Friends)(nil), "pb.Friends")
	proto.RegisterType((*SliceStringParam)(nil), "pb.SliceStringParam")
	proto.RegisterType((*FriendsReq)(nil), "pb.FriendsReq")
	proto.RegisterType((*Empty)(nil), "pb.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UServClient is the client API for UServ service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UServClient interface {
	CreateTable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	InsertUsers(ctx context.Context, opts ...grpc.CallOption) (UServ_InsertUsersClient, error)
	GetAllUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (UServ_GetAllUsersClient, error)
	SelectUserById(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	UpdateUserNames(ctx context.Context, opts ...grpc.CallOption) (UServ_UpdateUserNamesClient, error)
	UpdateNameToFoo(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error)
	UpdateAllNames(ctx context.Context, in *Empty, opts ...grpc.CallOption) (UServ_UpdateAllNamesClient, error)
	GetFriends(ctx context.Context, in *FriendsReq, opts ...grpc.CallOption) (UServ_GetFriendsClient, error)
	DropTable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type uServClient struct {
	cc *grpc.ClientConn
}

func NewUServClient(cc *grpc.ClientConn) UServClient {
	return &uServClient{cc}
}

func (c *uServClient) CreateTable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.UServ/CreateTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uServClient) InsertUsers(ctx context.Context, opts ...grpc.CallOption) (UServ_InsertUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UServ_serviceDesc.Streams[0], "/pb.UServ/InsertUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &uServInsertUsersClient{stream}
	return x, nil
}

type UServ_InsertUsersClient interface {
	Send(*User) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type uServInsertUsersClient struct {
	grpc.ClientStream
}

func (x *uServInsertUsersClient) Send(m *User) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uServInsertUsersClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *uServClient) GetAllUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (UServ_GetAllUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UServ_serviceDesc.Streams[1], "/pb.UServ/GetAllUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &uServGetAllUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UServ_GetAllUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type uServGetAllUsersClient struct {
	grpc.ClientStream
}

func (x *uServGetAllUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *uServClient) SelectUserById(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/pb.UServ/SelectUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uServClient) UpdateUserNames(ctx context.Context, opts ...grpc.CallOption) (UServ_UpdateUserNamesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UServ_serviceDesc.Streams[2], "/pb.UServ/UpdateUserNames", opts...)
	if err != nil {
		return nil, err
	}
	x := &uServUpdateUserNamesClient{stream}
	return x, nil
}

type UServ_UpdateUserNamesClient interface {
	Send(*User) error
	Recv() (*User, error)
	grpc.ClientStream
}

type uServUpdateUserNamesClient struct {
	grpc.ClientStream
}

func (x *uServUpdateUserNamesClient) Send(m *User) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uServUpdateUserNamesClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *uServClient) UpdateNameToFoo(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.UServ/UpdateNameToFoo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uServClient) UpdateAllNames(ctx context.Context, in *Empty, opts ...grpc.CallOption) (UServ_UpdateAllNamesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UServ_serviceDesc.Streams[3], "/pb.UServ/UpdateAllNames", opts...)
	if err != nil {
		return nil, err
	}
	x := &uServUpdateAllNamesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UServ_UpdateAllNamesClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type uServUpdateAllNamesClient struct {
	grpc.ClientStream
}

func (x *uServUpdateAllNamesClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *uServClient) GetFriends(ctx context.Context, in *FriendsReq, opts ...grpc.CallOption) (UServ_GetFriendsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UServ_serviceDesc.Streams[4], "/pb.UServ/GetFriends", opts...)
	if err != nil {
		return nil, err
	}
	x := &uServGetFriendsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UServ_GetFriendsClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type uServGetFriendsClient struct {
	grpc.ClientStream
}

func (x *uServGetFriendsClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *uServClient) DropTable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.UServ/DropTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UServServer is the server API for UServ service.
type UServServer interface {
	CreateTable(context.Context, *Empty) (*Empty, error)
	InsertUsers(UServ_InsertUsersServer) error
	GetAllUsers(*Empty, UServ_GetAllUsersServer) error
	SelectUserById(context.Context, *User) (*User, error)
	UpdateUserNames(UServ_UpdateUserNamesServer) error
	UpdateNameToFoo(context.Context, *User) (*Empty, error)
	UpdateAllNames(*Empty, UServ_UpdateAllNamesServer) error
	GetFriends(*FriendsReq, UServ_GetFriendsServer) error
	DropTable(context.Context, *Empty) (*Empty, error)
}

func RegisterUServServer(s *grpc.Server, srv UServServer) {
	s.RegisterService(&_UServ_serviceDesc, srv)
}

func _UServ_CreateTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UServServer).CreateTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UServ/CreateTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UServServer).CreateTable(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UServ_InsertUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UServServer).InsertUsers(&uServInsertUsersServer{stream})
}

type UServ_InsertUsersServer interface {
	SendAndClose(*Empty) error
	Recv() (*User, error)
	grpc.ServerStream
}

type uServInsertUsersServer struct {
	grpc.ServerStream
}

func (x *uServInsertUsersServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uServInsertUsersServer) Recv() (*User, error) {
	m := new(User)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UServ_GetAllUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UServServer).GetAllUsers(m, &uServGetAllUsersServer{stream})
}

type UServ_GetAllUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}

type uServGetAllUsersServer struct {
	grpc.ServerStream
}

func (x *uServGetAllUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func _UServ_SelectUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UServServer).SelectUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UServ/SelectUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UServServer).SelectUserById(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UServ_UpdateUserNames_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UServServer).UpdateUserNames(&uServUpdateUserNamesServer{stream})
}

type UServ_UpdateUserNamesServer interface {
	Send(*User) error
	Recv() (*User, error)
	grpc.ServerStream
}

type uServUpdateUserNamesServer struct {
	grpc.ServerStream
}

func (x *uServUpdateUserNamesServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uServUpdateUserNamesServer) Recv() (*User, error) {
	m := new(User)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UServ_UpdateNameToFoo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UServServer).UpdateNameToFoo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UServ/UpdateNameToFoo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UServServer).UpdateNameToFoo(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UServ_UpdateAllNames_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UServServer).UpdateAllNames(m, &uServUpdateAllNamesServer{stream})
}

type UServ_UpdateAllNamesServer interface {
	Send(*User) error
	grpc.ServerStream
}

type uServUpdateAllNamesServer struct {
	grpc.ServerStream
}

func (x *uServUpdateAllNamesServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func _UServ_GetFriends_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FriendsReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UServServer).GetFriends(m, &uServGetFriendsServer{stream})
}

type UServ_GetFriendsServer interface {
	Send(*User) error
	grpc.ServerStream
}

type uServGetFriendsServer struct {
	grpc.ServerStream
}

func (x *uServGetFriendsServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func _UServ_DropTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UServServer).DropTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UServ/DropTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UServServer).DropTable(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _UServ_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UServ",
	HandlerType: (*UServServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTable",
			Handler:    _UServ_CreateTable_Handler,
		},
		{
			MethodName: "SelectUserById",
			Handler:    _UServ_SelectUserById_Handler,
		},
		{
			MethodName: "UpdateNameToFoo",
			Handler:    _UServ_UpdateNameToFoo_Handler,
		},
		{
			MethodName: "DropTable",
			Handler:    _UServ_DropTable_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "InsertUsers",
			Handler:       _UServ_InsertUsers_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetAllUsers",
			Handler:       _UServ_GetAllUsers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UpdateUserNames",
			Handler:       _UServ_UpdateUserNames_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "UpdateAllNames",
			Handler:       _UServ_UpdateAllNames_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetFriends",
			Handler:       _UServ_GetFriends_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pb/user.proto",
}

func init() { proto.RegisterFile("pb/user.proto", fileDescriptor_user_f3305310eebd2b9a) }

var fileDescriptor_user_f3305310eebd2b9a = []byte{
	// 850 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xde, 0x71, 0xff, 0xc8, 0x31, 0x2d, 0x61, 0x28, 0xda, 0x90, 0x0b, 0x76, 0x64, 0x81, 0xea,
	0xae, 0x42, 0x52, 0x2d, 0x42, 0x02, 0x44, 0x57, 0x71, 0x8b, 0xdb, 0x46, 0xdb, 0x4d, 0xab, 0x89,
	0xb3, 0x50, 0x84, 0x14, 0x25, 0xf5, 0x34, 0x32, 0x72, 0x3c, 0xae, 0x3d, 0x45, 0xea, 0x25, 0x88,
	0x2b, 0xc4, 0x05, 0x15, 0xe2, 0x05, 0x78, 0x02, 0x2e, 0xfb, 0x2a, 0xbc, 0x00, 0xcf, 0x81, 0x66,
	0xc6, 0x8e, 0xdd, 0xd6, 0x44, 0x82, 0xab, 0xbd, 0x1b, 0xdb, 0xdf, 0xf7, 0x9d, 0xef, 0x9c, 0xf9,
	0x4e, 0x02, 0xeb, 0xf1, 0xa4, 0x73, 0x95, 0xb2, 0xa4, 0x1d, 0x27, 0x5c, 0x70, 0x6c, 0xc4, 0x93,
	0xe6, 0xbb, 0x31, 0x4b, 0xd2, 0x20, 0x15, 0x1d, 0x1e, 0x8b, 0x80, 0x47, 0xa9, 0xfe, 0xd4, 0x7c,
	0x32, 0xe5, 0x7c, 0x1a, 0xb2, 0x8e, 0x7a, 0x9a, 0x5c, 0x5d, 0x74, 0x44, 0x30, 0x63, 0xa9, 0x18,
	0xcf, 0x62, 0x0d, 0xb0, 0x7e, 0x41, 0xb0, 0x3c, 0x4c, 0x59, 0x82, 0x37, 0xc0, 0x08, 0xfc, 0x06,
	0x22, 0xc8, 0x5e, 0xa2, 0x46, 0xe0, 0x63, 0x0c, 0xcb, 0xd1, 0x78, 0xc6, 0x1a, 0x06, 0x41, 0x76,
	0x8d, 0xaa, 0x33, 0xfe, 0x10, 0xd6, 0x2e, 0x92, 0x80, 0x45, 0x7e, 0xda, 0x58, 0x22, 0xc8, 0x36,
	0x9f, 0x99, 0xed, 0x78, 0xd2, 0x3e, 0xd0, 0xaf, 0x68, 0xfe, 0x0d, 0x7f, 0x06, 0x70, 0x9e, 0xb0,
	0xb1, 0x60, 0xfe, 0x88, 0x47, 0x8d, 0x65, 0x85, 0x6c, 0xb6, 0xb5, 0x93, 0x76, 0xee, 0xa4, 0xed,
	0xe5, 0x4e, 0x68, 0x2d, 0x43, 0x9f, 0x44, 0xd6, 0xef, 0x08, 0xd6, 0x7b, 0x51, 0xca, 0x12, 0x21,
	0x4d, 0x51, 0x76, 0xf9, 0x9a, 0xf8, 0x7a, 0x02, 0x6b, 0x99, 0x1c, 0xde, 0x84, 0x15, 0x59, 0x34,
	0x6d, 0x20, 0xb2, 0x64, 0xd7, 0xa8, 0x7e, 0xb0, 0x6c, 0xa8, 0x0f, 0xc2, 0xe0, 0x9c, 0x0d, 0x44,
	0x12, 0x44, 0xd3, 0xd3, 0x71, 0x32, 0x9e, 0x49, 0x64, 0x2a, 0xdf, 0xe5, 0x48, 0xf5, 0x60, 0x7d,
	0x0a, 0x90, 0x3b, 0x63, 0x97, 0xf8, 0x69, 0xae, 0x66, 0x28, 0x3b, 0x9b, 0xd2, 0xf8, 0x7d, 0xa1,
	0xbc, 0xc6, 0x1a, 0xac, 0xb8, 0xb3, 0x58, 0x5c, 0x3f, 0xfb, 0xcb, 0x84, 0x95, 0xe1, 0x80, 0x25,
	0xdf, 0xe3, 0xe7, 0x60, 0xee, 0x2b, 0x93, 0xde, 0x78, 0x12, 0x32, 0x5c, 0x93, 0x74, 0x85, 0x69,
	0x16, 0x47, 0xeb, 0xbd, 0x3f, 0x6e, 0x6f, 0x8c, 0x4d, 0xc0, 0xba, 0x9d, 0x91, 0x4c, 0x4e, 0x3a,
	0x12, 0x8a, 0xb0, 0x0b, 0x66, 0x31, 0xee, 0x14, 0xbf, 0x21, 0x49, 0xf2, 0x58, 0xa6, 0x3f, 0x96,
	0x74, 0x0c, 0x6f, 0x06, 0x0a, 0xa8, 0xe9, 0xa7, 0xe8, 0x6b, 0x64, 0x23, 0x59, 0xfe, 0x90, 0x09,
	0x27, 0x0c, 0x35, 0xbd, 0x54, 0x7e, 0xae, 0x64, 0x35, 0x24, 0xfd, 0x1d, 0x58, 0x9f, 0x32, 0x31,
	0x1a, 0x87, 0x61, 0xc1, 0xdf, 0x91, 0xfc, 0x8d, 0x01, 0x0b, 0xd9, 0xb9, 0x2a, 0xbf, 0x77, 0xdd,
	0xf3, 0x4b, 0x0e, 0x1e, 0x28, 0xbc, 0x9d, 0x2a, 0xa8, 0x12, 0x18, 0x4d, 0xae, 0x47, 0x81, 0x8f,
	0xf7, 0xe0, 0xad, 0x61, 0xec, 0x8f, 0x05, 0x93, 0xb8, 0xbe, 0x1c, 0x52, 0xa5, 0x40, 0xd6, 0x41,
	0xfd, 0x4a, 0x61, 0xb5, 0x80, 0x1c, 0xa9, 0x2d, 0x3d, 0x38, 0xb9, 0x86, 0xe4, 0x7b, 0xfc, 0x80,
	0xf3, 0xea, 0x31, 0xe4, 0x53, 0xcc, 0x44, 0x24, 0x7f, 0x24, 0xf8, 0xe8, 0x82, 0x73, 0xbc, 0x0d,
	0x1b, 0x5a, 0xc2, 0x09, 0x43, 0xed, 0xa2, 0x72, 0x12, 0x8f, 0x76, 0x10, 0xfe, 0x02, 0xe0, 0x90,
	0x89, 0x3c, 0x4b, 0x1b, 0xe5, 0x9c, 0xb2, 0xcb, 0x12, 0x16, 0xcb, 0x6a, 0xeb, 0x60, 0xca, 0xa9,
	0x65, 0xf9, 0xdd, 0x41, 0xb8, 0x0d, 0xb5, 0x2f, 0x13, 0x1e, 0x2f, 0xba, 0x6c, 0x53, 0x12, 0x57,
	0x61, 0xd9, 0x4f, 0x78, 0xdc, 0xfc, 0x73, 0xed, 0xc7, 0xdb, 0x1b, 0xe3, 0xef, 0x55, 0xf8, 0x0d,
	0xc1, 0x8b, 0x7d, 0xea, 0x3a, 0x9e, 0x4b, 0x3c, 0x67, 0xef, 0xd8, 0x25, 0xea, 0x16, 0xec, 0xc0,
	0x27, 0x41, 0x24, 0xd8, 0x94, 0x25, 0xe4, 0x94, 0xf6, 0x5e, 0x3a, 0xf4, 0x8c, 0xbc, 0x70, 0xcf,
	0x5a, 0x44, 0x36, 0x47, 0x5e, 0x39, 0x74, 0xff, 0xc8, 0xa1, 0xf6, 0x27, 0x3b, 0xdb, 0x2d, 0x92,
	0x79, 0x20, 0x7b, 0x67, 0x9e, 0xeb, 0xb4, 0xe0, 0x71, 0xb1, 0x48, 0x65, 0xdc, 0x36, 0x46, 0x1f,
	0x34, 0x2b, 0x72, 0x66, 0x15, 0x06, 0x9f, 0x16, 0x47, 0xf8, 0x09, 0xc1, 0xb7, 0xbd, 0xfe, 0xc0,
	0xa5, 0x1e, 0xe9, 0xf5, 0xbd, 0x13, 0x6d, 0x8a, 0xd8, 0x81, 0xaf, 0x1d, 0xcc, 0xab, 0xb6, 0x48,
	0x51, 0x6e, 0x9b, 0xbc, 0x72, 0x8e, 0x87, 0xee, 0x80, 0xd8, 0x5d, 0x89, 0xeb, 0x6a, 0x60, 0x77,
	0x8e, 0xec, 0x96, 0xa0, 0xd2, 0xcd, 0x9d, 0xd8, 0x5a, 0xf3, 0x01, 0x97, 0x6d, 0x0c, 0xa1, 0x33,
	0x70, 0x8f, 0xdd, 0x7d, 0x8f, 0x2c, 0xac, 0x4c, 0x0e, 0xe8, 0xc9, 0x4b, 0xed, 0xb1, 0x79, 0x37,
	0xcd, 0xe5, 0xf6, 0xe6, 0x15, 0xe0, 0x3b, 0x78, 0xfe, 0x1f, 0x65, 0xc9, 0x57, 0x47, 0x2e, 0x75,
	0x49, 0xe0, 0x93, 0x5d, 0xd2, 0x0d, 0x7c, 0x69, 0xff, 0x61, 0xe8, 0x4b, 0x3d, 0x14, 0xb5, 0x7e,
	0x40, 0xf0, 0x91, 0x8e, 0x60, 0xa6, 0x94, 0x32, 0xa1, 0x6f, 0x71, 0x57, 0x8f, 0xe8, 0x9e, 0x34,
	0x81, 0x2d, 0xea, 0x7a, 0x43, 0xda, 0xef, 0xf5, 0x0f, 0x17, 0xdb, 0x93, 0x26, 0x1e, 0x2c, 0x4e,
	0xa5, 0x87, 0x6f, 0xa0, 0xf5, 0x6f, 0x16, 0xb6, 0x2e, 0x38, 0xdf, 0xaa, 0xe8, 0xae, 0x62, 0x99,
	0xaa, 0xaf, 0x48, 0xc0, 0xd1, 0xff, 0x9b, 0x65, 0x66, 0xc0, 0xe9, 0x9f, 0xd9, 0x6a, 0x0e, 0xa9,
	0x0a, 0x45, 0x79, 0xad, 0xac, 0x7b, 0x4b, 0x58, 0xea, 0xa8, 0x0d, 0x75, 0xb9, 0x46, 0x44, 0x65,
	0x38, 0xbb, 0x79, 0xb5, 0x58, 0xd5, 0x79, 0xfe, 0xf9, 0xf6, 0xc6, 0xf8, 0x1c, 0xde, 0x87, 0x05,
	0xff, 0x29, 0x75, 0x13, 0x1a, 0x50, 0xf9, 0x23, 0x5f, 0x37, 0x7f, 0xbd, 0xbd, 0x31, 0x1e, 0x4d,
	0x56, 0x15, 0xe9, 0xe3, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x56, 0x6c, 0xe8, 0xe4, 0x07,
	0x00, 0x00,
}

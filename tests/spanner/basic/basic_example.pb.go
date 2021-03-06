// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tests/spanner/basic/basic_example.proto

package basic // import "github.com/tcncloud/protoc-gen-persist/tests/spanner/basic"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/protoc-gen-go/descriptor"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/tcncloud/protoc-gen-persist/persist"
import test "github.com/tcncloud/protoc-gen-persist/tests/test"

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

type MyEnum int32

const (
	MyEnum_OPTION_0 MyEnum = 0
	MyEnum_OPTION_1 MyEnum = 1
)

var MyEnum_name = map[int32]string{
	0: "OPTION_0",
	1: "OPTION_1",
}
var MyEnum_value = map[string]int32{
	"OPTION_0": 0,
	"OPTION_1": 1,
}

func (x MyEnum) String() string {
	return proto.EnumName(MyEnum_name, int32(x))
}
func (MyEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_basic_example_2eb62dd7a8819437, []int{0}
}

type MappedEnum int32

const (
	MappedEnum_OPT_0 MappedEnum = 0
	MappedEnum_OPT_1 MappedEnum = 1
)

var MappedEnum_name = map[int32]string{
	0: "OPT_0",
	1: "OPT_1",
}
var MappedEnum_value = map[string]int32{
	"OPT_0": 0,
	"OPT_1": 1,
}

func (x MappedEnum) String() string {
	return proto.EnumName(MappedEnum_name, int32(x))
}
func (MappedEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_basic_example_2eb62dd7a8819437, []int{1}
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
	return fileDescriptor_basic_example_2eb62dd7a8819437, []int{0}
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

type Something struct {
	Thing                *Something_SomethingElse `protobuf:"bytes,2,opt,name=thing,proto3" json:"thing,omitempty"`
	Myenum               MyEnum                   `protobuf:"varint,3,opt,name=myenum,proto3,enum=tests.MyEnum" json:"myenum,omitempty"`
	Mappedenum           MappedEnum               `protobuf:"varint,4,opt,name=mappedenum,proto3,enum=tests.MappedEnum" json:"mappedenum,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Something) Reset()         { *m = Something{} }
func (m *Something) String() string { return proto.CompactTextString(m) }
func (*Something) ProtoMessage()    {}
func (*Something) Descriptor() ([]byte, []int) {
	return fileDescriptor_basic_example_2eb62dd7a8819437, []int{1}
}
func (m *Something) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Something.Unmarshal(m, b)
}
func (m *Something) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Something.Marshal(b, m, deterministic)
}
func (dst *Something) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Something.Merge(dst, src)
}
func (m *Something) XXX_Size() int {
	return xxx_messageInfo_Something.Size(m)
}
func (m *Something) XXX_DiscardUnknown() {
	xxx_messageInfo_Something.DiscardUnknown(m)
}

var xxx_messageInfo_Something proto.InternalMessageInfo

func (m *Something) GetThing() *Something_SomethingElse {
	if m != nil {
		return m.Thing
	}
	return nil
}

func (m *Something) GetMyenum() MyEnum {
	if m != nil {
		return m.Myenum
	}
	return MyEnum_OPTION_0
}

func (m *Something) GetMappedenum() MappedEnum {
	if m != nil {
		return m.Mappedenum
	}
	return MappedEnum_OPT_0
}

type Something_SomethingElse struct {
	Thing                string   `protobuf:"bytes,1,opt,name=thing,proto3" json:"thing,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Something_SomethingElse) Reset()         { *m = Something_SomethingElse{} }
func (m *Something_SomethingElse) String() string { return proto.CompactTextString(m) }
func (*Something_SomethingElse) ProtoMessage()    {}
func (*Something_SomethingElse) Descriptor() ([]byte, []int) {
	return fileDescriptor_basic_example_2eb62dd7a8819437, []int{1, 0}
}
func (m *Something_SomethingElse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Something_SomethingElse.Unmarshal(m, b)
}
func (m *Something_SomethingElse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Something_SomethingElse.Marshal(b, m, deterministic)
}
func (dst *Something_SomethingElse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Something_SomethingElse.Merge(dst, src)
}
func (m *Something_SomethingElse) XXX_Size() int {
	return xxx_messageInfo_Something_SomethingElse.Size(m)
}
func (m *Something_SomethingElse) XXX_DiscardUnknown() {
	xxx_messageInfo_Something_SomethingElse.DiscardUnknown(m)
}

var xxx_messageInfo_Something_SomethingElse proto.InternalMessageInfo

func (m *Something_SomethingElse) GetThing() string {
	if m != nil {
		return m.Thing
	}
	return ""
}

type HasTimestamp struct {
	Time                 *timestamp.Timestamp   `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Some                 *Something             `protobuf:"bytes,2,opt,name=some,proto3" json:"some,omitempty"`
	Str                  string                 `protobuf:"bytes,3,opt,name=str,proto3" json:"str,omitempty"`
	Table                *test.ExampleTable     `protobuf:"bytes,4,opt,name=table,proto3" json:"table,omitempty"`
	Strs                 []string               `protobuf:"bytes,5,rep,name=strs,proto3" json:"strs,omitempty"`
	Tables               []*test.ExampleTable   `protobuf:"bytes,6,rep,name=tables,proto3" json:"tables,omitempty"`
	Somes                []*Something           `protobuf:"bytes,7,rep,name=somes,proto3" json:"somes,omitempty"`
	Times                []*timestamp.Timestamp `protobuf:"bytes,8,rep,name=times,proto3" json:"times,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *HasTimestamp) Reset()         { *m = HasTimestamp{} }
func (m *HasTimestamp) String() string { return proto.CompactTextString(m) }
func (*HasTimestamp) ProtoMessage()    {}
func (*HasTimestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_basic_example_2eb62dd7a8819437, []int{2}
}
func (m *HasTimestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HasTimestamp.Unmarshal(m, b)
}
func (m *HasTimestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HasTimestamp.Marshal(b, m, deterministic)
}
func (dst *HasTimestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HasTimestamp.Merge(dst, src)
}
func (m *HasTimestamp) XXX_Size() int {
	return xxx_messageInfo_HasTimestamp.Size(m)
}
func (m *HasTimestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_HasTimestamp.DiscardUnknown(m)
}

var xxx_messageInfo_HasTimestamp proto.InternalMessageInfo

func (m *HasTimestamp) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *HasTimestamp) GetSome() *Something {
	if m != nil {
		return m.Some
	}
	return nil
}

func (m *HasTimestamp) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

func (m *HasTimestamp) GetTable() *test.ExampleTable {
	if m != nil {
		return m.Table
	}
	return nil
}

func (m *HasTimestamp) GetStrs() []string {
	if m != nil {
		return m.Strs
	}
	return nil
}

func (m *HasTimestamp) GetTables() []*test.ExampleTable {
	if m != nil {
		return m.Tables
	}
	return nil
}

func (m *HasTimestamp) GetSomes() []*Something {
	if m != nil {
		return m.Somes
	}
	return nil
}

func (m *HasTimestamp) GetTimes() []*timestamp.Timestamp {
	if m != nil {
		return m.Times
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "tests.Empty")
	proto.RegisterType((*Something)(nil), "tests.Something")
	proto.RegisterType((*Something_SomethingElse)(nil), "tests.Something.SomethingElse")
	proto.RegisterType((*HasTimestamp)(nil), "tests.HasTimestamp")
	proto.RegisterEnum("tests.MyEnum", MyEnum_name, MyEnum_value)
	proto.RegisterEnum("tests.MappedEnum", MappedEnum_name, MappedEnum_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExtraSrvClient is the client API for ExtraSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExtraSrvClient interface {
	ExtraUnary(ctx context.Context, in *test.NumRows, opts ...grpc.CallOption) (*test.ExampleTable, error)
	ExtraMethod(ctx context.Context, in *test.ExampleTable, opts ...grpc.CallOption) (*test.ExampleTable, error)
}

type extraSrvClient struct {
	cc *grpc.ClientConn
}

func NewExtraSrvClient(cc *grpc.ClientConn) ExtraSrvClient {
	return &extraSrvClient{cc}
}

func (c *extraSrvClient) ExtraUnary(ctx context.Context, in *test.NumRows, opts ...grpc.CallOption) (*test.ExampleTable, error) {
	out := new(test.ExampleTable)
	err := c.cc.Invoke(ctx, "/tests.ExtraSrv/ExtraUnary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extraSrvClient) ExtraMethod(ctx context.Context, in *test.ExampleTable, opts ...grpc.CallOption) (*test.ExampleTable, error) {
	out := new(test.ExampleTable)
	err := c.cc.Invoke(ctx, "/tests.ExtraSrv/ExtraMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExtraSrvServer is the server API for ExtraSrv service.
type ExtraSrvServer interface {
	ExtraUnary(context.Context, *test.NumRows) (*test.ExampleTable, error)
	ExtraMethod(context.Context, *test.ExampleTable) (*test.ExampleTable, error)
}

func RegisterExtraSrvServer(s *grpc.Server, srv ExtraSrvServer) {
	s.RegisterService(&_ExtraSrv_serviceDesc, srv)
}

func _ExtraSrv_ExtraUnary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(test.NumRows)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtraSrvServer).ExtraUnary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tests.ExtraSrv/ExtraUnary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtraSrvServer).ExtraUnary(ctx, req.(*test.NumRows))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExtraSrv_ExtraMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(test.ExampleTable)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtraSrvServer).ExtraMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tests.ExtraSrv/ExtraMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtraSrvServer).ExtraMethod(ctx, req.(*test.ExampleTable))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExtraSrv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tests.ExtraSrv",
	HandlerType: (*ExtraSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExtraUnary",
			Handler:    _ExtraSrv_ExtraUnary_Handler,
		},
		{
			MethodName: "ExtraMethod",
			Handler:    _ExtraSrv_ExtraMethod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tests/spanner/basic/basic_example.proto",
}

// MySpannerClient is the client API for MySpanner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MySpannerClient interface {
	UniaryInsert(ctx context.Context, in *test.ExampleTable, opts ...grpc.CallOption) (*test.ExampleTable, error)
	UniarySelect(ctx context.Context, in *test.ExampleTable, opts ...grpc.CallOption) (*test.ExampleTable, error)
	UniarySelectWithDirectives(ctx context.Context, in *test.ExampleTable, opts ...grpc.CallOption) (*test.ExampleTable, error)
	UniaryUpdate(ctx context.Context, in *test.ExampleTable, opts ...grpc.CallOption) (*test.PartialTable, error)
	UniaryDeleteRange(ctx context.Context, in *test.ExampleTableRange, opts ...grpc.CallOption) (*test.ExampleTable, error)
	ServerStream(ctx context.Context, in *test.Name, opts ...grpc.CallOption) (MySpanner_ServerStreamClient, error)
	ClientStreamInsert(ctx context.Context, opts ...grpc.CallOption) (MySpanner_ClientStreamInsertClient, error)
	ClientStreamDelete(ctx context.Context, opts ...grpc.CallOption) (MySpanner_ClientStreamDeleteClient, error)
	UniarySelectWithHooks(ctx context.Context, in *test.ExampleTable, opts ...grpc.CallOption) (*test.ExampleTable, error)
	ServerStreamWithHooks(ctx context.Context, in *test.Name, opts ...grpc.CallOption) (MySpanner_ServerStreamWithHooksClient, error)
	ClientStreamUpdateWithHooks(ctx context.Context, opts ...grpc.CallOption) (MySpanner_ClientStreamUpdateWithHooksClient, error)
}

type mySpannerClient struct {
	cc *grpc.ClientConn
}

func NewMySpannerClient(cc *grpc.ClientConn) MySpannerClient {
	return &mySpannerClient{cc}
}

func (c *mySpannerClient) UniaryInsert(ctx context.Context, in *test.ExampleTable, opts ...grpc.CallOption) (*test.ExampleTable, error) {
	out := new(test.ExampleTable)
	err := c.cc.Invoke(ctx, "/tests.MySpanner/UniaryInsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mySpannerClient) UniarySelect(ctx context.Context, in *test.ExampleTable, opts ...grpc.CallOption) (*test.ExampleTable, error) {
	out := new(test.ExampleTable)
	err := c.cc.Invoke(ctx, "/tests.MySpanner/UniarySelect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mySpannerClient) UniarySelectWithDirectives(ctx context.Context, in *test.ExampleTable, opts ...grpc.CallOption) (*test.ExampleTable, error) {
	out := new(test.ExampleTable)
	err := c.cc.Invoke(ctx, "/tests.MySpanner/UniarySelectWithDirectives", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mySpannerClient) UniaryUpdate(ctx context.Context, in *test.ExampleTable, opts ...grpc.CallOption) (*test.PartialTable, error) {
	out := new(test.PartialTable)
	err := c.cc.Invoke(ctx, "/tests.MySpanner/UniaryUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mySpannerClient) UniaryDeleteRange(ctx context.Context, in *test.ExampleTableRange, opts ...grpc.CallOption) (*test.ExampleTable, error) {
	out := new(test.ExampleTable)
	err := c.cc.Invoke(ctx, "/tests.MySpanner/UniaryDeleteRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mySpannerClient) ServerStream(ctx context.Context, in *test.Name, opts ...grpc.CallOption) (MySpanner_ServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MySpanner_serviceDesc.Streams[0], "/tests.MySpanner/ServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &mySpannerServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MySpanner_ServerStreamClient interface {
	Recv() (*test.ExampleTable, error)
	grpc.ClientStream
}

type mySpannerServerStreamClient struct {
	grpc.ClientStream
}

func (x *mySpannerServerStreamClient) Recv() (*test.ExampleTable, error) {
	m := new(test.ExampleTable)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mySpannerClient) ClientStreamInsert(ctx context.Context, opts ...grpc.CallOption) (MySpanner_ClientStreamInsertClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MySpanner_serviceDesc.Streams[1], "/tests.MySpanner/ClientStreamInsert", opts...)
	if err != nil {
		return nil, err
	}
	x := &mySpannerClientStreamInsertClient{stream}
	return x, nil
}

type MySpanner_ClientStreamInsertClient interface {
	Send(*test.ExampleTable) error
	CloseAndRecv() (*test.NumRows, error)
	grpc.ClientStream
}

type mySpannerClientStreamInsertClient struct {
	grpc.ClientStream
}

func (x *mySpannerClientStreamInsertClient) Send(m *test.ExampleTable) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mySpannerClientStreamInsertClient) CloseAndRecv() (*test.NumRows, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(test.NumRows)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mySpannerClient) ClientStreamDelete(ctx context.Context, opts ...grpc.CallOption) (MySpanner_ClientStreamDeleteClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MySpanner_serviceDesc.Streams[2], "/tests.MySpanner/ClientStreamDelete", opts...)
	if err != nil {
		return nil, err
	}
	x := &mySpannerClientStreamDeleteClient{stream}
	return x, nil
}

type MySpanner_ClientStreamDeleteClient interface {
	Send(*test.ExampleTable) error
	CloseAndRecv() (*test.NumRows, error)
	grpc.ClientStream
}

type mySpannerClientStreamDeleteClient struct {
	grpc.ClientStream
}

func (x *mySpannerClientStreamDeleteClient) Send(m *test.ExampleTable) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mySpannerClientStreamDeleteClient) CloseAndRecv() (*test.NumRows, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(test.NumRows)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mySpannerClient) UniarySelectWithHooks(ctx context.Context, in *test.ExampleTable, opts ...grpc.CallOption) (*test.ExampleTable, error) {
	out := new(test.ExampleTable)
	err := c.cc.Invoke(ctx, "/tests.MySpanner/UniarySelectWithHooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mySpannerClient) ServerStreamWithHooks(ctx context.Context, in *test.Name, opts ...grpc.CallOption) (MySpanner_ServerStreamWithHooksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MySpanner_serviceDesc.Streams[3], "/tests.MySpanner/ServerStreamWithHooks", opts...)
	if err != nil {
		return nil, err
	}
	x := &mySpannerServerStreamWithHooksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MySpanner_ServerStreamWithHooksClient interface {
	Recv() (*test.ExampleTable, error)
	grpc.ClientStream
}

type mySpannerServerStreamWithHooksClient struct {
	grpc.ClientStream
}

func (x *mySpannerServerStreamWithHooksClient) Recv() (*test.ExampleTable, error) {
	m := new(test.ExampleTable)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mySpannerClient) ClientStreamUpdateWithHooks(ctx context.Context, opts ...grpc.CallOption) (MySpanner_ClientStreamUpdateWithHooksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MySpanner_serviceDesc.Streams[4], "/tests.MySpanner/ClientStreamUpdateWithHooks", opts...)
	if err != nil {
		return nil, err
	}
	x := &mySpannerClientStreamUpdateWithHooksClient{stream}
	return x, nil
}

type MySpanner_ClientStreamUpdateWithHooksClient interface {
	Send(*test.ExampleTable) error
	CloseAndRecv() (*test.NumRows, error)
	grpc.ClientStream
}

type mySpannerClientStreamUpdateWithHooksClient struct {
	grpc.ClientStream
}

func (x *mySpannerClientStreamUpdateWithHooksClient) Send(m *test.ExampleTable) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mySpannerClientStreamUpdateWithHooksClient) CloseAndRecv() (*test.NumRows, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(test.NumRows)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MySpannerServer is the server API for MySpanner service.
type MySpannerServer interface {
	UniaryInsert(context.Context, *test.ExampleTable) (*test.ExampleTable, error)
	UniarySelect(context.Context, *test.ExampleTable) (*test.ExampleTable, error)
	UniarySelectWithDirectives(context.Context, *test.ExampleTable) (*test.ExampleTable, error)
	UniaryUpdate(context.Context, *test.ExampleTable) (*test.PartialTable, error)
	UniaryDeleteRange(context.Context, *test.ExampleTableRange) (*test.ExampleTable, error)
	ServerStream(*test.Name, MySpanner_ServerStreamServer) error
	ClientStreamInsert(MySpanner_ClientStreamInsertServer) error
	ClientStreamDelete(MySpanner_ClientStreamDeleteServer) error
	UniarySelectWithHooks(context.Context, *test.ExampleTable) (*test.ExampleTable, error)
	ServerStreamWithHooks(*test.Name, MySpanner_ServerStreamWithHooksServer) error
	ClientStreamUpdateWithHooks(MySpanner_ClientStreamUpdateWithHooksServer) error
}

func RegisterMySpannerServer(s *grpc.Server, srv MySpannerServer) {
	s.RegisterService(&_MySpanner_serviceDesc, srv)
}

func _MySpanner_UniaryInsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(test.ExampleTable)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MySpannerServer).UniaryInsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tests.MySpanner/UniaryInsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MySpannerServer).UniaryInsert(ctx, req.(*test.ExampleTable))
	}
	return interceptor(ctx, in, info, handler)
}

func _MySpanner_UniarySelect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(test.ExampleTable)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MySpannerServer).UniarySelect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tests.MySpanner/UniarySelect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MySpannerServer).UniarySelect(ctx, req.(*test.ExampleTable))
	}
	return interceptor(ctx, in, info, handler)
}

func _MySpanner_UniarySelectWithDirectives_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(test.ExampleTable)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MySpannerServer).UniarySelectWithDirectives(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tests.MySpanner/UniarySelectWithDirectives",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MySpannerServer).UniarySelectWithDirectives(ctx, req.(*test.ExampleTable))
	}
	return interceptor(ctx, in, info, handler)
}

func _MySpanner_UniaryUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(test.ExampleTable)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MySpannerServer).UniaryUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tests.MySpanner/UniaryUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MySpannerServer).UniaryUpdate(ctx, req.(*test.ExampleTable))
	}
	return interceptor(ctx, in, info, handler)
}

func _MySpanner_UniaryDeleteRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(test.ExampleTableRange)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MySpannerServer).UniaryDeleteRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tests.MySpanner/UniaryDeleteRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MySpannerServer).UniaryDeleteRange(ctx, req.(*test.ExampleTableRange))
	}
	return interceptor(ctx, in, info, handler)
}

func _MySpanner_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(test.Name)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MySpannerServer).ServerStream(m, &mySpannerServerStreamServer{stream})
}

type MySpanner_ServerStreamServer interface {
	Send(*test.ExampleTable) error
	grpc.ServerStream
}

type mySpannerServerStreamServer struct {
	grpc.ServerStream
}

func (x *mySpannerServerStreamServer) Send(m *test.ExampleTable) error {
	return x.ServerStream.SendMsg(m)
}

func _MySpanner_ClientStreamInsert_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MySpannerServer).ClientStreamInsert(&mySpannerClientStreamInsertServer{stream})
}

type MySpanner_ClientStreamInsertServer interface {
	SendAndClose(*test.NumRows) error
	Recv() (*test.ExampleTable, error)
	grpc.ServerStream
}

type mySpannerClientStreamInsertServer struct {
	grpc.ServerStream
}

func (x *mySpannerClientStreamInsertServer) SendAndClose(m *test.NumRows) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mySpannerClientStreamInsertServer) Recv() (*test.ExampleTable, error) {
	m := new(test.ExampleTable)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MySpanner_ClientStreamDelete_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MySpannerServer).ClientStreamDelete(&mySpannerClientStreamDeleteServer{stream})
}

type MySpanner_ClientStreamDeleteServer interface {
	SendAndClose(*test.NumRows) error
	Recv() (*test.ExampleTable, error)
	grpc.ServerStream
}

type mySpannerClientStreamDeleteServer struct {
	grpc.ServerStream
}

func (x *mySpannerClientStreamDeleteServer) SendAndClose(m *test.NumRows) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mySpannerClientStreamDeleteServer) Recv() (*test.ExampleTable, error) {
	m := new(test.ExampleTable)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MySpanner_UniarySelectWithHooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(test.ExampleTable)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MySpannerServer).UniarySelectWithHooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tests.MySpanner/UniarySelectWithHooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MySpannerServer).UniarySelectWithHooks(ctx, req.(*test.ExampleTable))
	}
	return interceptor(ctx, in, info, handler)
}

func _MySpanner_ServerStreamWithHooks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(test.Name)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MySpannerServer).ServerStreamWithHooks(m, &mySpannerServerStreamWithHooksServer{stream})
}

type MySpanner_ServerStreamWithHooksServer interface {
	Send(*test.ExampleTable) error
	grpc.ServerStream
}

type mySpannerServerStreamWithHooksServer struct {
	grpc.ServerStream
}

func (x *mySpannerServerStreamWithHooksServer) Send(m *test.ExampleTable) error {
	return x.ServerStream.SendMsg(m)
}

func _MySpanner_ClientStreamUpdateWithHooks_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MySpannerServer).ClientStreamUpdateWithHooks(&mySpannerClientStreamUpdateWithHooksServer{stream})
}

type MySpanner_ClientStreamUpdateWithHooksServer interface {
	SendAndClose(*test.NumRows) error
	Recv() (*test.ExampleTable, error)
	grpc.ServerStream
}

type mySpannerClientStreamUpdateWithHooksServer struct {
	grpc.ServerStream
}

func (x *mySpannerClientStreamUpdateWithHooksServer) SendAndClose(m *test.NumRows) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mySpannerClientStreamUpdateWithHooksServer) Recv() (*test.ExampleTable, error) {
	m := new(test.ExampleTable)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MySpanner_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tests.MySpanner",
	HandlerType: (*MySpannerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UniaryInsert",
			Handler:    _MySpanner_UniaryInsert_Handler,
		},
		{
			MethodName: "UniarySelect",
			Handler:    _MySpanner_UniarySelect_Handler,
		},
		{
			MethodName: "UniarySelectWithDirectives",
			Handler:    _MySpanner_UniarySelectWithDirectives_Handler,
		},
		{
			MethodName: "UniaryUpdate",
			Handler:    _MySpanner_UniaryUpdate_Handler,
		},
		{
			MethodName: "UniaryDeleteRange",
			Handler:    _MySpanner_UniaryDeleteRange_Handler,
		},
		{
			MethodName: "UniarySelectWithHooks",
			Handler:    _MySpanner_UniarySelectWithHooks_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStream",
			Handler:       _MySpanner_ServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStreamInsert",
			Handler:       _MySpanner_ClientStreamInsert_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ClientStreamDelete",
			Handler:       _MySpanner_ClientStreamDelete_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ServerStreamWithHooks",
			Handler:       _MySpanner_ServerStreamWithHooks_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStreamUpdateWithHooks",
			Handler:       _MySpanner_ClientStreamUpdateWithHooks_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "tests/spanner/basic/basic_example.proto",
}

func init() {
	proto.RegisterFile("tests/spanner/basic/basic_example.proto", fileDescriptor_basic_example_2eb62dd7a8819437)
}

var fileDescriptor_basic_example_2eb62dd7a8819437 = []byte{
	// 1206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x57, 0xdf, 0x6e, 0x1a, 0xc7,
	0x17, 0xfe, 0xad, 0x6d, 0x1c, 0x38, 0xc6, 0xf9, 0x91, 0x49, 0xad, 0x90, 0x4d, 0xd3, 0xac, 0x50,
	0x52, 0xaf, 0x51, 0x02, 0x84, 0x54, 0xea, 0x45, 0x65, 0x09, 0xc7, 0x6c, 0x1a, 0x94, 0x18, 0xc8,
	0x82, 0x93, 0xf4, 0x8f, 0x84, 0x06, 0x76, 0x02, 0xab, 0xec, 0x3f, 0xed, 0x0c, 0x6e, 0x50, 0x55,
	0xa9, 0xea, 0x45, 0x2b, 0x55, 0xaa, 0x54, 0xbf, 0x42, 0x2f, 0xfb, 0x04, 0xbe, 0xee, 0x45, 0x9f,
	0xa0, 0x37, 0x7d, 0x86, 0xbe, 0x44, 0x35, 0x33, 0x0b, 0x5e, 0xa8, 0x97, 0x92, 0xf4, 0x66, 0x34,
	0xcb, 0xf9, 0xe6, 0x3b, 0xe7, 0x7c, 0xe7, 0xcc, 0xd9, 0x05, 0x76, 0x19, 0xa1, 0x8c, 0x96, 0x69,
	0x80, 0x3d, 0x8f, 0x84, 0xe5, 0x3e, 0xa6, 0xf6, 0x40, 0xae, 0x3d, 0xf2, 0x06, 0xbb, 0x81, 0x43,
	0x4a, 0x41, 0xe8, 0x33, 0x1f, 0xa5, 0x04, 0x50, 0xdd, 0x09, 0x48, 0x48, 0x6d, 0xca, 0xca, 0x7e,
	0xc0, 0x6c, 0xdf, 0xa3, 0xd2, 0xaa, 0xde, 0x1a, 0xfa, 0xfe, 0xd0, 0x21, 0x65, 0xf1, 0xd4, 0x1f,
	0xbf, 0x2a, 0x33, 0xdb, 0x25, 0x94, 0x61, 0x37, 0x88, 0x00, 0xda, 0x22, 0xc0, 0x22, 0x74, 0x10,
	0xda, 0x01, 0xf3, 0xc3, 0x08, 0xb1, 0x23, 0x23, 0xe1, 0xab, 0x58, 0xe4, 0xcf, 0x85, 0x4b, 0x90,
	0x32, 0xdc, 0x80, 0x4d, 0x0a, 0xbf, 0x29, 0x90, 0xe9, 0xf8, 0x2e, 0x61, 0x23, 0xdb, 0x1b, 0xa2,
	0x8f, 0x20, 0x25, 0x36, 0xf9, 0x35, 0x4d, 0xd1, 0xb7, 0xaa, 0x1f, 0x94, 0xc4, 0xe9, 0xd2, 0x0c,
	0x70, 0xbe, 0x33, 0x1c, 0x4a, 0x4c, 0x09, 0x46, 0x77, 0x60, 0xd3, 0x9d, 0x10, 0x6f, 0xec, 0xe6,
	0xd7, 0x35, 0x45, 0xbf, 0x5c, 0xdd, 0x8e, 0x8e, 0x1d, 0x4d, 0x0c, 0x6f, 0xec, 0x9a, 0x91, 0x11,
	0xdd, 0x07, 0x70, 0x71, 0x10, 0x10, 0x4b, 0x40, 0x37, 0x04, 0xf4, 0xca, 0x14, 0x2a, 0x0c, 0x02,
	0x1e, 0x03, 0xa9, 0x77, 0x60, 0x7b, 0xce, 0x23, 0x7a, 0x6f, 0x1a, 0xa0, 0xa2, 0x29, 0x7a, 0x26,
	0x0a, 0xa0, 0xf0, 0xfb, 0x1a, 0x64, 0x1f, 0x63, 0xda, 0x9d, 0xaa, 0x83, 0x4a, 0xb0, 0xc1, 0xa5,
	0x12, 0xa8, 0xad, 0xaa, 0x5a, 0x92, 0x32, 0x95, 0xa6, 0x32, 0x95, 0x66, 0x48, 0x53, 0xe0, 0xd0,
	0x6d, 0xd8, 0xa0, 0xbe, 0x4b, 0xa2, 0xb4, 0x73, 0x8b, 0x69, 0x9b, 0xc2, 0x8a, 0x72, 0xb0, 0x4e,
	0x59, 0x28, 0x92, 0xcc, 0x98, 0x7c, 0x8b, 0x4a, 0x90, 0x62, 0xb8, 0xef, 0x10, 0x91, 0xcd, 0x56,
	0x35, 0x1f, 0x1d, 0x14, 0x42, 0x1b, 0xb2, 0xd0, 0x5d, 0x6e, 0x37, 0x25, 0x0c, 0x21, 0xd8, 0xa0,
	0x2c, 0xa4, 0xf9, 0x94, 0xb6, 0xae, 0x67, 0x4c, 0xb1, 0x47, 0x15, 0xd8, 0x14, 0x46, 0x9a, 0xdf,
	0xd4, 0xd6, 0x97, 0x92, 0x44, 0x38, 0xf4, 0x21, 0xa4, 0x78, 0x3c, 0x34, 0x7f, 0x49, 0x1c, 0xf8,
	0x67, 0xb8, 0xd2, 0x8c, 0x2a, 0x90, 0x12, 0x0d, 0x93, 0x4f, 0x0b, 0xdc, 0x32, 0x19, 0x24, 0xb0,
	0x78, 0x1b, 0x36, 0x65, 0xd1, 0x50, 0x16, 0xd2, 0xad, 0x76, 0xb7, 0xd1, 0x6a, 0xf6, 0x2a, 0xb9,
	0xff, 0xc5, 0x9e, 0xee, 0xe7, 0x94, 0x62, 0x01, 0xe0, 0xbc, 0x5e, 0x28, 0x03, 0xa9, 0x56, 0xbb,
	0x2b, 0x60, 0xd1, 0xf6, 0x7e, 0x4e, 0xa9, 0xfe, 0xa5, 0x40, 0xda, 0x78, 0xc3, 0x42, 0xdc, 0x09,
	0x4f, 0x50, 0x03, 0x40, 0xec, 0x8f, 0x3d, 0x1c, 0x4e, 0xd0, 0xd5, 0x78, 0x82, 0xcd, 0xb1, 0x6b,
	0xfa, 0x5f, 0x51, 0x35, 0x31, 0xeb, 0x42, 0xf6, 0x97, 0xb3, 0xd3, 0xb5, 0x4b, 0x90, 0x22, 0x9c,
	0x01, 0x1d, 0xc0, 0x96, 0xa0, 0x3a, 0x22, 0x6c, 0xe4, 0x5b, 0x28, 0xf1, 0x58, 0x32, 0xa1, 0xfa,
	0xec, 0xbb, 0xb3, 0xd3, 0xb5, 0x4f, 0xc1, 0x80, 0xeb, 0x1d, 0xe3, 0xa9, 0x71, 0xd8, 0xd5, 0x8a,
	0xda, 0x23, 0xb3, 0x75, 0xa4, 0x09, 0x07, 0xbd, 0x31, 0x8f, 0x51, 0x95, 0xde, 0x0a, 0xd9, 0x88,
	0x44, 0xdc, 0x95, 0xe2, 0x34, 0xf0, 0x78, 0xab, 0xfd, 0x7c, 0x76, 0xba, 0xa6, 0x54, 0xff, 0xc8,
	0x41, 0xe6, 0x68, 0xd2, 0x91, 0xd7, 0x1d, 0x3d, 0x83, 0xec, 0xb1, 0x67, 0xe3, 0x70, 0xd2, 0xf0,
	0x28, 0x09, 0xd9, 0xbb, 0x04, 0x59, 0xd8, 0xe6, 0x59, 0xa7, 0x61, 0xd3, 0x96, 0x14, 0x33, 0xca,
	0x0e, 0x71, 0xc8, 0xe0, 0xbf, 0x51, 0x52, 0x49, 0x41, 0x40, 0x8d, 0x53, 0xbe, 0xb0, 0xd9, 0xa8,
	0x6e, 0x87, 0x64, 0xc0, 0xec, 0x13, 0x42, 0xdf, 0xc9, 0xc1, 0x55, 0xee, 0xe0, 0x32, 0x64, 0xa5,
	0x83, 0x9e, 0xed, 0x59, 0xe4, 0xcd, 0x79, 0xe4, 0xc7, 0x81, 0x85, 0x19, 0x59, 0x95, 0xb8, 0x8d,
	0x43, 0x66, 0x63, 0x67, 0x3e, 0xf2, 0xb1, 0xa4, 0xf8, 0x02, 0xae, 0x48, 0xca, 0x3a, 0x71, 0x08,
	0x23, 0x26, 0xf6, 0x86, 0x04, 0xdd, 0x4c, 0xbc, 0x36, 0xdc, 0xbc, 0x82, 0x2c, 0x96, 0x20, 0x43,
	0x2d, 0xc8, 0x76, 0x48, 0x78, 0x42, 0xc2, 0x0e, 0x0b, 0x09, 0x76, 0x51, 0x6e, 0xae, 0x5b, 0xb1,
	0xbb, 0x8c, 0xea, 0x0a, 0xa7, 0xca, 0x02, 0x44, 0x02, 0x60, 0xc7, 0xa9, 0x28, 0xe8, 0x39, 0xa0,
	0x43, 0xc7, 0x26, 0x1e, 0x93, 0x84, 0xff, 0xda, 0x13, 0x17, 0x5d, 0x8f, 0xc2, 0xff, 0x39, 0x33,
	0x40, 0x5a, 0xb6, 0x43, 0xef, 0x81, 0xae, 0xa0, 0x17, 0xf3, 0xbc, 0x52, 0x8b, 0xb7, 0xe5, 0xcd,
	0x71, 0xde, 0x2d, 0xc8, 0xc8, 0xe4, 0x7b, 0xb6, 0xa5, 0x2b, 0xc8, 0x82, 0x9d, 0xc5, 0xc6, 0x78,
	0xec, 0xfb, 0xaf, 0xdf, 0xad, 0x27, 0xae, 0x71, 0x07, 0x68, 0xbe, 0x27, 0xda, 0xca, 0x4b, 0x05,
	0x7d, 0x09, 0x3b, 0x71, 0x9d, 0xcf, 0xbd, 0xbc, 0x8d, 0xe0, 0x3b, 0x9c, 0x3d, 0x17, 0x17, 0x9c,
	0x73, 0x57, 0x14, 0x34, 0x82, 0x1b, 0x71, 0x71, 0x64, 0xef, 0xad, 0x92, 0xc9, 0x85, 0x2a, 0xe5,
	0xb9, 0x9b, 0xab, 0xb0, 0x4d, 0x09, 0xeb, 0x79, 0xd8, 0x25, 0x3d, 0x4c, 0xad, 0x57, 0xdc, 0x93,
	0xae, 0xa8, 0x7f, 0xa6, 0xf9, 0x38, 0xf9, 0x35, 0x0d, 0x3f, 0x28, 0x50, 0x91, 0xc5, 0xd5, 0x6c,
	0x8f, 0xf9, 0x5a, 0xf4, 0xa2, 0xef, 0x89, 0xa1, 0xad, 0xe9, 0xb6, 0x75, 0x57, 0xa3, 0x0c, 0x87,
	0xac, 0xc7, 0x47, 0xed, 0x5d, 0x8d, 0xf3, 0xec, 0xc1, 0xed, 0xe7, 0xd8, 0x19, 0x13, 0xaa, 0xe9,
	0x35, 0x6e, 0xaf, 0xc5, 0x01, 0xbb, 0x7d, 0xec, 0x61, 0x0f, 0xd3, 0xdd, 0x3d, 0x35, 0x9a, 0x02,
	0x85, 0x6b, 0x09, 0x51, 0x17, 0xe7, 0xe6, 0x14, 0x30, 0xf8, 0x78, 0x36, 0xd9, 0x5e, 0x85, 0xbe,
	0xbb, 0x10, 0xca, 0x8b, 0x11, 0x09, 0x89, 0x66, 0x5b, 0xfb, 0x35, 0xdb, 0xd2, 0x0e, 0x9a, 0x75,
	0x11, 0xca, 0x7e, 0x8d, 0xaf, 0x6a, 0x34, 0x1c, 0x92, 0x3d, 0x25, 0x2a, 0x07, 0x3f, 0x29, 0xf0,
	0x64, 0x89, 0xdb, 0xda, 0xd7, 0x8f, 0x5a, 0xe6, 0xa1, 0xd1, 0x6b, 0x34, 0xeb, 0xc6, 0xcb, 0x7d,
	0xd1, 0x07, 0xdf, 0x2c, 0x0d, 0x65, 0xae, 0x65, 0x92, 0x03, 0x4a, 0x32, 0xc0, 0xf7, 0x0a, 0xec,
	0xc8, 0xba, 0x2f, 0xe4, 0x0f, 0x45, 0x4a, 0x58, 0xac, 0x16, 0xfb, 0xb5, 0xc5, 0xba, 0xec, 0xef,
	0xfa, 0x21, 0x9f, 0x1a, 0x74, 0x17, 0x50, 0x3b, 0xb4, 0x5d, 0x1c, 0x4e, 0x9e, 0x90, 0x89, 0x2e,
	0xe3, 0xdc, 0x53, 0xa3, 0x39, 0xb4, 0x6a, 0x39, 0x4e, 0xa0, 0x56, 0x37, 0x9e, 0x1a, 0x5d, 0x63,
	0xfa, 0x9a, 0x89, 0x07, 0xd3, 0xe9, 0x1e, 0x98, 0x5d, 0x3d, 0x8a, 0xc0, 0xb6, 0xf6, 0x34, 0xa3,
	0x59, 0xd7, 0x6b, 0xc4, 0xb3, 0xc4, 0xc3, 0x93, 0x46, 0xb3, 0xae, 0x1f, 0xb6, 0xf6, 0xd4, 0x68,
	0x3a, 0x15, 0xde, 0x5f, 0x36, 0xe4, 0x16, 0xfc, 0xb6, 0xe1, 0xc6, 0xe2, 0x0b, 0x2e, 0xe6, 0x59,
	0x8d, 0x5d, 0x9a, 0x85, 0xf7, 0x5c, 0xa2, 0xa4, 0xdf, 0x2a, 0x70, 0xd4, 0x68, 0x76, 0x0c, 0xb3,
	0xab, 0x35, 0x9a, 0xdd, 0xd6, 0x4a, 0x2d, 0xae, 0x3d, 0x3f, 0x78, 0x7a, 0x6c, 0x74, 0x2e, 0xea,
	0xf0, 0x07, 0x7b, 0xea, 0x6c, 0xa0, 0xad, 0x2a, 0xe6, 0x67, 0x70, 0x27, 0x59, 0x4c, 0xe9, 0x4a,
	0x17, 0x45, 0x3a, 0x1f, 0x69, 0xab, 0x52, 0x0f, 0xa1, 0x3c, 0xbe, 0xa8, 0x5f, 0x78, 0xb7, 0xc8,
	0xae, 0xe0, 0xd7, 0x7e, 0x57, 0x0b, 0x5e, 0xcf, 0x3a, 0x61, 0x7e, 0x22, 0xac, 0xe8, 0xe8, 0xc7,
	0xb3, 0xd3, 0xb5, 0x5b, 0x70, 0x13, 0xae, 0x27, 0x7e, 0x8e, 0xe5, 0xb6, 0xc4, 0x67, 0xc5, 0xc3,
	0x87, 0x9f, 0xd7, 0x86, 0x36, 0x1b, 0x8d, 0xfb, 0xa5, 0x81, 0xef, 0x96, 0xd9, 0xc0, 0x1b, 0x38,
	0xfe, 0xd8, 0x92, 0x5f, 0xfb, 0x83, 0x7b, 0x43, 0xe2, 0xdd, 0x9b, 0xfe, 0x6f, 0xb8, 0xe0, 0xff,
	0xc6, 0x27, 0x62, 0xed, 0x6f, 0x0a, 0xf4, 0x83, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x37,
	0x2b, 0xd6, 0x93, 0x0c, 0x00, 0x00,
}

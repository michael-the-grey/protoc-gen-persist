// Code generated by protoc-gen-go.
// source: examples/sql/basic/basic.proto
// DO NOT EDIT!

/*
Package basic is a generated protocol buffer package.

It is generated from these files:
	examples/sql/basic/basic.proto

It has these top-level messages:
	ExampleTable
	PartialTable
	Name
	NumRows
*/
package basic

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/tcncloud/protoc-gen-persist/persist"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/golang/protobuf/protoc-gen-go/descriptor"
import _ "github.com/tcncloud/protoc-gen-persist/examples/test"

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

type ExampleTable struct {
	Id        int64                       `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	StartTime *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	Name      string                      `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *ExampleTable) Reset()                    { *m = ExampleTable{} }
func (m *ExampleTable) String() string            { return proto.CompactTextString(m) }
func (*ExampleTable) ProtoMessage()               {}
func (*ExampleTable) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ExampleTable) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ExampleTable) GetStartTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *ExampleTable) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PartialTable struct {
	Id        int64                       `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	StartTime *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
}

func (m *PartialTable) Reset()                    { *m = PartialTable{} }
func (m *PartialTable) String() string            { return proto.CompactTextString(m) }
func (*PartialTable) ProtoMessage()               {}
func (*PartialTable) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PartialTable) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PartialTable) GetStartTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

type Name struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Name) Reset()                    { *m = Name{} }
func (m *Name) String() string            { return proto.CompactTextString(m) }
func (*Name) ProtoMessage()               {}
func (*Name) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Name) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type NumRows struct {
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *NumRows) Reset()                    { *m = NumRows{} }
func (m *NumRows) String() string            { return proto.CompactTextString(m) }
func (*NumRows) ProtoMessage()               {}
func (*NumRows) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *NumRows) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*ExampleTable)(nil), "examples.ExampleTable")
	proto.RegisterType((*PartialTable)(nil), "examples.PartialTable")
	proto.RegisterType((*Name)(nil), "examples.Name")
	proto.RegisterType((*NumRows)(nil), "examples.NumRows")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Amazing service

type AmazingClient interface {
	UniarySelect(ctx context.Context, in *PartialTable, opts ...grpc.CallOption) (*ExampleTable, error)
	ServerStream(ctx context.Context, in *Name, opts ...grpc.CallOption) (Amazing_ServerStreamClient, error)
	Bidirectional(ctx context.Context, opts ...grpc.CallOption) (Amazing_BidirectionalClient, error)
	ClientStream(ctx context.Context, opts ...grpc.CallOption) (Amazing_ClientStreamClient, error)
}

type amazingClient struct {
	cc *grpc.ClientConn
}

func NewAmazingClient(cc *grpc.ClientConn) AmazingClient {
	return &amazingClient{cc}
}

func (c *amazingClient) UniarySelect(ctx context.Context, in *PartialTable, opts ...grpc.CallOption) (*ExampleTable, error) {
	out := new(ExampleTable)
	err := grpc.Invoke(ctx, "/examples.Amazing/UniarySelect", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *amazingClient) ServerStream(ctx context.Context, in *Name, opts ...grpc.CallOption) (Amazing_ServerStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Amazing_serviceDesc.Streams[0], c.cc, "/examples.Amazing/ServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &amazingServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Amazing_ServerStreamClient interface {
	Recv() (*ExampleTable, error)
	grpc.ClientStream
}

type amazingServerStreamClient struct {
	grpc.ClientStream
}

func (x *amazingServerStreamClient) Recv() (*ExampleTable, error) {
	m := new(ExampleTable)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *amazingClient) Bidirectional(ctx context.Context, opts ...grpc.CallOption) (Amazing_BidirectionalClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Amazing_serviceDesc.Streams[1], c.cc, "/examples.Amazing/Bidirectional", opts...)
	if err != nil {
		return nil, err
	}
	x := &amazingBidirectionalClient{stream}
	return x, nil
}

type Amazing_BidirectionalClient interface {
	Send(*ExampleTable) error
	Recv() (*ExampleTable, error)
	grpc.ClientStream
}

type amazingBidirectionalClient struct {
	grpc.ClientStream
}

func (x *amazingBidirectionalClient) Send(m *ExampleTable) error {
	return x.ClientStream.SendMsg(m)
}

func (x *amazingBidirectionalClient) Recv() (*ExampleTable, error) {
	m := new(ExampleTable)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *amazingClient) ClientStream(ctx context.Context, opts ...grpc.CallOption) (Amazing_ClientStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Amazing_serviceDesc.Streams[2], c.cc, "/examples.Amazing/ClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &amazingClientStreamClient{stream}
	return x, nil
}

type Amazing_ClientStreamClient interface {
	Send(*ExampleTable) error
	CloseAndRecv() (*NumRows, error)
	grpc.ClientStream
}

type amazingClientStreamClient struct {
	grpc.ClientStream
}

func (x *amazingClientStreamClient) Send(m *ExampleTable) error {
	return x.ClientStream.SendMsg(m)
}

func (x *amazingClientStreamClient) CloseAndRecv() (*NumRows, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(NumRows)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Amazing service

type AmazingServer interface {
	UniarySelect(context.Context, *PartialTable) (*ExampleTable, error)
	ServerStream(*Name, Amazing_ServerStreamServer) error
	Bidirectional(Amazing_BidirectionalServer) error
	ClientStream(Amazing_ClientStreamServer) error
}

func RegisterAmazingServer(s *grpc.Server, srv AmazingServer) {
	s.RegisterService(&_Amazing_serviceDesc, srv)
}

func _Amazing_UniarySelect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PartialTable)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AmazingServer).UniarySelect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.Amazing/UniarySelect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AmazingServer).UniarySelect(ctx, req.(*PartialTable))
	}
	return interceptor(ctx, in, info, handler)
}

func _Amazing_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Name)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AmazingServer).ServerStream(m, &amazingServerStreamServer{stream})
}

type Amazing_ServerStreamServer interface {
	Send(*ExampleTable) error
	grpc.ServerStream
}

type amazingServerStreamServer struct {
	grpc.ServerStream
}

func (x *amazingServerStreamServer) Send(m *ExampleTable) error {
	return x.ServerStream.SendMsg(m)
}

func _Amazing_Bidirectional_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AmazingServer).Bidirectional(&amazingBidirectionalServer{stream})
}

type Amazing_BidirectionalServer interface {
	Send(*ExampleTable) error
	Recv() (*ExampleTable, error)
	grpc.ServerStream
}

type amazingBidirectionalServer struct {
	grpc.ServerStream
}

func (x *amazingBidirectionalServer) Send(m *ExampleTable) error {
	return x.ServerStream.SendMsg(m)
}

func (x *amazingBidirectionalServer) Recv() (*ExampleTable, error) {
	m := new(ExampleTable)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Amazing_ClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AmazingServer).ClientStream(&amazingClientStreamServer{stream})
}

type Amazing_ClientStreamServer interface {
	SendAndClose(*NumRows) error
	Recv() (*ExampleTable, error)
	grpc.ServerStream
}

type amazingClientStreamServer struct {
	grpc.ServerStream
}

func (x *amazingClientStreamServer) SendAndClose(m *NumRows) error {
	return x.ServerStream.SendMsg(m)
}

func (x *amazingClientStreamServer) Recv() (*ExampleTable, error) {
	m := new(ExampleTable)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Amazing_serviceDesc = grpc.ServiceDesc{
	ServiceName: "examples.Amazing",
	HandlerType: (*AmazingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UniarySelect",
			Handler:    _Amazing_UniarySelect_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStream",
			Handler:       _Amazing_ServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Bidirectional",
			Handler:       _Amazing_Bidirectional_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ClientStream",
			Handler:       _Amazing_ClientStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "examples/sql/basic/basic.proto",
}

func init() { proto.RegisterFile("examples/sql/basic/basic.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x6e, 0xd3, 0x30,
	0x1c, 0xc7, 0x49, 0xf7, 0x8f, 0x79, 0x65, 0x02, 0x0b, 0x50, 0x95, 0x03, 0xb3, 0xa2, 0x1c, 0xb2,
	0xa9, 0x4b, 0xd6, 0x4e, 0x9a, 0x34, 0xa1, 0x0a, 0x75, 0x5b, 0x80, 0x89, 0x2d, 0x9b, 0x92, 0x14,
	0x04, 0x42, 0x0c, 0x37, 0xf1, 0x3a, 0x43, 0x12, 0x07, 0xc7, 0x05, 0xc6, 0x91, 0x23, 0xa7, 0xed,
	0x05, 0x78, 0x06, 0x8e, 0x7b, 0x0a, 0x9e, 0x09, 0x39, 0x49, 0xbb, 0xa8, 0x53, 0x91, 0x90, 0xb8,
	0x58, 0xfd, 0xd5, 0xdf, 0xe8, 0xf3, 0xf1, 0x57, 0x4e, 0xc0, 0x23, 0xf2, 0x15, 0xc7, 0x69, 0x44,
	0x32, 0x2b, 0xfb, 0x14, 0x59, 0x7d, 0x9c, 0xd1, 0xa0, 0x58, 0xcd, 0x94, 0x33, 0xc1, 0xe0, 0xed,
	0xd1, 0xbe, 0xfa, 0x20, 0x25, 0x3c, 0xa3, 0x99, 0xb0, 0x58, 0x2a, 0x28, 0x4b, 0xb2, 0x22, 0xa0,
	0xae, 0x0c, 0x18, 0x1b, 0x44, 0xc4, 0xca, 0xa7, 0xfe, 0xf0, 0xd4, 0x12, 0x34, 0x26, 0x99, 0xc0,
	0x71, 0x5a, 0x06, 0xd0, 0x64, 0x20, 0x24, 0x59, 0xc0, 0x69, 0x2a, 0x18, 0x2f, 0x13, 0x8d, 0xb1,
	0x83, 0x20, 0x99, 0xc8, 0x97, 0x62, 0x47, 0x8b, 0x41, 0xdd, 0x2e, 0xf6, 0x7c, 0xdc, 0x8f, 0x08,
	0x5c, 0x06, 0x35, 0x1a, 0x36, 0x14, 0xa4, 0x18, 0x33, 0x6e, 0x8d, 0x86, 0x70, 0x1b, 0x80, 0x4c,
	0x60, 0x2e, 0x4e, 0x24, 0xb4, 0x51, 0x43, 0x8a, 0xb1, 0xd4, 0x56, 0xcd, 0x02, 0x68, 0x8e, 0x80,
	0xa6, 0x3f, 0x32, 0x72, 0x17, 0xf3, 0xb4, 0x9c, 0x21, 0x04, 0xb3, 0x09, 0x8e, 0x49, 0x63, 0x06,
	0x29, 0xc6, 0xa2, 0x9b, 0xff, 0xd6, 0x5e, 0x83, 0xfa, 0x31, 0xe6, 0x82, 0xe2, 0xe8, 0x7f, 0xe3,
	0x34, 0x15, 0xcc, 0x3a, 0xb8, 0x82, 0x55, 0x2a, 0xd8, 0x15, 0xb0, 0xe0, 0x0c, 0x63, 0x97, 0x7d,
	0xc9, 0xe0, 0x7d, 0x30, 0x17, 0xb0, 0x61, 0x22, 0x4a, 0x68, 0x31, 0xb4, 0x7f, 0xcf, 0x81, 0x85,
	0x6e, 0x8c, 0xbf, 0xd1, 0x64, 0x00, 0x2f, 0x14, 0x50, 0xef, 0x25, 0x14, 0xf3, 0x73, 0x8f, 0x44,
	0x24, 0x10, 0xf0, 0xa1, 0x39, 0xaa, 0xcf, 0xac, 0xca, 0xab, 0x95, 0xff, 0xab, 0x1d, 0x6a, 0xc7,
	0xdf, 0xaf, 0x2e, 0x6b, 0x2f, 0xc0, 0xb6, 0x67, 0x1f, 0xd8, 0xbb, 0x3e, 0x5a, 0x43, 0xa7, 0x9c,
	0xc5, 0xa8, 0xcc, 0x9e, 0x08, 0x99, 0x41, 0xaf, 0xce, 0x08, 0x27, 0x88, 0x86, 0x1d, 0xbd, 0x85,
	0xba, 0xce, 0x1e, 0xba, 0x3e, 0x75, 0x47, 0x6f, 0x43, 0x59, 0x45, 0xa5, 0x08, 0x98, 0x82, 0xba,
	0x47, 0xf8, 0x67, 0xc2, 0x3d, 0xc1, 0x09, 0x8e, 0xe1, 0xf2, 0x35, 0x59, 0x9e, 0x79, 0xaa, 0xc9,
	0x96, 0x34, 0x69, 0x81, 0xd5, 0xb1, 0xc9, 0x53, 0xf7, 0xe8, 0x70, 0xd2, 0xe4, 0xb9, 0xed, 0xda,
	0x48, 0xd6, 0xd4, 0xd1, 0x5b, 0x30, 0xef, 0x6b, 0x43, 0x81, 0xbf, 0x14, 0x70, 0x67, 0x87, 0x86,
	0x94, 0x93, 0x40, 0xde, 0x45, 0x1c, 0xc1, 0x29, 0x8c, 0xa9, 0xec, 0x8f, 0x92, 0x7d, 0x0a, 0x9c,
	0xde, 0xf1, 0x5e, 0xd7, 0xb7, 0x27, 0xa0, 0x9e, 0xed, 0x23, 0xe3, 0xfa, 0x80, 0xcd, 0x9c, 0xbf,
	0x8a, 0x3a, 0xc8, 0xd0, 0xdb, 0x4d, 0xa4, 0x6f, 0xae, 0x96, 0x5a, 0x45, 0x41, 0xae, 0xed, 0xf7,
	0x5c, 0x67, 0xdf, 0x79, 0x86, 0xd6, 0x6e, 0x54, 0x93, 0x0b, 0x1b, 0xca, 0x86, 0x02, 0x7f, 0x2a,
	0xa0, 0xbe, 0x1b, 0x51, 0x92, 0x88, 0xb2, 0xa5, 0x69, 0xc6, 0xf7, 0x2a, 0xed, 0x15, 0xb7, 0x42,
	0xc3, 0x52, 0xf6, 0x2d, 0xd8, 0xdb, 0x77, 0x3c, 0xdb, 0xf5, 0xd1, 0xbe, 0xe3, 0x1f, 0x4d, 0x18,
	0x1b, 0x34, 0x6c, 0xa2, 0x9b, 0xc6, 0x2f, 0xbb, 0x07, 0x3d, 0xdb, 0x43, 0x86, 0xde, 0x6a, 0xa2,
	0x52, 0x7d, 0x9a, 0xa2, 0xfa, 0xe1, 0xc7, 0xd5, 0x65, 0xed, 0x3d, 0x78, 0x07, 0xfe, 0x72, 0xa3,
	0xef, 0x2e, 0x69, 0xf3, 0x87, 0xe7, 0x72, 0x5c, 0xdb, 0x1a, 0x50, 0x71, 0x36, 0xec, 0x9b, 0x01,
	0x8b, 0x2d, 0x11, 0x24, 0x41, 0xc4, 0x86, 0x61, 0xf1, 0x8e, 0x07, 0xeb, 0x03, 0x92, 0xac, 0x8f,
	0xbe, 0x16, 0xe3, 0x77, 0x3b, 0x3e, 0x97, 0xb8, 0x8b, 0xab, 0xcb, 0xda, 0xad, 0x9d, 0xee, 0x9b,
	0x27, 0xff, 0xfa, 0xf8, 0xf8, 0xf3, 0xf4, 0x38, 0x5f, 0xfb, 0xf3, 0x79, 0x78, 0xf3, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x08, 0xd1, 0x19, 0xca, 0xc1, 0x04, 0x00, 0x00,
}
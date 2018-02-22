// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tests/spanner/bob_example/bobs.proto

/*
Package bob_example is a generated protocol buffer package.

It is generated from these files:
	tests/spanner/bob_example/bobs.proto

It has these top-level messages:
	Bob
	NumRows
	Names
	Empty
*/
package bob_example

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/tcncloud/protoc-gen-persist/persist"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

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

type Bob struct {
	Id        int64                       `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	StartTime *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	Name      string                      `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *Bob) Reset()                    { *m = Bob{} }
func (m *Bob) String() string            { return proto.CompactTextString(m) }
func (*Bob) ProtoMessage()               {}
func (*Bob) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Bob) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Bob) GetStartTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Bob) GetName() string {
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
func (*NumRows) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NumRows) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Names struct {
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
}

func (m *Names) Reset()                    { *m = Names{} }
func (m *Names) String() string            { return proto.CompactTextString(m) }
func (*Names) ProtoMessage()               {}
func (*Names) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Names) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Bob)(nil), "bob.Bob")
	proto.RegisterType((*NumRows)(nil), "bob.NumRows")
	proto.RegisterType((*Names)(nil), "bob.Names")
	proto.RegisterType((*Empty)(nil), "bob.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Bobs service

type BobsClient interface {
	// Delete all bob events before the time in the end key
	DeleteBobs(ctx context.Context, in *Bob, opts ...grpc.CallOption) (*Empty, error)
	PutBobs(ctx context.Context, opts ...grpc.CallOption) (Bobs_PutBobsClient, error)
	GetBobs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Bobs_GetBobsClient, error)
	GetPeopleFromNames(ctx context.Context, in *Names, opts ...grpc.CallOption) (Bobs_GetPeopleFromNamesClient, error)
}

type bobsClient struct {
	cc *grpc.ClientConn
}

func NewBobsClient(cc *grpc.ClientConn) BobsClient {
	return &bobsClient{cc}
}

func (c *bobsClient) DeleteBobs(ctx context.Context, in *Bob, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/bob.Bobs/DeleteBobs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bobsClient) PutBobs(ctx context.Context, opts ...grpc.CallOption) (Bobs_PutBobsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Bobs_serviceDesc.Streams[0], c.cc, "/bob.Bobs/PutBobs", opts...)
	if err != nil {
		return nil, err
	}
	x := &bobsPutBobsClient{stream}
	return x, nil
}

type Bobs_PutBobsClient interface {
	Send(*Bob) error
	CloseAndRecv() (*NumRows, error)
	grpc.ClientStream
}

type bobsPutBobsClient struct {
	grpc.ClientStream
}

func (x *bobsPutBobsClient) Send(m *Bob) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bobsPutBobsClient) CloseAndRecv() (*NumRows, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(NumRows)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bobsClient) GetBobs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Bobs_GetBobsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Bobs_serviceDesc.Streams[1], c.cc, "/bob.Bobs/GetBobs", opts...)
	if err != nil {
		return nil, err
	}
	x := &bobsGetBobsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Bobs_GetBobsClient interface {
	Recv() (*Bob, error)
	grpc.ClientStream
}

type bobsGetBobsClient struct {
	grpc.ClientStream
}

func (x *bobsGetBobsClient) Recv() (*Bob, error) {
	m := new(Bob)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bobsClient) GetPeopleFromNames(ctx context.Context, in *Names, opts ...grpc.CallOption) (Bobs_GetPeopleFromNamesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Bobs_serviceDesc.Streams[2], c.cc, "/bob.Bobs/GetPeopleFromNames", opts...)
	if err != nil {
		return nil, err
	}
	x := &bobsGetPeopleFromNamesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Bobs_GetPeopleFromNamesClient interface {
	Recv() (*Bob, error)
	grpc.ClientStream
}

type bobsGetPeopleFromNamesClient struct {
	grpc.ClientStream
}

func (x *bobsGetPeopleFromNamesClient) Recv() (*Bob, error) {
	m := new(Bob)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Bobs service

type BobsServer interface {
	// Delete all bob events before the time in the end key
	DeleteBobs(context.Context, *Bob) (*Empty, error)
	PutBobs(Bobs_PutBobsServer) error
	GetBobs(*Empty, Bobs_GetBobsServer) error
	GetPeopleFromNames(*Names, Bobs_GetPeopleFromNamesServer) error
}

func RegisterBobsServer(s *grpc.Server, srv BobsServer) {
	s.RegisterService(&_Bobs_serviceDesc, srv)
}

func _Bobs_DeleteBobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bob)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BobsServer).DeleteBobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bob.Bobs/DeleteBobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BobsServer).DeleteBobs(ctx, req.(*Bob))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bobs_PutBobs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BobsServer).PutBobs(&bobsPutBobsServer{stream})
}

type Bobs_PutBobsServer interface {
	SendAndClose(*NumRows) error
	Recv() (*Bob, error)
	grpc.ServerStream
}

type bobsPutBobsServer struct {
	grpc.ServerStream
}

func (x *bobsPutBobsServer) SendAndClose(m *NumRows) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bobsPutBobsServer) Recv() (*Bob, error) {
	m := new(Bob)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Bobs_GetBobs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BobsServer).GetBobs(m, &bobsGetBobsServer{stream})
}

type Bobs_GetBobsServer interface {
	Send(*Bob) error
	grpc.ServerStream
}

type bobsGetBobsServer struct {
	grpc.ServerStream
}

func (x *bobsGetBobsServer) Send(m *Bob) error {
	return x.ServerStream.SendMsg(m)
}

func _Bobs_GetPeopleFromNames_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Names)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BobsServer).GetPeopleFromNames(m, &bobsGetPeopleFromNamesServer{stream})
}

type Bobs_GetPeopleFromNamesServer interface {
	Send(*Bob) error
	grpc.ServerStream
}

type bobsGetPeopleFromNamesServer struct {
	grpc.ServerStream
}

func (x *bobsGetPeopleFromNamesServer) Send(m *Bob) error {
	return x.ServerStream.SendMsg(m)
}

var _Bobs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bob.Bobs",
	HandlerType: (*BobsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteBobs",
			Handler:    _Bobs_DeleteBobs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PutBobs",
			Handler:       _Bobs_PutBobs_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetBobs",
			Handler:       _Bobs_GetBobs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetPeopleFromNames",
			Handler:       _Bobs_GetPeopleFromNames_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "tests/spanner/bob_example/bobs.proto",
}

func init() { proto.RegisterFile("tests/spanner/bob_example/bobs.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 550 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x5d, 0x6b, 0xd4, 0x40,
	0x14, 0x25, 0xd9, 0xb6, 0x6b, 0x6f, 0x45, 0x24, 0x28, 0xae, 0x01, 0xed, 0x10, 0x84, 0x66, 0x4b,
	0x9b, 0x94, 0x56, 0x04, 0xe9, 0x4b, 0xf7, 0x63, 0x5a, 0x57, 0x77, 0xb3, 0x25, 0x9b, 0x2a, 0x28,
	0x52, 0x32, 0xbb, 0xd3, 0x35, 0x90, 0x64, 0x42, 0x66, 0x82, 0xee, 0xab, 0x4f, 0xe2, 0x93, 0xfb,
	0xea, 0x8b, 0x7f, 0x63, 0x7f, 0x9e, 0xcc, 0x64, 0xb7, 0x0d, 0x45, 0x04, 0xc1, 0xb7, 0x73, 0xb8,
	0x77, 0xce, 0xb9, 0xe7, 0x72, 0x07, 0x9e, 0x09, 0xca, 0x05, 0x77, 0x79, 0x16, 0xa6, 0x29, 0xcd,
	0x5d, 0xc2, 0xc8, 0x25, 0xfd, 0x12, 0x26, 0x59, 0x4c, 0x25, 0xe6, 0x4e, 0x96, 0x33, 0xc1, 0x8c,
	0x1a, 0x61, 0xc4, 0x7c, 0x98, 0xd1, 0x9c, 0x47, 0x5c, 0xb8, 0x2c, 0x13, 0x11, 0x4b, 0x97, 0x35,
	0x73, 0x7b, 0xca, 0xd8, 0x34, 0xa6, 0xae, 0x62, 0xa4, 0xb8, 0x72, 0x45, 0x94, 0x50, 0x2e, 0xc2,
	0x24, 0x2b, 0x1b, 0xac, 0x09, 0xd4, 0xda, 0x8c, 0x18, 0xf7, 0x40, 0x8f, 0x26, 0x0d, 0x0d, 0x69,
	0x76, 0xcd, 0xd7, 0xa3, 0x89, 0xf1, 0x12, 0x80, 0x8b, 0x30, 0x17, 0x97, 0xb2, 0xbf, 0xa1, 0x23,
	0xcd, 0xde, 0x3a, 0x34, 0x9d, 0x52, 0xcc, 0x59, 0x89, 0x39, 0xc1, 0x4a, 0xcc, 0xdf, 0x54, 0xdd,
	0x92, 0x1b, 0x06, 0xac, 0xa5, 0x61, 0x42, 0x1b, 0x35, 0xa4, 0xd9, 0x9b, 0xbe, 0xc2, 0xd6, 0x36,
	0xd4, 0xbd, 0x22, 0xf1, 0xd9, 0x67, 0x6e, 0x3c, 0x80, 0xf5, 0x31, 0x2b, 0x52, 0xb1, 0x34, 0x2b,
	0x89, 0xf5, 0x04, 0xd6, 0xbd, 0x30, 0xa1, 0xaa, 0x2c, 0x5f, 0xf0, 0x86, 0x86, 0x6a, 0xf6, 0xa6,
	0x5f, 0x12, 0xab, 0x0e, 0xeb, 0x38, 0xc9, 0xc4, 0xec, 0xf0, 0xdb, 0x1a, 0xac, 0xb5, 0x19, 0xe1,
	0x46, 0x0a, 0xd0, 0xa5, 0x31, 0x15, 0x54, 0xb1, 0x3b, 0x0e, 0x61, 0xc4, 0x69, 0x33, 0x62, 0x82,
	0x42, 0xaa, 0xd9, 0x7a, 0xfd, 0x75, 0x31, 0xd7, 0x31, 0x74, 0xba, 0xb8, 0x8f, 0x03, 0x8c, 0x4e,
	0xfd, 0xe1, 0x00, 0xc9, 0x1d, 0x8a, 0x90, 0xc4, 0x14, 0x8d, 0x82, 0x96, 0x1f, 0xd8, 0x3b, 0x6d,
	0x46, 0x76, 0x9a, 0x08, 0x7b, 0xdd, 0x12, 0xee, 0xa1, 0x93, 0x9b, 0xe0, 0x4d, 0xf4, 0xa6, 0xe7,
	0x75, 0xed, 0xce, 0xb0, 0x69, 0xcc, 0xa0, 0x7e, 0x5e, 0x88, 0x5b, 0x66, 0x77, 0x15, 0x5a, 0x26,
	0xb3, 0x02, 0x69, 0x37, 0x84, 0x41, 0xcf, 0x1b, 0x61, 0x3f, 0x40, 0x3d, 0x2f, 0x18, 0x56, 0xec,
	0xec, 0x68, 0xb2, 0x87, 0x64, 0x9a, 0x3d, 0x54, 0x35, 0x78, 0xdb, 0xea, 0x5f, 0xe0, 0x11, 0xb2,
	0x4f, 0x64, 0xf9, 0xa4, 0xac, 0x57, 0x27, 0xb0, 0x35, 0xa3, 0x05, 0xf5, 0x33, 0x5a, 0x5a, 0x57,
	0xd2, 0x99, 0xd7, 0x63, 0x58, 0x4f, 0xa5, 0xf1, 0x63, 0x78, 0x34, 0xc2, 0x7d, 0xdc, 0x09, 0xd0,
	0x2e, 0xba, 0xca, 0x59, 0x72, 0x63, 0x7d, 0xa0, 0x19, 0x11, 0x18, 0x67, 0x54, 0x9c, 0x53, 0x96,
	0xc5, 0xf4, 0x34, 0x67, 0x49, 0xb9, 0xeb, 0x52, 0x4d, 0xe1, 0x8a, 0xda, 0xb1, 0x54, 0x7b, 0x01,
	0xcf, 0xaf, 0xd5, 0x6e, 0xed, 0xed, 0xdd, 0x2b, 0xec, 0x63, 0x95, 0x04, 0xf5, 0x3c, 0x74, 0xe1,
	0x79, 0x78, 0x14, 0xd8, 0x6a, 0x72, 0xde, 0x3c, 0xd0, 0xcc, 0xe9, 0xf7, 0xc5, 0x5c, 0xff, 0x08,
	0x1f, 0xe0, 0x2f, 0xe7, 0x72, 0x7f, 0xcb, 0xda, 0x18, 0xcc, 0x24, 0xdd, 0x3d, 0x9a, 0x46, 0xe2,
	0x53, 0x41, 0x9c, 0x31, 0x4b, 0x5c, 0x31, 0x4e, 0xc7, 0x31, 0x2b, 0x26, 0xe5, 0xb5, 0x8e, 0xf7,
	0xa7, 0x34, 0xdd, 0x5f, 0x9d, 0x75, 0xf9, 0x13, 0x92, 0x99, 0x5c, 0xca, 0x8f, 0xc5, 0x5c, 0xd7,
	0xda, 0xbf, 0xb4, 0x9f, 0x8b, 0xb9, 0xde, 0xff, 0xa7, 0xf7, 0x7f, 0xf8, 0x49, 0xc7, 0x15, 0xfc,
	0xfe, 0xbf, 0xaa, 0x91, 0x0d, 0xf5, 0xf2, 0xe8, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x2d,
	0xf9, 0x75, 0xc7, 0x03, 0x00, 0x00,
}
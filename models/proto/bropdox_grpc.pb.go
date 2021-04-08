// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BropdoxClient is the client API for Bropdox service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BropdoxClient interface {
	CreateFile(ctx context.Context, in *File, opts ...grpc.CallOption) (*FileResponse, error)
	UpdateFile(ctx context.Context, in *File, opts ...grpc.CallOption) (*FileResponse, error)
	RemoveFile(ctx context.Context, in *File, opts ...grpc.CallOption) (*FileResponse, error)
	Notifications(ctx context.Context, in *File, opts ...grpc.CallOption) (Bropdox_NotificationsClient, error)
}

type bropdoxClient struct {
	cc grpc.ClientConnInterface
}

func NewBropdoxClient(cc grpc.ClientConnInterface) BropdoxClient {
	return &bropdoxClient{cc}
}

func (c *bropdoxClient) CreateFile(ctx context.Context, in *File, opts ...grpc.CallOption) (*FileResponse, error) {
	out := new(FileResponse)
	err := c.cc.Invoke(ctx, "/proto.Bropdox/CreateFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bropdoxClient) UpdateFile(ctx context.Context, in *File, opts ...grpc.CallOption) (*FileResponse, error) {
	out := new(FileResponse)
	err := c.cc.Invoke(ctx, "/proto.Bropdox/UpdateFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bropdoxClient) RemoveFile(ctx context.Context, in *File, opts ...grpc.CallOption) (*FileResponse, error) {
	out := new(FileResponse)
	err := c.cc.Invoke(ctx, "/proto.Bropdox/RemoveFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bropdoxClient) Notifications(ctx context.Context, in *File, opts ...grpc.CallOption) (Bropdox_NotificationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Bropdox_ServiceDesc.Streams[0], "/proto.Bropdox/Notifications", opts...)
	if err != nil {
		return nil, err
	}
	x := &bropdoxNotificationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Bropdox_NotificationsClient interface {
	Recv() (*FileResponse, error)
	grpc.ClientStream
}

type bropdoxNotificationsClient struct {
	grpc.ClientStream
}

func (x *bropdoxNotificationsClient) Recv() (*FileResponse, error) {
	m := new(FileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BropdoxServer is the server API for Bropdox service.
// All implementations should embed UnimplementedBropdoxServer
// for forward compatibility
type BropdoxServer interface {
	CreateFile(context.Context, *File) (*FileResponse, error)
	UpdateFile(context.Context, *File) (*FileResponse, error)
	RemoveFile(context.Context, *File) (*FileResponse, error)
	Notifications(*File, Bropdox_NotificationsServer) error
}

// UnimplementedBropdoxServer should be embedded to have forward compatible implementations.
type UnimplementedBropdoxServer struct {
}

func (UnimplementedBropdoxServer) CreateFile(context.Context, *File) (*FileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFile not implemented")
}
func (UnimplementedBropdoxServer) UpdateFile(context.Context, *File) (*FileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFile not implemented")
}
func (UnimplementedBropdoxServer) RemoveFile(context.Context, *File) (*FileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFile not implemented")
}
func (UnimplementedBropdoxServer) Notifications(*File, Bropdox_NotificationsServer) error {
	return status.Errorf(codes.Unimplemented, "method Notifications not implemented")
}

// UnsafeBropdoxServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BropdoxServer will
// result in compilation errors.
type UnsafeBropdoxServer interface {
	mustEmbedUnimplementedBropdoxServer()
}

func RegisterBropdoxServer(s grpc.ServiceRegistrar, srv BropdoxServer) {
	s.RegisterService(&Bropdox_ServiceDesc, srv)
}

func _Bropdox_CreateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(File)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BropdoxServer).CreateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Bropdox/CreateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BropdoxServer).CreateFile(ctx, req.(*File))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bropdox_UpdateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(File)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BropdoxServer).UpdateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Bropdox/UpdateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BropdoxServer).UpdateFile(ctx, req.(*File))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bropdox_RemoveFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(File)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BropdoxServer).RemoveFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Bropdox/RemoveFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BropdoxServer).RemoveFile(ctx, req.(*File))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bropdox_Notifications_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(File)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BropdoxServer).Notifications(m, &bropdoxNotificationsServer{stream})
}

type Bropdox_NotificationsServer interface {
	Send(*FileResponse) error
	grpc.ServerStream
}

type bropdoxNotificationsServer struct {
	grpc.ServerStream
}

func (x *bropdoxNotificationsServer) Send(m *FileResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Bropdox_ServiceDesc is the grpc.ServiceDesc for Bropdox service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bropdox_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Bropdox",
	HandlerType: (*BropdoxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFile",
			Handler:    _Bropdox_CreateFile_Handler,
		},
		{
			MethodName: "UpdateFile",
			Handler:    _Bropdox_UpdateFile_Handler,
		},
		{
			MethodName: "RemoveFile",
			Handler:    _Bropdox_RemoveFile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Notifications",
			Handler:       _Bropdox_Notifications_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "bropdox.proto",
}

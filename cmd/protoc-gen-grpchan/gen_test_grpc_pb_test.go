// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package main

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

// TestStreamsClient is the client API for TestStreams service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestStreamsClient interface {
	Unary1(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error)
	Unary2(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error)
	Stream1(ctx context.Context, opts ...grpc.CallOption) (TestStreams_Stream1Client, error)
	Unary3(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error)
	Stream2(ctx context.Context, opts ...grpc.CallOption) (TestStreams_Stream2Client, error)
	Unary4(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error)
	Unary5(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error)
	Stream3(ctx context.Context, in *Test, opts ...grpc.CallOption) (TestStreams_Stream3Client, error)
	Unary6(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error)
	Unary7(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error)
	Unary8(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error)
	Stream4(ctx context.Context, opts ...grpc.CallOption) (TestStreams_Stream4Client, error)
	Stream5(ctx context.Context, in *Test, opts ...grpc.CallOption) (TestStreams_Stream5Client, error)
	Unary9(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error)
	Unary10(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error)
	Unary11(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error)
	Stream6(ctx context.Context, opts ...grpc.CallOption) (TestStreams_Stream6Client, error)
	Stream7(ctx context.Context, opts ...grpc.CallOption) (TestStreams_Stream7Client, error)
	Unary12(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error)
}

type testStreamsClient struct {
	cc grpc.ClientConnInterface
}

func NewTestStreamsClient(cc grpc.ClientConnInterface) TestStreamsClient {
	return &testStreamsClient{cc}
}

func (c *testStreamsClient) Unary1(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error) {
	out := new(Test)
	err := c.cc.Invoke(ctx, "/main.TestStreams/Unary1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testStreamsClient) Unary2(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error) {
	out := new(Test)
	err := c.cc.Invoke(ctx, "/main.TestStreams/Unary2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testStreamsClient) Stream1(ctx context.Context, opts ...grpc.CallOption) (TestStreams_Stream1Client, error) {
	stream, err := c.cc.NewStream(ctx, &TestStreams_ServiceDesc.Streams[0], "/main.TestStreams/Stream1", opts...)
	if err != nil {
		return nil, err
	}
	x := &testStreamsStream1Client{stream}
	return x, nil
}

type TestStreams_Stream1Client interface {
	Send(*Test) error
	CloseAndRecv() (*Test, error)
	grpc.ClientStream
}

type testStreamsStream1Client struct {
	grpc.ClientStream
}

func (x *testStreamsStream1Client) Send(m *Test) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testStreamsStream1Client) CloseAndRecv() (*Test, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Test)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testStreamsClient) Unary3(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error) {
	out := new(Test)
	err := c.cc.Invoke(ctx, "/main.TestStreams/Unary3", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testStreamsClient) Stream2(ctx context.Context, opts ...grpc.CallOption) (TestStreams_Stream2Client, error) {
	stream, err := c.cc.NewStream(ctx, &TestStreams_ServiceDesc.Streams[1], "/main.TestStreams/Stream2", opts...)
	if err != nil {
		return nil, err
	}
	x := &testStreamsStream2Client{stream}
	return x, nil
}

type TestStreams_Stream2Client interface {
	Send(*Test) error
	Recv() (*Test, error)
	grpc.ClientStream
}

type testStreamsStream2Client struct {
	grpc.ClientStream
}

func (x *testStreamsStream2Client) Send(m *Test) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testStreamsStream2Client) Recv() (*Test, error) {
	m := new(Test)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testStreamsClient) Unary4(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error) {
	out := new(Test)
	err := c.cc.Invoke(ctx, "/main.TestStreams/Unary4", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testStreamsClient) Unary5(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error) {
	out := new(Test)
	err := c.cc.Invoke(ctx, "/main.TestStreams/Unary5", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testStreamsClient) Stream3(ctx context.Context, in *Test, opts ...grpc.CallOption) (TestStreams_Stream3Client, error) {
	stream, err := c.cc.NewStream(ctx, &TestStreams_ServiceDesc.Streams[2], "/main.TestStreams/Stream3", opts...)
	if err != nil {
		return nil, err
	}
	x := &testStreamsStream3Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestStreams_Stream3Client interface {
	Recv() (*Test, error)
	grpc.ClientStream
}

type testStreamsStream3Client struct {
	grpc.ClientStream
}

func (x *testStreamsStream3Client) Recv() (*Test, error) {
	m := new(Test)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testStreamsClient) Unary6(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error) {
	out := new(Test)
	err := c.cc.Invoke(ctx, "/main.TestStreams/Unary6", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testStreamsClient) Unary7(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error) {
	out := new(Test)
	err := c.cc.Invoke(ctx, "/main.TestStreams/Unary7", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testStreamsClient) Unary8(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error) {
	out := new(Test)
	err := c.cc.Invoke(ctx, "/main.TestStreams/Unary8", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testStreamsClient) Stream4(ctx context.Context, opts ...grpc.CallOption) (TestStreams_Stream4Client, error) {
	stream, err := c.cc.NewStream(ctx, &TestStreams_ServiceDesc.Streams[3], "/main.TestStreams/Stream4", opts...)
	if err != nil {
		return nil, err
	}
	x := &testStreamsStream4Client{stream}
	return x, nil
}

type TestStreams_Stream4Client interface {
	Send(*Test) error
	CloseAndRecv() (*Test, error)
	grpc.ClientStream
}

type testStreamsStream4Client struct {
	grpc.ClientStream
}

func (x *testStreamsStream4Client) Send(m *Test) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testStreamsStream4Client) CloseAndRecv() (*Test, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Test)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testStreamsClient) Stream5(ctx context.Context, in *Test, opts ...grpc.CallOption) (TestStreams_Stream5Client, error) {
	stream, err := c.cc.NewStream(ctx, &TestStreams_ServiceDesc.Streams[4], "/main.TestStreams/Stream5", opts...)
	if err != nil {
		return nil, err
	}
	x := &testStreamsStream5Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestStreams_Stream5Client interface {
	Recv() (*Test, error)
	grpc.ClientStream
}

type testStreamsStream5Client struct {
	grpc.ClientStream
}

func (x *testStreamsStream5Client) Recv() (*Test, error) {
	m := new(Test)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testStreamsClient) Unary9(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error) {
	out := new(Test)
	err := c.cc.Invoke(ctx, "/main.TestStreams/Unary9", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testStreamsClient) Unary10(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error) {
	out := new(Test)
	err := c.cc.Invoke(ctx, "/main.TestStreams/Unary10", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testStreamsClient) Unary11(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error) {
	out := new(Test)
	err := c.cc.Invoke(ctx, "/main.TestStreams/Unary11", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testStreamsClient) Stream6(ctx context.Context, opts ...grpc.CallOption) (TestStreams_Stream6Client, error) {
	stream, err := c.cc.NewStream(ctx, &TestStreams_ServiceDesc.Streams[5], "/main.TestStreams/Stream6", opts...)
	if err != nil {
		return nil, err
	}
	x := &testStreamsStream6Client{stream}
	return x, nil
}

type TestStreams_Stream6Client interface {
	Send(*Test) error
	Recv() (*Test, error)
	grpc.ClientStream
}

type testStreamsStream6Client struct {
	grpc.ClientStream
}

func (x *testStreamsStream6Client) Send(m *Test) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testStreamsStream6Client) Recv() (*Test, error) {
	m := new(Test)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testStreamsClient) Stream7(ctx context.Context, opts ...grpc.CallOption) (TestStreams_Stream7Client, error) {
	stream, err := c.cc.NewStream(ctx, &TestStreams_ServiceDesc.Streams[6], "/main.TestStreams/Stream7", opts...)
	if err != nil {
		return nil, err
	}
	x := &testStreamsStream7Client{stream}
	return x, nil
}

type TestStreams_Stream7Client interface {
	Send(*Test) error
	Recv() (*Test, error)
	grpc.ClientStream
}

type testStreamsStream7Client struct {
	grpc.ClientStream
}

func (x *testStreamsStream7Client) Send(m *Test) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testStreamsStream7Client) Recv() (*Test, error) {
	m := new(Test)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testStreamsClient) Unary12(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error) {
	out := new(Test)
	err := c.cc.Invoke(ctx, "/main.TestStreams/Unary12", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestStreamsServer is the server API for TestStreams service.
// All implementations must embed UnimplementedTestStreamsServer
// for forward compatibility
type TestStreamsServer interface {
	Unary1(context.Context, *Test) (*Test, error)
	Unary2(context.Context, *Test) (*Test, error)
	Stream1(TestStreams_Stream1Server) error
	Unary3(context.Context, *Test) (*Test, error)
	Stream2(TestStreams_Stream2Server) error
	Unary4(context.Context, *Test) (*Test, error)
	Unary5(context.Context, *Test) (*Test, error)
	Stream3(*Test, TestStreams_Stream3Server) error
	Unary6(context.Context, *Test) (*Test, error)
	Unary7(context.Context, *Test) (*Test, error)
	Unary8(context.Context, *Test) (*Test, error)
	Stream4(TestStreams_Stream4Server) error
	Stream5(*Test, TestStreams_Stream5Server) error
	Unary9(context.Context, *Test) (*Test, error)
	Unary10(context.Context, *Test) (*Test, error)
	Unary11(context.Context, *Test) (*Test, error)
	Stream6(TestStreams_Stream6Server) error
	Stream7(TestStreams_Stream7Server) error
	Unary12(context.Context, *Test) (*Test, error)
	mustEmbedUnimplementedTestStreamsServer()
}

// UnimplementedTestStreamsServer must be embedded to have forward compatible implementations.
type UnimplementedTestStreamsServer struct {
}

func (UnimplementedTestStreamsServer) Unary1(context.Context, *Test) (*Test, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unary1 not implemented")
}
func (UnimplementedTestStreamsServer) Unary2(context.Context, *Test) (*Test, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unary2 not implemented")
}
func (UnimplementedTestStreamsServer) Stream1(TestStreams_Stream1Server) error {
	return status.Errorf(codes.Unimplemented, "method Stream1 not implemented")
}
func (UnimplementedTestStreamsServer) Unary3(context.Context, *Test) (*Test, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unary3 not implemented")
}
func (UnimplementedTestStreamsServer) Stream2(TestStreams_Stream2Server) error {
	return status.Errorf(codes.Unimplemented, "method Stream2 not implemented")
}
func (UnimplementedTestStreamsServer) Unary4(context.Context, *Test) (*Test, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unary4 not implemented")
}
func (UnimplementedTestStreamsServer) Unary5(context.Context, *Test) (*Test, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unary5 not implemented")
}
func (UnimplementedTestStreamsServer) Stream3(*Test, TestStreams_Stream3Server) error {
	return status.Errorf(codes.Unimplemented, "method Stream3 not implemented")
}
func (UnimplementedTestStreamsServer) Unary6(context.Context, *Test) (*Test, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unary6 not implemented")
}
func (UnimplementedTestStreamsServer) Unary7(context.Context, *Test) (*Test, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unary7 not implemented")
}
func (UnimplementedTestStreamsServer) Unary8(context.Context, *Test) (*Test, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unary8 not implemented")
}
func (UnimplementedTestStreamsServer) Stream4(TestStreams_Stream4Server) error {
	return status.Errorf(codes.Unimplemented, "method Stream4 not implemented")
}
func (UnimplementedTestStreamsServer) Stream5(*Test, TestStreams_Stream5Server) error {
	return status.Errorf(codes.Unimplemented, "method Stream5 not implemented")
}
func (UnimplementedTestStreamsServer) Unary9(context.Context, *Test) (*Test, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unary9 not implemented")
}
func (UnimplementedTestStreamsServer) Unary10(context.Context, *Test) (*Test, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unary10 not implemented")
}
func (UnimplementedTestStreamsServer) Unary11(context.Context, *Test) (*Test, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unary11 not implemented")
}
func (UnimplementedTestStreamsServer) Stream6(TestStreams_Stream6Server) error {
	return status.Errorf(codes.Unimplemented, "method Stream6 not implemented")
}
func (UnimplementedTestStreamsServer) Stream7(TestStreams_Stream7Server) error {
	return status.Errorf(codes.Unimplemented, "method Stream7 not implemented")
}
func (UnimplementedTestStreamsServer) Unary12(context.Context, *Test) (*Test, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unary12 not implemented")
}
func (UnimplementedTestStreamsServer) mustEmbedUnimplementedTestStreamsServer() {}

// UnsafeTestStreamsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestStreamsServer will
// result in compilation errors.
type UnsafeTestStreamsServer interface {
	mustEmbedUnimplementedTestStreamsServer()
}

func RegisterTestStreamsServer(s grpc.ServiceRegistrar, srv TestStreamsServer) {
	s.RegisterService(&TestStreams_ServiceDesc, srv)
}

func _TestStreams_Unary1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Test)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestStreamsServer).Unary1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TestStreams/Unary1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestStreamsServer).Unary1(ctx, req.(*Test))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestStreams_Unary2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Test)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestStreamsServer).Unary2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TestStreams/Unary2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestStreamsServer).Unary2(ctx, req.(*Test))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestStreams_Stream1_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestStreamsServer).Stream1(&testStreamsStream1Server{stream})
}

type TestStreams_Stream1Server interface {
	SendAndClose(*Test) error
	Recv() (*Test, error)
	grpc.ServerStream
}

type testStreamsStream1Server struct {
	grpc.ServerStream
}

func (x *testStreamsStream1Server) SendAndClose(m *Test) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testStreamsStream1Server) Recv() (*Test, error) {
	m := new(Test)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestStreams_Unary3_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Test)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestStreamsServer).Unary3(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TestStreams/Unary3",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestStreamsServer).Unary3(ctx, req.(*Test))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestStreams_Stream2_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestStreamsServer).Stream2(&testStreamsStream2Server{stream})
}

type TestStreams_Stream2Server interface {
	Send(*Test) error
	Recv() (*Test, error)
	grpc.ServerStream
}

type testStreamsStream2Server struct {
	grpc.ServerStream
}

func (x *testStreamsStream2Server) Send(m *Test) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testStreamsStream2Server) Recv() (*Test, error) {
	m := new(Test)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestStreams_Unary4_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Test)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestStreamsServer).Unary4(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TestStreams/Unary4",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestStreamsServer).Unary4(ctx, req.(*Test))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestStreams_Unary5_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Test)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestStreamsServer).Unary5(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TestStreams/Unary5",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestStreamsServer).Unary5(ctx, req.(*Test))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestStreams_Stream3_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Test)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestStreamsServer).Stream3(m, &testStreamsStream3Server{stream})
}

type TestStreams_Stream3Server interface {
	Send(*Test) error
	grpc.ServerStream
}

type testStreamsStream3Server struct {
	grpc.ServerStream
}

func (x *testStreamsStream3Server) Send(m *Test) error {
	return x.ServerStream.SendMsg(m)
}

func _TestStreams_Unary6_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Test)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestStreamsServer).Unary6(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TestStreams/Unary6",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestStreamsServer).Unary6(ctx, req.(*Test))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestStreams_Unary7_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Test)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestStreamsServer).Unary7(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TestStreams/Unary7",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestStreamsServer).Unary7(ctx, req.(*Test))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestStreams_Unary8_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Test)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestStreamsServer).Unary8(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TestStreams/Unary8",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestStreamsServer).Unary8(ctx, req.(*Test))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestStreams_Stream4_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestStreamsServer).Stream4(&testStreamsStream4Server{stream})
}

type TestStreams_Stream4Server interface {
	SendAndClose(*Test) error
	Recv() (*Test, error)
	grpc.ServerStream
}

type testStreamsStream4Server struct {
	grpc.ServerStream
}

func (x *testStreamsStream4Server) SendAndClose(m *Test) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testStreamsStream4Server) Recv() (*Test, error) {
	m := new(Test)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestStreams_Stream5_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Test)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestStreamsServer).Stream5(m, &testStreamsStream5Server{stream})
}

type TestStreams_Stream5Server interface {
	Send(*Test) error
	grpc.ServerStream
}

type testStreamsStream5Server struct {
	grpc.ServerStream
}

func (x *testStreamsStream5Server) Send(m *Test) error {
	return x.ServerStream.SendMsg(m)
}

func _TestStreams_Unary9_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Test)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestStreamsServer).Unary9(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TestStreams/Unary9",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestStreamsServer).Unary9(ctx, req.(*Test))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestStreams_Unary10_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Test)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestStreamsServer).Unary10(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TestStreams/Unary10",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestStreamsServer).Unary10(ctx, req.(*Test))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestStreams_Unary11_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Test)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestStreamsServer).Unary11(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TestStreams/Unary11",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestStreamsServer).Unary11(ctx, req.(*Test))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestStreams_Stream6_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestStreamsServer).Stream6(&testStreamsStream6Server{stream})
}

type TestStreams_Stream6Server interface {
	Send(*Test) error
	Recv() (*Test, error)
	grpc.ServerStream
}

type testStreamsStream6Server struct {
	grpc.ServerStream
}

func (x *testStreamsStream6Server) Send(m *Test) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testStreamsStream6Server) Recv() (*Test, error) {
	m := new(Test)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestStreams_Stream7_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestStreamsServer).Stream7(&testStreamsStream7Server{stream})
}

type TestStreams_Stream7Server interface {
	Send(*Test) error
	Recv() (*Test, error)
	grpc.ServerStream
}

type testStreamsStream7Server struct {
	grpc.ServerStream
}

func (x *testStreamsStream7Server) Send(m *Test) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testStreamsStream7Server) Recv() (*Test, error) {
	m := new(Test)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestStreams_Unary12_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Test)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestStreamsServer).Unary12(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.TestStreams/Unary12",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestStreamsServer).Unary12(ctx, req.(*Test))
	}
	return interceptor(ctx, in, info, handler)
}

// TestStreams_ServiceDesc is the grpc.ServiceDesc for TestStreams service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestStreams_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.TestStreams",
	HandlerType: (*TestStreamsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unary1",
			Handler:    _TestStreams_Unary1_Handler,
		},
		{
			MethodName: "Unary2",
			Handler:    _TestStreams_Unary2_Handler,
		},
		{
			MethodName: "Unary3",
			Handler:    _TestStreams_Unary3_Handler,
		},
		{
			MethodName: "Unary4",
			Handler:    _TestStreams_Unary4_Handler,
		},
		{
			MethodName: "Unary5",
			Handler:    _TestStreams_Unary5_Handler,
		},
		{
			MethodName: "Unary6",
			Handler:    _TestStreams_Unary6_Handler,
		},
		{
			MethodName: "Unary7",
			Handler:    _TestStreams_Unary7_Handler,
		},
		{
			MethodName: "Unary8",
			Handler:    _TestStreams_Unary8_Handler,
		},
		{
			MethodName: "Unary9",
			Handler:    _TestStreams_Unary9_Handler,
		},
		{
			MethodName: "Unary10",
			Handler:    _TestStreams_Unary10_Handler,
		},
		{
			MethodName: "Unary11",
			Handler:    _TestStreams_Unary11_Handler,
		},
		{
			MethodName: "Unary12",
			Handler:    _TestStreams_Unary12_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream1",
			Handler:       _TestStreams_Stream1_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Stream2",
			Handler:       _TestStreams_Stream2_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Stream3",
			Handler:       _TestStreams_Stream3_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Stream4",
			Handler:       _TestStreams_Stream4_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Stream5",
			Handler:       _TestStreams_Stream5_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Stream6",
			Handler:       _TestStreams_Stream6_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Stream7",
			Handler:       _TestStreams_Stream7_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "gen_test.proto",
}

// Code generated by protoc-gen-grpchan. DO NOT EDIT.
// source: test.proto

package grpchantesting

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/solarwinds/grpchan"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func RegisterHandlerTestService(reg grpchan.ServiceRegistry, srv TestServiceServer) {
	reg.RegisterService(&_TestService_serviceDesc, srv)
}

type testServiceChannelClient struct {
	ch grpchan.Channel
}

func NewTestServiceChannelClient(ch grpchan.Channel) TestServiceClient {
	return &testServiceChannelClient{ch: ch}
}

func (c *testServiceChannelClient) Unary(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.ch.Invoke(ctx, "/grpchantesting.TestService/Unary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceChannelClient) ClientStream(ctx context.Context, opts ...grpc.CallOption) (TestService_ClientStreamClient, error) {
	stream, err := c.ch.NewStream(ctx, &_TestService_serviceDesc.Streams[0], "/grpchantesting.TestService/ClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceClientStreamClient{stream}
	return x, nil
}

func (c *testServiceChannelClient) ServerStream(ctx context.Context, in *Message, opts ...grpc.CallOption) (TestService_ServerStreamClient, error) {
	stream, err := c.ch.NewStream(ctx, &_TestService_serviceDesc.Streams[1], "/grpchantesting.TestService/ServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

func (c *testServiceChannelClient) BidiStream(ctx context.Context, opts ...grpc.CallOption) (TestService_BidiStreamClient, error) {
	stream, err := c.ch.NewStream(ctx, &_TestService_serviceDesc.Streams[2], "/grpchantesting.TestService/BidiStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceBidiStreamClient{stream}
	return x, nil
}

func (c *testServiceChannelClient) UseExternalMessageTwice(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.ch.Invoke(ctx, "/grpchantesting.TestService/UseExternalMessageTwice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

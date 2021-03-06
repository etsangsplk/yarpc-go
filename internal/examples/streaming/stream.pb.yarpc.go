// Code generated by protoc-gen-yarpc-go
// source: internal/examples/streaming/stream.proto
// DO NOT EDIT!

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package streaming

import (
	"context"
	"io/ioutil"
	"reflect"

	"github.com/gogo/protobuf/proto"
	"go.uber.org/fx"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/encoding/protobuf"
)

var _ = ioutil.NopCloser

// HelloYARPCClient is the YARPC client-side interface for the Hello service.
type HelloYARPCClient interface {
	HelloUnary(context.Context, *HelloRequest, ...yarpc.CallOption) (*HelloResponse, error)
	HelloOutStream(context.Context, ...yarpc.CallOption) (HelloServiceHelloOutStreamYARPCClient, error)
	HelloInStream(context.Context, *HelloRequest, ...yarpc.CallOption) (HelloServiceHelloInStreamYARPCClient, error)
	HelloThere(context.Context, ...yarpc.CallOption) (HelloServiceHelloThereYARPCClient, error)
}

// HelloServiceHelloOutStreamYARPCClient sends HelloRequests and receives the single HelloResponse when sending is done.
type HelloServiceHelloOutStreamYARPCClient interface {
	Context() context.Context
	Send(*HelloRequest, ...yarpc.StreamOption) error
	CloseAndRecv(...yarpc.StreamOption) (*HelloResponse, error)
}

// HelloServiceHelloInStreamYARPCClient receives HelloResponses, returning io.EOF when the stream is complete.
type HelloServiceHelloInStreamYARPCClient interface {
	Context() context.Context
	Recv(...yarpc.StreamOption) (*HelloResponse, error)
	CloseSend(...yarpc.StreamOption) error
}

// HelloServiceHelloThereYARPCClient sends HelloRequests and receives HelloResponses, returning io.EOF when the stream is complete.
type HelloServiceHelloThereYARPCClient interface {
	Context() context.Context
	Send(*HelloRequest, ...yarpc.StreamOption) error
	Recv(...yarpc.StreamOption) (*HelloResponse, error)
	CloseSend(...yarpc.StreamOption) error
}

// NewHelloYARPCClient builds a new YARPC client for the Hello service.
func NewHelloYARPCClient(clientConfig transport.ClientConfig, options ...protobuf.ClientOption) HelloYARPCClient {
	return &_HelloYARPCCaller{protobuf.NewStreamClient(
		protobuf.ClientParams{
			ServiceName:  "uber.yarpc.internal.examples.streaming.Hello",
			ClientConfig: clientConfig,
			Options:      options,
		},
	)}
}

// HelloYARPCServer is the YARPC server-side interface for the Hello service.
type HelloYARPCServer interface {
	HelloUnary(context.Context, *HelloRequest) (*HelloResponse, error)
	HelloOutStream(HelloServiceHelloOutStreamYARPCServer) (*HelloResponse, error)
	HelloInStream(*HelloRequest, HelloServiceHelloInStreamYARPCServer) error
	HelloThere(HelloServiceHelloThereYARPCServer) error
}

// HelloServiceHelloOutStreamYARPCServer receives HelloRequests.
type HelloServiceHelloOutStreamYARPCServer interface {
	Context() context.Context
	Recv(...yarpc.StreamOption) (*HelloRequest, error)
}

// HelloServiceHelloInStreamYARPCServer sends HelloResponses.
type HelloServiceHelloInStreamYARPCServer interface {
	Context() context.Context
	Send(*HelloResponse, ...yarpc.StreamOption) error
}

// HelloServiceHelloThereYARPCServer receives HelloRequests and sends HelloResponse.
type HelloServiceHelloThereYARPCServer interface {
	Context() context.Context
	Recv(...yarpc.StreamOption) (*HelloRequest, error)
	Send(*HelloResponse, ...yarpc.StreamOption) error
}

// BuildHelloYARPCProcedures prepares an implementation of the Hello service for YARPC registration.
func BuildHelloYARPCProcedures(server HelloYARPCServer) []transport.Procedure {
	handler := &_HelloYARPCHandler{server}
	return protobuf.BuildProcedures(
		protobuf.BuildProceduresParams{
			ServiceName: "uber.yarpc.internal.examples.streaming.Hello",
			UnaryHandlerParams: []protobuf.BuildProceduresUnaryHandlerParams{
				{
					MethodName: "HelloUnary",
					Handler: protobuf.NewUnaryHandler(
						protobuf.UnaryHandlerParams{
							Handle:     handler.HelloUnary,
							NewRequest: newHelloServiceHelloUnaryYARPCRequest,
						},
					),
				},
			},
			OnewayHandlerParams: []protobuf.BuildProceduresOnewayHandlerParams{},
			StreamHandlerParams: []protobuf.BuildProceduresStreamHandlerParams{
				{
					MethodName: "HelloThere",
					Handler: protobuf.NewStreamHandler(
						protobuf.StreamHandlerParams{
							Handle: handler.HelloThere,
						},
					),
				},

				{
					MethodName: "HelloInStream",
					Handler: protobuf.NewStreamHandler(
						protobuf.StreamHandlerParams{
							Handle: handler.HelloInStream,
						},
					),
				},

				{
					MethodName: "HelloOutStream",
					Handler: protobuf.NewStreamHandler(
						protobuf.StreamHandlerParams{
							Handle: handler.HelloOutStream,
						},
					),
				},
			},
		},
	)
}

// FxHelloYARPCClientParams defines the input
// for NewFxHelloYARPCClient. It provides the
// paramaters to get a HelloYARPCClient in an
// Fx application.
type FxHelloYARPCClientParams struct {
	fx.In

	Provider yarpc.ClientConfig
}

// FxHelloYARPCClientResult defines the output
// of NewFxHelloYARPCClient. It provides a
// HelloYARPCClient to an Fx application.
type FxHelloYARPCClientResult struct {
	fx.Out

	Client HelloYARPCClient

	// We are using an fx.Out struct here instead of just returning a client
	// so that we can add more values or add named versions of the client in
	// the future without breaking any existing code.
}

// NewFxHelloYARPCClient provides a HelloYARPCClient
// to an Fx application using the given name for routing.
//
//  fx.Provide(
//    streaming.NewFxHelloYARPCClient("service-name"),
//    ...
//  )
func NewFxHelloYARPCClient(name string, options ...protobuf.ClientOption) interface{} {
	return func(params FxHelloYARPCClientParams) FxHelloYARPCClientResult {
		return FxHelloYARPCClientResult{
			Client: NewHelloYARPCClient(params.Provider.ClientConfig(name), options...),
		}
	}
}

// FxHelloYARPCProceduresParams defines the input
// for NewFxHelloYARPCProcedures. It provides the
// paramaters to get HelloYARPCServer procedures in an
// Fx application.
type FxHelloYARPCProceduresParams struct {
	fx.In

	Server HelloYARPCServer
}

// FxHelloYARPCProceduresResult defines the output
// of NewFxHelloYARPCProcedures. It provides
// HelloYARPCServer procedures to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group.
// Dig 1.2 or newer must be used for this feature to work.
type FxHelloYARPCProceduresResult struct {
	fx.Out

	Procedures []transport.Procedure `group:"yarpcfx"`
}

// NewFxHelloYARPCProcedures provides HelloYARPCServer procedures to an Fx application.
// It expects a HelloYARPCServer to be present in the container.
//
//  fx.Provide(
//    streaming.NewFxHelloYARPCProcedures(),
//    ...
//  )
func NewFxHelloYARPCProcedures() interface{} {
	return func(params FxHelloYARPCProceduresParams) FxHelloYARPCProceduresResult {
		return FxHelloYARPCProceduresResult{
			Procedures: BuildHelloYARPCProcedures(params.Server),
		}
	}
}

type _HelloYARPCCaller struct {
	streamClient protobuf.StreamClient
}

func (c *_HelloYARPCCaller) HelloUnary(ctx context.Context, request *HelloRequest, options ...yarpc.CallOption) (*HelloResponse, error) {
	responseMessage, err := c.streamClient.Call(ctx, "HelloUnary", request, newHelloServiceHelloUnaryYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*HelloResponse)
	if !ok {
		return nil, protobuf.CastError(emptyHelloServiceHelloUnaryYARPCResponse, responseMessage)
	}
	return response, err
}

func (c *_HelloYARPCCaller) HelloOutStream(ctx context.Context, options ...yarpc.CallOption) (HelloServiceHelloOutStreamYARPCClient, error) {
	stream, err := c.streamClient.CallStream(ctx, "HelloOutStream", options...)
	if err != nil {
		return nil, err
	}
	return &_HelloServiceHelloOutStreamYARPCClient{stream: stream}, nil
}

func (c *_HelloYARPCCaller) HelloInStream(ctx context.Context, request *HelloRequest, options ...yarpc.CallOption) (HelloServiceHelloInStreamYARPCClient, error) {
	stream, err := c.streamClient.CallStream(ctx, "HelloInStream", options...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(request); err != nil {
		return nil, err
	}
	return &_HelloServiceHelloInStreamYARPCClient{stream: stream}, nil
}

func (c *_HelloYARPCCaller) HelloThere(ctx context.Context, options ...yarpc.CallOption) (HelloServiceHelloThereYARPCClient, error) {
	stream, err := c.streamClient.CallStream(ctx, "HelloThere", options...)
	if err != nil {
		return nil, err
	}
	return &_HelloServiceHelloThereYARPCClient{stream: stream}, nil
}

type _HelloYARPCHandler struct {
	server HelloYARPCServer
}

func (h *_HelloYARPCHandler) HelloUnary(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
	var request *HelloRequest
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*HelloRequest)
		if !ok {
			return nil, protobuf.CastError(emptyHelloServiceHelloUnaryYARPCRequest, requestMessage)
		}
	}
	response, err := h.server.HelloUnary(ctx, request)
	if response == nil {
		return nil, err
	}
	return response, err
}

func (h *_HelloYARPCHandler) HelloOutStream(serverStream *protobuf.ServerStream) error {
	response, err := h.server.HelloOutStream(&_HelloServiceHelloOutStreamYARPCServer{serverStream: serverStream})
	if err != nil {
		return err
	}
	return serverStream.Send(response)
}

func (h *_HelloYARPCHandler) HelloInStream(serverStream *protobuf.ServerStream) error {
	requestMessage, err := serverStream.Receive(newHelloServiceHelloInStreamYARPCRequest)
	if requestMessage == nil {
		return err
	}

	request, ok := requestMessage.(*HelloRequest)
	if !ok {
		return protobuf.CastError(emptyHelloServiceHelloInStreamYARPCRequest, requestMessage)
	}
	return h.server.HelloInStream(request, &_HelloServiceHelloInStreamYARPCServer{serverStream: serverStream})
}

func (h *_HelloYARPCHandler) HelloThere(serverStream *protobuf.ServerStream) error {
	return h.server.HelloThere(&_HelloServiceHelloThereYARPCServer{serverStream: serverStream})
}

type _HelloServiceHelloOutStreamYARPCClient struct {
	stream *protobuf.ClientStream
}

func (c *_HelloServiceHelloOutStreamYARPCClient) Context() context.Context {
	return c.stream.Context()
}

func (c *_HelloServiceHelloOutStreamYARPCClient) Send(request *HelloRequest, options ...yarpc.StreamOption) error {
	return c.stream.Send(request, options...)
}

func (c *_HelloServiceHelloOutStreamYARPCClient) CloseAndRecv(options ...yarpc.StreamOption) (*HelloResponse, error) {
	if err := c.stream.Close(options...); err != nil {
		return nil, err
	}
	responseMessage, err := c.stream.Receive(newHelloServiceHelloOutStreamYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*HelloResponse)
	if !ok {
		return nil, protobuf.CastError(emptyHelloServiceHelloOutStreamYARPCResponse, responseMessage)
	}
	return response, err
}

type _HelloServiceHelloInStreamYARPCClient struct {
	stream *protobuf.ClientStream
}

func (c *_HelloServiceHelloInStreamYARPCClient) Context() context.Context {
	return c.stream.Context()
}

func (c *_HelloServiceHelloInStreamYARPCClient) Recv(options ...yarpc.StreamOption) (*HelloResponse, error) {
	responseMessage, err := c.stream.Receive(newHelloServiceHelloInStreamYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*HelloResponse)
	if !ok {
		return nil, protobuf.CastError(emptyHelloServiceHelloInStreamYARPCResponse, responseMessage)
	}
	return response, err
}

func (c *_HelloServiceHelloInStreamYARPCClient) CloseSend(options ...yarpc.StreamOption) error {
	return c.stream.Close(options...)
}

type _HelloServiceHelloThereYARPCClient struct {
	stream *protobuf.ClientStream
}

func (c *_HelloServiceHelloThereYARPCClient) Context() context.Context {
	return c.stream.Context()
}

func (c *_HelloServiceHelloThereYARPCClient) Send(request *HelloRequest, options ...yarpc.StreamOption) error {
	return c.stream.Send(request, options...)
}

func (c *_HelloServiceHelloThereYARPCClient) Recv(options ...yarpc.StreamOption) (*HelloResponse, error) {
	responseMessage, err := c.stream.Receive(newHelloServiceHelloThereYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*HelloResponse)
	if !ok {
		return nil, protobuf.CastError(emptyHelloServiceHelloThereYARPCResponse, responseMessage)
	}
	return response, err
}

func (c *_HelloServiceHelloThereYARPCClient) CloseSend(options ...yarpc.StreamOption) error {
	return c.stream.Close(options...)
}

type _HelloServiceHelloOutStreamYARPCServer struct {
	serverStream *protobuf.ServerStream
}

func (s *_HelloServiceHelloOutStreamYARPCServer) Context() context.Context {
	return s.serverStream.Context()
}

func (s *_HelloServiceHelloOutStreamYARPCServer) Recv(options ...yarpc.StreamOption) (*HelloRequest, error) {
	requestMessage, err := s.serverStream.Receive(newHelloServiceHelloOutStreamYARPCRequest, options...)
	if requestMessage == nil {
		return nil, err
	}
	request, ok := requestMessage.(*HelloRequest)
	if !ok {
		return nil, protobuf.CastError(emptyHelloServiceHelloOutStreamYARPCRequest, requestMessage)
	}
	return request, err
}

type _HelloServiceHelloInStreamYARPCServer struct {
	serverStream *protobuf.ServerStream
}

func (s *_HelloServiceHelloInStreamYARPCServer) Context() context.Context {
	return s.serverStream.Context()
}

func (s *_HelloServiceHelloInStreamYARPCServer) Send(response *HelloResponse, options ...yarpc.StreamOption) error {
	return s.serverStream.Send(response, options...)
}

type _HelloServiceHelloThereYARPCServer struct {
	serverStream *protobuf.ServerStream
}

func (s *_HelloServiceHelloThereYARPCServer) Context() context.Context {
	return s.serverStream.Context()
}

func (s *_HelloServiceHelloThereYARPCServer) Recv(options ...yarpc.StreamOption) (*HelloRequest, error) {
	requestMessage, err := s.serverStream.Receive(newHelloServiceHelloThereYARPCRequest, options...)
	if requestMessage == nil {
		return nil, err
	}
	request, ok := requestMessage.(*HelloRequest)
	if !ok {
		return nil, protobuf.CastError(emptyHelloServiceHelloThereYARPCRequest, requestMessage)
	}
	return request, err
}

func (s *_HelloServiceHelloThereYARPCServer) Send(response *HelloResponse, options ...yarpc.StreamOption) error {
	return s.serverStream.Send(response, options...)
}

func newHelloServiceHelloUnaryYARPCRequest() proto.Message {
	return &HelloRequest{}
}

func newHelloServiceHelloUnaryYARPCResponse() proto.Message {
	return &HelloResponse{}
}

func newHelloServiceHelloThereYARPCRequest() proto.Message {
	return &HelloRequest{}
}

func newHelloServiceHelloThereYARPCResponse() proto.Message {
	return &HelloResponse{}
}

func newHelloServiceHelloOutStreamYARPCRequest() proto.Message {
	return &HelloRequest{}
}

func newHelloServiceHelloOutStreamYARPCResponse() proto.Message {
	return &HelloResponse{}
}

func newHelloServiceHelloInStreamYARPCRequest() proto.Message {
	return &HelloRequest{}
}

func newHelloServiceHelloInStreamYARPCResponse() proto.Message {
	return &HelloResponse{}
}

var (
	emptyHelloServiceHelloUnaryYARPCRequest      = &HelloRequest{}
	emptyHelloServiceHelloUnaryYARPCResponse     = &HelloResponse{}
	emptyHelloServiceHelloThereYARPCRequest      = &HelloRequest{}
	emptyHelloServiceHelloThereYARPCResponse     = &HelloResponse{}
	emptyHelloServiceHelloOutStreamYARPCRequest  = &HelloRequest{}
	emptyHelloServiceHelloOutStreamYARPCResponse = &HelloResponse{}
	emptyHelloServiceHelloInStreamYARPCRequest   = &HelloRequest{}
	emptyHelloServiceHelloInStreamYARPCResponse  = &HelloResponse{}
)

func init() {
	yarpc.RegisterClientBuilder(
		func(clientConfig transport.ClientConfig, structField reflect.StructField) HelloYARPCClient {
			return NewHelloYARPCClient(clientConfig, protobuf.ClientBuilderOptions(clientConfig, structField)...)
		},
	)
}

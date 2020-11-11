// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// CommandServiceClient is the client API for CommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommandServiceClient interface {
	Send(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandResponse, error)
}

type commandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommandServiceClient(cc grpc.ClientConnInterface) CommandServiceClient {
	return &commandServiceClient{cc}
}

func (c *commandServiceClient) Send(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandResponse, error) {
	out := new(CommandResponse)
	err := c.cc.Invoke(ctx, "/CommandService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandServiceServer is the server API for CommandService service.
// All implementations must embed UnimplementedCommandServiceServer
// for forward compatibility
type CommandServiceServer interface {
	Send(context.Context, *CommandRequest) (*CommandResponse, error)
	mustEmbedUnimplementedCommandServiceServer()
}

// UnimplementedCommandServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommandServiceServer struct {
}

func (UnimplementedCommandServiceServer) Send(context.Context, *CommandRequest) (*CommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedCommandServiceServer) mustEmbedUnimplementedCommandServiceServer() {}

// UnsafeCommandServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommandServiceServer will
// result in compilation errors.
type UnsafeCommandServiceServer interface {
	mustEmbedUnimplementedCommandServiceServer()
}

func RegisterCommandServiceServer(s grpc.ServiceRegistrar, srv CommandServiceServer) {
	s.RegisterService(&_CommandService_serviceDesc, srv)
}

func _CommandService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CommandService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServiceServer).Send(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommandService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CommandService",
	HandlerType: (*CommandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _CommandService_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "command.proto",
}
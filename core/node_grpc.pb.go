package core

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

// NodeServiceClient is the client API for NodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeServiceClient interface {
	ReportStatus(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	AssignTask(ctx context.Context, in *Request, opts ...grpc.CallOption) (NodeService_AssignTaskClient, error)
}

type nodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeServiceClient(cc grpc.ClientConnInterface) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) ReportStatus(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/core.NodeService/ReportStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) AssignTask(ctx context.Context, in *Request, opts ...grpc.CallOption) (NodeService_AssignTaskClient, error) {
	stream, err := c.cc.NewStream(ctx, &NodeService_ServiceDesc.Streams[0], "/core.NodeService/AssignTask", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceAssignTaskClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_AssignTaskClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type nodeServiceAssignTaskClient struct {
	grpc.ClientStream
}

func (x *nodeServiceAssignTaskClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NodeServiceServer is the server API for NodeService service.
// All implementations must embed UnimplementedNodeServiceServer
// for forward compatibility
type NodeServiceServer interface {
	ReportStatus(context.Context, *Request) (*Response, error)
	AssignTask(*Request, NodeService_AssignTaskServer) error
	mustEmbedUnimplementedNodeServiceServer()
}

// UnimplementedNodeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNodeServiceServer struct {
}

func (UnimplementedNodeServiceServer) ReportStatus(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportStatus not implemented")
}
func (UnimplementedNodeServiceServer) AssignTask(*Request, NodeService_AssignTaskServer) error {
	return status.Errorf(codes.Unimplemented, "method AssignTask not implemented")
}
func (UnimplementedNodeServiceServer) mustEmbedUnimplementedNodeServiceServer() {}

// UnsafeNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeServiceServer will
// result in compilation errors.
type UnsafeNodeServiceServer interface {
	mustEmbedUnimplementedNodeServiceServer()
}

func RegisterNodeServiceServer(s grpc.ServiceRegistrar, srv NodeServiceServer) {
	s.RegisterService(&NodeService_ServiceDesc, srv)
}

func _NodeService_ReportStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).ReportStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.NodeService/ReportStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).ReportStatus(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_AssignTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).AssignTask(m, &nodeServiceAssignTaskServer{stream})
}

type NodeService_AssignTaskServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type nodeServiceAssignTaskServer struct {
	grpc.ServerStream
}

func (x *nodeServiceAssignTaskServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

// NodeService_ServiceDesc is the grpc.ServiceDesc for NodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportStatus",
			Handler:    _NodeService_ReportStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AssignTask",
			Handler:       _NodeService_AssignTask_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "node.proto",
}

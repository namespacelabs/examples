// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: todos/api/todos/service.proto

package todos

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

// TodosServiceClient is the client API for TodosService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodosServiceClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	StreamList(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (TodosService_StreamListClient, error)
	GetRelatedData(ctx context.Context, in *GetRelatedDataRequest, opts ...grpc.CallOption) (*GetRelatedDataResponse, error)
}

type todosServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodosServiceClient(cc grpc.ClientConnInterface) TodosServiceClient {
	return &todosServiceClient{cc}
}

func (c *todosServiceClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/api.todos.TodosService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosServiceClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	out := new(RemoveResponse)
	err := c.cc.Invoke(ctx, "/api.todos.TodosService/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/api.todos.TodosService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosServiceClient) StreamList(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (TodosService_StreamListClient, error) {
	stream, err := c.cc.NewStream(ctx, &TodosService_ServiceDesc.Streams[0], "/api.todos.TodosService/StreamList", opts...)
	if err != nil {
		return nil, err
	}
	x := &todosServiceStreamListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TodosService_StreamListClient interface {
	Recv() (*ListResponse, error)
	grpc.ClientStream
}

type todosServiceStreamListClient struct {
	grpc.ClientStream
}

func (x *todosServiceStreamListClient) Recv() (*ListResponse, error) {
	m := new(ListResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *todosServiceClient) GetRelatedData(ctx context.Context, in *GetRelatedDataRequest, opts ...grpc.CallOption) (*GetRelatedDataResponse, error) {
	out := new(GetRelatedDataResponse)
	err := c.cc.Invoke(ctx, "/api.todos.TodosService/GetRelatedData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodosServiceServer is the server API for TodosService service.
// All implementations should embed UnimplementedTodosServiceServer
// for forward compatibility
type TodosServiceServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	StreamList(*ListRequest, TodosService_StreamListServer) error
	GetRelatedData(context.Context, *GetRelatedDataRequest) (*GetRelatedDataResponse, error)
}

// UnimplementedTodosServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTodosServiceServer struct {
}

func (UnimplementedTodosServiceServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedTodosServiceServer) Remove(context.Context, *RemoveRequest) (*RemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedTodosServiceServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTodosServiceServer) StreamList(*ListRequest, TodosService_StreamListServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamList not implemented")
}
func (UnimplementedTodosServiceServer) GetRelatedData(context.Context, *GetRelatedDataRequest) (*GetRelatedDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRelatedData not implemented")
}

// UnsafeTodosServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodosServiceServer will
// result in compilation errors.
type UnsafeTodosServiceServer interface {
	mustEmbedUnimplementedTodosServiceServer()
}

func RegisterTodosServiceServer(s grpc.ServiceRegistrar, srv TodosServiceServer) {
	s.RegisterService(&TodosService_ServiceDesc, srv)
}

func _TodosService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.todos.TodosService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServiceServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodosService_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServiceServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.todos.TodosService/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServiceServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodosService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.todos.TodosService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodosService_StreamList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TodosServiceServer).StreamList(m, &todosServiceStreamListServer{stream})
}

type TodosService_StreamListServer interface {
	Send(*ListResponse) error
	grpc.ServerStream
}

type todosServiceStreamListServer struct {
	grpc.ServerStream
}

func (x *todosServiceStreamListServer) Send(m *ListResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TodosService_GetRelatedData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRelatedDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServiceServer).GetRelatedData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.todos.TodosService/GetRelatedData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServiceServer).GetRelatedData(ctx, req.(*GetRelatedDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TodosService_ServiceDesc is the grpc.ServiceDesc for TodosService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodosService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.todos.TodosService",
	HandlerType: (*TodosServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _TodosService_Add_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _TodosService_Remove_Handler,
		},
		{
			MethodName: "List",
			Handler:    _TodosService_List_Handler,
		},
		{
			MethodName: "GetRelatedData",
			Handler:    _TodosService_GetRelatedData_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamList",
			Handler:       _TodosService_StreamList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "todos/api/todos/service.proto",
}

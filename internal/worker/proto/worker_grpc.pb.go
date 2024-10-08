// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: proto/worker.proto

package worker

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Worker_StartJob_FullMethodName        = "/Worker/StartJob"
	Worker_StopJob_FullMethodName         = "/Worker/StopJob"
	Worker_GetJobStatus_FullMethodName    = "/Worker/GetJobStatus"
	Worker_StreamJobOutput_FullMethodName = "/Worker/StreamJobOutput"
)

// WorkerClient is the client API for Worker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkerClient interface {
	// StartJob starts a job.
	StartJob(ctx context.Context, in *StartJobRequest, opts ...grpc.CallOption) (*StartJobReponse, error)
	// StopJob stops a job.
	StopJob(ctx context.Context, in *StopJobRequest, opts ...grpc.CallOption) (*StopJobResponse, error)
	// GetJobStatus gets the status of a job.
	GetJobStatus(ctx context.Context, in *GetJobStatusRequest, opts ...grpc.CallOption) (*GetJobStatusResponse, error)
	// StreamJobOutput streams the output of a job.
	// Note: If the job is still active when StreamOutputRequest
	// is sent the cli will tail the output.
	StreamJobOutput(ctx context.Context, in *StreamJobOutputRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamJobOutputResponse], error)
}

type workerClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerClient(cc grpc.ClientConnInterface) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) StartJob(ctx context.Context, in *StartJobRequest, opts ...grpc.CallOption) (*StartJobReponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartJobReponse)
	err := c.cc.Invoke(ctx, Worker_StartJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) StopJob(ctx context.Context, in *StopJobRequest, opts ...grpc.CallOption) (*StopJobResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StopJobResponse)
	err := c.cc.Invoke(ctx, Worker_StopJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) GetJobStatus(ctx context.Context, in *GetJobStatusRequest, opts ...grpc.CallOption) (*GetJobStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetJobStatusResponse)
	err := c.cc.Invoke(ctx, Worker_GetJobStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) StreamJobOutput(ctx context.Context, in *StreamJobOutputRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamJobOutputResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Worker_ServiceDesc.Streams[0], Worker_StreamJobOutput_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamJobOutputRequest, StreamJobOutputResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Worker_StreamJobOutputClient = grpc.ServerStreamingClient[StreamJobOutputResponse]

// WorkerServer is the server API for Worker service.
// All implementations must embed UnimplementedWorkerServer
// for forward compatibility.
type WorkerServer interface {
	// StartJob starts a job.
	StartJob(context.Context, *StartJobRequest) (*StartJobReponse, error)
	// StopJob stops a job.
	StopJob(context.Context, *StopJobRequest) (*StopJobResponse, error)
	// GetJobStatus gets the status of a job.
	GetJobStatus(context.Context, *GetJobStatusRequest) (*GetJobStatusResponse, error)
	// StreamJobOutput streams the output of a job.
	// Note: If the job is still active when StreamOutputRequest
	// is sent the cli will tail the output.
	StreamJobOutput(*StreamJobOutputRequest, grpc.ServerStreamingServer[StreamJobOutputResponse]) error
	mustEmbedUnimplementedWorkerServer()
}

// UnimplementedWorkerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWorkerServer struct{}

func (UnimplementedWorkerServer) StartJob(context.Context, *StartJobRequest) (*StartJobReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartJob not implemented")
}
func (UnimplementedWorkerServer) StopJob(context.Context, *StopJobRequest) (*StopJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopJob not implemented")
}
func (UnimplementedWorkerServer) GetJobStatus(context.Context, *GetJobStatusRequest) (*GetJobStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobStatus not implemented")
}
func (UnimplementedWorkerServer) StreamJobOutput(*StreamJobOutputRequest, grpc.ServerStreamingServer[StreamJobOutputResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamJobOutput not implemented")
}
func (UnimplementedWorkerServer) mustEmbedUnimplementedWorkerServer() {}
func (UnimplementedWorkerServer) testEmbeddedByValue()                {}

// UnsafeWorkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkerServer will
// result in compilation errors.
type UnsafeWorkerServer interface {
	mustEmbedUnimplementedWorkerServer()
}

func RegisterWorkerServer(s grpc.ServiceRegistrar, srv WorkerServer) {
	// If the following call pancis, it indicates UnimplementedWorkerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Worker_ServiceDesc, srv)
}

func _Worker_StartJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).StartJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Worker_StartJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).StartJob(ctx, req.(*StartJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_StopJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).StopJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Worker_StopJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).StopJob(ctx, req.(*StopJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_GetJobStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).GetJobStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Worker_GetJobStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).GetJobStatus(ctx, req.(*GetJobStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_StreamJobOutput_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamJobOutputRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WorkerServer).StreamJobOutput(m, &grpc.GenericServerStream[StreamJobOutputRequest, StreamJobOutputResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Worker_StreamJobOutputServer = grpc.ServerStreamingServer[StreamJobOutputResponse]

// Worker_ServiceDesc is the grpc.ServiceDesc for Worker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Worker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartJob",
			Handler:    _Worker_StartJob_Handler,
		},
		{
			MethodName: "StopJob",
			Handler:    _Worker_StopJob_Handler,
		},
		{
			MethodName: "GetJobStatus",
			Handler:    _Worker_GetJobStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamJobOutput",
			Handler:       _Worker_StreamJobOutput_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/worker.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/itr/V1/itr.proto

// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative itr.proto

package gen

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

const (
	InteractiveService_IncrReadCnt_FullMethodName = "/itr.V1.InteractiveService/IncrReadCnt"
	InteractiveService_Like_FullMethodName        = "/itr.V1.InteractiveService/Like"
	InteractiveService_CancelLike_FullMethodName  = "/itr.V1.InteractiveService/CancelLike"
	InteractiveService_Collect_FullMethodName     = "/itr.V1.InteractiveService/Collect"
	InteractiveService_Get_FullMethodName         = "/itr.V1.InteractiveService/Get"
	InteractiveService_GetByIds_FullMethodName    = "/itr.V1.InteractiveService/GetByIds"
)

// InteractiveServiceClient is the client API for InteractiveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InteractiveServiceClient interface {
	IncrReadCnt(ctx context.Context, in *IncrReadCntRequest, opts ...grpc.CallOption) (*IncrReadCntResponse, error)
	// Like 点赞
	Like(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeResponse, error)
	// CancelLike 取消点赞
	CancelLike(ctx context.Context, in *CancelLikeRequest, opts ...grpc.CallOption) (*CancelLikeResponse, error)
	// Collect 收藏
	Collect(ctx context.Context, in *CollectRequest, opts ...grpc.CallOption) (*CollectResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetByIds(ctx context.Context, in *GetByIdsRequest, opts ...grpc.CallOption) (*GetByIdsResponse, error)
}

type interactiveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInteractiveServiceClient(cc grpc.ClientConnInterface) InteractiveServiceClient {
	return &interactiveServiceClient{cc}
}

func (c *interactiveServiceClient) IncrReadCnt(ctx context.Context, in *IncrReadCntRequest, opts ...grpc.CallOption) (*IncrReadCntResponse, error) {
	out := new(IncrReadCntResponse)
	err := c.cc.Invoke(ctx, InteractiveService_IncrReadCnt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveServiceClient) Like(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeResponse, error) {
	out := new(LikeResponse)
	err := c.cc.Invoke(ctx, InteractiveService_Like_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveServiceClient) CancelLike(ctx context.Context, in *CancelLikeRequest, opts ...grpc.CallOption) (*CancelLikeResponse, error) {
	out := new(CancelLikeResponse)
	err := c.cc.Invoke(ctx, InteractiveService_CancelLike_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveServiceClient) Collect(ctx context.Context, in *CollectRequest, opts ...grpc.CallOption) (*CollectResponse, error) {
	out := new(CollectResponse)
	err := c.cc.Invoke(ctx, InteractiveService_Collect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, InteractiveService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveServiceClient) GetByIds(ctx context.Context, in *GetByIdsRequest, opts ...grpc.CallOption) (*GetByIdsResponse, error) {
	out := new(GetByIdsResponse)
	err := c.cc.Invoke(ctx, InteractiveService_GetByIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InteractiveServiceServer is the server API for InteractiveService service.
// All implementations must embed UnimplementedInteractiveServiceServer
// for forward compatibility
type InteractiveServiceServer interface {
	IncrReadCnt(context.Context, *IncrReadCntRequest) (*IncrReadCntResponse, error)
	// Like 点赞
	Like(context.Context, *LikeRequest) (*LikeResponse, error)
	// CancelLike 取消点赞
	CancelLike(context.Context, *CancelLikeRequest) (*CancelLikeResponse, error)
	// Collect 收藏
	Collect(context.Context, *CollectRequest) (*CollectResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetByIds(context.Context, *GetByIdsRequest) (*GetByIdsResponse, error)
	mustEmbedUnimplementedInteractiveServiceServer()
}

// UnimplementedInteractiveServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInteractiveServiceServer struct {
}

func (UnimplementedInteractiveServiceServer) IncrReadCnt(context.Context, *IncrReadCntRequest) (*IncrReadCntResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrReadCnt not implemented")
}
func (UnimplementedInteractiveServiceServer) Like(context.Context, *LikeRequest) (*LikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Like not implemented")
}
func (UnimplementedInteractiveServiceServer) CancelLike(context.Context, *CancelLikeRequest) (*CancelLikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelLike not implemented")
}
func (UnimplementedInteractiveServiceServer) Collect(context.Context, *CollectRequest) (*CollectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collect not implemented")
}
func (UnimplementedInteractiveServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedInteractiveServiceServer) GetByIds(context.Context, *GetByIdsRequest) (*GetByIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIds not implemented")
}
func (UnimplementedInteractiveServiceServer) mustEmbedUnimplementedInteractiveServiceServer() {}

// UnsafeInteractiveServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InteractiveServiceServer will
// result in compilation errors.
type UnsafeInteractiveServiceServer interface {
	mustEmbedUnimplementedInteractiveServiceServer()
}

func RegisterInteractiveServiceServer(s grpc.ServiceRegistrar, srv InteractiveServiceServer) {
	s.RegisterService(&InteractiveService_ServiceDesc, srv)
}

func _InteractiveService_IncrReadCnt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrReadCntRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServiceServer).IncrReadCnt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractiveService_IncrReadCnt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServiceServer).IncrReadCnt(ctx, req.(*IncrReadCntRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractiveService_Like_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServiceServer).Like(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractiveService_Like_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServiceServer).Like(ctx, req.(*LikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractiveService_CancelLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServiceServer).CancelLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractiveService_CancelLike_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServiceServer).CancelLike(ctx, req.(*CancelLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractiveService_Collect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServiceServer).Collect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractiveService_Collect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServiceServer).Collect(ctx, req.(*CollectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractiveService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractiveService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractiveService_GetByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServiceServer).GetByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractiveService_GetByIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServiceServer).GetByIds(ctx, req.(*GetByIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InteractiveService_ServiceDesc is the grpc.ServiceDesc for InteractiveService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InteractiveService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "itr.V1.InteractiveService",
	HandlerType: (*InteractiveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IncrReadCnt",
			Handler:    _InteractiveService_IncrReadCnt_Handler,
		},
		{
			MethodName: "Like",
			Handler:    _InteractiveService_Like_Handler,
		},
		{
			MethodName: "CancelLike",
			Handler:    _InteractiveService_CancelLike_Handler,
		},
		{
			MethodName: "Collect",
			Handler:    _InteractiveService_Collect_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _InteractiveService_Get_Handler,
		},
		{
			MethodName: "GetByIds",
			Handler:    _InteractiveService_GetByIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/itr/V1/itr.proto",
}

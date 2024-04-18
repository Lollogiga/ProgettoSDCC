// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: registry/election.proto

package registry

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

// RegistryClient is the client API for Registry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegistryClient interface {
	JoinNetwork(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinReply, error)
}

type registryClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistryClient(cc grpc.ClientConnInterface) RegistryClient {
	return &registryClient{cc}
}

func (c *registryClient) JoinNetwork(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinReply, error) {
	out := new(JoinReply)
	err := c.cc.Invoke(ctx, "/registry.Registry/JoinNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistryServer is the server API for Registry service.
// All implementations must embed UnimplementedRegistryServer
// for forward compatibility
type RegistryServer interface {
	JoinNetwork(context.Context, *JoinRequest) (*JoinReply, error)
	mustEmbedUnimplementedRegistryServer()
}

// UnimplementedRegistryServer must be embedded to have forward compatible implementations.
type UnimplementedRegistryServer struct {
}

func (UnimplementedRegistryServer) JoinNetwork(context.Context, *JoinRequest) (*JoinReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinNetwork not implemented")
}
func (UnimplementedRegistryServer) mustEmbedUnimplementedRegistryServer() {}

// UnsafeRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegistryServer will
// result in compilation errors.
type UnsafeRegistryServer interface {
	mustEmbedUnimplementedRegistryServer()
}

func RegisterRegistryServer(s grpc.ServiceRegistrar, srv RegistryServer) {
	s.RegisterService(&Registry_ServiceDesc, srv)
}

func _Registry_JoinNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).JoinNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Registry/JoinNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).JoinNetwork(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Registry_ServiceDesc is the grpc.ServiceDesc for Registry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Registry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "registry.Registry",
	HandlerType: (*RegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinNetwork",
			Handler:    _Registry_JoinNetwork_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry/election.proto",
}

// UpdateClient is the client API for Update service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UpdateClient interface {
	UpdateNetwork(ctx context.Context, in *UpdateMessage, opts ...grpc.CallOption) (*UpdateResponse, error)
}

type updateClient struct {
	cc grpc.ClientConnInterface
}

func NewUpdateClient(cc grpc.ClientConnInterface) UpdateClient {
	return &updateClient{cc}
}

func (c *updateClient) UpdateNetwork(ctx context.Context, in *UpdateMessage, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/registry.Update/UpdateNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateServer is the server API for Update service.
// All implementations must embed UnimplementedUpdateServer
// for forward compatibility
type UpdateServer interface {
	UpdateNetwork(context.Context, *UpdateMessage) (*UpdateResponse, error)
	mustEmbedUnimplementedUpdateServer()
}

// UnimplementedUpdateServer must be embedded to have forward compatible implementations.
type UnimplementedUpdateServer struct {
}

func (UnimplementedUpdateServer) UpdateNetwork(context.Context, *UpdateMessage) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNetwork not implemented")
}
func (UnimplementedUpdateServer) mustEmbedUnimplementedUpdateServer() {}

// UnsafeUpdateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UpdateServer will
// result in compilation errors.
type UnsafeUpdateServer interface {
	mustEmbedUnimplementedUpdateServer()
}

func RegisterUpdateServer(s grpc.ServiceRegistrar, srv UpdateServer) {
	s.RegisterService(&Update_ServiceDesc, srv)
}

func _Update_UpdateNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateServer).UpdateNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Update/UpdateNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateServer).UpdateNetwork(ctx, req.(*UpdateMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// Update_ServiceDesc is the grpc.ServiceDesc for Update service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Update_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "registry.Update",
	HandlerType: (*UpdateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateNetwork",
			Handler:    _Update_UpdateNetwork_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry/election.proto",
}

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	GetTime(ctx context.Context, in *TimeRequest, opts ...grpc.CallOption) (*TimeReply, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) GetTime(ctx context.Context, in *TimeRequest, opts ...grpc.CallOption) (*TimeReply, error) {
	out := new(TimeReply)
	err := c.cc.Invoke(ctx, "/registry.Service/GetTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	GetTime(context.Context, *TimeRequest) (*TimeReply, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) GetTime(context.Context, *TimeRequest) (*TimeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTime not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_GetTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.Service/GetTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetTime(ctx, req.(*TimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "registry.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTime",
			Handler:    _Service_GetTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry/election.proto",
}

// ElectionClient is the client API for Election service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ElectionClient interface {
	BullyElection(ctx context.Context, in *ElectionRequest, opts ...grpc.CallOption) (*ElectionReply, error)
	UpdateRegistry(ctx context.Context, in *IdLeader, opts ...grpc.CallOption) (*Nil, error)
	DolevElection(ctx context.Context, in *ElectionRequest, opts ...grpc.CallOption) (*ElectionReply, error)
}

type electionClient struct {
	cc grpc.ClientConnInterface
}

func NewElectionClient(cc grpc.ClientConnInterface) ElectionClient {
	return &electionClient{cc}
}

func (c *electionClient) BullyElection(ctx context.Context, in *ElectionRequest, opts ...grpc.CallOption) (*ElectionReply, error) {
	out := new(ElectionReply)
	err := c.cc.Invoke(ctx, "/registry.election/BullyElection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *electionClient) UpdateRegistry(ctx context.Context, in *IdLeader, opts ...grpc.CallOption) (*Nil, error) {
	out := new(Nil)
	err := c.cc.Invoke(ctx, "/registry.election/UpdateRegistry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *electionClient) DolevElection(ctx context.Context, in *ElectionRequest, opts ...grpc.CallOption) (*ElectionReply, error) {
	out := new(ElectionReply)
	err := c.cc.Invoke(ctx, "/registry.election/DolevElection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ElectionServer is the server API for Election service.
// All implementations must embed UnimplementedElectionServer
// for forward compatibility
type ElectionServer interface {
	BullyElection(context.Context, *ElectionRequest) (*ElectionReply, error)
	UpdateRegistry(context.Context, *IdLeader) (*Nil, error)
	DolevElection(context.Context, *ElectionRequest) (*ElectionReply, error)
	mustEmbedUnimplementedElectionServer()
}

// UnimplementedElectionServer must be embedded to have forward compatible implementations.
type UnimplementedElectionServer struct {
}

func (UnimplementedElectionServer) BullyElection(context.Context, *ElectionRequest) (*ElectionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BullyElection not implemented")
}
func (UnimplementedElectionServer) UpdateRegistry(context.Context, *IdLeader) (*Nil, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRegistry not implemented")
}
func (UnimplementedElectionServer) DolevElection(context.Context, *ElectionRequest) (*ElectionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DolevElection not implemented")
}
func (UnimplementedElectionServer) mustEmbedUnimplementedElectionServer() {}

// UnsafeElectionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ElectionServer will
// result in compilation errors.
type UnsafeElectionServer interface {
	mustEmbedUnimplementedElectionServer()
}

func RegisterElectionServer(s grpc.ServiceRegistrar, srv ElectionServer) {
	s.RegisterService(&Election_ServiceDesc, srv)
}

func _Election_BullyElection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ElectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServer).BullyElection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.election/BullyElection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServer).BullyElection(ctx, req.(*ElectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Election_UpdateRegistry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdLeader)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServer).UpdateRegistry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.election/UpdateRegistry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServer).UpdateRegistry(ctx, req.(*IdLeader))
	}
	return interceptor(ctx, in, info, handler)
}

func _Election_DolevElection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ElectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServer).DolevElection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.election/DolevElection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServer).DolevElection(ctx, req.(*ElectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Election_ServiceDesc is the grpc.ServiceDesc for Election service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Election_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "registry.election",
	HandlerType: (*ElectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BullyElection",
			Handler:    _Election_BullyElection_Handler,
		},
		{
			MethodName: "UpdateRegistry",
			Handler:    _Election_UpdateRegistry_Handler,
		},
		{
			MethodName: "DolevElection",
			Handler:    _Election_DolevElection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry/election.proto",
}

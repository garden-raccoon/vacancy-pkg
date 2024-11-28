// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: vacancy-service.proto

package vacancy

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
	OrderAdminService_CreateVacancy_FullMethodName = "/models.OrderAdminService/CreateVacancy"
)

// OrderAdminServiceClient is the client API for OrderAdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// OrderService is
type OrderAdminServiceClient interface {
	CreateVacancy(ctx context.Context, in *Vacancy, opts ...grpc.CallOption) (*VacancyDbEmpty, error)
}

type orderAdminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderAdminServiceClient(cc grpc.ClientConnInterface) OrderAdminServiceClient {
	return &orderAdminServiceClient{cc}
}

func (c *orderAdminServiceClient) CreateVacancy(ctx context.Context, in *Vacancy, opts ...grpc.CallOption) (*VacancyDbEmpty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VacancyDbEmpty)
	err := c.cc.Invoke(ctx, OrderAdminService_CreateVacancy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderAdminServiceServer is the server API for OrderAdminService service.
// All implementations must embed UnimplementedOrderAdminServiceServer
// for forward compatibility.
//
// OrderService is
type OrderAdminServiceServer interface {
	CreateVacancy(context.Context, *Vacancy) (*VacancyDbEmpty, error)
	mustEmbedUnimplementedOrderAdminServiceServer()
}

// UnimplementedOrderAdminServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOrderAdminServiceServer struct{}

func (UnimplementedOrderAdminServiceServer) CreateVacancy(context.Context, *Vacancy) (*VacancyDbEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVacancy not implemented")
}
func (UnimplementedOrderAdminServiceServer) mustEmbedUnimplementedOrderAdminServiceServer() {}
func (UnimplementedOrderAdminServiceServer) testEmbeddedByValue()                           {}

// UnsafeOrderAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderAdminServiceServer will
// result in compilation errors.
type UnsafeOrderAdminServiceServer interface {
	mustEmbedUnimplementedOrderAdminServiceServer()
}

func RegisterOrderAdminServiceServer(s grpc.ServiceRegistrar, srv OrderAdminServiceServer) {
	// If the following call pancis, it indicates UnimplementedOrderAdminServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OrderAdminService_ServiceDesc, srv)
}

func _OrderAdminService_CreateVacancy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Vacancy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderAdminServiceServer).CreateVacancy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderAdminService_CreateVacancy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderAdminServiceServer).CreateVacancy(ctx, req.(*Vacancy))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderAdminService_ServiceDesc is the grpc.ServiceDesc for OrderAdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderAdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "models.OrderAdminService",
	HandlerType: (*OrderAdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVacancy",
			Handler:    _OrderAdminService_CreateVacancy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vacancy-service.proto",
}

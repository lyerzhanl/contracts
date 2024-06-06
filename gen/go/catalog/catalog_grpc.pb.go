// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.0
// source: catalog/catalog.proto

package catalog

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

// CatalogueServiceClient is the client API for CatalogueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatalogueServiceClient interface {
	CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error)
	ListItems(ctx context.Context, in *ListItemsRequest, opts ...grpc.CallOption) (*ListItemsResponse, error)
	GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*GetItemResponse, error)
}

type catalogueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogueServiceClient(cc grpc.ClientConnInterface) CatalogueServiceClient {
	return &catalogueServiceClient{cc}
}

func (c *catalogueServiceClient) CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error) {
	out := new(CreateItemResponse)
	err := c.cc.Invoke(ctx, "/catalogue.CatalogueService/CreateItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogueServiceClient) ListItems(ctx context.Context, in *ListItemsRequest, opts ...grpc.CallOption) (*ListItemsResponse, error) {
	out := new(ListItemsResponse)
	err := c.cc.Invoke(ctx, "/catalogue.CatalogueService/ListItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogueServiceClient) GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*GetItemResponse, error) {
	out := new(GetItemResponse)
	err := c.cc.Invoke(ctx, "/catalogue.CatalogueService/GetItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogueServiceServer is the server API for CatalogueService service.
// All implementations must embed UnimplementedCatalogueServiceServer
// for forward compatibility
type CatalogueServiceServer interface {
	CreateItem(context.Context, *CreateItemRequest) (*CreateItemResponse, error)
	ListItems(context.Context, *ListItemsRequest) (*ListItemsResponse, error)
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)
	mustEmbedUnimplementedCatalogueServiceServer()
}

// UnimplementedCatalogueServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCatalogueServiceServer struct {
}

func (UnimplementedCatalogueServiceServer) CreateItem(context.Context, *CreateItemRequest) (*CreateItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateItem not implemented")
}
func (UnimplementedCatalogueServiceServer) ListItems(context.Context, *ListItemsRequest) (*ListItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListItems not implemented")
}
func (UnimplementedCatalogueServiceServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItem not implemented")
}
func (UnimplementedCatalogueServiceServer) mustEmbedUnimplementedCatalogueServiceServer() {}

// UnsafeCatalogueServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatalogueServiceServer will
// result in compilation errors.
type UnsafeCatalogueServiceServer interface {
	mustEmbedUnimplementedCatalogueServiceServer()
}

func RegisterCatalogueServiceServer(s grpc.ServiceRegistrar, srv CatalogueServiceServer) {
	s.RegisterService(&CatalogueService_ServiceDesc, srv)
}

func _CatalogueService_CreateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogueServiceServer).CreateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalogue.CatalogueService/CreateItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogueServiceServer).CreateItem(ctx, req.(*CreateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogueService_ListItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogueServiceServer).ListItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalogue.CatalogueService/ListItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogueServiceServer).ListItems(ctx, req.(*ListItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogueService_GetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogueServiceServer).GetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalogue.CatalogueService/GetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogueServiceServer).GetItem(ctx, req.(*GetItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CatalogueService_ServiceDesc is the grpc.ServiceDesc for CatalogueService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CatalogueService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "catalogue.CatalogueService",
	HandlerType: (*CatalogueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateItem",
			Handler:    _CatalogueService_CreateItem_Handler,
		},
		{
			MethodName: "ListItems",
			Handler:    _CatalogueService_ListItems_Handler,
		},
		{
			MethodName: "GetItem",
			Handler:    _CatalogueService_GetItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "catalog/catalog.proto",
}

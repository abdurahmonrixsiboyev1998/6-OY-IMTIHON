// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: proto/income/income.proto

package income

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	TransactionService_CreateTransaction_FullMethodName         = "/TransactionService/CreateTransaction"
	TransactionService_UpdateTransactionByID_FullMethodName     = "/TransactionService/UpdateTransactionByID"
	TransactionService_DeleteTransactionByID_FullMethodName     = "/TransactionService/DeleteTransactionByID"
	TransactionService_GetTransactionByID_FullMethodName        = "/TransactionService/GetTransactionByID"
	TransactionService_GetTransactionsByCategory_FullMethodName = "/TransactionService/GetTransactionsByCategory"
	TransactionService_GetTransactionByDate_FullMethodName      = "/TransactionService/GetTransactionByDate"
	TransactionService_GetReport_FullMethodName                 = "/TransactionService/GetReport"
)

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionServiceClient interface {
	CreateTransaction(ctx context.Context, in *TransactionInfo, opts ...grpc.CallOption) (*TransactionID, error)
	UpdateTransactionByID(ctx context.Context, in *TransactionWithID, opts ...grpc.CallOption) (*TransactionResponse, error)
	DeleteTransactionByID(ctx context.Context, in *TransactionID, opts ...grpc.CallOption) (*TransactionResponse, error)
	GetTransactionByID(ctx context.Context, in *TransactionID, opts ...grpc.CallOption) (*TransactionWithID, error)
	GetTransactionsByCategory(ctx context.Context, in *TransactionCategory, opts ...grpc.CallOption) (*ListTransactions, error)
	GetTransactionByDate(ctx context.Context, in *TransactionDate, opts ...grpc.CallOption) (*ListTransactions, error)
	GetReport(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ReportResponse, error)
}

type transactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionServiceClient(cc grpc.ClientConnInterface) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) CreateTransaction(ctx context.Context, in *TransactionInfo, opts ...grpc.CallOption) (*TransactionID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransactionID)
	err := c.cc.Invoke(ctx, TransactionService_CreateTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) UpdateTransactionByID(ctx context.Context, in *TransactionWithID, opts ...grpc.CallOption) (*TransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, TransactionService_UpdateTransactionByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) DeleteTransactionByID(ctx context.Context, in *TransactionID, opts ...grpc.CallOption) (*TransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, TransactionService_DeleteTransactionByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetTransactionByID(ctx context.Context, in *TransactionID, opts ...grpc.CallOption) (*TransactionWithID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransactionWithID)
	err := c.cc.Invoke(ctx, TransactionService_GetTransactionByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetTransactionsByCategory(ctx context.Context, in *TransactionCategory, opts ...grpc.CallOption) (*ListTransactions, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTransactions)
	err := c.cc.Invoke(ctx, TransactionService_GetTransactionsByCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetTransactionByDate(ctx context.Context, in *TransactionDate, opts ...grpc.CallOption) (*ListTransactions, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTransactions)
	err := c.cc.Invoke(ctx, TransactionService_GetTransactionByDate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetReport(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ReportResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReportResponse)
	err := c.cc.Invoke(ctx, TransactionService_GetReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceServer is the server API for TransactionService service.
// All implementations must embed UnimplementedTransactionServiceServer
// for forward compatibility
type TransactionServiceServer interface {
	CreateTransaction(context.Context, *TransactionInfo) (*TransactionID, error)
	UpdateTransactionByID(context.Context, *TransactionWithID) (*TransactionResponse, error)
	DeleteTransactionByID(context.Context, *TransactionID) (*TransactionResponse, error)
	GetTransactionByID(context.Context, *TransactionID) (*TransactionWithID, error)
	GetTransactionsByCategory(context.Context, *TransactionCategory) (*ListTransactions, error)
	GetTransactionByDate(context.Context, *TransactionDate) (*ListTransactions, error)
	GetReport(context.Context, *Empty) (*ReportResponse, error)
	mustEmbedUnimplementedTransactionServiceServer()
}

// UnimplementedTransactionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionServiceServer struct {
}

func (UnimplementedTransactionServiceServer) CreateTransaction(context.Context, *TransactionInfo) (*TransactionID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) UpdateTransactionByID(context.Context, *TransactionWithID) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTransactionByID not implemented")
}
func (UnimplementedTransactionServiceServer) DeleteTransactionByID(context.Context, *TransactionID) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTransactionByID not implemented")
}
func (UnimplementedTransactionServiceServer) GetTransactionByID(context.Context, *TransactionID) (*TransactionWithID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionByID not implemented")
}
func (UnimplementedTransactionServiceServer) GetTransactionsByCategory(context.Context, *TransactionCategory) (*ListTransactions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionsByCategory not implemented")
}
func (UnimplementedTransactionServiceServer) GetTransactionByDate(context.Context, *TransactionDate) (*ListTransactions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionByDate not implemented")
}
func (UnimplementedTransactionServiceServer) GetReport(context.Context, *Empty) (*ReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReport not implemented")
}
func (UnimplementedTransactionServiceServer) mustEmbedUnimplementedTransactionServiceServer() {}

// UnsafeTransactionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServiceServer will
// result in compilation errors.
type UnsafeTransactionServiceServer interface {
	mustEmbedUnimplementedTransactionServiceServer()
}

func RegisterTransactionServiceServer(s grpc.ServiceRegistrar, srv TransactionServiceServer) {
	s.RegisterService(&TransactionService_ServiceDesc, srv)
}

func _TransactionService_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_CreateTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).CreateTransaction(ctx, req.(*TransactionInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_UpdateTransactionByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionWithID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).UpdateTransactionByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_UpdateTransactionByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).UpdateTransactionByID(ctx, req.(*TransactionWithID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_DeleteTransactionByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).DeleteTransactionByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_DeleteTransactionByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).DeleteTransactionByID(ctx, req.(*TransactionID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GetTransactionByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetTransactionByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_GetTransactionByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetTransactionByID(ctx, req.(*TransactionID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GetTransactionsByCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionCategory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetTransactionsByCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_GetTransactionsByCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetTransactionsByCategory(ctx, req.(*TransactionCategory))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GetTransactionByDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionDate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetTransactionByDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_GetTransactionByDate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetTransactionByDate(ctx, req.(*TransactionDate))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GetReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_GetReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetReport(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionService_ServiceDesc is the grpc.ServiceDesc for TransactionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTransaction",
			Handler:    _TransactionService_CreateTransaction_Handler,
		},
		{
			MethodName: "UpdateTransactionByID",
			Handler:    _TransactionService_UpdateTransactionByID_Handler,
		},
		{
			MethodName: "DeleteTransactionByID",
			Handler:    _TransactionService_DeleteTransactionByID_Handler,
		},
		{
			MethodName: "GetTransactionByID",
			Handler:    _TransactionService_GetTransactionByID_Handler,
		},
		{
			MethodName: "GetTransactionsByCategory",
			Handler:    _TransactionService_GetTransactionsByCategory_Handler,
		},
		{
			MethodName: "GetTransactionByDate",
			Handler:    _TransactionService_GetTransactionByDate_Handler,
		},
		{
			MethodName: "GetReport",
			Handler:    _TransactionService_GetReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/income/income.proto",
}

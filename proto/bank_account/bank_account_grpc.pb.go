// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: bank_account.proto

package bankAccountService

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

// BankAccountServiceClient is the client API for BankAccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BankAccountServiceClient interface {
	CreateBankAccount(ctx context.Context, in *CreateBankAccountRequest, opts ...grpc.CallOption) (*CreateBankAccountResponse, error)
	DepositBalance(ctx context.Context, in *DepositBalanceRequest, opts ...grpc.CallOption) (*DepositBalanceResponse, error)
	ChangeEmail(ctx context.Context, in *ChangeEmailRequest, opts ...grpc.CallOption) (*ChangeEmailResponse, error)
	GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetByIdResponse, error)
	GetBankAccountByStatus(ctx context.Context, in *GetBankAccountByStatusRequest, opts ...grpc.CallOption) (*GetBankAccountByStatusResponse, error)
}

type bankAccountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBankAccountServiceClient(cc grpc.ClientConnInterface) BankAccountServiceClient {
	return &bankAccountServiceClient{cc}
}

func (c *bankAccountServiceClient) CreateBankAccount(ctx context.Context, in *CreateBankAccountRequest, opts ...grpc.CallOption) (*CreateBankAccountResponse, error) {
	out := new(CreateBankAccountResponse)
	err := c.cc.Invoke(ctx, "/orderService.bankAccountService/CreateBankAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankAccountServiceClient) DepositBalance(ctx context.Context, in *DepositBalanceRequest, opts ...grpc.CallOption) (*DepositBalanceResponse, error) {
	out := new(DepositBalanceResponse)
	err := c.cc.Invoke(ctx, "/orderService.bankAccountService/DepositBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankAccountServiceClient) ChangeEmail(ctx context.Context, in *ChangeEmailRequest, opts ...grpc.CallOption) (*ChangeEmailResponse, error) {
	out := new(ChangeEmailResponse)
	err := c.cc.Invoke(ctx, "/orderService.bankAccountService/ChangeEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankAccountServiceClient) GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetByIdResponse, error) {
	out := new(GetByIdResponse)
	err := c.cc.Invoke(ctx, "/orderService.bankAccountService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankAccountServiceClient) GetBankAccountByStatus(ctx context.Context, in *GetBankAccountByStatusRequest, opts ...grpc.CallOption) (*GetBankAccountByStatusResponse, error) {
	out := new(GetBankAccountByStatusResponse)
	err := c.cc.Invoke(ctx, "/orderService.bankAccountService/GetBankAccountByStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankAccountServiceServer is the server API for BankAccountService service.
// All implementations should embed UnimplementedBankAccountServiceServer
// for forward compatibility
type BankAccountServiceServer interface {
	CreateBankAccount(context.Context, *CreateBankAccountRequest) (*CreateBankAccountResponse, error)
	DepositBalance(context.Context, *DepositBalanceRequest) (*DepositBalanceResponse, error)
	ChangeEmail(context.Context, *ChangeEmailRequest) (*ChangeEmailResponse, error)
	GetById(context.Context, *GetByIdRequest) (*GetByIdResponse, error)
	GetBankAccountByStatus(context.Context, *GetBankAccountByStatusRequest) (*GetBankAccountByStatusResponse, error)
}

// UnimplementedBankAccountServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBankAccountServiceServer struct {
}

func (UnimplementedBankAccountServiceServer) CreateBankAccount(context.Context, *CreateBankAccountRequest) (*CreateBankAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBankAccount not implemented")
}
func (UnimplementedBankAccountServiceServer) DepositBalance(context.Context, *DepositBalanceRequest) (*DepositBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DepositBalance not implemented")
}
func (UnimplementedBankAccountServiceServer) ChangeEmail(context.Context, *ChangeEmailRequest) (*ChangeEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeEmail not implemented")
}
func (UnimplementedBankAccountServiceServer) GetById(context.Context, *GetByIdRequest) (*GetByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedBankAccountServiceServer) GetBankAccountByStatus(context.Context, *GetBankAccountByStatusRequest) (*GetBankAccountByStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBankAccountByStatus not implemented")
}

// UnsafeBankAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BankAccountServiceServer will
// result in compilation errors.
type UnsafeBankAccountServiceServer interface {
	mustEmbedUnimplementedBankAccountServiceServer()
}

func RegisterBankAccountServiceServer(s grpc.ServiceRegistrar, srv BankAccountServiceServer) {
	s.RegisterService(&BankAccountService_ServiceDesc, srv)
}

func _BankAccountService_CreateBankAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBankAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankAccountServiceServer).CreateBankAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderService.bankAccountService/CreateBankAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankAccountServiceServer).CreateBankAccount(ctx, req.(*CreateBankAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankAccountService_DepositBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankAccountServiceServer).DepositBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderService.bankAccountService/DepositBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankAccountServiceServer).DepositBalance(ctx, req.(*DepositBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankAccountService_ChangeEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankAccountServiceServer).ChangeEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderService.bankAccountService/ChangeEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankAccountServiceServer).ChangeEmail(ctx, req.(*ChangeEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankAccountService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankAccountServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderService.bankAccountService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankAccountServiceServer).GetById(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankAccountService_GetBankAccountByStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBankAccountByStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankAccountServiceServer).GetBankAccountByStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderService.bankAccountService/GetBankAccountByStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankAccountServiceServer).GetBankAccountByStatus(ctx, req.(*GetBankAccountByStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BankAccountService_ServiceDesc is the grpc.ServiceDesc for BankAccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BankAccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orderService.bankAccountService",
	HandlerType: (*BankAccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBankAccount",
			Handler:    _BankAccountService_CreateBankAccount_Handler,
		},
		{
			MethodName: "DepositBalance",
			Handler:    _BankAccountService_DepositBalance_Handler,
		},
		{
			MethodName: "ChangeEmail",
			Handler:    _BankAccountService_ChangeEmail_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _BankAccountService_GetById_Handler,
		},
		{
			MethodName: "GetBankAccountByStatus",
			Handler:    _BankAccountService_GetBankAccountByStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bank_account.proto",
}

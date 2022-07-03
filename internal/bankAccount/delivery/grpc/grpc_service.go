package grpc

import (
	"context"
	"github.com/AleksK1NG/go-cqrs-eventsourcing/config"
	"github.com/AleksK1NG/go-cqrs-eventsourcing/internal/bankAccount/commands"
	"github.com/AleksK1NG/go-cqrs-eventsourcing/internal/bankAccount/queries"
	"github.com/AleksK1NG/go-cqrs-eventsourcing/internal/bankAccount/service"
	"github.com/AleksK1NG/go-cqrs-eventsourcing/internal/mappers"
	"github.com/AleksK1NG/go-cqrs-eventsourcing/pkg/grpc_errors"
	"github.com/AleksK1NG/go-cqrs-eventsourcing/pkg/logger"
	"github.com/AleksK1NG/go-cqrs-eventsourcing/pkg/tracing"
	"github.com/AleksK1NG/go-cqrs-eventsourcing/pkg/utils"
	"github.com/AleksK1NG/go-cqrs-eventsourcing/proto/bank_account"
	"github.com/opentracing/opentracing-go/log"
	uuid "github.com/satori/go.uuid"
)

type grpcService struct {
	log                logger.Logger
	cfg                *config.Config
	bankAccountService *service.BankAccountService
}

func NewGrpcService(log logger.Logger, cfg *config.Config, bankAccountService *service.BankAccountService) *grpcService {
	return &grpcService{log: log, cfg: cfg, bankAccountService: bankAccountService}
}

func (g *grpcService) CreateBankAccount(ctx context.Context, request *bankAccountService.CreateBankAccountRequest) (*bankAccountService.CreateBankAccountResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "grpcService.CreateBankAccount")
	defer span.Finish()
	span.LogFields(log.String("req", request.String()))

	aggregateID := uuid.NewV4().String()
	command := commands.CreateBankAccountCommand{
		AggregateID: aggregateID,
		Email:       request.GetEmail(),
		Address:     request.GetAddress(),
		FirstName:   request.GetFirstName(),
		LastName:    request.GetLastName(),
		Balance:     0,
		Status:      request.GetStatus(),
	}

	err := g.bankAccountService.Commands.CreateBankAccount.Handle(ctx, command)
	if err != nil {
		g.log.Errorf("(CreateBankAccount.Handle) err: %v", err)
		return nil, grpc_errors.ErrResponse(tracing.TraceWithErr(span, err))
	}

	return &bankAccountService.CreateBankAccountResponse{Id: aggregateID}, nil
}

func (g *grpcService) DepositBalance(ctx context.Context, request *bankAccountService.DepositBalanceRequest) (*bankAccountService.DepositBalanceResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "grpcService.DepositBalance")
	defer span.Finish()
	span.LogFields(log.String("req", request.String()))

	command := commands.DepositBalanceCommand{
		AggregateID: request.GetId(),
		Amount:      request.GetAmount(),
		PaymentID:   request.GetPaymentId(),
	}

	err := g.bankAccountService.Commands.DepositBalance.Handle(ctx, command)
	if err != nil {
		g.log.Errorf("(DepositBalance.Handle) err: %v", err)
		return nil, grpc_errors.ErrResponse(tracing.TraceWithErr(span, err))
	}

	return new(bankAccountService.DepositBalanceResponse), nil
}

func (g *grpcService) WithdrawBalance(ctx context.Context, request *bankAccountService.WithdrawBalanceRequest) (*bankAccountService.WithdrawBalanceResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "grpcService.WithdrawBalance")
	defer span.Finish()
	span.LogFields(log.String("req", request.String()))

	command := commands.WithdrawBalanceCommand{
		AggregateID: request.GetId(),
		Amount:      request.GetAmount(),
		PaymentID:   request.GetPaymentId(),
	}

	err := g.bankAccountService.Commands.WithdrawBalance.Handle(ctx, command)
	if err != nil {
		g.log.Errorf("(WithdrawBalance.Handle) err: %v", err)
		return nil, grpc_errors.ErrResponse(tracing.TraceWithErr(span, err))
	}

	return new(bankAccountService.WithdrawBalanceResponse), nil
}

func (g *grpcService) ChangeEmail(ctx context.Context, request *bankAccountService.ChangeEmailRequest) (*bankAccountService.ChangeEmailResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "grpcService.ChangeEmail")
	defer span.Finish()
	span.LogFields(log.String("req", request.String()))

	command := commands.ChangeEmailCommand{AggregateID: request.GetId(), NewEmail: request.GetEmail()}

	err := g.bankAccountService.Commands.ChangeEmail.Handle(ctx, command)
	if err != nil {
		g.log.Errorf("(ChangeEmail.Handle) err: %v", err)
		return nil, grpc_errors.ErrResponse(tracing.TraceWithErr(span, err))
	}

	return new(bankAccountService.ChangeEmailResponse), nil
}

func (g *grpcService) GetById(ctx context.Context, request *bankAccountService.GetByIdRequest) (*bankAccountService.GetByIdResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "grpcService.GetById")
	defer span.Finish()
	span.LogFields(log.String("req", request.String()))

	query := queries.GetBankAccountByIDQuery{AggregateID: request.GetId(), FromEventStore: request.IsOwner}
	bankAccountProjection, err := g.bankAccountService.Queries.GetBankAccountByID.Handle(ctx, query)
	if err != nil {
		g.log.Errorf("(GetBankAccountByID.Handle) err: %v", err)
		return nil, grpc_errors.ErrResponse(tracing.TraceWithErr(span, err))
	}

	g.log.Infof("GetById bankAccountProjection: %#v", bankAccountProjection)

	return &bankAccountService.GetByIdResponse{BankAccount: mappers.BankAccountMongoProjectionToProto(bankAccountProjection)}, nil
}

func (g *grpcService) SearchBankAccounts(ctx context.Context, request *bankAccountService.SearchBankAccountsRequest) (*bankAccountService.SearchBankAccountsResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "grpcService.SearchBankAccounts")
	defer span.Finish()
	span.LogFields(log.String("req", request.String()))

	query := queries.SearchBankAccountsQuery{
		QueryTerm: request.GetSearchText(),
		Pagination: &utils.Pagination{
			Size: int(request.GetSize()),
			Page: int(request.GetPage()),
		},
	}
	searchQueryResult, err := g.bankAccountService.Queries.SearchBankAccounts.Handle(ctx, query)
	if err != nil {
		g.log.Errorf("(SearchBankAccounts.Handle) err: %v", err)
		return nil, grpc_errors.ErrResponse(tracing.TraceWithErr(span, err))
	}

	g.log.Infof("SearchBankAccounts result: %#v", searchQueryResult.PaginationResponse)
	return &bankAccountService.SearchBankAccountsResponse{
		BankAccounts: mappers.SearchBankAccountsListToProto(searchQueryResult.List),
		Pagination:   mappers.PaginationResponseToProto(searchQueryResult.PaginationResponse),
	}, nil
}

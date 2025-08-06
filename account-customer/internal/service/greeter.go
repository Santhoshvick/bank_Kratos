package service

import (
	"context"

	pb "account-customer/api/helloworld/v1"
	"account-customer/internal/biz"
)

// GreeterService is a greeter service.
type AccountCustomerService struct {
	pb.UnimplementedAccountCustomerServer

	uc *biz.AccountCustomerUsecase
}

// NewGreeterService new a greeter service.
func NewAccountCustomerService(uc *biz.AccountCustomerUsecase) *AccountCustomerService {
	return &AccountCustomerService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *AccountCustomerService) CreateAccountCustomer(ctx context.Context, in *pb.CreateCustomerAccountRequest) (*pb.CreateCustomerAccountReply, error) {
	g, err := s.uc.CreateAccountCustomer(ctx,nil)
	if err != nil {
		return nil, err
	}

	return &pb.CreateCustomerAccountReply{
		AccountId:g.AccountId,
		CustomerId:g.CustomerId,

	},nil


}

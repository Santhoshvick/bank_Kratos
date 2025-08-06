package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)


// Greeter is a Greeter model.
type Rel struct {
	CustomerId string
	AccountId int64
}

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	CreateAccountCustomer(context.Context, *Rel) (*Rel, error)
}

// GreeterUsecase is a Greeter usecase.
type AccountCustomerUsecase struct {
	repo GreeterRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo GreeterRepo, logger log.Logger) *AccountCustomerUsecase {
	return &AccountCustomerUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *AccountCustomerUsecase) CreateAccountCustomer(ctx context.Context, g *Rel) (*Rel,error) {
	return uc.repo.CreateAccountCustomer(ctx, g)
}



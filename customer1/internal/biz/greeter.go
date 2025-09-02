package biz

import (
	"context"

	v1 "customer1/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type Customer struct {
	Name string
	ID string
	Email string
	Phone string
	Address string

}

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	Create(context.Context, *Customer) (*Customer, error)
	Fetch(context.Context, *Customer) (*Customer, error)
	FindByID(context.Context, string) (*Customer, error)
	ListByHello(context.Context, string) ([]*Customer, error)
	ListAll(context.Context) ([]*Customer, error)
}

// GreeterUsecase is a Greeter usecase.
type CustomerUsecase struct {
	repo GreeterRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo GreeterRepo, logger log.Logger) *CustomerUsecase {
	return &CustomerUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
// func (uc *CustomerUsecase) CreateGreeter(ctx context.Context, g *Customer) (*Customer, error) {
// 	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Address)
// 	return uc.repo.Fetch(ctx, g)
// }

func(uc *CustomerUsecase)CreateCustomer(ctx context.Context,g *Customer)(*Customer,error){
	return uc.repo.Create(ctx,g)
}

func(uc *CustomerUsecase)CreateGreeter(ctx context.Context,Id string)(*Customer,error){
	return uc.repo.FindByID(ctx,Id)
}



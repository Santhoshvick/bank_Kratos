package biz

import (
	"context"
	"time"

	
	"github.com/go-kratos/kratos/v2/log"
)



// Greeter is a Greeter model.
type Payment struct {
	PaymentId int64
	FromAccountId int64
	ToAccountId int64
	PaymentType string
	Amount string
	Status string
	Currency string
	PaymentMethod string
	ReferenceNumber string
	ExternalReference string
	Fee string
	ScheduledAt time.Time
	ProcessedAt time.Time
	CreatedAt time.Time
}

// GreeterRepo is a Greater repo.
type PaymentRepo interface {
	CreatePayment(context.Context, *Payment) (*Payment, error)
	UpdatePayment(context.Context, *Payment) (*Payment, error)
	DeletePayment(context.Context, int64) (*Payment, error)
	FindPayment(context.Context, int64) (*Payment, error)
}

type PaymentUsecase struct {
	repo PaymentRepo
	log  *log.Helper
}

func NewPaymentUsecase(repo PaymentRepo, logger log.Logger) *PaymentUsecase {
	return &PaymentUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *PaymentUsecase) CreatePayment(ctx context.Context, payment *Payment) (*Payment, error) {
	return uc.repo.CreatePayment(ctx, payment)
}

func (uc *PaymentUsecase) UpdatePayment(ctx context.Context, g *Payment) (*Payment, error) {
	return uc.repo.UpdatePayment(ctx, g)
}

func (uc *PaymentUsecase) DeletePayment(ctx context.Context,PaymentId int64) (*Payment, error) {
	return uc.repo.DeletePayment(ctx, PaymentId)
}

func (uc *PaymentUsecase) FindPayment(ctx context.Context, PaymentId int64) (*Payment, error) {
	return uc.repo.FindPayment(ctx,PaymentId)
}





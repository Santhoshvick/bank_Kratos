package biz

import (
	"context"
	"time"

	v1 "payment-service/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
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





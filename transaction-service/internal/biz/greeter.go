package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Greeter is a Greeter model.
type Transaction struct {
	TransactionId int64
	AccountId int64
	RelatedId int64
	TransactionType string
	Amount string
	Currency string
	Status string
	Description string
	ReferenceNumber string
	PostingDate string
	TransactionDate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

// GreeterRepo is a Greater repo.
type TransactionRepo interface {
	CreateTransaction(context.Context, *Transaction) (*Transaction, error)
	UpdateTransaction(context.Context, *Transaction) (*Transaction, error)
	DeleteTransaction(context.Context, int64) (*Transaction, error)
	FindTransaction(context.Context, int64) (*Transaction, error)
}

// GreeterUsecase is a Greeter usecase.
type TransactionUsecase struct {
	repo TransactionRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewTransactionUsecase(repo TransactionRepo, logger log.Logger) *TransactionUsecase {
	return &TransactionUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *TransactionUsecase) CreateTransaction(ctx context.Context, customer *Transaction) (*Transaction, error) {
	return uc.repo.CreateTransaction(ctx, customer)
}
func (uc *TransactionUsecase) UpdateTransaction(ctx context.Context, g *Transaction) (*Transaction, error) {
	return uc.repo.UpdateTransaction(ctx, g)
}
func (uc *TransactionUsecase) DeleteTransaction(ctx context.Context, TransactionId int64) (*Transaction, error) {
	return uc.repo.DeleteTransaction(ctx, TransactionId)
}

func (uc *TransactionUsecase) FindTransaction(ctx context.Context, TransactionId int64) (*Transaction, error) {
	return uc.repo.FindTransaction(ctx, TransactionId)
}


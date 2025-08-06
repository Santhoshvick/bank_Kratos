package biz

import (
	"context"
	"time"
	"github.com/go-kratos/kratos/v2/log"
)



// Greeter is a Greeter model.
type Card struct {
	CardId int64 `gorm:"column:card_id;type:bigint;primaryKey"`
	AccountNumber int64
	CardNumber string
	CardType string
	CardStatus string
	ExpiryDate string
	DailyLimit string
	MonthlyLimit string
	PinAttempt string	
	CreatedAt time.Time
	LastUsed time.Time
}

// GreeterRepo is a Greater repo.
type CardRepo interface {
	CreateCard(context.Context, *Card) (*Card, error)
	UpdateCard(context.Context, *Card) (*Card, error)
    DeleteCard(context.Context, int64) (*Card, error)
	FindCard(context.Context, int64) (*Card, error)
}


type CardUsecase struct {
	repo CardRepo
	log  *log.Helper
}

func NewCardUsecase(repo CardRepo, logger log.Logger) *CardUsecase {
	return &CardUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CardUsecase) CreateCard(ctx context.Context, g *Card) (*Card, error) {
	return uc.repo.CreateCard(ctx, g)
}

func (uc *CardUsecase) UpdateCard(ctx context.Context, g *Card) (*Card, error) {
	return uc.repo.UpdateCard(ctx, g)
}


func (uc *CardUsecase) DeleteCard(ctx context.Context, CardId int64) (*Card, error) {
	return uc.repo.DeleteCard(ctx,CardId)
}


func (uc *CardUsecase) FindCard(ctx context.Context, CardId int64) (*Card, error) {
	return uc.repo.FindCard(ctx, CardId)
}








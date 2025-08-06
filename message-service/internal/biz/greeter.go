package biz

import (
	"context"

	v1 "message-service/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type Notification struct {
	Email string
	OTP int64
	
}

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	CreateNotification(context.Context, *Notification) (*Notification, error)
	CreateTransactionNotification(context.Context, *Notification) (*Notification, error)
	
}

// GreeterUsecase is a Greeter usecase.
type NotificationUsecase struct {
	repo GreeterRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo GreeterRepo, logger log.Logger) *NotificationUsecase {
	return &NotificationUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *NotificationUsecase) CreateNotification(ctx context.Context, g *Notification) (*Notification, error) {
	uc.log.WithContext(ctx).Infof("Email: %v", g.Email)
	return uc.repo.CreateNotification(ctx, g)
}

func (uc *NotificationUsecase) CreateTransactionNotification(ctx context.Context, g *Notification) (*Notification, error) {
	uc.log.WithContext(ctx).Infof("Email: %v", g.Email)
	return uc.repo.CreateNotification(ctx, g)
}

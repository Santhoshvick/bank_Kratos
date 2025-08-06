package data

import (
	"context"

	"message-service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) CreateNotification(ctx context.Context, g *biz.Notification) (*biz.Notification, error) {
	return g, nil
}

func (r *greeterRepo) CreateTransactionNotification(ctx context.Context, g *biz.Notification) (*biz.Notification, error) {
	return g, nil
}


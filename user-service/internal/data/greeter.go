package data

import (
	"context"

	"user-service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
	table *gorm.DB
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
		table: data.db.Table("users"),
	}
}



func (r *UserRepo) CreateUser(ctx context.Context, g *biz.User) (*biz.User, error) {
	result := r.table.WithContext(ctx).Create(g)
	if result.Error != nil {
		return nil, result.Error
	}
	return g, nil
}

func (r *UserRepo) FindUser(ctx context.Context,Email string) (*biz.User, error) {
	var tx biz.User
	result := r.table.WithContext(ctx).Where("email = ?", Email).First(&tx)
	if result.Error != nil {
		return nil, result.Error
	}
	return &tx, nil
	
	
}

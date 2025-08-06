package data

import (
	"context"
	"fmt"

	"account-customer/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type accountCustomerRepo struct {
	data *Data
	log  *log.Helper
	table *gorm.DB
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &accountCustomerRepo{
		data: data,
		log:  log.NewHelper(logger),
		table: data.db.Table("rel_account_customer"),
	}
}

func (r *accountCustomerRepo) CreateAccountCustomer(ctx context.Context, g *biz.Rel) (*biz.Rel, error) {
     result := r.table.WithContext(ctx).Create(g)
	if result.Error != nil {
		return nil, fmt.Errorf("customer already exist in our bank")
	}

	if result.RowsAffected==0{
		return nil,fmt.Errorf("customer already exist in our bank")
	}
	return g, nil
}



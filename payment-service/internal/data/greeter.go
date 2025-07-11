package data

import (
	"context"
    "fmt"
	"payment-service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type PaymentRepo struct {
	data *Data
	log  *log.Helper
	table *gorm.DB
}

// NewGreeterRepo .
func NewPaymentRepo(data *Data, logger log.Logger) biz.PaymentRepo {
	return &PaymentRepo{
		data: data,
		log:  log.NewHelper(logger),
		table:data.db.Table("payment"),
	}
}

func (r *PaymentRepo) CreatePayment(ctx context.Context, g *biz.Payment) (*biz.Payment, error) {
	result := r.table.WithContext(ctx).Create(g)
	if result.Error != nil {
		return nil, result.Error
	}
	return g, nil
}

func (r *PaymentRepo) UpdatePayment(ctx context.Context, g *biz.Payment) (*biz.Payment, error) {
	result:=r.table.Model(g).Where("payment_id = ?", g.PaymentId).Updates(biz.Payment{PaymentId:g.PaymentId,PaymentType: g.PaymentType,Amount: g.Amount,PaymentMethod: g.PaymentMethod,Currency: g.Currency,FromAccountId: g.FromAccountId,ToAccountId: g.ToAccountId,})
	if result.Error!=nil{
		return nil,result.Error
	}
	return g, nil
}

func (r *PaymentRepo) DeletePayment(ctx context.Context, PaymentId int64) (*biz.Payment, error) {
	result := r.table.WithContext(ctx).Where("payment_id = ?", PaymentId).Delete(&biz.Payment{})
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.Payment{}, nil
}

func (r *PaymentRepo) FindPayment(ctx context.Context,PaymentId int64) (*biz.Payment, error) {
	var tx biz.Payment
	result := r.table.WithContext(ctx).Where("payment_id = ?", PaymentId).First(&tx)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("payment with ID %d not found", PaymentId)
	}
	return &tx, nil
}



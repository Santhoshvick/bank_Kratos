package data

import (
	"context"
    "fmt"
	"transaction-service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type TransactionRepo struct {
	data *Data
	log  *log.Helper
	table  *gorm.DB
}

// NewGreeterRepo .
func NewTransactionRepo(data *Data, logger log.Logger) biz.TransactionRepo {
	return &TransactionRepo{
		data: data,
		log:  log.NewHelper(logger),
		table:data.db.Table("transaction"),
	}
}

func (r *TransactionRepo) CreateTransaction(ctx context.Context, g *biz.Transaction) (*biz.Transaction, error) {
	result := r.table.WithContext(ctx).Create(g)
	if result.Error != nil {
		return nil, result.Error
	}
	return g, nil
	
}

func (r *TransactionRepo) UpdateTransaction(ctx context.Context, g *biz.Transaction) (*biz.Transaction, error) {
	result:=r.table.Model(g).Where("transaction_id = ?", g.TransactionId).Updates(biz.Transaction{TransactionId: g.TransactionId,TransactionType: g.TransactionType,Description: g.Description,Amount: g.Amount,Currency: g.Currency})
	if result.Error!=nil{
		return nil,result.Error
	}
	return g, nil
}

func (r *TransactionRepo) DeleteTransaction(ctx context.Context, TransactionId int64) (*biz.Transaction, error) {
	var tx biz.Transaction
	result := r.table.WithContext(ctx).Where("transaction_id = ?", TransactionId).Delete(&biz.Transaction{})
	if result.Error != nil {
		return nil, result.Error
	}
	return &tx, nil
}

func (r *TransactionRepo) FindTransaction(ctx context.Context, TransactionId int64) (*biz.Transaction, error) {
	var tx biz.Transaction

	result := r.table.WithContext(ctx).Where("transaction_id = ?", TransactionId).First(&tx) 
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("payment with ID %d not found", TransactionId)
	}
	return &tx, nil
}


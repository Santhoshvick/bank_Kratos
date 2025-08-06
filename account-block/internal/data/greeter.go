package data

import (
	"context"
	"fmt"

	pb "account-block/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)
type AccountBlock struct {
	AccountId   int64  `gorm:"column:account_id;primaryKey"`
	AccountNumber int64 `gorm:"column:account_number"`
	AccountBlock string `gorm:"column:account_block"`
}

type GreeterRepo struct {
	data *Data
	log  *log.Helper
	table *gorm.DB

}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) *GreeterRepo {
	return &GreeterRepo{
		data: data,
		log:  log.NewHelper(logger),
		table:data.DB.Table("account_block"),
	}
}

func (r *GreeterRepo) CreateAccountBlock(ctx context.Context, g *pb.BlockRequest) error {
	account := &AccountBlock{
		AccountId:   g.AccountId,
		AccountNumber: g.AccountNumber,
		AccountBlock: g.Accblock.String(),
	}
	result := r.table.WithContext(ctx).Create(account)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GreeterRepo) AdminAccountBlock(ctx context.Context, g *pb.BlockRequest) error {
	account := &AccountBlock{
		AccountId:   g.AccountId,
		AccountNumber: g.AccountNumber,
		AccountBlock: g.Accblock.String(),
	}
	result := r.table.WithContext(ctx).Create(account)
	if result.Error != nil {
		return result.Error
	}
	return nil
}


func (r *GreeterRepo) FindAccountBlock(ctx context.Context, g *pb.FindBlockRequest) (*pb.FindBlockResponse, error) {
	var tx AccountBlock

	err := r.table.WithContext(ctx).Where("account_number = ?", g.AccountNumber).First(&tx).Error
	if err != nil {
		return nil, fmt.Errorf("DB error: %w", err)
	}
	return &pb.FindBlockResponse{
		AccountNumber: tx.AccountNumber,
		AccountId:     tx.AccountId, 
	}, nil
}
func (r *GreeterRepo) UpdateAccountBlock(ctx context.Context, g *pb.UpdateBlockRequest) (*pb.UpdateBlockResponse, error) {
	var tx AccountBlock
	acc:=AccountBlock{
		AccountBlock: tx.AccountBlock,
	}

	err := r.table.WithContext(ctx).Where("account_number = ?", g.AccountNumber).Updates(acc).Error
	if err != nil {
		return nil, fmt.Errorf("DB error: %w", err)
	}

	return &pb.UpdateBlockResponse{
		AccountNumber: tx.AccountNumber,
		AccountId:     tx.AccountId, 
	}, nil
}


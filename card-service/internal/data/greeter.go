package data

import (
	"context"

	"card-service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type CardRepo struct {
	data  *Data
	log   *log.Helper
	table *gorm.DB
}

// NewGreeterRepo .
func NewCardRepo(data *Data, logger log.Logger) biz.CardRepo {
	return &CardRepo{
		data:  data,
		log:   log.NewHelper(logger),
		table: data.db.Table("card"),
	}
}

func (r *CardRepo) CreateCard(ctx context.Context, g *biz.Card) (*biz.Card, error) {
	result := r.table.WithContext(ctx).Create(g)
	if result.Error != nil {
		return nil, result.Error
	}
	return g, nil
}

func (r *CardRepo) UpdateCard(ctx context.Context, g *biz.Card) (*biz.Card, error) {
	result := r.table.Model(g).Where("card_id = ?", g.CardId).Updates(biz.Card{CardId: g.CardId, CardNumber: g.CardNumber, CardType: g.CardType, CardStatus: g.CardStatus, DailyLimit: g.DailyLimit})
	if result.Error != nil {
		return nil, result.Error
	}
	return g, nil
}

func (r *CardRepo) DeleteCard(ctx context.Context, CardId int64) (*biz.Card, error) {
	result := r.table.WithContext(ctx).Where("card_id = ?", CardId).Delete(&biz.Card{})
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.Card{}, nil
}

func (r *CardRepo) FindCard(ctx context.Context, CardId int64) (*biz.Card, error) {
	var tx biz.Card
	result := r.table.WithContext(ctx).Where("card_id = ?", CardId).First(&tx)
	if result.Error != nil {
		return nil, result.Error
	}
	return &tx, nil
}

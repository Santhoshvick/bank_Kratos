package data

import (
	"context"

	"rule-engine1/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type RuleRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewRuleRepo(data *Data, logger log.Logger) biz.RuleRepo {
	return &RuleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *RuleRepo) RuleEngine(ctx context.Context, g *biz.Rule) (*biz.Rule, error) {
	return g, nil
}

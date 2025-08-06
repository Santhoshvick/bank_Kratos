package biz

import (
	"context"

	v1 "rule-engine1/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type Rule struct {
	Description string
	CustomerId int64
	AvailableBalance string
}

// GreeterRepo is a Greater repo.
type RuleRepo interface {
	RuleEngine(context.Context, *Rule) (*Rule, error)
}

// GreeterUsecase is a Greeter usecase.
type RuleUsecase struct {
	repo RuleRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewRuleUsecase(repo RuleRepo, logger log.Logger) *RuleUsecase {
	return &RuleUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *RuleUsecase) RuleEngine(ctx context.Context, g *Rule) (*Rule, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Description)
	return uc.repo.RuleEngine(ctx, g)
}

package biz

import (
	"context"

	v1 "user-service/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type User struct {
	UserId int64 `gorm:"column:user_id;type:bigint;primaryKey"`
	FirstName string
	LastName string
	DateOfBirth string
	Nationality string
	Email string
	Phone int64
	Address1 string
	Address2 string
	City string
	State string
	PostalCode string
	Country string
	UserName string
	Password string
}

// GreeterRepo is a Greater repo.
type UserRepo interface {
	CreateUser(context.Context, *User) (*User, error)
	FindUser(context.Context,string)(*User,error)
	
}

// GreeterUsecase is a Greeter usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *UserUsecase) CreateUser(ctx context.Context, g *User) (*User, error) {

	return uc.repo.CreateUser(ctx, g)
}

func (uc *UserUsecase) FindUser(ctx context.Context, Email string) (*User, error) {
	return uc.repo.FindUser(ctx, Email)
}

package data

import (
	"context"
	"fmt"

	"customer1/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type CustomerRepo struct {
	data *Data
	log  *log.Helper
	table *gorm.DB
}

// NewGreeterRepo .
func NewCustomerRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &CustomerRepo{
		data: data,
		log:  log.NewHelper(logger),
		table: data.db.Table("customer1"),
	}
}

func (r *CustomerRepo) Create(ctx context.Context, g *biz.Customer) (*biz.Customer, error) {
	result := r.table.WithContext(ctx).Create(g)
	fmt.Println(result)
	return g, nil
}

func (r *CustomerRepo) Fetch(ctx context.Context, g *biz.Customer) (*biz.Customer, error) {
	return g,nil
	
}

func (r *CustomerRepo) FindByID(ctx context.Context, Id string) (*biz.Customer, error) {
	var customer biz.Customer
	// id1:="1001"
	fmt.Println("Id Number:",Id)
	result := r.table.WithContext(ctx).Where("id = ?",Id).First(&customer)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}


func (r *CustomerRepo) ListByHello(context.Context, string) ([]*biz.Customer, error) {
	return nil, nil
}

func (r *CustomerRepo) ListAll(context.Context) ([]*biz.Customer, error) {
	return nil, nil
}

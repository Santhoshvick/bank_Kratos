package service

import (
	"context"
	"fmt"

	v1 "customer1/api/helloworld/v1"
	"customer1/internal/biz"
)

// GreeterService is a greeter service.
type CustomerService struct {
	v1.UnimplementedGreeterServer
	uc *biz.CustomerUsecase
}

// NewGreeterService new a greeter service.
func NewCustomerService(uc *biz.CustomerUsecase) *CustomerService {
	return &CustomerService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *CustomerService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	fmt.Println("Id1:",in.Id)
	g, err := s.uc.CreateGreeter(ctx, in.Id)
	if err != nil { 
		return nil, err
	}

	return &v1.HelloReply{
		Id: g.ID,
		Name: g.Name,
		Address: g.Address,
		Email: g.Email,
		Phone: g.Phone,
	}, nil
}


func(s * CustomerService)CreateCustomer(ctx context.Context,in *v1.CreateCustomerRequest)(*v1.CreateCustomerResponse,error){
	g,err:=s.uc.CreateCustomer(ctx,&biz.Customer{
		Name: in.Name,
		ID: in.Id,
		Email: in.Email,
		Address: in.Address,
		Phone: in.Phone,
	})  

	if err!=nil{
            return nil, err
	}

	return &v1.CreateCustomerResponse{
		Id: g.ID,
		Name:g.Name,
		Email: g.Email,
		Address: g.Address,
		Phone: g.Phone,
	},err
}

func(s *CustomerService)DeleteCustomer(ctx context.Context,in *v1.DeleteCustomerRequest)(*v1.DeleteCustomerResponse,error){
	_,err:=s.uc.DelCustomer(ctx,in.Id)
	if err!=nil{
		return nil,err
	}
	return &v1.DeleteCustomerResponse{
		Message: "The customer for the mentioned customer id is deleted successfully",
	},err
}


func(s *CustomerService)UpdateCustomer(ctx context.Context,in *v1.UpdateCustomerRequest)(*v1.UpdateCustomerResponse,error){
	g,err:=s.uc.UpdCustomer(ctx, &biz.Customer{
		ID: in.Id,
		Name: in.Name,
		Email: in.Email,
		Phone: in.Phone,
		Address: in.Address,
	})
	if err!=nil{
		return nil,err
	}

	return &v1.UpdateCustomerResponse{
		Id: g.ID,
		Name:g.Name,
		Address:g.Address,
		Email:g.Email,
		Phone:g.Phone,
	},err
}

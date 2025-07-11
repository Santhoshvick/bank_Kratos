package service

import (
	"context"

	pb "payment-service/api/helloworld/v1"
	"payment-service/internal/biz"
)

// GreeterService is a greeter service.
type PaymentService struct {
	pb.UnimplementedPaymentServer

	uc *biz.PaymentUsecase
}

// NewGreeterService new a greeter service.
func NewPaymentService(uc *biz.PaymentUsecase) *PaymentService {
	return &PaymentService{uc: uc}
}

func (s *PaymentService) CreatePayment(ctx context.Context,req *pb.CreateRequest) (*pb.CreateResponse,error){
	g,err:=s.uc.CreatePayment(ctx,&biz.Payment{
		PaymentId:req.PaymentId,
		FromAccountId:req.FromAccountId,
		ToAccountId:req.ToAccountId,
		PaymentType:req.PaymentType,
		Currency: req.Currency,
		Amount:req.Amount,
		Status:req.Status,
		PaymentMethod:req.PaymentMethod,
		ReferenceNumber:req.ReferenceNumber,
	})
	if err!=nil{
		return nil,nil
	}

	return &pb.CreateResponse{
		PaymentId:g.PaymentId,
		FromAccountId:g.FromAccountId,
		ToAccountId:g.ToAccountId,
		Currency: req.Currency,
		PaymentType:g.PaymentType,
		Amount:g.Amount,
		Status:g.Status,
		PaymentMethod:g.PaymentMethod,
		ReferenceNumber:g.ReferenceNumber,

	},nil
}


func (s *PaymentService) UpdatePayment(ctx context.Context,req *pb.UpdateRequest) (*pb.UpdateResponse,error){
	g,err:=s.uc.UpdatePayment(ctx,&biz.Payment{
		PaymentId:req.PaymentId,
		FromAccountId:req.FromAccountId,
		ToAccountId:req.ToAccountId,
		Currency: req.Currency,
		PaymentType:req.PaymentType,
		Amount:req.Amount,
		Status:req.Status,
		PaymentMethod:req.PaymentMethod,
		ReferenceNumber:req.ReferenceNumber,
	})
	if err!=nil{
		return nil,nil
	}

	return &pb.UpdateResponse{
		PaymentId:g.PaymentId,
		FromAccountId:g.FromAccountId,
		ToAccountId:g.ToAccountId,
		Currency: req.Currency,
		PaymentType:g.PaymentType,
		Amount:g.Amount,
		Status:g.Status,
		PaymentMethod:g.PaymentMethod,
		ReferenceNumber:g.ReferenceNumber,

	},nil
}


func (s *PaymentService) DeletePayment(ctx context.Context,req *pb.DeleteRequest) (*pb.DeleteResponse,error){
	g,err:=s.uc.DeletePayment(ctx,req.PaymentId)
	if err!=nil{
		return nil,nil
	}

	return &pb.DeleteResponse{
		PaymentId:g.PaymentId,
		FromAccountId:g.FromAccountId,
		Currency: req.Currency,
		ToAccountId:g.ToAccountId,
		PaymentType:g.PaymentType,
		Amount:g.Amount,
		Status:g.Status,
		PaymentMethod:g.PaymentMethod,
		ReferenceNumber:g.ReferenceNumber,

	},nil
}

func (s *PaymentService) FindPayment(ctx context.Context,req *pb.FindRequest) (*pb.FindResponse,error){
	g,err:=s.uc.FindPayment(ctx,req.PaymentId)
	if err!=nil{
		return nil,nil
	}

	return &pb.FindResponse{
		PaymentId:g.PaymentId,
		FromAccountId:g.FromAccountId,
		Currency: req.Currency,
		ToAccountId:g.ToAccountId,
		PaymentType:g.PaymentType,
		Amount:g.Amount,
		Status:g.Status,
		PaymentMethod:g.PaymentMethod,
		ReferenceNumber:g.ReferenceNumber,

	},nil
}


package service

import (
	"context"

	pb "transaction-service/api/helloworld/v1"
	"transaction-service/internal/biz"
)

// GreeterService is a greeter service.
type TransactionService struct {
	pb.UnimplementedTransactionServer

	uc *biz.TransactionUsecase
}

func NewTransactionService(uc *biz.TransactionUsecase) *TransactionService {
	return &TransactionService{uc: uc}
}


func (s *TransactionService) CreateTransaction(ctx context.Context,req *pb.CreateRequest) (*pb.CreateResponse,error){
	g,err:=s.uc.CreateTransaction(ctx,&biz.Transaction{
		TransactionId:req.TransactionId,
		AccountId:req.AccountId,
		RelatedId:req.RelatedId,
		TransactionType:req.TransactionType,
		Amount:req.Amount,
		Currency:req.Currency,
		Status:req.Status,
		Description:req.Description,
		ReferenceNumber:req.ReferenceNumber,
		PostingDate:req.PostingDate,
	})
	if err!=nil{
		return nil,err
	}
	return &pb.CreateResponse{
		TransactionId:g.TransactionId,
		AccountId:g.AccountId,
		RelatedId:g.RelatedId,
		TransactionType:g.TransactionType,
		Amount:g.Amount,
		Currency:g.Currency,
		Status:g.Status,
		Description:g.Description,
		ReferenceNumber:g.ReferenceNumber,
		PostingDate:g.PostingDate,
	},nil
}


func (s *TransactionService) UpdateTransaction(ctx context.Context,req *pb.UpdateRequest) (*pb.UpdateResponse,error){
	g,err:=s.uc.UpdateTransaction(ctx,&biz.Transaction{
		TransactionId:req.TransactionId,
		AccountId:req.AccountId,
		RelatedId:req.RelatedId,
		TransactionType:req.TransactionType,
		Amount:req.Amount,
		Currency:req.Currency,
		Status:req.Status,
		Description:req.Description,
		ReferenceNumber:req.ReferenceNumber,
		PostingDate:req.PostingDate,
	})
	if err!=nil{
		return nil,err
	}
	return &pb.UpdateResponse{
		TransactionId:g.TransactionId,
		AccountId:g.AccountId,
		RelatedId:g.RelatedId,
		TransactionType:g.TransactionType,
		Amount:g.Amount,
		Currency:g.Currency,
		Status:g.Status,
		Description:g.Description,
		ReferenceNumber:g.ReferenceNumber,
		PostingDate:g.PostingDate,
	},nil
}


func (s *TransactionService) DeleteTransaction(ctx context.Context,req *pb.DeleteRequest) (*pb.DeleteResponse,error){
	g,err:=s.uc.DeleteTransaction(ctx,req.TransactionId)
	if err!=nil{
		return nil,err
	}
	return &pb.DeleteResponse{
		TransactionId:g.TransactionId,
		AccountId:g.AccountId,
		RelatedId:g.RelatedId,
		TransactionType:g.TransactionType,
		Amount:g.Amount,
		Currency:g.Currency,
		Status:g.Status,
		Description:g.Description,
		ReferenceNumber:g.ReferenceNumber,
		PostingDate:g.PostingDate,
	},nil
}


func (s *TransactionService) FindTransaction(ctx context.Context,req *pb.FindRequest) (*pb.FindReponse,error){
	g,err:=s.uc.FindTransaction(ctx,req.TransactionId)
	if err!=nil{
		return nil,err
	}
	return &pb.FindReponse{
		TransactionId:g.TransactionId,
		AccountId:g.AccountId,
		RelatedId:g.RelatedId,
		TransactionType:g.TransactionType,
		Amount:g.Amount,
		Currency:g.Currency,
		Status:g.Status,
		Description:g.Description,
		ReferenceNumber:g.ReferenceNumber,
		PostingDate:g.PostingDate,
	},nil
}








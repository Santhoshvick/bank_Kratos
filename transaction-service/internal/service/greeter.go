package service

import (
	"context"
	"encoding/json"
	"fmt"
	// "log"

	paymetpb "payment-service/api/helloworld/v1"
	pb "transaction-service/api/helloworld/v1"
	"transaction-service/internal/biz"

	"github.com/nats-io/nats.go"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/credentials/insecure"
)

// GreeterService is a greeter service.
type TransactionService struct {
	pb.UnimplementedTransactionServer
    paymentclient paymetpb.PaymentClient
	uc *biz.TransactionUsecase
}

type TransactionMsg struct{
	TransactionId int64
	Amount string
	Currency string
	ReferenceNumber string
}
func NewTransactionService(uc *biz.TransactionUsecase,paymentClient paymetpb.PaymentClient) *TransactionService {
	return &TransactionService{uc: uc,
		paymentclient: paymentClient,
	}
}

func connectToNats()(*nats.Conn,error){
	return nats.Connect(nats.DefaultURL)
}

func (s *TransactionService) CreateTransaction(ctx context.Context,req *pb.CreateRequestTransaction) (*pb.CreateResponseTransaction,error){

	nc,_:=connectToNats()
	val:=TransactionMsg{
		TransactionId: req.TransactionId,
		Amount: req.Amount,
		ReferenceNumber: req.ReferenceNumber,
		Currency: req.Currency,
	}
	data, err := json.Marshal(val)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// conn1,err:=grpc.NewClient("localhost:9000",grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err!=nil{
	// 	log.Fatalf("GRPC connection is not establishded %v",err)
	// 	return nil,err
	// }
	// client1:=paymetpb.NewPaymentClient(conn1)

	// cust,err:=client1.FindPayment(ctx,&paymetpb.FindRequest{PaymentId: req.TransactionId})
	// if err!=nil{
	// 	return nil,err
	// }

	nc.Publish("Helloworld",data)
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
	fmt.Println("Transaction Successful")
	if err!=nil{
		return &pb.CreateResponseTransaction{
		   Message: "Transaction Details is not created",
	},nil
	}
	return &pb.CreateResponseTransaction{
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
		Message: "Trasaction is created Successfully",
	},nil
}


func (s *TransactionService) UpdateTransaction(ctx context.Context,req *pb.UpdateRequestTransaction) (*pb.UpdateResponseTransaction,error){
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
		return &pb.UpdateResponseTransaction{
			Message: "Updation Failed",
		},nil
	}
	return &pb.UpdateResponseTransaction{
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
func (s *TransactionService) DeleteTransaction(ctx context.Context,req *pb.DeleteRequestTransaction) (*pb.DeleteResponseTransaction,error){
	g,err:=s.uc.DeleteTransaction(ctx,req.TransactionId)
	if err!=nil{
		return nil,err
	}
	return &pb.DeleteResponseTransaction{
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


func (s *TransactionService) FindTransaction(ctx context.Context,req *pb.FindRequestTransaction) (*pb.FindReponseTransaction,error){
	g,err:=s.uc.FindTransaction(ctx,req.TransactionId)
	if err!=nil{
		return &pb.FindReponseTransaction{
		Message: "Transaction details not find please verify one the transaction is happend or not",
	},nil
		
	}
	return &pb.FindReponseTransaction{
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








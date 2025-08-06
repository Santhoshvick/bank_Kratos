package service

import (
	accountpb "account-service/api/helloworld/v1"
	"context"
	customerpb "customer-service/api/helloworld/v1"
	"encoding/json"
	"fmt"
	"log"
	 
	pb "payment-service/api/helloworld/v1"
	"payment-service/internal/biz"
	"payment-service/internal/handler"
	"strconv"
	"time"
	transactionpb "transaction-service/api/helloworld/v1"

	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GreeterService is a greeter service.
type PaymentService struct {
	pb.UnimplementedPaymentServer
	accountClient accountpb.AccountClient
	// customerClient customerpb.CustomerClient
	transactionclient transactionpb.TransactionClient
	uc *biz.PaymentUsecase
}

type FindPaymentRedis struct{
	PaymentId int64
	FromAccountId int64
	ToAccountId int64
	PaymentType string
	Amount string
	Status string
	Currency string
	PaymentMethod string
	ReferenceNumber string
	ExternalReference string
}

type PaymentStatus struct{
	FromEmail string
	ToEmail string
	FromBalance string
	ToBalance string
	FromFirstName string
	ToFirstName string
	Amount string
	FromAccountNumber int64
	ToAccountNumber int64
}

// NewGreeterService new a greeter service.
func NewPaymentService(uc *biz.PaymentUsecase,accClient accountpb.AccountClient,transactionClient transactionpb.TransactionClient) *PaymentService {
	return &PaymentService{uc: uc,
	accountClient: accClient,
	transactionclient: transactionClient,
}
}

func connectToNATS() (*nats.Conn, error) {
	fmt.Println("Nats is connected successfully")
    return nats.Connect(nats.DefaultURL)
}

func (s *PaymentService) CreatePayment(ctx context.Context,req *pb.CreateRequest) (*pb.CreateResponse,error){

	//Payment Method Validations
	if err:=handler.PaymentMethod(req.PaymentMethod);err!=nil{
		return &pb.CreateResponse{
			Message: ""+err.Error(),
		},nil
	}
	//Payment Type Validation
	if err:=handler.PaymentType(req.PaymentType);err!=nil{
		return &pb.CreateResponse{
			Message:""+err.Error(),
		},nil
	}

	//Currency Validation
	if err:=handler.CurrencyValidation(req.Currency);err!=nil{
		return &pb.CreateResponse{
			Message:""+err.Error(),
		},nil
	}
	// From Account Number Validation
	if err:=handler.AccountNumberValidation(req.FromAccountId);err!=nil{
		return &pb.CreateResponse{
			Message: ""+err.Error(),
		},nil
	}

	// //To account Number Validation
	if err:=handler.AccountNumberValidation(req.ToAccountId);err!=nil{
		return &pb.CreateResponse{
			Message: ""+err.Error(),
		},nil
	}

	
	//Account Number Random Generator
	randomNumber:=handler.GenerateRefNumber()


	//Transaction GRPC Client
	conn3,err:=grpc.NewClient("localhost:9005",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.Fatalf("GRPC connection is not establishded %v",err)
		return nil,err
	}
	client3:=transactionpb.NewTransactionClient(conn3)

	trans1,err:=client3.CreateTransaction(ctx,&transactionpb.CreateRequestTransaction{TransactionId: req.PaymentId,AccountId: req.FromAccountId,RelatedId: req.ToAccountId,Amount: req.Amount,TransactionType: req.PaymentType,Currency: req.Currency,Status: req.Status,Description:"transaction details is stored successfully",ReferenceNumber: req.ReferenceNumber,PostingDate: time.DateTime})
	if err!=nil{
		return &pb.CreateResponse{
			Message: "There is an issue with creating a transaction",
		},nil
	}
	fmt.Println(trans1)
	nc,_:=connectToNATS()
	nc.Subscribe("Helloworld",func(msg *nats.Msg) {
		fmt.Println("Payment is transfered successfully",string(msg.Data))
	})

	fmt.Println("---------ACCOUNT SERVICE-----------")
	conn, err := grpc.NewClient("localhost:9002", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Printf("GRPC Connection is not established %v", err)
        return nil, err
    }
	client:=accountpb.NewAccountClient(conn)
	val1,err1:=client.FindAccountNumber(ctx,&accountpb.FindRequest{AccountNumber: req.FromAccountId})
	val2,err2:=client.FindAccountNumber(ctx,&accountpb.FindRequest{AccountNumber: req.ToAccountId})


    fmt.Println("-------------------Account---------------")
	fmt.Println(val1,val2)
	fmt.Println("--------------------Account Ends----------")
	if err1!=nil||err2!=nil{
		log.Println("Account number doesnot exist")
		return nil,fmt.Errorf("fromAccountId or To accountID does not exist %v",err1)
	}
	fmt.Println(req.ToAccountId)
    fmt.Println("---------CUSTOMER SERVICE-----------")
	fmt.Println("Customer Id1:",val1.CustomerId)
	fmt.Println("Customer Id2:",val2.CustomerId)
	conn1,err:=grpc.NewClient("localhost:9000",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.Fatalf("GRPC connection is not establishded %v",err)
		return nil,err
	}
	client1:=customerpb.NewCustomerClient(conn1)
	id1,err3:=strconv.ParseInt(val1.CustomerId,10,64)
	if err3!=nil{
		log.Fatalf("Cannot convert string to int %v",err3)
	}

	id2,err4:=strconv.ParseInt(val2.CustomerId,10,64)
	if err4!=nil{
		log.Fatalf("Cannot convert string to int %v",err4)
	}
	cust1,err:=client1.FindCustomer(ctx,&customerpb.FindRequest{CustomerId: int64(id1)})
	if err!=nil{
		log.Fatalf("Customer not found %v",err)
	}

	cust2,err:=client1.FindCustomer(ctx,&customerpb.FindRequest{CustomerId: int64(id2)})
	if err!=nil{
		log.Fatalf("Customer not found %v",err)
	}
	fmt.Println(cust1,cust2)
	amount1:=req.Amount

	if (amount1>val1.CreditLimit){
		return &pb.CreateResponse{
              Message: "Amount Couldn't your transfer is greater than your credit limit,Please ensure the amount should be with in the credit limit",
		},nil
	}
	amount2:=val1.AvailableBalance
	amount3:=val2.AvailableBalance
     if(amount2<amount1){
	return &pb.CreateResponse{
              Message: " insufficient Balance",
	},nil
}
	fmt.Println("Amount1",amount1)
	fmt.Println("Amount2",amount2)
	fmt.Println("Account3",amount3)

	newAmount1,err:=strconv.ParseFloat(amount1,64)
	if err!=nil{
		log.Fatalf("Cannot convert string to int %v ",err)
	}

	fees:=(newAmount1*0.2)/100

	newAmount2,err:=strconv.ParseFloat(amount2,64)
	if err!=nil{
		log.Fatalf("Cannot convert string to int %v ",err)
	}

	newAmount3,err:=strconv.ParseFloat(amount3,64)
	if err!=nil{
		log.Fatalf("Cannot convert string to int %v ",err)
	}



	fromBalance:=newAmount2
	toBalance:=newAmount3
	
	 if(newAmount2>=newAmount1){
		fromBalance=newAmount2-newAmount1-fees
	 }else{
		log.Fatalf("Insufficient Balance")
	 }

		toBalance=newAmount3+newAmount1

	
	fmt.Println("FromBalance",fromBalance)
	fmt.Println("ToBalance",toBalance)

	

	    
	amt1,err3:=client.UpdateAccount(ctx,&accountpb.UpdateRequest{AccountNumber:req.FromAccountId,AvailableBalance:fmt.Sprintf("%.2f", fromBalance)})
	amt2,err4:=client.UpdateAccount(ctx,&accountpb.UpdateRequest{AccountNumber:req.ToAccountId,AvailableBalance: fmt.Sprintf("%.2f", toBalance)})
	if err3!=nil && err4!=nil{
		log.Fatalf("There is an issue with Update account %v %v",err2,err4)
	}

	fmt.Println("----------------Update Account----------------------")
	fmt.Println(amt1,amt2)
	fmt.Println("----------Update Account----------------------")

	balanceData:=PaymentStatus{
		FromEmail:cust1.Email,
		ToEmail: cust2.Email,
		FromBalance: amt1.AvailableBalance,
		ToBalance: amt2.AvailableBalance,
		Amount:req.Amount,
		FromAccountNumber: val1.AccountNumber,
		ToAccountNumber: val2.AccountNumber,
	}
	fmt.Println(balanceData.FromBalance,balanceData.ToBalance)
	data1, err := json.Marshal(balanceData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %w", err)
	}
	fmt.Println(val1.AccountNumber,val1.AccountType)
	fmt.Println(val2.AccountNumber,val2.AccountType)
	nc.Publish("Real",data1)
	g,err:=s.uc.CreatePayment(ctx,&biz.Payment{
		PaymentId:req.PaymentId,
		FromAccountId:req.FromAccountId,
		ToAccountId:req.ToAccountId,
		PaymentType:req.PaymentType,
		Currency: req.Currency,
		Amount:req.Amount,
		Status:req.Status,
		PaymentMethod:req.PaymentMethod,
	})
	if err!=nil{
		return &pb.CreateResponse{
			Message:"Payment is Not successfully Please Verify the payment Details Once ",
		
	},nil	
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
		ReferenceNumber:randomNumber,
		Message: "Payment Done Successfully",
		Fee:fmt.Sprintf("%.2f", fees),
	},nil
}


func (s *PaymentService) UpdatePayment(ctx context.Context,req *pb.UpdateRequest) (*pb.UpdateResponse,error){
	//From Account Number Validation
	if err:=handler.AccountNumberValidation(req.FromAccountId);err!=nil{
		return &pb.UpdateResponse{
			Message: ""+err.Error(),
		},nil
	}

	//To account Number Validation
	if err:=handler.AccountNumberValidation(req.ToAccountId);err!=nil{
		return &pb.UpdateResponse{
			Message: ""+err.Error(),
		},nil
	}
	//Payment Method Validations
	if err:=handler.PaymentMethod(req.PaymentMethod);err!=nil{
		return &pb.UpdateResponse{
			Message: ""+err.Error(),
		},nil
	}
	//Payment Type Validation
	if err:=handler.PaymentType(req.PaymentType);err!=nil{
		return &pb.UpdateResponse{
			Message:""+err.Error(),
		},nil
	}

	//Currency Validation
	if err:=handler.CurrencyValidation(req.Currency);err!=nil{
		return &pb.UpdateResponse{
			Message:""+err.Error(),
		},nil
		}

	
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
		return &pb.UpdateResponse{
		Message: "Payment Updation Failed",
	},nil
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
		Message: "Payment Details is updated successfully",
	},nil
}


func (s *PaymentService) DeletePayment(ctx context.Context,req *pb.DeleteRequest) (*pb.DeleteResponse,error){
	g,err:=s.uc.DeletePayment(ctx,req.PaymentId)
	if err!=nil{
		return &pb.DeleteResponse{
		Message: "Payment Details Not Found",
	},nil
		
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
		Message: "Payment details is deleted Successfully",
	},nil
}

func (s *PaymentService) FindPayment(ctx context.Context,req *pb.FindRequest) (*pb.FindResponse,error){
	g,err:=s.uc.FindPayment(ctx,req.PaymentId)
	if err!=nil{
		return &pb.FindResponse{
		Message: "Payment details not found",
	},nil
		
	}
    ctx = context.Background()
	rdb:=redis.NewClient(&redis.Options{
		Addr:"localhost:6379",
		DB:0,
	})
	paymentRedis:=FindPaymentRedis{
		PaymentId: req.PaymentId,
		FromAccountId: req.FromAccountId,
		ToAccountId: req.ToAccountId,
		PaymentType: req.PaymentType,
		Amount: req.Amount,
		Status: req.Status,
		Currency: req.Currency,
		PaymentMethod: req.PaymentMethod,
		ReferenceNumber: req.ReferenceNumber,
		ExternalReference: req.ReferenceNumber,
	}
	PaymentMarshall,err:=json.Marshal(paymentRedis)
	if err!=nil{
		log.Fatalf("Issue with converting the go data to the json data fromat %v",err)
	}

	var timeout=100000*time.Second

	paymentFindKey,err:=rdb.Set(ctx,"Payment1",string(PaymentMarshall),timeout).Result()
	if err!=nil{
		log.Fatalf("There is an issue with Setting key as Payment1 %v",err)
	}
	fmt.Println(paymentFindKey)

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
		Message: "Payment details Retrived Successfully",
	},nil
}








// Account Number Validations
	// fromAccount:=string(req.FromAccountId)
	// if len(fromAccount)!=10{
	// 	return &pb.CreateResponse{
	// 	 Message: "FromAccount is not of valid Digit,Please Enter the correct Account Number",
	// },nil
	// }

	// ToAccount:=string(req.FromAccountId)
	// if len(ToAccount)!=10{
	// 	return &pb.CreateResponse{
	// 	 Message: "FromAccount is not of valid Digit,Please Enter the correct Account Number",
	// },nil
	// }
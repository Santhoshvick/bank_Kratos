package service

import (
	"context"
	"fmt"
	"log"

	accountpb "account-service/api/helloworld/v1"
	pb "card-service/api/helloworld/v1"
	"card-service/internal/biz"
	"card-service/internal/handler"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GreeterService is a greeter service.
type CardService struct {
	pb.UnimplementedGreeterServer
	accountClient  accountpb.AccountClient
	uc *biz.CardUsecase
}

func NewCardService(uc *biz.CardUsecase,accClient accountpb.AccountClient) *CardService {
	return &CardService{uc: uc,
		accountClient: accClient,
	}
}



func(s  *CardService) CreateCard(ctx context.Context,req *pb.CreateCardRequest) (*pb.CreateCardResponse,error){

	conn1,err:=grpc.NewClient("localhost:9002",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.Fatalf("GRPC connection is not establishded %v",err)
		return nil,err
	}
	fmt.Println(req.AccountNumber)

   //Account Number Validation
	if _, err := handler.ValidateAccountNumber(req.AccountNumber); err != nil {
    return &pb.CreateCardResponse{
        Message: "Validation failed for Account Fields: " + err.Error(),
    }, nil
}

    //Card Number Random Generator
      cardnum:=handler.RandomGenerator()

   // Card Type Validation
   if err:=handler.CardType(req.CardType);err!=nil{

	return &pb.CreateCardResponse{
        Message: "" + err.Error(),
    }, nil
   }

   // Card Status Validation

   if err:=handler.CardStatus(req.CardStatus);err!=nil{
	  return &pb.CreateCardResponse{
        Message: "" + err.Error(),
    }, nil
   }


	client:=accountpb.NewAccountClient(conn1)
	val3,err:=client.FindAccountNumber(ctx,&accountpb.FindRequest{AccountNumber: req.AccountNumber})
	if err!=nil{
		log.Fatalf("%v",err)
		return &pb.CreateCardResponse{
		Message:"Card Creation failed Account for the customer does not exist,Please contact the branch to create an account before applying for the card",
	},nil
	}

	fmt.Println(val3)
	 if val3.Message=="Account Details Not Found"{
		return &pb.CreateCardResponse{
		Message:"Card Creation failed Account for the customer does not exist,Please contact the branch to create an account before applying for the card",
	},nil
	 }
	 g,err:=s.uc.CreateCard(ctx,&biz.Card{
		AccountNumber:req.AccountNumber,
		CardNumber: cardnum,
		CardType:req.CardType,
		CardStatus:req.CardStatus,
		ExpiryDate:req.ExpiryDate,
		DailyLimit:req.DailyLimit,
		MonthlyLimit:req.MonthlyLimit,
		PinAttempt:req.PinAttempt,
	})
	if err!=nil{
		 return &pb.CreateCardResponse{
		Message:"Card Creation failed Account Does not exist",
	},nil
	}
	return &pb.CreateCardResponse{
		CardId: g.CardId,
		CardNumber:g.CardNumber,
		AccountNumber:req.AccountNumber,
		CardType:g.CardType,
		CardStatus:g.CardStatus,
		ExpiryDate:g.ExpiryDate,
		DailyLimit:g.DailyLimit,
		MonthlyLimit:g.MonthlyLimit,
		PinAttempt:g.PinAttempt,
		Message:"Card Created Successfully",
	},nil
}


func(s  *CardService) UpdateCard(ctx context.Context,req *pb.UpdateCardRequest) (*pb.UpdateCardResponse,error){
	//Account Number Validation
	if _, err := handler.ValidateAccountNumber(req.AccountNumber); err != nil {
    return &pb.UpdateCardResponse{
        Message: "Validation failed for Account Fields: " + err.Error(),
    }, nil
}

   // Card Type Validation
   if err:=handler.CardType(req.CardType);err!=nil{
	return &pb.UpdateCardResponse{
        Message: "" + err.Error(),
    }, nil
   }

   // Card Status Validation
   if err:=handler.CardStatus(req.CardStatus);err!=nil{
	  return &pb.UpdateCardResponse{
        Message: "" + err.Error(),
    }, nil
   }

	
	g,err:=s.uc.CreateCard(ctx,&biz.Card{
		CardNumber:req.CardNumber,
		CardId:req.CardId,
		AccountNumber:req.AccountNumber,
		CardType:req.CardType,
		CardStatus:req.CardStatus,
		ExpiryDate:req.ExpiryDate,
		DailyLimit:req.DailyLimit,
		MonthlyLimit:req.MonthlyLimit,
		PinAttempt:req.PinAttempt,
	})
	if err!=nil{
		return &pb.UpdateCardResponse{
			Message: "Card Does Not exist",
      
    }, nil
	}
	return &pb.UpdateCardResponse{
		CardNumber:g.CardNumber,
		AccountNumber:g.AccountNumber,
		CardType:g.CardType,
		CardStatus:g.CardStatus,
		ExpiryDate:g.ExpiryDate,
		DailyLimit:g.DailyLimit,
		MonthlyLimit:g.MonthlyLimit,
		PinAttempt:g.PinAttempt,
		Message: "Card Details is updated successfully",
	},nil
}
func (s *CardService) DeleteCard(ctx context.Context, req *pb.DeleteCardRequest) (*pb.DeleteCardResponse, error) {
    g, err := s.uc.DeleteCard(ctx, req.CardId)
    if err != nil {
		return &pb.DeleteCardResponse{
        Message: "Card Not Found",
    }, nil
}
    return &pb.DeleteCardResponse{
        CardNumber:    g.CardNumber,
        AccountNumber: g.AccountNumber,
        CardType:      g.CardType,
        CardStatus:    g.CardStatus,
        ExpiryDate:    g.ExpiryDate,
        DailyLimit:    g.DailyLimit,
        MonthlyLimit:  g.MonthlyLimit,
        PinAttempt:    g.PinAttempt,
		Message: "Card is Deleted Successfully",
    }, nil
}
func (s *CardService) FindCard(ctx context.Context, req *pb.FindCardRequest) (*pb.FindCardResponse, error) {
    g, err := s.uc.FindCard(ctx, req.CardId)
    if err != nil {
         return &pb.FindCardResponse{
			Message:"Card Details Does not Exist",
    
    }, nil
    }
    return &pb.FindCardResponse{
        CardNumber:    g.CardNumber,
        AccountNumber: g.AccountNumber,
        CardType:      g.CardType,
        CardStatus:    g.CardStatus,
        ExpiryDate:    g.ExpiryDate,
        DailyLimit:    g.DailyLimit,
        MonthlyLimit:  g.MonthlyLimit,
        PinAttempt:    g.PinAttempt,
		Message: "Card is Deleted Successfully",
    }, nil
}


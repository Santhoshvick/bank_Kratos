package service

import (
	"context"

	pb "card-service/api/helloworld/v1"
	"card-service/internal/biz"
)

// GreeterService is a greeter service.
type CardService struct {
	pb.UnimplementedGreeterServer

	uc *biz.CardUsecase
}

func NewCardService(uc *biz.CardUsecase) *CardService {
	return &CardService{uc: uc}
}




func(s  *CardService) CreateCard(ctx context.Context,req *pb.CreateRequest) (*pb.CreateResponse,error){
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
		return nil,err
	}

	return &pb.CreateResponse{
		CardNumber:g.CardNumber,
		AccountNumber:g.AccountNumber,
		CardType:g.CardType,
		CardStatus:g.CardStatus,
		ExpiryDate:g.ExpiryDate,
		DailyLimit:g.DailyLimit,
		MonthlyLimit:g.MonthlyLimit,
		PinAttempt:g.PinAttempt,

	},nil
}



func(s  *CardService) UpdateCard(ctx context.Context,req *pb.UpdateRequest) (*pb.UpdateResponse,error){
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
		return nil,err
	}

	return &pb.UpdateResponse{
		CardNumber:g.CardNumber,
		AccountNumber:g.AccountNumber,
		CardType:g.CardType,
		CardStatus:g.CardStatus,
		ExpiryDate:g.ExpiryDate,
		DailyLimit:g.DailyLimit,
		MonthlyLimit:g.MonthlyLimit,
		PinAttempt:g.PinAttempt,

	},nil
}


func (s *CardService) DeleteCard(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
    g, err := s.uc.DeleteCard(ctx, req.CardId)
    if err != nil {
        return nil, err
    }

    return &pb.DeleteResponse{
        CardNumber:    g.CardNumber,
        AccountNumber: g.AccountNumber,
        CardType:      g.CardType,
        CardStatus:    g.CardStatus,
        ExpiryDate:    g.ExpiryDate,
        DailyLimit:    g.DailyLimit,
        MonthlyLimit:  g.MonthlyLimit,
        PinAttempt:    g.PinAttempt,
    }, nil
}



func (s *CardService) FindCard(ctx context.Context, req *pb.FindRequest) (*pb.FindResponse, error) {
    g, err := s.uc.FindCard(ctx, req.CardId)
    if err != nil {
        return nil, err
    }

    return &pb.FindResponse{
        CardNumber:    g.CardNumber,
        AccountNumber: g.AccountNumber,
        CardType:      g.CardType,
        CardStatus:    g.CardStatus,
        ExpiryDate:    g.ExpiryDate,
        DailyLimit:    g.DailyLimit,
        MonthlyLimit:  g.MonthlyLimit,
        PinAttempt:    g.PinAttempt,
    }, nil
}


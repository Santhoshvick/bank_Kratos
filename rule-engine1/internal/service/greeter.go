package service

import (
	"context"
	"strconv"

	v3 "customer-service/api/helloworld/v1"
	"log"
	v1 "rule-engine1/api/helloworld/v1"
	"rule-engine1/internal/biz"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GreeterService is a greeter service.
type RuleService struct {
	v1.UnimplementedGreeterServer

	uc *biz.RuleUsecase
}

// NewGreeterService new a greeter service.
func NewRuleService(uc *biz.RuleUsecase) *RuleService {
	return &RuleService{uc: uc}
}


func (s *RuleService) RuleEngine(ctx context.Context,in *v1.RuleRequest)(*v1.RuleResponse,error){
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Printf("Error connecting to rule engine service: %v", err)
        return nil, err
    }
    defer conn.Close()
 
    client := v3.NewCustomerClient(conn)
 
    _, err = client.CreateCustomer(ctx, &v3.CreateRequest{CustomerNumber: strconv.Itoa(int(in.CustomerId))})
    if err != nil {
        log.Printf("Error calling GetUser: %v", err)
        return nil, err
    }
	g,err:=s.uc.RuleEngine(ctx,&biz.Rule{
		CustomerId:in.CustomerId,
		Description:in.Description,
		AvailableBalance:in.AvailableBalance,
	})
	if err != nil {
		return nil, err
	}
	return &v1.RuleResponse{
		CustomerId:g.CustomerId ,
		Description: g.Description,
		AvailableBalance: g.AvailableBalance,
	}, nil
}


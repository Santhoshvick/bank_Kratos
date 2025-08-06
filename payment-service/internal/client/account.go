package client

import (

	accountpb "account-service/api/helloworld/v1"
	transactionpb "transaction-service/api/helloworld/v1"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
)

var endpoint string = "0.0.0.0:9002"

func NewAccountClient(endpoint string) (accountpb.AccountClient, error) {
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(endpoint))
	if err != nil {
		return nil, err
	}
	return accountpb.NewAccountClient(conn), nil
	
}


var endpoint1 string = "0.0.0.0:9005"

func NewTransactionClient(endpoint1 string) (transactionpb.TransactionClient, error) {
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(endpoint1))
	if err != nil {
		return nil, err
	}
	return transactionpb.NewTransactionClient(conn), nil
	
}

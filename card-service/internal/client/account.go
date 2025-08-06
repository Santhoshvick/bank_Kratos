package client

import (
	accountpb "account-service/api/helloworld/v1"
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

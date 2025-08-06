package client

import (

	paymentpb "payment-service/api/helloworld/v1"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
)

var endpoint string = "0.0.0.0:9004"

func NewPaymentClient(endpoint string) (paymentpb.PaymentClient, error) {
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(endpoint))
	if err != nil {
		return nil, err
	}
	return paymentpb.NewPaymentClient(conn), nil	
}

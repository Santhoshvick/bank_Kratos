package mapper

import (
	pb "account-block/api/helloworld/v1"
	"account-block/internal/data"
)

type Mapper struct{
	mapper *data.GreeterRepo
}

func Mapper1(req *pb.BlockRequest) *pb.BlockRequest {
	return &pb.BlockRequest{
		AccountId:   req.AccountId,
		AccountNumber: req.AccountNumber,
	}

}
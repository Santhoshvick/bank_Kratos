package service
import (
	"context"
	accountpb "account-service/api/helloworld/v1"

	pb "account-block/api/helloworld/v1"
	"account-block/internal/handler"
)




// GreeterService is a greeter service.
type AccountBlockService struct {
	pb.UnimplementedGreeterServer
	acchandler *handler.AccountHandler
	accountClient accountpb.AccountClient

}

// NewGreeterService new a greeter service.
func NewAccountBlockService(acchandler *handler.AccountHandler,accClient accountpb.AccountClient) *AccountBlockService {
	return &AccountBlockService{
		acchandler: acchandler,
		accountClient: accClient,
	}
}

func (s *AccountBlockService) CreateAccountBlock(ctx context.Context, req *pb.BlockRequest)(*pb.BlockResponse,error){
	return s.acchandler.CreateAccountBlock(ctx,req)
	
}

func (s *AccountBlockService) AdminAccountBlock(ctx context.Context, req *pb.BlockRequest)(*pb.BlockResponse,error){
	return s.acchandler.CreateAccountBlock(ctx,req)
	
}

func (s *AccountBlockService) RiskAccountBlock(ctx context.Context, req *pb.BlockRequest)(*pb.BlockResponse,error){
	return s.acchandler.CreateAccountBlock(ctx,req)
	
}

func (s *AccountBlockService) CaseAccountBlock(ctx context.Context, req *pb.BlockRequest)(*pb.BlockResponse,error){
	return s.acchandler.CreateAccountBlock(ctx,req)
	
}

func (s *AccountBlockService) FindAccountBlock(ctx context.Context, req *pb.FindBlockRequest)(*pb.FindBlockResponse,error){
	return s.acchandler.FindAccountBlock(ctx,req)
	
}

func (s *AccountBlockService) UpdateAccountBlock(ctx context.Context, req *pb.UpdateBlockRequest)(*pb.UpdateBlockResponse,error){
	return s.acchandler.UpdateAccountBlock(ctx,req)
	
}
 
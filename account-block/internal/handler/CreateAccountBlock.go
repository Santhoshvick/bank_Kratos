package handler

import (
	pb "account-block/api/helloworld/v1"
	accountpb "account-service/api/helloworld/v1"
	"context"
	"fmt"


	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func (h *AccountHandler)CreateAccountBlock(ctx context.Context, req *pb.BlockRequest)(*pb.BlockResponse,error){

	
	
	conn, err := grpc.NewClient("localhost:9002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error Connecting: %v", err)
		return nil, err
	}
	client :=accountpb.NewAccountClient(conn)
	val,err:=client.FindAccountNumber(ctx,&accountpb.FindRequest{AccountNumber:req.AccountNumber})
	if err!=nil{
		return nil,err
	}
	err1:=h.repo.CreateAccountBlock(ctx,Mapper(ctx,req))
	if err1!=nil{
		return nil,err1
	}
	fmt.Println(val)
	if val.Message=="Account Details Not Found"{
		return &pb.BlockResponse{
			Message:"The Account does not exist",
		},nil
	}
	findFunc := func(accountNumber int64) (*pb.FindBlockResponse, error) {
		return h.repo.FindAccountBlock(ctx, &pb.FindBlockRequest{AccountNumber: req.AccountNumber})
	}
        
	blockInfo, err := findFunc(req.AccountNumber)
	if err != nil {
		return &pb.BlockResponse{
			AccountNumber: req.AccountNumber,
			AccountId: req.AccountId,
			Accblock: req.Accblock,
			Message:"The Account is Blocked Successfully",
		 },nil
	}
	

	val1:=blockInfo.Accblock
	if val1==pb.AccountBlock_UNKNOWN_BLOCK||val1==pb.AccountBlock_adminBlock||val1==pb.AccountBlock_caseBlock{
		// return &pb.BlockResponse{
		// 	Message:"The Account is Already Block",
		// },nil
	
		updateFunc := func(accountNumber int64) (*pb.UpdateBlockResponse, error) {
		return h.repo.UpdateAccountBlock(ctx, &pb.UpdateBlockRequest{AccountNumber: req.AccountNumber})
	}
        
	     blockInfo, err := updateFunc(req.AccountNumber)
		 fmt.Println(blockInfo)
		 if err!=nil{
			return nil,err
		 }
		 if blockInfo.Accblock!=pb.AccountBlock_UNKNOWN_BLOCK||val1!=pb.AccountBlock_adminBlock||val1!=pb.AccountBlock_caseBlock{
	     return &pb.BlockResponse{
			AccountNumber: req.AccountNumber,
			AccountId: req.AccountId,
			Accblock: pb.AccountBlock_adminBlock,
			Message:"The Account is Blocked Successfully",

		 },nil
		 }
	}else{
		return &pb.BlockResponse{
			AccountNumber: req.AccountNumber,
			AccountId: req.AccountId,
			Accblock: pb.AccountBlock_release,
			Message:"The Account is Released Successfully",
		 },nil


	}
	if val.Status=="Active"{
		return &pb.BlockResponse{
			AccountId: req.AccountId,
			AccountNumber: req.AccountNumber,
			Message:"Account is Already in Active",
		},nil
	}
	if val.Status=="Inactive"{
	      return &pb.BlockResponse{
			AccountId: req.AccountId,
			AccountNumber: req.AccountNumber,
			Accblock:pb.AccountBlock_adminBlock,
			Message:"Account is InActive",
		},nil
	}

	if val.Status=="Pending"{
		return &pb.BlockResponse{
			AccountId: req.AccountId,
			AccountNumber: req.AccountNumber,
			Accblock: pb.AccountBlock_riskBlock,
			Message:"Account is not Creted there is issue with the personal Details",
		},nil
	}
	return &pb.BlockResponse{
		AccountId: val.AccountId,
		AccountNumber: val.AccountNumber,
		Accblock:pb.AccountBlock_adminBlock,
	},nil	
}

//admin account block
func (h *AccountHandler) AdminAccountBlock(ctx context.Context,req *pb.BlockRequest)(*pb.BlockResponse,error){
	conn, err := grpc.NewClient("localhost:9002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error Connecting: %v", err)
		return nil, err
	}
	client :=accountpb.NewAccountClient(conn)
	val,err:=client.FindAccountNumber(ctx,&accountpb.FindRequest{AccountNumber:req.AccountNumber})
	if err!=nil{
		return nil,err
	}
	err1:=h.repo.AdminAccountBlock(ctx,Mapper(ctx,req))
	if err1!=nil{
		return nil,err1
	}
	fmt.Println(val)
	findFunc := func(accountNumber int64) (*pb.FindBlockResponse, error) {
		return h.repo.FindAccountBlock(ctx, &pb.FindBlockRequest{AccountNumber: req.AccountNumber})
	}
        
	blockInfo, err := findFunc(req.AccountNumber)
	if err != nil {
		return &pb.BlockResponse{
			AccountNumber: req.AccountNumber,
			AccountId: req.AccountId,
			Accblock: req.Accblock,
			Message:"The Account is Admin  Blocked Successfully",
		 },nil
	}
	if blockInfo.Accblock==pb.AccountBlock_adminBlock{
		return &pb.BlockResponse{
			AccountNumber: req.AccountNumber,
			AccountId: req.AccountId,
			Accblock: req.Accblock,
			Message:"The Account is Admin  Blocked Successfully",
		 },nil
	}

	return &pb.BlockResponse{
			AccountNumber: req.AccountNumber,
			AccountId: req.AccountId,
			Accblock: req.Accblock,
			Message:"The Account is Admin  Blocked Successfully",
		 },nil
	
	

}

func(h *AccountHandler) FindAccountBlock(ctx context.Context,req *pb.FindBlockRequest)(*pb.FindBlockResponse,error){
	g,err:=h.repo.FindAccountBlock(ctx,&pb.FindBlockRequest{
		AccountNumber: req.AccountNumber,

	})
	if err!=nil{
		return nil,fmt.Errorf(err.Error())
	}
	fmt.Println(g)

	return &pb.FindBlockResponse{
		AccountNumber: g.AccountNumber,
		AccountId: g.AccountNumber,
		Message:"Account Details Retrived Successfully",
		Accblock: g.Accblock,
	},nil
}
func(h *AccountHandler) UpdateAccountBlock(ctx context.Context,req *pb.UpdateBlockRequest)(*pb.UpdateBlockResponse,error){
	g,err:=h.repo.UpdateAccountBlock(ctx,&pb.UpdateBlockRequest{
		AccountNumber: req.AccountNumber,
	})
	if err!=nil{
		return nil,fmt.Errorf(err.Error())
	}
	fmt.Println(g)

	return &pb.UpdateBlockResponse{
		AccountNumber: g.AccountNumber,
		AccountId: g.AccountNumber,
		Message:"Account Details Retrived Successfully",
		Accblock: g.Accblock,
	},nil
}


func Mapper(ctx context.Context,req *pb.BlockRequest) *pb.BlockRequest {
	conn, err := grpc.NewClient("localhost:9002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error Connecting: %v", err)
		return nil
	}
	client :=accountpb.NewAccountClient(conn)
	val,err:=client.FindAccountNumber(ctx,&accountpb.FindRequest{AccountNumber:req.AccountNumber})
	if err!=nil{
		return nil
	}
	if val.Status=="Pending"{
		return &pb.BlockRequest{
		AccountId:   req.AccountId,
		AccountNumber: req.AccountNumber,
		Accblock: pb.AccountBlock_riskBlock,
	}
} else if val.Status=="Inactive"{
		return &pb.BlockRequest{
		AccountId:   req.AccountId,
		AccountNumber: req.AccountNumber,
		Accblock: pb.AccountBlock_adminBlock,
	}
}
	return &pb.BlockRequest{
		AccountId: req.AccountId,
		AccountNumber: req.AccountNumber,
		Accblock: pb.AccountBlock_UNKNOWN_BLOCK,
	}
}





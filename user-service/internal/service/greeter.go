package service

import (
	"context"

	pb "user-service/api/helloworld/v1"
	"user-service/internal/biz"
)

// GreeterService is a greeter service.
type UserService struct {
	pb.UnimplementedUserServer
	uc *biz.UserUsecase
}

// NewGreeterService new a greeter service.
func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	g, err := s.uc.CreateUser(ctx, &biz.User{
		FirstName:req.FirstName,
		LastName:req.LastName,
		DateOfBirth:req.DateofBirth,
		Nationality:req.Nationality,
		Email:req.Email,
		Phone:req.Phone,
		Address1:req.Address1,
		Address2:req.Address2,
		City:req.City,
		Country:req.Country,
		UserName:req.UserName,
		Password:req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateResponse{
		FirstName: g.FirstName,
		LastName: g.LastName,
		DateofBirth: g.DateOfBirth,
		Nationality: g.Nationality,
		Email:g.Email,
		Phone:g.Phone,		
		Address1: g.Address1,
		Address2: g.Address2,
		City:g.City,
		Country: g.Country,
		UserName: g.UserName,
		Password: g.Password,
		UserId: g.UserId,
		Message: "User Created Successfully",
	}, nil
}



func (s *UserService) FindUser(ctx context.Context, req *pb.FindRequest) (*pb.FindResponse, error) {
	g, err := s.uc.FindUser(ctx, req.Email)
	if err != nil {
		return &pb.FindResponse{
			Message:"User Does Not Exist",
		
	}, nil
	}
	return &pb.FindResponse{
		FirstName: g.FirstName,
		LastName: g.LastName,
		DateofBirth: g.DateOfBirth,
		Nationality: g.Nationality,
		Email:g.Email,
		Phone:g.Phone,		
		Address1: g.Address1,
		Address2: g.Address2,
		City:g.City,
		Country: g.Country,
		UserName: g.UserName,
		Password: g.Password,
	}, nil
}

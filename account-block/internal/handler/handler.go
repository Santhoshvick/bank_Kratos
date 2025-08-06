package handler

import "account-block/internal/data"

// import "account-block/internal/data"

type AccountHandler struct{
	repo *data.GreeterRepo
}


func NewAccountHandler(
	repo *data.GreeterRepo,
)*AccountHandler{
	return &AccountHandler{repo:repo}
}
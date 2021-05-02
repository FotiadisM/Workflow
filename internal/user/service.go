package user

import (
	"context"
	"fmt"
)

type Service struct {
}

func NewService() Service {
	return Service{}
}

func (s Service) NewUser(ctx context.Context, req NewUserRequest) (res NewUserResponse, err error) {

	return
}

func (s Service) GetUser(ctx context.Context, req GetUserRequest) (res GetUserResponse, err error) {

	fmt.Println(req.ID)

	res.User.FName = "mike"
	res.User.LName = "fotiadis"

	return
}

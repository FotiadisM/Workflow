package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	NewUserEndpoint endpoint.Endpoint
	GetUserEndpoint endpoint.Endpoint
}

func NewEndpoints(s Service) Endpoints {
	return Endpoints{
		NewUserEndpoint: makeNewUserEndpoint(s),
		GetUserEndpoint: makeGetUserEndpoint(s),
	}
}

func makeNewUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(NewUserRequest)
		res, err := s.NewUser(ctx, req)

		return res, err
	}
}

func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetUserRequest)
		res, err := s.GetUser(ctx, req)

		return res, err
	}
}

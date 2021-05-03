package auth

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	signInEndpoint endpoint.Endpoint
	signUpEndpoint endpoint.Endpoint
}

func NewEndpoints(s Service) Endpoints {
	return Endpoints{
		makeSignInEndpoint(s),
		makeSignUpEndpoint(s),
	}
}

func makeSignInEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(signInRequest)
		res, err := s.signIn(ctx, req)

		return res, err
	}
}

func makeSignUpEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(signUpRequest)
		res, err := s.signUp(ctx, req)

		return res, err
	}
}

package auth

import (
	"context"
	"fmt"

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

func (e *Endpoints) WrapAllExcept(middleware endpoint.Middleware, excluded ...string) {
	included := map[string]struct{}{
		"signIn": {},
		"signUp": {},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist", ex))
		}
		delete(included, ex)
	}

	for inc := range included {
		switch inc {
		case "signIn":
			e.signInEndpoint = middleware(e.signInEndpoint)
		case "signUp":
			e.signUpEndpoint = middleware(e.signUpEndpoint)
		default:
			panic(fmt.Sprintf("Endpoint '%s' is missing", inc))
		}
	}
}

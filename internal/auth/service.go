package auth

import "context"

type Service interface {
	signIn(ctx context.Context, req signInRequest) (res signInResponse, err error)
	signUp(ctx context.Context, req signUpRequest) (res signUpResponse, err error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return service{r}
}

func (s service) signIn(ctx context.Context, req signInRequest) (res signInResponse, err error) {
	return
}

func (s service) signUp(ctx context.Context, req signUpRequest) (res signUpResponse, err error) {
	return
}

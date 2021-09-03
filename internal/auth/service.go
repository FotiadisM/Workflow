package auth

import (
	"context"
	"strings"

	"github.com/FotiadisM/workflow-server/internal/user"
	"golang.org/x/crypto/bcrypt"
)

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
	hashedPassword, err := s.repo.GetUserPassword(ctx, req.Email)
	if err != nil {
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)); err != nil {
		return
	}

	u, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return
	}

	res.User.ID = u.ID
	res.User.FName = u.FName
	res.User.LName = u.LName
	res.User.Email = u.Email
	res.User.Company = u.Company
	res.User.Position = u.Position
	res.User.Role = u.Role

	return
}

func (s service) signUp(ctx context.Context, req signUpRequest) (res signUpResponse, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	fileID, err := s.repo.PostFile(ctx, req.ProfilePic)
	if err != nil {
		return
	}

	fName := strings.Title(req.FName)
	lName := strings.Title(req.LName)
	company := strings.Title(req.Company)
	posistion := strings.Title(req.Position)
	id, err := s.repo.CreateUser(ctx, fName, lName, req.Email, company, posistion, fileID, string(hashedPassword))
	if err != nil {
		return
	}

	res.User.ID = id
	res.User.FName = fName
	res.User.LName = lName
	res.User.Email = req.Email
	res.User.Company = company
	res.User.Position = posistion
	res.User.ProfilePic = fileID
	res.User.Role = user.Normal

	return
}

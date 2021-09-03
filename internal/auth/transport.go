package auth

import (
	"io"

	"github.com/FotiadisM/workflow-server/internal/user"
)

type signInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signInResponse struct {
	AccessToken  string    `json:"access_token"`
	RefressToken string    `json:"refress_token"`
	User         user.User `json:"user"`
}

type signUpRequest struct {
	FName      string `json:"f_name"`
	LName      string `json:"l_name"`
	Email      string `json:"email"`
	Company    string `json:"company"`
	Position   string `json:"position"`
	Password   string `json:"password"`
	ProfilePic io.Reader
}

type signUpResponse struct {
	AccessToken  string    `json:"access_token"`
	RefressToken string    `json:"refress_token"`
	User         user.User `json:"user"`
}

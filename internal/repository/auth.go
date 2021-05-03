package repository

import (
	"context"

	"github.com/FotiadisM/workflow-server/internal/user"
)

func (r Repository) GetUserPassword(ctx context.Context, email string) (password string, err error) {
	return
}

func (r Repository) CreateUserCredentials(ctx context.Context, email string, password string) (err error) {
	return
}

func (r Repository) CreateUser(ctx context.Context, fName string, lName string, email string) (id string, err error) {
	return
}

func (r Repository) GetUser(ctx context.Context, email string) (u user.User, err error) {
	return
}

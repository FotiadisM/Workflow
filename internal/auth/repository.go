package auth

import (
	"context"

	"github.com/FotiadisM/workflow-server/internal/user"
)

type Repository interface {
	GetUserPassword(ctx context.Context, email string) (password string, err error)
	CreateUser(ctx context.Context, fName string, lName string, email string, password string) (id string, err error)
	GetUser(ctx context.Context, email string) (u user.User, err error)
}

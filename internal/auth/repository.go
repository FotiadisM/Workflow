package auth

import (
	"context"

	"github.com/FotiadisM/workflow-server/internal/user"
)

type Repository interface {
	GetUserPassword(ctx context.Context, email string) (password string, err error)
	CreateUser(ctx context.Context, fName, lName, email, password string) (id string, err error)
	GetUserByEmail(ctx context.Context, email string) (u user.User, err error)
}

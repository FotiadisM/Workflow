package auth

import (
	"context"
	"io"

	"github.com/FotiadisM/workflow-server/internal/user"
)

type Repository interface {
	GetUserPassword(ctx context.Context, email string) (password string, err error)
	CreateUser(ctx context.Context, fName, lName, email, company, position, profilePic, password string) (id string, err error)
	GetUserByEmail(ctx context.Context, email string) (u user.User, err error)
	PostFile(ctx context.Context, file io.Reader) (id string, err error)
}

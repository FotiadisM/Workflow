package user

import (
	"context"
)

type Role string

const (
	Admin  Role = "admin"
	Normal Role = "normal"
)

type User struct {
	ID       string `json:"id"`
	FName    string `json:"f_name"`
	LName    string `json:"l_name"`
	Email    string `json:"email"`
	Company  string `json:"company"`
	Position string `json:"position"`
	Role     Role   `json:"role"`
}

type Connection struct {
	ConnID string
	UserID string
}

type Repository interface {
	GetUserByID(ctx context.Context, id string) (u User, err error)
	GetPerpetator(ctx context.Context, id string) (u User, err error)
	GetConnections(ctx context.Context, userID string) (cons []Connection, err error)
	CreateConnectionRequest(ctx context.Context, userID, user2ID string) (ConnID string, err error)
	AcceptConnectionRequest(ctx context.Context, ConnID string) (err error)
	RejectConnectionRequest(ctx context.Context, ConnID string) (err error)
}

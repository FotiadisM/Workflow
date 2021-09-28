package feed

import (
	"context"
	"fmt"

	"github.com/FotiadisM/workflow-server/internal/user"
)

type ChannelFeed struct {
	PostID        string
	PerpetratorID string
	// post || comment || like
	Type string
}

type Repository interface {
	GetConnections(ctx context.Context, userID string) (cons []user.Connection, err error)
	UpdateUserFeed(ctx context.Context, userID, perpID, postID, feedType string) (id string, err error)
}

type Service interface {
	Run(ctx context.Context, ch chan ChannelFeed) (err error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return service{r}
}

func (s service) Run(ctx context.Context, ch chan ChannelFeed) (err error) {
	f := <-ch

	cons, err := s.r.GetConnections(ctx, f.PerpetratorID)
	if err != nil {
		return
	}

	for i, j := range cons {
		fmt.Println(i, j)
		_, err = s.r.UpdateUserFeed(ctx, j.UserID, f.PerpetratorID, f.PostID, f.Type)
		if err != nil {
			return
		}
	}

	fmt.Sprintln(f)
	return
}

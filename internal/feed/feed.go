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
	Run(ctx context.Context, ch chan ChannelFeed)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return service{r}
}

func (s service) Run(ctx context.Context, ch chan ChannelFeed) {
	for {
		f := <-ch

		_, err := s.r.UpdateUserFeed(ctx, f.PerpetratorID, f.PerpetratorID, f.PostID, f.Type)
		if err != nil {
			fmt.Println("Failed to update user feed", err)
			return
		}

		cons, err := s.r.GetConnections(ctx, f.PerpetratorID)
		if err != nil {
			fmt.Println("Failed to get connections", err)
			return
		}

		for i, j := range cons {
			fmt.Println(i, j)
			_, err = s.r.UpdateUserFeed(ctx, j.UserID, f.PerpetratorID, f.PostID, f.Type)
			if err != nil {
				fmt.Println("Failed to update user feed", err)
				return
			}
		}
	}
}

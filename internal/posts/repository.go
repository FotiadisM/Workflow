package posts

import (
	"context"
	"io"
	"time"
)

type PostVisibility string

const (
	All     PostVisibility = "all"
	Friends                = "friends"
	Nopne                  = "none"
)

type Post struct {
	ID         string         `json:"id"`
	UserID     string         `json:"user_id"`
	Text       string         `json:"text"`
	Images     []string       `json:"images"`
	Videos     []string       `json:"videos"`
	Visivility PostVisibility `json:"-"`
	Likes      []string       `json:"likes"`
	Comments   []string       `json:"comments"`
	Created    string         `json:"created"`
}

type Comment struct {
	ID      string   `json:"id"`
	PostID  string   `json:"post_id"`
	UserID  string   `json:"user_id"`
	Text    string   `json:"text"`
	Likes   []string `json:"likes"`
	Created string   `json:"created"`
}

type Feed struct {
	ID            string `json:"id"`
	UserID        string `json:"user_id"`
	PostID        string `json:"post_id"`
	PerpetratorID string `json:"perpetator_id"`
	FeedType      string `json:"type"`
}

type Repository interface {
	getFeed(ctx context.Context, userID string) (fs []Feed, err error)
	CreatePost(ctx context.Context, userID, text, visivility string, images, videos []string) (id string, created time.Time, err error)
	GetPost(ctx context.Context, postID string) (p *Post, err error)
	GetUserPosts(ctx context.Context, userID, FromUserID string) (posts []Post, err error)
	TogglePostLike(ctx context.Context, postID, userID string) (err error)
	CreatePostComment(ctx context.Context, postID, userID, text string) (id string, created time.Time, err error)
	GetPostComment(ctx context.Context, postID string) (c *Comment, err error)
	ToggleCommentLike(ctx context.Context, commentID, userID string) (err error)
	PostFile(ctx context.Context, file io.Reader) (id string, err error)
}

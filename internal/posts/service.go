package posts

import "context"

type Service interface {
	getPost(ctx context.Context, req getPostsRequest) (res getPostsResponse, err error)
	postPost(ctx context.Context, req postPostRequest) (res postPostResponse, err error)

	getComments(ctx context.Context, req getCommentsRequest) (res getCommentsResponse, err error)
	postComment(ctx context.Context, req postCommentRequest) (res postCommentResponse, err error)

	likePost(ctx context.Context, req likePostRequest) (res likePostRequest, err error)
	likeComment(ctx context.Context, req likeCommentRequest) (res likeCommentRequest, err error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return service{r}
}

func (s service) getPost(ctx context.Context, req getPostsRequest) (res getPostsResponse, err error) {
	return
}

func (s service) postPost(ctx context.Context, req postPostRequest) (res postPostResponse, err error) {
	return
}

func (s service) getComments(ctx context.Context, req getCommentsRequest) (res getCommentsResponse, err error) {
	return
}

func (s service) postComment(ctx context.Context, req postCommentRequest) (res postCommentResponse, err error) {
	return
}

func (s service) likePost(ctx context.Context, req likePostRequest) (res likePostRequest, err error) {
	return
}

func (s service) likeComment(ctx context.Context, req likeCommentRequest) (res likeCommentRequest, err error) {
	return
}

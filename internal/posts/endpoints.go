package posts

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	getPostsEndpoint    endpoint.Endpoint
	postPostsEndpoint   endpoint.Endpoint
	getCommentsEndpoint endpoint.Endpoint
	postCommentEndpoint endpoint.Endpoint
	likePostEndpoint    endpoint.Endpoint
	likeCommentEndpoint endpoint.Endpoint
}

func NewEndpoints(s Service) Endpoints {
	return Endpoints{
		makeGetPostEndpoint(s),
		makePostPostEndpoint(s),
		makeGetCommentsEndpoint(s),
		makePostCommentEndpoint(s),
		makeLikePostEndpoint(s),
		makeLikeCommentEndpoint(s),
	}
}

func makeGetPostEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getPostsRequest)
		res, err := s.getPost(ctx, req)

		return res, err
	}
}

func makePostPostEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postPostRequest)
		res, err := s.postPost(ctx, req)

		return res, err
	}
}

func makeGetCommentsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getCommentsRequest)
		res, err := s.getComments(ctx, req)

		return res, err
	}
}

func makePostCommentEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postCommentRequest)
		res, err := s.postComment(ctx, req)

		return res, err
	}
}

func makeLikePostEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(likePostRequest)
		res, err := s.likePost(ctx, req)

		return res, err
	}
}

func makeLikeCommentEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(likeCommentRequest)
		res, err := s.likeComment(ctx, req)

		return res, err
	}
}

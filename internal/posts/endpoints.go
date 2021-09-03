package posts

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	getPostsEndpoint          endpoint.Endpoint
	createPostsEndpoint       endpoint.Endpoint
	getPostCommentEndpoint    endpoint.Endpoint
	createPostCommentEndpoint endpoint.Endpoint
	togglePostLikeEndpoint    endpoint.Endpoint
	toggleCommentLikeEndpoint endpoint.Endpoint
}

func NewEndpoints(s Service) Endpoints {
	return Endpoints{
		makeGetPostEndpoint(s),
		makeCreatePostEndpoint(s),
		makeGetPostCommentEndpoint(s),
		makeCreatePostCommentEndpoint(s),
		makeTogglePostLikeEndpoint(s),
		makeToggleCommentLikeEndpoint(s),
	}
}

func makeGetPostEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getPostsRequest)
		res, err := s.getPost(ctx, req)

		return res, err
	}
}

func makeCreatePostEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createPostRequest)
		res, err := s.createPost(ctx, req)

		return res, err
	}
}

func makeGetPostCommentEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getPostCommentRequest)
		res, err := s.getPostComment(ctx, req)

		return res, err
	}
}

func makeCreatePostCommentEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createPostCommentRequest)
		res, err := s.createPostComment(ctx, req)

		return res, err
	}
}

func makeTogglePostLikeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(togglePostLikeRequest)
		res, err := s.togglePostLike(ctx, req)

		return res, err
	}
}

func makeToggleCommentLikeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(toggleCommentLikeRequest)
		res, err := s.toggleCommentLike(ctx, req)

		return res, err
	}
}

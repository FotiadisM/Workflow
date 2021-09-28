package posts

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	getFeedsEndpoint          endpoint.Endpoint
	getPostsEndpoint          endpoint.Endpoint
	getUserPostsEndpoint      endpoint.Endpoint
	createPostsEndpoint       endpoint.Endpoint
	getPostCommentEndpoint    endpoint.Endpoint
	createPostCommentEndpoint endpoint.Endpoint
	togglePostLikeEndpoint    endpoint.Endpoint
	toggleCommentLikeEndpoint endpoint.Endpoint
}

func NewEndpoints(s Service) Endpoints {
	return Endpoints{
		makeGetFeedEndpoint(s),
		makeGetPostEndpoint(s),
		makeGetUserPostsEndpoint(s),
		makeCreatePostEndpoint(s),
		makeGetPostCommentEndpoint(s),
		makeCreatePostCommentEndpoint(s),
		makeTogglePostLikeEndpoint(s),
		makeToggleCommentLikeEndpoint(s),
	}
}

func makeGetFeedEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getFeedsRequest)
		res, err := s.getFeed(ctx, req)

		return res, err
	}
}

func makeGetPostEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getPostsRequest)
		res, err := s.getPost(ctx, req)

		return res, err
	}
}

func makeGetUserPostsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getUserPostsRequest)
		res, err := s.getUserPosts(ctx, req)

		return res, err
	}
}

func makeCreatePostEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreatePostRequest)
		res, err := s.CreatePost(ctx, req)

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
		req := request.(TogglePostLikeRequest)
		res, err := s.TogglePostLike(ctx, req)

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

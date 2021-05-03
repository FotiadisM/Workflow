package posts

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPRouter(e Endpoints, r *mux.Router, options ...httptransport.ServerOption) {
	r.Methods("GET").Path("/").Handler(httptransport.NewServer(
		e.getPostsEndpoint,
		decodeGetPostRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("POST").Path("/").Handler(httptransport.NewServer(
		e.postPostsEndpoint,
		decodePostPostRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("GET").Path("/comments").Handler(httptransport.NewServer(
		e.getCommentsEndpoint,
		decodeGetCommentsRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("POST").Path("/comments").Handler(httptransport.NewServer(
		e.postCommentEndpoint,
		decodePostCommentRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("POST").Path("/like/").Handler(httptransport.NewServer(
		e.likePostEndpoint,
		decodeLikePostRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("POST").Path("/comments/like").Handler(httptransport.NewServer(
		e.likeCommentEndpoint,
		decodeLikeCommentRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))
}

func decodeGetPostRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	return
}

func decodePostPostRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	return
}

func decodeGetCommentsRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	return
}

func decodePostCommentRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	return
}

func decodeLikePostRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	return
}

func decodeLikeCommentRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	return
}

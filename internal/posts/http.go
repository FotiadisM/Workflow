package posts

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPRouter(e Endpoints, r *mux.Router, options ...httptransport.ServerOption) {
	r.Methods("POST").Path("/").Handler(httptransport.NewServer(
		e.createPostsEndpoint,
		decodeCreatePostRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("POST").Path("/like").Handler(httptransport.NewServer(
		e.togglePostLikeEndpoint,
		decodeTogglePostLikeRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("POST").Path("/comments").Handler(httptransport.NewServer(
		e.createPostCommentEndpoint,
		decodeCreatePostCommentRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("POST").Path("/comments/like").Handler(httptransport.NewServer(
		e.toggleCommentLikeEndpoint,
		decodeToggleCommentLikeRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("GET").Path("/comments/{id}").Handler(httptransport.NewServer(
		e.getPostCommentEndpoint,
		decodeGetPostCommentRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("GET").Path("/{id}").Handler(httptransport.NewServer(
		e.getPostsEndpoint,
		decodeGetPostRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))
}

func decodeGetPostRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req getPostsRequest
	vars := mux.Vars(r)
	req.PostID = vars["id"]
	return req, err
}

func decodeCreatePostRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req createPostRequest
	err = r.ParseMultipartForm(32 << 22)
	if err != nil {
		return
	}

	images := []io.ReadCloser{}
	fhs := r.MultipartForm.File["images"]
	for _, f := range fhs {
		f, err := f.Open()
		if err != nil {
			return req, err
		}
		images = append(images, f)
	}

	videos := []io.ReadCloser{}
	fhs = r.MultipartForm.File["videos"]
	for _, f := range fhs {
		f, err := f.Open()
		if err != nil {
			return req, err
		}
		videos = append(videos, f)
	}

	req.UserID = r.MultipartForm.Value["user_id"][0]
	req.Text = r.MultipartForm.Value["text"][0]
	req.Images = images
	req.Videos = videos
	req.Visibility = PostVisibility(r.MultipartForm.Value["visibility"][0])

	return req, nil
}

func decodeTogglePostLikeRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req togglePostLikeRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeGetPostCommentRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req getPostCommentRequest
	vars := mux.Vars(r)
	req.CommentID = vars["id"]
	return req, err
}

func decodeCreatePostCommentRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req createPostCommentRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeToggleCommentLikeRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req toggleCommentLikeRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

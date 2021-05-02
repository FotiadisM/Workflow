package posts

import (
	"context"
	"net/http"
)

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

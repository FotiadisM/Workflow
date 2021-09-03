package posts

import "io"

type getPostsRequest struct {
	PostID string `json:"post_id"`
}
type getPostsResponse struct {
	Post *Post `json:"post,omitempty"`
	Err  error `json:"err,omitempty"`
}

type createPostRequest struct {
	UserID     string
	Text       string
	Images     []io.ReadCloser
	Videos     []io.ReadCloser
	Visibility PostVisibility
}
type createPostResponse struct {
	Post *Post `json:"post,omitempty"`
	Err  error `json:"err,omitempty"`
}

type getPostCommentRequest struct {
	CommentID string `json:"comment_id"`
}
type getPostCommentResponse struct {
	Comment *Comment `json:"comment,omitempty"`
	Err     error    `json:"err,omitempty"`
}

type createPostCommentRequest struct {
	UserID string `json:"user_id"`
	PostID string `json:"post_id"`
	Text   string `json:"text"`
}
type createPostCommentResponse struct {
	Comment *Comment `json:"comment,omitempty"`
	Err     error    `json:"err,omitempty"`
}

type togglePostLikeRequest struct {
	UserID string `json:"user_id"`
	PostID string `json:"post_id"`
}
type togglePostLikeResponse struct {
	Err error `json:"err,omitempty"`
}

type toggleCommentLikeRequest struct {
	UserID    string `json:"user_id"`
	CommentID string `json:"comment_id"`
}
type toggleCommentLikeResponse struct {
	Err error `json:"err,omitempty"`
}

package posts

import "io"

type getPostsRequest struct {
	PostID string `json:"post_id"`
}
type getPostsResponse struct {
	Post *Post `json:"post,omitempty"`
	Err  error `json:"err,omitempty"`
}

type getUserPostsRequest struct {
	UserID     string `json:"user_id"`
	FromUserID string `json:"from_user_id"`
}
type getUserPostsResponse struct {
	Posts []Post `json:"posts,omitempty"`
	Err   error  `json:"err,omitempty"`
}

type CreatePostRequest struct {
	UserID     string
	Text       string
	Images     []io.ReadCloser
	Videos     []io.ReadCloser
	Visibility PostVisibility
}
type CreatePostResponse struct {
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

type TogglePostLikeRequest struct {
	UserID string `json:"user_id"`
	PostID string `json:"post_id"`
}
type TogglePostLikeResponse struct {
	Err error `json:"err,omitempty"`
}

type toggleCommentLikeRequest struct {
	UserID    string `json:"user_id"`
	CommentID string `json:"comment_id"`
}
type toggleCommentLikeResponse struct {
	Err error `json:"err,omitempty"`
}

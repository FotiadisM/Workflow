package user

type NewUserRequest struct{}
type NewUserResponse struct{}

type GetUserRequest struct {
	ID string `json:"id"`
}

type GetUserResponse struct {
	User  User  `json:"user"`
	Error error `json:"error"`
}

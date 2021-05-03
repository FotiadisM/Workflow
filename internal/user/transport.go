package user

type getUserRequest struct {
	ID string `json:"id"`
}

type getUserResponse struct {
	User  User  `json:"user"`
	Error error `json:"error"`
}

type getPerpetatorRequest struct{}
type getPerpetatorResponse struct{}

type getConnectionsRequest struct{}
type getConnectionsResponse struct{}

type postConnectionRequest struct{}
type postConnectionResponse struct{}

type changeConnectionRequest struct{}
type changeConnectionResponse struct{}

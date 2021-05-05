package user

type getUserRequest struct {
	UserID string `json:"user_id"`
}
type getUserResponse struct {
	User User  `json:"user"`
	Err  error `json:"err,omitempty"`
}

type getPerpetatorRequest struct {
	PerpID string `json:"perp_id"`
}
type getPerpetatorResponse struct {
	FName    string `json:"f_name"`
	LName    string `json:"l_name"`
	Company  string `json:"company"`
	Position string `json:"position"`
	Err      error  `json:"err,omitempty"`
}

type getConnectionsRequest struct {
	UserID string
}
type getConnectionsResponse struct {
	Connections []struct {
		ConnID string `json:"conn_id"`
		UserID string `json:"user_id"`
	} `json:"connections"`
	Err error `json:"err,omitempty"`
}

type postConnectionRequest struct {
	UserID  string `json:"user_id"`
	User2ID string `json:"user_2_id"`
}
type postConnectionResponse struct {
	ConnID string `json:"conn_id"`
	Err    error  `json:"err,omitempty"`
}

type changeConnectionRequest struct {
	UserID string `json:"user_id"`
	ConnID string `json:"conn_id"`
	Accept bool   `json:"accept"`
	Err    error  `json:"err,omitempty"`
}
type changeConnectionResponse struct {
	Err error `json:"err,omitempty"`
}

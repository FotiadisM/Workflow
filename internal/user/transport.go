package user

type getUserRequest struct {
	UserID string `json:"user_id"`
}
type getUserResponse struct {
	User *User `json:"user,omitempty"`
	Err  error `json:"err,omitempty"`
}

type getPerpetatorRequest struct {
	PerpID string `json:"perp_id"`
}
type getPerpetatorResponse struct {
	FName    string `json:"f_name,omitempty"`
	LName    string `json:"l_name,omitempty"`
	Company  string `json:"company,omitempty"`
	Position string `json:"position,omitempty"`
	Err      error  `json:"err,omitempty"`
}

type getConnectionsRequest struct {
	UserID string
}
type getConnectionsResponse struct {
	Connections []Connection `json:"connections,omitempty"`
	Err         error        `json:"err,omitempty"`
}

type postConnectionRequest struct {
	UserID  string `json:"user_id"`
	User2ID string `json:"user_2_id"`
}
type postConnectionResponse struct {
	ConnID string `json:"conn_id,omitempty"`
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

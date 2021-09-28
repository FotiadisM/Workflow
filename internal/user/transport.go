package user

type getUserRequest struct {
	UserID string `json:"user_id"`
}
type getUserResponse struct {
	User *User `json:"user,omitempty"`
	Err  error `json:"err,omitempty"`
}

type getUsersRequest struct {
}
type getUsersResponse struct {
	Users []User `json:"users,omitempty"`
	Err   error  `json:"err,omitempty"`
}

type getPerpetatorRequest struct {
	PerpID string `json:"perp_id"`
}
type getPerpetatorResponse struct {
	FName      string `json:"f_name,omitempty"`
	LName      string `json:"l_name,omitempty"`
	Company    string `json:"company,omitempty"`
	Position   string `json:"position,omitempty"`
	ProfilePic string `json:"profile_pic"`
	Err        error  `json:"err,omitempty"`
}

type getConnectionsRequest struct {
	UserID string `json:"user_id"`
}
type getConnectionsResponse struct {
	Connections []Connection `json:"connections,omitempty"`
	Err         error        `json:"err,omitempty"`
}

type postConnectionRequest struct {
	UserID     string `json:"user_id"`
	ReceiverID string `json:"receiver_id"`
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

type getConnectionRequestsRequst struct {
	UserID string `json:"user_id"`
}
type getConnectionRequestsResponse struct {
	Connections []Connection `json:"connections,omitempty"`
	Err         error        `json:"err,omitempty"`
}

type decideConnectionRequestRequst struct {
	ConnID string `json:"conn_id"`
	Accept bool   `json:"accept"`
}
type decideConnectionRequestResponse struct {
	ConnID string `json:"conn_id,omitempty"`
	Err    error  `json:"err,omitempty"`
}

package conversations

type getConversationsRequest struct {
	UserID string `json:"user_id"`
}
type getConversationsResponse struct {
	Conversations []Conversation `json:"conversations"`
	Err           error          `json:"err,omitempty"`
}

type postConversationsRequest struct {
	UserID     string `json:"user_id"`
	ConvUserID string `json:"conv_user_id"`
}
type postConversationsResponse struct {
	ConvID string `json:"conv_id,omitempty"`
	Err    error  `json:"err,omitempty"`
}

type getMessagesRequest struct {
	ConvID string `json:"conv_id"`
}
type getMessagesResponse struct {
	Messages []Message `json:"messages"`
	Err      error     `json:"err,omitempty"`
}

type postMessageRequest struct {
	ConvID   string `json:"conv_id"`
	SenterID string `json:"senter_id"`
	Text     string `json:"text"`
}
type postMessageResponse struct {
	MesgID   string `json:"mesg_id,omitempty"`
	TimeSent string `json:"time_sent,omitempty"`
	Err      error  `json:"err,omitempty"`
}

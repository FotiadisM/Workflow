package conversations

import "context"

type Conversation struct {
	ID         string `json:"id"`
	ConvUserID string `json:"conv_user_id"`
}

type Message struct {
	ID       string `json:"id"`
	ConvID   string `json:"conv_id"`
	SenterID string `json:"senter_id"`
	Text     string `json:"text"`
	Time     string `json:"time"`
}

type Repository interface {
	GetConversations(ctx context.Context, userID string) (convs []Conversation, err error)
	CreateConversation(ctx context.Context, userID, convUserID string) (convID string, err error)
	GetConversationMessages(ctx context.Context, convID string) (msgs []Message, err error)
	CreateConversationMessage(ctx context.Context, convID, senterID, text string) (msgID, timeSent string, err error)
}

package conversations

import (
	"context"
	"time"
)

type Conversation struct {
	ID         string `json:"id"`
	ConvUserID string `json:"conv_user_id"`
}

type Message struct {
	ID       string    `json:"id"`
	ConvID   string    `json:"-"`
	SenterID string    `json:"senter_id"`
	Text     string    `json:"text"`
	Time     time.Time `json:"time"`
}

type BroadCastMessage struct {
	ID       string    `json:"id"`
	ConvID   string    `json:"conn_id"`
	SenterID string    `json:"senter_id"`
	Receiver string    `json:"-"`
	Text     string    `json:"text"`
	Time     time.Time `json:"time"`
}

type Repository interface {
	GetReceiver(ctx context.Context, convID, userID string) (receiverID string, err error)
	GetConversations(ctx context.Context, userID string) (convs []Conversation, err error)
	CreateConversation(ctx context.Context, userID, convUserID string) (convID string, err error)
	GetConversationMessages(ctx context.Context, convID string) (msgs []Message, err error)
	CreateConversationMessage(ctx context.Context, convID, senterID, text string) (msgID string, timeSent time.Time, err error)
}

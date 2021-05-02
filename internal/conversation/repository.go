package conversation

import "context"

type Conversation struct {
	ID      string
	User1ID string
	User2ID string
}

type Message struct {
	ID       string
	ConvID   string
	SenterID string
	Text     string
	Time     string
}

type Repository interface {
	getConversations(ctx context.Context, userID string) (convs []Conversation, err error)
	postConversations(ctx context.Context, c Conversation) (err error)
	getMessages(ctx context.Context, convID string) (msgs []Message, err error)
	postMessage(ctx context.Context, msg Message) (m Message, err error)
}

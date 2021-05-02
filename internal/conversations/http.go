package conversations

import (
	"context"
	"net/http"
)

func decodeGetConversations(ctx context.Context, r *http.Request) (request interface{}, err error) {

	return
}

func decodePostConversations(ctx context.Context, req postConversationsRequest) (request interface{}, err error) {

	return
}

func decodeGetMessages(ctx context.Context, req getMessagesRequest) (request interface{}, err error) {
	return
}

func decodePostMessage(ctx context.Context, req postMessageRequest) (request interface{}, err error) {
	return
}

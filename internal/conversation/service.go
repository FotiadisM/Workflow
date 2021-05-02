package conversation

import "context"

type Service interface {
	getConversations(ctx context.Context, req getConversationsRequest) (res getConversationsResponse, err error)
	postConversations(ctx context.Context, req postConversationsRequest) (res postConversationsResponse, err error)
	getMessages(ctx context.Context, req getMessagesRequest) (res getMessagesResponse, err error)
	postMessage(ctx context.Context, req postMessageRequest) (res postMessageResponse, err error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return service{repo: r}
}

func (s service) getConversations(ctx context.Context, req getConversationsRequest) (res getConversationsResponse, err error) {
	return
}

func (s service) postConversations(ctx context.Context, req postConversationsRequest) (res postConversationsResponse, err error) {
	return
}

func (s service) getMessages(ctx context.Context, req getMessagesRequest) (res getMessagesResponse, err error) {
	return
}

func (s service) postMessage(ctx context.Context, req postMessageRequest) (res postMessageResponse, err error) {
	return
}

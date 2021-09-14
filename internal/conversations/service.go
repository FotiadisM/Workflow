package conversations

import (
	"context"
)

type Service interface {
	getConversations(ctx context.Context, req getConversationsRequest) (res getConversationsResponse, err error)
	postConversations(ctx context.Context, req postConversationsRequest) (res postConversationsResponse, err error)
	getMessages(ctx context.Context, req getMessagesRequest) (res getMessagesResponse, err error)
	postMessage(ctx context.Context, req postMessageRequest) (res postMessageResponse, err error)
}

type service struct {
	repo      Repository
	broadcast chan BroadCastMessage
}

func NewService(r Repository, b chan BroadCastMessage) Service {
	return service{repo: r, broadcast: b}
}

func (s service) getConversations(ctx context.Context, req getConversationsRequest) (res getConversationsResponse, err error) {
	convs, err := s.repo.GetConversations(ctx, req.UserID)
	if err != nil {
		return
	}

	if convs == nil {
		convs = []Conversation{}
	}
	res.Conversations = convs

	return
}

func (s service) postConversations(ctx context.Context, req postConversationsRequest) (res postConversationsResponse, err error) {
	id, err := s.repo.CreateConversation(ctx, req.UserID, req.ConvUserID)
	if err != nil {
		return
	}

	res.ConvID = id

	return
}

func (s service) getMessages(ctx context.Context, req getMessagesRequest) (res getMessagesResponse, err error) {
	msgs, err := s.repo.GetConversationMessages(ctx, req.ConvID)
	if err != nil {
		return
	}

	if msgs == nil {
		msgs = []Message{}
	}
	res.Messages = msgs

	return
}

func (s service) postMessage(ctx context.Context, req postMessageRequest) (res postMessageResponse, err error) {
	msgID, timeSent, err := s.repo.CreateConversationMessage(ctx, req.ConvID, req.SenterID, req.Text)
	if err != nil {
		return
	}

	receiverID, err := s.repo.GetReceiver(ctx, req.ConvID, req.SenterID)

	res.MesgID = msgID
	res.TimeSent = timeSent.String()
	s.broadcast <- BroadCastMessage{ID: msgID, ConvID: req.ConvID, SenterID: req.SenterID, Receiver: receiverID, Text: req.Text, Time: timeSent}

	return
}

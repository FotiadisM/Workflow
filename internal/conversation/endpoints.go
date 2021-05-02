package conversation

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	getConversationsEndpoint  endpoint.Endpoint
	postConversationsEndpoint endpoint.Endpoint
	getMessagesEndpoint       endpoint.Endpoint
	postMessageEndpoint       endpoint.Endpoint
}

func NewEndpoints(s Service) Endpoints {
	return Endpoints{
		makeGetConversationsEndpoint(s),
		makePostConversationsEndpoint(s),
		makeGetMessagesEndpoint(s),
		makePostMessageEndpoint(s),
	}
}

func makeGetConversationsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getConversationsRequest)
		res, err := s.getConversations(ctx, req)

		return res, err
	}
}

func makePostConversationsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postConversationsRequest)
		res, err := s.postConversations(ctx, req)

		return res, err
	}
}

func makeGetMessagesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getMessagesRequest)
		res, err := s.getMessages(ctx, req)

		return res, err
	}
}

func makePostMessageEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postMessageRequest)
		res, err := s.postMessage(ctx, req)

		return res, err
	}
}

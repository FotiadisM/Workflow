package conversations

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPRouter(e Endpoints, r *mux.Router, options ...httptransport.ServerOption) {
	r.Methods("GET").Path("/").Handler(httptransport.NewServer(
		e.getConversationsEndpoint,
		decodeGetConversations,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("POST").Path("/").Handler(httptransport.NewServer(
		e.postConversationsEndpoint,
		decodePostConversations,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("GET").Path("/messages").Handler(httptransport.NewServer(
		e.getMessagesEndpoint,
		decodeGetMessages,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("POST").Path("/messages").Handler(httptransport.NewServer(
		e.postMessageEndpoint,
		decodePostMessage,
		httptransport.EncodeJSONResponse,
		options...,
	))
}

func decodeGetConversations(ctx context.Context, r *http.Request) (request interface{}, err error) {

	return
}

func decodePostConversations(ctx context.Context, r *http.Request) (request interface{}, err error) {

	return
}

func decodeGetMessages(ctx context.Context, r *http.Request) (request interface{}, err error) {
	return
}

func decodePostMessage(ctx context.Context, r *http.Request) (request interface{}, err error) {
	return
}

package user

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPHandler(endpoints Endpoints, options ...httptransport.ServerOption) http.Handler {

	m := mux.NewRouter()

	m.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoints.GetUserEndpoint,
		decodeGetUserRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	m.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		endpoints.NewUserEndpoint,
		decodeNewUserRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	return m
}

func decodeNewUserRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)

	id := vars["id"]

	req := GetUserRequest{id}

	return req, nil
}

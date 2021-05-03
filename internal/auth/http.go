package auth

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPRouter(e Endpoints, r *mux.Router, options ...httptransport.ServerOption) {

	r.Methods("POST").Path("/signIn").Handler(httptransport.NewServer(
		e.signInEndpoint,
		decodeSignInRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("POST").Path("/signUp").Handler(httptransport.NewServer(
		e.signInEndpoint,
		decodeSignUpRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))
}

func decodeSignInRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var res signInRequest
	err = json.NewDecoder(r.Body).Decode(&res)

	return res, err
}

func decodeSignUpRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var res signUpRequest
	err = json.NewDecoder(r.Body).Decode(&res)

	return res, err
}

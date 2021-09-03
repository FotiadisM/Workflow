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
		e.signUpEndpoint,
		decodeSignUpRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))
}

func decodeSignInRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req signInRequest
	err = json.NewDecoder(r.Body).Decode(&req)

	return req, err
}

func decodeSignUpRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req signUpRequest
	err = r.ParseMultipartForm(32 << 20)
	if err != nil {
		return
	}

	f, _, err := r.FormFile("profile")
	if err != nil {
		return
	}

	req.FName = r.MultipartForm.Value["f_name"][0]
	req.LName = r.MultipartForm.Value["l_name"][0]
	req.Email = r.MultipartForm.Value["email"][0]
	req.Company = r.MultipartForm.Value["company"][0]
	req.Position = r.MultipartForm.Value["position"][0]
	req.Password = r.MultipartForm.Value["password"][0]
	req.ProfilePic = f

	return req, err
}

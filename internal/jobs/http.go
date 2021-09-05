package jobs

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPRouter(e Endpoints, r *mux.Router, options ...httptransport.ServerOption) {
	r.Methods("GET").Path("/").Handler(httptransport.NewServer(
		e.getJobsEndpoint,
		decodeGetJobsRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("POST").Path("/").Handler(httptransport.NewServer(
		e.createJobEndpoint,
		decodeCreateJobRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("POST").Path("/interested").Handler(httptransport.NewServer(
		e.toggleJobInterestedEndpoint,
		decodeToggleJobsInterestedRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("POST").Path("/apply").Handler(httptransport.NewServer(
		e.applyJobEndpoint,
		decodeApplyJobRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("PUT").Path("/").Handler(httptransport.NewServer(
		e.updateJobStatusEndpoint,
		decodeUpdateJobRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))
}

func decodeGetJobsRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return getJobsRequest{}, nil
}

func decodeCreateJobRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req createJobRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeToggleJobsInterestedRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	panic("not implemented") // TODO: Implement
}

func decodeApplyJobRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	panic("not implemented") // TODO: Implement
}

func decodeUpdateJobRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	panic("not implemented") // TODO: Implement
}

package jobs

import (
	"context"
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
		e.postJobEndpoint,
		decodePostJobRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("GET").Path("/interested").Handler(httptransport.NewServer(
		e.getJobsInterestedEndpoint,
		decodeGetJobsInterestedRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("GET").Path("/applied").Handler(httptransport.NewServer(
		e.getJobsAppliedEndpoint,
		decodeGetJobsAppliedRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods("PUT").Path("/change/status").Handler(httptransport.NewServer(
		e.changeJobStatusEndpoint,
		decodeChangeJobStatusRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))
}

func decodeGetJobsRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return getJobsRequest{}, nil
}

func decodePostJobRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return
}

func decodeGetJobsInterestedRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return
}

func decodeGetJobsAppliedRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return
}

func decodeChangeJobStatusRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return
}

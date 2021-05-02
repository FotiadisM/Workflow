package jobs

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	getJobsEndpoint           endpoint.Endpoint
	postJobEndpoint           endpoint.Endpoint
	getJobsInterestedEndpoint endpoint.Endpoint
	getJobsAppliedEndpoint    endpoint.Endpoint
	changeJobStatusEndpoint   endpoint.Endpoint
}

func NewEndpoints(s Service) Endpoints {
	return Endpoints{
		makeGetJobsEndpoint(s),
		makePostJobEndpoint(s),
		makeGetJobsInterestedEndpoint(s),
		makeGetJobsAppliedEndpoint(s),
		makeChangeJobStatusEndpoint(s),
	}
}

func makeGetJobsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getJobsRequest)
		res, err := s.getJobs(ctx, req)

		return res, err
	}
}

func makePostJobEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postJobRequest)
		res, err := s.postJob(ctx, req)

		return res, err
	}
}

func makeGetJobsInterestedEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getJobsInterestedRequest)
		res, err := s.getJobsInterested(ctx, req)

		return res, err
	}
}

func makeGetJobsAppliedEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getJobsAppliedRequest)
		res, err := s.getJobsApplied(ctx, req)

		return res, err
	}
}

func makeChangeJobStatusEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(changeJobStatusRequest)
		res, err := s.changeJobStatus(ctx, req)

		return res, err
	}
}

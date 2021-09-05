package jobs

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	getJobsEndpoint           endpoint.Endpoint
	createJobEndpoint           endpoint.Endpoint
	toggleJobInterestedEndpoint endpoint.Endpoint
	applyJobEndpoint endpoint.Endpoint
	updateJobStatusEndpoint   endpoint.Endpoint
}

func NewEndpoints(s Service) Endpoints {
	return Endpoints{
		makeGetJobsEndpoint(s),
		makeCreateJobEndpoint(s),
		makeToggleJobInterested(s),
		makeApplyJob(s),
		makeUpdateJobEndpoint(s),
	}
}

func makeGetJobsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getJobsRequest)
		res, err := s.getJobs(ctx, req)

		return res, err
	}
}

func makeCreateJobEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createJobRequest)
		res, err := s.createJob(ctx, req)

		return res, err
	}
}

func makeToggleJobInterested(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(toggleJobInterestedRequest)
		res, err := s.toggleJobInterested(ctx, req)

		return res, err
	}
}

func makeApplyJob(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(applyJobRequest)
		res, err := s.applyJob(ctx, req)

		return res, err
	}
}

func makeUpdateJobEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(updateJobRequest)
		res, err := s.updateJob(ctx, req)

		return res, err
	}
}

package jobs

import "context"

type Service interface {
	getJobs(ctx context.Context, req getJobsRequest) (res getJobsResponse, err error)
	postJob(ctx context.Context, req postJobRequest) (res postJobResponse, err error)
	getJobsInterested(ctx context.Context, req getJobsInterestedRequest) (res getJobsInterestedResponse, err error)
	getJobsApplied(ctx context.Context, req getJobsAppliedRequest) (res getJobsAppliedResponse, err error)
	changeJobStatus(ctx context.Context, req changeJobStatusRequest) (res changeJobStatusResponse, err error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return service{r}
}

func (s service) getJobs(ctx context.Context, req getJobsRequest) (res getJobsResponse, err error) {
	res.Name = "test worked!"
	return
}

func (s service) postJob(ctx context.Context, req postJobRequest) (res postJobResponse, err error) {
	return
}

func (s service) getJobsInterested(ctx context.Context, req getJobsInterestedRequest) (res getJobsInterestedResponse, err error) {
	return
}

func (s service) getJobsApplied(ctx context.Context, req getJobsAppliedRequest) (res getJobsAppliedResponse, err error) {
	return
}

func (s service) changeJobStatus(ctx context.Context, req changeJobStatusRequest) (res changeJobStatusResponse, err error) {
	return
}

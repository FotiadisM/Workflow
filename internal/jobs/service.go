package jobs

import (
	"context"
)

type Service interface {
	getJobs(ctx context.Context, req getJobsRequest) (res getJobsResponse, err error)
	createJob(ctx context.Context, req createJobRequest) (res createJobResponse, err error)
	toggleJobInterested(ctx context.Context, req toggleJobInterestedRequest) (res toggleJobInterestedResponse, err error)
	applyJob(ctx context.Context, req applyJobRequest) (res applyJobResponse, err error)
	updateJob(ctx context.Context, req updateJobRequest) (res updateJobResponse, err error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return service{r}
}

func (s service) getJobs(ctx context.Context, req getJobsRequest) (res getJobsResponse, err error) {
	js, err := s.repo.GetJobs(ctx)
	if err != nil {
		res.Err = err
		return
	}

	res.Jobs = js
	return
}

func (s service) createJob(ctx context.Context, req createJobRequest) (res createJobResponse, err error) {
	id, t, err := s.repo.CreateJob(ctx, req.UserID, req.Title, string(req.Type), req.Location, req.Company, req.Description, req.MinSalary, req.MaxSalary, req.Skills)
	if err != nil {
		res.Err = err
		return
	}

	res.ID = id
	res.Created = t.Format("Jan 2")
	return
}

func (s service) toggleJobInterested(ctx context.Context, req toggleJobInterestedRequest) (res toggleJobInterestedResponse, err error) {
	err = s.repo.ToggleJobInterested(ctx, req.UserID, req.JobID)

	res.Err = err
	return
}

func (s service) applyJob(ctx context.Context, req applyJobRequest) (res applyJobResponse, err error) {
	err = s.repo.ApplyJob(ctx, req.UserID, req.JobID)

	res.Err = err
	return
}

func (s service) updateJob(ctx context.Context, req updateJobRequest) (res updateJobResponse, err error) {
	panic("not implemented") // TODO: Implement
}

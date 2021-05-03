package jobs

type getJobsRequest struct{}
type getJobsResponse struct {
	Name string `json:"name"`
}

type postJobRequest struct{}
type postJobResponse struct{}

type getJobsInterestedRequest struct{}
type getJobsInterestedResponse struct{}

type getJobsAppliedRequest struct{}
type getJobsAppliedResponse struct{}

type changeJobStatusRequest struct{}
type changeJobStatusResponse struct{}

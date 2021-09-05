package jobs

type getJobsRequest struct{}
type getJobsResponse struct {
	Jobs []Job `json:"jobs"`
	Err  error `json:"err,omitempty"`
}

type createJobRequest struct {
	UserID      string   `json:"user_id"`
	Title       string   `json:"title"`
	Type        JobType  `json:"type"`
	Location    string   `json:"location"`
	Company     string   `json:"company"`
	MinSalary   float64  `json:"min_salary"`
	MaxSalary   float64  `json:"max_salary"`
	Description string   `json:"description"`
	Skills      []string `json:"skills"`
}
type createJobResponse struct {
	ID      string `json:"id,omitempty"`
	Created string `json:"created,omitempty"`
	Err     error  `json:"err,omitempty"`
}

type toggleJobInterestedRequest struct {
	UserID string `json:"user_id"`
	JobID  string `json:"job_id"`
}
type toggleJobInterestedResponse struct {
	Err error `json:"err,omitempty"`
}

type applyJobRequest struct {
	UserID string `json:"user_id"`
	JobID  string `json:"job_id"`
}
type applyJobResponse struct {
	Err error `json:"err,omitempty"`
}

// TODO: complete update structs
type updateJobRequest struct{}
type updateJobResponse struct{}

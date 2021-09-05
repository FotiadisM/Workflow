package jobs

import (
	"context"
	"time"
)

type JobType string

const (
	FullTime   JobType = "full_time"
	PartTime           = "part_time"
	Internship         = "internship"
)

type Job struct {
	ID          string   `json:"id"`
	UserID      string   `json:"user_id"`
	Title       string   `json:"title"`
	Type        JobType  `json:"type"`
	Location    string   `json:"location"`
	Company     string   `json:"company"`
	MinSalary   float64  `json:"min_salary"`
	MaxSalary   float64  `json:"max_salary"`
	Description string   `json:"description"`
	Skills      []string `json:"skills"`
	Interested  []string `json:"interested"`
	Applied     []string `json:"applied"`
	Created     string   `json:"created"`
}

type Repository interface {
	GetJobs(ctx context.Context) (jobs []Job, err error)
	CreateJob(ctx context.Context, userID, title, jType, location, company, description string, min, max float64, skills []string) (id string, created time.Time, err error)
	ToggleJobInterested(ctx context.Context, userID, jobID string) (err error)
	ApplyJob(ctx context.Context, userID, jobID string) (err error)
}

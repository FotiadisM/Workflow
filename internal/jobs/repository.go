package jobs

type JobType int

const (
	FullTime JobType = iota
	PartTime
	Internship
)

type Job struct {
	ID          string
	UserID      string
	Company     string
	Title       string
	Location    string
	Type        JobType
	Description string
	Skills      []string
	MinSalary   float64
	MaxSalary   float64
}

type Repository interface{}

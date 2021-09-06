package repository

import (
	"context"
	"time"

	"github.com/FotiadisM/workflow-server/internal/jobs"
)

func (r Repository) GetJobs(ctx context.Context) (jobsArr []jobs.Job, err error) {
	rows, err := r.db.Query(ctx, `SELECT * from jobs`)
	if err != nil {
		return
	}
	defer rows.Close()

	jobsArr = []jobs.Job{}
	for rows.Next() {
		var j jobs.Job
		var t time.Time
		if err = rows.Scan(&j.ID,
			&j.UserID,
			&j.Title,
			&j.Type,
			&j.Location,
			&j.Company,
			&j.MinSalary,
			&j.MaxSalary,
			&j.Description,
			&j.Skills,
			&j.Interested,
			&j.Applied,
			&t); err != nil {
			return
		}
		j.Created = t.Format("Jan 2")
		jobsArr = append(jobsArr, j)
	}

	return
}

func (r Repository) CreateJob(ctx context.Context, userID string, title string, jType string, location string, company string, description string, min float64, max float64, skills []string) (id string, created time.Time, err error) {
	err = r.db.QueryRow(ctx, `
	INSERT INTO jobs 
		(user_id, title, type, location, company, min, max, description, skills, interested, applied)
	VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	RETURNING id, created
	`, userID, title, jType, location, company, min, max, description, skills, []string{}, []string{}).Scan(&id, &created)

	return
}

func (r Repository) ToggleJobInterested(ctx context.Context, userID string, jobID string) (err error) {
	indx := []int{}
	err = r.db.QueryRow(ctx, `
	SELECT
		array_positions(interested, $1)
	FROM
		jobs
	WHERE
		id=$2
	;`, userID, jobID).Scan(&indx)
	if err != nil {
		return
	}

	if len(indx) == 0 {
		// append interested
		_, err = r.db.Exec(ctx, `
		UPDATE jobs SET
			interested = array_append(interested, $1)
		WHERE
			id = $2
		;`, userID, jobID)
	} else {
		// remove interested
		_, err = r.db.Exec(ctx, `
		UPDATE jobs SET
			interested = array_remove(interested, $1)
		WHERE
			id = $2
		;`, userID, jobID)
	}

	return
}

func (r Repository) ApplyJob(ctx context.Context, userID string, jobID string) (err error) {
	_, err = r.db.Exec(ctx, `
	UPDATE jobs SET
		applied = array_append(applied, $1)
	WHERE
		id = $2
	;`, userID, jobID)

	return
}

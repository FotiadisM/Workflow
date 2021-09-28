package repository

import (
	"context"

	"github.com/FotiadisM/workflow-server/internal/posts"
)

func (r Repository) UpdateUserFeed(ctx context.Context, userID, perpID, postID, feedType string) (id string, err error) {
	err = r.db.QueryRow(ctx, `
	INSERT INTO feed
		(user_id, post_id, perpetator_id, type)
	VALUES
		($1, $2, $3, $4)
	RETURNING id
	;`, userID, postID, perpID, feedType).Scan(&id)

	return
}

func (r Repository) getFeed(ctx context.Context, userID string) (fs []posts.Feed, err error) {
	rows, err := r.db.Query(ctx, `
	SELECT
		id, post_id, perpetrator_id, type
	FROM
		feed
	WHERE
		user_id=$1
	SORT BY
		created
	LIMIT
		25
		`, &userID)
	if err != nil {
		return
	}
	defer rows.Close()

	fs = []posts.Feed{}
	for rows.Next() {
		f := posts.Feed{}
		err = rows.Scan(&f.ID, &f.PostID, &f.PerpetratorID, &f.FeedType)
		if err != nil {
			return
		}
		fs = append(fs, f)
	}

	return
}

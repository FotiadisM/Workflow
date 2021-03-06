package repository

import (
	"context"
	"time"

	"github.com/FotiadisM/workflow-server/internal/posts"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"
	"github.com/jackc/pgx/v4"
)

func (r Repository) CreatePost(ctx context.Context, userID, text, visibility string, images, videos []string) (id string, created time.Time, err error) {
	err = r.db.QueryRow(ctx, `
	INSERT INTO posts
		(user_id, text, images, videos, visibility, likes, comments)
	VALUES
		($1, $2, $3, $4, $5, $6, $7)
	RETURNING id, created
	;`, userID, text, images, videos, visibility, []string{userID}, []string{}).Scan(&id, &created)

	return
}

func (r Repository) GetUserPosts(ctx context.Context, userID, fromUserID string) (ps []posts.Post, err error) {
	var rows pgx.Rows
	if userID == fromUserID {
		rows, err = r.db.Query(ctx, ` SELECT * FROM posts WHERE user_id=$1`, userID)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var t time.Time
		for rows.Next() {
			p := posts.Post{}
			if err = rows.Scan(&(p.ID), &(p.UserID), &(p.Text), &(p.Images), &(p.Videos), &(p.Visivility), &(p.Likes), &(p.Comments), &t); err != nil {
				return nil, err
			}
			p.Created = t.Format("1/2 15:04")
			ps = append(ps, p)
		}

		return ps, nil
	}

	var fr int
	if fromUserID != "" {
		err = r.db.QueryRow(ctx, `
		SELECT
			COUNT(id)
		FROM
			connections
		WHERE
			(user1_id=$1 AND user2_id=$2)
			OR
			(user1_id=$2 AND user2_id=$1)
	`, userID, fromUserID).Scan(&fr)
		if err != nil {
			return nil, err
		}
	} else {
		fr = 0
	}

	if fr == 1 {
		rows, err = r.db.Query(ctx, ` SELECT * FROM posts WHERE user_id=$1 AND (visibility='friends' OR visibility='all')`, userID)
		defer rows.Close()
	} else {
		rows, err = r.db.Query(ctx, ` SELECT * FROM posts WHERE user_id=$1 AND visibility='all'`, userID)
		defer rows.Close()
	}

	if err != nil {
		return nil, err
	}

	var t time.Time
	for rows.Next() {
		p := posts.Post{}
		if err = rows.Scan(&(p.ID), &(p.UserID), &(p.Text), &(p.Images), &(p.Videos), &(p.Visivility), &(p.Likes), &(p.Comments), &t); err != nil {
			return nil, err
		}
		p.Created = t.Format("1/2 15:04")
		ps = append(ps, p)
	}

	return
}

func (r Repository) GetPost(ctx context.Context, postID string) (p *posts.Post, err error) {
	var t time.Time
	p = &posts.Post{}
	err = r.db.QueryRow(ctx, `
		SELECT * FROM posts WHERE id=$1
	;`, postID).Scan(&(p.ID), &(p.UserID), &(p.Text), &(p.Images), &(p.Videos), &(p.Visivility), &(p.Likes), &(p.Comments), &t)
	p.Created = t.Format("1/2 15:04")

	return
}

func (r Repository) TogglePostLike(ctx context.Context, postID, userID string) (err error) {
	indx := []int{}
	err = r.db.QueryRow(ctx, `
	SELECT
		array_positions(likes, $1)
	FROM
		posts
	WHERE
		id=$2
	;`, userID, postID).Scan(&indx)
	if err != nil {
		return
	}

	if len(indx) == 0 {
		// append interested
		_, err = r.db.Exec(ctx, `
		UPDATE posts SET
			likes = array_append(likes, $1)
		WHERE
			id = $2
		;`, userID, postID)
	} else {
		// remove interested
		_, err = r.db.Exec(ctx, `
		UPDATE posts SET
			likes = array_remove(likes, $1)
		WHERE
			id = $2
		;`, userID, postID)
	}

	return
}

func (r Repository) CreatePostComment(ctx context.Context, postID, userID, text string) (id string, created time.Time, err error) {
	err = crdbpgx.ExecuteTx(ctx, r.db, pgx.TxOptions{}, func(tx pgx.Tx) error {
		err = r.db.QueryRow(ctx, `
			INSERT INTO comments
				(post_id, user_id, text, likes)
			VALUES
				($1, $2, $3, $4)
			RETURNING id, created
		;`, postID, userID, text, []string{userID}).Scan(&id, &created)

		_, err = tx.Exec(ctx, `
			UPDATE posts SET
				comments = array_append(comments, $1)
			WHERE
				id = $2
		`, id, postID)

		return err
	})

	return
}

func (r Repository) GetPostComment(ctx context.Context, commentID string) (c *posts.Comment, err error) {
	var t time.Time
	c = &posts.Comment{}
	err = r.db.QueryRow(ctx, `
		SELECT * FROM Comments WHERE id=$1
	;`, commentID).Scan(&(c.ID), &(c.PostID), &(c.UserID), &(c.Text), &(c.Likes), &t)
	c.Created = t.Format("1/2 15:04")

	return
}

func (r Repository) ToggleCommentLike(ctx context.Context, commentID, userID string) (err error) {
	panic("not implemented")
}

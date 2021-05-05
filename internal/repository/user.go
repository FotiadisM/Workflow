package repository

import (
	"context"

	"github.com/FotiadisM/workflow-server/internal/user"
	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"
	"github.com/jackc/pgx/v4"
)

func (r Repository) GetUserByID(ctx context.Context, id string) (u user.User, err error) {
	err = r.db.QueryRow(ctx, `
	SELECT
		(id, f_name, l_name, email, company, position)
	FROM users WHERE id=$1`, id).Scan(&u.ID, &u.FName, &u.LName, &u.Email, &u.Company, &u.Position)

	return
}

func (r Repository) GetPerpetator(ctx context.Context, id string) (u user.User, err error) {
	err = r.db.QueryRow(ctx, `
	SELECT
		(f_name, l_name, email, company, position)
	FROM users WHERE id=$1`, id).Scan(&u.FName, &u.LName, &u.Email, &u.Company, &u.Position)

	return
}

func (r Repository) GetConnections(ctx context.Context, userID string) (cons []user.Connection, err error) {
	rows, err := r.db.Query(ctx, ` SELECT * FROM connections WHERE user1_id=$1 OR user2_id=$1`)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		con := user.Connection{}
		var us1, us2 string
		if err = rows.Scan(&con.ConnID, &us1, &us2); err != nil {
			return
		}

		switch userID {
		case us1:
			con.UserID = us1
		case us2:
			con.UserID = us2
		}
		cons = append(cons, con)
	}

	return
}

func (r Repository) CreateConnectionRequest(ctx context.Context, userID string, user2ID string) (ConnID string, err error) {
	err = r.db.QueryRow(ctx, `INSEER INTO connection_requests (user1_id, user2_id) VALUES ($1, $2) RETURNING id`, userID, user2ID).Scan(&ConnID)

	return
}

func (r Repository) AcceptConnectionRequest(ctx context.Context, ConnID string) (err error) {
	err = crdbpgx.ExecuteTx(ctx, r.db, pgx.TxOptions{}, func(tx pgx.Tx) error {
		var id, user1ID, user2ID string
		err := tx.QueryRow(ctx, `DELETE FROM connection_requests WHERE id=$1 RETURNING id, user1_id, user2_id`).Scan(&id, &user1ID, &user2ID)
		if err != nil {
			return err
		}

		_, err = tx.Exec(ctx, `INSERT INTO connections VALUES ($1, $2, $3)`, id, user1ID, user2ID)
		if err != nil {
			return err
		}

		return nil
	})

	return
}

func (r Repository) RejectConnectionRequest(ctx context.Context, ConnID string) (err error) {
	_, err = r.db.Exec(ctx, `DELETE FROM connection_requests WHERE id=$1`, ConnID)

	return
}

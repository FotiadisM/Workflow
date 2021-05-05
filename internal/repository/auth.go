package repository

import (
	"context"

	"github.com/FotiadisM/workflow-server/internal/user"
	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

func (r Repository) GetUserPassword(ctx context.Context, email string) (password string, err error) {
	err = r.db.QueryRow(ctx, `SELECT (password) FROM auth WHERE email=$1`, email).Scan(&password)

	return
}

func (r Repository) CreateUser(ctx context.Context, fName, lName, email, password string) (id string, err error) {
	err = crdbpgx.ExecuteTx(ctx, r.db, pgx.TxOptions{}, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, `INSERT INTO auth (email, password) VALUES ($1, $2)`, email, password); err != nil {
			pgErr, ok := err.(*pgconn.PgError)
			if ok {
				if pgErr.Code == "23505" { // duplicate key
					return ErrKeyExists
				}
			}
			return err
		}

		err := tx.QueryRow(ctx, `INSERT INTO users
			(f_name, l_name, email, company, position, role)
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id`,
			fName, lName, email, "-", "-", user.Normal).Scan(&id)

		return err
	})

	return
}

func (r Repository) GetUserByEmail(ctx context.Context, email string) (u user.User, err error) {
	err = r.db.QueryRow(ctx, `SELECT * FROM users WHERE email=$1`, email).Scan(&u.ID, &u.FName, &u.LName, &u.Email, &u.Company, &u.Position, &u.Role)
	return
}

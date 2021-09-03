package repository

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	ErrKeyExists = errors.New("key already exist")
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(ctx context.Context, dbURL string) (r Repository, err error) {

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return
	}

	ctx1, cncl1 := context.WithTimeout(ctx, 2*time.Second)
	defer cncl1()

	pool, err := pgxpool.ConnectConfig(ctx1, config)
	if err != nil {
		return
	}
	r.db = pool

	ctx2, cncl2 := context.WithTimeout(ctx, 2*time.Second)
	defer cncl2()
	if err = r.initDatabase(ctx2); err != nil {
		return
	}

	return
}

func (r Repository) initDatabase(ctx context.Context) (err error) {
	if _, err = r.db.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS auth (
		email STRING(90) PRIMARY KEY,
		password STRING NOT NULL
	);`); err != nil {
		return
	}

	if _, err = r.db.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS users (
		id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
		f_name STRING(60) NOT NULL,
		l_name STRING(60) NOT NULL,
		email STRING REFERENCES public.auth(email),
		company STRING(60) NOT NULL,
		position STRING(60) NOT NULL,
		profile_pic STRING NOT NULL,
		role STRING(60) NOT NULL
	);`); err != nil {
		return
	}

	// connections
	if _, err = r.db.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS connections (
		id UUID UNIQUE,
		user1_id UUID REFERENCES public.users(id),
		user2_id UUID REFERENCES public.users(id),
	  
		CONSTRAINT "primary" PRIMARY KEY (user1_id, user2_id)
	);`); err != nil {
		return
	}

	if _, err = r.db.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS connection_requests (
		id UUID UNIQUE DEFAULT gen_random_uuid(),
		user1_id UUID REFERENCES public.users(id),
		user2_id UUID REFERENCES public.users(id),
  
		CONSTRAINT "primary" PRIMARY KEY (user1_id, user2_id)
	);`); err != nil {
		return
	}

	if _, err = r.db.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS conversations (
		id UUID UNIQUE DEFAULT gen_random_uuid(),
		user1_id UUID REFERENCES public.users(id),
		user2_id UUID REFERENCES public.users(id),
  
		CONSTRAINT "primary" PRIMARY KEY (user1_id, user2_id)
	);`); err != nil {
		return
	}

	if _, err = r.db.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS messages (
		id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
		conv_id UUID REFERENCES public.conversations(id),
		senter_id UUID REFERENCES public.users(id),
		text STRING NOT NULL,
		time_sent TIMESTAMP DEFAULT now()
	);`); err != nil {
		return
	}

	// posts
	if _, err = r.db.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS posts (
		id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
		user_id UUID REFERENCES public.users(id),
		text STRING,
		images STRING[],
		videos STRING[],
		visibility STRING,
		likes STRING[],
		comments STRING[],
		created TIMESTAMP DEFAULT now()
	);`); err != nil {
		return
	}

	if _, err = r.db.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS comments (
		id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
		post_id UUID REFERENCES public.posts(id),
		user_id UUID REFERENCES public.users(id),
		text STRING,
		likes STRING[],
		created TIMESTAMP DEFAULT now()
	);`); err != nil {
		return
	}

	return
}

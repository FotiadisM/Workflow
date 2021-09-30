package repository

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/FotiadisM/workflow-server/internal/posts"
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

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return
	}
	r.db = pool

	if err = r.initDatabase(ctx); err != nil {
		return
	}

	// err = r.populateDB(ctx)
	// if err != nil {
	// 	return r, fmt.Errorf("Failed to populate db: %w", err)
	// }

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
		id UUID UNIQUE DEFAULT gen_random_uuid(),
		user1_id UUID REFERENCES public.users(id),
		user2_id UUID REFERENCES public.users(id),
	  
		CONSTRAINT "primary" PRIMARY KEY (user1_id, user2_id)
	);`); err != nil {
		return
	}

	if _, err = r.db.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS connection_requests (
		id UUID UNIQUE DEFAULT gen_random_uuid(),
		user_id UUID REFERENCES public.users(id),
		receiver_id UUID REFERENCES public.users(id),
  
		CONSTRAINT "primary" PRIMARY KEY (user_id, receiver_id)
	);`); err != nil {
		return
	}

	if _, err = r.db.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS messages (
		id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
		conv_id UUID REFERENCES public.connections(id),
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

	// jobs
	if _, err = r.db.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS jobs (
		id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
		user_id UUID REFERENCES public.users(id),
		title STRING NOT NULL,
		type STRING NOT NULL,
		location STRING NOT NULL,
		company STRING NOT NULL,
		min FLOAT8 NOT NULL,
		max FLOAT8 NOT NULL,
		description STRING,
		skills STRING[],
		interested STRING[],
		applied STRING[],
		created TIMESTAMP DEFAULT now()
	);`); err != nil {
		return
	}

	// experience
	if _, err = r.db.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS experience (
		id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
		user_id UUID REFERENCES public.users(id),
		time_from STRING,
		time_to STRING,
		company STRING,
		position STRING
	);`); err != nil {
		return
	}

	// education
	if _, err = r.db.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS education (
		id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
		user_id UUID REFERENCES public.users(id),
		time STRING,
		title STRING
	);`); err != nil {
		return
	}

	// feed
	// type = "post" | "share" | "comment" | "like"
	if _, err = r.db.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS feed (
		id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
		user_id UUID REFERENCES public.users(id),
		post_id UUID REFERENCES public.posts(id),
		perpetrator_id UUID REFERENCES public.users(id),
		type STRING,
		created TIMESTAMP DEFAULT now()
	);`); err != nil {
		return
	}

	return
}

func (r Repository) populateDB(ctx context.Context) (err error) {
	// USERS

	// 	f, err := os.Open("./mock-data/user_mock_data.csv")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	defer f.Close()

	// 	autSvc := auth.NewService(r)

	// 	cr := csv.NewReader(f)
	// 	var record []string
	// 	for {
	// 		record, err = cr.Read()
	// 		if record == nil {
	// 			break
	// 		}

	// 		req := auth.SignUpRequest{
	// 			FName:      record[0],
	// 			LName:      record[1],
	// 			Email:      record[2],
	// 			Company:    record[3],
	// 			Position:   record[4],
	// 			Password:   "1234",
	// 			ProfilePic: nil,
	// 		}
	// 		res, err := autSvc.SignUp(ctx, req)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		fmt.Println(res.User.ID)

	// 	}
	// 	if err != nil {
	// 		if !errors.Is(err, io.EOF) {
	// 			return
	// 		}
	// 	}

	// POSTS

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	rows, err := r.db.Query(ctx, `SELECT id FROM users`)
	if err != nil {
		return
	}
	defer rows.Close()

	userIDs := []string{}
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			return
		}
		userIDs = append(userIDs, id)
	}

	postsSvc := posts.NewService(r)

	// 	f, err := os.Open("./mock-data/text_mock_data.csv")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	defer f.Close()

	// 	cr := csv.NewReader(f)
	// 	var record []string
	// 	for {
	// 		record, err = cr.Read()
	// 		if record == nil {
	// 			break
	// 		}

	// 		rID := r1.Int() % len(ids)
	// 		req := posts.CreatePostRequest{
	// 			UserID:     ids[rID],
	// 			Text:       record[0],
	// 			Images:     []io.ReadCloser{},
	// 			Videos:     []io.ReadCloser{},
	// 			Visibility: posts.All,
	// 		}

	// 		rID = r1.Int() % 10
	// 		if rID == 0 {
	// 			req.Visibility = posts.Friends
	// 		}
	// 		_, err := postsSvc.CreatePost(ctx, req)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	// 	if err != nil {
	// 		if !errors.Is(err, io.EOF) {
	// 			return
	// 		}
	// 	}

	// LIKES

	rows, err = r.db.Query(ctx, `SELECT id FROM posts`)
	if err != nil {
		return
	}
	defer rows.Close()

	postsIDs := []string{}
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			return
		}
		postsIDs = append(postsIDs, id)
	}

	for index := range postsIDs {
		likesCount := r1.Int() % 250
		for i := 0; i < likesCount; i++ {
			id := r1.Int() % len(userIDs)
			req := posts.TogglePostLikeRequest{
				UserID: userIDs[id],
				PostID: postsIDs[index],
			}
			_, err = postsSvc.TogglePostLike(ctx, req)
			if err != nil {
				return
			}
		}
	}

	return nil
}

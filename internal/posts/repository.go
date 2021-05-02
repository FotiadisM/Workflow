package posts

type Post struct {
	ID      string
	UserID  string
	Created string
	Likes   int
	Text    string
}

type Comment struct {
	ID      string
	PostID  string
	UserID  string
	Created string
	Likes   int
	Text    string
}

type Repository interface {
}

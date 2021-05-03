package user

type Role int

const (
	Admin Role = iota
	Normal
)

type User struct {
	ID       string `json:"id"`
	FName    string `json:"f_name"`
	LName    string `json:"l_name"`
	Email    string `json:"email"`
	Company  string `json:"company"`
	Position string `json:"position"`
	Role     Role   `json:"role"`
}

type Repository interface {
}

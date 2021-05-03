package auth

type signInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signInResponse struct {
	AccessToken  string `json:"access_token"`
	RefressToken string `json:"refress_token"`
	User         struct {
		ID       string `json:"id"`
		FName    string `json:"f_name"`
		LName    string `json:"l_name"`
		Email    string `json:"email"`
		Company  string `json:"company"`
		Position string `json:"position"`
	} `json:"user"`
}

type signUpRequest struct {
	FName    string `json:"f_name"`
	LName    string `json:"l_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signUpResponse struct {
	AccessToken  string `json:"access_token"`
	RefressToken string `json:"refress_token"`
	User         struct {
		ID       string `json:"id"`
		FName    string `json:"f_name"`
		LName    string `json:"l_name"`
		Email    string `json:"email"`
		Company  string `json:"company"`
		Position string `json:"position"`
	} `json:"user"`
}

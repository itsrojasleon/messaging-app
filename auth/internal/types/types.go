package types

type SignupCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninCredentials struct {
	SignupCredentials
}

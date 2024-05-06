package types

type SignupCreds struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninCreds struct {
	SignupCreds
}

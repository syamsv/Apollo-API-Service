package schema

type LoginCreds struct {
	Email    string `json:"email" validate:"required,email" `
	Password string `json:"password" validate:"required,min=6,max=15"`
}

type JwtToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

package entity

import "time"

type SignUpInput struct {
	FirstName       string `json:"name" binding:"required"`
	LastName        string `json:"surname" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type SignInResult struct {
	Role            string `json:"role"`
	AccessToken     string `json:"access_token"`
	RefreshToken    string `json:"refresh_token"`
	AccessTokenAge  int    `json:"access_token_age"`
	RefreshTokenAge int    `json:"refresh_token_age"`
}

type SignUpResult struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"name"`
	LastName  string    `json:"surname"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Message struct {
	Message string `json:"message"`
}

type RefreshTokenResult struct {
	AccessToken    string `json:"access_token"`
	AccessTokenAge int    `json:"access_token_age"`
}

type ForgotPasswordInput struct {
	Email string `json:"email" binding:"required"`
}

type ResetPasswordInput struct {
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

type EmailRegisterInput struct {
	URL       string `json:"url" binding:"required"`
	FirstName string `json:"firstname" binding:"required"`
	Subject   string `json:"subject" binding:"required"`
	To        string `json:"to"`
	Template  string `json:"template"`
}

type CustomResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type CustomResponseWithData struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

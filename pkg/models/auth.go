package models

import "github.com/golang-jwt/jwt/v4"

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Login      string `json:"login"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type RecoveryPasswordRequest struct {
	Email string `json:"email"`
}

type AuthTokenClaims struct {
	jwt.StandardClaims
	AuthUserData
}

type AuthUserData struct {
	UserID     int    `json:"id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	Status     string `json:"status"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Phone     string    `json:"phoneNumber"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SignupRequest struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phoneNumber"`
}

type SigninRequest struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	gorm.Model
	Token string `json:"token"`
	User  User   `json:"user"`
}
type ErrorResponse struct {
	gorm.Model
	Error string `json:"error"`
}

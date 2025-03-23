package models

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}
type User struct {
	gorm.Model
	ID                  int       `json:"id"`
	Email               string    `json:"email" gorm:"unique"`
	Phone               string    `json:"phoneNumber"`
	Password            string    `json:"password"`
	Name                string    `json:"name"`
	Address             Address   `json:"address"`
	DateOfBirth         time.Time `json:"date_of_birth"`
	Nationality         string    `json:"nationality"`
	PassportNumber      string    `json:"passport_number,omitempty"`
	PassportExpiry      time.Time `json:"passport_expiry,omitempty"`
	FrequentFlyerNumber string    `json:"frequent_flyer_number,omitempty"`
	PreferredSeatType   string    `json:"preferred_seat_type,omitempty"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type SigninRequest struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UpdatePassword struct {
	NewPassword string `json:"newPassword"`
}
type AuthResponse struct {
	gorm.Model
	Token string `json:"token"`
	User  User   `json:"user"`
}
type ErrorResponse struct {
	gorm.Model
	Error   string `json:"error"`
	Message string `json:"message"`
}

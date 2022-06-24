package model

import "time"

type User struct {
	ID        *uint      `json:"id,omitempty"`
	Username  string     `json:"username,omitempty"`
	Email     string     `json:"email,omitempty"`
	Password  string     `json:"password,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// Codes for user HttpErrors
const (
	IncorrectCreds        = "ERR_INCORRECT_CREDS"
	UsernameAlreadyExists = "ERR_USERNAME_ALREADY_EXISTS"
	EmailAlreadyExists    = "ERR_EMAIL_ALREADY_EXISTS"
	UserNonExist          = "ERR_USER_NON_EXIST"
)

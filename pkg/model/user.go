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
	ErrIncorrectCreds = "ERR_INCORRECT_CREDS"
)

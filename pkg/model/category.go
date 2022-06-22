package model

import "time"

// Category structure definition
type Category struct {
	ID        *uint      `json:"id,omitempty"`
	UserID    *uint      `json:"userId,omitempty"`
	Title     string     `json:"title,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

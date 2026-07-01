package domain

import "time"

type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"-"`
	GoogleID     string    `json:"google_id"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}
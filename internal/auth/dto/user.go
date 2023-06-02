package dto

import "github.com/google/uuid"

// User is a data transfer object for user.
// It is used for internal communication between packages.
type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Verified  bool   `json:"verified"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// UserID is a user id type.
type UserID struct {
	ID uuid.UUID `json:"id"`
}

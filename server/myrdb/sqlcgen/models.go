// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlcgen

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	CreatedAt time.Time
	ID        uuid.UUID
}

type UserAuthPassword struct {
	CreatedAt    time.Time
	UpdatedAt    time.Time
	UserID       uuid.UUID
	PasswordHash string
}

type UserEmail struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	Email     string
}

type UserEmailMutation struct {
	CreatedAt time.Time
	UserID    uuid.UUID
	Email     string
}

type UserProfile struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	Username  string
	Bio       string
	ImageUrl  string
}

type UserProfileMutation struct {
	CreatedAt time.Time
	UserID    uuid.UUID
	Username  string
	Bio       string
	ImageUrl  string
}
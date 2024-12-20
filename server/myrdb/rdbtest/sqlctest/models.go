// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlctest

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ArticleFavoriteMutationType string

const (
	ArticleFavoriteMutationTypeFavorite   ArticleFavoriteMutationType = "favorite"
	ArticleFavoriteMutationTypeUnfavorite ArticleFavoriteMutationType = "unfavorite"
)

func (e *ArticleFavoriteMutationType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ArticleFavoriteMutationType(s)
	case string:
		*e = ArticleFavoriteMutationType(s)
	default:
		return fmt.Errorf("unsupported scan type for ArticleFavoriteMutationType: %T", src)
	}
	return nil
}

type NullArticleFavoriteMutationType struct {
	ArticleFavoriteMutationType ArticleFavoriteMutationType
	Valid                       bool // Valid is true if ArticleFavoriteMutationType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullArticleFavoriteMutationType) Scan(value interface{}) error {
	if value == nil {
		ns.ArticleFavoriteMutationType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ArticleFavoriteMutationType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullArticleFavoriteMutationType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ArticleFavoriteMutationType), nil
}

type UserFollowMutationType string

const (
	UserFollowMutationTypeFollow   UserFollowMutationType = "follow"
	UserFollowMutationTypeUnfollow UserFollowMutationType = "unfollow"
)

func (e *UserFollowMutationType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserFollowMutationType(s)
	case string:
		*e = UserFollowMutationType(s)
	default:
		return fmt.Errorf("unsupported scan type for UserFollowMutationType: %T", src)
	}
	return nil
}

type NullUserFollowMutationType struct {
	UserFollowMutationType UserFollowMutationType
	Valid                  bool // Valid is true if UserFollowMutationType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserFollowMutationType) Scan(value interface{}) error {
	if value == nil {
		ns.UserFollowMutationType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserFollowMutationType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserFollowMutationType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserFollowMutationType), nil
}

type Article struct {
	CreatedAt time.Time
	ID        uuid.UUID
}

type ArticleComment struct {
	CreatedAt time.Time
	ID        uuid.UUID
}

type ArticleCommentContent struct {
	CreatedAt        time.Time
	ArticleCommentID uuid.UUID
	ArticleID        uuid.UUID
	Body             string
	UserID           uuid.UUID
}

type ArticleCommentContentMutation struct {
	CreatedAt        time.Time
	ArticleCommentID uuid.UUID
	ArticleID        uuid.UUID
	Body             string
	UserID           uuid.UUID
}

type ArticleCommentDeleted struct {
	CreatedAt        time.Time
	ArticleCommentID uuid.UUID
}

type ArticleContent struct {
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ArticleID    uuid.UUID
	Title        string
	Description  string
	Body         string
	AuthorUserID uuid.UUID
}

type ArticleContentMutation struct {
	CreatedAt    time.Time
	ArticleID    uuid.UUID
	Title        string
	Description  string
	Body         string
	AuthorUserID uuid.UUID
}

type ArticleDeleted struct {
	CreatedAt time.Time
	ArticleID uuid.UUID
}

type ArticleFavorite struct {
	CreatedAt time.Time
	ArticleID uuid.UUID
	UserID    uuid.UUID
}

type ArticleFavoriteMutation struct {
	CreatedAt time.Time
	ArticleID uuid.UUID
	UserID    uuid.UUID
	Type      ArticleFavoriteMutationType
}

type ArticleStats struct {
	CreatedAt      time.Time
	UpdatedAt      time.Time
	ArticleID      uuid.UUID
	FavoritesCount int64
}

type ArticleTag struct {
	CreatedAt time.Time
	ArticleID uuid.UUID
	SeqNo     int32
	Tag       string
}

type ArticleTagMutation struct {
	CreatedAt time.Time
	ArticleID uuid.UUID
	Tags      []byte
}

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

type UserFollow struct {
	CreatedAt      time.Time
	UserID         uuid.UUID
	FollowedUserID uuid.UUID
}

type UserFollowMutation struct {
	CreatedAt      time.Time
	UserID         uuid.UUID
	FollowedUserID uuid.UUID
	Type           UserFollowMutationType
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

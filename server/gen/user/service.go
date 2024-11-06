// Code generated by goa v3.19.1, DO NOT EDIT.
//
// user service
//
// Command:
// $ goa gen github.com/mrngsht/realworld-goa-react/design

package user

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// user
type Service interface {
	// Login implements login.
	Login(context.Context, *LoginPayload) (res *LoginResult, err error)
	// Register implements register.
	Register(context.Context, *RegisterPayload) (res *RegisterResult, err error)
	// GetCurrentUser implements getCurrentUser.
	GetCurrentUser(context.Context) (res *GetCurrentUserResult, err error)
	// UpdateUser implements updateUser.
	UpdateUser(context.Context, *UpdateUserPayload) (res *UpdateUserResult, err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "readlworld"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "user"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [4]string{"login", "register", "getCurrentUser", "updateUser"}

// GetCurrentUserResult is the result type of the user service getCurrentUser
// method.
type GetCurrentUserResult struct {
	User *User
}

// LoginPayload is the payload type of the user service login method.
type LoginPayload struct {
	Email    string
	Password string
}

// LoginResult is the result type of the user service login method.
type LoginResult struct {
	User *User
}

// RegisterPayload is the payload type of the user service register method.
type RegisterPayload struct {
	Username string
	Email    string
	Password string
}

// RegisterResult is the result type of the user service register method.
type RegisterResult struct {
	User *User
}

// UpdateUserPayload is the payload type of the user service updateUser method.
type UpdateUserPayload struct {
	Username *string
	Email    *string
	Password *string
	Image    *string
	Bio      *string
}

// UpdateUserResult is the result type of the user service updateUser method.
type UpdateUserResult struct {
	User *User
}

type User struct {
	Email    string
	Token    string
	Username string
	Bio      string
	Image    string
}

// MakeUsernameAlreadyUsed builds a goa.ServiceError from an error.
func MakeUsernameAlreadyUsed(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "UsernameAlreadyUsed", false, false, false)
}

// MakeEmailAlreadyUsed builds a goa.ServiceError from an error.
func MakeEmailAlreadyUsed(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "EmailAlreadyUsed", false, false, false)
}

// MakeEmailNotFound builds a goa.ServiceError from an error.
func MakeEmailNotFound(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "EmailNotFound", false, false, false)
}

// MakePasswordIsIncorrect builds a goa.ServiceError from an error.
func MakePasswordIsIncorrect(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "PasswordIsIncorrect", false, false, false)
}

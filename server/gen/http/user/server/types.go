// Code generated by goa v3.19.1, DO NOT EDIT.
//
// user HTTP server types
//
// Command:
// $ goa gen github.com/mrngsht/realworld-goa-react/design

package server

import (
	"unicode/utf8"

	user "github.com/mrngsht/realworld-goa-react/gen/user"
	goa "goa.design/goa/v3/pkg"
)

// LoginRequestBody is the type of the "user" service "login" endpoint HTTP
// request body.
type LoginRequestBody struct {
	Email    *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// RegisterRequestBody is the type of the "user" service "register" endpoint
// HTTP request body.
type RegisterRequestBody struct {
	Username *string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
	Email    *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// UpdateRequestBody is the type of the "user" service "update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	Username *string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
	Email    *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	Image    *string `form:"image,omitempty" json:"image,omitempty" xml:"image,omitempty"`
	Bio      *string `form:"bio,omitempty" json:"bio,omitempty" xml:"bio,omitempty"`
}

// LoginResponseBody is the type of the "user" service "login" endpoint HTTP
// response body.
type LoginResponseBody struct {
	User *UserResponseBody `form:"user" json:"user" xml:"user"`
}

// RegisterResponseBody is the type of the "user" service "register" endpoint
// HTTP response body.
type RegisterResponseBody struct {
	User *UserResponseBody `form:"user" json:"user" xml:"user"`
}

// GetCurrentResponseBody is the type of the "user" service "getCurrent"
// endpoint HTTP response body.
type GetCurrentResponseBody struct {
	User *UserResponseBody `form:"user" json:"user" xml:"user"`
}

// UpdateResponseBody is the type of the "user" service "update" endpoint HTTP
// response body.
type UpdateResponseBody struct {
	User *UserResponseBody `form:"user" json:"user" xml:"user"`
}

// LoginUserLoginBadRequestResponseBody is the type of the "user" service
// "login" endpoint HTTP response body for the "UserLoginBadRequest" error.
type LoginUserLoginBadRequestResponseBody struct {
	Code string `form:"code" json:"code" xml:"code"`
}

// RegisterUserRegisterBadRequestResponseBody is the type of the "user" service
// "register" endpoint HTTP response body for the "UserRegisterBadRequest"
// error.
type RegisterUserRegisterBadRequestResponseBody struct {
	Code string `form:"code" json:"code" xml:"code"`
}

// UserResponseBody is used to define fields on response body types.
type UserResponseBody struct {
	Email    string `form:"email" json:"email" xml:"email"`
	Token    string `form:"token" json:"token" xml:"token"`
	Username string `form:"username" json:"username" xml:"username"`
	Bio      string `form:"bio" json:"bio" xml:"bio"`
	Image    string `form:"image" json:"image" xml:"image"`
}

// NewLoginResponseBody builds the HTTP response body from the result of the
// "login" endpoint of the "user" service.
func NewLoginResponseBody(res *user.LoginResult) *LoginResponseBody {
	body := &LoginResponseBody{}
	if res.User != nil {
		body.User = marshalUserUserToUserResponseBody(res.User)
	}
	return body
}

// NewRegisterResponseBody builds the HTTP response body from the result of the
// "register" endpoint of the "user" service.
func NewRegisterResponseBody(res *user.RegisterResult) *RegisterResponseBody {
	body := &RegisterResponseBody{}
	if res.User != nil {
		body.User = marshalUserUserToUserResponseBody(res.User)
	}
	return body
}

// NewGetCurrentResponseBody builds the HTTP response body from the result of
// the "getCurrent" endpoint of the "user" service.
func NewGetCurrentResponseBody(res *user.GetCurrentResult) *GetCurrentResponseBody {
	body := &GetCurrentResponseBody{}
	if res.User != nil {
		body.User = marshalUserUserToUserResponseBody(res.User)
	}
	return body
}

// NewUpdateResponseBody builds the HTTP response body from the result of the
// "update" endpoint of the "user" service.
func NewUpdateResponseBody(res *user.UpdateResult) *UpdateResponseBody {
	body := &UpdateResponseBody{}
	if res.User != nil {
		body.User = marshalUserUserToUserResponseBody(res.User)
	}
	return body
}

// NewLoginUserLoginBadRequestResponseBody builds the HTTP response body from
// the result of the "login" endpoint of the "user" service.
func NewLoginUserLoginBadRequestResponseBody(res *user.UserLoginBadRequest) *LoginUserLoginBadRequestResponseBody {
	body := &LoginUserLoginBadRequestResponseBody{
		Code: res.Code,
	}
	return body
}

// NewRegisterUserRegisterBadRequestResponseBody builds the HTTP response body
// from the result of the "register" endpoint of the "user" service.
func NewRegisterUserRegisterBadRequestResponseBody(res *user.UserRegisterBadRequest) *RegisterUserRegisterBadRequestResponseBody {
	body := &RegisterUserRegisterBadRequestResponseBody{
		Code: res.Code,
	}
	return body
}

// NewLoginPayload builds a user service login endpoint payload.
func NewLoginPayload(body *LoginRequestBody) *user.LoginPayload {
	v := &user.LoginPayload{
		Email:    *body.Email,
		Password: *body.Password,
	}

	return v
}

// NewRegisterPayload builds a user service register endpoint payload.
func NewRegisterPayload(body *RegisterRequestBody) *user.RegisterPayload {
	v := &user.RegisterPayload{
		Username: *body.Username,
		Email:    *body.Email,
		Password: *body.Password,
	}

	return v
}

// NewUpdatePayload builds a user service update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody) *user.UpdatePayload {
	v := &user.UpdatePayload{
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
		Image:    body.Image,
		Bio:      body.Bio,
	}

	return v
}

// ValidateLoginRequestBody runs the validations defined on LoginRequestBody
func ValidateLoginRequestBody(body *LoginRequestBody) (err error) {
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Password == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("password", "body"))
	}
	if body.Email != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", *body.Email, goa.FormatEmail))
	}
	if body.Password != nil {
		if utf8.RuneCountInString(*body.Password) < 6 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.password", *body.Password, utf8.RuneCountInString(*body.Password), 6, true))
		}
	}
	if body.Password != nil {
		if utf8.RuneCountInString(*body.Password) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.password", *body.Password, utf8.RuneCountInString(*body.Password), 128, false))
		}
	}
	return
}

// ValidateRegisterRequestBody runs the validations defined on
// RegisterRequestBody
func ValidateRegisterRequestBody(body *RegisterRequestBody) (err error) {
	if body.Username == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("username", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Password == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("password", "body"))
	}
	if body.Username != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.username", *body.Username, "^[a-zA-Z0-9_]{3,32}$"))
	}
	if body.Email != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", *body.Email, goa.FormatEmail))
	}
	if body.Password != nil {
		if utf8.RuneCountInString(*body.Password) < 6 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.password", *body.Password, utf8.RuneCountInString(*body.Password), 6, true))
		}
	}
	if body.Password != nil {
		if utf8.RuneCountInString(*body.Password) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.password", *body.Password, utf8.RuneCountInString(*body.Password), 128, false))
		}
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Username != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.username", *body.Username, "^[a-zA-Z0-9_]{3,32}$"))
	}
	if body.Email != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", *body.Email, goa.FormatEmail))
	}
	if body.Password != nil {
		if utf8.RuneCountInString(*body.Password) < 6 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.password", *body.Password, utf8.RuneCountInString(*body.Password), 6, true))
		}
	}
	if body.Password != nil {
		if utf8.RuneCountInString(*body.Password) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.password", *body.Password, utf8.RuneCountInString(*body.Password), 128, false))
		}
	}
	if body.Image != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.image", *body.Image, "^https?://.+$"))
	}
	if body.Bio != nil {
		if utf8.RuneCountInString(*body.Bio) > 4096 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.bio", *body.Bio, utf8.RuneCountInString(*body.Bio), 4096, false))
		}
	}
	return
}

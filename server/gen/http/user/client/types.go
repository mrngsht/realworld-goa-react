// Code generated by goa v3.19.1, DO NOT EDIT.
//
// user HTTP client types
//
// Command:
// $ goa gen github.com/mrngsht/realworld-goa-react/design

package client

import (
	user "github.com/mrngsht/realworld-goa-react/gen/user"
	goa "goa.design/goa/v3/pkg"
)

// LoginRequestBody is the type of the "user" service "login" endpoint HTTP
// request body.
type LoginRequestBody struct {
	Email    string `form:"email" json:"email" xml:"email"`
	Password string `form:"password" json:"password" xml:"password"`
}

// RegisterRequestBody is the type of the "user" service "register" endpoint
// HTTP request body.
type RegisterRequestBody struct {
	Username string `form:"username" json:"username" xml:"username"`
	Email    string `form:"email" json:"email" xml:"email"`
	Password string `form:"password" json:"password" xml:"password"`
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
	User *UserResponseBody `form:"user,omitempty" json:"user,omitempty" xml:"user,omitempty"`
}

// RegisterResponseBody is the type of the "user" service "register" endpoint
// HTTP response body.
type RegisterResponseBody struct {
	User *UserResponseBody `form:"user,omitempty" json:"user,omitempty" xml:"user,omitempty"`
}

// GetCurrentResponseBody is the type of the "user" service "getCurrent"
// endpoint HTTP response body.
type GetCurrentResponseBody struct {
	User *UserResponseBody `form:"user,omitempty" json:"user,omitempty" xml:"user,omitempty"`
}

// UpdateResponseBody is the type of the "user" service "update" endpoint HTTP
// response body.
type UpdateResponseBody struct {
	User *UserResponseBody `form:"user,omitempty" json:"user,omitempty" xml:"user,omitempty"`
}

// LoginEmailNotFoundResponseBody is the type of the "user" service "login"
// endpoint HTTP response body for the "EmailNotFound" error.
type LoginEmailNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// LoginPasswordIsIncorrectResponseBody is the type of the "user" service
// "login" endpoint HTTP response body for the "PasswordIsIncorrect" error.
type LoginPasswordIsIncorrectResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// RegisterUsernameAlreadyUsedResponseBody is the type of the "user" service
// "register" endpoint HTTP response body for the "UsernameAlreadyUsed" error.
type RegisterUsernameAlreadyUsedResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// RegisterEmailAlreadyUsedResponseBody is the type of the "user" service
// "register" endpoint HTTP response body for the "EmailAlreadyUsed" error.
type RegisterEmailAlreadyUsedResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// UserResponseBody is used to define fields on response body types.
type UserResponseBody struct {
	Email    *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	Token    *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
	Username *string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
	Bio      *string `form:"bio,omitempty" json:"bio,omitempty" xml:"bio,omitempty"`
	Image    *string `form:"image,omitempty" json:"image,omitempty" xml:"image,omitempty"`
}

// NewLoginRequestBody builds the HTTP request body from the payload of the
// "login" endpoint of the "user" service.
func NewLoginRequestBody(p *user.LoginPayload) *LoginRequestBody {
	body := &LoginRequestBody{
		Email:    p.Email,
		Password: p.Password,
	}
	return body
}

// NewRegisterRequestBody builds the HTTP request body from the payload of the
// "register" endpoint of the "user" service.
func NewRegisterRequestBody(p *user.RegisterPayload) *RegisterRequestBody {
	body := &RegisterRequestBody{
		Username: p.Username,
		Email:    p.Email,
		Password: p.Password,
	}
	return body
}

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "update" endpoint of the "user" service.
func NewUpdateRequestBody(p *user.UpdatePayload) *UpdateRequestBody {
	body := &UpdateRequestBody{
		Username: p.Username,
		Email:    p.Email,
		Password: p.Password,
		Image:    p.Image,
		Bio:      p.Bio,
	}
	return body
}

// NewLoginResultOK builds a "user" service "login" endpoint result from a HTTP
// "OK" response.
func NewLoginResultOK(body *LoginResponseBody) *user.LoginResult {
	v := &user.LoginResult{}
	v.User = unmarshalUserResponseBodyToUserUser(body.User)

	return v
}

// NewLoginEmailNotFound builds a user service login endpoint EmailNotFound
// error.
func NewLoginEmailNotFound(body *LoginEmailNotFoundResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewLoginPasswordIsIncorrect builds a user service login endpoint
// PasswordIsIncorrect error.
func NewLoginPasswordIsIncorrect(body *LoginPasswordIsIncorrectResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewRegisterResultOK builds a "user" service "register" endpoint result from
// a HTTP "OK" response.
func NewRegisterResultOK(body *RegisterResponseBody) *user.RegisterResult {
	v := &user.RegisterResult{}
	v.User = unmarshalUserResponseBodyToUserUser(body.User)

	return v
}

// NewRegisterUsernameAlreadyUsed builds a user service register endpoint
// UsernameAlreadyUsed error.
func NewRegisterUsernameAlreadyUsed(body *RegisterUsernameAlreadyUsedResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewRegisterEmailAlreadyUsed builds a user service register endpoint
// EmailAlreadyUsed error.
func NewRegisterEmailAlreadyUsed(body *RegisterEmailAlreadyUsedResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewGetCurrentResultOK builds a "user" service "getCurrent" endpoint result
// from a HTTP "OK" response.
func NewGetCurrentResultOK(body *GetCurrentResponseBody) *user.GetCurrentResult {
	v := &user.GetCurrentResult{}
	v.User = unmarshalUserResponseBodyToUserUser(body.User)

	return v
}

// NewUpdateResultOK builds a "user" service "update" endpoint result from a
// HTTP "OK" response.
func NewUpdateResultOK(body *UpdateResponseBody) *user.UpdateResult {
	v := &user.UpdateResult{}
	v.User = unmarshalUserResponseBodyToUserUser(body.User)

	return v
}

// ValidateLoginResponseBody runs the validations defined on LoginResponseBody
func ValidateLoginResponseBody(body *LoginResponseBody) (err error) {
	if body.User == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user", "body"))
	}
	if body.User != nil {
		if err2 := ValidateUserResponseBody(body.User); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateRegisterResponseBody runs the validations defined on
// RegisterResponseBody
func ValidateRegisterResponseBody(body *RegisterResponseBody) (err error) {
	if body.User == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user", "body"))
	}
	if body.User != nil {
		if err2 := ValidateUserResponseBody(body.User); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateGetCurrentResponseBody runs the validations defined on
// GetCurrentResponseBody
func ValidateGetCurrentResponseBody(body *GetCurrentResponseBody) (err error) {
	if body.User == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user", "body"))
	}
	if body.User != nil {
		if err2 := ValidateUserResponseBody(body.User); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateUpdateResponseBody runs the validations defined on UpdateResponseBody
func ValidateUpdateResponseBody(body *UpdateResponseBody) (err error) {
	if body.User == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user", "body"))
	}
	if body.User != nil {
		if err2 := ValidateUserResponseBody(body.User); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateLoginEmailNotFoundResponseBody runs the validations defined on
// login_EmailNotFound_response_body
func ValidateLoginEmailNotFoundResponseBody(body *LoginEmailNotFoundResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateLoginPasswordIsIncorrectResponseBody runs the validations defined on
// login_PasswordIsIncorrect_response_body
func ValidateLoginPasswordIsIncorrectResponseBody(body *LoginPasswordIsIncorrectResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateRegisterUsernameAlreadyUsedResponseBody runs the validations defined
// on register_UsernameAlreadyUsed_response_body
func ValidateRegisterUsernameAlreadyUsedResponseBody(body *RegisterUsernameAlreadyUsedResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateRegisterEmailAlreadyUsedResponseBody runs the validations defined on
// register_EmailAlreadyUsed_response_body
func ValidateRegisterEmailAlreadyUsedResponseBody(body *RegisterEmailAlreadyUsedResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateUserResponseBody runs the validations defined on UserResponseBody
func ValidateUserResponseBody(body *UserResponseBody) (err error) {
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Token == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("token", "body"))
	}
	if body.Username == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("username", "body"))
	}
	if body.Bio == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("bio", "body"))
	}
	if body.Image == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("image", "body"))
	}
	return
}

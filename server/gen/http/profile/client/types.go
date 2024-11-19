// Code generated by goa v3.19.1, DO NOT EDIT.
//
// profile HTTP client types
//
// Command:
// $ goa gen github.com/mrngsht/realworld-goa-react/design

package client

import (
	profile "github.com/mrngsht/realworld-goa-react/gen/profile"
	goa "goa.design/goa/v3/pkg"
)

// FollowUserRequestBody is the type of the "profile" service "followUser"
// endpoint HTTP request body.
type FollowUserRequestBody struct {
	Username string `form:"username" json:"username" xml:"username"`
}

// UnfollowUserRequestBody is the type of the "profile" service "unfollowUser"
// endpoint HTTP request body.
type UnfollowUserRequestBody struct {
	Username string `form:"username" json:"username" xml:"username"`
}

// FollowUserResponseBody is the type of the "profile" service "followUser"
// endpoint HTTP response body.
type FollowUserResponseBody struct {
	Profile *ProfileResponseBody `form:"profile,omitempty" json:"profile,omitempty" xml:"profile,omitempty"`
}

// UnfollowUserResponseBody is the type of the "profile" service "unfollowUser"
// endpoint HTTP response body.
type UnfollowUserResponseBody struct {
	Profile *ProfileResponseBody `form:"profile,omitempty" json:"profile,omitempty" xml:"profile,omitempty"`
}

// FollowUserUserNotFoundResponseBody is the type of the "profile" service
// "followUser" endpoint HTTP response body for the "UserNotFound" error.
type FollowUserUserNotFoundResponseBody struct {
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

// FollowUserUserAlreadyFollowingResponseBody is the type of the "profile"
// service "followUser" endpoint HTTP response body for the
// "UserAlreadyFollowing" error.
type FollowUserUserAlreadyFollowingResponseBody struct {
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

// FollowUserCannotFollowYourselfResponseBody is the type of the "profile"
// service "followUser" endpoint HTTP response body for the
// "CannotFollowYourself" error.
type FollowUserCannotFollowYourselfResponseBody struct {
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

// UnfollowUserUserNotFoundResponseBody is the type of the "profile" service
// "unfollowUser" endpoint HTTP response body for the "UserNotFound" error.
type UnfollowUserUserNotFoundResponseBody struct {
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

// UnfollowUserUserNotFollowingResponseBody is the type of the "profile"
// service "unfollowUser" endpoint HTTP response body for the
// "UserNotFollowing" error.
type UnfollowUserUserNotFollowingResponseBody struct {
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

// ProfileResponseBody is used to define fields on response body types.
type ProfileResponseBody struct {
	Username  *string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
	Bio       *string `form:"bio,omitempty" json:"bio,omitempty" xml:"bio,omitempty"`
	Image     *string `form:"image,omitempty" json:"image,omitempty" xml:"image,omitempty"`
	Following *bool   `form:"following,omitempty" json:"following,omitempty" xml:"following,omitempty"`
}

// NewFollowUserRequestBody builds the HTTP request body from the payload of
// the "followUser" endpoint of the "profile" service.
func NewFollowUserRequestBody(p *profile.FollowUserPayload) *FollowUserRequestBody {
	body := &FollowUserRequestBody{
		Username: p.Username,
	}
	return body
}

// NewUnfollowUserRequestBody builds the HTTP request body from the payload of
// the "unfollowUser" endpoint of the "profile" service.
func NewUnfollowUserRequestBody(p *profile.UnfollowUserPayload) *UnfollowUserRequestBody {
	body := &UnfollowUserRequestBody{
		Username: p.Username,
	}
	return body
}

// NewFollowUserResultOK builds a "profile" service "followUser" endpoint
// result from a HTTP "OK" response.
func NewFollowUserResultOK(body *FollowUserResponseBody) *profile.FollowUserResult {
	v := &profile.FollowUserResult{}
	v.Profile = unmarshalProfileResponseBodyToProfileProfile(body.Profile)

	return v
}

// NewFollowUserUserNotFound builds a profile service followUser endpoint
// UserNotFound error.
func NewFollowUserUserNotFound(body *FollowUserUserNotFoundResponseBody) *goa.ServiceError {
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

// NewFollowUserUserAlreadyFollowing builds a profile service followUser
// endpoint UserAlreadyFollowing error.
func NewFollowUserUserAlreadyFollowing(body *FollowUserUserAlreadyFollowingResponseBody) *goa.ServiceError {
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

// NewFollowUserCannotFollowYourself builds a profile service followUser
// endpoint CannotFollowYourself error.
func NewFollowUserCannotFollowYourself(body *FollowUserCannotFollowYourselfResponseBody) *goa.ServiceError {
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

// NewUnfollowUserResultOK builds a "profile" service "unfollowUser" endpoint
// result from a HTTP "OK" response.
func NewUnfollowUserResultOK(body *UnfollowUserResponseBody) *profile.UnfollowUserResult {
	v := &profile.UnfollowUserResult{}
	v.Profile = unmarshalProfileResponseBodyToProfileProfile(body.Profile)

	return v
}

// NewUnfollowUserUserNotFound builds a profile service unfollowUser endpoint
// UserNotFound error.
func NewUnfollowUserUserNotFound(body *UnfollowUserUserNotFoundResponseBody) *goa.ServiceError {
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

// NewUnfollowUserUserNotFollowing builds a profile service unfollowUser
// endpoint UserNotFollowing error.
func NewUnfollowUserUserNotFollowing(body *UnfollowUserUserNotFollowingResponseBody) *goa.ServiceError {
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

// ValidateFollowUserResponseBody runs the validations defined on
// FollowUserResponseBody
func ValidateFollowUserResponseBody(body *FollowUserResponseBody) (err error) {
	if body.Profile == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("profile", "body"))
	}
	if body.Profile != nil {
		if err2 := ValidateProfileResponseBody(body.Profile); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateUnfollowUserResponseBody runs the validations defined on
// UnfollowUserResponseBody
func ValidateUnfollowUserResponseBody(body *UnfollowUserResponseBody) (err error) {
	if body.Profile == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("profile", "body"))
	}
	if body.Profile != nil {
		if err2 := ValidateProfileResponseBody(body.Profile); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateFollowUserUserNotFoundResponseBody runs the validations defined on
// followUser_UserNotFound_response_body
func ValidateFollowUserUserNotFoundResponseBody(body *FollowUserUserNotFoundResponseBody) (err error) {
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

// ValidateFollowUserUserAlreadyFollowingResponseBody runs the validations
// defined on followUser_UserAlreadyFollowing_response_body
func ValidateFollowUserUserAlreadyFollowingResponseBody(body *FollowUserUserAlreadyFollowingResponseBody) (err error) {
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

// ValidateFollowUserCannotFollowYourselfResponseBody runs the validations
// defined on followUser_CannotFollowYourself_response_body
func ValidateFollowUserCannotFollowYourselfResponseBody(body *FollowUserCannotFollowYourselfResponseBody) (err error) {
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

// ValidateUnfollowUserUserNotFoundResponseBody runs the validations defined on
// unfollowUser_UserNotFound_response_body
func ValidateUnfollowUserUserNotFoundResponseBody(body *UnfollowUserUserNotFoundResponseBody) (err error) {
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

// ValidateUnfollowUserUserNotFollowingResponseBody runs the validations
// defined on unfollowUser_UserNotFollowing_response_body
func ValidateUnfollowUserUserNotFollowingResponseBody(body *UnfollowUserUserNotFollowingResponseBody) (err error) {
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

// ValidateProfileResponseBody runs the validations defined on
// ProfileResponseBody
func ValidateProfileResponseBody(body *ProfileResponseBody) (err error) {
	if body.Username == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("username", "body"))
	}
	if body.Bio == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("bio", "body"))
	}
	if body.Image == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("image", "body"))
	}
	if body.Following == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("following", "body"))
	}
	return
}

package design

import . "goa.design/goa/v3/dsl"

var _ = Service("user", func() {
	Description("user")

	Error(ErrorUser_UsernameAlreadyUsed)
	Error(ErrorUser_EmailAlreadyUsed)
	Error(ErrorUser_EmailNotFound)
	Error(ErrorUser_PasswordIsIncorrect)

	Method("login", func() {
		HTTP(func() {
			POST("users/login")
			Response(StatusOK)
			Response(ErrorUser_EmailNotFound, StatusBadRequest)
			Response(ErrorUser_PasswordIsIncorrect, StatusBadRequest)
		})

		Payload(func() {
			Required(
				AttributeUser_RequestEmail(),
				AttributeUser_RequestPassword(),
			)
		})

		Result(func() {
			Required(
				AttributeWithName("user", UserType),
			)
		})
	})

	Method("register", func() {
		HTTP(func() {
			POST("users")
			Response(StatusOK)
			Response(ErrorUser_UsernameAlreadyUsed, StatusBadRequest)
			Response(ErrorUser_EmailAlreadyUsed, StatusBadRequest)
		})

		Payload(func() {
			Required(
				AttributeUser_RequestUsername(),
				AttributeUser_RequestEmail(),
				AttributeUser_RequestPassword(),
			)
		})

		Result(func() {
			Required(
				AttributeWithName("user", UserType),
			)
		})
	})
})

const (
	ErrorUser_UsernameAlreadyUsed = "UsernameAlreadyUsed"
	ErrorUser_EmailAlreadyUsed    = "EmailAlreadyUsed"
	ErrorUser_EmailNotFound       = "EmailNotFound"
	ErrorUser_PasswordIsIncorrect = "PasswordIsIncorrect"
)

var (
	AttributeUser_RequestUsername = func() string {
		return AttributeWithName("username", String, func() {
			Pattern(`^[a-z0-9_]{3, 32}$`)
		})
	}
	AttributeUser_RequestEmail = func() string {
		return AttributeWithName("email", String, func() {
			Format(FormatEmail)
		})
	}
	AttributeUser_RequestPassword = func() string {
		return AttributeWithName("password", String, func() {
			MinLength(6)
			MaxLength(128)
		})
	}
)

var UserType = Type("UserType", func() {
	Required(
		AttributeWithName("email", String),
		AttributeWithName("token", String),
		AttributeWithName("username", String),
		AttributeWithName("bio", String),
		AttributeWithName("image", String),
	)
})

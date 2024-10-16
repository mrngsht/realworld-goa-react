package design

import . "goa.design/goa/v3/dsl"

const (
	ErrorUserUsernameAlreadyUsed = "UsernameAlreadyUsed"
	ErrorUserEmailAlreadyUsed    = "EmailAlreadyUsed"
)

var (
	AttributeUserRequestUsername = func() string {
		return AttributeWithName("username", String, func() {
			Pattern(`^[a-z0-9_]{3, 32}$`)
		})
	}
	AttributeUserRequestEmail = func() string {
		return AttributeWithName("email", String, func() {
			Format(FormatEmail)
		})
	}
	AttributeUserRequestPassword = func() string {
		return AttributeWithName("password", String, func() {
			MinLength(6)
			MaxLength(128)
		})
	}
)

var _ = Service("user", func() {
	Description("user")

	Error(ErrorUserUsernameAlreadyUsed)
	Error(ErrorUserEmailAlreadyUsed)

	Method("login", func() {
		HTTP(func() {
			POST("users/login")
			Response(StatusOK)
		})

		Payload(func() {
			Required(
				AttributeWithName("email", String, func() {
					Format(FormatEmail)
				}),
				AttributeWithName("password", String),
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
			Response(ErrorUserUsernameAlreadyUsed, StatusBadRequest)
			Response(ErrorUserEmailAlreadyUsed, StatusBadRequest)
		})

		Payload(func() {
			Required(
				AttributeUserRequestUsername(),
				AttributeUserRequestEmail(),
				AttributeUserRequestPassword(),
			)
		})

		Result(func() {
			Required(
				AttributeWithName("user", UserType),
			)
		})
	})
})

var UserType = Type("UserType", func() {
	Required(
		AttributeWithName("email", String),
		AttributeWithName("token", String),
		AttributeWithName("username", String),
		AttributeWithName("bio", String),
		AttributeWithName("image", String),
	)
})

package design

import . "goa.design/goa/v3/dsl"

var _ = Service("user", func() {
	Description("user")

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
		})

		Payload(func() {
			Required(
				AttributeWithName("username", String),
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
})

var UserType = Type("UserType", func() {
	Required(
		AttributeWithName("email", String, func() {
			Format(FormatEmail)
		}),
		AttributeWithName("token", String),
		AttributeWithName("username", String),
		AttributeWithName("bio", String),
		AttributeWithName("image", String),
	)
})

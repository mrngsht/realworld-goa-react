package design

import . "goa.design/goa/v3/dsl"

var _ = Service("user", func() {
	Description("user")

	HTTP(func() {
		Path("users")
	})

	Method("login", func() {
		HTTP(func() {
			POST("/login")
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

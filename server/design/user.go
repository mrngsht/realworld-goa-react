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
			POST("user/login")
			Response(StatusOK)
			Response(ErrorUser_EmailNotFound, StatusBadRequest)
			Response(ErrorUser_PasswordIsIncorrect, StatusBadRequest)
		})

		Payload(func() {
			Required(
				AttributeWithName("email", String, DefUser_RequestEmail),
				AttributeWithName("password", String, DefUser_RequestPassword),
			)
		})

		Result(func() {
			Required(
				AttributeWithName("user", Type_User),
			)
		})
	})

	Method("register", func() {
		HTTP(func() {
			POST("user/register")
			Response(StatusOK)
			Response(ErrorUser_UsernameAlreadyUsed, StatusBadRequest)
			Response(ErrorUser_EmailAlreadyUsed, StatusBadRequest)
		})

		Payload(func() {
			Required(
				AttributeWithName("username", String, DefUser_RequestUsername),
				AttributeWithName("email", String, DefUser_RequestEmail),
				AttributeWithName("password", String, DefUser_RequestPassword),
			)
		})

		Result(func() {
			Required(
				AttributeWithName("user", Type_User),
			)
		})
	})

	Method("getCurrentUser", func() {
		HTTP(func() {
			GET("user/current")
			Response(StatusOK)
		})

		Result(func() {
			Required(
				AttributeWithName("user", Type_User),
			)
		})
	})

	Method("updateUser", func() {
		HTTP(func() {
			POST("user/update")
			Response(StatusOK)
		})

		Payload(func() {
			AttributeWithName("username", String, DefUser_RequestUsername)
			AttributeWithName("email", String, DefUser_RequestEmail)
			AttributeWithName("password", String, DefUser_RequestPassword)
			AttributeWithName("image", String, DefUser_RequestImage)
			AttributeWithName("bio", String, DefUser_RequestBio)
		})

		Result(func() {
			Required(
				AttributeWithName("user", Type_User),
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
	DefUser_RequestUsername = func() {
		Pattern(`^[a-zA-Z0-9_]{3,32}$`)
	}
	DefUser_RequestEmail = func() {
		Format(FormatEmail)
	}
	DefUser_RequestPassword = func() {
		MinLength(6)
		MaxLength(128)
	}
	DefUser_RequestImage = func() {
		Pattern(`^http://.+$`)
	}
	DefUser_RequestBio = func() {
		MaxLength(4096)
	}
)

var (
	Type_User = Type("User", func() {
		Required(
			AttributeWithName("email", String),
			AttributeWithName("token", String),
			AttributeWithName("username", String),
			AttributeWithName("bio", String),
			AttributeWithName("image", String),
		)
	})
)

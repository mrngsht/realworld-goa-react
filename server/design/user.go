package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("user", func() {
	Description("user")

	Method("login", func() {
		HTTP(func() {
			POST("user/login")
			Response(StatusOK)
			Response(errType_UserLoginBadRequest.Name(), StatusBadRequest)
		})

		Payload(func() {
			Required(
				AttributeWithName("email", String, def_User_RequestEmail),
				AttributeWithName("password", String, def_User_RequestPassword),
			)
		})

		Result(func() {
			Required(
				AttributeWithName("user", type_User),
			)
		})

		Error(errType_UserLoginBadRequest.Name(), errType_UserLoginBadRequest)
	})

	Method("register", func() {
		HTTP(func() {
			POST("user/register")
			Response(StatusOK)
			Response(errType_UserRegisterBadRequest.Name(), StatusBadRequest)
		})

		Payload(func() {
			Required(
				AttributeWithName("username", String, def_User_RequestUsername),
				AttributeWithName("email", String, def_User_RequestEmail),
				AttributeWithName("password", String, def_User_RequestPassword),
			)
		})

		Result(func() {
			Required(
				AttributeWithName("user", type_User),
			)
		})

		Error(errType_UserRegisterBadRequest.Name(), errType_UserRegisterBadRequest)
	})

	Method("getCurrent", func() {
		HTTP(func() {
			GET("user/current")
			Response(StatusOK)
		})

		Result(func() {
			Required(
				AttributeWithName("user", type_User),
			)
		})
	})

	Method("update", func() {
		HTTP(func() {
			POST("user/update")
			Response(StatusOK)
		})

		Payload(func() {
			AttributeWithName("username", String, def_User_RequestUsername)
			AttributeWithName("email", String, def_User_RequestEmail)
			AttributeWithName("password", String, def_User_RequestPassword)
			AttributeWithName("image", String, def_User_RequestImage)
			AttributeWithName("bio", String, def_User_RequestBio)
		})

		Result(func() {
			Required(
				AttributeWithName("user", type_User),
			)
		})
	})

})

var (
	def_User_RequestUsername = func() {
		Pattern(`^[a-zA-Z0-9_]{3,32}$`)
	}
	def_User_RequestEmail = func() {
		Format(FormatEmail)
	}
	def_User_RequestPassword = func() {
		MinLength(6)
		MaxLength(128)
	}
	def_User_RequestImage = func() {
		Pattern(`^https?://.+$`)
	}
	def_User_RequestBio = func() {
		MaxLength(4096)
	}
)

var (
	type_User = Type("User", func() {
		Required(
			AttributeWithName("email", String),
			AttributeWithName("token", String),
			AttributeWithName("username", String),
			AttributeWithName("bio", String),
			AttributeWithName("image", String),
		)
	})
)

var (
	errType_UserLoginBadRequest = myErrorType("UserLoginBadRequest", []any{
		ErrCode_User_EmailNotFound,
		ErrCode_User_PasswordIsIncorrect,
	}, nil)

	errType_UserRegisterBadRequest = myErrorType("UserRegisterBadRequest", []any{
		ErrCode_User_UsernameAlreadyUsed,
		ErrCode_User_EmailAlreadyUsed,
	}, nil)
)

const (
	ErrCode_User_EmailNotFound       = "EmailNotFound"
	ErrCode_User_PasswordIsIncorrect = "PasswordIsIncorrect"
	ErrCode_User_UsernameAlreadyUsed = "UsernameAlreadyUsed"
	ErrCode_User_EmailAlreadyUsed    = "EmailAlreadyUsed"
)

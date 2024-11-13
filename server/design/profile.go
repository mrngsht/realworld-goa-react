package design

import . "goa.design/goa/v3/dsl"

var _ = Service("profile", func() {
	Description("profile")

	Error(ErrorProfile_UserNotFound)
	Error(ErrorProfile_UserAlreadyFollowing)
	Error(ErrorProfile_UserNotFollowing)

	Method("followUser", func() {
		HTTP(func() {
			POST("profile/follow_user")
			Response(StatusOK)
			Response(ErrorProfile_UserNotFound, StatusBadRequest)
			Response(ErrorProfile_UserAlreadyFollowing, StatusBadRequest)
		})

		Payload(func() {
			Required(
				AttributeWithName("username", String, DefUser_RequestUsername),
			)
		})

		Result(func() {
			Required(
				AttributeWithName("profile", Type_Profile),
			)
		})
	})

	Method("unfollowUser", func() {
		HTTP(func() {
			POST("profile/unfollow_user")
			Response(StatusOK)
			Response(ErrorProfile_UserNotFound, StatusBadRequest)
			Response(ErrorProfile_UserNotFollowing, StatusBadRequest)
		})

		Payload(func() {
			Required(
				AttributeWithName("username", String, DefUser_RequestUsername),
			)
		})

		Result(func() {
			Required(
				AttributeWithName("profile", Type_Profile),
			)
		})
	})

})

const (
	ErrorProfile_UserNotFound         = "UserNotFound"
	ErrorProfile_UserAlreadyFollowing = "UserAlreadyFollowing"
	ErrorProfile_UserNotFollowing     = "UserNotFollowing"
)

var (
	Type_Profile = Type("Profile", func() {
		Required(
			AttributeWithName("username", String),
			AttributeWithName("bio", String),
			AttributeWithName("image", String),
			AttributeWithName("following", Boolean),
		)
	})
)

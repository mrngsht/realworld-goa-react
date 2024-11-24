package design

import . "goa.design/goa/v3/dsl"

var _ = Service("profile", func() {
	Description("profile")

	Method("followUser", func() {
		HTTP(func() {
			POST("profile/follow_user")
			Response(StatusOK)
			Response(errType_ProfileFollowUserBadRequest.Name(), StatusBadRequest)
		})

		Payload(func() {
			Required(
				AttributeWithName("username", String, def_User_RequestUsername),
			)
		})

		Result(func() {
			Required(
				AttributeWithName("profile", type_Profile),
			)
		})

		Error(errType_ProfileFollowUserBadRequest.Name(), errType_ProfileFollowUserBadRequest)
	})

	Method("unfollowUser", func() {
		HTTP(func() {
			POST("profile/unfollow_user")
			Response(StatusOK)
			Response(errType_ProfileUnfollowUserBadRequest.Name(), StatusBadRequest)
		})

		Payload(func() {
			Required(
				AttributeWithName("username", String, def_User_RequestUsername),
			)
		})

		Result(func() {
			Required(
				AttributeWithName("profile", type_Profile),
			)
		})

		Error(errType_ProfileUnfollowUserBadRequest.Name(), errType_ProfileUnfollowUserBadRequest)
	})

})

var (
	type_Profile = Type("Profile", func() {
		Required(
			AttributeWithName("username", String),
			AttributeWithName("bio", String),
			AttributeWithName("image", String),
			AttributeWithName("following", Boolean),
		)
	})
)

var (
	errType_ProfileFollowUserBadRequest = myErrorType("ProfileFollowUserBadRequest", []any{
		ErrCode_Profile_UserNotFound,
		ErrCode_Profile_UserAlreadyFollowing,
		ErrCode_Profile_UserCannotFollowYourself,
	}, nil)

	errType_ProfileUnfollowUserBadRequest = myErrorType("ProfileUnfollowUserBadRequest", []any{
		ErrCode_Profile_UserNotFound,
		ErrCode_Profile_UserNotFollowing,
	}, nil)
)

const (
	ErrCode_Profile_UserNotFound             = "UserNotFound"
	ErrCode_Profile_UserAlreadyFollowing     = "UserAlreadyFollowing"
	ErrCode_Profile_UserNotFollowing         = "UserNotFollowing"
	ErrCode_Profile_UserCannotFollowYourself = "CannotFollowYourself"
)

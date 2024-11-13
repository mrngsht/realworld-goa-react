// Code generated by goa v3.19.1, DO NOT EDIT.
//
// user HTTP client CLI support package
//
// Command:
// $ goa gen github.com/mrngsht/realworld-goa-react/design

package client

import (
	"encoding/json"
	"fmt"
	"unicode/utf8"

	user "github.com/mrngsht/realworld-goa-react/gen/user"
	goa "goa.design/goa/v3/pkg"
)

// BuildLoginPayload builds the payload for the user login endpoint from CLI
// flags.
func BuildLoginPayload(userLoginBody string) (*user.LoginPayload, error) {
	var err error
	var body LoginRequestBody
	{
		err = json.Unmarshal([]byte(userLoginBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"ophelia@stehr.biz\",\n      \"password\": \"sm4\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", body.Email, goa.FormatEmail))
		if utf8.RuneCountInString(body.Password) < 6 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.password", body.Password, utf8.RuneCountInString(body.Password), 6, true))
		}
		if utf8.RuneCountInString(body.Password) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.password", body.Password, utf8.RuneCountInString(body.Password), 128, false))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &user.LoginPayload{
		Email:    body.Email,
		Password: body.Password,
	}

	return v, nil
}

// BuildRegisterPayload builds the payload for the user register endpoint from
// CLI flags.
func BuildRegisterPayload(userRegisterBody string) (*user.RegisterPayload, error) {
	var err error
	var body RegisterRequestBody
	{
		err = json.Unmarshal([]byte(userRegisterBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"raymond@roobschimmel.biz\",\n      \"password\": \"uip\",\n      \"username\": \"q4I\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidatePattern("body.username", body.Username, "^[a-zA-Z0-9_]{3,32}$"))
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", body.Email, goa.FormatEmail))
		if utf8.RuneCountInString(body.Password) < 6 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.password", body.Password, utf8.RuneCountInString(body.Password), 6, true))
		}
		if utf8.RuneCountInString(body.Password) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.password", body.Password, utf8.RuneCountInString(body.Password), 128, false))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &user.RegisterPayload{
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
	}

	return v, nil
}

// BuildUpdatePayload builds the payload for the user update endpoint from CLI
// flags.
func BuildUpdatePayload(userUpdateBody string) (*user.UpdatePayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(userUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"bio\": \"w3n\",\n      \"email\": \"eleonore@reichel.biz\",\n      \"image\": \"https://o1\",\n      \"password\": \"1po\",\n      \"username\": \"05i1T\"\n   }'")
		}
	}
	v := &user.UpdatePayload{
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
		Image:    body.Image,
		Bio:      body.Bio,
	}

	return v, nil
}

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
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"hermann.swift@bernhard.com\",\n      \"password\": \"Pariatur saepe placeat a perferendis occaecati assumenda.\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", body.Email, goa.FormatEmail))
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

// Code generated by goa v3.19.1, DO NOT EDIT.
//
// realworld HTTP client CLI support package
//
// Command:
// $ goa gen github.com/mrngsht/realworld-goa-react/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	profilec "github.com/mrngsht/realworld-goa-react/gen/http/profile/client"
	userc "github.com/mrngsht/realworld-goa-react/gen/http/user/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `profile (follow-user|unfollow-user)
user (login|register|get-current|update)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` profile follow-user --body '{
      "username": "zn5"
   }'` + "\n" +
		os.Args[0] + ` user login --body '{
      "email": "ophelia@stehr.biz",
      "password": "sm4"
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, any, error) {
	var (
		profileFlags = flag.NewFlagSet("profile", flag.ContinueOnError)

		profileFollowUserFlags    = flag.NewFlagSet("follow-user", flag.ExitOnError)
		profileFollowUserBodyFlag = profileFollowUserFlags.String("body", "REQUIRED", "")

		profileUnfollowUserFlags    = flag.NewFlagSet("unfollow-user", flag.ExitOnError)
		profileUnfollowUserBodyFlag = profileUnfollowUserFlags.String("body", "REQUIRED", "")

		userFlags = flag.NewFlagSet("user", flag.ContinueOnError)

		userLoginFlags    = flag.NewFlagSet("login", flag.ExitOnError)
		userLoginBodyFlag = userLoginFlags.String("body", "REQUIRED", "")

		userRegisterFlags    = flag.NewFlagSet("register", flag.ExitOnError)
		userRegisterBodyFlag = userRegisterFlags.String("body", "REQUIRED", "")

		userGetCurrentFlags = flag.NewFlagSet("get-current", flag.ExitOnError)

		userUpdateFlags    = flag.NewFlagSet("update", flag.ExitOnError)
		userUpdateBodyFlag = userUpdateFlags.String("body", "REQUIRED", "")
	)
	profileFlags.Usage = profileUsage
	profileFollowUserFlags.Usage = profileFollowUserUsage
	profileUnfollowUserFlags.Usage = profileUnfollowUserUsage

	userFlags.Usage = userUsage
	userLoginFlags.Usage = userLoginUsage
	userRegisterFlags.Usage = userRegisterUsage
	userGetCurrentFlags.Usage = userGetCurrentUsage
	userUpdateFlags.Usage = userUpdateUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "profile":
			svcf = profileFlags
		case "user":
			svcf = userFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "profile":
			switch epn {
			case "follow-user":
				epf = profileFollowUserFlags

			case "unfollow-user":
				epf = profileUnfollowUserFlags

			}

		case "user":
			switch epn {
			case "login":
				epf = userLoginFlags

			case "register":
				epf = userRegisterFlags

			case "get-current":
				epf = userGetCurrentFlags

			case "update":
				epf = userUpdateFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "profile":
			c := profilec.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "follow-user":
				endpoint = c.FollowUser()
				data, err = profilec.BuildFollowUserPayload(*profileFollowUserBodyFlag)
			case "unfollow-user":
				endpoint = c.UnfollowUser()
				data, err = profilec.BuildUnfollowUserPayload(*profileUnfollowUserBodyFlag)
			}
		case "user":
			c := userc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "login":
				endpoint = c.Login()
				data, err = userc.BuildLoginPayload(*userLoginBodyFlag)
			case "register":
				endpoint = c.Register()
				data, err = userc.BuildRegisterPayload(*userRegisterBodyFlag)
			case "get-current":
				endpoint = c.GetCurrent()
			case "update":
				endpoint = c.Update()
				data, err = userc.BuildUpdatePayload(*userUpdateBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// profileUsage displays the usage of the profile command and its subcommands.
func profileUsage() {
	fmt.Fprintf(os.Stderr, `profile
Usage:
    %[1]s [globalflags] profile COMMAND [flags]

COMMAND:
    follow-user: FollowUser implements followUser.
    unfollow-user: UnfollowUser implements unfollowUser.

Additional help:
    %[1]s profile COMMAND --help
`, os.Args[0])
}
func profileFollowUserUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] profile follow-user -body JSON

FollowUser implements followUser.
    -body JSON: 

Example:
    %[1]s profile follow-user --body '{
      "username": "zn5"
   }'
`, os.Args[0])
}

func profileUnfollowUserUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] profile unfollow-user -body JSON

UnfollowUser implements unfollowUser.
    -body JSON: 

Example:
    %[1]s profile unfollow-user --body '{
      "username": "C2PiK"
   }'
`, os.Args[0])
}

// userUsage displays the usage of the user command and its subcommands.
func userUsage() {
	fmt.Fprintf(os.Stderr, `user
Usage:
    %[1]s [globalflags] user COMMAND [flags]

COMMAND:
    login: Login implements login.
    register: Register implements register.
    get-current: GetCurrent implements getCurrent.
    update: Update implements update.

Additional help:
    %[1]s user COMMAND --help
`, os.Args[0])
}
func userLoginUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] user login -body JSON

Login implements login.
    -body JSON: 

Example:
    %[1]s user login --body '{
      "email": "ophelia@stehr.biz",
      "password": "sm4"
   }'
`, os.Args[0])
}

func userRegisterUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] user register -body JSON

Register implements register.
    -body JSON: 

Example:
    %[1]s user register --body '{
      "email": "raymond@roobschimmel.biz",
      "password": "uip",
      "username": "q4I"
   }'
`, os.Args[0])
}

func userGetCurrentUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] user get-current

GetCurrent implements getCurrent.

Example:
    %[1]s user get-current
`, os.Args[0])
}

func userUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] user update -body JSON

Update implements update.
    -body JSON: 

Example:
    %[1]s user update --body '{
      "bio": "w3n",
      "email": "eleonore@reichel.biz",
      "image": "https://o1",
      "password": "1po",
      "username": "05i1T"
   }'
`, os.Args[0])
}

package design

import . "goa.design/goa/v3/dsl"

var _ = API("readlworld", func() {
	Title("readworld app example")
	Description("readworld app example")

	Error(ErrorCommon_AuthenticationRequired)

	HTTP(func() {
		Path("api")
		Response(ErrorCommon_AuthenticationRequired, StatusUnauthorized)
	})

	Server("realworld", func() {
		Host("localhost", func() { URI("http://localhost:8080") })
	})
})

const (
	ErrorCommon_AuthenticationRequired = "AuthenticationRequired"
)

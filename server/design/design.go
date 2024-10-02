package design

import . "goa.design/goa/v3/dsl"

var _ = API("readlworld", func() {
	Title("readworld app example")
	Description("readworld app example")

	HTTP(func() {
		Path("api")
	})

	Server("realworld", func() {
		Host("localhost", func() { URI("http://localhost:8080") })
	})
})

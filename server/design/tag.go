package design

import . "goa.design/goa/v3/dsl"

var _ = Service("tag", func() {
	Description("tag")

	Method("getTags", func() {
		HTTP(func() {
			GET("tags")
			Response(StatusOK)
		})

		Result(func() {
			Required(
				AttributeWithName("tags", ArrayOf(String)),
			)
		})
	})
})

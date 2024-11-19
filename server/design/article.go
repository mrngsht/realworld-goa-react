package design

import . "goa.design/goa/v3/dsl"

var _ = Service("article", func() {
	Description("article")

	Method("create", func() {
		HTTP(func() {
			POST("article/create")
			Response(StatusOK)
		})

		Payload(func() {
			Required(
				AttributeWithName("title", String, DefArticle_RequestTitle),
				AttributeWithName("description", String),
				AttributeWithName("body", String),
				AttributeWithName("tagList", ArrayOf(String)),
			)
		})

		Result(func() {
			Required(
				AttributeWithName("article", Type_ArticleDetail),
			)
		})
	})
})

var (
	DefArticle_RequestTitle = func() {
		MaxLength(128)
	}
)

var (
	Type_ArticleDetail = Type("ArticleDetail", func() {
		Required(
			AttributeWithName("id", String, func() {
				Format(FormatUUID)
			}),
			AttributeWithName("title", String),
			AttributeWithName("description", String),
			AttributeWithName("body", String),
			AttributeWithName("tagList", ArrayOf(String)),
			AttributeWithName("createdAt", String, func() {
				Format(FormatDateTime)
			}),
			AttributeWithName("updatedAt", String, func() {
				Format(FormatDateTime)
			}),
			AttributeWithName("favorited", Boolean),
			AttributeWithName("favoritesCount", UInt),
			AttributeWithName("author", Type_Profile),
		)
	})
)

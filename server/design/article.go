package design

import . "goa.design/goa/v3/dsl"

var _ = Service("article", func() {
	Description("article")

	Method("get", func() {
		HTTP(func() {
			GET("article/{articleId}")
			Response(StatusOK)
		})

		Payload(func() {
			Required(
				AttributeWithName("articleId", String, DefArticle_RequestArticleID),
			)
		})

		Result(func() {
			Required(
				AttributeWithName("article", Type_ArticleDetail),
			)
		})
	})

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
	DefArticle_RequestArticleID = func() {
		Format(FormatUUID)
	}
)

var (
	Type_ArticleDetail = Type("ArticleDetail", func() {
		Required(
			AttributeWithName("articleId", String, func() {
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

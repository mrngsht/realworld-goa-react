package design

import . "goa.design/goa/v3/dsl"

var _ = Service("article", func() {
	Description("article")

	Method("get", func() {
		HTTP(func() {
			GET("article/{articleId}")
			Response(StatusOK)
			Response(errType_ArticleGetArticleBadRequest.Name(), StatusBadRequest)
		})

		Payload(func() {
			Required(
				AttributeWithName("articleId", String, def_Article_RequestArticleID),
			)
		})

		Result(func() {
			Required(
				AttributeWithName("article", type_ArticleDetail),
			)
		})

		Error(errType_ArticleGetArticleBadRequest.Name(), errType_ArticleGetArticleBadRequest)
	})

	Method("create", func() {
		HTTP(func() {
			POST("article/create")
			Response(StatusOK)
		})

		Payload(func() {
			Required(
				AttributeWithName("title", String, def_Article_RequestTitle),
				AttributeWithName("description", String),
				AttributeWithName("body", String),
				AttributeWithName("tagList", ArrayOf(String)),
			)
		})

		Result(func() {
			Required(
				AttributeWithName("article", type_ArticleDetail),
			)
		})
	})

	Method("favorite", func() {
		HTTP(func() {
			POST("article/{articleId}/favorite")
			Response(StatusOK)
		})

		Payload(func() {
			Required(
				AttributeWithName("articleId", String, def_Article_RequestArticleID),
			)
		})

		Result(func() {
			Required(
				AttributeWithName("article", type_ArticleDetail),
			)
		})
	})
})

var (
	def_Article_RequestTitle = func() {
		MaxLength(128)
	}
	def_Article_RequestArticleID = func() {
		Format(FormatUUID)
	}
)

var (
	type_ArticleDetail = Type("ArticleDetail", func() {
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
			AttributeWithName("author", type_Profile),
		)
	})
)

var (
	errType_ArticleGetArticleBadRequest = myErrorType("ArticleGetArticleBadRequest", []any{
		ErrCode_Article_ArticleNotFound,
	}, nil)
)

const (
	ErrCode_Article_ArticleNotFound = "ArticleNotFound"
)

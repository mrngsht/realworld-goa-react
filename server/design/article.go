package design

import . "goa.design/goa/v3/dsl"

var _ = Service("article", func() {
	Description("article")

	Method("get", func() {
		errorBadRequest := ErrorByErrorType(errType_ArticleGetArticleBadRequest)

		HTTP(func() {
			GET("article/{articleId}")
			Response(StatusOK)
			Response(errorBadRequest, StatusBadRequest)
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

	Method("list", func() {
		HTTP(func() {
			GET("articles")
			Response(StatusOK)
		})

		Payload(func() {
			AttributeWithName("tag", String)
			AttributeWithName("author", String)
			AttributeWithName("favorited", String)
			AttributeWithName("limit", Int32, func() {
				Default(20)
				Minimum(1)
				Maximum(1000)
			})
			AttributeWithName("offset", Int32, func() {
				Default(0)
				Minimum(0)
			})
		})

		Result(func() {
			Required(
				AttributeWithName("articles", ArrayOf(type_ArticleSummary)),
			)
		})
	})

	Method("feed", func() {
		HTTP(func() {
			GET("articles/feed")
			Response(StatusOK)
		})

		Payload(func() {
			AttributeWithName("limit", Int32, func() {
				Default(20)
				Minimum(1)
				Maximum(1000)
			})
			AttributeWithName("offset", Int32, func() {
				Default(0)
				Minimum(0)
			})
		})

		Result(func() {
			Required(
				AttributeWithName("articles", ArrayOf(type_ArticleSummary)),
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

	Method("update", func() {
		errorBadRequest := ErrorByErrorType(errType_ArticleUpdateArticleBadRequest)

		HTTP(func() {
			POST("article/{articleId}/update")
			Response(StatusOK)
			Response(errorBadRequest, StatusBadRequest)
		})

		Payload(func() {
			Required(
				AttributeWithName("articleId", String, def_Article_RequestArticleID),
			)
			AttributeWithName("title", String, def_Article_RequestTitle)
			AttributeWithName("description", String)
			AttributeWithName("body", String)
		})

		Result(func() {
			Required(
				AttributeWithName("article", type_ArticleDetail),
			)
		})
	})

	Method("delete", func() {
		errorBadRequest := ErrorByErrorType(errType_ArticleDeleteArticleBadRequest)

		HTTP(func() {
			POST("article/{articleId}/delete")
			Response(StatusOK)
			Response(errorBadRequest, StatusBadRequest)
		})

		Payload(func() {
			Required(
				AttributeWithName("articleId", String, def_Article_RequestArticleID),
			)
		})
	})

	Method("favorite", func() {
		errorBadRequest := ErrorByErrorType(errType_ArticleFavoriteArticleBadRequest)

		HTTP(func() {
			POST("article/{articleId}/favorite")
			Response(StatusOK)
			Response(errorBadRequest, StatusBadRequest)
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

	Method("unfavorite", func() {
		errorBadRequest := ErrorByErrorType(errType_ArticleUnfavoriteArticleBadRequest)

		HTTP(func() {
			POST("article/{articleId}/unfavorite")
			Response(StatusOK)
			Response(errorBadRequest, StatusBadRequest)
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

	Method("addComments", func() {
		errorBadRequest := ErrorByErrorType(errType_ArticleAddCommentsBadRequest)

		HTTP(func() {
			POST("article/{articleId}/addComments")
			Response(StatusOK)
			Response(errorBadRequest, StatusBadRequest)
		})

		Payload(func() {
			Required(
				AttributeWithName("articleId", String, def_Article_RequestArticleID),
				AttributeWithName("body", String),
			)
		})

		Result(func() {
			Required(
				AttributeWithName("comment", type_Comment),
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
	type_ArticleSummary = Type("ArticleSummary", func() {
		Required(
			AttributeWithName("articleId", String, func() {
				Format(FormatUUID)
			}),
			AttributeWithName("title", String),
			AttributeWithName("description", String),
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

	type_Comment = Type("Comment", func() {
		Required(
			AttributeWithName("id", String, func() {
				Format(FormatUUID)
			}),
			AttributeWithName("createdAt", String, func() {
				Format(FormatDateTime)
			}),
			AttributeWithName("updatedAt", String, func() {
				Format(FormatDateTime)
			}),
			AttributeWithName("body", String),
			AttributeWithName("author", type_Profile),
		)
	})
)

var (
	errType_ArticleGetArticleBadRequest = myErrorType("ArticleGetArticleBadRequest", []any{
		ErrCode_Article_ArticleNotFound,
	}, nil)
	errType_ArticleUpdateArticleBadRequest = myErrorType("ArticleUpdateArticleBadRequest", []any{
		ErrCode_Article_ArticleNotFound,
		ErrCode_Article_ForbiddenOperation,
	}, nil)
	errType_ArticleDeleteArticleBadRequest = myErrorType("ArticleDeleteArticleBadRequest", []any{
		ErrCode_Article_ArticleNotFound,
		ErrCode_Article_ForbiddenOperation,
	}, nil)
	errType_ArticleFavoriteArticleBadRequest = myErrorType("ArticleFavoriteArticleBadRequest", []any{
		ErrCode_Article_ArticleNotFound,
	}, nil)
	errType_ArticleUnfavoriteArticleBadRequest = myErrorType("ArticleUnfavoriteArticleBadRequest", []any{
		ErrCode_Article_ArticleNotFound,
	}, nil)
	errType_ArticleAddCommentsBadRequest = myErrorType("ArticleAddCommentsBadRequest", []any{
		ErrCode_Article_ArticleNotFound,
	}, nil)
)

const (
	ErrCode_Article_ArticleNotFound    = "ArticleNotFound"
	ErrCode_Article_ForbiddenOperation = "ForbiddenOperation"
)

package design

import (
	"fmt"
	"slices"

	. "goa.design/goa/v3/dsl"
	"goa.design/goa/v3/expr"
)

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

	CodeCommon_Unspecified = "Unspecified"
)

func myErrorType(name string, codes []any, ext func()) expr.UserType {
	if slices.Contains(codes, CodeCommon_Unspecified) {
		panic(fmt.Sprintf("%s is already included implicitly", CodeCommon_Unspecified))
	}

	codesWithDefaults := slices.Concat([]any{"Unspecified"}, codes)

	return Type(name, func() {
		Required(
			AttributeWithName("code", String, func() {
				Enum(codesWithDefaults...)
			}),
		)
		if ext != nil {
			ext()
		}
	})
}

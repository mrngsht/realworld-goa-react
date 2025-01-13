package design

import (
	. "goa.design/goa/v3/dsl"
	"goa.design/goa/v3/expr"
)

func AttributeWithName(name string, args ...any) string {
	Attribute(name, args...)
	return name
}

func ErrorByErrorType(et expr.UserType) string {
	n := et.Name()
	Error(n, et)
	return n
}

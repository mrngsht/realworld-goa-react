package design

import . "goa.design/goa/v3/dsl"

func AttributeWithName(name string, args ...any) string {
	Attribute(name, args...)
	return name
}

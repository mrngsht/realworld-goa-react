//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func GoaGen() error {
	return sh.RunV("goa", "gen", "github.com/mrngsht/realworld-goa-react/design")
}

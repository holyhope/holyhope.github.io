// +build mage

package main

import (
	"github.com/gohugoio/hugo/commands"
	"github.com/magefile/mage/mg"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
var Default = Build

// Build the website
func Build() error {
	mg.Deps(OptimizeImages)
	resp := commands.Execute([]string{})
	return resp.Err
}

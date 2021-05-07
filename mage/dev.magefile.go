// +build mage

package main

import (
	"github.com/gohugoio/hugo/commands"
	"github.com/magefile/mage/mg"
)

// Serve the website on localhost
func Dev() error {
	mg.Deps(OptimizeImages)
	resp := commands.Execute([]string{"server", "-D"})
	return resp.Err
}

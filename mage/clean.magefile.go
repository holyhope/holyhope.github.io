// +build mage

package main

import (
	"fmt"
	"os"
)

// Clean up generated files
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("public")
	os.RemoveAll(optimizedImagesDestPath)
}

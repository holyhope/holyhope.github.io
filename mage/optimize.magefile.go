// +build mage

package main

import (
	"errors"
	"fmt"
	"image"
    "image/jpeg"
	"os"
	"path"

	"github.com/nfnt/resize"
	"github.com/magefile/mage/mg"
)

// Optimize images to web
func Optimize() error {
	mg.Deps(OptimizeImages)

	return nil
}

func loadImage(path string) (image.Image, error) {
	f, err := os.Open(path)
    if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
    }

    defer f.Close()

	image, _, err := image.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("image decode: %w", err)
	}

	return image, nil
}

const (
	width  = 150
	height = 150
)

var optimizedImagesDestPath = fmt.Sprintf("static/%dx%d", width, height)

// Optimize images to web
func OptimizeImages() error {
	fmt.Printf("Optimizing images to %s...\n", optimizedImagesDestPath)

    image, err := loadImage("static/me.jpg")
    if err != nil {
		return err
    }

	if err := os.Mkdir(optimizedImagesDestPath, 0755); err != nil && !errors.Is(err, os.ErrExist) {
		return fmt.Errorf("mkdir: %w", err)
	}

	newImage := resize.Resize(width, height, image, resize.Lanczos3)

    f, err := os.OpenFile(path.Join(optimizedImagesDestPath, "me.jpg"), os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
		return fmt.Errorf("open file: %w", err)
    }

    defer f.Close()

	return jpeg.Encode(f, newImage, nil)
}

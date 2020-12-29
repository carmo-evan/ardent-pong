package pong

import (
	"image"
	"image/color"

	"github.com/split-cube-studios/ardent/engine"
	"golang.org/x/image/font"
)

type ImageLoader interface {
	NewImageFromPath(string) (engine.Image, error)
	NewImageFromAssetPath(string) (engine.Image, error)
	NewImageFromImage(image.Image) engine.Image

	NewTextImage(string, int, int, font.Face, color.Color) engine.Image
}

func generateRect(imageLoader ImageLoader, width int, height int, color color.Color) engine.Image {
	i := image.NewNRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			i.Set(x, y, color)
		}
	}

	return imageLoader.NewImageFromImage(i)
}

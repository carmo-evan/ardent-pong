package main

import (
	"log"

	"github.com/madxmike/ardent-pong/pong"
	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/engine"
)

var (
	screenWidth  = 800
	screenHeight = 800
)

func main() {

	pong := new(pong.Game)
	var camera engine.Camera
	game := ardent.NewGame(
		"Ardent - Pong",
		screenWidth,
		screenHeight,
		engine.FlagResizable,
		pong.Tick,
		func(ow, oh int) (int, int) {
			rAspect := float64(ow) / float64(oh)

			var w, h int
			if rAspect >= 1 {
				w = int(float64(screenHeight) * rAspect)
				h = screenHeight
			} else {
				w = screenWidth
				h = int(float64(screenWidth) / rAspect)
			}

			if camera != nil {
				camera.LookAt(float64(w/2), float64(h/2), 1)
			}
			return w, h
		},
	)

	renderer := game.NewRenderer()
	game.AddRenderer(renderer)
	camera = game.NewCamera()
	context := engine.NewContext(renderer, game.NewCollider())

	pong.Setup(game, context, camera)
	err := pong.Run()
	if err != nil {
		log.Fatal(err)
	}
}

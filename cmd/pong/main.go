package main

import (
	"log"

	"github.com/madxmike/ardent-pong/pong"
	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/engine"
)

func main() {
	context := engine.NewContext(nil, nil)

	game := ardent.NewGame(
		"Ardent - Pong",
		800,
		600,
		engine.FlagResizable,
		context.Tick,
		func(ow int, oh int) (int, int) {
			return 800, 600
		},
	)

	renderer := game.NewRenderer()
	game.AddRenderer(renderer)
	context.Renderer = renderer
	pong := pong.Pong{
		Game:  game,
		Board: pong.NewBoard(game, context),
	}
	err := pong.Run()
	if err != nil {
		log.Fatal(err)
	}
}

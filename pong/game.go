package pong

import (
	"github.com/split-cube-studios/ardent/engine"
)

type Game struct {
	engine.Game
	context *engine.Context
	camera  engine.Camera
	scene   Scene

	sceneInitialized bool
}

func (g *Game) Setup(game engine.Game, ctx *engine.Context, camera engine.Camera) {
	g.Game = game
	g.context = ctx
	g.camera = camera

	g.scene = &GameScene{
		leftPaddle: Paddle{
			controller: g,
			speed:      10,
		},
		rightPaddle: Paddle{
			controller: g,
			speed:      10,
		},
	}
}

func (g *Game) Tick() {
	if !g.sceneInitialized {
		g.scene.Setup(g.context)
		g.scene.LayoutUI(g, g.camera)
		g.scene.LayoutEntities(g, g.camera)
		g.sceneInitialized = true
	}

	g.context.Tick()
	g.scene.Tick(g, g.camera)
}

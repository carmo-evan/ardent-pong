package pong

import (
	"image/color"

	"github.com/split-cube-studios/ardent/engine"
)

// A scene represents a set of ui and entity elements such as menus, game boards, and more
// The lifecycle of a scene is as follows:
// Setup
// LayoutUI
// LayoutEntities
// Tick

type Scene interface {
	// Create all entities and add them to the context
	Setup(*engine.Context)

	// Layout all ui elements in the scene and load textures
	LayoutUI(engine.Component, engine.Camera)

	// Layout all entities in the scene and load textures
	LayoutEntities(engine.Component, engine.Camera)

	// Do any per-tick updates
	// TODO - its kinda jank to do this
	Tick(engine.Component, engine.Camera)
}

type GameScene struct {
	ball        Ball
	leftPaddle  Paddle
	rightPaddle Paddle

	leftScore  int
	rightScore int
}

func (s *GameScene) Setup(ctx *engine.Context) {
	ctx.AddEntity(&s.ball)
	ctx.AddEntity(&s.leftPaddle)
	ctx.AddEntity(&s.rightPaddle)
}

func (s *GameScene) LayoutUI(component engine.Component, camera engine.Camera) {

}

func (s *GameScene) LayoutEntities(component engine.Component, camera engine.Camera) {
	cameraX, cameraY := camera.Position()
	s.ball.Vec2 = engine.Vec2{
		X: cameraX,
		Y: cameraY,
	}

	paddleDistance := 0.2
	s.leftPaddle.Vec2 = engine.Vec2{
		X: cameraX - (cameraX * (1 - paddleDistance)),
		Y: cameraY,
	}

	s.rightPaddle.Vec2 = engine.Vec2{
		X: cameraX + (cameraX * (1 - paddleDistance)),
		Y: cameraY,
	}

	s.ball.AddImage(generateRect(component, 20, 20, color.White))
	s.leftPaddle.AddImage(generateRect(component, 20, 200, color.White))
	s.rightPaddle.AddImage(generateRect(component, 20, 200, color.White))

}

func (s *GameScene) Tick(component engine.Component, camera engine.Camera) {
	// do scoring logic
}

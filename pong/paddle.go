package pong

import (
	"github.com/split-cube-studios/ardent/engine"
)

type PaddleController interface {
	engine.Input
}

type Paddle struct {
	engine.CoreEntity
	controller PaddleController
	speed      int
}

func (p *Paddle) Tick() {
	p.CoreEntity.Tick()

	if p.controller.IsKeyPressed(engine.KeyUp) {
		p.Y += float64(p.speed)
	} else if p.controller.IsKeyPressed(engine.KeyDown) {
		p.Y -= float64(p.speed)
	}
}

func (b *Paddle) Class() string {
	return "Paddle"
}

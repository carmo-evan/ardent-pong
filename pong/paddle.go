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
}

func (p *Paddle) Tick() {
	// TODO - Handle movement
}

func (b *Paddle) Class() string {
	return "Paddle"
}

package pong

import (
	"fmt"

	"github.com/split-cube-studios/ardent/engine"
)

type Ball struct {
	engine.CoreEntity
	angle float64
}

func (b *Ball) Tick() {

	fmt.Println(b.Vec2)
	b.CoreEntity.Tick()
}

func (b *Ball) Class() string {
	return "Ball"
}

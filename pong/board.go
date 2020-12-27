package pong

import (
	"image/color"

	"github.com/split-cube-studios/ardent/engine"
	"golang.org/x/image/font/basicfont"
)

type Board struct {
	engine.CoreEntity

	ball        Ball
	leftPaddle  Paddle
	rightPaddle Paddle

	leftScore  int
	rightScore int
}

func NewBoard(cmp engine.Component, ctx *engine.Context) *Board {
	b := new(Board)
	b.ball.CoreEntity.X = 400
	b.ball.CoreEntity.Y = 300
	b.ball.angle = 3
	ctx.AddEntity(b)
	ctx.AddEntity(&b.ball)
	ctx.AddEntity(&b.leftPaddle)
	ctx.AddEntity(&b.rightPaddle)

	b.ball.AddImage(cmp.NewTextImage("ball", 50, 50, basicfont.Face7x13, color.White))
	return b
}

func (b *Board) SetLeftPaddleController(pc PaddleController) {
	b.leftPaddle.controller = pc
}

func (b *Board) SetRightPaddleController(pc PaddleController) {
	b.rightPaddle.controller = pc
}

func (b *Board) Class() string {
	return "Board"
}

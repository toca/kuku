package views

import (
	"../models"
)

type Frame struct {
	model    *models.Frame
	renderer *Renderer
}

func NewFrame(renderer *Renderer, model *models.Frame) *Frame {
	return &Frame{model, renderer}
}

// 何かのイベントでコールバックがいいか木がする
func (this *Frame) Load() {
	if (*this.model).Width <= 0 || (*this.model).Height <= 0 {
		return
	}
	// corner
	this.renderer.Set(0, 0, '┌')
	this.renderer.Set(this.model.Width-1, 0, '┐')
	this.renderer.Set(0, this.model.Height-1, '└')
	this.renderer.Set(this.model.Width-1, this.model.Height-1, '┘')
	// north & south
	for i := 1; i < this.model.Width-1; i++ {
		this.renderer.Set(i, 0, '─')
		this.renderer.Set(i, this.model.Height-1, '─')
	}
	// west & east
	for i := 1; i < this.model.Height-1; i++ {
		this.renderer.Set(0, i, '│')
		this.renderer.Set(this.model.Width-1, i, '│')
	}

}

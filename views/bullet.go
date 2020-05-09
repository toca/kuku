package views

import (
	"../models"
)

type Bullet struct {
	model    *models.Bullet
	renderer *Renderer
}

func NewBullet(renderer *Renderer, model *models.Bullet) *Bullet {
	return &Bullet{model, renderer}
}

func (this *Bullet) Load() {
	this.renderer.Set(this.model.Rect().Min.X, this.model.Rect().Min.Y, '*')
}

package views

import (
	"../models"
)

type Saucer struct {
	model    *models.Saucer
	renderer *Renderer
}

func NewSaucer(renderer *Renderer, model *models.Saucer) *Saucer {
	return &Saucer{model, renderer}
}

func (this *Saucer) Load() {
	for x := this.model.Rect().Min.X; x < this.model.Rect().Max.X; x++ {
		this.renderer.Set(x, this.model.Rect().Min.Y, '▬')
	}
}

package views

import (
	"fmt"

	"kuku/models"
)

type Saucer struct {
	model    *models.Saucer
	renderer *Renderer
}

func NewSaucer(renderer *Renderer, model *models.Saucer) *Saucer {
	return &Saucer{model, renderer}
}

func (this *Saucer) Load() {
	if this.model.Rect().Min.X <= 0 {
		panic(fmt.Sprintf("saucer index out. %v", this.model.Rect()))
	}
	for x := this.model.Rect().Min.X; x < this.model.Rect().Max.X; x++ {
		this.renderer.Set(x, this.model.Rect().Min.Y, '▬')
	}
}

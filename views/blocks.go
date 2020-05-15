package views

import (
	"kuku/models"
)

type Blocks struct {
	renderer *Renderer
	model    *models.Blocks
}

func NewBlocks(r *Renderer, m *models.Blocks) *Blocks {
	return &Blocks{r, m}
}

func (this *Blocks) Load() {
	for _, block := range this.model.List() {
		for x := block.Rect().Min.X; x < block.Rect().Max.X; x++ {
			this.renderer.Set(x, block.Rect().Min.Y, '▮')
		}
	}
}

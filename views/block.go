package views

import (
	"kuku/models"
)

type Block struct {
	renderer *Renderer
	model    *models.Block
}

func NewBlock(r *Renderer, m *models.Block) *Block {
	return &Block{r, m}
}

func (this *Block) Load() {
	// this.renderer.Set(this.model.Rect().Min.X-1, this.model.Rect().Min.Y, '\033[33m')
	for x := this.model.Rect().Min.X; x < this.model.Rect().Max.X; x++ {
		this.renderer.Set(x, this.model.Rect().Min.Y, '▮')
	}
	// this.renderer.Set(this.model.Rect().Max.X, this.model.Rect().Max.Y, '\033[0m')
}

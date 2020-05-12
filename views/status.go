package views

import (
	"fmt"
	"kuku/models"
	"strings"
)

type Status struct {
	renderer *Renderer
	model    *models.Status
}

func NewStatus(r *Renderer, model *models.Status) *Status {
	return &Status{r, model}
}

func (this *Status) Load() {
	bullets := strings.Repeat("* ", this.model.BulletCount)
	s := fmt.Sprintf("| %s| %s", bullets, this.model.Message)
	for i, r := range s {
		this.renderer.Set(this.model.Rect().Min.X+i, this.model.Rect().Min.Y, r)
	}
}

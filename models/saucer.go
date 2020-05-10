package models

import (
	"image"
)

type Saucer struct {
	rect image.Rectangle
}

func NewSaucer(x0, y0, x1, y1 int) *Saucer {
	r := image.Rect(x0, y0, x1, y1)
	return &Saucer{r}
}

func (this *Saucer) Left() {
	this.rect.Min.X--
	this.rect.Max.X--
}

func (this *Saucer) Right() {
	this.rect.Min.X++
	this.rect.Max.X++
}

// implememt models.Object
func (this *Saucer) Rect() image.Rectangle {
	return this.rect
}
func (this *Saucer) HitTest(o Object) bool {
	return this.rect.Overlaps(o.Rect())
}
func (this *Saucer) Affect(o Object) {

}

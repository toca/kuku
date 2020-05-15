package models

import (
	"image"
)

type Saucer struct {
	rect *image.Rectangle
}

func NewSaucer(x0, y0, x1, y1 int) *Saucer {
	rect := image.Rect(x0, y0, x1, y1)
	return &Saucer{&rect}
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
	return *this.rect
}

func (this *Saucer) HitTest(o Object) bool {
	reflector := NewReflector(this.rect.Min.X, this.rect.Min.Y, this.rect.Max.X, this.rect.Max.Y)
	return reflector.HitTest(o)
}

func (this *Saucer) Affect(o Object) {
	// fmt.Printf("saucer Affect\n") 何とぶつかっているのだ?
	if bullet, ok := o.(*Bullet); ok {
		reflector := NewReflectorWithAcceleration(this.rect.Min.X, this.rect.Min.Y, this.rect.Max.X, this.rect.Max.Y, 1)
		reflector.Affect(bullet)
	}
}

func (this *Saucer) MarkedForDeath() bool {
	return false
}

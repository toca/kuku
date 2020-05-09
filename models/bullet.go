package models

import (
	"fmt"
	"image"
)

type Bullet struct {
	rectangle image.Rectangle
	vector    image.Point
}

func NewBullet(x, y int, vx, vy int) *Bullet {
	return &Bullet{
		image.Rect(x, y, x+1, y+1),
		image.Pt(vx, vy)}
}

// object interface
func (this *Bullet) Rect() image.Rectangle {
	return this.rectangle
}

func (this *Bullet) Vect() *image.Point {
	return &this.vector
}
func (this *Bullet) Action() {
	this.rectangle.Min.X += this.Vect().X
	this.rectangle.Min.Y += this.Vect().Y
	this.rectangle.Max.X += this.Vect().X
	this.rectangle.Max.Y += this.Vect().Y
}

// object interface
func (this *Bullet) HitTest(o Object) bool {
	return this.Rect().Overlaps(o.Rect())
}

// object interface
func (this *Bullet) Affect(o Object) {
	fmt.Println(this.Rect())
}

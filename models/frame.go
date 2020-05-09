package models

import (
	"image"
)

type Frame struct {
	Width    int
	Height   int
	elements map[string]image.Rectangle
}

func NewFrame(width int, height int) *Frame {
	elements := make(map[string]image.Rectangle)
	elements["north"] = image.Rect(1, 0, width-2, 2)
	elements["south"] = image.Rect(1, height-1, width-2, height-2)
	elements["west"] = image.Rect(0, 1, 2, height-2)
	elements["east"] = image.Rect(width-2, 1, width-1, height-2)
	elements["upper_left"] = image.Rect(0, 0, 1, 1)
	elements["upper_right"] = image.Rect(width-1, 0, width, 1)
	elements["lower_left"] = image.Rect(0, height-1, 1, height)
	elements["lower_right"] = image.Rect(width-1, height-1, width, height)
	return &Frame{width, height, elements}
}

// Object interface
func (this *Frame) HitTest(o Object) bool {
	for _, v := range this.elements {
		if v.Overlaps(o.Rect()) {
			// fmt.Printf("hit %s\n", k)
			return true
		}
	}
	return false
}

func (this *Frame) Rect() image.Rectangle {
	return image.Rect(0, 0, this.Width-1, this.Height-1)
}

func (this *Frame) Affect(o Object) {
	if bullet, ok := o.(*Bullet); ok {
		if this.elements["north"].Overlaps(bullet.Rect()) {
			bullet.Vect().Y *= -1
		} else if this.elements["south"].Overlaps(bullet.Rect()) {
			bullet.Vect().Y *= -1
		} else if this.elements["west"].Overlaps(bullet.Rect()) {
			bullet.Vect().X *= -1
		} else if this.elements["east"].Overlaps(bullet.Rect()) {
			bullet.Vect().X *= -1
		} else {
			bullet.Vect().X *= -1
			bullet.Vect().Y *= -1
		}
	}
}

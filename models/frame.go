package models

import (
	"image"
	"math"
)

type Frame struct {
	Width    int
	Height   int
	elements map[string]image.Rectangle
}

func NewFrame(width int, height int) *Frame {
	elements := make(map[string]image.Rectangle)
	elements["north"] = image.Rect(1, 0, width-2, 0)
	elements["south"] = image.Rect(1, height-1, width-2, height-1)
	elements["west"] = image.Rect(int(math.Inf(-1)), 1, 0, height-2)
	elements["east"] = image.Rect(width-1, 1, math.MaxInt32, height-2)
	elements["upper_left"] = image.Rect(0, 0, 0, 0)
	elements["upper_right"] = image.Rect(width-1, 0, width, 0)
	elements["lower_left"] = image.Rect(0, height-1, 0, height-1)
	elements["lower_right"] = image.Rect(width-1, height-1, width-1, height-1)
	return &Frame{width, height, elements}
}

// Object interface
func (this *Frame) HitTest(o Object) bool {
	for k, v := range this.elements {
		_ = k
		if hit(v, o.Rect()) {
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

		if hit(this.elements["north"], bullet.Rect()) {
			bullet.Vect().Y *= -1
		} else if hit(this.elements["south"], bullet.Rect()) {
			bullet.Vect().Y *= -1
		} else if hit(this.elements["west"], bullet.Rect()) {
			bullet.Vect().X *= -1
		} else if hit(this.elements["east"], bullet.Rect()) {
			bullet.Vect().X *= -1
		} else {
			bullet.Vect().X *= -1
			bullet.Vect().Y *= -1
		}
	} else if saucer, ok := o.(*Saucer); ok {
		for hit(this.elements["west"], saucer.Rect()) {
			saucer.Right()
		}
		for hit(this.elements["east"], saucer.Rect()) {
			saucer.Left()
		}
	}
}

func hit(lhs, rhs image.Rectangle) bool {
	return Overlap(lhs, rhs)
	// if (lhs.Min.X < min(rhs.Min.X, rhs.Max.X) || max(rhs.Min.X, rhs.Max.X) < lhs.Min.X) &&
	// 	(lhs.Max.X < min(rhs.Min.X, rhs.Max.X) || max(rhs.Min.X, rhs.Max.X) < lhs.Max.X) {
	// 	return false
	// }

	// if (lhs.Min.Y < min(rhs.Min.Y, rhs.Max.Y) || max(rhs.Min.Y, rhs.Max.Y) < lhs.Min.Y) &&
	// 	(lhs.Max.Y < min(rhs.Min.Y, rhs.Max.Y) || max(rhs.Min.Y, rhs.Max.Y) < lhs.Max.Y) {
	// 	return false
	// }
	// return true

	// if (min(rhs.Min.X, rhs.Max.X) <= lhs.Min.X && lhs.Min.X <= max(rhs.Min.X, rhs.Max.X) ||
	// 	min(rhs.Min.X, rhs.Max.X) <= lhs.Max.X && lhs.Max.X <= max(rhs.Min.X, rhs.Max.X)) &&
	// 	(min(rhs.Min.Y, rhs.Max.Y) <= lhs.Min.Y && lhs.Min.Y <= max(rhs.Min.Y, rhs.Max.Y) ||
	// 		min(rhs.Min.Y, rhs.Max.Y) <= lhs.Max.Y && lhs.Max.Y <= max(rhs.Min.Y, rhs.Max.Y)) {
	// 	return true
	// } else {
	// 	return false
	// }
}

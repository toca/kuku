package models

import (
	"image"
	"math"
)

type Frame struct {
	Width    int
	Height   int
	elements map[string]*image.Rectangle
}

func NewFrame(width int, height int) *Frame {
	elements := make(map[string]*image.Rectangle)
	n := image.Rect(1, 0, width-2, 0)
	elements["north"] = &n
	s := image.Rect(1, height-1, width-2, height-1)
	elements["south"] = &s
	w := image.Rect(int(math.Inf(-1)), 1, 0, height-2)
	elements["west"] = &w
	e := image.Rect(width-1, 1, math.MaxInt32, height-2)
	elements["east"] = &e
	ul := image.Rect(0, 0, 0, 0)
	elements["upper_left"] = &ul
	ur := image.Rect(width-1, 0, width, 0)
	elements["upper_right"] = &ur
	ll := image.Rect(0, height-1, 0, height-1)
	elements["lower_left"] = &ll
	lr := image.Rect(width-1, height-1, width-1, height-1)
	elements["lower_right"] = &lr
	return &Frame{width, height, elements}
}

// Object interface
func (this *Frame) HitTest(o Object) bool {
	for k, v := range this.elements {
		_ = k
		r := o.Rect()
		if hit(v, &r) {
			// fmt.Printf("frameModel:Hit %s\n", k)
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
		currentVect := bullet.Vect()
		switch this.GetHitElement(&o) {

		case "north":
			if currentVect.Y < 0 {
				currentVect.Y *= -1
				bullet.SetVect(&currentVect)
			}
		case "south":
			if 0 < currentVect.Y {
				currentVect.Y *= -1
				bullet.SetVect(&currentVect)
			}
		case "west":
			if currentVect.X < 0 {
				currentVect.X *= -1
				bullet.SetVect(&currentVect)
			}
		case "east":
			if 0 < currentVect.X {
				currentVect.X *= -1
				bullet.SetVect(&currentVect)
			}
		case "upper_left":
			if currentVect.X < 0 && currentVect.Y < 0 {
				currentVect.X *= -1
				currentVect.Y *= -1
				bullet.SetVect(&currentVect)
			}
		case "upper_right":
			if 0 < currentVect.X && currentVect.Y < 0 {
				currentVect.X *= -1
				currentVect.Y *= -1
				bullet.SetVect(&currentVect)
			}
		case "lower_left":
			if currentVect.X < 0 && 0 < currentVect.Y {
				currentVect.X *= -1
				currentVect.Y *= -1
				bullet.SetVect(&currentVect)
			}
		case "lower_right":
			if 0 < currentVect.X && 0 < currentVect.Y {
				currentVect.X *= -1
				currentVect.Y *= -1
				bullet.SetVect(&currentVect)
			}
		default:
			panic("models.frame unknown element name")
		}
	} else if saucer, ok := o.(*Saucer); ok {
		r := saucer.Rect()
		for hit(this.elements["west"], &r) {
			saucer.Right()
		}
		for hit(this.elements["east"], &r) {
			saucer.Left()
		}
	}
}

func (this *Frame) GetHitElement(o *Object) string {
	rect := (*o).Rect()
	for k, v := range this.elements {
		if hit(v, &rect) {
			return k
		}
	}
	return "null"
}
func hit(lhs, rhs *image.Rectangle) bool {
	return Overlap(lhs, rhs)
}

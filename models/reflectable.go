package models

import "image"

type Reflectable interface {
	Vect() image.Point
	SetVect(v *image.Point)
	Rect() image.Rectangle
}

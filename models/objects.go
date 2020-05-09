package models

import (
	"image"
)

type Object interface {
	Rect() image.Rectangle
	HitTest(Object) bool
	Affect(Object)
}

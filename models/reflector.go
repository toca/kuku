package models

import (
	"fmt"
	"image"
)

// // factory
// type ReflectorFactory struct {
// 	rect   *image.Rectangle
// 	top    *image.Rectangle
// 	bottom *image.Rectangle
// 	left   *image.Rectangle
// 	right  *image.Rectangle
// }

// func NewReflectorFactory() *ReflectorFactory {
// 	return &ReflectorFactory{}
// }

// func (this *ReflectorFactory) North(x0, y0, x1, y1 int) *ReflectorFactory {
// 	this.top = image.Rect(x0, y0, x1, y1)
// 	return this
// }

// func (this *ReflectorFactory) South(x0, y0, x1, y1 int) *ReflectorFactory {
// 	this.elements["south"] = image.Rect(x0, y0, x1, y1)
// 	return this
// }

// func (this *ReflectorFactory) West(x0, y0, x1, y1 int) *ReflectorFactory {
// 	this.elements["west"] = image.Rect(x0, y0, x1, y1)
// 	return this
// }

// func (this *ReflectorFactory) East(x0, y0, x1, y1 int) *ReflectorFactory {
// 	this.elements["east"] = image.Rect(x0, y0, x1, y1)
// 	return this
// }

// func (this *ReflectorFactory) Northwest(x0, y0, x1, y1 int) *ReflectorFactory {
// 	this.elements["northwest"] = image.Rect(x0, y0, x1, y1)
// 	return this
// }

// func (this *ReflectorFactory) Northeast(x0, y0, x1, y1 int) *ReflectorFactory {
// 	this.elements["northeast"] = image.Rect(x0, y0, x1, y1)
// 	return this
// }

// func (this *ReflectorFactory) Southwest(x0, y0, x1, y1 int) *ReflectorFactory {
// 	this.elements["southwest"] = image.Rect(x0, y0, x1, y1)
// 	return this
// }

// func (this *ReflectorFactory) Southeast(x0, y0, x1, y1 int) *ReflectorFactory {
// 	this.elements["southeast"] = image.Rect(x0, y0, x1, y1)
// 	return this
// }

// func (this *ReflectorFactory) Create() *Reflector {
// 	return &Reflector{&this.elements}
// }

// recrector
type Reflector struct {
	rect        *image.Rectangle
	top         *image.Rectangle
	bottom      *image.Rectangle
	left        *image.Rectangle
	right       *image.Rectangle
	topLeft     *image.Rectangle
	topRight    *image.Rectangle
	bottomLeft  *image.Rectangle
	bottomRight *image.Rectangle
}

func NewReflector(x0, y0, x1, y1 int) *Reflector {
	var (
		rect        = image.Rect(x0-1, y0-1, x1+1, y1+1)
		top         = image.Rect(x0, y0-1, x1, y0-1)
		bottom      = image.Rect(x0, y1+1, x1, y1+1)
		left        = image.Rect(x0-1, y0, x0-1, y0)
		right       = image.Rect(x1+1, y0, x1+1, y0)
		topLeft     = image.Rect(x0-1, y0-1, x0-1, y0-1)
		topRight    = image.Rect(x1+1, y0-1, x1+1, y0-1)
		bottomLeft  = image.Rect(x0-1, y1+1, x0-1, y1+1)
		bottomRight = image.Rect(x1+1, y1+1, x1+1, y1+1)
	)
	return &Reflector{&rect, &top, &bottom, &left, &right, &topLeft, &topRight, &bottomLeft, &bottomRight}
}

func (this *Reflector) HitTest(o Object) bool {
	r := o.Rect()
	return Overlap(this.rect, &r)
}

func (this *Reflector) Affect(reflectable Reflectable) {
	currentVect := reflectable.Vect()
	r := reflectable.Rect()
	fmt.Printf("Reflector.Hit: %s\n", this.GetHitPos(&r, &currentVect))
	switch this.GetHitPos(&r, &currentVect) {

	case "top":
		currentVect.Y *= -1
		reflectable.SetVect(&currentVect)
	case "bottom":
		currentVect.Y *= -1
		reflectable.SetVect(&currentVect)
	case "left":
		currentVect.X *= -1
		reflectable.SetVect(&currentVect)
	case "right":
		currentVect.X *= -1
		reflectable.SetVect(&currentVect)
	case "top_left":
		currentVect.X *= -1
		currentVect.Y *= -1
		reflectable.SetVect(&currentVect)
	case "top_right":
		currentVect.X *= -1
		currentVect.Y *= -1
		reflectable.SetVect(&currentVect)
	case "bottom_left":
		currentVect.X *= -1
		currentVect.Y *= -1
		reflectable.SetVect(&currentVect)
	case "bottom_right":
		currentVect.X *= -1
		currentVect.Y *= -1
		reflectable.SetVect(&currentVect)
	default:
		panic("models.frame unknown element name")
	}
}

func (this *Reflector) GetHitPos(r *image.Rectangle, v *image.Point) string {
	if Overlap(this.top, r) {
		return "top"
	} else if Overlap(this.bottom, r) {
		return "bottom"
	} else if Overlap(this.left, r) {
		return "left"
	} else if Overlap(this.right, r) {
		return "right"
	} else if Overlap(this.topLeft, r) {
		return "top_left"
	} else if Overlap(this.topRight, r) {
		return "top_right"
	} else if Overlap(this.bottomLeft, r) {
		return "bottom_left"
	} else if Overlap(this.bottomRight, r) {
		return "bottom_right"
	} else {
		return "day_after_tommorow"
	}
}

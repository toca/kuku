package models

import (
	"image"
)

type Status struct {
	Message     string
	BulletCount int
	FrameRate   int
	pos         image.Point
}

var instance = Status{"", 3, 0, image.Pt(0, 0)}

func GetStatus() *Status {
	return &instance
}
func (this *Status) SetPos(x, y int) {
	this.pos = image.Pt(x, y)
}
func (this *Status) Pos() image.Point {
	return this.pos
}

// object impl
func (this *Status) Rect() image.Rectangle {
	return image.Rectangle{this.pos, this.pos}
}
func (this *Status) HitTest(Object) bool {
	return false
}
func (this *Status) Affect(Object) {}

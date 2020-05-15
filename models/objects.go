package models

import (
	"image"
)

type Object interface {
	Rect() image.Rectangle
	HitTest(Object) bool
	Affect(Object)
	MarkedForDeath() bool
}

// to collision detection
func Overlap(lhs, rhs *image.Rectangle) bool {
	// x
	if !lineOverlap(lhs.Min.X, lhs.Max.X, rhs.Min.X, rhs.Max.X) {
		return false
	}
	// y
	if !lineOverlap(lhs.Min.Y, lhs.Max.Y, rhs.Min.Y, rhs.Max.Y) {
		return false
	}
	return true
}

func lineOverlap(a0, a1, b0, b1 int) bool {
	if min(a0, a1) <= b0 && b0 <= max(a0, a1) {
		return true
	}
	if min(a0, a1) <= b1 && b1 <= max(a0, a1) {
		return true
	}
	return false
}

func min(lhs, rhs int) int {
	if lhs < rhs {
		return lhs
	} else {
		return rhs
	}
}
func max(lhs, rhs int) int {
	if lhs < rhs {
		return rhs
	} else {
		return lhs
	}
}

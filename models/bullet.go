package models

import (
	"image"
	"math"
	"time"
)

type Bullet struct {
	rect      *image.Rectangle
	vector    *image.Point
	lastMoved time.Time
	lastHited time.Time
	changing  bool
}

// 移動量は 1/secにしてみる
func NewBullet(x, y int, vx, vy int) *Bullet {
	v := image.Pt(vx, vy)
	r := image.Rect(x, y, x, y)
	now := time.Now()
	return &Bullet{&r, &v, now, now, false} // image.Rect(x, y, x, y),

}

// object interface
func (this *Bullet) Rect() image.Rectangle {
	return *this.rect
}

func (this *Bullet) Vect() image.Point {
	return *this.vector
}
func (this *Bullet) SetVect(v *image.Point) {
	if !this.changing {
		this.vector.X = v.X
		this.vector.Y = v.Y
		this.changing = true
	} else {
		// fmt.Printf("bullet.setvect ignore\n")
	}
}

func (this *Bullet) Action() {
	d := time.Now().Sub(this.lastMoved)
	xMoves := float64(d.Milliseconds()) * float64(this.Vect().X) / 1000.0
	if math.Abs(xMoves) < 1.0 {
		return
	}
	yMoves := float64(d.Milliseconds()) * float64(this.Vect().Y) / 1000.0
	if math.Abs(yMoves) < 1.0 {
		return
	}
	this.rect.Min.X += int(xMoves)
	this.rect.Max.X += int(xMoves)
	this.rect.Min.Y += int(yMoves)
	this.rect.Max.Y += int(yMoves)

	this.lastMoved = this.lastMoved.Add(d)
	this.changing = false
}

// object interface
func (this *Bullet) HitTest(o Object) bool {
	return this.Rect().Overlaps(o.Rect())
}

// object interface
func (this *Bullet) Affect(o Object) {
	if block, ok := o.(*Block); ok {
		if this.lastHited != this.lastMoved {
			block.Hit()
			this.lastHited = this.lastMoved
		}
	}
}

func (this *Bullet) MarkedForDeath() bool {
	return false
}

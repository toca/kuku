package models

import (
	"image"
)

type BlockType int

const (
	NORMAL_BLOCK BlockType = 1
	HARD_BLOCK   BlockType = 2
)

type Block struct {
	rect       *image.Rectangle
	durability int
	blockType  BlockType
	reflector  *Reflector
}

func NewBlock(x0, y0, x1, y1 int, blockType BlockType) *Block {
	reflector := NewReflector(x0, y0, x1, y1)
	d := durability(blockType)
	r := image.Rect(x0, y0, x1, y1)
	return &Block{&r, d, blockType, reflector}
}

func (this *Block) Hit() {
	this.durability--
	if this.durability == 0 {
		GetStatus().Score += uint32(100 * this.blockType)
	}
}

func durability(t BlockType) int {
	switch t {
	case NORMAL_BLOCK:
		return 3
	case HARD_BLOCK:
		return 6
	default:
		panic("models.Block: Unknown BlockType")
	}
}

// implememt models.Object
func (this *Block) Rect() image.Rectangle {
	return *this.rect
}
func (this *Block) HitTest(o Object) bool {
	return this.reflector.HitTest(o)
}
func (this *Block) Affect(o Object) {
	// frame と共通化
	if bullet, ok := o.(*Bullet); ok {
		this.reflector.Affect(bullet)
	}
}
func (this *Block) MarkedForDeath() bool {
	return this.durability <= 0
}

package controllers

import (
	"fmt"

	"kuku/models"
	"kuku/operation"
)

type Saucer struct {
	model *models.Saucer
}

func NewSaucer(m *models.Saucer) *Saucer {
	return &Saucer{m}
}

func (this *Saucer) Input(keyInput operation.KeyInput) {
	fmt.Printf("RepeatCount:%d\n", keyInput.Repeat)
	fmt.Println(this.model.Rect())
	for i := 0; i < keyInput.Repeat; i++ {
		switch keyInput.Key {
		case operation.VK_LEFT:
			this.model.Left()
		case operation.VK_RIGHT:
			this.model.Right()
		default:
			fmt.Printf("ignored:%v\n", keyInput.Key)
		}
	}
	fmt.Println(this.model.Rect())
}

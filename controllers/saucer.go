package controllers

import (
	"fmt"

	"../input"
	"../models"
)

type Saucer struct {
	model *models.Saucer
}

func NewSaucer(m *models.Saucer) *Saucer {
	return &Saucer{m}
}

func (this *Saucer) Input(ipt input.Input) {
	fmt.Printf("RepeatCount:%d\n", ipt.Repeat)
	// for i := 0; i <= ipt.Repeat; i++ {
	switch ipt.Key {
	case input.VK_LEFT:
		this.model.Left()
	case input.VK_RIGHT:
		this.model.Right()
	default:
	}
	// }
}

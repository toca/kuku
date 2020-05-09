package views

import (
	"../console"
)

type Renderer struct {
	console *console.Console
	width   int
	height  int
	buffer  [][]rune
}

func NewRenderer(width int, height int) (*Renderer, error) {
	c, err := console.NewConsole(width, height)
	if err != nil {
		return nil, err
	}
	buffer := make([][]rune, height)
	for i, _ := range buffer {
		buffer[i] = make([]rune, width)
	}
	renderer := Renderer{c, width, height, buffer}
	renderer.Clear()
	return &renderer, nil
}

func (this *Renderer) Set(x int, y int, character rune) {
	this.buffer[y][x] = character
}

func (this *Renderer) Render() {
	lines := ""
	for row := range this.buffer {
		lines += string(this.buffer[row]) + "\n"
	}
	this.console.Write(lines)
}

func (this *Renderer) Clear() {
	for row, _ := range this.buffer {
		for col, _ := range this.buffer[row] {
			this.buffer[row][col] = ' '
		}
	}
}

func (this *Renderer) Close() {
	this.console.Close()
}

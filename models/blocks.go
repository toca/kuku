package models

import "fmt"

type Blocks struct {
	list []*Block
}

// TODO add good arg
func NewBlocks() *Blocks {
	list := make([]*Block, 0)
	list = append(list, NewBlock(9, 6, 24, 6, NORMAL_BLOCK))
	list = append(list, NewBlock(29, 6, 44, 6, NORMAL_BLOCK))
	list = append(list, NewBlock(54, 6, 69, 6, NORMAL_BLOCK))
	list = append(list, NewBlock(75, 6, 89, 6, NORMAL_BLOCK))

	list = append(list, NewBlock(9, 12, 24, 12, NORMAL_BLOCK))
	list = append(list, NewBlock(29, 12, 44, 12, NORMAL_BLOCK))
	list = append(list, NewBlock(54, 12, 69, 12, NORMAL_BLOCK))
	list = append(list, NewBlock(75, 12, 89, 12, NORMAL_BLOCK))

	list = append(list, NewBlock(9, 20, 24, 20, NORMAL_BLOCK))
	list = append(list, NewBlock(29, 20, 44, 20, NORMAL_BLOCK))
	list = append(list, NewBlock(54, 20, 69, 20, NORMAL_BLOCK))
	list = append(list, NewBlock(75, 20, 89, 20, NORMAL_BLOCK))

	list = append(list, NewBlock(9, 26, 24, 26, NORMAL_BLOCK))
	list = append(list, NewBlock(29, 26, 44, 26, NORMAL_BLOCK))
	list = append(list, NewBlock(54, 26, 69, 26, NORMAL_BLOCK))
	list = append(list, NewBlock(75, 26, 89, 26, NORMAL_BLOCK))

	fmt.Println(*list[0])
	return &Blocks{list}
}

func (this *Blocks) List() []*Block {
	for i := 0; i < len(this.list); i++ {
		if this.list[i].MarkedForDeath() {
			this.list = append(this.list[:i], this.list[i+1:]...)
			i++
		}
	}
	return this.list
}

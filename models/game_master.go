package models

type GameMaster struct {
	blocks   []*Block
	finished bool
}

func NewGameMaster(blocks []*Block) *GameMaster {
	return &GameMaster{blocks, false}
}
func (this *GameMaster) Action() {
	for i := 0; i < len(this.blocks); i++ {
		if this.blocks[i].MarkedForDeath() {
			this.blocks = append(this.blocks[:i], this.blocks[i+1:]...)
			i++
		}
	}
	if 0 == len(this.blocks) {
		this.finished = true
		GetStatus().Message = "GAME CLEAR"
	}
}
func (this *GameMaster) ToStop() bool {
	return this.finished
}

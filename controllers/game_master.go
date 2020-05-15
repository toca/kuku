package controllers

import (
	"fmt"
	"kuku/models"
	"kuku/views"
	"time"
)

type GameMaster struct {
	renderer    *views.Renderer
	bulletView  *views.Bullet
	bulletModel *models.Bullet
	blocks      *models.Blocks
	blocksView  *views.Blocks
	blockList   []*models.Block
	finished    bool
}

func NewGameMaster() *GameMaster {
	renderer, err := views.NewRenderer(100, 41)
	if err != nil {
		panic(err)
	}
	bulletModel := models.NewBullet(2, 2, 20, 20)
	bulletView := views.NewBullet(renderer, bulletModel)

	blocks := models.NewBlocks()
	blocksView := views.NewBlocks(renderer, blocks)

	return &GameMaster{renderer, bulletView, bulletModel, blocks, blocksView, blocks.List(), false}
}

func (this *GameMaster) MainLoop() {

	defer this.renderer.Close()
	console := this.renderer.Console()

	frameModel := models.NewFrame(100, 40)
	frameView := views.NewFrame(this.renderer, frameModel)

	saucerModel := models.NewSaucer(14, 38, 84, 38)
	saucerView := views.NewSaucer(this.renderer, saucerModel)
	saucerController := NewSaucer(saucerModel)

	statusModel := models.GetStatus()
	statusModel.SetPos(0, 40)
	statusView := views.NewStatus(this.renderer, statusModel)
	statusModel.Message = "start! (to stop Ctrl + c)"

	collider := NewCollider()
	collider.AddActive(this.bulletModel)
	collider.AddActive(saucerModel)
	collider.AddMotionless(saucerModel)
	collider.AddMotionless(frameModel)
	for _, v := range this.blockList {
		collider.AddMotionless(v)
	}

	// main loop
	for !this.finished {
		// render
		frameView.Load()
		this.bulletView.Load()
		saucerView.Load()
		this.blocksView.Load()
		statusView.Load()

		this.renderer.Render()
		this.renderer.Clear()

		// operation
		if ok, key := console.ReadInput(); ok {
			saucerController.Input(key)
		}

		// action
		this.bulletModel.Action()
		// detect collidion
		collider.Detect()
		this.Action()
		time.Sleep(5 * time.Millisecond)
	}
	console.Free()
	status := models.GetStatus()
	fmt.Printf("%s. SCORE:%d\n", status.Message, status.Score)
}

func (this *GameMaster) Action() {
	status := models.GetStatus()
	for i := 0; i < len(this.blockList); i++ {
		if this.blockList[i].MarkedForDeath() {
			this.blockList = append(this.blockList[:i], this.blockList[i+1:]...)
			i++
		}
	}
	if 0 == len(this.blockList) {
		this.finished = true
		status.Message = "CLEAR!!"
		return
	}
	// bullet
	if this.bulletModel.MarkedForDeath() {
		status.BulletCount -= 1
		this.bulletModel.Reset(30, 30, 20, 20)
	}
	if status.BulletCount <= 0 {
		this.finished = true
		status.Message = "GAME OVER"
		return
	}
}

func (this *GameMaster) Stop() {
	this.finished = true
}

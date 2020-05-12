package main

import (
	"flag"
	"fmt"
	"image"
	"os"
	"os/signal"
	"syscall"
	"time"

	"kuku/controllers"
	"kuku/models"
	"kuku/views"
)

// how to build!
// env GOOS=windows GOARCH=amd64 go build
func main() {
	// TODO
	// * cursor mode?
	// * speed up
	// * colider
	// * how to die
	// * shot (space)
	// * define paramater json

	// test
	testFlag := flag.Bool("test", false, "run unit test")
	flag.Parse()
	if *testFlag {
		test()
		return
	}

	renderer, err := views.NewRenderer(100, 41)
	if err != nil {
		panic(err)
	}
	defer renderer.Close()
	console := renderer.Console()

	bulletModel := models.NewBullet(1, 1, 20, 20)
	bulletView := views.NewBullet(renderer, bulletModel)

	frameModel := models.NewFrame(100, 40)
	frameView := views.NewFrame(renderer, frameModel)

	saucerModel := models.NewSaucer(40, 38, 50, 38)
	saucerView := views.NewSaucer(renderer, saucerModel)
	saucerController := controllers.NewSaucer(saucerModel)

	block := models.NewBlock(20, 20, 50, 20, models.NORMAL_BLOCK)
	blockView := views.NewBlock(renderer, block)

	statusModel := models.GetStatus()
	statusModel.SetPos(0, 40)
	statusView := views.NewStatus(renderer, statusModel)
	statusModel.Message = "start! (to stop Ctrl + c)"

	collider := NewCollider()
	collider.AddActive(bulletModel)
	collider.AddActive(saucerModel)
	collider.AddMotionless(frameModel)
	collider.AddMotionless(saucerModel)
	collider.AddMotionless(block)
	// break
	stop := false
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	go func() {
		<-sigChan
		stop = true
		fmt.Println("interrupted!!")
	}()

	// main loop
	for !stop {
		// render
		frameView.Load()
		bulletView.Load()
		saucerView.Load()
		blockView.Load()
		statusView.Load()

		renderer.Render()
		renderer.Clear()

		// operation
		if ok, key := console.ReadInput(); ok {
			saucerController.Input(key)
		}

		// action
		bulletModel.Action()

		// detect collidion
		collider.Detect()

		time.Sleep(5 * time.Millisecond)
	}
}

func test() {
	w := 100
	h := 40
	south := image.Rect(1, h-2, w-2, h-2)
	b := image.Rect(38, 38, 38, 38)
	if !models.Overlap(&south, &b) {
		fmt.Printf("not overlaped: %v %v\n", south, b)
	} else {
		fmt.Println("ok")
	}
}

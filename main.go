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
	testFlag := flag.Bool("test", false, "run unit test")
	flag.Parse()
	if *testFlag {
		test()
		return
	}

	renderer, err := views.NewRenderer(100, 40)
	if err != nil {
		panic(err)
	}
	defer renderer.Close()
	console := renderer.Console()

	bulletModel := models.NewBullet(1, 1, 1, 1)
	bulletView := views.NewBullet(renderer, bulletModel)
	frameModel := models.NewFrame(100, 40)
	frameView := views.NewFrame(renderer, frameModel)
	saucerModel := models.NewSaucer(40, 38, 50, 38)
	saucerView := views.NewSaucer(renderer, saucerModel)
	saucerController := controllers.NewSaucer(saucerModel)

	collider := NewCollider()
	collider.AddActive(bulletModel)
	collider.AddActive(saucerModel)
	collider.AddMotionless(frameModel)

	// break
	stop := false
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	go func() {
		<-sigChan
		stop = true
		fmt.Println("interrupted!!")
	}()

	for i := 0; i < 2000 && !stop; i++ {
		time.Sleep(30 * time.Millisecond)

		// render
		frameView.Load()
		bulletView.Load()
		saucerView.Load()

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

	}
}

func test() {
	w := 100
	h := 40
	south := image.Rect(1, h-2, w-2, h-2)
	b := image.Rect(38, 38, 38, 38)
	if !models.Overlap(south, b) {
		fmt.Printf("not overlaped: %v %v\n", south, b)
	} else {
		fmt.Println("ok")
	}
}

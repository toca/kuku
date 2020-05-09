package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"./models"
	"./views"
)

// how to build!
// env GOOS=windows GOARCH=amd64 go build
func main() {
	// TODO
	// * cursor mode?
	// * speed up
	// * colider

	renderer, err := views.NewRenderer(100, 40)
	if err != nil {
		panic(err)
	}
	defer renderer.Close()

	bulletModel := models.NewBullet(1, 1, 1, 1)
	bulletView := views.NewBullet(renderer, bulletModel)
	frameModel := models.NewFrame(100, 40)
	frameView := views.NewFrame(renderer, frameModel)

	collider := NewCollider()
	collider.AddActive(bulletModel)
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

	for i := 0; i < 1000 && !stop; i++ {
		time.Sleep(5 * time.Millisecond)

		// render
		frameView.Load()
		bulletView.Load()
		renderer.Render()
		renderer.Clear()

		// action
		bulletModel.Action()

		// detect collidion
		collider.Detect()

	}
}

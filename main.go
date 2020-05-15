package main

import (
	"flag"
	"fmt"
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
	// * colider
	// * score
	// * how to die
	// * shot (space)
	// * define paramater json? for stage
	// * frame -> block base?
	// * DONE saucer ->
	// * block appearlance
	// * unbrakable block

	// test
	testFlag := flag.Bool("test", false, "run unit test")
	flag.Parse()
	if *testFlag {
		test()
		return
	}
	{
		renderer, err := views.NewRenderer(100, 41)
		if err != nil {
			panic(err)
		}
		defer renderer.Close()
		console := renderer.Console()

		bulletModel := models.NewBullet(1, 1, 60, 60)
		bulletView := views.NewBullet(renderer, bulletModel)

		frameModel := models.NewFrame(100, 40)
		frameView := views.NewFrame(renderer, frameModel)

		saucerModel := models.NewSaucer(4, 38, 94, 38)
		saucerView := views.NewSaucer(renderer, saucerModel)
		saucerController := controllers.NewSaucer(saucerModel)

		blocks := models.NewBlocks()
		blocksView := views.NewBlocks(renderer, blocks)

		statusModel := models.GetStatus()
		statusModel.SetPos(0, 40)
		statusView := views.NewStatus(renderer, statusModel)
		statusModel.Message = "start! (to stop Ctrl + c)"

		collider := NewCollider()
		collider.AddActive(bulletModel)
		collider.AddActive(saucerModel)
		collider.AddMotionless(saucerModel)
		collider.AddMotionless(frameModel)
		for _, v := range blocks.List() {
			collider.AddMotionless(v)
		}
		// break
		stop := false
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT)
		go func() {
			<-sigChan
			stop = true
			fmt.Println("interrupted!!")
		}()

		master := models.NewGameMaster(blocks.List())

		// main loop
		for !stop {
			// render
			frameView.Load()
			bulletView.Load()
			saucerView.Load()
			blocksView.Load()
			statusView.Load()

			renderer.Render()
			renderer.Clear()

			// operation
			if ok, key := console.ReadInput(); ok {
				saucerController.Input(key)
			}

			// action
			bulletModel.Action()
			master.Action()
			// detect collidion
			collider.Detect()

			stop = master.ToStop()

			time.Sleep(5 * time.Millisecond)
		}
		fmt.Printf("Finished. SCORE:%d\n>", models.GetStatus().Score)
	}

	var str string
	fmt.Scanf("%s", &str)
}

func test() {

	var str string
	fmt.Scanf("%s", &str)
	fmt.Println(str)

	// renderer, err := views.NewRenderer(100, 41)
	// if err != nil {
	// 	panic(err)
	// }
	// defer renderer.Close()

	// const PrintColor = "\033[38;5;%dm%s\033[39;49m\n"
	// for j := 0; j < 256; j++ {
	// 	fmt.Printf(PrintColor, j, "Hello!")
	// }

	// w := 100
	// h := 40
	// south := image.Rect(1, h-2, w-2, h-2)
	// b := image.Rect(38, 38, 38, 38)
	// if !models.Overlap(&south, &b) {
	// 	fmt.Printf("not overlaped: %v %v\n", south, b)
	// } else {
	// 	fmt.Println("ok")
	// }
}

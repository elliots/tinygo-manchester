package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	WindowTitle  = "tinygo-lifi"
	WindowWidth  = 800
	WindowHeight = 600

	RectWidth  = 20
	RectHeight = 20
	NumRects   = WindowHeight / RectHeight
	Interval   = time.Millisecond * 200

	Message = "hello"
)

var rects [NumRects]sdl.Rect
var runningMutex sync.Mutex

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var err error

	sdl.Do(func() {
		window, err = sdl.CreateWindow(WindowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, WindowWidth, WindowHeight, sdl.WINDOW_OPENGL)
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer func() {
		sdl.Do(func() {
			window.Destroy()
		})
	}()

	sdl.SetHint(sdl.HINT_RENDER_VSYNC, "1")

	sdl.Do(func() {
		renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	})
	if err != nil {
		log.Printf("Failed to create renderer: %s\n", err)
		return 2
	}
	defer func() {
		sdl.Do(func() {
			renderer.Destroy()
		})
	}()

	sdl.Do(func() {
		renderer.Clear()
	})

	running := true

	var outputStream []bool

	sendBool := func(data bool) {
		outputStream = append(outputStream, !data)
		outputStream = append(outputStream, data)
	}

	sendBytes := func(data []byte) {
		for _, b := range data {
			log.Printf("out: %08b\n", b)
			for j := 0; j < 8; j++ {
				sendBool(b>>uint(7-j)&0x01 == 1)
			}
		}
	}

	sendMessage := func(data []byte) {
		sendBytes([]byte{0b11111111})
		sendBytes([]byte{0b11110100})
		sendBytes([]byte{uint8(len(data))})
		sendBytes(data)
		sendBool(true)
	}

	sendMessage([]byte(Message))

	time.Sleep(time.Now().Round(time.Second).Add(time.Second).Sub(time.Now()))

	t := time.NewTicker(Interval)

	var val2 bool
	var last bool
	var first bool

	var lastTime = time.Now()
	for running {

		sdl.Do(func() {
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch event.(type) {
				case *sdl.QuitEvent:
					runningMutex.Lock()
					running = false
					runningMutex.Unlock()
				}
			}

			if len(outputStream) < 2 {
				return
			}

			val2, outputStream = outputStream[0], outputStream[1:]

			change := val2 != last
			log.Printf("val: %t last: %t change: %v", val2, last, change)

			if change || first {
				log.Printf("writing %t", val2)

				first = false
				// renderer.Clear()

				if val2 {
					renderer.SetDrawColor(0xff, 0xff, 0xff, 0xff)
				} else {
					renderer.SetDrawColor(0, 0, 0, 0xff)
				}

				renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: WindowWidth, H: WindowHeight})

				<-t.C

				renderer.Present()

				last = val2
				log.Printf("set last to %t", val2)
			} else {
				<-t.C
			}
			log.Printf("diff: %s", time.Since(lastTime))
			lastTime = time.Now()
		})

	}

	return 0
}

func main() {
	var exitcode int
	sdl.Main(func() {
		exitcode = run()
	})
	os.Exit(exitcode)
}

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell"
)

func main() {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e = s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	vi := view{Screen: s, LeftCorner: [2]int{0, 0}, RightCorner: [2]int{100, 30}, IsPaused: false}

	sn := snake{Direction: make(chan int), Body: [][2]int{{4, 2}, {2, 2}}, IsCanMove: true, Speed: 100}
	fo := food{}
	fo.creatFood(sn.Body, vi.LeftCorner, vi.RightCorner)

	quit := make(chan struct{})
	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch {
				case ev.Key() == tcell.KeyTab:
					vi.IsPaused = !vi.IsPaused
				case ev.Key() == tcell.KeyEscape && !vi.IsPaused:
					close(quit)
					close(sn.Direction)
				case ev.Key() == tcell.KeyUp && !vi.IsPaused && sn.checkDirectionSame(0):
					sn.Direction <- 0
				case ev.Key() == tcell.KeyLeft && !vi.IsPaused && sn.checkDirectionSame(1):
					sn.Direction <- 1
				case ev.Key() == tcell.KeyDown && !vi.IsPaused && sn.checkDirectionSame(2):
					sn.Direction <- 2
				case ev.Key() == tcell.KeyRight && !vi.IsPaused && sn.checkDirectionSame(3):
					sn.Direction <- 3
				}
			}
		}
	}()

loop:
	for {
		if !vi.IsPaused {
			select {
			case <-quit:
				break loop
			case direction := <-sn.Direction:
				sn.moveStep(direction)
			case <-time.After(time.Millisecond * sn.Speed):
				sn.moveStep(-1)
			}
			if fo.isTouchMe(sn.Body) {
				fo.creatFood(sn.Body, vi.LeftCorner, vi.RightCorner)
				sn.eatFood()
			}

			if sn.isTouchSelf() || sn.isTouchWall(vi.LeftCorner, vi.RightCorner) {
				break loop
			}

			vi.updateView(sn, fo)
		}
	}

	s.Fini()
}

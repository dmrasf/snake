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
	vi := view{Screen: s, LeftCorner: [2]int{0, 0}, RightCorner: [2]int{50, 16}}

	sn := snake{Direction: 3, Body: [][2]int{{4, 2}, {2, 2}}, IsCanMove: true}
	fo := food{}
	fo.creatFood(sn.Body, vi.LeftCorner, vi.RightCorner)

	quit := make(chan struct{})
	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					close(quit)
					return
				case tcell.KeyUp:
					sn.Direction = 0
				case tcell.KeyLeft:
					sn.Direction = 1
				case tcell.KeyDown:
					sn.Direction = 2
				case tcell.KeyRight:
					sn.Direction = 3
				}
			}
		}
	}()

loop:
	for {
		select {
		case <-quit:
			break loop
		case <-time.After(time.Millisecond * 100):
		}
		sn.moveStep()

		if fo.isTouchMe(sn.Body) {
			fo.creatFood(sn.Body, vi.LeftCorner, vi.RightCorner)
			sn.eatFood()
		}

		if sn.isTouchSelf() || sn.isTouchWall(vi.LeftCorner, vi.RightCorner) {
			break loop
		}

		vi.updateView(sn.Body, fo.Pos)
	}

	s.Fini()
}

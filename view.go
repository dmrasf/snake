package main

import (
	"github.com/gdamore/tcell"
)

type view struct {
	Screen      tcell.Screen
	LeftCorner  [2]int
	RightCorner [2]int
}

func (v *view) getSize() (int, int) {
	return v.Screen.Size()
}

func (v *view) drawSnake(snake [][2]int) {
	rgb := tcell.NewHexColor(0xffffff)
	st := tcell.StyleDefault.Background(rgb)
	for _, item := range snake {
		v.Screen.SetCell(item[0], item[1], st, ' ')
		v.Screen.SetCell(item[0]+1, item[1], st, ' ')
	}
}

func (v *view) drawBorder() {
	rgb := tcell.NewHexColor(0x00ffff)
	st := tcell.StyleDefault.Background(rgb)
	for y := v.LeftCorner[1]; y <= v.RightCorner[1]; y++ {
		v.Screen.SetCell(v.LeftCorner[0], y, st, ' ')
		v.Screen.SetCell(v.LeftCorner[0]+1, y, st, ' ')
		v.Screen.SetCell(v.RightCorner[0], y, st, ' ')
		v.Screen.SetCell(v.RightCorner[0]+1, y, st, ' ')
	}
	for x := v.LeftCorner[0]; x <= v.RightCorner[0]; x++ {
		v.Screen.SetCell(x, v.LeftCorner[1], st, ' ')
		v.Screen.SetCell(x+1, v.LeftCorner[1], st, ' ')
		v.Screen.SetCell(x, v.RightCorner[1], st, ' ')
		v.Screen.SetCell(x+1, v.RightCorner[1], st, ' ')
	}
}

func (v *view) drawFood(f [2]int) {
	rgb := tcell.NewHexColor(0x00ff00)
	st := tcell.StyleDefault.Background(rgb)
	v.Screen.SetCell(f[0], f[1], st, ' ')
	v.Screen.SetCell(f[0]+1, f[1], st, ' ')
}

func (v *view) updateView(snake [][2]int, food [2]int) {
	v.Screen.Clear()
	v.drawSnake(snake)
	v.drawBorder()
	v.drawFood(food)
	v.Screen.Show()
}

package main

import (
	"strconv"

	"github.com/gdamore/tcell"
)

type view struct {
	Screen      tcell.Screen
	LeftCorner  [2]int
	RightCorner [2]int
	IsPaused    bool
}

func (v *view) getSize() (int, int) {
	return v.Screen.Size()
}

func (v *view) drawSnake(snake [][2]int) {
	rgb := tcell.NewHexColor(0xffffff)
	st := tcell.StyleDefault.Background(rgb)
	v.Screen.SetCell(snake[0][0], snake[0][1], st, ' ')
	v.Screen.SetCell(snake[0][0]+1, snake[0][1], st, ' ')
	rgb = tcell.NewHexColor(0xf0ff0f)
	st = tcell.StyleDefault.Background(rgb)
	for _, item := range snake[1:] {
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

func (v *view) drawInstruction() {
	str := "Controll: up left down right; Exit: esc; Pause: tab"
	for i, item := range str {
		v.Screen.SetContent(v.LeftCorner[0]+i, v.RightCorner[1]+1,
			item, nil, tcell.StyleDefault)
	}
}

func (v *view) drawFood(f [2]int) {
	rgb := tcell.NewHexColor(0x00ff00)
	st := tcell.StyleDefault.Background(rgb)
	v.Screen.SetCell(f[0], f[1], st, ' ')
	v.Screen.SetCell(f[0]+1, f[1], st, ' ')
}

func (v *view) drawScore(sn snake) {
	score := len(sn.Body)
	str := "Score: " + strconv.Itoa(score)
	for i, item := range str {
		v.Screen.SetContent(v.LeftCorner[0]+i, v.RightCorner[1]+2,
			item, nil, tcell.StyleDefault)
	}
}

func (v *view) updateView(sn snake, fo food) {
	v.Screen.Clear()
	v.drawSnake(sn.Body)
	v.drawBorder()
	v.drawFood(fo.Pos)
	v.drawInstruction()
	v.drawScore(sn)
	v.Screen.Show()
}

package main

import (
	"math/rand"
	"time"
)

type food struct {
	Pos [2]int
}

func (f *food) creatFood(snake [][2]int, L [2]int, R [2]int) {
	rand.Seed(time.Now().UnixNano())
	for {
		posX := (rand.Intn((R[0]-L[0]-2)/2)+1)*2 + L[0]
		posY := rand.Intn(R[1]-L[1]-1) + L[1] + 1
		for _, item := range snake {
			if posX != item[0] || posY != item[1] {
				f.Pos[0], f.Pos[1] = posX, posY
				return
			}
		}
	}
}

func (f *food) isTouchMe(snake [][2]int) bool {
	if snake[0][0] == f.Pos[0] && snake[0][1] == f.Pos[1] {
		return true
	}
	return false
}

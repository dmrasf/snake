package main

type snake struct {
	//  0; 1; 2; 3
	Direction     int
	Body          [][2]int
	IsCanMove     bool
	TailDirection int
}

// Get the current direction of the snake
func (sn *snake) getCurrentDirection() int {
	ox := sn.Body[0][0] - sn.Body[1][0]
	oy := sn.Body[0][1] - sn.Body[1][1]
	switch {
	case ox == 0 && oy > 0:
		return 2
	case ox == 0 && oy < 0:
		return 0
	case ox > 0 && oy == 0:
		return 3
	case ox < 0 && oy == 0:
		return 1
	default:
		return -1
	}
}

func (sn *snake) setTailDirection() {
	l := len(sn.Body)
	ox := sn.Body[l-1][0] - sn.Body[l-2][0]
	oy := sn.Body[l-1][1] - sn.Body[l-2][1]
	switch {
	case ox == 0 && oy > 0:
		sn.TailDirection = 2
	case ox == 0 && oy < 0:
		sn.TailDirection = 0
	case ox > 0 && oy == 0:
		sn.TailDirection = 3
	case ox < 0 && oy == 0:
		sn.TailDirection = 1
	}
}

// Snake moves one step
func (sn *snake) moveStep() {
	sn.setTailDirection()

	currentDirection := sn.getCurrentDirection()
	var dir int
	for i := len(sn.Body) - 1; i > 0; i-- {
		sn.Body[i][0] = sn.Body[i-1][0]
		sn.Body[i][1] = sn.Body[i-1][1]
	}
	if currentDirection%2 == sn.Direction%2 {
		dir = currentDirection
	} else {
		dir = sn.Direction
	}
	switch dir {
	case 0:
		sn.Body[0][1] -= 1
	case 1:
		sn.Body[0][0] -= 2
	case 2:
		sn.Body[0][1] += 1
	case 3:
		sn.Body[0][0] += 2
	}

}

// Add food to the tail
func (sn *snake) eatFood() {
	newBody := sn.Body[len(sn.Body)-1]
	switch sn.TailDirection {
	case 0:
		newBody[1] -= 1
	case 1:
		newBody[0] -= 2
	case 2:
		newBody[1] += 1
	case 3:
		newBody[0] += 2
	}
	sn.Body = append(sn.Body, newBody)
}

func (sn *snake) isTouchSelf() bool {
	for _, item := range sn.Body[1:] {
		if item[0] == sn.Body[0][0] && item[1] == sn.Body[0][1] {
			return true
		}
	}
	return false
}

func (sn *snake) isTouchWall(L, R [2]int) bool {
	headX, headY := sn.Body[0][0], sn.Body[0][1]
	if headX <= L[0] || headX >= R[0] || headY <= L[1] || headY >= R[1] {
		return true
	}
	return false
}

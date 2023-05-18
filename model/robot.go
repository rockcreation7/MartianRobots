package model

import (
	"fmt"
)

const (
	North = iota
	East
	South
	West
)

type Robot struct {
	X, Y         int16
	PrevX, PrevY int16
	Direction    int16
	IsLost       bool
}

func (robot *Robot) Left() {
	if robot.Direction == North {
		robot.Direction = West
		return
	}
	robot.Direction = robot.Direction - 1
}

func (robot *Robot) Right() {
	if robot.Direction == West {
		robot.Direction = North
		return
	}
	robot.Direction = robot.Direction + 1
}

func (robot *Robot) Forward(mars *Mars) {

	//Stop when scent on lost direction
	posKey := fmt.Sprintf("%d%d", robot.X, robot.PrevY)
	noMove := false
	for _, scent := range mars.Scent[posKey] {
		if scent == robot.Direction {
			noMove = true
			break
		}
	}
	if noMove {
		return
	}

	// Save current location
	robot.PrevX = robot.X
	robot.PrevY = robot.Y

	// Command execution
	switch robot.Direction {
	case North:
		robot.Y += 1
	case East:
		robot.X += 1
	case South:
		robot.Y -= 1
	case West:
		robot.X -= 1
	}

	// If it is lost
	if mars.isOutBound(robot) {
		robot.IsLost = true
	}
	return
}

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
	prevX, prevY int16
	Direction    int16
}

var DirectionToEnum = map[string]int16{
	"N": North,
	"E": East,
	"S": South,
	"W": West,
}

var EnumToDirection = map[int16]string{
	North: "N",
	East:  "E",
	South: "S",
	West:  "W",
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

func (robot *Robot) Forward(mars *Mars) (lost bool) {

	//Stop when scent on lost direction
	posKey := fmt.Sprintf("%d%d", robot.X, robot.prevY)
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
	robot.prevX = robot.X
	robot.prevY = robot.Y

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
		lost = true
	}
	return
}

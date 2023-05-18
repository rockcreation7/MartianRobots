package model

import "fmt"

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

type Mars struct {
	UpperRightX, UpperRightY int16
	Scent                    map[string][]int16
}

func (mars *Mars) isOutBound(robot *Robot) (out bool) {
	posKey := fmt.Sprintf("%d%d", robot.prevX, robot.prevY)

	if robot.X < 0 || (robot.X == 0 && robot.Y < 0) {
		mars.Scent[posKey] = append(mars.Scent[posKey], robot.Direction)
		out = true
		return
	}

	if robot.Y < 0 || (robot.Y < 0 && robot.X < 0) {
		mars.Scent[posKey] = append(mars.Scent[posKey], robot.Direction)
		out = true
		return
	}

	if robot.X > mars.UpperRightX || (robot.X == mars.UpperRightX && robot.Y > mars.UpperRightY) {
		mars.Scent[posKey] = append(mars.Scent[posKey], robot.Direction)
		out = true
		return
	}

	if robot.Y > mars.UpperRightY || (robot.Y == mars.UpperRightY && robot.X > mars.UpperRightX) {
		mars.Scent[posKey] = append(mars.Scent[posKey], robot.Direction)
		out = true
		return
	}

	return
}

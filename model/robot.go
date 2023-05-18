package model

const (
	North = iota
	East
	South
	West
)

type Robot struct {
	X, Y      int16
	Direction int16
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

	if mars.isOutBound(robot) {
		lost = true
	}
	return
}

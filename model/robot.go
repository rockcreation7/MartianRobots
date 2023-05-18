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

func (robot *Robot) Left() {
	if robot.Direction == North {
		robot.Direction = East
		return
	}
	robot.Direction = robot.Direction - 1
}

package model

type Mars struct {
	UpperRightX, UpperRightY int16
	Scent                    map[string]int16
}

func (mars *Mars) isOutBound(robot *Robot) (out bool) {

	if robot.X < 0 || robot.Y < 0 || robot.X > mars.UpperRightX || robot.Y > mars.UpperRightY {
		out = true
		return
	}

	return
}

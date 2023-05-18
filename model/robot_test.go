package model

import "testing"

// go test -run Test_Robot -v
func Test_Robot(t *testing.T) {
	robot := Robot{
		X:         0,
		Y:         0,
		Direction: East,
	}

	robot.Left()
	robot.Left()

	if robot.Direction != West {
		t.Errorf("Robot turn left have problem ")
	}

}

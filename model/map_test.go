package model

import "testing"

// go test -run Test_Map -v
func Test_Map(t *testing.T) {
	mars := &Mars{
		UpperRightX: 10,
		UpperRightY: 10,
		Scent:       map[string]int16{},
	}
	robot := &Robot{
		X:         -1,
		Y:         0,
		Direction: South,
	}

	if !mars.isOutBound(robot) {
		t.Errorf("isOutBound error on robot %d %d", robot.X, robot.Y)
	}

	if len(mars.Scent) != 1 {
		t.Error("Scent error")
	}
}

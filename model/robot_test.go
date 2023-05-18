package model

import (
	"fmt"
	"testing"
)

// go test -run Test_Robot -v
func Test_Robot(t *testing.T) {
	robot := Robot{
		X:         0,
		Y:         0,
		Direction: South,
	}

	robot.Left()
	robot.Left()
	robot.Left()

	if robot.Direction != West {
		t.Errorf("Robot turn left have problem ")
	}
	fmt.Println(EnumToDirection[robot.Direction])

	robot.Right()
	robot.Right()
	robot.Right()

	if robot.Direction != South {
		t.Errorf("Robot turn Right have problem ")
	}
	fmt.Println(EnumToDirection[robot.Direction])

	robot.Left()
	robot.Left()
	fmt.Println(EnumToDirection[robot.Direction])
	robot.Forward()
	robot.Forward()

	if robot.Y != 2 {
		t.Errorf(fmt.Sprintf("Robot Forward have problem X: %d Y: %d", robot.X, robot.Y))
	}

}

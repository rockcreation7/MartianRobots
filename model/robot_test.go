package model

import (
	"fmt"
	"testing"
)

// go test -run Test_Robot_Lost -v
func Test_Robot_Lost(t *testing.T) {
	robot := &Robot{
		X:         10,
		Y:         10,
		Direction: East,
	}
	mars := &Mars{
		UpperRightX: 10,
		UpperRightY: 10,
		Scent:       map[string][]int16{},
	}
	if !robot.Forward(mars) {
		t.Error("Some issue on bound calculate - Robot")
	} else {
		fmt.Println("Robot Lost!")
	}

	robot2 := &Robot{
		X:         10,
		Y:         10,
		Direction: East,
	}

	if robot2.Forward(mars) && robot2.Y == 10 {
		fmt.Println("Robot2 Scent! No Move!")
	} else {
		t.Error("Some issue on bound calculate - Robot2")
	}
}

// go test -run Test_Robot -v
func Test_Robot(t *testing.T) {
	robot := &Robot{
		X:         0,
		Y:         0,
		Direction: South,
	}

	robot.Left()
	robot.Left()
	robot.Left()

	if robot.Direction != West {
		t.Error("Robot turn left have problem ")
	}
	fmt.Println(EnumToDirection[robot.Direction])

	robot.Right()
	robot.Right()
	robot.Right()

	if robot.Direction != South {
		t.Error("Robot turn Right have problem ")
	}
	fmt.Println(EnumToDirection[robot.Direction])
	mars := &Mars{
		UpperRightX: 10,
		UpperRightY: 10,
		Scent:       map[string][]int16{},
	}
	robot.Left()
	robot.Left()
	fmt.Println(EnumToDirection[robot.Direction])
	robot.Forward(mars)
	robot.Forward(mars)

	if robot.Y != 2 {
		t.Errorf("Robot Forward have problem X: %d Y: %d", robot.X, robot.Y)
	}

}

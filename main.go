package main

import (
	"fmt"
	"strings"

	"github.com/rockcreation7/MartianRobots/model"
)

func main() {
	// fmt.Println("Start")
	command := &model.Command{
		UpperRightCorner: [2]int16{5, 3},
		IntructionList: []*model.RobotIntruction{
			{
				RobotPoint:     [2]int16{1, 100},
				RobotDirection: "E",
				Intruction:     "RFRFRFRF",
			},
			{
				RobotPoint:     [2]int16{3, 2},
				RobotDirection: "N",
				Intruction:     "FRRFLLFFRRFLL",
			},
			{
				RobotPoint:     [2]int16{0, 3},
				RobotDirection: "W",
				Intruction:     "LLFFFLFLFL",
			},
		},
	}

	mars := &model.Mars{
		UpperRightX: command.UpperRightCorner[0],
		UpperRightY: command.UpperRightCorner[1],
		Scent:       map[string][]int16{},
	}

	execute(command, mars)
	// fmt.Println(command)
}

func execute(command *model.Command, mars *model.Mars) {
	if len(command.IntructionList) > 0 {
		IntructionData := command.IntructionList[0]
		command.IntructionList = command.IntructionList[1:]

		// input vaildation
		if len(IntructionData.RobotPoint) != 2 || IntructionData.RobotPoint[0] > 50 ||
			IntructionData.RobotPoint[1] > 50 || len(IntructionData.Intruction) > 99 {
			fmt.Println("Command Invaild, try next command")
			execute(command, mars)
			return
		}

		robot := model.Robot{
			X:         IntructionData.RobotPoint[0],
			Y:         IntructionData.RobotPoint[1],
			Direction: model.DirectionToEnum[IntructionData.RobotDirection],
		}

		executable := strings.Split(IntructionData.Intruction, "")

		// fmt.Println(executable, robot, "executable robot")
	L:
		for _, execute := range executable {
			switch execute {
			case "L":
				robot.Left()
			case "R":
				robot.Right()
			case "F":
				robot.Forward(mars)
				if robot.IsLost {
					fmt.Println(robot.PrevX, robot.PrevY, model.EnumToDirection[robot.Direction], "LOST")
					break L
				}
			}
		}
		if !robot.IsLost {
			fmt.Println(robot.X, robot.Y, model.EnumToDirection[robot.Direction])
		}
		execute(command, mars)
	}
}

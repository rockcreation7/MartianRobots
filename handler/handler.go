package handler

import (
	"fmt"
	"strings"

	"github.com/rockcreation7/MartianRobots/model"
)

func Execute(command *model.Command, mars *model.Mars) {
	if len(command.IntructionList) > 0 {
		IntructionData := command.IntructionList[0]
		command.IntructionList = command.IntructionList[1:]

		// input vaildation
		if len(IntructionData.RobotPoint) != 2 || IntructionData.RobotPoint[0] > 50 ||
			IntructionData.RobotPoint[1] > 50 || len(IntructionData.Intruction) > 99 {
			fmt.Println("Command Invaild, try next command")
			Execute(command, mars)
			return
		}

		robot := model.Robot{
			X:         IntructionData.RobotPoint[0],
			Y:         IntructionData.RobotPoint[1],
			Direction: model.DirectionToEnum[IntructionData.RobotDirection],
		}

		executable := strings.Split(IntructionData.Intruction, "")

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
		Execute(command, mars)
	}
}

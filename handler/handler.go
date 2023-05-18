package handler

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
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

func Readtxt(cmd *model.Command) error {
	data, err := ioutil.ReadFile("command.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		// fmt.Println(line, i)
		if i == 0 {
			str := strings.Split(line, " ")
			if len(str) != 2 {
				return errors.New(fmt.Sprintf("Read Error line: %d %d", i, len(str)))
			}

			i1, err := strconv.Atoi(str[0])
			if err != nil {
				return err
			}

			i2, err := strconv.Atoi(str[1])
			if err != nil {
				return err
			}
			cmd.UpperRightCorner = [2]int16{int16(i1), int16(i2)}
			cmd.IntructionList = append(cmd.IntructionList, &model.RobotIntruction{})
			continue
		}
		if cmd.IntructionList[len(cmd.IntructionList)-1].RobotPoint == [2]int16{} {
			str := strings.Split(line, " ")

			if len(str) != 3 {
				return errors.New(fmt.Sprintf("Read Error line: %d %d", i, len(str)))
			}

			i1, err := strconv.Atoi(str[0])
			if err != nil {
				return err
			}

			i2, err := strconv.Atoi(str[1])
			if err != nil {
				return err
			}
			cmd.IntructionList[len(cmd.IntructionList)-1].RobotPoint = [2]int16{int16(i1), int16(i2)}
			cmd.IntructionList[len(cmd.IntructionList)-1].RobotDirection = str[2]
			continue
		} else if cmd.IntructionList[len(cmd.IntructionList)-1].Intruction == "" {
			cmd.IntructionList[len(cmd.IntructionList)-1].Intruction = line
			continue
		} else if len(line) == 0 {
			cmd.IntructionList = append(cmd.IntructionList, &model.RobotIntruction{})
		}
	}

	return nil
}

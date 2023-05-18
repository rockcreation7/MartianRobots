package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/rockcreation7/MartianRobots/handler"
	"github.com/rockcreation7/MartianRobots/model"
)

func main() {
	command := &model.Command{
		/* UpperRightCorner: [2]int16{5, 3},
		IntructionList: []*model.RobotIntruction{
			{
				RobotPoint:     [2]int16{1, 1},
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
		}, */
	}
	err := Readtxt(command)
	mars := &model.Mars{
		UpperRightX: command.UpperRightCorner[0],
		UpperRightY: command.UpperRightCorner[1],
		Scent:       map[string][]int16{},
	}

	handler.Execute(command, mars)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func Readtxt(cmd *model.Command) error {
	data, err := ioutil.ReadFile("command.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		fmt.Println(line, i)
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

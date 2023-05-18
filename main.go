package main

import (
	"github.com/rockcreation7/MartianRobots/handler"
	"github.com/rockcreation7/MartianRobots/model"
)

func main() {
	command := &model.Command{
		UpperRightCorner: [2]int16{5, 3},
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
		},
	}

	mars := &model.Mars{
		UpperRightX: command.UpperRightCorner[0],
		UpperRightY: command.UpperRightCorner[1],
		Scent:       map[string][]int16{},
	}

	handler.Execute(command, mars)
}

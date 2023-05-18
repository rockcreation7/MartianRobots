package main

import (
	"fmt"

	"github.com/rockcreation7/MartianRobots/handler"
	"github.com/rockcreation7/MartianRobots/model"
)

func main() {
	command := &model.Command{}
	err := handler.Readtxt(command)

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

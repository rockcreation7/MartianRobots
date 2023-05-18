package model

type Command struct {
	UpperRightCorner [2]int16
	IntructionList   []RobotIntruction
}

type RobotIntruction struct {
	RobotPoint     [2]int16
	RobotDirection string
	Intruction     string
}

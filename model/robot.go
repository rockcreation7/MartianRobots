package model

const (
	North = iota
	East
	South
	West
)

type Robot struct {
	X, Y      int16
	Direction int16
}

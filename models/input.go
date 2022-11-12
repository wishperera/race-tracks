package models

type Input struct {
	GridLength int
	GridWidth  int
	Start      Coordinate
	Target     Coordinate
	Obstacles  []Obstacles
}

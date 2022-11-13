package models

// Input : holds the input information
type Input struct {
	// GridLength : length of the 2d matrix
	GridLength int
	// GridWidth : width of the 2d matrix
	GridWidth int
	// Start : hopper's starting position on the matrix
	Start Coordinate
	// Target : hopper's ending position on the matrix
	Target Coordinate
	// Obstacles : array of obstacle boundaries
	Obstacles []Obstacles
}

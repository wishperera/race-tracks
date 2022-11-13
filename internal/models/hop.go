package models

// Hop : represents the state of the hopper at a given time
type Hop struct {
	// CurrentVelocity : the current velocity of the hopper in x and y directions
	CurrentVelocity Velocity
	// CurrentPosition : coordinates of the current position of the hopper
	CurrentPosition Coordinate
	// HopCount : number of hops taken by the hopper from start to reach the current state
	HopCount int
}

package models

// Grid : represents the board the hopper is moving on
type Grid struct {
	length    int
	width     int
	obstacles map[int]map[int]bool
	// visited is a map of all visited nodes with the same velocity during a minimum hop count search
	visited map[Coordinate]map[Velocity]bool
}

// Obstacles : represent the boundaries for obstacles.
// all cells with coordinates (x,y) such that  X1 <= x <= X2 and Y1 <= y <= Y2 qualify as obstacles
type Obstacles struct {
	X1 int
	X2 int
	Y1 int
	Y2 int
}

// NewGrid : returns a new grid with given length, width and defined obstacles
func NewGrid(length, width int, obstacles []Obstacles) *Grid {
	grid := new(Grid)
	grid.length = length
	grid.width = width
	grid.obstacles = make(map[int]map[int]bool)
	grid.visited = make(map[Coordinate]map[Velocity]bool)

	for _, obs := range obstacles {
		for i := obs.X1; i <= obs.X2; i++ {
			for j := obs.Y1; j <= obs.Y2; j++ {
				if _, ok := grid.obstacles[i]; !ok {
					grid.obstacles[i] = make(map[int]bool)
				}

				grid.obstacles[i][j] = true
			}
		}
	}

	return grid
}

// IsInside : checks if a coordinate is inside the grid
func (g *Grid) IsInside(coordinate Coordinate) bool {
	return g.length > coordinate.X && g.width > coordinate.Y && (coordinate.X >= 0) && (coordinate.Y >= 0)
}

// IsBlocked : checks if a coordinate is blocked by an obstacle
func (g *Grid) IsBlocked(coordinate Coordinate) bool {
	return g.obstacles[coordinate.X][coordinate.Y]
}

// MarkVisited : mark the coordinate as visited
func (g *Grid) MarkVisited(coordinate Coordinate, velocity Velocity) {
	_, ok := g.visited[coordinate]
	if !ok {
		g.visited[coordinate] = make(map[Velocity]bool)
	}

	g.visited[coordinate][velocity] = true
}

// IsVisited : returns true if the coordinate has been visited before
func (g *Grid) IsVisited(coordinate Coordinate, velocity Velocity) bool {
	return g.visited[coordinate][velocity]
}

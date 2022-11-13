package models

type Grid struct {
	length    int
	width     int
	obstacles map[int]map[int]bool
	visited   map[int]map[int]bool
}

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
	grid.visited = make(map[int]map[int]bool)

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
func (g *Grid) MarkVisited(coordinate Coordinate) {
	_, ok := g.visited[coordinate.X]
	if !ok {
		g.visited[coordinate.X] = make(map[int]bool)
	}

	g.visited[coordinate.X][coordinate.Y] = true
}

// IsVisited : returns true if the coordinate has been visited before
func (g *Grid) IsVisited(coordinate Coordinate) bool {
	return g.visited[coordinate.X][coordinate.Y]
}

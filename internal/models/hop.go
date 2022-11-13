package models

import "fmt"

type Hop struct {
	CurrentVelocity Velocity
	CurrentPosition Coordinate
	HopCount        int
	Prev            *Hop
}

func (h *Hop) PrintPath() {
	order := make([]Coordinate, 0)
	tmp := h
	for tmp != nil {
		order = append(order, tmp.CurrentPosition)
		tmp = tmp.Prev
	}

	path := ""
	for i := len(order) - 1; i >= 0; i-- {
		path += fmt.Sprintf("(%d,%d)", order[i].X, order[i].Y)
		if i != 0 {
			path += "-->"
		}
	}

	fmt.Println(path)
}

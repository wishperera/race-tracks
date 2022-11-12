package service

import (
	"testing"

	models2 "github.com/wishperera/race-tracks/internal/models"
)

func TestFindMinimumHops(t *testing.T) {
	type args struct {
		start  models2.Coordinate
		target models2.Coordinate
		grid   *models2.Grid
	}
	tests := []struct {
		name         string
		args         args
		wantHopCount int
	}{
		{
			name: "optimal solution exists",
			args: args{
				start:  models2.Coordinate{X: 4, Y: 0},
				target: models2.Coordinate{X: 4, Y: 4},
				grid:   models2.NewGrid(5, 5, []models2.Obstacles{{X1: 1, X2: 4, Y1: 2, Y2: 3}}),
			},
			wantHopCount: 7,
		},
		{
			name: "no solution exists",
			args: args{
				start:  models2.Coordinate{X: 0, Y: 0},
				target: models2.Coordinate{X: 2, Y: 2},
				grid: models2.NewGrid(3, 3, []models2.Obstacles{
					{X1: 1, X2: 1, Y1: 0, Y2: 2},
					{X1: 0, X2: 2, Y1: 1, Y2: 1},
				}),
			},
			wantHopCount: -1,
		},
	}

	for _, tt := range tests {
		temp := tt
		t.Run(temp.name, func(t *testing.T) {
			if gotHopCount := FindMinimumHops(tt.args.start, tt.args.target, temp.args.grid); gotHopCount != tt.wantHopCount {
				t.Errorf("FindMinimumHops() = %v, want %v", gotHopCount, tt.wantHopCount)
			}
		})
	}
}

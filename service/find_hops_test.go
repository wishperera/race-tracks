package service

import (
	"github.com/wishperera/race-tracks/models"
	"testing"
)

func TestFindMinimumHops(t *testing.T) {
	type args struct {
		start  models.Coordinate
		target models.Coordinate
		grid   *models.Grid
	}
	tests := []struct {
		name         string
		args         args
		wantHopCount int
	}{
		{
			name: "optimal solution exists",
			args: args{
				start:  models.Coordinate{X: 4, Y: 0},
				target: models.Coordinate{X: 4, Y: 4},
				grid:   models.NewGrid(5, 5, []models.Obstacles{{X1: 1, X2: 4, Y1: 2, Y2: 3}}),
			},
			wantHopCount: 7,
		},
		{
			name: "no solution exists",
			args: args{
				start:  models.Coordinate{X: 0, Y: 0},
				target: models.Coordinate{X: 2, Y: 2},
				grid: models.NewGrid(3, 3, []models.Obstacles{
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

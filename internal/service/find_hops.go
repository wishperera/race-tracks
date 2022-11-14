package service

import (
	"log"

	models2 "github.com/wishperera/race-tracks/internal/models"
	"github.com/wishperera/race-tracks/internal/pkg/queue"
)

// FindMinimumHops : returns the minimum number of hops required for a hopper to reach the target
// if a solution is found, otherwise returns -1
func FindMinimumHops(start, target models2.Coordinate, grid *models2.Grid) (hopCount int) {
	if isTargetReached(start, target) {
		log.Println("no hops necessary start and target are the same")
		return 0
	}

	hq := queue.NewQueue()
	grid.MarkVisited(start, models2.Velocity{})
	hq.Enqueue(models2.Hop{
		CurrentVelocity: models2.Velocity{
			XVelocity: 0,
			YVelocity: 0,
		},
		CurrentPosition: start,
		HopCount:        0,
	})

	for !hq.Empty() {
		temp := hq.Dequeue()

		for _, v := range getPossibleVelocities(temp.CurrentVelocity) {
			// the hopper cannot move if both velocities are zero
			if v.XVelocity == 0 && v.YVelocity == 0 {
				continue
			}

			nextPosition := models2.Coordinate{
				X: temp.CurrentPosition.X + v.XVelocity,
				Y: temp.CurrentPosition.Y + v.YVelocity,
			}

			if grid.IsInside(nextPosition) && !grid.IsBlocked(nextPosition) && !grid.IsVisited(nextPosition, v) {
				if isTargetReached(nextPosition, target) {
					return temp.HopCount + 1
				}

				grid.MarkVisited(nextPosition, v)

				hq.Enqueue(models2.Hop{
					CurrentPosition: nextPosition,
					CurrentVelocity: v,
					HopCount:        temp.HopCount + 1,
				})
			}
		}
	}

	return -1
}

func isTargetReached(position, target models2.Coordinate) bool {
	return position.X == target.X && position.Y == target.Y
}

func getPossibleVelocities(velocity models2.Velocity) []models2.Velocity {
	return []models2.Velocity{
		{
			XVelocity: velocity.XVelocity - 1,
			YVelocity: velocity.YVelocity - 1,
		},
		{
			XVelocity: velocity.XVelocity,
			YVelocity: velocity.YVelocity,
		},
		{
			XVelocity: velocity.XVelocity + 1,
			YVelocity: velocity.YVelocity + 1,
		},
		{
			XVelocity: velocity.XVelocity - 1,
			YVelocity: velocity.YVelocity,
		},
		{
			XVelocity: velocity.XVelocity,
			YVelocity: velocity.YVelocity - 1,
		},
		{
			XVelocity: velocity.XVelocity,
			YVelocity: velocity.YVelocity + 1,
		},
		{
			XVelocity: velocity.XVelocity + 1,
			YVelocity: velocity.YVelocity,
		},
		{
			XVelocity: velocity.XVelocity - 1,
			YVelocity: velocity.YVelocity + 1,
		},
		{
			XVelocity: velocity.XVelocity + 1,
			YVelocity: velocity.YVelocity - 1,
		},
	}
}

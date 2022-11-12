package service

import (
	"github.com/wishperera/race-tracks/models"
	"github.com/wishperera/race-tracks/queue"
	"log"
)

// FindMinimumHops : returns the minimum number of hops required for a hopper to reach the target
// if a solution is found, otherwise returns -1
func FindMinimumHops(start, target models.Coordinate, grid *models.Grid) (hopCount int) {
	if isTargetReached(start, target) {
		log.Println("no hops necessary start and target are the same")
		return 0
	}

	hq := queue.NewQueue()
	hq.Enqueue(models.Hop{
		CurrentVelocity: models.Velocity{
			XVelocity: 0,
			YVelocity: 0,
		},
		CurrentPosition: start,
		HopCount:        0,
	})

	for !hq.Empty() {
		temp := hq.Dequeue()

		if isTargetReached(temp.CurrentPosition, target) {
			return temp.HopCount
		}

		for _, v := range getPossibleVelocities(temp.CurrentVelocity) {
			nextPosition := models.Coordinate{
				X: temp.CurrentPosition.X + v.XVelocity,
				Y: temp.CurrentPosition.Y + v.YVelocity,
			}

			// the hopper cannot move if both velocities are zero
			if nextPosition.X == 0 && nextPosition.Y == 0 {
				continue
			}

			if grid.IsInside(nextPosition) && !grid.IsBlocked(nextPosition) {
				hq.Enqueue(models.Hop{
					CurrentPosition: nextPosition,
					CurrentVelocity: v,
					HopCount:        temp.HopCount + 1,
				})
			}
		}

	}

	return -1
}

func isTargetReached(position, target models.Coordinate) bool {
	return position.X == target.X && position.Y == target.Y
}

func getPossibleVelocities(velocity models.Velocity) []models.Velocity {
	return []models.Velocity{
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

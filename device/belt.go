package device

import (
	"math"
)

type Belt struct {
	Speed     float64
	Pos       Position
	Direction string
}

func (belt *Belt) BeltMove(car *CarModel, timeDuration int64, direction int) {
	if direction > 0 {
		if belt.Direction == "X" {
			car.NowPos.PosX = math.Min(car.Destination.PosX, car.NowPos.PosY+belt.Speed*float64(timeDuration/1e9))
		} else if belt.Direction == "Y" {
			car.NowPos.PosY = math.Min(car.Destination.PosY, car.NowPos.PosY+belt.Speed*float64(timeDuration/1e9))
		} else {
			car.NowPos.PosZ = math.Min(car.Destination.PosZ, car.NowPos.PosZ+belt.Speed*float64(timeDuration/1e9))
		}
	} else {
		if belt.Direction == "X" {
			car.NowPos.PosX = math.Max(car.Destination.PosX, car.NowPos.PosY-belt.Speed*float64(timeDuration/1e9))
		} else if belt.Direction == "Y" {
			car.NowPos.PosY = math.Max(car.Destination.PosY, car.NowPos.PosY-belt.Speed*float64(timeDuration/1e9))
		} else {
			car.NowPos.PosZ = math.Max(car.Destination.PosZ, car.NowPos.PosZ-belt.Speed*float64(timeDuration/1e9))
		}
	}
}

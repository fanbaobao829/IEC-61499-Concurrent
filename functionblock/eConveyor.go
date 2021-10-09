package functionblock

import "IEC-61499-Concurrent/device"

type EConveyor struct {
	FbInfo
}

const (
	PositiveDirection = 1
	NegativeDirection = -1
)

func (nowFb *EConveyor) Execute(car *device.CarModel, eventIn string) {
	nowFb.DeviceMapping.(*device.Belt).BeltMove(car, CycleTime, PositiveDirection)
}
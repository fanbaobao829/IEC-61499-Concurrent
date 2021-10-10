package functionblock

import "IEC-61499-Concurrent/device"

type EArm struct {
	FbInfo
}

func (nowFb *EArm) Exectue(car *device.CarModel, eventIn string) {
	if eventIn == "" {
		panic("empty event input")
	}
	nowFb.DeviceMapping.(*device.Arm).ArmMove(car, CycleTime, "X", PositiveDirection)
}

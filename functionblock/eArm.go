package functionblock

import "IEC-61499-Concurrent/device"

type EArm struct {
	FbInfo
}

func (nowFb *EArm) Execute(car *device.CarModel, eventIn string) {
	nowFb.DeviceMapping.(*device.Arm).ArmMove(car, CycleTime, "X", PositiveDirection)
}

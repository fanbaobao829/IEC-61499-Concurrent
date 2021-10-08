package functionblock

import "IEC-61499-Concurrent/device"

type EArm struct {
	FbInfo
}

func (nowFb *EArm) Execute(eventIn string) {
	nowFb.DeviceMapping.(*device.Arm).ArmMove(device.GlobalCarModel, CycleTime, "X", PositiveDirection)

}

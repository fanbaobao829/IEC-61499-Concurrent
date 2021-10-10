package functionblock

import "IEC-61499-Concurrent/device"

type ESensor struct {
	FbInfo
}

func (nowFb *ESensor) Execute(car *device.CarModel, eventIn string) {
	if eventIn == "" {
		panic("empty event input")
	}
	nowFb.DeviceMapping.(*device.Sensor).Execute(car)
}

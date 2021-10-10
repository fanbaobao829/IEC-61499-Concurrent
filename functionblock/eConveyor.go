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
	if eventIn == "" {
		panic("empty event input")
	}
	nowFb.DeviceMapping.(*device.Belt).BeltMove(car, CycleTime, PositiveDirection)
}

func (nowFb *EConveyor) DeviceMap(device interface{}) {
	nowFb.DeviceMapping = device
}

func (nowFb *EConveyor) EventMap(fb Fb) {
	for _, inputEvent := range nowFb.EventIn {
		EventMap[inputEvent.Name] = fb
	}
	for _, outputEvent := range nowFb.EventOut {
		EventMap[outputEvent.Name] = fb
	}
	for _, inputData := range nowFb.DataIn {
		DataMap[inputData.Name] = fb
	}
	for _, outputData := range nowFb.DataOut {
		DataMap[outputData.Name] = fb
	}
}

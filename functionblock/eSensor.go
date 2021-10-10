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

func (nowFb *ESensor) DeviceMap(device interface{}) {
	nowFb.DeviceMapping = device
}

func (nowFb *ESensor) EventMap(fb Fb) {
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

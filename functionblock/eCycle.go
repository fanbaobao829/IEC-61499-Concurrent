package functionblock

import (
	"IEC-61499-Concurrent/communication"
	"IEC-61499-Concurrent/communication/channel"
	"IEC-61499-Concurrent/device"
	"IEC-61499-Concurrent/event"
	"time"
)

type ECycle struct {
	FbInfo
}

func (nowFb *ECycle) Execute(car *device.CarModel, eventIn string) {
	if eventIn == "" || car == nil {
		panic("empty event input")
	}
	for {
		select {
		case <-channel.GlobalExitChannel:
			return
		default:
			for _, eventOut := range nowFb.EventOut {
				go communication.GlobalEventBus.Publish(eventOut.Name, event.DiscreteEvent{Name: eventOut.Name, Tlast: time.Now().UnixNano(), Tddl: time.Now().UnixNano() + CycleTime, Priority: BasePriority})
			}
		}
		time.Sleep(time.Duration(ScanCycle) * time.Nanosecond)
	}
}

func (nowFb *ECycle) DeviceMap(device interface{}) {
	nowFb.DeviceMapping = device
}

func (nowFb *ECycle) EventMap(fb Fb) {
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

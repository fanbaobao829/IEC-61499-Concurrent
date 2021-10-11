package functionblock

import (
	"IEC-61499-Concurrent/communication"
	"IEC-61499-Concurrent/device"
	"IEC-61499-Concurrent/event"
	"time"
)

type EMerge struct {
	FbInfo
}

func (nowFb *EMerge) Execute(car *device.CarModel, eventIn string) {
	if car == nil {
		panic("empty car model")
	}
	nowFbPrivate := nowFb.FbPrivate.(EMergeAndServiceValue)
	println(233)
	for eventInIndex, eventInInterface := range nowFb.EventOut {
		if eventIn == eventInInterface.Name {
			if nowFbPrivate.FbLast+nowFbPrivate.FbTtl < time.Now().UnixNano() {
				go error(nowFb)
			}
			nowFbPrivate.FbCache |= 1 << eventInIndex
			if nowFbPrivate.FbCache >= nowFbPrivate.FbThreshold {
				for _, eventOut := range nowFb.EventOut {
					go communication.GlobalEventBus.Publish(eventOut.Name, event.DiscreteEvent{Name: eventOut.Name, Tlast: time.Now().UnixNano(), Tddl: time.Now().UnixNano() + 3*1e9, Priority: BasePriority})
					//data refresh
				}
				nowFbPrivate.FbCache = 0
			}
			nowFbPrivate.FbLast = time.Now().UnixNano()
			nowFb.FbPrivate = nowFbPrivate
		}
	}
}

func (nowFb *EMerge) DeviceMap(device interface{}) {
	nowFb.DeviceMapping = device
}

func (nowFb *EMerge) EventMap(fb Fb) {
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

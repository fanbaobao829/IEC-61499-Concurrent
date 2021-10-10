package functionblock

import (
	"IEC-61499-Concurrent/communication"
	"IEC-61499-Concurrent/device"
	"IEC-61499-Concurrent/event"
	"time"
)

type EStart struct {
	FbInfo
}

func (nowFb *EStart) Exectue(car *device.CarModel, eventIn string) {
	if eventIn == "" || car == nil {
		panic("empty event input")
	}
	for _, eventOut := range nowFb.EventOut {
		go communication.GlobalEventBus.Publish(eventOut.Name, event.DiscreteEvent{Name: eventOut.Name, Tlast: time.Now().UnixNano(), Tddl: time.Now().UnixNano() + CycleTime, Priority: BasePriority})
		//data refresh
	}
}

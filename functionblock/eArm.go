package functionblock

import (
	"IEC-61499-Concurrent/communication"
	"IEC-61499-Concurrent/device"
	"IEC-61499-Concurrent/event"
	"strings"
	"time"
)

type EArm struct {
	FbInfo
}

func (nowFb *EArm) Exectue(car *device.CarModel, eventIn string) {
	if eventIn == "" {
		panic("empty event input")
	}
	if strings.Contains(eventIn, "arm_in") {
		//refresh data
		for _, eventOut := range nowFb.EventOut {
			for _, dataOut := range nowFb.NameToInterface[eventOut.Name].(FbOutputEventInterface).DataLink {
				nowFb.NameToInterface[dataOut].(*FbOutputDataInterface).Value = true
			}
		}
	}
	if strings.Contains(eventIn, "arm_cycle") {
		//check pos
		if nowFb.DeviceMapping.(*device.Arm).AxisXoY.Angular >= 120 {
			for _, eventOut := range nowFb.EventOut {
				go communication.GlobalEventBus.Publish(eventOut.Name, event.DiscreteEvent{Name: eventOut.Name, Tlast: time.Now().UnixNano(), Tddl: time.Now().UnixNano() + CycleTime, Priority: BasePriority})
			}
		}
	}
}

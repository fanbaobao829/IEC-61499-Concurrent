package functionblock

import (
	"IEC-61499-Concurrent/communication"
	"IEC-61499-Concurrent/device"
	"IEC-61499-Concurrent/event"
	"time"
)

type EArm struct {
	FbInfo
}

func (nowFb *EArm) Exectue(car *device.CarModel, eventIn string) {
	if eventIn == "" {
		panic("empty event input")
	}
	nowFb.DeviceMapping.(*device.Arm).ArmMove(car, CycleTime, "X", PositiveDirection)
	if car.NowPos == car.Destination {
		for _, eventOut := range nowFb.EventOut {
			go communication.GlobalEventBus.Publish(eventOut.Name, event.DiscreteEvent{Name: eventOut.Name, Tlast: time.Now().UnixNano(), Tddl: time.Now().UnixNano() + CycleTime, Priority: BasePriority})
		}
	}
}

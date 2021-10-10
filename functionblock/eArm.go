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

var ch chan bool

func (nowFb *EArm) Execute(car *device.CarModel, eventIn string) {
	if eventIn == "" {
		panic("empty event input")
	}
	if strings.Contains(eventIn, "arm_in") {
		//refresh data
		for _, eventOut := range nowFb.EventOut {
			for _, dataOut := range nowFb.NameToInterface[eventOut.Name].(FbOutputEventInterface).DataLink {
				nowFb.NameToInterface[dataOut].(*FbOutputDataInterface).Value = true
				go func() {
					for {
						select {
						case <-ch:
							return
						default:
							nowFb.DeviceMapping.(*device.Arm).ArmMove(car, CycleTime, "XoY", PositiveDirection)
							time.Sleep(CycleTime * time.Nanosecond)
						}
					}
				}()
			}
		}
	}
	if strings.Contains(eventIn, "arm_cycle") {
		//check pos
		if nowFb.DeviceMapping.(*device.Arm).AxisXoY.Angular >= 120 {
			ch <- true
			for _, eventOut := range nowFb.EventOut {
				go communication.GlobalEventBus.Publish(eventOut.Name, event.DiscreteEvent{Name: eventOut.Name, Tlast: time.Now().UnixNano(), Tddl: time.Now().UnixNano() + CycleTime, Priority: BasePriority})
			}
		}
	}
}

func (nowFb *EArm) DeviceMap(device interface{}) {
	nowFb.DeviceMapping = device
}

func (nowFb *EArm) EventMap(fb Fb) {
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

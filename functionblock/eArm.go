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

func (nowFb *EArm) Execute(car *device.CarModel, eventIn string) {
	if eventIn == "" {
		panic("empty event input")
	}
	if strings.Contains(eventIn, "arm_in") {
		//refresh data
		for _, eventOut := range nowFb.EventOut {
			for _, dataOut := range nowFb.NameToInterface[eventOut.Name].(*FbOutputEventInterface).DataLink {
				nowFb.NameToInterface[dataOut].(*FbOutputDataInterface).Value = true
				nowFb.FbPrivate.(*EArmServiceValue).FbActiveTimeStamp = time.Now().UnixNano()
				nowFb.FbPrivate.(*EArmServiceValue).FbLastTimeStamp = time.Now().UnixNano()
			}
		}
	}
	if strings.Contains(eventIn, "arm_cycle") {
		//check pos
		if nowFb.DataOut[0].Value != nil && nowFb.DataOut[0].Value.(bool) {
			if RunMode == "serial" {
				nowFb.DeviceMapping.(*device.Arm).ArmMove(car, CycleTime, "XoY", PositiveDirection)
				time.Sleep(time.Duration(CycleTime) * time.Nanosecond)
			} else {
				nowFb.DeviceMapping.(*device.Arm).ArmMove(car, time.Now().UnixNano()-nowFb.FbPrivate.(*EArmServiceValue).FbLastTimeStamp, "XoY", PositiveDirection)
			}
			nowFb.FbPrivate.(*EArmServiceValue).FbLastTimeStamp = time.Now().UnixNano()
		} else {
			return
		}
		if nowFb.DeviceMapping.(*device.Arm).AxisXoY.Angular >= 120 {
			nowFb.DataOut[0].Value = false
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

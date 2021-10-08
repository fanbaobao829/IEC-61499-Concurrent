package functionblock

import (
	"IEC-61499-Concurrent/communication"
	"IEC-61499-Concurrent/event"
	"time"
)

type Fb interface {
	Exectue()
}

type ESplit struct {
	FbInfo
}

const (
	CycleTime    = 20000000
	BasePriority = 1
)

func (nowFb *ESplit) Execute(eventIn string) {
	for _, eventOut := range nowFb.EventOut {
		go communication.GlobalEventBus.Publish(eventOut.Name, event.DiscreteEvent{Name: eventOut.Name, Tlast: time.Now().UnixNano(), Tddl: time.Now().UnixNano() + CycleTime, Priority: BasePriority})
		//data refresh
	}
}

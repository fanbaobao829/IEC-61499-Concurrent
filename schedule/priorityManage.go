package schedule

import (
	"IEC-61499-Concurrent/communication/channel"
	"IEC-61499-Concurrent/device"
	"IEC-61499-Concurrent/event"
	"IEC-61499-Concurrent/functionblock"
	"IEC-61499-Concurrent/schedule/skiplist"
	"time"
)

func init() {
	go func() {
		for {
			select {
			case <-channel.GlobalExitChannel:
				return
			default:
				AdjustPriority(skiplist.GlobalEventQueue)
				ActiveFunctionBlock(skiplist.GlobalEventQueue)
				time.Sleep(5 * time.Millisecond)
			}
		}
	}()
}
func AdjustPriority(list *skiplist.EventQueue) {
	//上锁
	list.Rm.Lock()
	newList := skiplist.PriorityQueue()
	for !list.Queue.Empty() {
		newList.Push(adjustPriority(list.Queue.Top()))
		list.Queue.Pop()
	}
	list.Queue = newList
	//解锁
	list.Rm.Unlock()
}

func adjustPriority(top *event.DiscreteEvent) event.DiscreteEvent {
	top.Priority -= int(functionblock.BasePriority - (top.Tddl-time.Now().UnixNano())/(top.Tddl-top.Tlast)*4)
	return *top
}

func ActiveFunctionBlock(list *skiplist.EventQueue) {
	list.Rm.Lock()
	for {
		if list.Queue.Empty() {
			break
		}
		nowEvent := list.Queue.Top()
		go functionblock.EventMap[nowEvent.Name].Execute(device.GlobalCarModel, nowEvent.Name)
		list.Queue.Pop()
	}
	list.Rm.Unlock()
}

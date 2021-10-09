package schedule

import (
	"IEC-61499-Concurrent/event"
	"IEC-61499-Concurrent/functionblock"
	"IEC-61499-Concurrent/schedule/skiplist"
	"time"
)

func init() {
	go func() {
		for {
			AdjustPriority(skiplist.EventQueue)
			time.Sleep(5 * time.Millisecond)
		}
	}()
}
func AdjustPriority(list *skiplist.SkipList) *skiplist.SkipList {
	//上锁
	newList := skiplist.PriorityQueue()
	for !list.Empty() {
		newList.Push(adjustPriority(list.Top()))
		list.Pop()
	}
	return newList
	//解锁
}

func adjustPriority(top *event.DiscreteEvent) event.DiscreteEvent {
	top.Priority -= int(functionblock.BasePriority - (top.Tddl-top.Tlast)/5)
	return *top
}

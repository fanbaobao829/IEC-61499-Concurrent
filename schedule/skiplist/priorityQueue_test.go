package skiplist

import (
	"IEC-61499-Concurrent/event"
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	list := PriorityQueue()
	list.Push(event.DiscreteEvent{Tlast: 5, Tddl: 5, Priority: 5})
	list.Push(event.DiscreteEvent{Tlast: 3, Tddl: 9, Priority: 1})
	list.Push(event.DiscreteEvent{Tlast: 1, Tddl: 7, Priority: 3})
	list.Push(event.DiscreteEvent{Tlast: 2, Tddl: 4, Priority: 1})
	list.Push(event.DiscreteEvent{Tlast: 5, Tddl: 6, Priority: 4})
	list.Push(event.DiscreteEvent{Tlast: 5, Tddl: 7, Priority: 2})

	printPriorityQueue(list)
	fmt.Println(list.Top())
	list.Pop()

	printPriorityQueue(list)

}

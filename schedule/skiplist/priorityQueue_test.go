package skiplist

import (
	"IEC-61499-Concurrent/event"
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	list := PriorityQueue()
	list.Push(event.DiscreteEvent{Priority: 5, Tlast: 1633955556384319000, Tddl: 1633955556404319000, Name: "arm_cycle1"})
	list.Push(event.DiscreteEvent{Priority: 5, Tlast: 1633955556384319000, Tddl: 1633955556404319000, Name: "arm_cycle1"})
	list.Push(event.DiscreteEvent{Tlast: 1, Tddl: 7, Priority: 3})
	list.Push(event.DiscreteEvent{Tlast: 2, Tddl: 4, Priority: 1})
	list.Push(event.DiscreteEvent{Tlast: 5, Tddl: 6, Priority: 4})
	list.Push(event.DiscreteEvent{Tlast: 5, Tddl: 7, Priority: 2})

	PrintPriorityQueue(list)
	fmt.Println(list.Top())
	list.Pop()
	fmt.Println(list.Top())
	list.Pop()
	fmt.Println(list.Top())
	list.Pop()
	fmt.Println(list.Top())
	list.Pop()
	fmt.Println(list.Top())
	list.Pop()
	fmt.Println(list.Top())
	list.Pop()

	PrintPriorityQueue(list)
	println(1633958720407969000 - 1633958717415162000)
}

package event

import "testing"

func TestEqual(t *testing.T) {
	println(Equal(DiscreteEvent{Name: "", Tlast: 1, Tddl: 1, Priority: 1}, DiscreteEvent{Name: "", Tlast: 1, Tddl: 1, Priority: 1}))
}

func TestGreat(t *testing.T) {
	println(Greater(DiscreteEvent{Name: "", Tlast: 1, Tddl: 1, Priority: 1}, DiscreteEvent{Name: "", Tlast: 1, Tddl: 1, Priority: 1}))
}

func TestSmall(t *testing.T) {
	println(Smaller(DiscreteEvent{Name: "", Tlast: 1, Tddl: 1, Priority: 1}, DiscreteEvent{Name: "", Tlast: 1, Tddl: 1, Priority: 1}))
}

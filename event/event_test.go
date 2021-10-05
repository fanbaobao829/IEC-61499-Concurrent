package event

import "testing"

func TestEqual(t *testing.T) {
	println(Equal(DiscreteEvent{1, 1, 1}, DiscreteEvent{1, 1, 1}))
}

func TestGreat(t *testing.T) {
	println(Greater(DiscreteEvent{1, 1, 1}, DiscreteEvent{1, 1, 1}))
}

func TestSmall(t *testing.T) {
	println(Smaller(DiscreteEvent{1, 1, 1}, DiscreteEvent{1, 1, 1}))
}

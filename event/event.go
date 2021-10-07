package event

type DiscreteEvent struct {
	Name     string
	Tlast    int64
	Tddl     int64
	Priority int
}

func Greater(event1 DiscreteEvent, event2 DiscreteEvent) bool {
	if event1.Priority == event2.Priority {
		if event1.Tlast == event2.Tlast {
			return event1.Tddl < event2.Tddl
		}
		return event1.Tlast < event2.Tlast
	}
	return event1.Priority < event2.Priority
}

func Smaller(event1 DiscreteEvent, event2 DiscreteEvent) bool {
	if event1.Priority == event2.Priority {
		if event1.Tlast == event2.Tlast {
			return event1.Tddl > event2.Tddl
		}
		return event1.Tlast > event2.Tlast
	}
	return event1.Priority > event2.Priority
}

func Equal(event1 DiscreteEvent, event2 DiscreteEvent) bool {
	return event1.Priority == event2.Priority && event1.Tlast == event2.Tlast && event1.Tddl == event2.Tddl
}

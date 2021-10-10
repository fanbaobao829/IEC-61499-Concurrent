package device

type Sensor struct {
	Value     float64
	ScopeMax  float64
	ScopeMin  float64
	Direction string
	Active    bool
}

func (sensor *Sensor) Execute(car *CarModel) {
	if sensor.Direction == "X" {
		if car.NowPos.PosX >= sensor.ScopeMin && car.NowPos.PosX <= sensor.ScopeMax {
			sensor.Active = true
		}
	} else if sensor.Direction == "Y" {
		if car.NowPos.PosY >= sensor.ScopeMin && car.NowPos.PosY <= sensor.ScopeMax {
			sensor.Active = true
		}
	} else {
		if car.NowPos.PosZ >= sensor.ScopeMin && car.NowPos.PosZ <= sensor.ScopeMax {
			sensor.Active = true
		}
	}
}

package device

type Sensor struct {
	Value     float64
	ScopeMax  float64
	ScopeMin  float64
	Direction string
}

func (sensor *Sensor) Execute(car *CarModel) bool {
	if sensor.Direction == "X" {
		if car.NowPos.PosX >= sensor.ScopeMin && car.NowPos.PosX <= sensor.ScopeMax {
			return true
		}
	} else if sensor.Direction == "Y" {
		if car.NowPos.PosY >= sensor.ScopeMin && car.NowPos.PosY <= sensor.ScopeMax {
			return true
		}
	} else {
		if car.NowPos.PosZ >= sensor.ScopeMin && car.NowPos.PosZ <= sensor.ScopeMax {
			return true
		}
	}
	return false
}

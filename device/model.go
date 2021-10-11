package device

type Position struct {
	PosX float64
	PosY float64
	PosZ float64
}

type CarModel struct {
	NowPos      Position
	Destination Position
}

var GlobalCarModel *CarModel

func init() {
	GlobalCarModel = new(CarModel)
}

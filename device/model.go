package device

type Position struct {
	PosX float64
	PosY float64
	PosZ float64
}

type carModel struct {
	NowPos      Position
	Destination Position
}

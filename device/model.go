package device

type Position struct {
	PosX float32
	PosY float32
	PosZ float32
}

type carModel struct {
	NowPos      Position
	Destination Position
}

func GetNowPos() {

}

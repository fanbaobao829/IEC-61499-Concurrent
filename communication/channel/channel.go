package channel

var GlobalExitChannel chan bool

func init() {
	GlobalExitChannel = make(chan bool)
}

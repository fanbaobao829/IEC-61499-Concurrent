package device

func init() {

}

type Axis struct {
	Speed      float32
	Angular    float32
	length     float32
	MaxAngular float32
	MinAngular float32
	Direction  string
}

type Arm struct {
	ActuatorPos Position
	AxisX       Axis
	AxisY       Axis
	AxisZ       Axis
	Height      float32
}

func StepArmMoveTo(car carModel) {

}

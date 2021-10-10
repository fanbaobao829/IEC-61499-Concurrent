package device

import (
	"math"
)

type Axis struct {
	Speed      float64
	Angular    float64
	Length     float64
	MaxAngular float64
	MinAngular float64
}

type Arm struct {
	ActuatorPos Position
	AxisXoY     Axis
	AxisYoZ     Axis
	AxisXoZ     Axis
	BasePos     Position
}

func (arm *Arm) ArmMove(car *CarModel, timeDuration int, axis string, direction int) {
	if axis == "XoY" {
		if direction > 0 {
			arm.AxisXoY.Angular = math.Min(arm.AxisXoY.MaxAngular, arm.AxisXoY.Angular+arm.AxisXoY.Speed*float64(timeDuration/1e9))
		} else {
			arm.AxisXoY.Angular = math.Max(arm.AxisXoY.MinAngular, arm.AxisXoY.Angular-arm.AxisXoY.Speed*float64(timeDuration/1e9))
		}
	} else if axis == "XoZ" {
		if direction > 0 {
			arm.AxisXoZ.Angular = math.Min(arm.AxisXoZ.MaxAngular, arm.AxisXoZ.Angular+arm.AxisXoZ.Speed*float64(timeDuration/1e9))
		} else {
			arm.AxisXoZ.Angular = math.Max(arm.AxisXoZ.MinAngular, arm.AxisXoZ.Angular-arm.AxisXoZ.Speed*float64(timeDuration/1e9))
		}
	} else {
		if direction > 0 {
			arm.AxisYoZ.Angular = math.Min(arm.AxisYoZ.MaxAngular, arm.AxisYoZ.Angular+arm.AxisYoZ.Speed*float64(timeDuration/1e9))
		} else {
			arm.AxisYoZ.Angular = math.Max(arm.AxisYoZ.MinAngular, arm.AxisYoZ.Angular-arm.AxisYoZ.Speed*float64(timeDuration/1e9))
		}
	}
	arm.ActuatorPos.PosX = arm.BasePos.PosX + arm.AxisXoY.Length*math.Cos(arm.AxisXoY.Angular) + arm.AxisXoZ.Length*math.Cos(arm.AxisXoZ.Angular)
	arm.ActuatorPos.PosY = arm.BasePos.PosY + arm.AxisXoY.Length*math.Sin(arm.AxisXoY.Angular) + arm.AxisYoZ.Length*math.Cos(arm.AxisYoZ.Angular)
	arm.ActuatorPos.PosZ = arm.BasePos.PosZ + arm.AxisXoZ.Length*math.Sin(arm.AxisXoZ.Angular) + arm.AxisYoZ.Length*math.Sin(arm.AxisYoZ.Angular)
	car.NowPos = arm.ActuatorPos
}

//新建和监听？

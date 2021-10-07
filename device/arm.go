package device

import (
	"math"
	"time"
)

type Axis struct {
	Speed      float64
	Angular    float64
	length     float64
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

func (arm *Arm) ArmMove(car *carModel, timeDuration time.Duration, axis string, direction int) {
	if axis == "XoY" {
		if direction > 0 {
			arm.AxisXoY.Angular = math.Min(arm.AxisXoY.MaxAngular, arm.AxisXoY.Angular+arm.AxisXoY.Speed*float64(timeDuration/time.Second))
		} else {
			arm.AxisXoY.Angular = math.Max(arm.AxisXoY.MinAngular, arm.AxisXoY.Angular-arm.AxisXoY.Speed*float64(timeDuration/time.Second))
		}
	} else if axis == "XoZ" {
		if direction > 0 {
			arm.AxisXoZ.Angular = math.Min(arm.AxisXoZ.MaxAngular, arm.AxisXoZ.Angular+arm.AxisXoZ.Speed*float64(timeDuration/time.Second))
		} else {
			arm.AxisXoZ.Angular = math.Max(arm.AxisXoZ.MinAngular, arm.AxisXoZ.Angular-arm.AxisXoZ.Speed*float64(timeDuration/time.Second))
		}
	} else {
		if direction > 0 {
			arm.AxisYoZ.Angular = math.Min(arm.AxisYoZ.MaxAngular, arm.AxisYoZ.Angular+arm.AxisYoZ.Speed*float64(timeDuration/time.Second))
		} else {
			arm.AxisYoZ.Angular = math.Max(arm.AxisYoZ.MinAngular, arm.AxisYoZ.Angular-arm.AxisYoZ.Speed*float64(timeDuration/time.Second))
		}
	}
	arm.ActuatorPos.PosX = arm.BasePos.PosX + arm.AxisXoY.length*math.Cos(arm.AxisXoY.Angular) + arm.AxisXoZ.length*math.Cos(arm.AxisXoZ.Angular)
	arm.ActuatorPos.PosY = arm.BasePos.PosY + arm.AxisXoY.length*math.Sin(arm.AxisXoY.Angular) + arm.AxisYoZ.length*math.Cos(arm.AxisYoZ.Angular)
	arm.ActuatorPos.PosZ = arm.BasePos.PosZ + arm.AxisXoZ.length*math.Sin(arm.AxisXoZ.Angular) + arm.AxisYoZ.length*math.Sin(arm.AxisYoZ.Angular)
	car.NowPos = arm.ActuatorPos
}

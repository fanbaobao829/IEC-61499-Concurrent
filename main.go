package main

import (
	"IEC-61499-Concurrent/communication"
	"IEC-61499-Concurrent/device"
	"IEC-61499-Concurrent/functionblock"
)

var (
	E_Split   functionblock.Fb
	E_Merge   functionblock.Fb
	E_Arm     [6]functionblock.Fb
	E_Start   functionblock.Fb
	E_Cycle   [6]functionblock.Fb
	D_Arm     [6]*device.Arm
	Car_Model *device.CarModel
)

func init() {
	//创建功能块
	E_Split = &functionblock.ESplit{*functionblock.AddFb("split", nil, []string{"split_in"}, []string{"split_out1", "split_out2", "split_out3", "split_out4", "split_out5", "split_out6"}, []string{}, []string{})}
	E_Merge = &functionblock.EMerge{*functionblock.AddFb("merge", functionblock.EMergeAndServiceValue{FbThreshold: 1 << 6, FbTtl: functionblock.CycleTime}, []string{"merge_in1", "merge_in2", "merge_in3", "merge_in4", "merge_in5", "merge_in6"}, []string{"merge_out"}, []string{}, []string{})}
	//功能块内事件驱动
	E_Arm[0] = &functionblock.EArm{*functionblock.AddFb("arm1", nil, []string{"arm_in1", "arm_cycle1"}, []string{"arm_out1"}, []string{}, []string{"arm_execute1"}).AddFbOutputEventDataLink("arm_out1", []string{"arm_execute1"})}
	E_Arm[1] = &functionblock.EArm{*functionblock.AddFb("arm2", nil, []string{"arm_in2", "arm_cycle2"}, []string{"arm_out2"}, []string{}, []string{"arm_execute2"}).AddFbOutputEventDataLink("arm_out2", []string{"arm_execute2"})}
	E_Arm[2] = &functionblock.EArm{*functionblock.AddFb("arm3", nil, []string{"arm_in3", "arm_cycle3"}, []string{"arm_out3"}, []string{}, []string{"arm_execute3"}).AddFbOutputEventDataLink("arm_out3", []string{"arm_execute3"})}
	E_Arm[3] = &functionblock.EArm{*functionblock.AddFb("arm4", nil, []string{"arm_in4", "arm_cycle4"}, []string{"arm_out4"}, []string{}, []string{"arm_execute4"}).AddFbOutputEventDataLink("arm_out4", []string{"arm_execute4"})}
	E_Arm[4] = &functionblock.EArm{*functionblock.AddFb("arm5", nil, []string{"arm_in5", "arm_cycle5"}, []string{"arm_out5"}, []string{}, []string{"arm_execute5"}).AddFbOutputEventDataLink("arm_out5", []string{"arm_execute5"})}
	E_Arm[5] = &functionblock.EArm{*functionblock.AddFb("arm6", nil, []string{"arm_in6", "arm_cycle6"}, []string{"arm_out6"}, []string{}, []string{"arm_execute6"}).AddFbOutputEventDataLink("arm_out6", []string{"arm_execute6"})}
	E_Start = &functionblock.EStart{*functionblock.AddFb("start", nil, []string{}, []string{"start_out"}, []string{}, []string{})}
	E_Cycle[0] = &functionblock.ECycle{*functionblock.AddFb("cycle1", nil, []string{}, []string{"cycle_out1"}, []string{}, []string{})}
	E_Cycle[1] = &functionblock.ECycle{*functionblock.AddFb("cycle2", nil, []string{}, []string{"cycle_out2"}, []string{}, []string{})}
	E_Cycle[2] = &functionblock.ECycle{*functionblock.AddFb("cycle3", nil, []string{}, []string{"cycle_out3"}, []string{}, []string{})}
	E_Cycle[3] = &functionblock.ECycle{*functionblock.AddFb("cycle4", nil, []string{}, []string{"cycle_out4"}, []string{}, []string{})}
	E_Cycle[4] = &functionblock.ECycle{*functionblock.AddFb("cycle5", nil, []string{}, []string{"cycle_out5"}, []string{}, []string{})}
	E_Cycle[5] = &functionblock.ECycle{*functionblock.AddFb("cycle6", nil, []string{}, []string{"cycle_out6"}, []string{}, []string{})}
	//功能块链接
	communication.AddFbEventLink("start_out", "split")
	communication.AddFbEventLink("split_out1", "arm_in1")
	communication.AddFbEventLink("split_out2", "arm_in2")
	communication.AddFbEventLink("split_out3", "arm_in3")
	communication.AddFbEventLink("split_out4", "arm_in4")
	communication.AddFbEventLink("split_out5", "arm_in5")
	communication.AddFbEventLink("split_out6", "arm_in6")
	communication.AddFbEventLink("cycle_out1", "arm_cycle1")
	communication.AddFbEventLink("cycle_out2", "arm_cycle2")
	communication.AddFbEventLink("cycle_out3", "arm_cycle3")
	communication.AddFbEventLink("cycle_out4", "arm_cycle4")
	communication.AddFbEventLink("cycle_out5", "arm_cycle5")
	communication.AddFbEventLink("cycle_out6", "arm_cycle6")
	communication.AddFbEventLink("arm_out1", "merge_in1")
	communication.AddFbEventLink("arm_out2", "merge_in2")
	communication.AddFbEventLink("arm_out3", "merge_in3")
	communication.AddFbEventLink("arm_out4", "merge_in4")
	communication.AddFbEventLink("arm_out5", "merge_in5")
	communication.AddFbEventLink("arm_out6", "merge_in6")
	//创建设备
	D_Arm[0] = &device.Arm{ActuatorPos: device.Position{PosX: 0, PosY: 0, PosZ: 0}, AxisXoY: device.Axis{Angular: 90, Speed: 10, Length: 10, MaxAngular: 180, MinAngular: 0}, AxisXoZ: device.Axis{Angular: 90, Speed: 10, Length: 10, MaxAngular: 180, MinAngular: 0}, AxisYoZ: device.Axis{Angular: 90, Speed: 10, Length: 10, MaxAngular: 180, MinAngular: 0}, BasePos: device.Position{PosX: 0, PosY: 0, PosZ: 0}}
	D_Arm[1] = &device.Arm{ActuatorPos: device.Position{PosX: 0, PosY: 0, PosZ: 0}, AxisXoY: device.Axis{Angular: 90, Speed: 10, Length: 10, MaxAngular: 180, MinAngular: 0}, AxisXoZ: device.Axis{Angular: 90, Speed: 10, Length: 10, MaxAngular: 180, MinAngular: 0}, AxisYoZ: device.Axis{Angular: 90, Speed: 10, Length: 10, MaxAngular: 180, MinAngular: 0}, BasePos: device.Position{PosX: 0, PosY: 0, PosZ: 0}}
	D_Arm[2] = &device.Arm{ActuatorPos: device.Position{PosX: 0, PosY: 0, PosZ: 0}, AxisXoY: device.Axis{Angular: 90, Speed: 10, Length: 10, MaxAngular: 180, MinAngular: 0}, AxisXoZ: device.Axis{Angular: 90, Speed: 10, Length: 10, MaxAngular: 180, MinAngular: 0}, AxisYoZ: device.Axis{Angular: 90, Speed: 10, Length: 10, MaxAngular: 180, MinAngular: 0}, BasePos: device.Position{PosX: 0, PosY: 0, PosZ: 0}}
	D_Arm[3] = &device.Arm{ActuatorPos: device.Position{PosX: 0, PosY: 0, PosZ: 0}, AxisXoY: device.Axis{Angular: 90, Speed: 10, Length: 10, MaxAngular: 180, MinAngular: 0}, AxisXoZ: device.Axis{Angular: 90, Speed: 10, Length: 10, MaxAngular: 180, MinAngular: 0}, AxisYoZ: device.Axis{Angular: 90, Speed: 10, Length: 10, MaxAngular: 180, MinAngular: 0}, BasePos: device.Position{PosX: 0, PosY: 0, PosZ: 0}}
	D_Arm[4] = &device.Arm{ActuatorPos: device.Position{PosX: 0, PosY: 0, PosZ: 0}, AxisXoY: device.Axis{Angular: 90, Speed: 10, Length: 10, MaxAngular: 180, MinAngular: 0}, AxisXoZ: device.Axis{Angular: 90, Speed: 10, Length: 10, MaxAngular: 180, MinAngular: 0}, AxisYoZ: device.Axis{Angular: 90, Speed: 10, Length: 10, MaxAngular: 180, MinAngular: 0}, BasePos: device.Position{PosX: 0, PosY: 0, PosZ: 0}}
	D_Arm[5] = &device.Arm{ActuatorPos: device.Position{PosX: 0, PosY: 0, PosZ: 0}, AxisXoY: device.Axis{Angular: 90, Speed: 10, Length: 10, MaxAngular: 180, MinAngular: 0}, AxisXoZ: device.Axis{Angular: 90, Speed: 10, Length: 10, MaxAngular: 180, MinAngular: 0}, AxisYoZ: device.Axis{Angular: 90, Speed: 10, Length: 10, MaxAngular: 180, MinAngular: 0}, BasePos: device.Position{PosX: 0, PosY: 0, PosZ: 0}}
	//设备与功能块映射
	functionblock.AddMappingFbToDevice(E_Arm[0], D_Arm[0])
	functionblock.AddMappingFbToDevice(E_Arm[1], D_Arm[1])
	functionblock.AddMappingFbToDevice(E_Arm[2], D_Arm[2])
	functionblock.AddMappingFbToDevice(E_Arm[3], D_Arm[3])
	functionblock.AddMappingFbToDevice(E_Arm[4], D_Arm[4])
	functionblock.AddMappingFbToDevice(E_Arm[5], D_Arm[5])
	//初始化元器件
	Car_Model = &device.CarModel{NowPos: device.Position{PosX: 0, PosY: 0, PosZ: 0}, Destination: device.Position{PosX: 0, PosY: 0, PosZ: 0}}
}

func main() {
	//初始触发
	E_Start.Exectue(Car_Model, "")

}

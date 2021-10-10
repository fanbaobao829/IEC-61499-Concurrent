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
	E_Arm[0] = &functionblock.EArm{*functionblock.AddFb("arm1", nil, []string{"arm1_in"}, []string{"arm1_out"}, []string{}, []string{})}
	E_Arm[1] = &functionblock.EArm{*functionblock.AddFb("arm1", nil, []string{"arm2_in"}, []string{"arm2_out"}, []string{}, []string{})}
	E_Arm[2] = &functionblock.EArm{*functionblock.AddFb("arm1", nil, []string{"arm3_in"}, []string{"arm3_out"}, []string{}, []string{})}
	E_Arm[3] = &functionblock.EArm{*functionblock.AddFb("arm1", nil, []string{"arm4_in"}, []string{"arm4_out"}, []string{}, []string{})}
	E_Arm[4] = &functionblock.EArm{*functionblock.AddFb("arm1", nil, []string{"arm5_in"}, []string{"arm5_out"}, []string{}, []string{})}
	E_Arm[5] = &functionblock.EArm{*functionblock.AddFb("arm1", nil, []string{"arm6_in"}, []string{"arm6_out"}, []string{}, []string{})}
	E_Start = &functionblock.EArm{*functionblock.AddFb("start", nil, []string{}, []string{"cycle_out"}, []string{}, []string{})}
	E_Cycle[0] = &functionblock.EArm{*functionblock.AddFb("cycle1", nil, []string{}, []string{"cycle1_out"}, []string{}, []string{})}
	E_Cycle[1] = &functionblock.EArm{*functionblock.AddFb("cycle1", nil, []string{}, []string{"cycle2_out"}, []string{}, []string{})}
	E_Cycle[2] = &functionblock.EArm{*functionblock.AddFb("cycle1", nil, []string{}, []string{"cycle3_out"}, []string{}, []string{})}
	E_Cycle[3] = &functionblock.EArm{*functionblock.AddFb("cycle1", nil, []string{}, []string{"cycle4_out"}, []string{}, []string{})}
	E_Cycle[4] = &functionblock.EArm{*functionblock.AddFb("cycle1", nil, []string{}, []string{"cycle5_out"}, []string{}, []string{})}
	E_Cycle[5] = &functionblock.EArm{*functionblock.AddFb("cycle1", nil, []string{}, []string{"cycle6_out"}, []string{}, []string{})}
	//功能块链接
	communication.AddFbEventLink("split_out1", "arm1_in")
	communication.AddFbEventLink("split_out2", "arm2_in")
	communication.AddFbEventLink("split_out3", "arm3_in")
	communication.AddFbEventLink("split_out4", "arm4_in")
	communication.AddFbEventLink("split_out5", "arm5_in")
	communication.AddFbEventLink("split_out6", "arm6_in")
	communication.AddFbEventLink("arm1_out", "merge_in1")
	communication.AddFbEventLink("arm2_out", "merge_in2")
	communication.AddFbEventLink("arm3_out", "merge_in3")
	communication.AddFbEventLink("arm4_out", "merge_in4")
	communication.AddFbEventLink("arm5_out", "merge_in5")
	communication.AddFbEventLink("arm6_out", "merge_in6")
	//创建设备
	D_Arm[0] = &device.Arm{}
	D_Arm[1] = &device.Arm{}
	D_Arm[2] = &device.Arm{}
	D_Arm[3] = &device.Arm{}
	D_Arm[4] = &device.Arm{}
	D_Arm[5] = &device.Arm{}
	//设备与功能块映射
	functionblock.AddMappingFbToDevice(E_Arm[0], D_Arm[0])
	functionblock.AddMappingFbToDevice(E_Arm[1], D_Arm[1])
	functionblock.AddMappingFbToDevice(E_Arm[2], D_Arm[2])
	functionblock.AddMappingFbToDevice(E_Arm[3], D_Arm[3])
	functionblock.AddMappingFbToDevice(E_Arm[4], D_Arm[4])
	functionblock.AddMappingFbToDevice(E_Arm[5], D_Arm[5])
	//初始化元器件
	Car_Model = &device.CarModel{}
}

func main() {
	//初始触发
	E_Start.Exectue(Car_Model, "")

}

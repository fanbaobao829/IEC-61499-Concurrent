package functionblock

import (
	"IEC-61499-Concurrent/device"
	"testing"
)

func TestAddMappingFbToDevice(t *testing.T) {
	var fb Fb
	fb = &EArm{*AddFb("split", nil, []string{"split_in"}, []string{"split_out1", "split_out2", "split_out3", "split_out4", "split_out5", "split_out6"}, []string{}, []string{})}
	device := &device.Arm{}
	AddMappingFbToDevice(fb, device)
}

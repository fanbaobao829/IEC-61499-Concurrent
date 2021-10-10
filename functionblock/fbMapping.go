package functionblock

func AddMappingFbToDevice(fb interface{}, device interface{}) {
	switch fb.(type) {
	case *EArm:
		fb.(*EArm).DeviceMapping = device
	case *ESplit:
		fb.(*ESplit).DeviceMapping = device
	case *EMerge:
		fb.(*EMerge).DeviceMapping = device
	case *EConveyor:
		fb.(*EConveyor).DeviceMapping = device
	case *ESensor:
		fb.(*ESensor).DeviceMapping = device
	}
}

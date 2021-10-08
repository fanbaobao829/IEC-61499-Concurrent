package functionblock

import "reflect"

func addMappingFbToDevice(fb interface{}, device interface{}) {
	fbType := reflect.TypeOf(fb)
	fbValue := reflect.ValueOf(fb)
	deviceType := reflect.TypeOf(device)
	deviceValue := reflect.ValueOf(device)

}

package functionblock

import "IEC-61499-Concurrent/device"

type FbInputEventInterface struct {
	Name     string
	DataLink []string
}

type FbOutputEventInterface struct {
	Name     string
	DataLink []string
}

type FbInputDataInterface struct {
	Name  string
	Value interface{}
}

type FbOutputDataInterface struct {
	Name  string
	Value interface{}
}

type FbInfo struct {
	FbName          string
	FbPointer       *FbInfo
	FbPrivate       interface{}
	DeviceMapping   interface{}
	NameToInterface map[string]interface{}
	EventIn         []FbInputEventInterface
	EventOut        []FbOutputEventInterface
	DataIn          []FbInputDataInterface
	DataOut         []FbOutputDataInterface
}

type EMergeAndServiceValue struct {
	FbCache     int
	FbThreshold int
	FbTtl       int64
	FbLast      int64
}
type Fb interface {
	Execute(car *device.CarModel, eventIn string)
	DeviceMap(device interface{})
	EventMap(fb Fb)
}

var EventMap map[string]Fb
var DataMap map[string]Fb

func init() {
	EventMap = make(map[string]Fb)
	DataMap = make(map[string]Fb)
}

func AddFb(name string, privateValue interface{}, inputEventInterface []string, outputEventInterface []string, inputDataInterface []string, outputDataInterface []string) *FbInfo {
	nowFb := FbInfo{FbName: name, FbPrivate: privateValue, NameToInterface: make(map[string]interface{})}
	nowFb.FbPointer = &nowFb
	nowFb.EventIn = make([]FbInputEventInterface, len(inputEventInterface))
	for inputEventIndex, inputEvent := range inputEventInterface {
		nowFb.EventIn[inputEventIndex] = FbInputEventInterface{Name: inputEvent}
		nowFb.NameToInterface[inputEvent] = &nowFb.EventIn[inputEventIndex]
	}
	nowFb.EventOut = make([]FbOutputEventInterface, len(outputEventInterface))
	for outputEventIndex, outputEvent := range outputEventInterface {
		nowFb.EventOut[outputEventIndex] = FbOutputEventInterface{Name: outputEvent}
		nowFb.NameToInterface[outputEvent] = &nowFb.EventOut[outputEventIndex]
	}
	nowFb.DataIn = make([]FbInputDataInterface, len(inputDataInterface))
	for inputDataIndex, inputData := range inputDataInterface {
		nowFb.DataIn[inputDataIndex] = FbInputDataInterface{Name: inputData}
		nowFb.NameToInterface[inputData] = &nowFb.DataIn[inputDataIndex]
	}
	nowFb.DataOut = make([]FbOutputDataInterface, len(outputDataInterface))
	for outputDataIndex, outputData := range outputDataInterface {
		nowFb.DataOut[outputDataIndex] = FbOutputDataInterface{Name: outputData}
		nowFb.NameToInterface[outputData] = &nowFb.DataOut[outputDataIndex]
	}
	return nowFb.FbPointer
}

func (nowFb *FbInfo) AddFbInputEventDataLink(inputEvent string, inputDataInterface []string) *FbInfo {
	nowFbInputEventInterface := nowFb.NameToInterface[inputEvent].(*FbInputEventInterface)
	nowFbInputEventInterface.DataLink = make([]string, len(inputDataInterface))
	for inputDataIndex, inputData := range inputDataInterface {
		nowFbInputEventInterface.DataLink[inputDataIndex] = inputData
	}
	return nowFb
}

func (nowFb *FbInfo) AddFbOutputEventDataLink(outputEvent string, outputDataInterface []string) *FbInfo {
	nowFbOutputEventInterface := nowFb.NameToInterface[outputEvent].(*FbOutputEventInterface)
	nowFbOutputEventInterface.DataLink = make([]string, len(outputDataInterface))
	for outputDataIndex, outputData := range outputDataInterface {
		nowFbOutputEventInterface.DataLink[outputDataIndex] = outputData
	}
	return nowFb
}

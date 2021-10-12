package functionblock

import (
	"IEC-61499-Concurrent/device"
	"gopkg.in/ini.v1"
	"sync"
)

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
	Rm          sync.Mutex
}

type EArmServiceValue struct {
	FbActiveTimeStamp int64
	FbLastTimeStamp   int64
}

type Fb interface {
	Execute(car *device.CarModel, eventIn string)
	DeviceMap(device interface{})
	EventMap(fb Fb)
}

const (
	PositiveDirection = 1
	NegativeDirection = -1
)

var (
	CycleTime    int64
	ScanCycle    int64
	BasePriority int
	RunMode      string
	EventMap     map[string]Fb
	DataMap      map[string]Fb
)

func init() {
	EventMap = make(map[string]Fb)
	DataMap = make(map[string]Fb)
	cfg, _ := ini.Load("./conf/config.ini")
	//创建功能块
	RunMode = cfg.Section("default").Key("mode").String()
	if RunMode == "serial" {
		ScanCycle, _ = cfg.Section("serial").Key("scan_cycle").Int64()
	} else {
		ScanCycle, _ = cfg.Section("concurrency").Key("scan_cycle").Int64()
	}
	CycleTime, _ = cfg.Section("default").Key("cycle_time").Int64()
	BasePriority, _ = cfg.Section("default").Key("base_priority").Int()
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

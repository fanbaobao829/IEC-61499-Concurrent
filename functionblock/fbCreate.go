package functionblock

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
	NameToInterface map[string]interface{}
	EventIn         []FbInputEventInterface
	EventOut        []FbOutputEventInterface
	DataIn          []FbInputDataInterface
	DataOut         []FbOutputDataInterface
}

type EMergeAndServiceValue struct {
	FbCache     int
	FbThreshold int
	FbTtl       int
	FbLast      int
}

var EventMapping map[string]*FbInfo
var DataMapping map[string]*FbInfo

func AddFb(name string, privateValue interface{}, inputEventInterface []string, outputEventInterface []string, inputDataInterface []string, outputDataInterface []string) *FbInfo {
	nowFb := FbInfo{FbName: name, FbPrivate: privateValue}
	nowFb.FbPointer = &nowFb
	nowFb.EventIn = make([]FbInputEventInterface, len(inputEventInterface))
	for inputEventIndex, inputEvent := range inputEventInterface {
		nowFb.EventIn[inputEventIndex] = FbInputEventInterface{Name: inputEvent}
		nowFb.NameToInterface[inputEvent] = &nowFb.EventIn[inputEventIndex]
		EventMapping[inputEvent] = &nowFb
	}
	nowFb.EventOut = make([]FbOutputEventInterface, len(outputEventInterface))
	for outputEventIndex, outputEvent := range outputEventInterface {
		nowFb.EventOut[outputEventIndex] = FbOutputEventInterface{Name: outputEvent}
		nowFb.NameToInterface[outputEvent] = &nowFb.EventOut[outputEventIndex]
		EventMapping[outputEvent] = &nowFb
	}
	nowFb.DataIn = make([]FbInputDataInterface, len(inputDataInterface))
	for inputDataIndex, inputData := range inputDataInterface {
		nowFb.DataIn[inputDataIndex] = FbInputDataInterface{Name: inputData}
		nowFb.NameToInterface[inputData] = &nowFb.DataIn[inputDataIndex]
		DataMapping[inputData] = &nowFb
	}
	nowFb.DataOut = make([]FbOutputDataInterface, len(outputDataInterface))
	for outputDataIndex, outputData := range outputDataInterface {
		nowFb.DataOut[outputDataIndex] = FbOutputDataInterface{Name: outputData}
		nowFb.NameToInterface[outputData] = &nowFb.DataOut[outputDataIndex]
		DataMapping[outputData] = &nowFb
	}
	return nowFb.FbPointer
}

func (nowFb *FbInfo) AddFbInputEventDataLink(inputEvent string, inputDataInterface []string) {
	nowFbInputEventInterface := nowFb.NameToInterface[inputEvent].(FbInputEventInterface)
	nowFbInputEventInterface.DataLink = make([]string, len(inputDataInterface))
	for inputDataIndex, inputData := range inputDataInterface {
		nowFbInputEventInterface.DataLink[inputDataIndex] = inputData
	}
}

func (nowFb *FbInfo) AddFbOutputEventDataLink(outputEvent string, outputDataInterface []string) {
	nowFbOutputEventInterface := nowFb.NameToInterface[outputEvent].(FbOutputEventInterface)
	nowFbOutputEventInterface.DataLink = make([]string, len(outputDataInterface))
	for outputDataIndex, outputData := range outputDataInterface {
		nowFbOutputEventInterface.DataLink[outputDataIndex] = outputData
	}
}

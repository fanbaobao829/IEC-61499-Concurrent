package functionblock

import "IEC-61499-Concurrent/communication"

var EventLinkMapping map[string][]string
var DataLinkMapping map[string][]string

func AddFbEventLink(FbFromEventInterface string, FbToEventInterface string) {
	EventLinkMapping[FbFromEventInterface] = append(EventLinkMapping[FbFromEventInterface], FbToEventInterface)
	communication.GlobalEventBus.Subscribe(FbFromEventInterface, communication.GlobalChannel)
}

func AddFbDataLink(FbFromDataInterface string, FbToDataInterface string) {
	DataLinkMapping[FbFromDataInterface] = append(DataLinkMapping[FbFromDataInterface], FbToDataInterface)
}

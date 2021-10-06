package communication

import (
	"IEC-61499-Concurrent/event"
	"IEC-61499-Concurrent/functionblock"
	"IEC-61499-Concurrent/skiplist"
	"sync"
)

// 定义数据结构
type DataEvent struct {
	Data  interface{}
	Topic string
}

// DataChannel是一个能接收 DataEvent 的 channel
type DataChannel chan DataEvent

// DataChannelSlice 是一个包含 DataChannels 数据的切片
type DataChannelSlice []DataChannel

// 定义事件总线  EventBus 存储有关订阅者感兴趣的特定主题的信息
type EventBus struct {
	Subscribers map[string]DataChannelSlice
	rm          sync.RWMutex
}

//处理订阅消息
func DealDataEvent(data DataEvent) {
	eventdata := data.Data.(event.DiscreteEvent)
	for _, linkedEvent := range functionblock.EventLinkMapping[data.Topic] {
		eventdata.Name = linkedEvent
		skiplist.EventQueue.Push(eventdata)
	}
}

// 订阅主题  如传统方法回调一样。当发布者向主题发布数据时，channel将接收数据。
func (eb *EventBus) Subscribe(topic string, ch DataChannel) {
	eb.rm.Lock()
	if prev, found := eb.Subscribers[topic]; found {
		eb.Subscribers[topic] = append(prev, ch)
	} else {
		eb.Subscribers[topic] = append([]DataChannel{}, ch)
	}
	eb.rm.Unlock()
}

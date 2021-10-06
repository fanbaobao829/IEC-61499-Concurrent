package communication

import "time"

var GlobalEventBus *EventBus
var GlobalChannel DataChannel

func init() {
	// 声明事件总线对象
	GlobalEventBus = &EventBus{
		Subscribers: map[string]DataChannelSlice{},
	}
	GlobalChannel = make(DataChannel)
	go func() {
		for {
			select {
			case d := <-GlobalChannel:
				go DealDataEvent(d)
			}
			time.Sleep(5 * time.Millisecond)
		}
	}()
}

// 发布主题 发布者需要提供广播给订阅者所需要的主题和数据
func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.rm.RLock()
	if chans, found := eb.Subscribers[topic]; found {
		// 这样做是因为切片引用相同的数组，即使它们是按值传递的
		// 因此我们正在使用我们的元素创建一个新切片，从而正确地保持锁定
		channels := append(DataChannelSlice{}, chans...) //切片赋值

		//使用Goroutine 来避免阻塞发布者
		go func(data DataEvent, dataChannelSlices DataChannelSlice) {
			for _, ch := range dataChannelSlices {
				ch <- data
			}
		}(DataEvent{Data: data, Topic: topic}, channels)
	}
	eb.rm.RUnlock()
}

package skiplist

import (
	"IEC-61499-Concurrent/event"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

const MaxLevel = 7

type Node struct {
	Event event.DiscreteEvent
	Prev  *Node // 同层前节点
	Next  *Node // 同层后节点
	Down  *Node // 下层同节点
	Up    *Node // 上层同节点
}

type SkipList struct {
	Level       int
	HeadNodeArr []*Node
}

type EventQueue struct {
	Queue *SkipList
	Rm    sync.Mutex
}

var GlobalEventQueue *EventQueue

func init() {
	GlobalEventQueue = new(EventQueue)
	GlobalEventQueue.Queue = PriorityQueue()
}

// HasNode 是否包含节点
func (list SkipList) HasNode(eventValue event.DiscreteEvent) *Node {
	if list.Level >= 0 {
		// 只有层级在大于等于 0 的时候在进行循环判断，如果层级小于 0 说明是没有任何数据
		level := list.Level
		node := list.HeadNodeArr[level].Next
		for node != nil {
			if event.Equal(node.Event, eventValue) {
				// 如果节点的值 等于 传入的值 就说明包含这个节点 返回 true
				return node
			} else if event.Smaller(node.Event, eventValue) {
				// 如果节点的值大于传入的值，就应该返回上个节点并进入下一层
				if node.Prev.Down == nil {
					if level-1 >= 0 {
						node = list.HeadNodeArr[level-1].Next
					} else {
						node = nil
					}
				} else {
					node = node.Prev.Down
				}
				level -= 1
			} else if event.Greater(node.Event, eventValue) {
				// 如果节点的值小于传入的值就进入下一个节点，如果下一个节点是 nil，说明本层已经查完了，进入下一层
				if node.Next == nil {
					level -= 1
					if level >= 0 {
						// 如果不是最底层继续进入到下一层
						node = node.Down
					} else {
						return nil
					}
				} else {
					node = node.Next
				}
			}
		}
	}
	return nil
}

// Pop 删除节点
func (list *SkipList) Pop() {
	// 如果没有节点就删除
	if list.Top() == nil {
		return
	}
	// 如果有节点就删除
	node := list.top()
	for node != nil {
		prevNode := node.Prev
		nextNode := node.Next

		prevNode.Next = nextNode
		if nextNode != nil {
			nextNode.Prev = prevNode
		}
		node = node.Up
	}
	if list.HeadNodeArr[list.Level].Next == nil {
		list.Level--
	}
}

func (list *SkipList) Empty() bool {
	return list.Level < 0
}
func (list *SkipList) top() *Node {
	if list.Level >= 0 {
		return list.HeadNodeArr[0].Next
	}
	return nil
}

func (list *SkipList) Top() *event.DiscreteEvent {
	if list.Level >= 0 && list.HeadNodeArr[0].Next != nil {
		return &list.HeadNodeArr[0].Next.Event
	}
	return nil
}

// Push 添加数据到跳表中
func (list *SkipList) Push(eventValue event.DiscreteEvent) {
	if list.HasNode(eventValue) != nil {
		// 如果包含相同的数据，就返回，不用添加了
		return
	}
	headNodeInsertPositionArr := make([]*Node, MaxLevel)
	// 如果不包含数据，就查找每一层的插入位置
	if list.Level >= 0 {
		// 只有层级在大于等于 0 的时候在进行循环判断，如果层级小于 0 说明是没有任何数据
		level := list.Level
		node := list.HeadNodeArr[level].Next
		for node != nil && level >= 0 {
			if event.Smaller(node.Event, eventValue) {
				// 如果节点的值大于传入的值，就应该返回上个节点并进入下一层
				headNodeInsertPositionArr[level] = node.Prev
				if node.Prev.Down == nil {
					if level-1 >= 0 {
						node = list.HeadNodeArr[level-1].Next
					} else {
						node = nil
					}
				} else {
					node = node.Prev.Down
				}
				level -= 1
			} else if event.Greater(node.Event, eventValue) {
				// 如果节点的值小于传入的值就进入下一个节点，如果下一个节点是 nil，说明本层已经查完了，进入下一层，且从下一层的头部开始
				if node.Next == nil {
					headNodeInsertPositionArr[level] = node
					level -= 1
					if level >= 0 {
						// 如果不是最底层继续进入到下一层
						node = node.Down
					} else {
						node = nil
					}
				} else {
					node = node.Next
				}
			}
		}
	}

	list.InsertValue(eventValue, headNodeInsertPositionArr)

}

func (list *SkipList) InsertValue(eventValue event.DiscreteEvent, headNodeInsertPositionArr []*Node) {
	// 插入最底层
	node := new(Node)
	node.Event = eventValue
	if list.Level < 0 {
		// 如果是空的就插入最底层数据
		list.Level = 0
		list.HeadNodeArr[0] = new(Node)
		list.HeadNodeArr[0].Next = node
		node.Prev = list.HeadNodeArr[0]
	} else {
		// 如果不是空的，就插入每一层
		// 插入最底层，
		rootNode := headNodeInsertPositionArr[0]
		nextNode := rootNode.Next

		rootNode.Next = node

		node.Prev = rootNode
		node.Next = nextNode

		if nextNode != nil {
			nextNode.Prev = node
		}

		currentLevel := 1
		for randLevel() && currentLevel <= list.Level+1 && currentLevel < MaxLevel {

			// 通过摇点 和 顶层判断是否创建新层，顶层判断有两种判断，一、不能超过预定的最高层，二、不能比当前层多出过多层，也就是说最多只能增加1层
			if headNodeInsertPositionArr[currentLevel] == nil {
				rootNode = new(Node)
				list.HeadNodeArr[currentLevel] = rootNode
			} else {
				rootNode = headNodeInsertPositionArr[currentLevel]
			}

			nextNode = rootNode.Next

			upNode := new(Node)
			upNode.Event = eventValue
			upNode.Down = node
			node.Up = upNode
			upNode.Prev = rootNode
			upNode.Next = nextNode

			rootNode.Next = upNode
			if nextNode != nil {
				nextNode.Prev = node
			}

			node = upNode

			// 增加层数
			currentLevel++
		}

		list.Level = currentLevel - 1
	}
}

// 通过抛硬币决定是否加入下一层
func randLevel() bool {
	randNum := rand.Intn(2)
	if randNum == 0 {
		return true
	}
	return false
}

// PriorityQueue 初始化跳表
func PriorityQueue() *SkipList {
	list := new(SkipList)
	list.Level = -1                            // 设置层级别
	list.HeadNodeArr = make([]*Node, MaxLevel) // 初始化头节点数组
	rand.Seed(time.Now().UnixNano())

	return list
}

func printPriorityQueue(list *SkipList) {
	fmt.Println("====================start===============" + strconv.Itoa(list.Level))
	for i := list.Level; i >= 0; i-- {
		node := list.HeadNodeArr[i].Next
		for node != nil {
			fmt.Print("{Priority:" + strconv.Itoa(node.Event.Priority) + ", last:" + strconv.Itoa(int(node.Event.Tlast)) + ", ddl:" + strconv.Itoa(int(node.Event.Tddl)) + "} -> ")
			node = node.Next
		}
		fmt.Println()
	}
	fmt.Println("====================end===============")

}

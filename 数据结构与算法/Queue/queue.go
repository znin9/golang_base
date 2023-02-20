package main

type Queue interface {
	EnQueue(interface{}) bool
	DeQueue() interface{}
	Size() int
	Peek() interface{}
}

// ArrayQueue 数组实现的队列,自己实现
type ArrayQueue struct {
	data   [10]any
	count  int
	inIdx  int
	outIdx int
}

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{}
}

// EnQueue 入队
func (aq *ArrayQueue) EnQueue(item any) bool {
	// [0,1,2,3,4,5,6,7,8,9]
	// outIdx inIdx
	if aq.count >= 10 {
		return false
	}
	idx := aq.inIdx
	aq.data[idx] = item
	aq.inIdx++
	if aq.inIdx > len(aq.data)-1 {
		aq.inIdx = 0
	}
	aq.count++
	return true
}

// DeQueue 出队
func (aq *ArrayQueue) DeQueue() (item any) {
	if aq.count == 0 {
		return nil
	}
	item = aq.data[aq.outIdx]
	aq.outIdx++
	if aq.outIdx > len(aq.data)-1 {
		aq.outIdx = 0
	}
	aq.count--
	return item
}

func (aq *ArrayQueue) Size() int {
	return aq.count
}

func (aq *ArrayQueue) Peek() interface{} {
	item := aq.data[aq.outIdx]
	return item
}

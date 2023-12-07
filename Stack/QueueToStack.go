package stack

type QueueToStack struct {
	SQueue []int
}

// 用队列实现栈
// 只需要一个队列，关键是如何让队列先进先出变为栈的先进后出
// 因为队列有头有尾，尾进头出，所以需要保证最晚压入的数据在头部
// 即压入第n个数据的时候，队列中前n-1个元素依次弹出然后再从尾巴进去即可，

func (qts *QueueToStack) Push(value int) {
	// check how many values in the queue
	n := len(qts.SQueue)
	// push the value to the queue
	qts.SQueue = append(qts.SQueue, value)
	//	make sure that the last push value is on the head of queue
	for i := 0; i < n; i++ {
		qts.SQueue = append(qts.SQueue, qts.Pop())
	}
}

func (qts *QueueToStack) Pop() int {
	ans := qts.SQueue[0]
	qts.SQueue = qts.SQueue[1:]
	return ans
}

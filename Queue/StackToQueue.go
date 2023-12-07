package Queue

type StackToQueue struct {
	InStack  []int
	OutStack []int
}

// 用栈实现队列，需要用到双栈，一个为in栈，一个为out栈
// 倒数据时，需要满足两个条件
// in倒则倒空，out非空则不倒

func (stq *StackToQueue) Push(value int) {
	stq.InStack = append(stq.InStack, value)
}

func (stq *StackToQueue) InToOut() bool {
	if len(stq.OutStack) != 0 {
		//	out栈非空，不倒数据
		return false
	}
	n := len(stq.InStack) - 1
	for i := n - 1; i >= 0; i-- {
		// in倒则倒空
		stq.OutStack = append(stq.OutStack, stq.InStack[i])
		stq.InStack = stq.InStack[:i]
	}
	return true
}

func (stq *StackToQueue) Pop() int {
	ans := stq.OutStack[len(stq.OutStack)-1]
	stq.OutStack = stq.OutStack[:len(stq.OutStack)-1]
	return ans
}

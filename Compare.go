package main

import (
	"fmt"
	stack "leetcode/Stack"
)

func main() {
	var qts stack.QueueToStack
	for i := 1; i <= 5; i++ {
		qts.Push(i)
	}
	fmt.Println(qts.SQueue)
	for i := 1; i <= 5; i++ {
		fmt.Println(qts.Pop())
	}
}

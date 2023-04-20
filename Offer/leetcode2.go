package leetcode

/*
剑指offer02：包含min函数的栈
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，
调用 min、push 及 pop 的时间复杂度都是 O(1)
*/
/*解题思路:
一开始想着用一个变量记住最小的数就可以了，但忘了pop操作之后min可能会更新，因此利用一个辅助栈min来同步顺序记录最小的数
即正常的压入，如果min栈为空或者压入的数小于min的栈顶，则压入min栈。
正常的压出，如果压出的数等于min栈的栈顶，则更新min栈
返回min值则直接返回栈顶即可*/
// type MinStack struct {
// 	stack []int
// 	//辅助的min栈
// 	min []int
// }

// /** initialize your data structure here. */
// func Constructor2() MinStack {
// 	return MinStack{}
// }

// func (this *MinStack) Push(x int) {
// 	this.stack = append(this.stack, x)
// 	//如果min栈为空或者压入的数小于min的栈顶，则压入min栈。
// 	if len(this.min) == 0 || x <= this.min[len(this.min)-1] {
// 		this.min = append(this.min, x)
// 	}
// }

// func (this *MinStack) Pop() {
// 	top := this.stack[len(this.stack)-1]
// 	this.stack = this.stack[:len(this.stack)-1]
// 	//如果压出的数等于min栈的栈顶，则更新min栈，保证min栈的栈顶始终为最小数
// 	if top == this.min[len(this.min)-1] {
// 		this.min = this.min[:len(this.min)-1]
// 	}
// }

// func (this *MinStack) Top() int {
// 	top := this.stack[len(this.stack)-1]
// 	return top
// }

// func (this *MinStack) Min() int {
// 	//最小数即min栈栈顶的数
// 	return this.min[len(this.min)-1]
// }

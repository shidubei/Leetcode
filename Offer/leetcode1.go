package leetcode
//剑指Offer01：双栈模拟队列 2023.02.05
/*解题思路：用一个栈作为输入，一个栈作为输出
输入栈存储队列的pop操作，用来存输入的数据
输出栈用于删除操作，当输出栈为空则将输入栈的数据依次弹入输出栈并进行删除*/
// type CQueue struct {
// 	//用切片来模拟栈
// 	inStack, outStack []int
// }

// func Constructor() CQueue {
// 	//创建栈的函数
// 	return CQueue{}
// }

// //队列的AppendTail方法
// func (this *CQueue) AppendTail(val int) {
// 	//将元素压入输入栈
// 	this.inStack = append(this.inStack, val)
// }

// //队列的DeleteHead方法
// func (this *CQueue) DeleteHead() int {
// 	//检查输出栈和输出栈
// 	if len(this.outStack) == 0 {
// 		if len(this.inStack) == 0 {
// 			//双栈皆为空，无数据，返回-1
// 			return -1
// 		}
// 		//输入栈不为空，则将输入栈的数据传进输出栈
// 		this.in2out()
// 	}
// 	//拿队首数据
// 	val := this.outStack[len(this.outStack)-1]
// 	//删除队首数据
// 	this.outStack = this.outStack[:len(this.outStack)-1]
// 	return val
// }

// //队列的in2out方法，用于对栈元素进行交换
// func (this *CQueue) in2out() {
// 	//将输入栈的数据依次压入输出栈
// 	for i := len(this.inStack) - 1; i >= 0; i-- {
// 		this.outStack = append(this.outStack, this.inStack[i])
// 		this.inStack = this.inStack[:len(this.inStack)-1]
// 	}
// }

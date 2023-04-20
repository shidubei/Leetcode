package leetcode

//剑指offer03：输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
/*解题思路：即逆序返回一个链表的数据
思路1：直接递归，以链表为空为终止条件，想着栈FIFO的原则，想起
defer就是栈属性，但是defer后不能接赋值语句作罢。
思路2：顺序输出再逆序。*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/*解法1：顺序输出接着逆序*/
// func reversePrint(head *ListNode) []int {
// 	var rec []int
// 	for head != nil {
// 		rec = append(rec, head.Val)
// 		head = head.Next
// 	}
// 	return rev(rec)
// }
// func rev(rec []int) []int {
// 	rev := make([]int, len(rec))
// 	for i := len(rec) - 1; i >= 0; i-- {
// 		rev[len(rec)-i] = rec[i]
// 	}
// 	return rev
// }
/*解法2：递归调用(出错)*/
// func reversePrint(head *ListNode) []int{
// 	if head ==nil{
// 		return nil
// 	}
// 	var rec []int
// 	reversePrint(head.Next)
// 	rec=append(rec, head.Val)
// 	return rec
// }

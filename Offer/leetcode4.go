package leetcode

/*剑指offer4：反转链表
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
解题思路：基本思路是利用三个指针来实现，前指针、当前指针、后指针。
反转即先记录目标的后一指针，让后让当前指针指向前一指针，前指针后移到当前位置，当前位置后移到下一位置，循环至结束*/
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func reverseList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	var pre *ListNode = nil
// 	var now *ListNode = head

// 	for now != nil {
// 		nxt := now.Next
// 		now.Next = pre
// 		pre = now
// 		now = nxt
// 	}

// 	return pre
// }

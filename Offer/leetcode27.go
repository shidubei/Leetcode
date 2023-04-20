package leetcode

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func deleteNode(head *ListNode, val int) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	if val == head.Val {
// 		return head.Next
// 	}
// 	pre, cur := head, head.Next

// 	for cur.Val != val {
// 		pre, cur = cur, cur.Next
// 		if cur == nil {
// 			break
// 		}
// 	}
// 	if cur != nil {
// 		pre.Next = cur.Next
// 	}
// 	return head
// }

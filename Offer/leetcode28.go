package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针：快慢指针
//
//	func getKthFromEnd(head *ListNode, k int) *ListNode {
//		fast,low:=head,head
//		for fast!=nil&&k>0{
//			fast=fast.Next
//			k--
//		}
//		for fast!=nil{
//			fast=fast.Next
//			low=low.Next
//		}
//		return low
//	}
//
// 直接查找
// func getKthFromEnd(head *ListNode, k int) *ListNode {
// 	cnt := 0
// 	for cur := head; cur != nil; cur = cur.Next {
// 		cnt++
// 	}
// 	for i := 0; i < cnt-k; i++ {
// 		head = head.Next
// 	}
// 	return head
// }

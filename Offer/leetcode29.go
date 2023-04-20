package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/*递归法*/
// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	if l1 == nil {
// 		return l2
// 	}else if l2 == nil {
// 		return l1
// 	}else if l1.Val<l2.Val {
// 		l1.Next=mergeTwoLists(l1.Next,l2)
// 		return l1
// 	}else {
// 		l2.Next=mergeTwoLists(l1,l2.Next)
// 		return l2
// 	}

// }
//依次迭代
// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	var prehead *ListNode
// 	prehead = &ListNode{
// 		Val:  -1,
// 		Next: nil,
// 	}
// 	prev := prehead
// 	for l1 != nil && l2 != nil {
// 		if l1.Val < l2.Val {
// 			prev.Next = l1
// 			l1 = l1.Next
// 		} else {
// 			prev.Next = l2
// 			l2 = l2.Next
// 		}
// 		prev = prev.Next
// 	}
// 	if l1 == nil {
// 		prev.Next = l2
// 	} else if l2 == nil {
// 		prev.Next = l1
// 	}
// 	return prehead.Next
// }

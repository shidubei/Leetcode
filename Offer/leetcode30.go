package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/*哈希表*/
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	m := make(map[*ListNode]bool)
// 	for cur := headA; cur != nil; cur = cur.Next {
// 		m[cur] = true
// 	}
// 	for cur := headB; cur != nil; cur = cur.Next {
// 		if m[cur] {
// 			return cur
// 		}
// 	}
// 	return nil
// }
/*双指针*/
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	if headA == nil || headB == nil {
// 		return nil
// 	}
// 	pA, pB := headA, headB
// 	for pA != pB {
// 		if pA != nil {
// 			pA = pA.Next
// 		} else {
// 			pA = headB
// 		}
// 		if pB != nil {
// 			pB = pB.Next
// 		} else {
// 			pB = headA
// 		}
// 	}
// 	return pA
// }

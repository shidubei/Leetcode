package LinkList

type LinkList struct {
	Val  int
	Next *LinkList
}

// /*题目：给定一个链表，判断该链表是否为回文链表*/
// /*思路1：利用栈先进后出的思想，逆序输出判断*/
//
//	func isPalindrome(head *LinkList) bool {
//		var stack []int
//		for i := head; i != nil; i = i.Next {
//			// 先存储进栈
//			stack = append(stack, i.Val)
//		}
//		for i := head; i != nil; i = i.Next {
//			temp := stack[len(stack)-1]
//			if temp == i.Val {
//				stack = stack[:len(stack)-1]
//			} else {
//				return false
//			}
//		}
//		return true
//	}
//
// 思路2: 双指针之快慢指针，利用两个指针一快一慢的移动来确定位置
func isPalindrome2(head *LinkList) bool {
	slow := new(LinkList)
	fast := new(LinkList)

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var temp *LinkList
	fast = slow.Next // fast 指向中位的下一个开头
	slow.Next = nil
	// 反转链表
	for fast != nil {
		// 暂存下一指针
		temp = fast.Next
		// 指向前一指针
		fast.Next = slow
		// 前一指针后移
		slow = fast
		// 当前指针后移
		fast = temp
	}
	temp = slow
	fast = head
	for fast != nil && slow != nil {
		if fast.Val != slow.Val {
			return false
		}
		slow = slow.Next
		fast = fast.Next
	}
	slow = temp.Next
	temp.Next = nil
	for slow != nil {
		fast = slow.Next
		slow.Next = temp
		temp = slow
		slow = fast
	}
	return true
}

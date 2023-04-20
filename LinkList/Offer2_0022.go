package LinkList

type ListNode struct {
	Val  int
	Next *ListNode
}

/*题目：给一个可能有环的链表，要求如果有环，返回入环的节点*/
/*思路1：快慢指针*/
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	isCircle := false
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		// 快慢指针对上，说明存在环，退出
		if fast == slow {
			isCircle = true
			break
		}
	}
	fast = head
	for fast.Next != nil && isCircle {
		if slow == fast {
			return fast
		}
		fast = fast.Next
		slow = slow.Next
	}
	return nil
}

/*思路2：哈希表*/
func detectCycle2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	m := make(map[*ListNode]int)
	cur := head
	for cur.Next != nil {
		if m[cur] == 2 {
			return cur
		}
		m[cur]++
		cur = cur.Next
	}
	return nil
}

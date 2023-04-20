package LinkList

//type LinkList struct {
//	val  int
//	Next *LinkList
//}
//
//// 快慢指针，利用两个指针的移动速率不同，从而达到在快指针达到某一位置时，慢指针到达另一位置
//// 快慢指针常用于中位，即快指针的移动速率是慢指针的两倍
//
//func basic_FastAndSlow(head *LinkList) {
//	slow := head
//	fast := head
//	// 两倍
//	for fast.Next != nil && fast.Next.Next != nil {
//		// 慢指针移动一次，快指针移动两次
//		slow = slow.Next
//		fast = fast.Next.Next
//	}
//}

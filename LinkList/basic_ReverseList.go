package LinkList

//type LinkList struct {
//	val  int
//	next *LinkList
//}
//
//// 链表基础-反转链表
//func basic_ReverseList(head *LinkList) *LinkList {
//	if head == nil {
//		return nil
//	}
//	var pre, cur *LinkList // 初始化前节点和现节点
//	cur = head             // 把链表赋值给现节点
//	for cur != nil {
//		// 记录下一节点
//		next := cur.next
//		// 现节点指向前节点
//		cur.next = pre
//		// 前节点后移至现节点
//		pre = cur
//		// 现节点后移至下一节点
//		cur = next
//	}
//	return pre
//}

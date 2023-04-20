package leetcode

/*剑指offer5：复杂链表的复制（偏困难）*/
/*解题思路：
思路1：用哈希表来记录复杂链表的指针信息，并还原
思路2：
*/// type Node struct {
// 	Val    int
// 	Next   *Node
// 	Random *Node
// }

// func copyRandomList(head *Node) *Node {
// 	if head == nil {
// 		return nil
// 	}
// 	m := make(map[*Node]*Node)
//
// 	for cur := head; cur != nil; cur = cur.Next {
// 		m[cur] = &Node{Val: cur.Val}
// 	}
// 	for cur := head; cur != nil; cur = cur.Next {
// 		m[cur].Next = m[cur.Next]
// 		m[cur].Random = m[cur.Random]
// 	}
// 	return m[head]
// }

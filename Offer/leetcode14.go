package leetcode

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// 方法一：通过一个队列来记录，FIFO，左右节点依次入队列
// func levelOrder(root *TreeNode) []int {
// 	if root == nil {
// 		return nil
// 	}
// 	var res []int
// 	quene := []*TreeNode{root}

//		for len(quene) > 0 {
//			q := quene
//			quene = nil
//			for _, v := range q {
//				res = append(res, v.Val)
//				if v.Left != nil {
//					quene = append(quene, v.Left)
//				}
//				if v.Right != nil {
//					quene = append(quene, v.Right)
//				}
//			}
//		}
//		return res
//	}
//
// 方法二、递归？
// func levelOrder(root *TreeNode) []int {
// 	if root == nil {
// 		return nil
// 	}
// 	var res []int
// 	res = append(res, root.Val)
// 	for root.Left != nil || root.Right != nil {	
// 		if root.Left != nil && root.Right != nil {
// 			levelOrder(root.Left)
// 	}
// 	return res

// }

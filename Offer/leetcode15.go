package leetcode

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func levelOrder(root *TreeNode2) [][]int {
// 	if root == nil {
// 		return nil
// 	}
// 	quene := []*TreeNode{root}
// 	var res [][]int

// 	for i := 0; len(quene) > 0; i++ {
// 		res = append(res, []int{})
// 		p := []*TreeNode{}
// 		for j := 0; j < len(quene); j++ {
// 			res[i] = append(res[i], quene[j].Val)
// 			if quene[j].Left != nil {
// 				p = append(p, quene[j].Left)
// 			}
// 			if quene[j].Right != nil {
// 				p = append(p, quene[j].Right)
// 			}
// 		}
// 		quene = p
// 	}

// 	return res
// }

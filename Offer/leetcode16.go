package leetcode

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func levelOrder(root *TreeNode) [][]int {
// 	if root == nil {
// 		return nil
// 	}
// 	quene := []*TreeNode{root}
// 	var res [][]int

// 	for i := 0; len(quene) > 0; i++ {
// 		res = append(res, []int{})
// 		p := []*TreeNode{}
// 		if i%2 != 0 {
// 			for j := 0; j < len(quene); j++ {
// 				res[i] = revList(append(res[i], quene[j].Val))
// 				if quene[j].Left != nil {
// 					p = append(p, quene[j].Left)
// 				}
// 				if quene[j].Right != nil {
// 					p = append(p, quene[j].Right)
// 				}
// 			}
// 		} else {
// 			for j := 0; j < len(quene); j++ {
// 				res[i] = append(res[i], quene[j].Val)
// 				if quene[j].Left != nil {
// 					p = append(p, quene[j].Left)
// 				}
// 				if quene[j].Right != nil {
// 					p = append(p, quene[j].Right)
// 				}
// 			}
// 		}

// 		quene = p
// 	}
// 	return res
// }

// func revList(list []int) []int {
// 	var ans []int

// 	for i := len(list) - 1; i >= 0; i-- {
// 		ans = append(ans, list[i])
// 	}

// 	return ans
// }

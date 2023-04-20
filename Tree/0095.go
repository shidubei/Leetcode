package Tree

/*题目：根据传入的数字，生成对应的所有二叉搜索树*/

/*思路：*/
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func generateTrees(n int) []*TreeNode {
//	if n == 0 {
//		return []*TreeNode{}
//	}
//	return findGenerateTrees(1, n)
//}
//
//func findGenerateTrees(start, end int) []*TreeNode {
//	tree := []*TreeNode{}
//	if start > end {
//		tree = append(tree, nil)
//		return tree
//	}
//
//	for i := start; i <= end; i++ {
//		left := findGenerateTrees(start, i-1)
//		right := findGenerateTrees(i+1, end)
//		for _, l := range left {
//			for _, r := range right {
//				root := &TreeNode{Val: i, Left: l, Right: r}
//				tree = append(tree, root)
//			}
//		}
//	}
//	return tree
//}

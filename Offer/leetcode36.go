package leetcode

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// /*深度优先DFS*/
// func pathSum(root *TreeNode, target int) (res [][]int) {
// 	var path []int
// 	// var res [][]int
// 	var dfs func(*TreeNode, int)
// 	dfs = func(node *TreeNode, diff int) {
// 		if node == nil {
// 			return
// 		}
// 		diff -= node.Val
// 		path = append(path, node.Val)
// 		defer func() { path = path[:len(path)-1] }()
// 		if node.Left == nil && node.Right == nil && diff == 0 {
// 			// res = append(res, path)
// 			res = append(res, append([]int(nil),path...))
// 			return
// 		}
// 		dfs(node.Left, diff)
// 		dfs(node.Right, diff)
// 	}
// 	dfs(root, target)
// 	return 
// }

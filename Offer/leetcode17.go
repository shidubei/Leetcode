package leetcode

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func isSubStructure(A *TreeNode, B *TreeNode) bool {
// 	if A == nil || B == nil {
// 		return false
// 	} else {
// 		//B在A的根节点或左子树或右子树上
// 		return (recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
// 	}
// }
// func recur(A, B *TreeNode) bool {
// 	if B == nil {
// 		return true
// 	}
// 	if A == nil || A.Val != B.Val {
// 		return false
// 	} else {
// 		return recur(A.Left, B.Left) && recur(A.Right, B.Right)
// 	}
// }

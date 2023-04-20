package Tree

func IsSortTree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	res := InOrderTraversal2(root)
	for i := 0; i < len(res)-1; i++ {
		if res[i] > res[i+1] {
			return false
		}
	}
	return true
}

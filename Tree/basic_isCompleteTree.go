package Tree

/*
思路：
1.满足完全二叉树的第一个条件，右孩子存在的时候左孩子必须存在
2.满足完全二叉树的第二个条件，第一个左右不全的节点之后全为叶子节点
*/
func basic_isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 层序队列
	queue := []*TreeNode{}
	// 判断是否存在非双全节点
	leaf := false
	var l *TreeNode
	var r *TreeNode

	queue = append(queue, root)
	// 层序遍历中校验
	for len(queue) > 0 {
		// 弹出
		root = queue[0]
		queue = queue[1:]
		l = root.Left
		r = root.Right
		// 校验
		// 条件1：l == nil && r !=nil
		// 条件2：leaf && (l != nil || r != nil)
		if (l == nil && r != nil) || (leaf && (l != nil || r != nil)) {
			return false
		}
		if l != nil {
			queue = append(queue, l)
		}
		if r != nil {
			queue = append(queue, r)
		}
		if l == nil || r == nil {
			leaf = true
		}
	}
	return true
}

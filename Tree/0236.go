package Tree

/*题目：在二叉树中找两个指定节点的最近公共节点*/
/*思路1：利用两个数组来存储遍历的路径，记录二者重合的最深的节点*/

/*
思路2：根据两个指定的节点位置，可以分为以下两种
2.1：节点1和节点2为另一节点祖先,此时节点1和节点2在同一子树上
2.2：节点1和节点2不互为祖先
*/

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	// 如果当前节点为空或节点1或节点2，直接返回当前节点
	if root == nil || root == p || root == q {
		return root
	}
	// 查询左子树和右子树
	Left := lowestCommonAncestor2(root.Left, p, q)
	Right := lowestCommonAncestor2(root.Right, p, q)

	// 如果左孩子返回不为空且右孩子返回不为空，说明左子树和右子树都存在待查节点
	// 说明两个节点不在一个子树上，那么返回它们所在树的父节点root
	if Left != nil && Right != nil {
		return root
	}
	// 返回一个空值
	if Left != nil {
		return Left
	} else {
		return Right
	}
}

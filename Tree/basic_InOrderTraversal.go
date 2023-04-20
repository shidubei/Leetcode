package Tree

import "fmt"

// 中序遍历-递归
func InOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	InOrderTraversal(root.Left)
	fmt.Println(root.Val)
	InOrderTraversal(root.Right)
	return
}

// 中序遍历-非递归
func InOrderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		if root != nil {
			// 如果当前节点非空，那么先将该节点入栈
			stack = append(stack, root)
			// 一直入栈左子树知道为空
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 依次弹出且进入右子树重复操作
			res = append(res, root.Val)
			root = root.Right
		}
	}
	return res
}

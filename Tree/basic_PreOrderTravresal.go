package Tree

import "fmt"

// 先序遍历-递归操作
func basic_PreOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val) // 打印当前
	basic_PreOrderTraversal(root.Left)
	basic_PreOrderTraversal(root.Right)
	return
}

// 先序遍历-非递归操纵-自己实现栈
func basic_PreOrderTraversal2(root *TreeNode) {
	if root == nil {
		return
	}
	// 初始化栈
	stack := []*TreeNode{}
	stack = append(stack, root)
	for len(stack) > 0 {
		// 弹出当前节点
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(root.Val)
		// 右不为空，先压入右孩子
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		// 左不为空，先压入左孩子
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
	}
}

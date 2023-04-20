package Tree

import "fmt"

// 后序遍历-递归
func basic_PostOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	basic_PostOrderTraversal(root.Left)
	basic_PostOrderTraversal(root.Right)
	fmt.Println(root.Val)
	return
}

// 后序遍历-非递归
func basic_PostOrderTraversal2(root *TreeNode) {
	if root == nil {
		return
	}
	stack := []*TreeNode{}
	var pre *TreeNode
	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if root.Right == nil || root.Right == pre {
				fmt.Println(root.Val)
				pre = root
				root = nil
			} else {
				stack = append(stack, root)
				root = root.Right
			}
		}
	}
}

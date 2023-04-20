package Tree

// 层序遍历-利用队列

func basic_LayerOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	res := [][]int{}
	//  i表示层数，每层入栈
	for i := 0; len(queue) > 0; i++ {
		res = append(res, []int{})
		// temp记录下层的节点
		temp := []*TreeNode{}
		for j := 0; j < len(queue); j++ {
			root = queue[j] // 每层出栈
			res[i] = append(res[i], root.Val)
			if root.Left != nil {
				temp = append(temp, root.Left)
			}
			if root.Right != nil {
				temp = append(temp, root.Right)
			}
		}
		queue = temp // 进入下一层
	}
}

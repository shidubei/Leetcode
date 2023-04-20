package Graphic

import "fmt"

func basic_DFS(node *Node) {
	if node == nil {
		return
	}
	stack := []*Node{}
	m := make(map[*Node]int)

	stack = append(stack, node)
	m[node]++

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, next := range cur.nexts {
			if _, ok := m[next]; !ok {
				// 先压回当前点
				stack = append(stack, cur)
				// 压入下一节点
				stack = append(stack, next)
				m[next]++
				// 处理但不弹出
				fmt.Println(next.Val)
				// 跳出循环
				break
			}
		}
	}
}

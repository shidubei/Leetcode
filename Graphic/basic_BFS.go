package Graphic

import "fmt"

// 图的广度优先遍历
func basic_BFS(node *Node) {
	if node == nil {
		return
	}
	// 利用队列
	queue := []*Node{}
	// hash表保证该节点依旧遍历
	m := make(map[*Node]int)
	// 第一个节点如队列
	queue = append(queue, node)
	m[node]++

	for len(queue) > 0 {
		cur := queue[0]
		fmt.Println(cur.Val)
		// 开始查找当前节点的相连节点表
		for _, next := range cur.nexts {
			// 如果该节点没有被遍历
			if _, ok := m[next]; !ok {
				m[next]++
				queue = append(queue, next)
			}
		}
	}
}

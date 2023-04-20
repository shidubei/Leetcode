package Graphic

/* 拓补排序适用于有向无环图,确定一个顺序使得所有的节点都能完全打印*/
// 思路1：需要检测入度和出度，找到第一个入度为0的点，然后更新图，继续
func basic_Toplogy(grap *Graph) []*Node {
	// 入度的hash表,判断当前节点的入度是否为0
	inMap := make(map[*Node]int)
	// 入度为0的队列，先进先出
	zeroInQueue := []*Node{}

	// 检测一个图的节点，找到所有入度为0的节点，将其入队列
	for node, _ := range grap.nodes {
		inMap[node] = node.In
		if node.In == 0 {
			zeroInQueue = append(zeroInQueue, node)
		}
	}
	res := []*Node{}
	// 开始遍历
	for len(zeroInQueue) > 0 {
		// 当前节点出队列
		cur := zeroInQueue[0]
		zeroInQueue = zeroInQueue[1:]

		res = append(res, cur)
		// 检查该节点的后续节点
		for _, next := range cur.nexts {
			// 将该节点的后续节点的入度更新
			inMap[next] = next.In - 1
			// 如果更新后的入度为0，入队等待弹出
			if inMap[next] == 0 {
				zeroInQueue = append(zeroInQueue)
			}
		}
	}
	return res
}

package Graphic

// 图的节点模板
type Node struct {
	Val   int
	In    int     // 入度
	Out   int     // 出度
	nexts []*Node //
	edges []*Edge //
}

func (n *Node) newNode(val int) {
	n.Val = val
}

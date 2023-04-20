package Graphic

// 图的边模板
type Edge struct {
	Weight int
	from   *Node
	to     *Node
}

func (e *Edge) newEdge(weight int, from, to *Node) {
	e.Weight = weight
	e.from = from
	e.to = to
}

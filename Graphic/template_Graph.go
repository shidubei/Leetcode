package Graphic

type Graph struct {
	nodes map[*Node]int
	edges map[*Edge]int
}

func (g *Graph) newGraph() {
	g.nodes = make(map[*Node]int)
	g.edges = make(map[*Edge]int)
}

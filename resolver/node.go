package resolver

type Node struct {
	Name string
	Edges []*Node
}

func (n *Node) AddEdge(edge *Node){
	n.Edges = append(n.Edges, edge)
}
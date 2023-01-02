package graph

type Node interface {
	ID() int
}

type Edge interface {
	From() Node
	To() Node
}

type Graph interface {
	Nodes() []Node
	Edge(from, to Node) Edge

	From(Node) []Node

	Has(Node) bool
	HasEdgeBetween(u, v Node) bool
}

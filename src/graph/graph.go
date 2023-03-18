package graph

import (
	"graph/src/data_structure"
	"graph/src/domain"
	"graph/src/value_object"
)

type Edge[T domain.Aggregate] struct {
	Node *Node[T]
	Cost int
}

type Node[T domain.Aggregate] struct {
	value *T
	Edges []*Edge[T]
}

func NewNode[T domain.Aggregate](value T) *Node[T] {
	return &Node[T]{
		value: &value,
		Edges: make([]*Edge[T], 0),
	}
}

func (n *Node[T]) AddEdge(node *Node[T], cost int) {
	n.Edges = append(n.Edges, &Edge[T]{
		Node: node,
		Cost: cost,
	})
}

type Graph[T domain.Aggregate] struct {
	Nodes map[value_object.Identifier[any]]*Node[T]
}

func NewGraph[T domain.Aggregate]() *Graph[T] {
	return &Graph[T]{
		Nodes: map[value_object.Identifier[any]]*Node[T]{},
	}
}

func (g *Graph[T]) AddNode(node *Node[T]) {
	value := *node.value

	g.Nodes[value.Id()] = node
}

func (g *Graph[T]) Search(initial T, predicate func(value T) bool) *T {
	visited := map[value_object.Identifier[any]]*Node[T]{}

	searchStack := data_structure.NewStack[Node[T]]()

	for _, edge := range g.Nodes[initial.Id()].Edges {
		searchStack.Push(edge.Node)
	}

	for !searchStack.IsEmpty() {
		node := searchStack.Pop()
		value := *node.value

		if visited[value.Id()] == nil {

			if predicate(value) {
				return &value
			}

			visited[value.Id()] = node
			for _, edge := range node.Edges {
				searchStack.Push(edge.Node)
			}
		}
	}

	return nil
}

func (g *Graph[T]) GetDistance() int {
	return 0
}

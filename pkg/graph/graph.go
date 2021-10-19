package graph

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type Graph interface {
	GetID() string
	AddNode(n Node)
	AddEdge(parent Node, child Node)
	GetNodeByID(id string) (Node, error)
	String() string
}

type graph struct {
	id    string
	nodes map[string]Node
}

func NewGraph() Graph {
	return &graph{
		id:    uuid.New().String(),
		nodes: map[string]Node{},
	}
}

func (g *graph) GetID() string {
	return g.String() // lol
}

func (g *graph) AddNode(n Node) {
	// Don't double-add nodes
	if _, exists := g.nodes[n.ID()]; !exists {
		g.nodes[n.ID()] = n
	}
}

func (g *graph) AddEdge(parent Node, child Node) {
	if _, ok := g.nodes[parent.ID()]; !ok {
		g.nodes[parent.ID()] = parent
	}
	if _, ok := g.nodes[child.ID()]; !ok {
		g.nodes[child.ID()] = child
	}

	g.nodes[parent.ID()].AddChild(child)
}

func (g *graph) GetNodeByID(id string) (Node, error) {
	n, ok := g.nodes[id]
	if !ok {
		return &node{}, fmt.Errorf("unable to find node with id %s", id)
	}
	return n, nil
}

/*
[
    {
        "id0": ["id1", "id2", "id3"]
    },
    {
        "id1": ["id2", "id3"]
    },
    {
        "id2": ["id4"]
    },
    {
        "id3": ["id4"]
    },
    {
        "id4": []
    }
]
*/
func (g *graph) String() string {
	repr := []string{}
	for _, n := range g.nodes {
		repr = append(repr, n.String())
	}
	s := fmt.Sprintf("[%s]", strings.Join(repr, ","))
	return s
}

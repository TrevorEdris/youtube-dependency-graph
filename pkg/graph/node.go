package graph

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Node interface {
	ID() string
	GetChildren() []Node
	AddChild(c Node)
	String() string
}

type node struct {
	id       string
	children map[string]Node
}

func NewNode(id string) (Node, error) {
	if strings.TrimSpace(id) == "" {
		return &node{}, fmt.Errorf("node ID cannot be empty")
	}
	return &node{
		id:       id,
		children: map[string]Node{},
	}, nil
}

func (n *node) ID() string {
	return n.id
}

func (n *node) GetChildren() []Node {
	children := make([]Node, len(n.children))
	for _, child := range n.children {
		children = append(children, child)
	}
	return children
}

func (n *node) AddChild(c Node) {
	// Don't double-insert nodes
	if _, exists := n.children[c.ID()]; !exists {
		n.children[c.ID()] = c
	}
}

func (n *node) String() string {
	/*
		{
		    "ParentID": ["id0", "id1", "id2"]
		}
	*/
	childIDs := []string{}
	for childID, _ := range n.children {
		childIDs = append(childIDs, childID)
	}
	repr := make(map[string][]string)
	repr[n.ID()] = childIDs
	b, _ := json.Marshal(repr)
	return string(b)
}

package graph

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Node interface {
	GetID() string
	GetLabel() string
	GetChildren() []Node
	AddChild(c Node)
	String() string
	ToJSON() string
	ToCustomJSON() string
}

// Adhere to jsongraphformatv2
// https://jsongraphformat.info/v2.0/json-graph-schema.json
/*
{
    "label": "Descriptive Label Here",
    "metadata": {
        "id": "unique_node_id"
    }
}
*/
type node struct {
	Label    string          `json:"label"`
	ID       string          `json:"id"`
	Metadata nodeMetadata    `json:"metadata"`
	children map[string]Node // Deprecated: Parent graph structure tracks edges
}

type nodeCustomJSON struct {
	Label string `json:"label"`
	ID    string `json:"id"`
}

type nodeMetadata struct {
	ID string `json:"id"`
}

func NewNode(id string, label string) (Node, error) {
	if strings.TrimSpace(id) == "" {
		return &node{}, fmt.Errorf("node ID cannot be empty")
	}
	if strings.TrimSpace(label) == "" {
		return &node{}, fmt.Errorf("node label cannot be empty")
	}
	return &node{
		ID:       id,
		Metadata: nodeMetadata{ID: id},
		Label:    label,
		children: map[string]Node{},
	}, nil
}

func (n *node) GetID() string {
	return n.Metadata.ID
}

func (n *node) GetLabel() string {
	return n.Label
}

// Deprecated: Use graph structure to track edges.
func (n *node) GetChildren() []Node {
	children := make([]Node, len(n.children))
	for _, child := range n.children {
		children = append(children, child)
	}
	return children
}

func (n *node) AddChild(c Node) {
	// Don't double-insert nodes
	if _, exists := n.children[c.GetID()]; !exists {
		n.children[c.GetID()] = c
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
	repr[n.GetID()] = childIDs
	b, _ := json.Marshal(repr)
	return string(b)
}

func (n *node) ToJSON() string {
	b, _ := json.Marshal(n)
	return string(b)
}

func (n *node) ToCustomJSON() string {
	nc := &nodeCustomJSON{
		ID:    n.Metadata.ID,
		Label: n.Label,
	}
	return nc.ToJSON()
}

func (nc *nodeCustomJSON) ToJSON() string {
	b, _ := json.Marshal(nc)
	return string(b)
}

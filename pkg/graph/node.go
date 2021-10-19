package graph

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Node defines an interface for an object that can be inserted into a graph.
type Node interface {
	GetID() string
	GetLabel() string
	ToJSON() string
}

// Adhere to jsongraphformatv2
// https://jsongraphformat.info/v2.0/json-graph-schema.json
// The ID field is in both the top-level struct as well as Metadata, to allow for extreme
// ease of visualizing the graph with various visualization libraries.
/*
{
    "label": "Descriptive Label Here",
    "id": "unique_node_id",
    "metadata": {
        "id": "unique_node_id"
    }
}
*/
type node struct {
	Label    string       `json:"label"`
	ID       string       `json:"id"`
	Metadata nodeMetadata `json:"metadata"`
}

type nodeMetadata struct {
	ID string `json:"id"`
}

// NewNode creates an instance of node, which implements the Node interface.
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
		//children: map[string]Node{},
	}, nil
}

// GetID returns the ID of the node, preferring Metadata.ID over ID, though the values are equivalent.
func (n *node) GetID() string {
	return n.Metadata.ID
}

// GetLabel returns the Label of the node.
func (n *node) GetLabel() string {
	return n.Label
}

// ToJSON returns a string representation of the node, following the json graph schema v2.
func (n *node) ToJSON() string {
	b, _ := json.Marshal(n)
	return string(b)
}

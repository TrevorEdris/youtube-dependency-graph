package graph

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type Graph interface {
	GetID() string
	AddNode(n Node)
	AddEdge(parent Node, child Node, relation string)
	GetNodeByID(id string) (Node, error)
	ToJSON() string
	ToCustomJSON() string
}

// Adhere to jsongraphformatv2
// https://jsongraphformat.info/v2.0/json-graph-schema.json
// https://github.com/jsongraph/json-graph-specification/blob/master/examples/car_graphs.json
/*
   {
       "id": "car-manufacturer-relationships",
       "type": "car",
       "label": "Car Manufacturer Relationships",
       "nodes": {
           "nissan": {
               "label": "Nissan"
           },
           "infiniti": {
               "label": "Infiniti"
           },
           "toyota": {
               "label": "Toyota"
           },
           "lexus": {
               "label": "Lexus"
           }
       },
       "edges": [
           {
               "source": "nissan",
               "target": "infiniti",
               "relation": "has_luxury_division"
           },
           {
               "source": "toyota",
               "target": "lexus",
               "relation": "has_luxury_division"
           }
       ]
   },
*/
type graph struct {
	ID    string          `json:"id"`
	Label string          `json:"label"`
	Type  string          `json:"type"`
	Nodes map[string]Node `json:"nodes"`
	Edges []Edge          `json:"edges"`
}

type graphCustomJSON struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Type  string `json:"type"`
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

// NewGraph creates an instance of graph, which implements the Graph interface.
func NewGraph(id, label, graphType string) Graph {
	if strings.TrimSpace(id) == "" {
		id = uuid.New().String()
	}
	if strings.TrimSpace(label) == "" {
		label = "Default label"
	}
	if strings.TrimSpace(graphType) == "" {
		graphType = "default_type"
	}

	return &graph{
		ID:    id,
		Label: label,
		Type:  graphType,
		Nodes: map[string]Node{},
		Edges: []Edge{},
	}
}

// GetID returns the ID of the graph.
func (g *graph) GetID() string {
	return g.ID
}

// AddNode adds Node n to the graph if it is not already in the graph.
func (g *graph) AddNode(n Node) {
	// Don't double-add nodes
	if _, exists := g.Nodes[n.GetID()]; !exists {
		g.Nodes[n.GetID()] = n
	}
}

// AddEdge adds a directed edge between Node parent and Node child, labeling it with the given relation.
// If (parent|child) do not exist in the graph yet, they will be added.
func (g *graph) AddEdge(parent Node, child Node, relation string) {
	// No-op if the graph already contains this edge
	if g.containsEdge(parent, child) {
		return
	}

	// If the graph doesn't have the edge, ensure that the parent/child nodes
	// exist in the graph (adding them if either doesn't exist yet)
	// TODO: Use g.AddNode(...) here?
	if _, ok := g.Nodes[parent.GetID()]; !ok {
		g.Nodes[parent.GetID()] = parent
	}
	if _, ok := g.Nodes[child.GetID()]; !ok {
		g.Nodes[child.GetID()] = child
	}

	// Enforce non-empty string for relation
	if strings.TrimSpace(relation) == "" {
		relation = "references"
	}

	e := NewEdge(parent.GetID(), child.GetID(), relation)
	g.Edges = append(g.Edges, e)
}

// GetNodeByID returns the Node whose ID is equivalent to the given id, or nil.
func (g *graph) GetNodeByID(id string) (Node, error) {
	n, ok := g.Nodes[id]
	if !ok {
		return &node{}, fmt.Errorf("unable to find node with id %s", id)
	}
	return n, nil
}

// ToJSON returns a string representation of the graph, following the json graph schema v2.
func (g *graph) ToJSON() string {
	b, _ := json.Marshal(g)
	return string(b)
}

/* ToCustomJSON returns a string representation of the graph, following a format similar to
json graph schema v2, however rather than
{
    "nodes": {
        "node_ID": {
            <node_details>
        }
    }
}

the format is

{
    "nodes": [
        {
            <node_details>
        }
    ]
}
*/
func (g *graph) ToCustomJSON() string {
	gc := &graphCustomJSON{
		ID:    g.ID,
		Label: g.Label,
		Type:  g.Type,
		Edges: g.Edges,
	}
	nodes := []Node{}
	for _, node := range g.Nodes {
		nodes = append(nodes, node)
	}
	gc.Nodes = nodes
	return gc.ToJSON()
}

// ToJSON returns a string representation of a graphCustomJSON struct. Refer to the comment on
// ToCustomJSON for an example of the format.
func (gc *graphCustomJSON) ToJSON() string {
	b, _ := json.Marshal(gc)
	return string(b)
}

// containsEdge returns a bool, indicating if a directed edge between Node parent and Node child exists.
func (g *graph) containsEdge(parent Node, child Node) bool {
	for _, e := range g.Edges {
		if e.GetSource() == parent.GetID() && e.GetTarget() == child.GetID() {
			return true
		}
	}
	return false
}

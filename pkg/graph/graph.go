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
	String() string
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

func (g *graph) GetID() string {
	return g.String() // lol
}

func (g *graph) AddNode(n Node) {
	// Don't double-add nodes
	if _, exists := g.Nodes[n.GetID()]; !exists {
		g.Nodes[n.GetID()] = n
	}
}

func (g *graph) AddEdge(parent Node, child Node, relation string) {
	// No-op if the graph already contains this edge
	if g.containsEdge(parent, child) {
		return
	}

	// If the graph doesn't have the edge, ensure that the parent/child nodes
	// exist in the graph (adding them if either doesn't exist yet)
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

func (g *graph) GetNodeByID(id string) (Node, error) {
	n, ok := g.Nodes[id]
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
// Deprecated: String() is deprecated. Use ToJSON() instead.
func (g *graph) String() string {
	repr := []string{}
	for _, n := range g.Nodes {
		repr = append(repr, n.String())
	}
	s := fmt.Sprintf("[%s]", strings.Join(repr, ","))
	return s
}

func (g *graph) ToJSON() string {
	b, _ := json.Marshal(g)
	return string(b)
}

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

func (gc *graphCustomJSON) ToJSON() string {
	b, _ := json.Marshal(gc)
	return string(b)
}

func (g *graph) containsEdge(parent Node, child Node) bool {
	for _, e := range g.Edges {
		if e.GetSource() == parent.GetID() && e.GetTarget() == child.GetID() {
			return true
		}
	}
	return false
}

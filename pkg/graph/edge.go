package graph

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Edge interface {
	GetID() string
	GetSource() string
	GetTarget() string
	GetRelation() string
	ToJSON() string
}

// Adhere to jsongraphformatv2
// https://jsongraphformat.info/v2.0/json-graph-schema.json
/*
{
    "source": "source_node_ID",
    "target": "target_node_ID",
    "relation": "relationship_between_source_and_target"
}
*/
type edge struct {
	ID       string `json:"id"`
	Source   string `json:"source"`
	Target   string `json:"target"`
	Relation string `json:"relation"`
	Directed bool   `json:"directed"`
	Label    string `json:"label"`
}

func NewEdge(source, target, relation string) Edge {
	return &edge{
		ID:       uuid.New().String(),
		Source:   source,
		Target:   target,
		Relation: relation,
		Label:    relation,
		Directed: true,
	}
}

func (e *edge) GetID() string {
	return e.ID
}

func (e *edge) GetSource() string {
	return e.Source
}

func (e *edge) GetTarget() string {
	return e.Target
}

func (e *edge) GetRelation() string {
	return e.Relation
}

func (e *edge) ToJSON() string {
	b, _ := json.Marshal(e)
	return string(b)
}

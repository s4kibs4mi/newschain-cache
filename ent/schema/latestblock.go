package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
)

// LatestBlock holds the schema definition for the LatestBlock entity.
type LatestBlock struct {
	ent.Schema
}

// Fields of the LatestBlock.
func (LatestBlock) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("block_number").StructTag(`json:"block_number"`),
	}
}

// Edges of the LatestBlock.
func (LatestBlock) Edges() []ent.Edge {
	return nil
}

// Indexes of the LatestBlock.
func (LatestBlock) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("block_number"),
	}
}

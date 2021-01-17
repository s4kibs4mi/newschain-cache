package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
	"github.com/google/uuid"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().StructTag(`json:"id"`),
		field.String("title").MinLen(0).StructTag(`json:"title"`),
		field.Time("created_at").StructTag(`json:"created_at"`),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return nil
}

// Indexes of the Post.
func (Post) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("title"),
		index.Fields("created_at"),
	}
}

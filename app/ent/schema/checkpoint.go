package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Checkpoint holds the schema definition for the Checkpoint entity.
type Checkpoint struct {
	ent.Schema
}

// Fields of the Checkpoint.
func (Checkpoint) Fields() []ent.Field {
	return []ent.Field{
		field.Uint8("point"),
		field.String("content").MaxLen(64),
		field.String("detail").MaxLen(255),
		field.Time("check_time").Default(time.Now),
	}
}

// Edges of the Checkpoint.
func (Checkpoint) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("task", Task.Type).
			Ref("checkpoints").
			Unique(),
	}
}

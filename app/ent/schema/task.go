package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/go-gosh/tomato/app/ent/mixin"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").MaxLen(32),
		field.String("category").MaxLen(16),
		field.Int8("star"),
		field.String("content").MaxLen(255),
		field.Time("join_time").Default(time.Now),
		field.Time("start_time").Nillable().Default(time.Now),
		field.Time("end_time").Nillable(),
		field.Time("deadline").Nillable(),
	}
}

// Mixin of Task.
func (Task) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return nil
}

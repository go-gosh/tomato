package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// UserConfig holds the schema definition for the UserConfig entity.
type UserConfig struct {
	ent.Schema
}

// Fields of the UserConfig.
func (UserConfig) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Uint8("rank"),
		field.Uint("working"),
		field.Uint("break"),
	}
}

// Edges of the UserConfig.
func (UserConfig) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("user_configs").
			Field("user_id").
			Required().
			Unique(),
	}
}

// Indexes of the UserConfig
func (UserConfig) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}

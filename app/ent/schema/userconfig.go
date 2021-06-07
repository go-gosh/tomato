package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserConfig holds the schema definition for the UserConfig entity.
type UserConfig struct {
	ent.Schema
}

// Fields of the UserConfig.
func (UserConfig) Fields() []ent.Field {
	return []ent.Field{
		field.Uint8("rank"),
		field.Uint8("working"),
		field.Uint8("break"),
	}
}

// Edges of the UserConfig.
func (UserConfig) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("user_configs").
			Unique(),
	}
}

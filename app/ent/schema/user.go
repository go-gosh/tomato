package schema

import (
	"cauliflower/app/ent/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			MaxLen(32).
			MinLen(3),
		field.String("password").
			Sensitive(),
		field.Bool("enabled"),
	}
}

// Mixin of User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_configs", UserConfig.Type),
		edge.To("user_tomatoes", UserTomato.Type),
	}
}

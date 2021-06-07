package schema

import (
	"time"

	"cauliflower/app/ent/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserTomato holds the schema definition for the UserTomato entity.
type UserTomato struct {
	ent.Schema
}

// Fields of the UserTomato.
func (UserTomato) Fields() []ent.Field {
	return []ent.Field{
		field.Time("start_time").
			Immutable().
			Default(time.Now),
		field.Time("end_time").
			Nillable(),
	}
}

// Mixin of UserTomato.
func (UserTomato) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Edges of the UserTomato.
func (UserTomato) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("user_tomatoes").
			Unique(),
	}
}

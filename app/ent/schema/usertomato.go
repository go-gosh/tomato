package schema

import (
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/index"
	"github.com/go-gosh/tomato/app/ent/mixin"

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
		field.Int("user_id"),
		field.Time("start_time").
			Immutable().
			Default(time.Now),
		field.Enum("color").
			NamedValues("red", "1", "green", "2").
			SchemaType(map[string]string{
				dialect.MySQL:  "tinyint",
				dialect.SQLite: "integer",
			}),
		field.Time("remain_time"),
		field.Time("end_time").
			Nillable().
			Optional(),
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
			Field("user_id").
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}).
			Required().
			Unique(),
	}
}

// Indexes of the UserTomato
func (UserTomato) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}

// Code generated by entc, DO NOT EDIT.

package userconfig

import (
	"github.com/go-gosh/tomato/app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Rank applies equality check predicate on the "rank" field. It's identical to RankEQ.
func Rank(v uint8) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRank), v))
	})
}

// Working applies equality check predicate on the "working" field. It's identical to WorkingEQ.
func Working(v uint8) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWorking), v))
	})
}

// Break applies equality check predicate on the "break" field. It's identical to BreakEQ.
func Break(v uint8) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBreak), v))
	})
}

// RankEQ applies the EQ predicate on the "rank" field.
func RankEQ(v uint8) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRank), v))
	})
}

// RankNEQ applies the NEQ predicate on the "rank" field.
func RankNEQ(v uint8) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRank), v))
	})
}

// RankIn applies the In predicate on the "rank" field.
func RankIn(vs ...uint8) predicate.UserConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRank), v...))
	})
}

// RankNotIn applies the NotIn predicate on the "rank" field.
func RankNotIn(vs ...uint8) predicate.UserConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRank), v...))
	})
}

// RankGT applies the GT predicate on the "rank" field.
func RankGT(v uint8) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRank), v))
	})
}

// RankGTE applies the GTE predicate on the "rank" field.
func RankGTE(v uint8) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRank), v))
	})
}

// RankLT applies the LT predicate on the "rank" field.
func RankLT(v uint8) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRank), v))
	})
}

// RankLTE applies the LTE predicate on the "rank" field.
func RankLTE(v uint8) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRank), v))
	})
}

// WorkingEQ applies the EQ predicate on the "working" field.
func WorkingEQ(v uint8) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWorking), v))
	})
}

// WorkingNEQ applies the NEQ predicate on the "working" field.
func WorkingNEQ(v uint8) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWorking), v))
	})
}

// WorkingIn applies the In predicate on the "working" field.
func WorkingIn(vs ...uint8) predicate.UserConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWorking), v...))
	})
}

// WorkingNotIn applies the NotIn predicate on the "working" field.
func WorkingNotIn(vs ...uint8) predicate.UserConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWorking), v...))
	})
}

// WorkingGT applies the GT predicate on the "working" field.
func WorkingGT(v uint8) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWorking), v))
	})
}

// WorkingGTE applies the GTE predicate on the "working" field.
func WorkingGTE(v uint8) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWorking), v))
	})
}

// WorkingLT applies the LT predicate on the "working" field.
func WorkingLT(v uint8) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWorking), v))
	})
}

// WorkingLTE applies the LTE predicate on the "working" field.
func WorkingLTE(v uint8) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWorking), v))
	})
}

// BreakEQ applies the EQ predicate on the "break" field.
func BreakEQ(v uint8) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBreak), v))
	})
}

// BreakNEQ applies the NEQ predicate on the "break" field.
func BreakNEQ(v uint8) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBreak), v))
	})
}

// BreakIn applies the In predicate on the "break" field.
func BreakIn(vs ...uint8) predicate.UserConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBreak), v...))
	})
}

// BreakNotIn applies the NotIn predicate on the "break" field.
func BreakNotIn(vs ...uint8) predicate.UserConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBreak), v...))
	})
}

// BreakGT applies the GT predicate on the "break" field.
func BreakGT(v uint8) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBreak), v))
	})
}

// BreakGTE applies the GTE predicate on the "break" field.
func BreakGTE(v uint8) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBreak), v))
	})
}

// BreakLT applies the LT predicate on the "break" field.
func BreakLT(v uint8) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBreak), v))
	})
}

// BreakLTE applies the LTE predicate on the "break" field.
func BreakLTE(v uint8) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBreak), v))
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserConfig) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserConfig) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserConfig) predicate.UserConfig {
	return predicate.UserConfig(func(s *sql.Selector) {
		p(s.Not())
	})
}

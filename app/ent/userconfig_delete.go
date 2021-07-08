// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/go-gosh/tomato/app/ent/predicate"
	"github.com/go-gosh/tomato/app/ent/userconfig"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserConfigDelete is the builder for deleting a UserConfig entity.
type UserConfigDelete struct {
	config
	hooks    []Hook
	mutation *UserConfigMutation
}

// Where appends a list predicates to the UserConfigDelete builder.
func (ucd *UserConfigDelete) Where(ps ...predicate.UserConfig) *UserConfigDelete {
	ucd.mutation.Where(ps...)
	return ucd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ucd *UserConfigDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ucd.hooks) == 0 {
		affected, err = ucd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserConfigMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ucd.mutation = mutation
			affected, err = ucd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ucd.hooks) - 1; i >= 0; i-- {
			if ucd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ucd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ucd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucd *UserConfigDelete) ExecX(ctx context.Context) int {
	n, err := ucd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ucd *UserConfigDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: userconfig.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userconfig.FieldID,
			},
		},
	}
	if ps := ucd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ucd.driver, _spec)
}

// UserConfigDeleteOne is the builder for deleting a single UserConfig entity.
type UserConfigDeleteOne struct {
	ucd *UserConfigDelete
}

// Exec executes the deletion query.
func (ucdo *UserConfigDeleteOne) Exec(ctx context.Context) error {
	n, err := ucdo.ucd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userconfig.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ucdo *UserConfigDeleteOne) ExecX(ctx context.Context) {
	ucdo.ucd.ExecX(ctx)
}

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/go-gosh/tomato/app/ent/user"
	"github.com/go-gosh/tomato/app/ent/usertomato"
)

// UserTomatoCreate is the builder for creating a UserTomato entity.
type UserTomatoCreate struct {
	config
	mutation *UserTomatoMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (utc *UserTomatoCreate) SetCreatedAt(t time.Time) *UserTomatoCreate {
	utc.mutation.SetCreatedAt(t)
	return utc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (utc *UserTomatoCreate) SetNillableCreatedAt(t *time.Time) *UserTomatoCreate {
	if t != nil {
		utc.SetCreatedAt(*t)
	}
	return utc
}

// SetUpdatedAt sets the "updated_at" field.
func (utc *UserTomatoCreate) SetUpdatedAt(t time.Time) *UserTomatoCreate {
	utc.mutation.SetUpdatedAt(t)
	return utc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (utc *UserTomatoCreate) SetNillableUpdatedAt(t *time.Time) *UserTomatoCreate {
	if t != nil {
		utc.SetUpdatedAt(*t)
	}
	return utc
}

// SetUserID sets the "user_id" field.
func (utc *UserTomatoCreate) SetUserID(i int) *UserTomatoCreate {
	utc.mutation.SetUserID(i)
	return utc
}

// SetStartTime sets the "start_time" field.
func (utc *UserTomatoCreate) SetStartTime(t time.Time) *UserTomatoCreate {
	utc.mutation.SetStartTime(t)
	return utc
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (utc *UserTomatoCreate) SetNillableStartTime(t *time.Time) *UserTomatoCreate {
	if t != nil {
		utc.SetStartTime(*t)
	}
	return utc
}

// SetColor sets the "color" field.
func (utc *UserTomatoCreate) SetColor(u usertomato.Color) *UserTomatoCreate {
	utc.mutation.SetColor(u)
	return utc
}

// SetRemainTime sets the "remain_time" field.
func (utc *UserTomatoCreate) SetRemainTime(t time.Time) *UserTomatoCreate {
	utc.mutation.SetRemainTime(t)
	return utc
}

// SetEndTime sets the "end_time" field.
func (utc *UserTomatoCreate) SetEndTime(t time.Time) *UserTomatoCreate {
	utc.mutation.SetEndTime(t)
	return utc
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (utc *UserTomatoCreate) SetNillableEndTime(t *time.Time) *UserTomatoCreate {
	if t != nil {
		utc.SetEndTime(*t)
	}
	return utc
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (utc *UserTomatoCreate) SetUsersID(id int) *UserTomatoCreate {
	utc.mutation.SetUsersID(id)
	return utc
}

// SetUsers sets the "users" edge to the User entity.
func (utc *UserTomatoCreate) SetUsers(u *User) *UserTomatoCreate {
	return utc.SetUsersID(u.ID)
}

// Mutation returns the UserTomatoMutation object of the builder.
func (utc *UserTomatoCreate) Mutation() *UserTomatoMutation {
	return utc.mutation
}

// Save creates the UserTomato in the database.
func (utc *UserTomatoCreate) Save(ctx context.Context) (*UserTomato, error) {
	var (
		err  error
		node *UserTomato
	)
	utc.defaults()
	if len(utc.hooks) == 0 {
		if err = utc.check(); err != nil {
			return nil, err
		}
		node, err = utc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserTomatoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = utc.check(); err != nil {
				return nil, err
			}
			utc.mutation = mutation
			if node, err = utc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(utc.hooks) - 1; i >= 0; i-- {
			if utc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = utc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, utc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (utc *UserTomatoCreate) SaveX(ctx context.Context) *UserTomato {
	v, err := utc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (utc *UserTomatoCreate) Exec(ctx context.Context) error {
	_, err := utc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (utc *UserTomatoCreate) ExecX(ctx context.Context) {
	if err := utc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (utc *UserTomatoCreate) defaults() {
	if _, ok := utc.mutation.CreatedAt(); !ok {
		v := usertomato.DefaultCreatedAt()
		utc.mutation.SetCreatedAt(v)
	}
	if _, ok := utc.mutation.UpdatedAt(); !ok {
		v := usertomato.DefaultUpdatedAt()
		utc.mutation.SetUpdatedAt(v)
	}
	if _, ok := utc.mutation.StartTime(); !ok {
		v := usertomato.DefaultStartTime()
		utc.mutation.SetStartTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (utc *UserTomatoCreate) check() error {
	if _, ok := utc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := utc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := utc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "user_id"`)}
	}
	if _, ok := utc.mutation.StartTime(); !ok {
		return &ValidationError{Name: "start_time", err: errors.New(`ent: missing required field "start_time"`)}
	}
	if _, ok := utc.mutation.Color(); !ok {
		return &ValidationError{Name: "color", err: errors.New(`ent: missing required field "color"`)}
	}
	if v, ok := utc.mutation.Color(); ok {
		if err := usertomato.ColorValidator(v); err != nil {
			return &ValidationError{Name: "color", err: fmt.Errorf(`ent: validator failed for field "color": %w`, err)}
		}
	}
	if _, ok := utc.mutation.RemainTime(); !ok {
		return &ValidationError{Name: "remain_time", err: errors.New(`ent: missing required field "remain_time"`)}
	}
	if _, ok := utc.mutation.UsersID(); !ok {
		return &ValidationError{Name: "users", err: errors.New("ent: missing required edge \"users\"")}
	}
	return nil
}

func (utc *UserTomatoCreate) sqlSave(ctx context.Context) (*UserTomato, error) {
	_node, _spec := utc.createSpec()
	if err := sqlgraph.CreateNode(ctx, utc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (utc *UserTomatoCreate) createSpec() (*UserTomato, *sqlgraph.CreateSpec) {
	var (
		_node = &UserTomato{config: utc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: usertomato.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usertomato.FieldID,
			},
		}
	)
	if value, ok := utc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usertomato.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := utc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usertomato.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := utc.mutation.StartTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usertomato.FieldStartTime,
		})
		_node.StartTime = value
	}
	if value, ok := utc.mutation.Color(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: usertomato.FieldColor,
		})
		_node.Color = value
	}
	if value, ok := utc.mutation.RemainTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usertomato.FieldRemainTime,
		})
		_node.RemainTime = value
	}
	if value, ok := utc.mutation.EndTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usertomato.FieldEndTime,
		})
		_node.EndTime = &value
	}
	if nodes := utc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usertomato.UsersTable,
			Columns: []string{usertomato.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserTomatoCreateBulk is the builder for creating many UserTomato entities in bulk.
type UserTomatoCreateBulk struct {
	config
	builders []*UserTomatoCreate
}

// Save creates the UserTomato entities in the database.
func (utcb *UserTomatoCreateBulk) Save(ctx context.Context) ([]*UserTomato, error) {
	specs := make([]*sqlgraph.CreateSpec, len(utcb.builders))
	nodes := make([]*UserTomato, len(utcb.builders))
	mutators := make([]Mutator, len(utcb.builders))
	for i := range utcb.builders {
		func(i int, root context.Context) {
			builder := utcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserTomatoMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, utcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, utcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, utcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (utcb *UserTomatoCreateBulk) SaveX(ctx context.Context) []*UserTomato {
	v, err := utcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (utcb *UserTomatoCreateBulk) Exec(ctx context.Context) error {
	_, err := utcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (utcb *UserTomatoCreateBulk) ExecX(ctx context.Context) {
	if err := utcb.Exec(ctx); err != nil {
		panic(err)
	}
}

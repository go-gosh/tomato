// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/go-gosh/tomato/app/ent/predicate"
	"github.com/go-gosh/tomato/app/ent/user"
	"github.com/go-gosh/tomato/app/ent/usertomato"
)

// UserTomatoUpdate is the builder for updating UserTomato entities.
type UserTomatoUpdate struct {
	config
	hooks    []Hook
	mutation *UserTomatoMutation
}

// Where appends a list predicates to the UserTomatoUpdate builder.
func (utu *UserTomatoUpdate) Where(ps ...predicate.UserTomato) *UserTomatoUpdate {
	utu.mutation.Where(ps...)
	return utu
}

// SetUpdatedAt sets the "updated_at" field.
func (utu *UserTomatoUpdate) SetUpdatedAt(t time.Time) *UserTomatoUpdate {
	utu.mutation.SetUpdatedAt(t)
	return utu
}

// SetUserID sets the "user_id" field.
func (utu *UserTomatoUpdate) SetUserID(i int) *UserTomatoUpdate {
	utu.mutation.SetUserID(i)
	return utu
}

// SetColor sets the "color" field.
func (utu *UserTomatoUpdate) SetColor(u usertomato.Color) *UserTomatoUpdate {
	utu.mutation.SetColor(u)
	return utu
}

// SetRemainTime sets the "remain_time" field.
func (utu *UserTomatoUpdate) SetRemainTime(t time.Time) *UserTomatoUpdate {
	utu.mutation.SetRemainTime(t)
	return utu
}

// SetEndTime sets the "end_time" field.
func (utu *UserTomatoUpdate) SetEndTime(t time.Time) *UserTomatoUpdate {
	utu.mutation.SetEndTime(t)
	return utu
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (utu *UserTomatoUpdate) SetNillableEndTime(t *time.Time) *UserTomatoUpdate {
	if t != nil {
		utu.SetEndTime(*t)
	}
	return utu
}

// ClearEndTime clears the value of the "end_time" field.
func (utu *UserTomatoUpdate) ClearEndTime() *UserTomatoUpdate {
	utu.mutation.ClearEndTime()
	return utu
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (utu *UserTomatoUpdate) SetUsersID(id int) *UserTomatoUpdate {
	utu.mutation.SetUsersID(id)
	return utu
}

// SetUsers sets the "users" edge to the User entity.
func (utu *UserTomatoUpdate) SetUsers(u *User) *UserTomatoUpdate {
	return utu.SetUsersID(u.ID)
}

// Mutation returns the UserTomatoMutation object of the builder.
func (utu *UserTomatoUpdate) Mutation() *UserTomatoMutation {
	return utu.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (utu *UserTomatoUpdate) ClearUsers() *UserTomatoUpdate {
	utu.mutation.ClearUsers()
	return utu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (utu *UserTomatoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	utu.defaults()
	if len(utu.hooks) == 0 {
		if err = utu.check(); err != nil {
			return 0, err
		}
		affected, err = utu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserTomatoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = utu.check(); err != nil {
				return 0, err
			}
			utu.mutation = mutation
			affected, err = utu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(utu.hooks) - 1; i >= 0; i-- {
			if utu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = utu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, utu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (utu *UserTomatoUpdate) SaveX(ctx context.Context) int {
	affected, err := utu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (utu *UserTomatoUpdate) Exec(ctx context.Context) error {
	_, err := utu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (utu *UserTomatoUpdate) ExecX(ctx context.Context) {
	if err := utu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (utu *UserTomatoUpdate) defaults() {
	if _, ok := utu.mutation.UpdatedAt(); !ok {
		v := usertomato.UpdateDefaultUpdatedAt()
		utu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (utu *UserTomatoUpdate) check() error {
	if v, ok := utu.mutation.Color(); ok {
		if err := usertomato.ColorValidator(v); err != nil {
			return &ValidationError{Name: "color", err: fmt.Errorf("ent: validator failed for field \"color\": %w", err)}
		}
	}
	if _, ok := utu.mutation.UsersID(); utu.mutation.UsersCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"users\"")
	}
	return nil
}

func (utu *UserTomatoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usertomato.Table,
			Columns: usertomato.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usertomato.FieldID,
			},
		},
	}
	if ps := utu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := utu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usertomato.FieldUpdatedAt,
		})
	}
	if value, ok := utu.mutation.Color(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: usertomato.FieldColor,
		})
	}
	if value, ok := utu.mutation.RemainTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usertomato.FieldRemainTime,
		})
	}
	if value, ok := utu.mutation.EndTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usertomato.FieldEndTime,
		})
	}
	if utu.mutation.EndTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: usertomato.FieldEndTime,
		})
	}
	if utu.mutation.UsersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := utu.mutation.UsersIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, utu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usertomato.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserTomatoUpdateOne is the builder for updating a single UserTomato entity.
type UserTomatoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserTomatoMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (utuo *UserTomatoUpdateOne) SetUpdatedAt(t time.Time) *UserTomatoUpdateOne {
	utuo.mutation.SetUpdatedAt(t)
	return utuo
}

// SetUserID sets the "user_id" field.
func (utuo *UserTomatoUpdateOne) SetUserID(i int) *UserTomatoUpdateOne {
	utuo.mutation.SetUserID(i)
	return utuo
}

// SetColor sets the "color" field.
func (utuo *UserTomatoUpdateOne) SetColor(u usertomato.Color) *UserTomatoUpdateOne {
	utuo.mutation.SetColor(u)
	return utuo
}

// SetRemainTime sets the "remain_time" field.
func (utuo *UserTomatoUpdateOne) SetRemainTime(t time.Time) *UserTomatoUpdateOne {
	utuo.mutation.SetRemainTime(t)
	return utuo
}

// SetEndTime sets the "end_time" field.
func (utuo *UserTomatoUpdateOne) SetEndTime(t time.Time) *UserTomatoUpdateOne {
	utuo.mutation.SetEndTime(t)
	return utuo
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (utuo *UserTomatoUpdateOne) SetNillableEndTime(t *time.Time) *UserTomatoUpdateOne {
	if t != nil {
		utuo.SetEndTime(*t)
	}
	return utuo
}

// ClearEndTime clears the value of the "end_time" field.
func (utuo *UserTomatoUpdateOne) ClearEndTime() *UserTomatoUpdateOne {
	utuo.mutation.ClearEndTime()
	return utuo
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (utuo *UserTomatoUpdateOne) SetUsersID(id int) *UserTomatoUpdateOne {
	utuo.mutation.SetUsersID(id)
	return utuo
}

// SetUsers sets the "users" edge to the User entity.
func (utuo *UserTomatoUpdateOne) SetUsers(u *User) *UserTomatoUpdateOne {
	return utuo.SetUsersID(u.ID)
}

// Mutation returns the UserTomatoMutation object of the builder.
func (utuo *UserTomatoUpdateOne) Mutation() *UserTomatoMutation {
	return utuo.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (utuo *UserTomatoUpdateOne) ClearUsers() *UserTomatoUpdateOne {
	utuo.mutation.ClearUsers()
	return utuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (utuo *UserTomatoUpdateOne) Select(field string, fields ...string) *UserTomatoUpdateOne {
	utuo.fields = append([]string{field}, fields...)
	return utuo
}

// Save executes the query and returns the updated UserTomato entity.
func (utuo *UserTomatoUpdateOne) Save(ctx context.Context) (*UserTomato, error) {
	var (
		err  error
		node *UserTomato
	)
	utuo.defaults()
	if len(utuo.hooks) == 0 {
		if err = utuo.check(); err != nil {
			return nil, err
		}
		node, err = utuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserTomatoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = utuo.check(); err != nil {
				return nil, err
			}
			utuo.mutation = mutation
			node, err = utuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(utuo.hooks) - 1; i >= 0; i-- {
			if utuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = utuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, utuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (utuo *UserTomatoUpdateOne) SaveX(ctx context.Context) *UserTomato {
	node, err := utuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (utuo *UserTomatoUpdateOne) Exec(ctx context.Context) error {
	_, err := utuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (utuo *UserTomatoUpdateOne) ExecX(ctx context.Context) {
	if err := utuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (utuo *UserTomatoUpdateOne) defaults() {
	if _, ok := utuo.mutation.UpdatedAt(); !ok {
		v := usertomato.UpdateDefaultUpdatedAt()
		utuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (utuo *UserTomatoUpdateOne) check() error {
	if v, ok := utuo.mutation.Color(); ok {
		if err := usertomato.ColorValidator(v); err != nil {
			return &ValidationError{Name: "color", err: fmt.Errorf("ent: validator failed for field \"color\": %w", err)}
		}
	}
	if _, ok := utuo.mutation.UsersID(); utuo.mutation.UsersCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"users\"")
	}
	return nil
}

func (utuo *UserTomatoUpdateOne) sqlSave(ctx context.Context) (_node *UserTomato, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usertomato.Table,
			Columns: usertomato.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usertomato.FieldID,
			},
		},
	}
	id, ok := utuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing UserTomato.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := utuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usertomato.FieldID)
		for _, f := range fields {
			if !usertomato.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != usertomato.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := utuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := utuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usertomato.FieldUpdatedAt,
		})
	}
	if value, ok := utuo.mutation.Color(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: usertomato.FieldColor,
		})
	}
	if value, ok := utuo.mutation.RemainTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usertomato.FieldRemainTime,
		})
	}
	if value, ok := utuo.mutation.EndTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usertomato.FieldEndTime,
		})
	}
	if utuo.mutation.EndTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: usertomato.FieldEndTime,
		})
	}
	if utuo.mutation.UsersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := utuo.mutation.UsersIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserTomato{config: utuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, utuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usertomato.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

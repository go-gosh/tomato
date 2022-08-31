// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/go-gosh/tomato/app/ent/checkpoint"
	"github.com/go-gosh/tomato/app/ent/predicate"
	"github.com/go-gosh/tomato/app/ent/task"
)

// CheckpointUpdate is the builder for updating Checkpoint entities.
type CheckpointUpdate struct {
	config
	hooks    []Hook
	mutation *CheckpointMutation
}

// Where appends a list predicates to the CheckpointUpdate builder.
func (cu *CheckpointUpdate) Where(ps ...predicate.Checkpoint) *CheckpointUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetPoint sets the "point" field.
func (cu *CheckpointUpdate) SetPoint(u uint8) *CheckpointUpdate {
	cu.mutation.ResetPoint()
	cu.mutation.SetPoint(u)
	return cu
}

// AddPoint adds u to the "point" field.
func (cu *CheckpointUpdate) AddPoint(u int8) *CheckpointUpdate {
	cu.mutation.AddPoint(u)
	return cu
}

// SetContent sets the "content" field.
func (cu *CheckpointUpdate) SetContent(s string) *CheckpointUpdate {
	cu.mutation.SetContent(s)
	return cu
}

// SetDetail sets the "detail" field.
func (cu *CheckpointUpdate) SetDetail(s string) *CheckpointUpdate {
	cu.mutation.SetDetail(s)
	return cu
}

// SetCheckTime sets the "check_time" field.
func (cu *CheckpointUpdate) SetCheckTime(t time.Time) *CheckpointUpdate {
	cu.mutation.SetCheckTime(t)
	return cu
}

// SetNillableCheckTime sets the "check_time" field if the given value is not nil.
func (cu *CheckpointUpdate) SetNillableCheckTime(t *time.Time) *CheckpointUpdate {
	if t != nil {
		cu.SetCheckTime(*t)
	}
	return cu
}

// SetTaskID sets the "task" edge to the Task entity by ID.
func (cu *CheckpointUpdate) SetTaskID(id int) *CheckpointUpdate {
	cu.mutation.SetTaskID(id)
	return cu
}

// SetNillableTaskID sets the "task" edge to the Task entity by ID if the given value is not nil.
func (cu *CheckpointUpdate) SetNillableTaskID(id *int) *CheckpointUpdate {
	if id != nil {
		cu = cu.SetTaskID(*id)
	}
	return cu
}

// SetTask sets the "task" edge to the Task entity.
func (cu *CheckpointUpdate) SetTask(t *Task) *CheckpointUpdate {
	return cu.SetTaskID(t.ID)
}

// Mutation returns the CheckpointMutation object of the builder.
func (cu *CheckpointUpdate) Mutation() *CheckpointMutation {
	return cu.mutation
}

// ClearTask clears the "task" edge to the Task entity.
func (cu *CheckpointUpdate) ClearTask() *CheckpointUpdate {
	cu.mutation.ClearTask()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CheckpointUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CheckpointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CheckpointUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CheckpointUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CheckpointUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CheckpointUpdate) check() error {
	if v, ok := cu.mutation.Content(); ok {
		if err := checkpoint.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Checkpoint.content": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Detail(); ok {
		if err := checkpoint.DetailValidator(v); err != nil {
			return &ValidationError{Name: "detail", err: fmt.Errorf(`ent: validator failed for field "Checkpoint.detail": %w`, err)}
		}
	}
	return nil
}

func (cu *CheckpointUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   checkpoint.Table,
			Columns: checkpoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: checkpoint.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Point(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: checkpoint.FieldPoint,
		})
	}
	if value, ok := cu.mutation.AddedPoint(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: checkpoint.FieldPoint,
		})
	}
	if value, ok := cu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: checkpoint.FieldContent,
		})
	}
	if value, ok := cu.mutation.Detail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: checkpoint.FieldDetail,
		})
	}
	if value, ok := cu.mutation.CheckTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: checkpoint.FieldCheckTime,
		})
	}
	if cu.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   checkpoint.TaskTable,
			Columns: []string{checkpoint.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   checkpoint.TaskTable,
			Columns: []string{checkpoint.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{checkpoint.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CheckpointUpdateOne is the builder for updating a single Checkpoint entity.
type CheckpointUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CheckpointMutation
}

// SetPoint sets the "point" field.
func (cuo *CheckpointUpdateOne) SetPoint(u uint8) *CheckpointUpdateOne {
	cuo.mutation.ResetPoint()
	cuo.mutation.SetPoint(u)
	return cuo
}

// AddPoint adds u to the "point" field.
func (cuo *CheckpointUpdateOne) AddPoint(u int8) *CheckpointUpdateOne {
	cuo.mutation.AddPoint(u)
	return cuo
}

// SetContent sets the "content" field.
func (cuo *CheckpointUpdateOne) SetContent(s string) *CheckpointUpdateOne {
	cuo.mutation.SetContent(s)
	return cuo
}

// SetDetail sets the "detail" field.
func (cuo *CheckpointUpdateOne) SetDetail(s string) *CheckpointUpdateOne {
	cuo.mutation.SetDetail(s)
	return cuo
}

// SetCheckTime sets the "check_time" field.
func (cuo *CheckpointUpdateOne) SetCheckTime(t time.Time) *CheckpointUpdateOne {
	cuo.mutation.SetCheckTime(t)
	return cuo
}

// SetNillableCheckTime sets the "check_time" field if the given value is not nil.
func (cuo *CheckpointUpdateOne) SetNillableCheckTime(t *time.Time) *CheckpointUpdateOne {
	if t != nil {
		cuo.SetCheckTime(*t)
	}
	return cuo
}

// SetTaskID sets the "task" edge to the Task entity by ID.
func (cuo *CheckpointUpdateOne) SetTaskID(id int) *CheckpointUpdateOne {
	cuo.mutation.SetTaskID(id)
	return cuo
}

// SetNillableTaskID sets the "task" edge to the Task entity by ID if the given value is not nil.
func (cuo *CheckpointUpdateOne) SetNillableTaskID(id *int) *CheckpointUpdateOne {
	if id != nil {
		cuo = cuo.SetTaskID(*id)
	}
	return cuo
}

// SetTask sets the "task" edge to the Task entity.
func (cuo *CheckpointUpdateOne) SetTask(t *Task) *CheckpointUpdateOne {
	return cuo.SetTaskID(t.ID)
}

// Mutation returns the CheckpointMutation object of the builder.
func (cuo *CheckpointUpdateOne) Mutation() *CheckpointMutation {
	return cuo.mutation
}

// ClearTask clears the "task" edge to the Task entity.
func (cuo *CheckpointUpdateOne) ClearTask() *CheckpointUpdateOne {
	cuo.mutation.ClearTask()
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CheckpointUpdateOne) Select(field string, fields ...string) *CheckpointUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Checkpoint entity.
func (cuo *CheckpointUpdateOne) Save(ctx context.Context) (*Checkpoint, error) {
	var (
		err  error
		node *Checkpoint
	)
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CheckpointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Checkpoint)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CheckpointMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CheckpointUpdateOne) SaveX(ctx context.Context) *Checkpoint {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CheckpointUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CheckpointUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CheckpointUpdateOne) check() error {
	if v, ok := cuo.mutation.Content(); ok {
		if err := checkpoint.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Checkpoint.content": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Detail(); ok {
		if err := checkpoint.DetailValidator(v); err != nil {
			return &ValidationError{Name: "detail", err: fmt.Errorf(`ent: validator failed for field "Checkpoint.detail": %w`, err)}
		}
	}
	return nil
}

func (cuo *CheckpointUpdateOne) sqlSave(ctx context.Context) (_node *Checkpoint, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   checkpoint.Table,
			Columns: checkpoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: checkpoint.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Checkpoint.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, checkpoint.FieldID)
		for _, f := range fields {
			if !checkpoint.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != checkpoint.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Point(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: checkpoint.FieldPoint,
		})
	}
	if value, ok := cuo.mutation.AddedPoint(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: checkpoint.FieldPoint,
		})
	}
	if value, ok := cuo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: checkpoint.FieldContent,
		})
	}
	if value, ok := cuo.mutation.Detail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: checkpoint.FieldDetail,
		})
	}
	if value, ok := cuo.mutation.CheckTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: checkpoint.FieldCheckTime,
		})
	}
	if cuo.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   checkpoint.TaskTable,
			Columns: []string{checkpoint.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   checkpoint.TaskTable,
			Columns: []string{checkpoint.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Checkpoint{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{checkpoint.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
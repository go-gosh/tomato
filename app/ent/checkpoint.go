// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-gosh/tomato/app/ent/checkpoint"
	"github.com/go-gosh/tomato/app/ent/task"
)

// Checkpoint is the model entity for the Checkpoint schema.
type Checkpoint struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Point holds the value of the "point" field.
	Point uint8 `json:"point,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Detail holds the value of the "detail" field.
	Detail string `json:"detail,omitempty"`
	// CheckTime holds the value of the "check_time" field.
	CheckTime time.Time `json:"check_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CheckpointQuery when eager-loading is set.
	Edges            CheckpointEdges `json:"edges"`
	task_checkpoints *int
}

// CheckpointEdges holds the relations/edges for other nodes in the graph.
type CheckpointEdges struct {
	// Task holds the value of the task edge.
	Task *Task `json:"task,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TaskOrErr returns the Task value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CheckpointEdges) TaskOrErr() (*Task, error) {
	if e.loadedTypes[0] {
		if e.Task == nil {
			// The edge task was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: task.Label}
		}
		return e.Task, nil
	}
	return nil, &NotLoadedError{edge: "task"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Checkpoint) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case checkpoint.FieldID, checkpoint.FieldPoint:
			values[i] = new(sql.NullInt64)
		case checkpoint.FieldContent, checkpoint.FieldDetail:
			values[i] = new(sql.NullString)
		case checkpoint.FieldCheckTime:
			values[i] = new(sql.NullTime)
		case checkpoint.ForeignKeys[0]: // task_checkpoints
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Checkpoint", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Checkpoint fields.
func (c *Checkpoint) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case checkpoint.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case checkpoint.FieldPoint:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field point", values[i])
			} else if value.Valid {
				c.Point = uint8(value.Int64)
			}
		case checkpoint.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				c.Content = value.String
			}
		case checkpoint.FieldDetail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field detail", values[i])
			} else if value.Valid {
				c.Detail = value.String
			}
		case checkpoint.FieldCheckTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field check_time", values[i])
			} else if value.Valid {
				c.CheckTime = value.Time
			}
		case checkpoint.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field task_checkpoints", value)
			} else if value.Valid {
				c.task_checkpoints = new(int)
				*c.task_checkpoints = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryTask queries the "task" edge of the Checkpoint entity.
func (c *Checkpoint) QueryTask() *TaskQuery {
	return (&CheckpointClient{config: c.config}).QueryTask(c)
}

// Update returns a builder for updating this Checkpoint.
// Note that you need to call Checkpoint.Unwrap() before calling this method if this Checkpoint
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Checkpoint) Update() *CheckpointUpdateOne {
	return (&CheckpointClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Checkpoint entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Checkpoint) Unwrap() *Checkpoint {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Checkpoint is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Checkpoint) String() string {
	var builder strings.Builder
	builder.WriteString("Checkpoint(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("point=")
	builder.WriteString(fmt.Sprintf("%v", c.Point))
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(c.Content)
	builder.WriteString(", ")
	builder.WriteString("detail=")
	builder.WriteString(c.Detail)
	builder.WriteString(", ")
	builder.WriteString("check_time=")
	builder.WriteString(c.CheckTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Checkpoints is a parsable slice of Checkpoint.
type Checkpoints []*Checkpoint

func (c Checkpoints) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
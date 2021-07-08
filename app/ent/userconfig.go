// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/go-gosh/tomato/app/ent/user"
	"github.com/go-gosh/tomato/app/ent/userconfig"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// UserConfig is the model entity for the UserConfig schema.
type UserConfig struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Rank holds the value of the "rank" field.
	Rank uint8 `json:"rank,omitempty"`
	// Working holds the value of the "working" field.
	Working uint8 `json:"working,omitempty"`
	// Break holds the value of the "break" field.
	Break uint8 `json:"break,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserConfigQuery when eager-loading is set.
	Edges             UserConfigEdges `json:"edges"`
	user_user_configs *int
}

// UserConfigEdges holds the relations/edges for other nodes in the graph.
type UserConfigEdges struct {
	// Users holds the value of the users edge.
	Users *User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserConfigEdges) UsersOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Users == nil {
			// The edge users was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserConfig) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case userconfig.FieldID, userconfig.FieldRank, userconfig.FieldWorking, userconfig.FieldBreak:
			values[i] = new(sql.NullInt64)
		case userconfig.ForeignKeys[0]: // user_user_configs
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserConfig", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserConfig fields.
func (uc *UserConfig) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userconfig.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			uc.ID = int(value.Int64)
		case userconfig.FieldRank:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rank", values[i])
			} else if value.Valid {
				uc.Rank = uint8(value.Int64)
			}
		case userconfig.FieldWorking:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field working", values[i])
			} else if value.Valid {
				uc.Working = uint8(value.Int64)
			}
		case userconfig.FieldBreak:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field break", values[i])
			} else if value.Valid {
				uc.Break = uint8(value.Int64)
			}
		case userconfig.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_user_configs", value)
			} else if value.Valid {
				uc.user_user_configs = new(int)
				*uc.user_user_configs = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUsers queries the "users" edge of the UserConfig entity.
func (uc *UserConfig) QueryUsers() *UserQuery {
	return (&UserConfigClient{config: uc.config}).QueryUsers(uc)
}

// Update returns a builder for updating this UserConfig.
// Note that you need to call UserConfig.Unwrap() before calling this method if this UserConfig
// was returned from a transaction, and the transaction was committed or rolled back.
func (uc *UserConfig) Update() *UserConfigUpdateOne {
	return (&UserConfigClient{config: uc.config}).UpdateOne(uc)
}

// Unwrap unwraps the UserConfig entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uc *UserConfig) Unwrap() *UserConfig {
	tx, ok := uc.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserConfig is not a transactional entity")
	}
	uc.config.driver = tx.drv
	return uc
}

// String implements the fmt.Stringer.
func (uc *UserConfig) String() string {
	var builder strings.Builder
	builder.WriteString("UserConfig(")
	builder.WriteString(fmt.Sprintf("id=%v", uc.ID))
	builder.WriteString(", rank=")
	builder.WriteString(fmt.Sprintf("%v", uc.Rank))
	builder.WriteString(", working=")
	builder.WriteString(fmt.Sprintf("%v", uc.Working))
	builder.WriteString(", break=")
	builder.WriteString(fmt.Sprintf("%v", uc.Break))
	builder.WriteByte(')')
	return builder.String()
}

// UserConfigs is a parsable slice of UserConfig.
type UserConfigs []*UserConfig

func (uc UserConfigs) config(cfg config) {
	for _i := range uc {
		uc[_i].config = cfg
	}
}

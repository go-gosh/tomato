// Code generated by ent, DO NOT EDIT.

package task

import (
	"time"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// FieldStar holds the string denoting the star field in the database.
	FieldStar = "star"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldJoinTime holds the string denoting the join_time field in the database.
	FieldJoinTime = "join_time"
	// FieldStartTime holds the string denoting the start_time field in the database.
	FieldStartTime = "start_time"
	// FieldEndTime holds the string denoting the end_time field in the database.
	FieldEndTime = "end_time"
	// FieldDeadline holds the string denoting the deadline field in the database.
	FieldDeadline = "deadline"
	// Table holds the table name of the task in the database.
	Table = "tasks"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldTitle,
	FieldCategory,
	FieldStar,
	FieldContent,
	FieldJoinTime,
	FieldStartTime,
	FieldEndTime,
	FieldDeadline,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// CategoryValidator is a validator for the "category" field. It is called by the builders before save.
	CategoryValidator func(string) error
	// ContentValidator is a validator for the "content" field. It is called by the builders before save.
	ContentValidator func(string) error
	// DefaultJoinTime holds the default value on creation for the "join_time" field.
	DefaultJoinTime func() time.Time
	// DefaultStartTime holds the default value on creation for the "start_time" field.
	DefaultStartTime func() time.Time
)

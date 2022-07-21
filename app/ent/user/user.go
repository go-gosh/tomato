// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldEnabled holds the string denoting the enabled field in the database.
	FieldEnabled = "enabled"
	// EdgeUserConfigs holds the string denoting the user_configs edge name in mutations.
	EdgeUserConfigs = "user_configs"
	// EdgeUserTomatoes holds the string denoting the user_tomatoes edge name in mutations.
	EdgeUserTomatoes = "user_tomatoes"
	// Table holds the table name of the user in the database.
	Table = "users"
	// UserConfigsTable is the table that holds the user_configs relation/edge.
	UserConfigsTable = "user_configs"
	// UserConfigsInverseTable is the table name for the UserConfig entity.
	// It exists in this package in order to avoid circular dependency with the "userconfig" package.
	UserConfigsInverseTable = "user_configs"
	// UserConfigsColumn is the table column denoting the user_configs relation/edge.
	UserConfigsColumn = "user_id"
	// UserTomatoesTable is the table that holds the user_tomatoes relation/edge.
	UserTomatoesTable = "user_tomatos"
	// UserTomatoesInverseTable is the table name for the UserTomato entity.
	// It exists in this package in order to avoid circular dependency with the "usertomato" package.
	UserTomatoesInverseTable = "user_tomatos"
	// UserTomatoesColumn is the table column denoting the user_tomatoes relation/edge.
	UserTomatoesColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUsername,
	FieldPassword,
	FieldEnabled,
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
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
)

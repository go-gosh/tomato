// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/go-gosh/tomato/app/ent/schema"
	"github.com/go-gosh/tomato/app/ent/task"
	"github.com/go-gosh/tomato/app/ent/user"
	"github.com/go-gosh/tomato/app/ent/usertomato"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	taskMixin := schema.Task{}.Mixin()
	taskMixinFields0 := taskMixin[0].Fields()
	_ = taskMixinFields0
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescCreatedAt is the schema descriptor for created_at field.
	taskDescCreatedAt := taskMixinFields0[0].Descriptor()
	// task.DefaultCreatedAt holds the default value on creation for the created_at field.
	task.DefaultCreatedAt = taskDescCreatedAt.Default.(func() time.Time)
	// taskDescUpdatedAt is the schema descriptor for updated_at field.
	taskDescUpdatedAt := taskMixinFields0[1].Descriptor()
	// task.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	task.DefaultUpdatedAt = taskDescUpdatedAt.Default.(func() time.Time)
	// task.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	task.UpdateDefaultUpdatedAt = taskDescUpdatedAt.UpdateDefault.(func() time.Time)
	// taskDescTitle is the schema descriptor for title field.
	taskDescTitle := taskFields[0].Descriptor()
	// task.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	task.TitleValidator = taskDescTitle.Validators[0].(func(string) error)
	// taskDescCategory is the schema descriptor for category field.
	taskDescCategory := taskFields[1].Descriptor()
	// task.CategoryValidator is a validator for the "category" field. It is called by the builders before save.
	task.CategoryValidator = taskDescCategory.Validators[0].(func(string) error)
	// taskDescContent is the schema descriptor for content field.
	taskDescContent := taskFields[3].Descriptor()
	// task.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	task.ContentValidator = taskDescContent.Validators[0].(func(string) error)
	// taskDescJoinTime is the schema descriptor for join_time field.
	taskDescJoinTime := taskFields[4].Descriptor()
	// task.DefaultJoinTime holds the default value on creation for the join_time field.
	task.DefaultJoinTime = taskDescJoinTime.Default.(func() time.Time)
	// taskDescStartTime is the schema descriptor for start_time field.
	taskDescStartTime := taskFields[5].Descriptor()
	// task.DefaultStartTime holds the default value on creation for the start_time field.
	task.DefaultStartTime = taskDescStartTime.Default.(func() time.Time)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = func() func(string) error {
		validators := userDescUsername.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(username string) error {
			for _, fn := range fns {
				if err := fn(username); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	usertomatoMixin := schema.UserTomato{}.Mixin()
	usertomatoMixinFields0 := usertomatoMixin[0].Fields()
	_ = usertomatoMixinFields0
	usertomatoFields := schema.UserTomato{}.Fields()
	_ = usertomatoFields
	// usertomatoDescCreatedAt is the schema descriptor for created_at field.
	usertomatoDescCreatedAt := usertomatoMixinFields0[0].Descriptor()
	// usertomato.DefaultCreatedAt holds the default value on creation for the created_at field.
	usertomato.DefaultCreatedAt = usertomatoDescCreatedAt.Default.(func() time.Time)
	// usertomatoDescUpdatedAt is the schema descriptor for updated_at field.
	usertomatoDescUpdatedAt := usertomatoMixinFields0[1].Descriptor()
	// usertomato.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	usertomato.DefaultUpdatedAt = usertomatoDescUpdatedAt.Default.(func() time.Time)
	// usertomato.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	usertomato.UpdateDefaultUpdatedAt = usertomatoDescUpdatedAt.UpdateDefault.(func() time.Time)
	// usertomatoDescStartTime is the schema descriptor for start_time field.
	usertomatoDescStartTime := usertomatoFields[1].Descriptor()
	// usertomato.DefaultStartTime holds the default value on creation for the start_time field.
	usertomato.DefaultStartTime = usertomatoDescStartTime.Default.(func() time.Time)
}

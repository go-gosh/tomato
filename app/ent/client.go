// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"cauliflower/app/ent/migrate"

	"cauliflower/app/ent/user"
	"cauliflower/app/ent/userconfig"
	"cauliflower/app/ent/usertomato"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// User is the client for interacting with the User builders.
	User *UserClient
	// UserConfig is the client for interacting with the UserConfig builders.
	UserConfig *UserConfigClient
	// UserTomato is the client for interacting with the UserTomato builders.
	UserTomato *UserTomatoClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.User = NewUserClient(c.config)
	c.UserConfig = NewUserConfigClient(c.config)
	c.UserTomato = NewUserTomatoClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		User:       NewUserClient(cfg),
		UserConfig: NewUserConfigClient(cfg),
		UserTomato: NewUserTomatoClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:     cfg,
		User:       NewUserClient(cfg),
		UserConfig: NewUserConfigClient(cfg),
		UserTomato: NewUserTomatoClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		User.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.User.Use(hooks...)
	c.UserConfig.Use(hooks...)
	c.UserTomato.Use(hooks...)
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUserConfigs queries the user_configs edge of a User.
func (c *UserClient) QueryUserConfigs(u *User) *UserConfigQuery {
	query := &UserConfigQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(userconfig.Table, userconfig.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.UserConfigsTable, user.UserConfigsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUserTomatoes queries the user_tomatoes edge of a User.
func (c *UserClient) QueryUserTomatoes(u *User) *UserTomatoQuery {
	query := &UserTomatoQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(usertomato.Table, usertomato.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.UserTomatoesTable, user.UserTomatoesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// UserConfigClient is a client for the UserConfig schema.
type UserConfigClient struct {
	config
}

// NewUserConfigClient returns a client for the UserConfig from the given config.
func NewUserConfigClient(c config) *UserConfigClient {
	return &UserConfigClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `userconfig.Hooks(f(g(h())))`.
func (c *UserConfigClient) Use(hooks ...Hook) {
	c.hooks.UserConfig = append(c.hooks.UserConfig, hooks...)
}

// Create returns a create builder for UserConfig.
func (c *UserConfigClient) Create() *UserConfigCreate {
	mutation := newUserConfigMutation(c.config, OpCreate)
	return &UserConfigCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserConfig entities.
func (c *UserConfigClient) CreateBulk(builders ...*UserConfigCreate) *UserConfigCreateBulk {
	return &UserConfigCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserConfig.
func (c *UserConfigClient) Update() *UserConfigUpdate {
	mutation := newUserConfigMutation(c.config, OpUpdate)
	return &UserConfigUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserConfigClient) UpdateOne(uc *UserConfig) *UserConfigUpdateOne {
	mutation := newUserConfigMutation(c.config, OpUpdateOne, withUserConfig(uc))
	return &UserConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserConfigClient) UpdateOneID(id int) *UserConfigUpdateOne {
	mutation := newUserConfigMutation(c.config, OpUpdateOne, withUserConfigID(id))
	return &UserConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserConfig.
func (c *UserConfigClient) Delete() *UserConfigDelete {
	mutation := newUserConfigMutation(c.config, OpDelete)
	return &UserConfigDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserConfigClient) DeleteOne(uc *UserConfig) *UserConfigDeleteOne {
	return c.DeleteOneID(uc.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserConfigClient) DeleteOneID(id int) *UserConfigDeleteOne {
	builder := c.Delete().Where(userconfig.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserConfigDeleteOne{builder}
}

// Query returns a query builder for UserConfig.
func (c *UserConfigClient) Query() *UserConfigQuery {
	return &UserConfigQuery{
		config: c.config,
	}
}

// Get returns a UserConfig entity by its id.
func (c *UserConfigClient) Get(ctx context.Context, id int) (*UserConfig, error) {
	return c.Query().Where(userconfig.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserConfigClient) GetX(ctx context.Context, id int) *UserConfig {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a UserConfig.
func (c *UserConfigClient) QueryUsers(uc *UserConfig) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := uc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(userconfig.Table, userconfig.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, userconfig.UsersTable, userconfig.UsersColumn),
		)
		fromV = sqlgraph.Neighbors(uc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserConfigClient) Hooks() []Hook {
	return c.hooks.UserConfig
}

// UserTomatoClient is a client for the UserTomato schema.
type UserTomatoClient struct {
	config
}

// NewUserTomatoClient returns a client for the UserTomato from the given config.
func NewUserTomatoClient(c config) *UserTomatoClient {
	return &UserTomatoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `usertomato.Hooks(f(g(h())))`.
func (c *UserTomatoClient) Use(hooks ...Hook) {
	c.hooks.UserTomato = append(c.hooks.UserTomato, hooks...)
}

// Create returns a create builder for UserTomato.
func (c *UserTomatoClient) Create() *UserTomatoCreate {
	mutation := newUserTomatoMutation(c.config, OpCreate)
	return &UserTomatoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserTomato entities.
func (c *UserTomatoClient) CreateBulk(builders ...*UserTomatoCreate) *UserTomatoCreateBulk {
	return &UserTomatoCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserTomato.
func (c *UserTomatoClient) Update() *UserTomatoUpdate {
	mutation := newUserTomatoMutation(c.config, OpUpdate)
	return &UserTomatoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserTomatoClient) UpdateOne(ut *UserTomato) *UserTomatoUpdateOne {
	mutation := newUserTomatoMutation(c.config, OpUpdateOne, withUserTomato(ut))
	return &UserTomatoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserTomatoClient) UpdateOneID(id int) *UserTomatoUpdateOne {
	mutation := newUserTomatoMutation(c.config, OpUpdateOne, withUserTomatoID(id))
	return &UserTomatoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserTomato.
func (c *UserTomatoClient) Delete() *UserTomatoDelete {
	mutation := newUserTomatoMutation(c.config, OpDelete)
	return &UserTomatoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserTomatoClient) DeleteOne(ut *UserTomato) *UserTomatoDeleteOne {
	return c.DeleteOneID(ut.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserTomatoClient) DeleteOneID(id int) *UserTomatoDeleteOne {
	builder := c.Delete().Where(usertomato.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserTomatoDeleteOne{builder}
}

// Query returns a query builder for UserTomato.
func (c *UserTomatoClient) Query() *UserTomatoQuery {
	return &UserTomatoQuery{
		config: c.config,
	}
}

// Get returns a UserTomato entity by its id.
func (c *UserTomatoClient) Get(ctx context.Context, id int) (*UserTomato, error) {
	return c.Query().Where(usertomato.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserTomatoClient) GetX(ctx context.Context, id int) *UserTomato {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a UserTomato.
func (c *UserTomatoClient) QueryUsers(ut *UserTomato) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ut.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(usertomato.Table, usertomato.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, usertomato.UsersTable, usertomato.UsersColumn),
		)
		fromV = sqlgraph.Neighbors(ut.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserTomatoClient) Hooks() []Hook {
	return c.hooks.UserTomato
}
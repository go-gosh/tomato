// Code generated by entc, DO NOT EDIT.

package ent

import (
	"cauliflower/app/ent/predicate"
	"cauliflower/app/ent/user"
	"cauliflower/app/ent/userconfig"
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserConfigQuery is the builder for querying UserConfig entities.
type UserConfigQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UserConfig
	// eager-loading edges.
	withUsers *UserQuery
	withFKs   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserConfigQuery builder.
func (ucq *UserConfigQuery) Where(ps ...predicate.UserConfig) *UserConfigQuery {
	ucq.predicates = append(ucq.predicates, ps...)
	return ucq
}

// Limit adds a limit step to the query.
func (ucq *UserConfigQuery) Limit(limit int) *UserConfigQuery {
	ucq.limit = &limit
	return ucq
}

// Offset adds an offset step to the query.
func (ucq *UserConfigQuery) Offset(offset int) *UserConfigQuery {
	ucq.offset = &offset
	return ucq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ucq *UserConfigQuery) Unique(unique bool) *UserConfigQuery {
	ucq.unique = &unique
	return ucq
}

// Order adds an order step to the query.
func (ucq *UserConfigQuery) Order(o ...OrderFunc) *UserConfigQuery {
	ucq.order = append(ucq.order, o...)
	return ucq
}

// QueryUsers chains the current query on the "users" edge.
func (ucq *UserConfigQuery) QueryUsers() *UserQuery {
	query := &UserQuery{config: ucq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ucq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ucq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userconfig.Table, userconfig.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, userconfig.UsersTable, userconfig.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(ucq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserConfig entity from the query.
// Returns a *NotFoundError when no UserConfig was found.
func (ucq *UserConfigQuery) First(ctx context.Context) (*UserConfig, error) {
	nodes, err := ucq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userconfig.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ucq *UserConfigQuery) FirstX(ctx context.Context) *UserConfig {
	node, err := ucq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserConfig ID from the query.
// Returns a *NotFoundError when no UserConfig ID was found.
func (ucq *UserConfigQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ucq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userconfig.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ucq *UserConfigQuery) FirstIDX(ctx context.Context) int {
	id, err := ucq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserConfig entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one UserConfig entity is not found.
// Returns a *NotFoundError when no UserConfig entities are found.
func (ucq *UserConfigQuery) Only(ctx context.Context) (*UserConfig, error) {
	nodes, err := ucq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userconfig.Label}
	default:
		return nil, &NotSingularError{userconfig.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ucq *UserConfigQuery) OnlyX(ctx context.Context) *UserConfig {
	node, err := ucq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserConfig ID in the query.
// Returns a *NotSingularError when exactly one UserConfig ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ucq *UserConfigQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ucq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userconfig.Label}
	default:
		err = &NotSingularError{userconfig.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ucq *UserConfigQuery) OnlyIDX(ctx context.Context) int {
	id, err := ucq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserConfigs.
func (ucq *UserConfigQuery) All(ctx context.Context) ([]*UserConfig, error) {
	if err := ucq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ucq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ucq *UserConfigQuery) AllX(ctx context.Context) []*UserConfig {
	nodes, err := ucq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserConfig IDs.
func (ucq *UserConfigQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ucq.Select(userconfig.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ucq *UserConfigQuery) IDsX(ctx context.Context) []int {
	ids, err := ucq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ucq *UserConfigQuery) Count(ctx context.Context) (int, error) {
	if err := ucq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ucq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ucq *UserConfigQuery) CountX(ctx context.Context) int {
	count, err := ucq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ucq *UserConfigQuery) Exist(ctx context.Context) (bool, error) {
	if err := ucq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ucq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ucq *UserConfigQuery) ExistX(ctx context.Context) bool {
	exist, err := ucq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserConfigQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ucq *UserConfigQuery) Clone() *UserConfigQuery {
	if ucq == nil {
		return nil
	}
	return &UserConfigQuery{
		config:     ucq.config,
		limit:      ucq.limit,
		offset:     ucq.offset,
		order:      append([]OrderFunc{}, ucq.order...),
		predicates: append([]predicate.UserConfig{}, ucq.predicates...),
		withUsers:  ucq.withUsers.Clone(),
		// clone intermediate query.
		sql:  ucq.sql.Clone(),
		path: ucq.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (ucq *UserConfigQuery) WithUsers(opts ...func(*UserQuery)) *UserConfigQuery {
	query := &UserQuery{config: ucq.config}
	for _, opt := range opts {
		opt(query)
	}
	ucq.withUsers = query
	return ucq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Rank uint8 `json:"rank,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserConfig.Query().
//		GroupBy(userconfig.FieldRank).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ucq *UserConfigQuery) GroupBy(field string, fields ...string) *UserConfigGroupBy {
	group := &UserConfigGroupBy{config: ucq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ucq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ucq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Rank uint8 `json:"rank,omitempty"`
//	}
//
//	client.UserConfig.Query().
//		Select(userconfig.FieldRank).
//		Scan(ctx, &v)
//
func (ucq *UserConfigQuery) Select(fields ...string) *UserConfigSelect {
	ucq.fields = append(ucq.fields, fields...)
	return &UserConfigSelect{UserConfigQuery: ucq}
}

func (ucq *UserConfigQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ucq.fields {
		if !userconfig.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ucq.path != nil {
		prev, err := ucq.path(ctx)
		if err != nil {
			return err
		}
		ucq.sql = prev
	}
	return nil
}

func (ucq *UserConfigQuery) sqlAll(ctx context.Context) ([]*UserConfig, error) {
	var (
		nodes       = []*UserConfig{}
		withFKs     = ucq.withFKs
		_spec       = ucq.querySpec()
		loadedTypes = [1]bool{
			ucq.withUsers != nil,
		}
	)
	if ucq.withUsers != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, userconfig.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &UserConfig{config: ucq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, ucq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ucq.withUsers; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*UserConfig)
		for i := range nodes {
			if nodes[i].user_user_configs == nil {
				continue
			}
			fk := *nodes[i].user_user_configs
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_user_configs" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Users = n
			}
		}
	}

	return nodes, nil
}

func (ucq *UserConfigQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ucq.querySpec()
	return sqlgraph.CountNodes(ctx, ucq.driver, _spec)
}

func (ucq *UserConfigQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ucq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ucq *UserConfigQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userconfig.Table,
			Columns: userconfig.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userconfig.FieldID,
			},
		},
		From:   ucq.sql,
		Unique: true,
	}
	if unique := ucq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ucq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userconfig.FieldID)
		for i := range fields {
			if fields[i] != userconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ucq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ucq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ucq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ucq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ucq *UserConfigQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ucq.driver.Dialect())
	t1 := builder.Table(userconfig.Table)
	columns := ucq.fields
	if len(columns) == 0 {
		columns = userconfig.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ucq.sql != nil {
		selector = ucq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range ucq.predicates {
		p(selector)
	}
	for _, p := range ucq.order {
		p(selector)
	}
	if offset := ucq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ucq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserConfigGroupBy is the group-by builder for UserConfig entities.
type UserConfigGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ucgb *UserConfigGroupBy) Aggregate(fns ...AggregateFunc) *UserConfigGroupBy {
	ucgb.fns = append(ucgb.fns, fns...)
	return ucgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ucgb *UserConfigGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ucgb.path(ctx)
	if err != nil {
		return err
	}
	ucgb.sql = query
	return ucgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ucgb *UserConfigGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ucgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ucgb *UserConfigGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ucgb.fields) > 1 {
		return nil, errors.New("ent: UserConfigGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ucgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ucgb *UserConfigGroupBy) StringsX(ctx context.Context) []string {
	v, err := ucgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ucgb *UserConfigGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ucgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userconfig.Label}
	default:
		err = fmt.Errorf("ent: UserConfigGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ucgb *UserConfigGroupBy) StringX(ctx context.Context) string {
	v, err := ucgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ucgb *UserConfigGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ucgb.fields) > 1 {
		return nil, errors.New("ent: UserConfigGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ucgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ucgb *UserConfigGroupBy) IntsX(ctx context.Context) []int {
	v, err := ucgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ucgb *UserConfigGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ucgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userconfig.Label}
	default:
		err = fmt.Errorf("ent: UserConfigGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ucgb *UserConfigGroupBy) IntX(ctx context.Context) int {
	v, err := ucgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ucgb *UserConfigGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ucgb.fields) > 1 {
		return nil, errors.New("ent: UserConfigGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ucgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ucgb *UserConfigGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ucgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ucgb *UserConfigGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ucgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userconfig.Label}
	default:
		err = fmt.Errorf("ent: UserConfigGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ucgb *UserConfigGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ucgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ucgb *UserConfigGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ucgb.fields) > 1 {
		return nil, errors.New("ent: UserConfigGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ucgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ucgb *UserConfigGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ucgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ucgb *UserConfigGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ucgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userconfig.Label}
	default:
		err = fmt.Errorf("ent: UserConfigGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ucgb *UserConfigGroupBy) BoolX(ctx context.Context) bool {
	v, err := ucgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ucgb *UserConfigGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ucgb.fields {
		if !userconfig.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ucgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ucgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ucgb *UserConfigGroupBy) sqlQuery() *sql.Selector {
	selector := ucgb.sql.Select()
	aggregation := make([]string, 0, len(ucgb.fns))
	for _, fn := range ucgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ucgb.fields)+len(ucgb.fns))
		for _, f := range ucgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ucgb.fields...)...)
}

// UserConfigSelect is the builder for selecting fields of UserConfig entities.
type UserConfigSelect struct {
	*UserConfigQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ucs *UserConfigSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ucs.prepareQuery(ctx); err != nil {
		return err
	}
	ucs.sql = ucs.UserConfigQuery.sqlQuery(ctx)
	return ucs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ucs *UserConfigSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ucs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ucs *UserConfigSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ucs.fields) > 1 {
		return nil, errors.New("ent: UserConfigSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ucs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ucs *UserConfigSelect) StringsX(ctx context.Context) []string {
	v, err := ucs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ucs *UserConfigSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ucs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userconfig.Label}
	default:
		err = fmt.Errorf("ent: UserConfigSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ucs *UserConfigSelect) StringX(ctx context.Context) string {
	v, err := ucs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ucs *UserConfigSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ucs.fields) > 1 {
		return nil, errors.New("ent: UserConfigSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ucs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ucs *UserConfigSelect) IntsX(ctx context.Context) []int {
	v, err := ucs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ucs *UserConfigSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ucs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userconfig.Label}
	default:
		err = fmt.Errorf("ent: UserConfigSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ucs *UserConfigSelect) IntX(ctx context.Context) int {
	v, err := ucs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ucs *UserConfigSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ucs.fields) > 1 {
		return nil, errors.New("ent: UserConfigSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ucs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ucs *UserConfigSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ucs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ucs *UserConfigSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ucs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userconfig.Label}
	default:
		err = fmt.Errorf("ent: UserConfigSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ucs *UserConfigSelect) Float64X(ctx context.Context) float64 {
	v, err := ucs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ucs *UserConfigSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ucs.fields) > 1 {
		return nil, errors.New("ent: UserConfigSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ucs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ucs *UserConfigSelect) BoolsX(ctx context.Context) []bool {
	v, err := ucs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ucs *UserConfigSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ucs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userconfig.Label}
	default:
		err = fmt.Errorf("ent: UserConfigSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ucs *UserConfigSelect) BoolX(ctx context.Context) bool {
	v, err := ucs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ucs *UserConfigSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ucs.sql.Query()
	if err := ucs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

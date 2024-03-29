// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/go-gosh/tomato/app/ent/predicate"
	"github.com/go-gosh/tomato/app/ent/user"
	"github.com/go-gosh/tomato/app/ent/usertomato"
)

// UserTomatoQuery is the builder for querying UserTomato entities.
type UserTomatoQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UserTomato
	// eager-loading edges.
	withUsers *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserTomatoQuery builder.
func (utq *UserTomatoQuery) Where(ps ...predicate.UserTomato) *UserTomatoQuery {
	utq.predicates = append(utq.predicates, ps...)
	return utq
}

// Limit adds a limit step to the query.
func (utq *UserTomatoQuery) Limit(limit int) *UserTomatoQuery {
	utq.limit = &limit
	return utq
}

// Offset adds an offset step to the query.
func (utq *UserTomatoQuery) Offset(offset int) *UserTomatoQuery {
	utq.offset = &offset
	return utq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (utq *UserTomatoQuery) Unique(unique bool) *UserTomatoQuery {
	utq.unique = &unique
	return utq
}

// Order adds an order step to the query.
func (utq *UserTomatoQuery) Order(o ...OrderFunc) *UserTomatoQuery {
	utq.order = append(utq.order, o...)
	return utq
}

// QueryUsers chains the current query on the "users" edge.
func (utq *UserTomatoQuery) QueryUsers() *UserQuery {
	query := &UserQuery{config: utq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := utq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := utq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usertomato.Table, usertomato.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, usertomato.UsersTable, usertomato.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(utq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserTomato entity from the query.
// Returns a *NotFoundError when no UserTomato was found.
func (utq *UserTomatoQuery) First(ctx context.Context) (*UserTomato, error) {
	nodes, err := utq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usertomato.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (utq *UserTomatoQuery) FirstX(ctx context.Context) *UserTomato {
	node, err := utq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserTomato ID from the query.
// Returns a *NotFoundError when no UserTomato ID was found.
func (utq *UserTomatoQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = utq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usertomato.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (utq *UserTomatoQuery) FirstIDX(ctx context.Context) int {
	id, err := utq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserTomato entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one UserTomato entity is not found.
// Returns a *NotFoundError when no UserTomato entities are found.
func (utq *UserTomatoQuery) Only(ctx context.Context) (*UserTomato, error) {
	nodes, err := utq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usertomato.Label}
	default:
		return nil, &NotSingularError{usertomato.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (utq *UserTomatoQuery) OnlyX(ctx context.Context) *UserTomato {
	node, err := utq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserTomato ID in the query.
// Returns a *NotSingularError when exactly one UserTomato ID is not found.
// Returns a *NotFoundError when no entities are found.
func (utq *UserTomatoQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = utq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usertomato.Label}
	default:
		err = &NotSingularError{usertomato.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (utq *UserTomatoQuery) OnlyIDX(ctx context.Context) int {
	id, err := utq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserTomatos.
func (utq *UserTomatoQuery) All(ctx context.Context) ([]*UserTomato, error) {
	if err := utq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return utq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (utq *UserTomatoQuery) AllX(ctx context.Context) []*UserTomato {
	nodes, err := utq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserTomato IDs.
func (utq *UserTomatoQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := utq.Select(usertomato.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (utq *UserTomatoQuery) IDsX(ctx context.Context) []int {
	ids, err := utq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (utq *UserTomatoQuery) Count(ctx context.Context) (int, error) {
	if err := utq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return utq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (utq *UserTomatoQuery) CountX(ctx context.Context) int {
	count, err := utq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (utq *UserTomatoQuery) Exist(ctx context.Context) (bool, error) {
	if err := utq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return utq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (utq *UserTomatoQuery) ExistX(ctx context.Context) bool {
	exist, err := utq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserTomatoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (utq *UserTomatoQuery) Clone() *UserTomatoQuery {
	if utq == nil {
		return nil
	}
	return &UserTomatoQuery{
		config:     utq.config,
		limit:      utq.limit,
		offset:     utq.offset,
		order:      append([]OrderFunc{}, utq.order...),
		predicates: append([]predicate.UserTomato{}, utq.predicates...),
		withUsers:  utq.withUsers.Clone(),
		// clone intermediate query.
		sql:  utq.sql.Clone(),
		path: utq.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (utq *UserTomatoQuery) WithUsers(opts ...func(*UserQuery)) *UserTomatoQuery {
	query := &UserQuery{config: utq.config}
	for _, opt := range opts {
		opt(query)
	}
	utq.withUsers = query
	return utq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserTomato.Query().
//		GroupBy(usertomato.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (utq *UserTomatoQuery) GroupBy(field string, fields ...string) *UserTomatoGroupBy {
	group := &UserTomatoGroupBy{config: utq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := utq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return utq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.UserTomato.Query().
//		Select(usertomato.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (utq *UserTomatoQuery) Select(fields ...string) *UserTomatoSelect {
	utq.fields = append(utq.fields, fields...)
	return &UserTomatoSelect{UserTomatoQuery: utq}
}

func (utq *UserTomatoQuery) prepareQuery(ctx context.Context) error {
	for _, f := range utq.fields {
		if !usertomato.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if utq.path != nil {
		prev, err := utq.path(ctx)
		if err != nil {
			return err
		}
		utq.sql = prev
	}
	return nil
}

func (utq *UserTomatoQuery) sqlAll(ctx context.Context) ([]*UserTomato, error) {
	var (
		nodes       = []*UserTomato{}
		_spec       = utq.querySpec()
		loadedTypes = [1]bool{
			utq.withUsers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &UserTomato{config: utq.config}
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
	if err := sqlgraph.QueryNodes(ctx, utq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := utq.withUsers; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*UserTomato)
		for i := range nodes {
			fk := nodes[i].UserID
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
				return nil, fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Users = n
			}
		}
	}

	return nodes, nil
}

func (utq *UserTomatoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := utq.querySpec()
	return sqlgraph.CountNodes(ctx, utq.driver, _spec)
}

func (utq *UserTomatoQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := utq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (utq *UserTomatoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usertomato.Table,
			Columns: usertomato.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usertomato.FieldID,
			},
		},
		From:   utq.sql,
		Unique: true,
	}
	if unique := utq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := utq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usertomato.FieldID)
		for i := range fields {
			if fields[i] != usertomato.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := utq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := utq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := utq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := utq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (utq *UserTomatoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(utq.driver.Dialect())
	t1 := builder.Table(usertomato.Table)
	columns := utq.fields
	if len(columns) == 0 {
		columns = usertomato.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if utq.sql != nil {
		selector = utq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range utq.predicates {
		p(selector)
	}
	for _, p := range utq.order {
		p(selector)
	}
	if offset := utq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := utq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserTomatoGroupBy is the group-by builder for UserTomato entities.
type UserTomatoGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (utgb *UserTomatoGroupBy) Aggregate(fns ...AggregateFunc) *UserTomatoGroupBy {
	utgb.fns = append(utgb.fns, fns...)
	return utgb
}

// Scan applies the group-by query and scans the result into the given value.
func (utgb *UserTomatoGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := utgb.path(ctx)
	if err != nil {
		return err
	}
	utgb.sql = query
	return utgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (utgb *UserTomatoGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := utgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (utgb *UserTomatoGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(utgb.fields) > 1 {
		return nil, errors.New("ent: UserTomatoGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := utgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (utgb *UserTomatoGroupBy) StringsX(ctx context.Context) []string {
	v, err := utgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (utgb *UserTomatoGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = utgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usertomato.Label}
	default:
		err = fmt.Errorf("ent: UserTomatoGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (utgb *UserTomatoGroupBy) StringX(ctx context.Context) string {
	v, err := utgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (utgb *UserTomatoGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(utgb.fields) > 1 {
		return nil, errors.New("ent: UserTomatoGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := utgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (utgb *UserTomatoGroupBy) IntsX(ctx context.Context) []int {
	v, err := utgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (utgb *UserTomatoGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = utgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usertomato.Label}
	default:
		err = fmt.Errorf("ent: UserTomatoGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (utgb *UserTomatoGroupBy) IntX(ctx context.Context) int {
	v, err := utgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (utgb *UserTomatoGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(utgb.fields) > 1 {
		return nil, errors.New("ent: UserTomatoGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := utgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (utgb *UserTomatoGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := utgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (utgb *UserTomatoGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = utgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usertomato.Label}
	default:
		err = fmt.Errorf("ent: UserTomatoGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (utgb *UserTomatoGroupBy) Float64X(ctx context.Context) float64 {
	v, err := utgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (utgb *UserTomatoGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(utgb.fields) > 1 {
		return nil, errors.New("ent: UserTomatoGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := utgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (utgb *UserTomatoGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := utgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (utgb *UserTomatoGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = utgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usertomato.Label}
	default:
		err = fmt.Errorf("ent: UserTomatoGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (utgb *UserTomatoGroupBy) BoolX(ctx context.Context) bool {
	v, err := utgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (utgb *UserTomatoGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range utgb.fields {
		if !usertomato.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := utgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := utgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (utgb *UserTomatoGroupBy) sqlQuery() *sql.Selector {
	selector := utgb.sql.Select()
	aggregation := make([]string, 0, len(utgb.fns))
	for _, fn := range utgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(utgb.fields)+len(utgb.fns))
		for _, f := range utgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(utgb.fields...)...)
}

// UserTomatoSelect is the builder for selecting fields of UserTomato entities.
type UserTomatoSelect struct {
	*UserTomatoQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (uts *UserTomatoSelect) Scan(ctx context.Context, v interface{}) error {
	if err := uts.prepareQuery(ctx); err != nil {
		return err
	}
	uts.sql = uts.UserTomatoQuery.sqlQuery(ctx)
	return uts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uts *UserTomatoSelect) ScanX(ctx context.Context, v interface{}) {
	if err := uts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (uts *UserTomatoSelect) Strings(ctx context.Context) ([]string, error) {
	if len(uts.fields) > 1 {
		return nil, errors.New("ent: UserTomatoSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := uts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uts *UserTomatoSelect) StringsX(ctx context.Context) []string {
	v, err := uts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (uts *UserTomatoSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = uts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usertomato.Label}
	default:
		err = fmt.Errorf("ent: UserTomatoSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (uts *UserTomatoSelect) StringX(ctx context.Context) string {
	v, err := uts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (uts *UserTomatoSelect) Ints(ctx context.Context) ([]int, error) {
	if len(uts.fields) > 1 {
		return nil, errors.New("ent: UserTomatoSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := uts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uts *UserTomatoSelect) IntsX(ctx context.Context) []int {
	v, err := uts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (uts *UserTomatoSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = uts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usertomato.Label}
	default:
		err = fmt.Errorf("ent: UserTomatoSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (uts *UserTomatoSelect) IntX(ctx context.Context) int {
	v, err := uts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (uts *UserTomatoSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(uts.fields) > 1 {
		return nil, errors.New("ent: UserTomatoSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := uts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uts *UserTomatoSelect) Float64sX(ctx context.Context) []float64 {
	v, err := uts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (uts *UserTomatoSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = uts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usertomato.Label}
	default:
		err = fmt.Errorf("ent: UserTomatoSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (uts *UserTomatoSelect) Float64X(ctx context.Context) float64 {
	v, err := uts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (uts *UserTomatoSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(uts.fields) > 1 {
		return nil, errors.New("ent: UserTomatoSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := uts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uts *UserTomatoSelect) BoolsX(ctx context.Context) []bool {
	v, err := uts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (uts *UserTomatoSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = uts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usertomato.Label}
	default:
		err = fmt.Errorf("ent: UserTomatoSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (uts *UserTomatoSelect) BoolX(ctx context.Context) bool {
	v, err := uts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uts *UserTomatoSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := uts.sql.Query()
	if err := uts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

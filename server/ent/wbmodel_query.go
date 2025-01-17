// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/kallydev/privacy/ent/predicate"
	"github.com/kallydev/privacy/ent/wbmodel"
)

// WBModelQuery is the builder for querying WBModel entities.
type WBModelQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.WBModel
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (wmq *WBModelQuery) Where(ps ...predicate.WBModel) *WBModelQuery {
	wmq.predicates = append(wmq.predicates, ps...)
	return wmq
}

// Limit adds a limit step to the query.
func (wmq *WBModelQuery) Limit(limit int) *WBModelQuery {
	wmq.limit = &limit
	return wmq
}

// Offset adds an offset step to the query.
func (wmq *WBModelQuery) Offset(offset int) *WBModelQuery {
	wmq.offset = &offset
	return wmq
}

// Order adds an order step to the query.
func (wmq *WBModelQuery) Order(o ...OrderFunc) *WBModelQuery {
	wmq.order = append(wmq.order, o...)
	return wmq
}

// First returns the first WBModel entity in the query. Returns *NotFoundError when no WBmodel was found.
func (wmq *WBModelQuery) First(ctx context.Context) (*WBModel, error) {
	nodes, err := wmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{wbmodel.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wmq *WBModelQuery) FirstX(ctx context.Context) *WBModel {
	node, err := wmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WBModel id in the query. Returns *NotFoundError when no id was found.
func (wmq *WBModelQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{wbmodel.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wmq *WBModelQuery) FirstIDX(ctx context.Context) int {
	id, err := wmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only WBModel entity in the query, returns an error if not exactly one entity was returned.
func (wmq *WBModelQuery) Only(ctx context.Context) (*WBModel, error) {
	nodes, err := wmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{wbmodel.Label}
	default:
		return nil, &NotSingularError{wbmodel.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wmq *WBModelQuery) OnlyX(ctx context.Context) *WBModel {
	node, err := wmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only WBModel id in the query, returns an error if not exactly one id was returned.
func (wmq *WBModelQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{wbmodel.Label}
	default:
		err = &NotSingularError{wbmodel.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wmq *WBModelQuery) OnlyIDX(ctx context.Context) int {
	id, err := wmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WBModels.
func (wmq *WBModelQuery) All(ctx context.Context) ([]*WBModel, error) {
	if err := wmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wmq *WBModelQuery) AllX(ctx context.Context) []*WBModel {
	nodes, err := wmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WBModel ids.
func (wmq *WBModelQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := wmq.Select(wbmodel.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wmq *WBModelQuery) IDsX(ctx context.Context) []int {
	ids, err := wmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wmq *WBModelQuery) Count(ctx context.Context) (int, error) {
	if err := wmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wmq *WBModelQuery) CountX(ctx context.Context) int {
	count, err := wmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wmq *WBModelQuery) Exist(ctx context.Context) (bool, error) {
	if err := wmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wmq *WBModelQuery) ExistX(ctx context.Context) bool {
	exist, err := wmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wmq *WBModelQuery) Clone() *WBModelQuery {
	if wmq == nil {
		return nil
	}
	return &WBModelQuery{
		config:     wmq.config,
		limit:      wmq.limit,
		offset:     wmq.offset,
		order:      append([]OrderFunc{}, wmq.order...),
		unique:     append([]string{}, wmq.unique...),
		predicates: append([]predicate.WBModel{}, wmq.predicates...),
		// clone intermediate query.
		sql:  wmq.sql.Clone(),
		path: wmq.path,
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		WbNumber int64 `json:"uid,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WBModel.Query().
//		GroupBy(wbmodel.FieldWbNumber).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (wmq *WBModelQuery) GroupBy(field string, fields ...string) *WBModelGroupBy {
	group := &WBModelGroupBy{config: wmq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wmq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		WbNumber int64 `json:"uid,omitempty"`
//	}
//
//	client.WBModel.Query().
//		Select(wbmodel.FieldWbNumber).
//		Scan(ctx, &v)
//
func (wmq *WBModelQuery) Select(field string, fields ...string) *WBModelSelect {
	selector := &WBModelSelect{config: wmq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wmq.sqlQuery(), nil
	}
	return selector
}

func (wmq *WBModelQuery) prepareQuery(ctx context.Context) error {
	if wmq.path != nil {
		prev, err := wmq.path(ctx)
		if err != nil {
			return err
		}
		wmq.sql = prev
	}
	return nil
}

func (wmq *WBModelQuery) sqlAll(ctx context.Context) ([]*WBModel, error) {
	var (
		nodes = []*WBModel{}
		_spec = wmq.querySpec()
	)
	_spec.ScanValues = func() []interface{} {
		node := &WBModel{config: wmq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, wmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (wmq *WBModelQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wmq.querySpec()
	return sqlgraph.CountNodes(ctx, wmq.driver, _spec)
}

func (wmq *WBModelQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := wmq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (wmq *WBModelQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   wbmodel.Table,
			Columns: wbmodel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: wbmodel.FieldID,
			},
		},
		From:   wmq.sql,
		Unique: true,
	}
	if ps := wmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, wbmodel.ValidColumn)
			}
		}
	}
	return _spec
}

func (wmq *WBModelQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(wmq.driver.Dialect())
	t1 := builder.Table(wbmodel.Table)
	selector := builder.Select(t1.Columns(wbmodel.Columns...)...).From(t1)
	if wmq.sql != nil {
		selector = wmq.sql
		selector.Select(selector.Columns(wbmodel.Columns...)...)
	}
	for _, p := range wmq.predicates {
		p(selector)
	}
	for _, p := range wmq.order {
		p(selector, wbmodel.ValidColumn)
	}
	if offset := wmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WBModelGroupBy is the builder for group-by WBModel entities.
type WBModelGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wmgb *WBModelGroupBy) Aggregate(fns ...AggregateFunc) *WBModelGroupBy {
	wmgb.fns = append(wmgb.fns, fns...)
	return wmgb
}

// Scan applies the group-by query and scan the result into the given value.
func (wmgb *WBModelGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := wmgb.path(ctx)
	if err != nil {
		return err
	}
	wmgb.sql = query
	return wmgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wmgb *WBModelGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := wmgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (wmgb *WBModelGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(wmgb.fields) > 1 {
		return nil, errors.New("ent: WBModelGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := wmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wmgb *WBModelGroupBy) StringsX(ctx context.Context) []string {
	v, err := wmgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (wmgb *WBModelGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wmgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{wbmodel.Label}
	default:
		err = fmt.Errorf("ent: WBModelGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wmgb *WBModelGroupBy) StringX(ctx context.Context) string {
	v, err := wmgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (wmgb *WBModelGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(wmgb.fields) > 1 {
		return nil, errors.New("ent: WBModelGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := wmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wmgb *WBModelGroupBy) IntsX(ctx context.Context) []int {
	v, err := wmgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (wmgb *WBModelGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wmgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{wbmodel.Label}
	default:
		err = fmt.Errorf("ent: WBModelGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wmgb *WBModelGroupBy) IntX(ctx context.Context) int {
	v, err := wmgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (wmgb *WBModelGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(wmgb.fields) > 1 {
		return nil, errors.New("ent: WBModelGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := wmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wmgb *WBModelGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := wmgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (wmgb *WBModelGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wmgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{wbmodel.Label}
	default:
		err = fmt.Errorf("ent: WBModelGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wmgb *WBModelGroupBy) Float64X(ctx context.Context) float64 {
	v, err := wmgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (wmgb *WBModelGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(wmgb.fields) > 1 {
		return nil, errors.New("ent: WBModelGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := wmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wmgb *WBModelGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := wmgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (wmgb *WBModelGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wmgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{wbmodel.Label}
	default:
		err = fmt.Errorf("ent: WBModelGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wmgb *WBModelGroupBy) BoolX(ctx context.Context) bool {
	v, err := wmgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wmgb *WBModelGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range wmgb.fields {
		if !wbmodel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := wmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wmgb *WBModelGroupBy) sqlQuery() *sql.Selector {
	selector := wmgb.sql
	columns := make([]string, 0, len(wmgb.fields)+len(wmgb.fns))
	columns = append(columns, wmgb.fields...)
	for _, fn := range wmgb.fns {
		columns = append(columns, fn(selector, wbmodel.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(wmgb.fields...)
}

// WBModelSelect is the builder for select fields of WBModel entities.
type WBModelSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (wms *WBModelSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := wms.path(ctx)
	if err != nil {
		return err
	}
	wms.sql = query
	return wms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wms *WBModelSelect) ScanX(ctx context.Context, v interface{}) {
	if err := wms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (wms *WBModelSelect) Strings(ctx context.Context) ([]string, error) {
	if len(wms.fields) > 1 {
		return nil, errors.New("ent: WBModelSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := wms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wms *WBModelSelect) StringsX(ctx context.Context) []string {
	v, err := wms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (wms *WBModelSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wms.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{wbmodel.Label}
	default:
		err = fmt.Errorf("ent: WBModelSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wms *WBModelSelect) StringX(ctx context.Context) string {
	v, err := wms.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (wms *WBModelSelect) Ints(ctx context.Context) ([]int, error) {
	if len(wms.fields) > 1 {
		return nil, errors.New("ent: WBModelSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := wms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wms *WBModelSelect) IntsX(ctx context.Context) []int {
	v, err := wms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (wms *WBModelSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wms.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{wbmodel.Label}
	default:
		err = fmt.Errorf("ent: WBModelSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wms *WBModelSelect) IntX(ctx context.Context) int {
	v, err := wms.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (wms *WBModelSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(wms.fields) > 1 {
		return nil, errors.New("ent: WBModelSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := wms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wms *WBModelSelect) Float64sX(ctx context.Context) []float64 {
	v, err := wms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (wms *WBModelSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wms.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{wbmodel.Label}
	default:
		err = fmt.Errorf("ent: WBModelSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wms *WBModelSelect) Float64X(ctx context.Context) float64 {
	v, err := wms.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (wms *WBModelSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(wms.fields) > 1 {
		return nil, errors.New("ent: WBModelSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := wms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wms *WBModelSelect) BoolsX(ctx context.Context) []bool {
	v, err := wms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (wms *WBModelSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wms.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{wbmodel.Label}
	default:
		err = fmt.Errorf("ent: WBModelSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wms *WBModelSelect) BoolX(ctx context.Context) bool {
	v, err := wms.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wms *WBModelSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range wms.fields {
		if !wbmodel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := wms.sqlQuery().Query()
	if err := wms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wms *WBModelSelect) sqlQuery() sql.Querier {
	selector := wms.sql
	selector.Select(selector.Columns(wms.fields...)...)
	return selector
}

// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/loadbalancer"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/loadbalancerstatus"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/predicate"
	"go.infratographer.com/x/gidx"
)

// LoadBalancerStatusQuery is the builder for querying LoadBalancerStatus entities.
type LoadBalancerStatusQuery struct {
	config
	ctx              *QueryContext
	order            []loadbalancerstatus.OrderOption
	inters           []Interceptor
	predicates       []predicate.LoadBalancerStatus
	withLoadBalancer *LoadBalancerQuery
	modifiers        []func(*sql.Selector)
	loadTotal        []func(context.Context, []*LoadBalancerStatus) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LoadBalancerStatusQuery builder.
func (lbsq *LoadBalancerStatusQuery) Where(ps ...predicate.LoadBalancerStatus) *LoadBalancerStatusQuery {
	lbsq.predicates = append(lbsq.predicates, ps...)
	return lbsq
}

// Limit the number of records to be returned by this query.
func (lbsq *LoadBalancerStatusQuery) Limit(limit int) *LoadBalancerStatusQuery {
	lbsq.ctx.Limit = &limit
	return lbsq
}

// Offset to start from.
func (lbsq *LoadBalancerStatusQuery) Offset(offset int) *LoadBalancerStatusQuery {
	lbsq.ctx.Offset = &offset
	return lbsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lbsq *LoadBalancerStatusQuery) Unique(unique bool) *LoadBalancerStatusQuery {
	lbsq.ctx.Unique = &unique
	return lbsq
}

// Order specifies how the records should be ordered.
func (lbsq *LoadBalancerStatusQuery) Order(o ...loadbalancerstatus.OrderOption) *LoadBalancerStatusQuery {
	lbsq.order = append(lbsq.order, o...)
	return lbsq
}

// QueryLoadBalancer chains the current query on the "load_balancer" edge.
func (lbsq *LoadBalancerStatusQuery) QueryLoadBalancer() *LoadBalancerQuery {
	query := (&LoadBalancerClient{config: lbsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lbsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lbsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(loadbalancerstatus.Table, loadbalancerstatus.FieldID, selector),
			sqlgraph.To(loadbalancer.Table, loadbalancer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, loadbalancerstatus.LoadBalancerTable, loadbalancerstatus.LoadBalancerColumn),
		)
		fromU = sqlgraph.SetNeighbors(lbsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LoadBalancerStatus entity from the query.
// Returns a *NotFoundError when no LoadBalancerStatus was found.
func (lbsq *LoadBalancerStatusQuery) First(ctx context.Context) (*LoadBalancerStatus, error) {
	nodes, err := lbsq.Limit(1).All(setContextOp(ctx, lbsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{loadbalancerstatus.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lbsq *LoadBalancerStatusQuery) FirstX(ctx context.Context) *LoadBalancerStatus {
	node, err := lbsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LoadBalancerStatus ID from the query.
// Returns a *NotFoundError when no LoadBalancerStatus ID was found.
func (lbsq *LoadBalancerStatusQuery) FirstID(ctx context.Context) (id gidx.PrefixedID, err error) {
	var ids []gidx.PrefixedID
	if ids, err = lbsq.Limit(1).IDs(setContextOp(ctx, lbsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{loadbalancerstatus.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lbsq *LoadBalancerStatusQuery) FirstIDX(ctx context.Context) gidx.PrefixedID {
	id, err := lbsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LoadBalancerStatus entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LoadBalancerStatus entity is found.
// Returns a *NotFoundError when no LoadBalancerStatus entities are found.
func (lbsq *LoadBalancerStatusQuery) Only(ctx context.Context) (*LoadBalancerStatus, error) {
	nodes, err := lbsq.Limit(2).All(setContextOp(ctx, lbsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{loadbalancerstatus.Label}
	default:
		return nil, &NotSingularError{loadbalancerstatus.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lbsq *LoadBalancerStatusQuery) OnlyX(ctx context.Context) *LoadBalancerStatus {
	node, err := lbsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LoadBalancerStatus ID in the query.
// Returns a *NotSingularError when more than one LoadBalancerStatus ID is found.
// Returns a *NotFoundError when no entities are found.
func (lbsq *LoadBalancerStatusQuery) OnlyID(ctx context.Context) (id gidx.PrefixedID, err error) {
	var ids []gidx.PrefixedID
	if ids, err = lbsq.Limit(2).IDs(setContextOp(ctx, lbsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{loadbalancerstatus.Label}
	default:
		err = &NotSingularError{loadbalancerstatus.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lbsq *LoadBalancerStatusQuery) OnlyIDX(ctx context.Context) gidx.PrefixedID {
	id, err := lbsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LoadBalancerStatusSlice.
func (lbsq *LoadBalancerStatusQuery) All(ctx context.Context) ([]*LoadBalancerStatus, error) {
	ctx = setContextOp(ctx, lbsq.ctx, "All")
	if err := lbsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LoadBalancerStatus, *LoadBalancerStatusQuery]()
	return withInterceptors[[]*LoadBalancerStatus](ctx, lbsq, qr, lbsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lbsq *LoadBalancerStatusQuery) AllX(ctx context.Context) []*LoadBalancerStatus {
	nodes, err := lbsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LoadBalancerStatus IDs.
func (lbsq *LoadBalancerStatusQuery) IDs(ctx context.Context) (ids []gidx.PrefixedID, err error) {
	if lbsq.ctx.Unique == nil && lbsq.path != nil {
		lbsq.Unique(true)
	}
	ctx = setContextOp(ctx, lbsq.ctx, "IDs")
	if err = lbsq.Select(loadbalancerstatus.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lbsq *LoadBalancerStatusQuery) IDsX(ctx context.Context) []gidx.PrefixedID {
	ids, err := lbsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lbsq *LoadBalancerStatusQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lbsq.ctx, "Count")
	if err := lbsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lbsq, querierCount[*LoadBalancerStatusQuery](), lbsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lbsq *LoadBalancerStatusQuery) CountX(ctx context.Context) int {
	count, err := lbsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lbsq *LoadBalancerStatusQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lbsq.ctx, "Exist")
	switch _, err := lbsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lbsq *LoadBalancerStatusQuery) ExistX(ctx context.Context) bool {
	exist, err := lbsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LoadBalancerStatusQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lbsq *LoadBalancerStatusQuery) Clone() *LoadBalancerStatusQuery {
	if lbsq == nil {
		return nil
	}
	return &LoadBalancerStatusQuery{
		config:           lbsq.config,
		ctx:              lbsq.ctx.Clone(),
		order:            append([]loadbalancerstatus.OrderOption{}, lbsq.order...),
		inters:           append([]Interceptor{}, lbsq.inters...),
		predicates:       append([]predicate.LoadBalancerStatus{}, lbsq.predicates...),
		withLoadBalancer: lbsq.withLoadBalancer.Clone(),
		// clone intermediate query.
		sql:  lbsq.sql.Clone(),
		path: lbsq.path,
	}
}

// WithLoadBalancer tells the query-builder to eager-load the nodes that are connected to
// the "load_balancer" edge. The optional arguments are used to configure the query builder of the edge.
func (lbsq *LoadBalancerStatusQuery) WithLoadBalancer(opts ...func(*LoadBalancerQuery)) *LoadBalancerStatusQuery {
	query := (&LoadBalancerClient{config: lbsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lbsq.withLoadBalancer = query
	return lbsq
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
//	client.LoadBalancerStatus.Query().
//		GroupBy(loadbalancerstatus.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (lbsq *LoadBalancerStatusQuery) GroupBy(field string, fields ...string) *LoadBalancerStatusGroupBy {
	lbsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LoadBalancerStatusGroupBy{build: lbsq}
	grbuild.flds = &lbsq.ctx.Fields
	grbuild.label = loadbalancerstatus.Label
	grbuild.scan = grbuild.Scan
	return grbuild
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
//	client.LoadBalancerStatus.Query().
//		Select(loadbalancerstatus.FieldCreatedAt).
//		Scan(ctx, &v)
func (lbsq *LoadBalancerStatusQuery) Select(fields ...string) *LoadBalancerStatusSelect {
	lbsq.ctx.Fields = append(lbsq.ctx.Fields, fields...)
	sbuild := &LoadBalancerStatusSelect{LoadBalancerStatusQuery: lbsq}
	sbuild.label = loadbalancerstatus.Label
	sbuild.flds, sbuild.scan = &lbsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LoadBalancerStatusSelect configured with the given aggregations.
func (lbsq *LoadBalancerStatusQuery) Aggregate(fns ...AggregateFunc) *LoadBalancerStatusSelect {
	return lbsq.Select().Aggregate(fns...)
}

func (lbsq *LoadBalancerStatusQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lbsq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lbsq); err != nil {
				return err
			}
		}
	}
	for _, f := range lbsq.ctx.Fields {
		if !loadbalancerstatus.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if lbsq.path != nil {
		prev, err := lbsq.path(ctx)
		if err != nil {
			return err
		}
		lbsq.sql = prev
	}
	return nil
}

func (lbsq *LoadBalancerStatusQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LoadBalancerStatus, error) {
	var (
		nodes       = []*LoadBalancerStatus{}
		_spec       = lbsq.querySpec()
		loadedTypes = [1]bool{
			lbsq.withLoadBalancer != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LoadBalancerStatus).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LoadBalancerStatus{config: lbsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(lbsq.modifiers) > 0 {
		_spec.Modifiers = lbsq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lbsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lbsq.withLoadBalancer; query != nil {
		if err := lbsq.loadLoadBalancer(ctx, query, nodes, nil,
			func(n *LoadBalancerStatus, e *LoadBalancer) { n.Edges.LoadBalancer = e }); err != nil {
			return nil, err
		}
	}
	for i := range lbsq.loadTotal {
		if err := lbsq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lbsq *LoadBalancerStatusQuery) loadLoadBalancer(ctx context.Context, query *LoadBalancerQuery, nodes []*LoadBalancerStatus, init func(*LoadBalancerStatus), assign func(*LoadBalancerStatus, *LoadBalancer)) error {
	ids := make([]gidx.PrefixedID, 0, len(nodes))
	nodeids := make(map[gidx.PrefixedID][]*LoadBalancerStatus)
	for i := range nodes {
		fk := nodes[i].LoadBalancerID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(loadbalancer.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "load_balancer_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (lbsq *LoadBalancerStatusQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lbsq.querySpec()
	if len(lbsq.modifiers) > 0 {
		_spec.Modifiers = lbsq.modifiers
	}
	_spec.Node.Columns = lbsq.ctx.Fields
	if len(lbsq.ctx.Fields) > 0 {
		_spec.Unique = lbsq.ctx.Unique != nil && *lbsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lbsq.driver, _spec)
}

func (lbsq *LoadBalancerStatusQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(loadbalancerstatus.Table, loadbalancerstatus.Columns, sqlgraph.NewFieldSpec(loadbalancerstatus.FieldID, field.TypeString))
	_spec.From = lbsq.sql
	if unique := lbsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lbsq.path != nil {
		_spec.Unique = true
	}
	if fields := lbsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, loadbalancerstatus.FieldID)
		for i := range fields {
			if fields[i] != loadbalancerstatus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if lbsq.withLoadBalancer != nil {
			_spec.Node.AddColumnOnce(loadbalancerstatus.FieldLoadBalancerID)
		}
	}
	if ps := lbsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lbsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lbsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lbsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lbsq *LoadBalancerStatusQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lbsq.driver.Dialect())
	t1 := builder.Table(loadbalancerstatus.Table)
	columns := lbsq.ctx.Fields
	if len(columns) == 0 {
		columns = loadbalancerstatus.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lbsq.sql != nil {
		selector = lbsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lbsq.ctx.Unique != nil && *lbsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lbsq.predicates {
		p(selector)
	}
	for _, p := range lbsq.order {
		p(selector)
	}
	if offset := lbsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lbsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LoadBalancerStatusGroupBy is the group-by builder for LoadBalancerStatus entities.
type LoadBalancerStatusGroupBy struct {
	selector
	build *LoadBalancerStatusQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lbsgb *LoadBalancerStatusGroupBy) Aggregate(fns ...AggregateFunc) *LoadBalancerStatusGroupBy {
	lbsgb.fns = append(lbsgb.fns, fns...)
	return lbsgb
}

// Scan applies the selector query and scans the result into the given value.
func (lbsgb *LoadBalancerStatusGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lbsgb.build.ctx, "GroupBy")
	if err := lbsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LoadBalancerStatusQuery, *LoadBalancerStatusGroupBy](ctx, lbsgb.build, lbsgb, lbsgb.build.inters, v)
}

func (lbsgb *LoadBalancerStatusGroupBy) sqlScan(ctx context.Context, root *LoadBalancerStatusQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lbsgb.fns))
	for _, fn := range lbsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lbsgb.flds)+len(lbsgb.fns))
		for _, f := range *lbsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lbsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lbsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LoadBalancerStatusSelect is the builder for selecting fields of LoadBalancerStatus entities.
type LoadBalancerStatusSelect struct {
	*LoadBalancerStatusQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (lbss *LoadBalancerStatusSelect) Aggregate(fns ...AggregateFunc) *LoadBalancerStatusSelect {
	lbss.fns = append(lbss.fns, fns...)
	return lbss
}

// Scan applies the selector query and scans the result into the given value.
func (lbss *LoadBalancerStatusSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lbss.ctx, "Select")
	if err := lbss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LoadBalancerStatusQuery, *LoadBalancerStatusSelect](ctx, lbss.LoadBalancerStatusQuery, lbss, lbss.inters, v)
}

func (lbss *LoadBalancerStatusSelect) sqlScan(ctx context.Context, root *LoadBalancerStatusQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(lbss.fns))
	for _, fn := range lbss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*lbss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lbss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

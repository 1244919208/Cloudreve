// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cloudreve/Cloudreve/v4/ent/entity"
	"github.com/cloudreve/Cloudreve/v4/ent/file"
	"github.com/cloudreve/Cloudreve/v4/ent/predicate"
	"github.com/cloudreve/Cloudreve/v4/ent/storagepolicy"
	"github.com/cloudreve/Cloudreve/v4/ent/user"
)

// EntityQuery is the builder for querying Entity entities.
type EntityQuery struct {
	config
	ctx               *QueryContext
	order             []entity.OrderOption
	inters            []Interceptor
	predicates        []predicate.Entity
	withFile          *FileQuery
	withUser          *UserQuery
	withStoragePolicy *StoragePolicyQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EntityQuery builder.
func (eq *EntityQuery) Where(ps ...predicate.Entity) *EntityQuery {
	eq.predicates = append(eq.predicates, ps...)
	return eq
}

// Limit the number of records to be returned by this query.
func (eq *EntityQuery) Limit(limit int) *EntityQuery {
	eq.ctx.Limit = &limit
	return eq
}

// Offset to start from.
func (eq *EntityQuery) Offset(offset int) *EntityQuery {
	eq.ctx.Offset = &offset
	return eq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eq *EntityQuery) Unique(unique bool) *EntityQuery {
	eq.ctx.Unique = &unique
	return eq
}

// Order specifies how the records should be ordered.
func (eq *EntityQuery) Order(o ...entity.OrderOption) *EntityQuery {
	eq.order = append(eq.order, o...)
	return eq
}

// QueryFile chains the current query on the "file" edge.
func (eq *EntityQuery) QueryFile() *FileQuery {
	query := (&FileClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(entity.Table, entity.FieldID, selector),
			sqlgraph.To(file.Table, file.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, entity.FileTable, entity.FilePrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (eq *EntityQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(entity.Table, entity.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, entity.UserTable, entity.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStoragePolicy chains the current query on the "storage_policy" edge.
func (eq *EntityQuery) QueryStoragePolicy() *StoragePolicyQuery {
	query := (&StoragePolicyClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(entity.Table, entity.FieldID, selector),
			sqlgraph.To(storagepolicy.Table, storagepolicy.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, entity.StoragePolicyTable, entity.StoragePolicyColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Entity entity from the query.
// Returns a *NotFoundError when no Entity was found.
func (eq *EntityQuery) First(ctx context.Context) (*Entity, error) {
	nodes, err := eq.Limit(1).All(setContextOp(ctx, eq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{entity.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eq *EntityQuery) FirstX(ctx context.Context) *Entity {
	node, err := eq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Entity ID from the query.
// Returns a *NotFoundError when no Entity ID was found.
func (eq *EntityQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eq.Limit(1).IDs(setContextOp(ctx, eq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{entity.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eq *EntityQuery) FirstIDX(ctx context.Context) int {
	id, err := eq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Entity entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Entity entity is found.
// Returns a *NotFoundError when no Entity entities are found.
func (eq *EntityQuery) Only(ctx context.Context) (*Entity, error) {
	nodes, err := eq.Limit(2).All(setContextOp(ctx, eq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{entity.Label}
	default:
		return nil, &NotSingularError{entity.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eq *EntityQuery) OnlyX(ctx context.Context) *Entity {
	node, err := eq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Entity ID in the query.
// Returns a *NotSingularError when more than one Entity ID is found.
// Returns a *NotFoundError when no entities are found.
func (eq *EntityQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eq.Limit(2).IDs(setContextOp(ctx, eq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{entity.Label}
	default:
		err = &NotSingularError{entity.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eq *EntityQuery) OnlyIDX(ctx context.Context) int {
	id, err := eq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Entities.
func (eq *EntityQuery) All(ctx context.Context) ([]*Entity, error) {
	ctx = setContextOp(ctx, eq.ctx, "All")
	if err := eq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Entity, *EntityQuery]()
	return withInterceptors[[]*Entity](ctx, eq, qr, eq.inters)
}

// AllX is like All, but panics if an error occurs.
func (eq *EntityQuery) AllX(ctx context.Context) []*Entity {
	nodes, err := eq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Entity IDs.
func (eq *EntityQuery) IDs(ctx context.Context) (ids []int, err error) {
	if eq.ctx.Unique == nil && eq.path != nil {
		eq.Unique(true)
	}
	ctx = setContextOp(ctx, eq.ctx, "IDs")
	if err = eq.Select(entity.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eq *EntityQuery) IDsX(ctx context.Context) []int {
	ids, err := eq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eq *EntityQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, eq.ctx, "Count")
	if err := eq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, eq, querierCount[*EntityQuery](), eq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (eq *EntityQuery) CountX(ctx context.Context) int {
	count, err := eq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eq *EntityQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, eq.ctx, "Exist")
	switch _, err := eq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (eq *EntityQuery) ExistX(ctx context.Context) bool {
	exist, err := eq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EntityQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eq *EntityQuery) Clone() *EntityQuery {
	if eq == nil {
		return nil
	}
	return &EntityQuery{
		config:            eq.config,
		ctx:               eq.ctx.Clone(),
		order:             append([]entity.OrderOption{}, eq.order...),
		inters:            append([]Interceptor{}, eq.inters...),
		predicates:        append([]predicate.Entity{}, eq.predicates...),
		withFile:          eq.withFile.Clone(),
		withUser:          eq.withUser.Clone(),
		withStoragePolicy: eq.withStoragePolicy.Clone(),
		// clone intermediate query.
		sql:  eq.sql.Clone(),
		path: eq.path,
	}
}

// WithFile tells the query-builder to eager-load the nodes that are connected to
// the "file" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EntityQuery) WithFile(opts ...func(*FileQuery)) *EntityQuery {
	query := (&FileClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withFile = query
	return eq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EntityQuery) WithUser(opts ...func(*UserQuery)) *EntityQuery {
	query := (&UserClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withUser = query
	return eq
}

// WithStoragePolicy tells the query-builder to eager-load the nodes that are connected to
// the "storage_policy" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EntityQuery) WithStoragePolicy(opts ...func(*StoragePolicyQuery)) *EntityQuery {
	query := (&StoragePolicyClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withStoragePolicy = query
	return eq
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
//	client.Entity.Query().
//		GroupBy(entity.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (eq *EntityQuery) GroupBy(field string, fields ...string) *EntityGroupBy {
	eq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EntityGroupBy{build: eq}
	grbuild.flds = &eq.ctx.Fields
	grbuild.label = entity.Label
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
//	client.Entity.Query().
//		Select(entity.FieldCreatedAt).
//		Scan(ctx, &v)
func (eq *EntityQuery) Select(fields ...string) *EntitySelect {
	eq.ctx.Fields = append(eq.ctx.Fields, fields...)
	sbuild := &EntitySelect{EntityQuery: eq}
	sbuild.label = entity.Label
	sbuild.flds, sbuild.scan = &eq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EntitySelect configured with the given aggregations.
func (eq *EntityQuery) Aggregate(fns ...AggregateFunc) *EntitySelect {
	return eq.Select().Aggregate(fns...)
}

func (eq *EntityQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range eq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, eq); err != nil {
				return err
			}
		}
	}
	for _, f := range eq.ctx.Fields {
		if !entity.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if eq.path != nil {
		prev, err := eq.path(ctx)
		if err != nil {
			return err
		}
		eq.sql = prev
	}
	return nil
}

func (eq *EntityQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Entity, error) {
	var (
		nodes       = []*Entity{}
		_spec       = eq.querySpec()
		loadedTypes = [3]bool{
			eq.withFile != nil,
			eq.withUser != nil,
			eq.withStoragePolicy != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Entity).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Entity{config: eq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := eq.withFile; query != nil {
		if err := eq.loadFile(ctx, query, nodes,
			func(n *Entity) { n.Edges.File = []*File{} },
			func(n *Entity, e *File) { n.Edges.File = append(n.Edges.File, e) }); err != nil {
			return nil, err
		}
	}
	if query := eq.withUser; query != nil {
		if err := eq.loadUser(ctx, query, nodes, nil,
			func(n *Entity, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := eq.withStoragePolicy; query != nil {
		if err := eq.loadStoragePolicy(ctx, query, nodes, nil,
			func(n *Entity, e *StoragePolicy) { n.Edges.StoragePolicy = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (eq *EntityQuery) loadFile(ctx context.Context, query *FileQuery, nodes []*Entity, init func(*Entity), assign func(*Entity, *File)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Entity)
	nids := make(map[int]map[*Entity]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(entity.FileTable)
		s.Join(joinT).On(s.C(file.FieldID), joinT.C(entity.FilePrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(entity.FilePrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(entity.FilePrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Entity]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*File](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "file" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (eq *EntityQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Entity, init func(*Entity), assign func(*Entity, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Entity)
	for i := range nodes {
		fk := nodes[i].CreatedBy
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "created_by" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (eq *EntityQuery) loadStoragePolicy(ctx context.Context, query *StoragePolicyQuery, nodes []*Entity, init func(*Entity), assign func(*Entity, *StoragePolicy)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Entity)
	for i := range nodes {
		fk := nodes[i].StoragePolicyEntities
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(storagepolicy.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "storage_policy_entities" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (eq *EntityQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eq.querySpec()
	_spec.Node.Columns = eq.ctx.Fields
	if len(eq.ctx.Fields) > 0 {
		_spec.Unique = eq.ctx.Unique != nil && *eq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, eq.driver, _spec)
}

func (eq *EntityQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(entity.Table, entity.Columns, sqlgraph.NewFieldSpec(entity.FieldID, field.TypeInt))
	_spec.From = eq.sql
	if unique := eq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if eq.path != nil {
		_spec.Unique = true
	}
	if fields := eq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entity.FieldID)
		for i := range fields {
			if fields[i] != entity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if eq.withUser != nil {
			_spec.Node.AddColumnOnce(entity.FieldCreatedBy)
		}
		if eq.withStoragePolicy != nil {
			_spec.Node.AddColumnOnce(entity.FieldStoragePolicyEntities)
		}
	}
	if ps := eq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eq *EntityQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eq.driver.Dialect())
	t1 := builder.Table(entity.Table)
	columns := eq.ctx.Fields
	if len(columns) == 0 {
		columns = entity.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eq.sql != nil {
		selector = eq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eq.ctx.Unique != nil && *eq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range eq.predicates {
		p(selector)
	}
	for _, p := range eq.order {
		p(selector)
	}
	if offset := eq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EntityGroupBy is the group-by builder for Entity entities.
type EntityGroupBy struct {
	selector
	build *EntityQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (egb *EntityGroupBy) Aggregate(fns ...AggregateFunc) *EntityGroupBy {
	egb.fns = append(egb.fns, fns...)
	return egb
}

// Scan applies the selector query and scans the result into the given value.
func (egb *EntityGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, egb.build.ctx, "GroupBy")
	if err := egb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EntityQuery, *EntityGroupBy](ctx, egb.build, egb, egb.build.inters, v)
}

func (egb *EntityGroupBy) sqlScan(ctx context.Context, root *EntityQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(egb.fns))
	for _, fn := range egb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*egb.flds)+len(egb.fns))
		for _, f := range *egb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*egb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := egb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EntitySelect is the builder for selecting fields of Entity entities.
type EntitySelect struct {
	*EntityQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (es *EntitySelect) Aggregate(fns ...AggregateFunc) *EntitySelect {
	es.fns = append(es.fns, fns...)
	return es
}

// Scan applies the selector query and scans the result into the given value.
func (es *EntitySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, es.ctx, "Select")
	if err := es.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EntityQuery, *EntitySelect](ctx, es.EntityQuery, es, es.inters, v)
}

func (es *EntitySelect) sqlScan(ctx context.Context, root *EntityQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(es.fns))
	for _, fn := range es.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*es.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := es.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

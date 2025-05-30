// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cloudreve/Cloudreve/v4/ent/entity"
	"github.com/cloudreve/Cloudreve/v4/ent/file"
	"github.com/cloudreve/Cloudreve/v4/ent/predicate"
	"github.com/cloudreve/Cloudreve/v4/ent/storagepolicy"
	"github.com/cloudreve/Cloudreve/v4/ent/user"
	"github.com/cloudreve/Cloudreve/v4/inventory/types"
	"github.com/gofrs/uuid"
)

// EntityUpdate is the builder for updating Entity entities.
type EntityUpdate struct {
	config
	hooks    []Hook
	mutation *EntityMutation
}

// Where appends a list predicates to the EntityUpdate builder.
func (eu *EntityUpdate) Where(ps ...predicate.Entity) *EntityUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetUpdatedAt sets the "updated_at" field.
func (eu *EntityUpdate) SetUpdatedAt(t time.Time) *EntityUpdate {
	eu.mutation.SetUpdatedAt(t)
	return eu
}

// SetDeletedAt sets the "deleted_at" field.
func (eu *EntityUpdate) SetDeletedAt(t time.Time) *EntityUpdate {
	eu.mutation.SetDeletedAt(t)
	return eu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (eu *EntityUpdate) SetNillableDeletedAt(t *time.Time) *EntityUpdate {
	if t != nil {
		eu.SetDeletedAt(*t)
	}
	return eu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (eu *EntityUpdate) ClearDeletedAt() *EntityUpdate {
	eu.mutation.ClearDeletedAt()
	return eu
}

// SetType sets the "type" field.
func (eu *EntityUpdate) SetType(i int) *EntityUpdate {
	eu.mutation.ResetType()
	eu.mutation.SetType(i)
	return eu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (eu *EntityUpdate) SetNillableType(i *int) *EntityUpdate {
	if i != nil {
		eu.SetType(*i)
	}
	return eu
}

// AddType adds i to the "type" field.
func (eu *EntityUpdate) AddType(i int) *EntityUpdate {
	eu.mutation.AddType(i)
	return eu
}

// SetSource sets the "source" field.
func (eu *EntityUpdate) SetSource(s string) *EntityUpdate {
	eu.mutation.SetSource(s)
	return eu
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (eu *EntityUpdate) SetNillableSource(s *string) *EntityUpdate {
	if s != nil {
		eu.SetSource(*s)
	}
	return eu
}

// SetSize sets the "size" field.
func (eu *EntityUpdate) SetSize(i int64) *EntityUpdate {
	eu.mutation.ResetSize()
	eu.mutation.SetSize(i)
	return eu
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (eu *EntityUpdate) SetNillableSize(i *int64) *EntityUpdate {
	if i != nil {
		eu.SetSize(*i)
	}
	return eu
}

// AddSize adds i to the "size" field.
func (eu *EntityUpdate) AddSize(i int64) *EntityUpdate {
	eu.mutation.AddSize(i)
	return eu
}

// SetReferenceCount sets the "reference_count" field.
func (eu *EntityUpdate) SetReferenceCount(i int) *EntityUpdate {
	eu.mutation.ResetReferenceCount()
	eu.mutation.SetReferenceCount(i)
	return eu
}

// SetNillableReferenceCount sets the "reference_count" field if the given value is not nil.
func (eu *EntityUpdate) SetNillableReferenceCount(i *int) *EntityUpdate {
	if i != nil {
		eu.SetReferenceCount(*i)
	}
	return eu
}

// AddReferenceCount adds i to the "reference_count" field.
func (eu *EntityUpdate) AddReferenceCount(i int) *EntityUpdate {
	eu.mutation.AddReferenceCount(i)
	return eu
}

// SetStoragePolicyEntities sets the "storage_policy_entities" field.
func (eu *EntityUpdate) SetStoragePolicyEntities(i int) *EntityUpdate {
	eu.mutation.SetStoragePolicyEntities(i)
	return eu
}

// SetNillableStoragePolicyEntities sets the "storage_policy_entities" field if the given value is not nil.
func (eu *EntityUpdate) SetNillableStoragePolicyEntities(i *int) *EntityUpdate {
	if i != nil {
		eu.SetStoragePolicyEntities(*i)
	}
	return eu
}

// SetCreatedBy sets the "created_by" field.
func (eu *EntityUpdate) SetCreatedBy(i int) *EntityUpdate {
	eu.mutation.SetCreatedBy(i)
	return eu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (eu *EntityUpdate) SetNillableCreatedBy(i *int) *EntityUpdate {
	if i != nil {
		eu.SetCreatedBy(*i)
	}
	return eu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (eu *EntityUpdate) ClearCreatedBy() *EntityUpdate {
	eu.mutation.ClearCreatedBy()
	return eu
}

// SetUploadSessionID sets the "upload_session_id" field.
func (eu *EntityUpdate) SetUploadSessionID(u uuid.UUID) *EntityUpdate {
	eu.mutation.SetUploadSessionID(u)
	return eu
}

// SetNillableUploadSessionID sets the "upload_session_id" field if the given value is not nil.
func (eu *EntityUpdate) SetNillableUploadSessionID(u *uuid.UUID) *EntityUpdate {
	if u != nil {
		eu.SetUploadSessionID(*u)
	}
	return eu
}

// ClearUploadSessionID clears the value of the "upload_session_id" field.
func (eu *EntityUpdate) ClearUploadSessionID() *EntityUpdate {
	eu.mutation.ClearUploadSessionID()
	return eu
}

// SetRecycleOptions sets the "recycle_options" field.
func (eu *EntityUpdate) SetRecycleOptions(tro *types.EntityRecycleOption) *EntityUpdate {
	eu.mutation.SetRecycleOptions(tro)
	return eu
}

// ClearRecycleOptions clears the value of the "recycle_options" field.
func (eu *EntityUpdate) ClearRecycleOptions() *EntityUpdate {
	eu.mutation.ClearRecycleOptions()
	return eu
}

// AddFileIDs adds the "file" edge to the File entity by IDs.
func (eu *EntityUpdate) AddFileIDs(ids ...int) *EntityUpdate {
	eu.mutation.AddFileIDs(ids...)
	return eu
}

// AddFile adds the "file" edges to the File entity.
func (eu *EntityUpdate) AddFile(f ...*File) *EntityUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return eu.AddFileIDs(ids...)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (eu *EntityUpdate) SetUserID(id int) *EntityUpdate {
	eu.mutation.SetUserID(id)
	return eu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (eu *EntityUpdate) SetNillableUserID(id *int) *EntityUpdate {
	if id != nil {
		eu = eu.SetUserID(*id)
	}
	return eu
}

// SetUser sets the "user" edge to the User entity.
func (eu *EntityUpdate) SetUser(u *User) *EntityUpdate {
	return eu.SetUserID(u.ID)
}

// SetStoragePolicyID sets the "storage_policy" edge to the StoragePolicy entity by ID.
func (eu *EntityUpdate) SetStoragePolicyID(id int) *EntityUpdate {
	eu.mutation.SetStoragePolicyID(id)
	return eu
}

// SetStoragePolicy sets the "storage_policy" edge to the StoragePolicy entity.
func (eu *EntityUpdate) SetStoragePolicy(s *StoragePolicy) *EntityUpdate {
	return eu.SetStoragePolicyID(s.ID)
}

// Mutation returns the EntityMutation object of the builder.
func (eu *EntityUpdate) Mutation() *EntityMutation {
	return eu.mutation
}

// ClearFile clears all "file" edges to the File entity.
func (eu *EntityUpdate) ClearFile() *EntityUpdate {
	eu.mutation.ClearFile()
	return eu
}

// RemoveFileIDs removes the "file" edge to File entities by IDs.
func (eu *EntityUpdate) RemoveFileIDs(ids ...int) *EntityUpdate {
	eu.mutation.RemoveFileIDs(ids...)
	return eu
}

// RemoveFile removes "file" edges to File entities.
func (eu *EntityUpdate) RemoveFile(f ...*File) *EntityUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return eu.RemoveFileIDs(ids...)
}

// ClearUser clears the "user" edge to the User entity.
func (eu *EntityUpdate) ClearUser() *EntityUpdate {
	eu.mutation.ClearUser()
	return eu
}

// ClearStoragePolicy clears the "storage_policy" edge to the StoragePolicy entity.
func (eu *EntityUpdate) ClearStoragePolicy() *EntityUpdate {
	eu.mutation.ClearStoragePolicy()
	return eu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EntityUpdate) Save(ctx context.Context) (int, error) {
	if err := eu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EntityUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EntityUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EntityUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eu *EntityUpdate) defaults() error {
	if _, ok := eu.mutation.UpdatedAt(); !ok {
		if entity.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized entity.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := entity.UpdateDefaultUpdatedAt()
		eu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (eu *EntityUpdate) check() error {
	if _, ok := eu.mutation.StoragePolicyID(); eu.mutation.StoragePolicyCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Entity.storage_policy"`)
	}
	return nil
}

func (eu *EntityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(entity.Table, entity.Columns, sqlgraph.NewFieldSpec(entity.FieldID, field.TypeInt))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.UpdatedAt(); ok {
		_spec.SetField(entity.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := eu.mutation.DeletedAt(); ok {
		_spec.SetField(entity.FieldDeletedAt, field.TypeTime, value)
	}
	if eu.mutation.DeletedAtCleared() {
		_spec.ClearField(entity.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := eu.mutation.GetType(); ok {
		_spec.SetField(entity.FieldType, field.TypeInt, value)
	}
	if value, ok := eu.mutation.AddedType(); ok {
		_spec.AddField(entity.FieldType, field.TypeInt, value)
	}
	if value, ok := eu.mutation.Source(); ok {
		_spec.SetField(entity.FieldSource, field.TypeString, value)
	}
	if value, ok := eu.mutation.Size(); ok {
		_spec.SetField(entity.FieldSize, field.TypeInt64, value)
	}
	if value, ok := eu.mutation.AddedSize(); ok {
		_spec.AddField(entity.FieldSize, field.TypeInt64, value)
	}
	if value, ok := eu.mutation.ReferenceCount(); ok {
		_spec.SetField(entity.FieldReferenceCount, field.TypeInt, value)
	}
	if value, ok := eu.mutation.AddedReferenceCount(); ok {
		_spec.AddField(entity.FieldReferenceCount, field.TypeInt, value)
	}
	if value, ok := eu.mutation.UploadSessionID(); ok {
		_spec.SetField(entity.FieldUploadSessionID, field.TypeUUID, value)
	}
	if eu.mutation.UploadSessionIDCleared() {
		_spec.ClearField(entity.FieldUploadSessionID, field.TypeUUID)
	}
	if value, ok := eu.mutation.RecycleOptions(); ok {
		_spec.SetField(entity.FieldRecycleOptions, field.TypeJSON, value)
	}
	if eu.mutation.RecycleOptionsCleared() {
		_spec.ClearField(entity.FieldRecycleOptions, field.TypeJSON)
	}
	if eu.mutation.FileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   entity.FileTable,
			Columns: entity.FilePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedFileIDs(); len(nodes) > 0 && !eu.mutation.FileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   entity.FileTable,
			Columns: entity.FilePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.FileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   entity.FileTable,
			Columns: entity.FilePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entity.UserTable,
			Columns: []string{entity.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entity.UserTable,
			Columns: []string{entity.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.StoragePolicyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entity.StoragePolicyTable,
			Columns: []string{entity.StoragePolicyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(storagepolicy.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.StoragePolicyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entity.StoragePolicyTable,
			Columns: []string{entity.StoragePolicyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(storagepolicy.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EntityUpdateOne is the builder for updating a single Entity entity.
type EntityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EntityMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (euo *EntityUpdateOne) SetUpdatedAt(t time.Time) *EntityUpdateOne {
	euo.mutation.SetUpdatedAt(t)
	return euo
}

// SetDeletedAt sets the "deleted_at" field.
func (euo *EntityUpdateOne) SetDeletedAt(t time.Time) *EntityUpdateOne {
	euo.mutation.SetDeletedAt(t)
	return euo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (euo *EntityUpdateOne) SetNillableDeletedAt(t *time.Time) *EntityUpdateOne {
	if t != nil {
		euo.SetDeletedAt(*t)
	}
	return euo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (euo *EntityUpdateOne) ClearDeletedAt() *EntityUpdateOne {
	euo.mutation.ClearDeletedAt()
	return euo
}

// SetType sets the "type" field.
func (euo *EntityUpdateOne) SetType(i int) *EntityUpdateOne {
	euo.mutation.ResetType()
	euo.mutation.SetType(i)
	return euo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (euo *EntityUpdateOne) SetNillableType(i *int) *EntityUpdateOne {
	if i != nil {
		euo.SetType(*i)
	}
	return euo
}

// AddType adds i to the "type" field.
func (euo *EntityUpdateOne) AddType(i int) *EntityUpdateOne {
	euo.mutation.AddType(i)
	return euo
}

// SetSource sets the "source" field.
func (euo *EntityUpdateOne) SetSource(s string) *EntityUpdateOne {
	euo.mutation.SetSource(s)
	return euo
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (euo *EntityUpdateOne) SetNillableSource(s *string) *EntityUpdateOne {
	if s != nil {
		euo.SetSource(*s)
	}
	return euo
}

// SetSize sets the "size" field.
func (euo *EntityUpdateOne) SetSize(i int64) *EntityUpdateOne {
	euo.mutation.ResetSize()
	euo.mutation.SetSize(i)
	return euo
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (euo *EntityUpdateOne) SetNillableSize(i *int64) *EntityUpdateOne {
	if i != nil {
		euo.SetSize(*i)
	}
	return euo
}

// AddSize adds i to the "size" field.
func (euo *EntityUpdateOne) AddSize(i int64) *EntityUpdateOne {
	euo.mutation.AddSize(i)
	return euo
}

// SetReferenceCount sets the "reference_count" field.
func (euo *EntityUpdateOne) SetReferenceCount(i int) *EntityUpdateOne {
	euo.mutation.ResetReferenceCount()
	euo.mutation.SetReferenceCount(i)
	return euo
}

// SetNillableReferenceCount sets the "reference_count" field if the given value is not nil.
func (euo *EntityUpdateOne) SetNillableReferenceCount(i *int) *EntityUpdateOne {
	if i != nil {
		euo.SetReferenceCount(*i)
	}
	return euo
}

// AddReferenceCount adds i to the "reference_count" field.
func (euo *EntityUpdateOne) AddReferenceCount(i int) *EntityUpdateOne {
	euo.mutation.AddReferenceCount(i)
	return euo
}

// SetStoragePolicyEntities sets the "storage_policy_entities" field.
func (euo *EntityUpdateOne) SetStoragePolicyEntities(i int) *EntityUpdateOne {
	euo.mutation.SetStoragePolicyEntities(i)
	return euo
}

// SetNillableStoragePolicyEntities sets the "storage_policy_entities" field if the given value is not nil.
func (euo *EntityUpdateOne) SetNillableStoragePolicyEntities(i *int) *EntityUpdateOne {
	if i != nil {
		euo.SetStoragePolicyEntities(*i)
	}
	return euo
}

// SetCreatedBy sets the "created_by" field.
func (euo *EntityUpdateOne) SetCreatedBy(i int) *EntityUpdateOne {
	euo.mutation.SetCreatedBy(i)
	return euo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (euo *EntityUpdateOne) SetNillableCreatedBy(i *int) *EntityUpdateOne {
	if i != nil {
		euo.SetCreatedBy(*i)
	}
	return euo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (euo *EntityUpdateOne) ClearCreatedBy() *EntityUpdateOne {
	euo.mutation.ClearCreatedBy()
	return euo
}

// SetUploadSessionID sets the "upload_session_id" field.
func (euo *EntityUpdateOne) SetUploadSessionID(u uuid.UUID) *EntityUpdateOne {
	euo.mutation.SetUploadSessionID(u)
	return euo
}

// SetNillableUploadSessionID sets the "upload_session_id" field if the given value is not nil.
func (euo *EntityUpdateOne) SetNillableUploadSessionID(u *uuid.UUID) *EntityUpdateOne {
	if u != nil {
		euo.SetUploadSessionID(*u)
	}
	return euo
}

// ClearUploadSessionID clears the value of the "upload_session_id" field.
func (euo *EntityUpdateOne) ClearUploadSessionID() *EntityUpdateOne {
	euo.mutation.ClearUploadSessionID()
	return euo
}

// SetRecycleOptions sets the "recycle_options" field.
func (euo *EntityUpdateOne) SetRecycleOptions(tro *types.EntityRecycleOption) *EntityUpdateOne {
	euo.mutation.SetRecycleOptions(tro)
	return euo
}

// ClearRecycleOptions clears the value of the "recycle_options" field.
func (euo *EntityUpdateOne) ClearRecycleOptions() *EntityUpdateOne {
	euo.mutation.ClearRecycleOptions()
	return euo
}

// AddFileIDs adds the "file" edge to the File entity by IDs.
func (euo *EntityUpdateOne) AddFileIDs(ids ...int) *EntityUpdateOne {
	euo.mutation.AddFileIDs(ids...)
	return euo
}

// AddFile adds the "file" edges to the File entity.
func (euo *EntityUpdateOne) AddFile(f ...*File) *EntityUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return euo.AddFileIDs(ids...)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (euo *EntityUpdateOne) SetUserID(id int) *EntityUpdateOne {
	euo.mutation.SetUserID(id)
	return euo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (euo *EntityUpdateOne) SetNillableUserID(id *int) *EntityUpdateOne {
	if id != nil {
		euo = euo.SetUserID(*id)
	}
	return euo
}

// SetUser sets the "user" edge to the User entity.
func (euo *EntityUpdateOne) SetUser(u *User) *EntityUpdateOne {
	return euo.SetUserID(u.ID)
}

// SetStoragePolicyID sets the "storage_policy" edge to the StoragePolicy entity by ID.
func (euo *EntityUpdateOne) SetStoragePolicyID(id int) *EntityUpdateOne {
	euo.mutation.SetStoragePolicyID(id)
	return euo
}

// SetStoragePolicy sets the "storage_policy" edge to the StoragePolicy entity.
func (euo *EntityUpdateOne) SetStoragePolicy(s *StoragePolicy) *EntityUpdateOne {
	return euo.SetStoragePolicyID(s.ID)
}

// Mutation returns the EntityMutation object of the builder.
func (euo *EntityUpdateOne) Mutation() *EntityMutation {
	return euo.mutation
}

// ClearFile clears all "file" edges to the File entity.
func (euo *EntityUpdateOne) ClearFile() *EntityUpdateOne {
	euo.mutation.ClearFile()
	return euo
}

// RemoveFileIDs removes the "file" edge to File entities by IDs.
func (euo *EntityUpdateOne) RemoveFileIDs(ids ...int) *EntityUpdateOne {
	euo.mutation.RemoveFileIDs(ids...)
	return euo
}

// RemoveFile removes "file" edges to File entities.
func (euo *EntityUpdateOne) RemoveFile(f ...*File) *EntityUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return euo.RemoveFileIDs(ids...)
}

// ClearUser clears the "user" edge to the User entity.
func (euo *EntityUpdateOne) ClearUser() *EntityUpdateOne {
	euo.mutation.ClearUser()
	return euo
}

// ClearStoragePolicy clears the "storage_policy" edge to the StoragePolicy entity.
func (euo *EntityUpdateOne) ClearStoragePolicy() *EntityUpdateOne {
	euo.mutation.ClearStoragePolicy()
	return euo
}

// Where appends a list predicates to the EntityUpdate builder.
func (euo *EntityUpdateOne) Where(ps ...predicate.Entity) *EntityUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EntityUpdateOne) Select(field string, fields ...string) *EntityUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Entity entity.
func (euo *EntityUpdateOne) Save(ctx context.Context) (*Entity, error) {
	if err := euo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EntityUpdateOne) SaveX(ctx context.Context) *Entity {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EntityUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EntityUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (euo *EntityUpdateOne) defaults() error {
	if _, ok := euo.mutation.UpdatedAt(); !ok {
		if entity.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized entity.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := entity.UpdateDefaultUpdatedAt()
		euo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (euo *EntityUpdateOne) check() error {
	if _, ok := euo.mutation.StoragePolicyID(); euo.mutation.StoragePolicyCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Entity.storage_policy"`)
	}
	return nil
}

func (euo *EntityUpdateOne) sqlSave(ctx context.Context) (_node *Entity, err error) {
	if err := euo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(entity.Table, entity.Columns, sqlgraph.NewFieldSpec(entity.FieldID, field.TypeInt))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Entity.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entity.FieldID)
		for _, f := range fields {
			if !entity.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != entity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.UpdatedAt(); ok {
		_spec.SetField(entity.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := euo.mutation.DeletedAt(); ok {
		_spec.SetField(entity.FieldDeletedAt, field.TypeTime, value)
	}
	if euo.mutation.DeletedAtCleared() {
		_spec.ClearField(entity.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := euo.mutation.GetType(); ok {
		_spec.SetField(entity.FieldType, field.TypeInt, value)
	}
	if value, ok := euo.mutation.AddedType(); ok {
		_spec.AddField(entity.FieldType, field.TypeInt, value)
	}
	if value, ok := euo.mutation.Source(); ok {
		_spec.SetField(entity.FieldSource, field.TypeString, value)
	}
	if value, ok := euo.mutation.Size(); ok {
		_spec.SetField(entity.FieldSize, field.TypeInt64, value)
	}
	if value, ok := euo.mutation.AddedSize(); ok {
		_spec.AddField(entity.FieldSize, field.TypeInt64, value)
	}
	if value, ok := euo.mutation.ReferenceCount(); ok {
		_spec.SetField(entity.FieldReferenceCount, field.TypeInt, value)
	}
	if value, ok := euo.mutation.AddedReferenceCount(); ok {
		_spec.AddField(entity.FieldReferenceCount, field.TypeInt, value)
	}
	if value, ok := euo.mutation.UploadSessionID(); ok {
		_spec.SetField(entity.FieldUploadSessionID, field.TypeUUID, value)
	}
	if euo.mutation.UploadSessionIDCleared() {
		_spec.ClearField(entity.FieldUploadSessionID, field.TypeUUID)
	}
	if value, ok := euo.mutation.RecycleOptions(); ok {
		_spec.SetField(entity.FieldRecycleOptions, field.TypeJSON, value)
	}
	if euo.mutation.RecycleOptionsCleared() {
		_spec.ClearField(entity.FieldRecycleOptions, field.TypeJSON)
	}
	if euo.mutation.FileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   entity.FileTable,
			Columns: entity.FilePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedFileIDs(); len(nodes) > 0 && !euo.mutation.FileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   entity.FileTable,
			Columns: entity.FilePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.FileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   entity.FileTable,
			Columns: entity.FilePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entity.UserTable,
			Columns: []string{entity.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entity.UserTable,
			Columns: []string{entity.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.StoragePolicyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entity.StoragePolicyTable,
			Columns: []string{entity.StoragePolicyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(storagepolicy.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.StoragePolicyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entity.StoragePolicyTable,
			Columns: []string{entity.StoragePolicyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(storagepolicy.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Entity{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}

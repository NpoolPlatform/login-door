// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/login-door/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/login-door/pkg/db/ent/provider"
)

// ProviderUpdate is the builder for updating Provider entities.
type ProviderUpdate struct {
	config
	hooks    []Hook
	mutation *ProviderMutation
}

// Where appends a list predicates to the ProviderUpdate builder.
func (pu *ProviderUpdate) Where(ps ...predicate.Provider) *ProviderUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetProviderName sets the "provider_name" field.
func (pu *ProviderUpdate) SetProviderName(s string) *ProviderUpdate {
	pu.mutation.SetProviderName(s)
	return pu
}

// SetClientID sets the "client_id" field.
func (pu *ProviderUpdate) SetClientID(s string) *ProviderUpdate {
	pu.mutation.SetClientID(s)
	return pu
}

// SetClientSecret sets the "client_secret" field.
func (pu *ProviderUpdate) SetClientSecret(s string) *ProviderUpdate {
	pu.mutation.SetClientSecret(s)
	return pu
}

// SetProviderURL sets the "provider_url" field.
func (pu *ProviderUpdate) SetProviderURL(s string) *ProviderUpdate {
	pu.mutation.SetProviderURL(s)
	return pu
}

// SetProviderLogo sets the "provider_logo" field.
func (pu *ProviderUpdate) SetProviderLogo(s string) *ProviderUpdate {
	pu.mutation.SetProviderLogo(s)
	return pu
}

// SetCreateAt sets the "create_at" field.
func (pu *ProviderUpdate) SetCreateAt(u uint32) *ProviderUpdate {
	pu.mutation.ResetCreateAt()
	pu.mutation.SetCreateAt(u)
	return pu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (pu *ProviderUpdate) SetNillableCreateAt(u *uint32) *ProviderUpdate {
	if u != nil {
		pu.SetCreateAt(*u)
	}
	return pu
}

// AddCreateAt adds u to the "create_at" field.
func (pu *ProviderUpdate) AddCreateAt(u uint32) *ProviderUpdate {
	pu.mutation.AddCreateAt(u)
	return pu
}

// SetUpdateAt sets the "update_at" field.
func (pu *ProviderUpdate) SetUpdateAt(u uint32) *ProviderUpdate {
	pu.mutation.ResetUpdateAt()
	pu.mutation.SetUpdateAt(u)
	return pu
}

// AddUpdateAt adds u to the "update_at" field.
func (pu *ProviderUpdate) AddUpdateAt(u uint32) *ProviderUpdate {
	pu.mutation.AddUpdateAt(u)
	return pu
}

// SetDeleteAt sets the "delete_at" field.
func (pu *ProviderUpdate) SetDeleteAt(u uint32) *ProviderUpdate {
	pu.mutation.ResetDeleteAt()
	pu.mutation.SetDeleteAt(u)
	return pu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (pu *ProviderUpdate) SetNillableDeleteAt(u *uint32) *ProviderUpdate {
	if u != nil {
		pu.SetDeleteAt(*u)
	}
	return pu
}

// AddDeleteAt adds u to the "delete_at" field.
func (pu *ProviderUpdate) AddDeleteAt(u uint32) *ProviderUpdate {
	pu.mutation.AddDeleteAt(u)
	return pu
}

// Mutation returns the ProviderMutation object of the builder.
func (pu *ProviderUpdate) Mutation() *ProviderMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProviderUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pu.defaults()
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProviderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProviderUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProviderUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProviderUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *ProviderUpdate) defaults() {
	if _, ok := pu.mutation.UpdateAt(); !ok {
		v := provider.UpdateDefaultUpdateAt()
		pu.mutation.SetUpdateAt(v)
	}
}

func (pu *ProviderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   provider.Table,
			Columns: provider.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: provider.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.ProviderName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: provider.FieldProviderName,
		})
	}
	if value, ok := pu.mutation.ClientID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: provider.FieldClientID,
		})
	}
	if value, ok := pu.mutation.ClientSecret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: provider.FieldClientSecret,
		})
	}
	if value, ok := pu.mutation.ProviderURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: provider.FieldProviderURL,
		})
	}
	if value, ok := pu.mutation.ProviderLogo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: provider.FieldProviderLogo,
		})
	}
	if value, ok := pu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: provider.FieldCreateAt,
		})
	}
	if value, ok := pu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: provider.FieldCreateAt,
		})
	}
	if value, ok := pu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: provider.FieldUpdateAt,
		})
	}
	if value, ok := pu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: provider.FieldUpdateAt,
		})
	}
	if value, ok := pu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: provider.FieldDeleteAt,
		})
	}
	if value, ok := pu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: provider.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{provider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ProviderUpdateOne is the builder for updating a single Provider entity.
type ProviderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProviderMutation
}

// SetProviderName sets the "provider_name" field.
func (puo *ProviderUpdateOne) SetProviderName(s string) *ProviderUpdateOne {
	puo.mutation.SetProviderName(s)
	return puo
}

// SetClientID sets the "client_id" field.
func (puo *ProviderUpdateOne) SetClientID(s string) *ProviderUpdateOne {
	puo.mutation.SetClientID(s)
	return puo
}

// SetClientSecret sets the "client_secret" field.
func (puo *ProviderUpdateOne) SetClientSecret(s string) *ProviderUpdateOne {
	puo.mutation.SetClientSecret(s)
	return puo
}

// SetProviderURL sets the "provider_url" field.
func (puo *ProviderUpdateOne) SetProviderURL(s string) *ProviderUpdateOne {
	puo.mutation.SetProviderURL(s)
	return puo
}

// SetProviderLogo sets the "provider_logo" field.
func (puo *ProviderUpdateOne) SetProviderLogo(s string) *ProviderUpdateOne {
	puo.mutation.SetProviderLogo(s)
	return puo
}

// SetCreateAt sets the "create_at" field.
func (puo *ProviderUpdateOne) SetCreateAt(u uint32) *ProviderUpdateOne {
	puo.mutation.ResetCreateAt()
	puo.mutation.SetCreateAt(u)
	return puo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (puo *ProviderUpdateOne) SetNillableCreateAt(u *uint32) *ProviderUpdateOne {
	if u != nil {
		puo.SetCreateAt(*u)
	}
	return puo
}

// AddCreateAt adds u to the "create_at" field.
func (puo *ProviderUpdateOne) AddCreateAt(u uint32) *ProviderUpdateOne {
	puo.mutation.AddCreateAt(u)
	return puo
}

// SetUpdateAt sets the "update_at" field.
func (puo *ProviderUpdateOne) SetUpdateAt(u uint32) *ProviderUpdateOne {
	puo.mutation.ResetUpdateAt()
	puo.mutation.SetUpdateAt(u)
	return puo
}

// AddUpdateAt adds u to the "update_at" field.
func (puo *ProviderUpdateOne) AddUpdateAt(u uint32) *ProviderUpdateOne {
	puo.mutation.AddUpdateAt(u)
	return puo
}

// SetDeleteAt sets the "delete_at" field.
func (puo *ProviderUpdateOne) SetDeleteAt(u uint32) *ProviderUpdateOne {
	puo.mutation.ResetDeleteAt()
	puo.mutation.SetDeleteAt(u)
	return puo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (puo *ProviderUpdateOne) SetNillableDeleteAt(u *uint32) *ProviderUpdateOne {
	if u != nil {
		puo.SetDeleteAt(*u)
	}
	return puo
}

// AddDeleteAt adds u to the "delete_at" field.
func (puo *ProviderUpdateOne) AddDeleteAt(u uint32) *ProviderUpdateOne {
	puo.mutation.AddDeleteAt(u)
	return puo
}

// Mutation returns the ProviderMutation object of the builder.
func (puo *ProviderUpdateOne) Mutation() *ProviderMutation {
	return puo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProviderUpdateOne) Select(field string, fields ...string) *ProviderUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Provider entity.
func (puo *ProviderUpdateOne) Save(ctx context.Context) (*Provider, error) {
	var (
		err  error
		node *Provider
	)
	puo.defaults()
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProviderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProviderUpdateOne) SaveX(ctx context.Context) *Provider {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProviderUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProviderUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *ProviderUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdateAt(); !ok {
		v := provider.UpdateDefaultUpdateAt()
		puo.mutation.SetUpdateAt(v)
	}
}

func (puo *ProviderUpdateOne) sqlSave(ctx context.Context) (_node *Provider, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   provider.Table,
			Columns: provider.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: provider.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Provider.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, provider.FieldID)
		for _, f := range fields {
			if !provider.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != provider.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.ProviderName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: provider.FieldProviderName,
		})
	}
	if value, ok := puo.mutation.ClientID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: provider.FieldClientID,
		})
	}
	if value, ok := puo.mutation.ClientSecret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: provider.FieldClientSecret,
		})
	}
	if value, ok := puo.mutation.ProviderURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: provider.FieldProviderURL,
		})
	}
	if value, ok := puo.mutation.ProviderLogo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: provider.FieldProviderLogo,
		})
	}
	if value, ok := puo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: provider.FieldCreateAt,
		})
	}
	if value, ok := puo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: provider.FieldCreateAt,
		})
	}
	if value, ok := puo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: provider.FieldUpdateAt,
		})
	}
	if value, ok := puo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: provider.FieldUpdateAt,
		})
	}
	if value, ok := puo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: provider.FieldDeleteAt,
		})
	}
	if value, ok := puo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: provider.FieldDeleteAt,
		})
	}
	_node = &Provider{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{provider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
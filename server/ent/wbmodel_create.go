// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/kallydev/privacy/ent/wbmodel"
)

// WBModelCreate is the builder for creating a WBModel entity.
type WBModelCreate struct {
	config
	mutation *WBModelMutation
	hooks    []Hook
}

// SetWbNumber sets the uid field.
func (wmc *WBModelCreate) SetWbNumber(i int64) *WBModelCreate {
	wmc.mutation.SetWbNumber(i)
	return wmc
}

// SetPhoneNumber sets the phone_number field.
func (wmc *WBModelCreate) SetPhoneNumber(i int64) *WBModelCreate {
	wmc.mutation.SetPhoneNumber(i)
	return wmc
}

// Mutation returns the WBModelMutation object of the builder.
func (wmc *WBModelCreate) Mutation() *WBModelMutation {
	return wmc.mutation
}

// Save creates the WBModel in the database.
func (wmc *WBModelCreate) Save(ctx context.Context) (*WBModel, error) {
	var (
		err  error
		node *WBModel
	)
	if len(wmc.hooks) == 0 {
		if err = wmc.check(); err != nil {
			return nil, err
		}
		node, err = wmc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WBModelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wmc.check(); err != nil {
				return nil, err
			}
			wmc.mutation = mutation
			node, err = wmc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wmc.hooks) - 1; i >= 0; i-- {
			mut = wmc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wmc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wmc *WBModelCreate) SaveX(ctx context.Context) *WBModel {
	v, err := wmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (wmc *WBModelCreate) check() error {
	if _, ok := wmc.mutation.WbNumber(); !ok {
		return &ValidationError{Name: "uid", err: errors.New("ent: missing required field \"uid\"")}
	}
	if _, ok := wmc.mutation.PhoneNumber(); !ok {
		return &ValidationError{Name: "phone_number", err: errors.New("ent: missing required field \"phone_number\"")}
	}
	return nil
}

func (wmc *WBModelCreate) sqlSave(ctx context.Context) (*WBModel, error) {
	_node, _spec := wmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wmc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (wmc *WBModelCreate) createSpec() (*WBModel, *sqlgraph.CreateSpec) {
	var (
		_node = &WBModel{config: wmc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: wbmodel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: wbmodel.FieldID,
			},
		}
	)
	if value, ok := wmc.mutation.WbNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: wbmodel.FieldWbNumber,
		})
		_node.WbNumber = value
	}
	if value, ok := wmc.mutation.PhoneNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: wbmodel.FieldPhoneNumber,
		})
		_node.PhoneNumber = value
	}
	return _node, _spec
}

// WBModelCreateBulk is the builder for creating a bulk of WBModel entities.
type WBModelCreateBulk struct {
	config
	builders []*WBModelCreate
}

// Save creates the WBModel entities in the database.
func (wmcb *WBModelCreateBulk) Save(ctx context.Context) ([]*WBModel, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wmcb.builders))
	nodes := make([]*WBModel, len(wmcb.builders))
	mutators := make([]Mutator, len(wmcb.builders))
	for i := range wmcb.builders {
		func(i int, root context.Context) {
			builder := wmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WBModelMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wmcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wmcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (wmcb *WBModelCreateBulk) SaveX(ctx context.Context) []*WBModel {
	v, err := wmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
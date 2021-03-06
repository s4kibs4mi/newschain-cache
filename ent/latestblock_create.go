// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/s4kibs4mi/newschain-cache/ent/latestblock"
)

// LatestBlockCreate is the builder for creating a LatestBlock entity.
type LatestBlockCreate struct {
	config
	mutation *LatestBlockMutation
	hooks    []Hook
}

// SetBlockNumber sets the "block_number" field.
func (lbc *LatestBlockCreate) SetBlockNumber(u uint32) *LatestBlockCreate {
	lbc.mutation.SetBlockNumber(u)
	return lbc
}

// Mutation returns the LatestBlockMutation object of the builder.
func (lbc *LatestBlockCreate) Mutation() *LatestBlockMutation {
	return lbc.mutation
}

// Save creates the LatestBlock in the database.
func (lbc *LatestBlockCreate) Save(ctx context.Context) (*LatestBlock, error) {
	var (
		err  error
		node *LatestBlock
	)
	if len(lbc.hooks) == 0 {
		if err = lbc.check(); err != nil {
			return nil, err
		}
		node, err = lbc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LatestBlockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lbc.check(); err != nil {
				return nil, err
			}
			lbc.mutation = mutation
			node, err = lbc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(lbc.hooks) - 1; i >= 0; i-- {
			mut = lbc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lbc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lbc *LatestBlockCreate) SaveX(ctx context.Context) *LatestBlock {
	v, err := lbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (lbc *LatestBlockCreate) check() error {
	if _, ok := lbc.mutation.BlockNumber(); !ok {
		return &ValidationError{Name: "block_number", err: errors.New("ent: missing required field \"block_number\"")}
	}
	return nil
}

func (lbc *LatestBlockCreate) sqlSave(ctx context.Context) (*LatestBlock, error) {
	_node, _spec := lbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lbc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (lbc *LatestBlockCreate) createSpec() (*LatestBlock, *sqlgraph.CreateSpec) {
	var (
		_node = &LatestBlock{config: lbc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: latestblock.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: latestblock.FieldID,
			},
		}
	)
	if value, ok := lbc.mutation.BlockNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: latestblock.FieldBlockNumber,
		})
		_node.BlockNumber = value
	}
	return _node, _spec
}

// LatestBlockCreateBulk is the builder for creating many LatestBlock entities in bulk.
type LatestBlockCreateBulk struct {
	config
	builders []*LatestBlockCreate
}

// Save creates the LatestBlock entities in the database.
func (lbcb *LatestBlockCreateBulk) Save(ctx context.Context) ([]*LatestBlock, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lbcb.builders))
	nodes := make([]*LatestBlock, len(lbcb.builders))
	mutators := make([]Mutator, len(lbcb.builders))
	for i := range lbcb.builders {
		func(i int, root context.Context) {
			builder := lbcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LatestBlockMutation)
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
					_, err = mutators[i+1].Mutate(root, lbcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lbcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lbcb *LatestBlockCreateBulk) SaveX(ctx context.Context) []*LatestBlock {
	v, err := lbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

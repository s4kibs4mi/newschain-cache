// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/s4kibs4mi/newschain-cache/ent/latestblock"
	"github.com/s4kibs4mi/newschain-cache/ent/predicate"
)

// LatestBlockUpdate is the builder for updating LatestBlock entities.
type LatestBlockUpdate struct {
	config
	hooks    []Hook
	mutation *LatestBlockMutation
}

// Where adds a new predicate for the LatestBlockUpdate builder.
func (lbu *LatestBlockUpdate) Where(ps ...predicate.LatestBlock) *LatestBlockUpdate {
	lbu.mutation.predicates = append(lbu.mutation.predicates, ps...)
	return lbu
}

// SetBlockNumber sets the "block_number" field.
func (lbu *LatestBlockUpdate) SetBlockNumber(u uint32) *LatestBlockUpdate {
	lbu.mutation.ResetBlockNumber()
	lbu.mutation.SetBlockNumber(u)
	return lbu
}

// AddBlockNumber adds u to the "block_number" field.
func (lbu *LatestBlockUpdate) AddBlockNumber(u uint32) *LatestBlockUpdate {
	lbu.mutation.AddBlockNumber(u)
	return lbu
}

// Mutation returns the LatestBlockMutation object of the builder.
func (lbu *LatestBlockUpdate) Mutation() *LatestBlockMutation {
	return lbu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lbu *LatestBlockUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lbu.hooks) == 0 {
		affected, err = lbu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LatestBlockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lbu.mutation = mutation
			affected, err = lbu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lbu.hooks) - 1; i >= 0; i-- {
			mut = lbu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lbu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lbu *LatestBlockUpdate) SaveX(ctx context.Context) int {
	affected, err := lbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lbu *LatestBlockUpdate) Exec(ctx context.Context) error {
	_, err := lbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lbu *LatestBlockUpdate) ExecX(ctx context.Context) {
	if err := lbu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lbu *LatestBlockUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   latestblock.Table,
			Columns: latestblock.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: latestblock.FieldID,
			},
		},
	}
	if ps := lbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lbu.mutation.BlockNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: latestblock.FieldBlockNumber,
		})
	}
	if value, ok := lbu.mutation.AddedBlockNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: latestblock.FieldBlockNumber,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{latestblock.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// LatestBlockUpdateOne is the builder for updating a single LatestBlock entity.
type LatestBlockUpdateOne struct {
	config
	hooks    []Hook
	mutation *LatestBlockMutation
}

// SetBlockNumber sets the "block_number" field.
func (lbuo *LatestBlockUpdateOne) SetBlockNumber(u uint32) *LatestBlockUpdateOne {
	lbuo.mutation.ResetBlockNumber()
	lbuo.mutation.SetBlockNumber(u)
	return lbuo
}

// AddBlockNumber adds u to the "block_number" field.
func (lbuo *LatestBlockUpdateOne) AddBlockNumber(u uint32) *LatestBlockUpdateOne {
	lbuo.mutation.AddBlockNumber(u)
	return lbuo
}

// Mutation returns the LatestBlockMutation object of the builder.
func (lbuo *LatestBlockUpdateOne) Mutation() *LatestBlockMutation {
	return lbuo.mutation
}

// Save executes the query and returns the updated LatestBlock entity.
func (lbuo *LatestBlockUpdateOne) Save(ctx context.Context) (*LatestBlock, error) {
	var (
		err  error
		node *LatestBlock
	)
	if len(lbuo.hooks) == 0 {
		node, err = lbuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LatestBlockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lbuo.mutation = mutation
			node, err = lbuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(lbuo.hooks) - 1; i >= 0; i-- {
			mut = lbuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lbuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (lbuo *LatestBlockUpdateOne) SaveX(ctx context.Context) *LatestBlock {
	node, err := lbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lbuo *LatestBlockUpdateOne) Exec(ctx context.Context) error {
	_, err := lbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lbuo *LatestBlockUpdateOne) ExecX(ctx context.Context) {
	if err := lbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lbuo *LatestBlockUpdateOne) sqlSave(ctx context.Context) (_node *LatestBlock, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   latestblock.Table,
			Columns: latestblock.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: latestblock.FieldID,
			},
		},
	}
	id, ok := lbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing LatestBlock.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := lbuo.mutation.BlockNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: latestblock.FieldBlockNumber,
		})
	}
	if value, ok := lbuo.mutation.AddedBlockNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: latestblock.FieldBlockNumber,
		})
	}
	_node = &LatestBlock{config: lbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{latestblock.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

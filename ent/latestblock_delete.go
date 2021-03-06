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

// LatestBlockDelete is the builder for deleting a LatestBlock entity.
type LatestBlockDelete struct {
	config
	hooks    []Hook
	mutation *LatestBlockMutation
}

// Where adds a new predicate to the LatestBlockDelete builder.
func (lbd *LatestBlockDelete) Where(ps ...predicate.LatestBlock) *LatestBlockDelete {
	lbd.mutation.predicates = append(lbd.mutation.predicates, ps...)
	return lbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lbd *LatestBlockDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lbd.hooks) == 0 {
		affected, err = lbd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LatestBlockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lbd.mutation = mutation
			affected, err = lbd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lbd.hooks) - 1; i >= 0; i-- {
			mut = lbd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lbd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (lbd *LatestBlockDelete) ExecX(ctx context.Context) int {
	n, err := lbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lbd *LatestBlockDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: latestblock.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: latestblock.FieldID,
			},
		},
	}
	if ps := lbd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, lbd.driver, _spec)
}

// LatestBlockDeleteOne is the builder for deleting a single LatestBlock entity.
type LatestBlockDeleteOne struct {
	lbd *LatestBlockDelete
}

// Exec executes the deletion query.
func (lbdo *LatestBlockDeleteOne) Exec(ctx context.Context) error {
	n, err := lbdo.lbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{latestblock.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (lbdo *LatestBlockDeleteOne) ExecX(ctx context.Context) {
	lbdo.lbd.ExecX(ctx)
}

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sundaytycoon/buttons-api/ent/predicate"
	"github.com/sundaytycoon/buttons-api/ent/usermeta"
)

// UserMetaDelete is the builder for deleting a UserMeta entity.
type UserMetaDelete struct {
	config
	hooks    []Hook
	mutation *UserMetaMutation
}

// Where appends a list predicates to the UserMetaDelete builder.
func (umd *UserMetaDelete) Where(ps ...predicate.UserMeta) *UserMetaDelete {
	umd.mutation.Where(ps...)
	return umd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (umd *UserMetaDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(umd.hooks) == 0 {
		affected, err = umd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMetaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			umd.mutation = mutation
			affected, err = umd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(umd.hooks) - 1; i >= 0; i-- {
			if umd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = umd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, umd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (umd *UserMetaDelete) ExecX(ctx context.Context) int {
	n, err := umd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (umd *UserMetaDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: usermeta.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: usermeta.FieldID,
			},
		},
	}
	if ps := umd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, umd.driver, _spec)
}

// UserMetaDeleteOne is the builder for deleting a single UserMeta entity.
type UserMetaDeleteOne struct {
	umd *UserMetaDelete
}

// Exec executes the deletion query.
func (umdo *UserMetaDeleteOne) Exec(ctx context.Context) error {
	n, err := umdo.umd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{usermeta.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (umdo *UserMetaDeleteOne) ExecX(ctx context.Context) {
	umdo.umd.ExecX(ctx)
}
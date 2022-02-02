// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sundaytycoon/buttons-api/ent/predicate"
	"github.com/sundaytycoon/buttons-api/ent/user"
	"github.com/sundaytycoon/buttons-api/ent/useroauthprovider"
)

// UserOAuthProviderUpdate is the builder for updating UserOAuthProvider entities.
type UserOAuthProviderUpdate struct {
	config
	hooks    []Hook
	mutation *UserOAuthProviderMutation
}

// Where appends a list predicates to the UserOAuthProviderUpdate builder.
func (uopu *UserOAuthProviderUpdate) Where(ps ...predicate.UserOAuthProvider) *UserOAuthProviderUpdate {
	uopu.mutation.Where(ps...)
	return uopu
}

// SetUserID sets the "user_id" field.
func (uopu *UserOAuthProviderUpdate) SetUserID(s string) *UserOAuthProviderUpdate {
	uopu.mutation.SetUserID(s)
	return uopu
}

// SetCreatedAt sets the "created_at" field.
func (uopu *UserOAuthProviderUpdate) SetCreatedAt(t time.Time) *UserOAuthProviderUpdate {
	uopu.mutation.SetCreatedAt(t)
	return uopu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uopu *UserOAuthProviderUpdate) SetNillableCreatedAt(t *time.Time) *UserOAuthProviderUpdate {
	if t != nil {
		uopu.SetCreatedAt(*t)
	}
	return uopu
}

// SetCreatedBy sets the "created_by" field.
func (uopu *UserOAuthProviderUpdate) SetCreatedBy(s string) *UserOAuthProviderUpdate {
	uopu.mutation.SetCreatedBy(s)
	return uopu
}

// SetUpdatedAt sets the "updated_at" field.
func (uopu *UserOAuthProviderUpdate) SetUpdatedAt(t time.Time) *UserOAuthProviderUpdate {
	uopu.mutation.SetUpdatedAt(t)
	return uopu
}

// SetUpdatedBy sets the "updated_by" field.
func (uopu *UserOAuthProviderUpdate) SetUpdatedBy(s string) *UserOAuthProviderUpdate {
	uopu.mutation.SetUpdatedBy(s)
	return uopu
}

// SetStatus sets the "status" field.
func (uopu *UserOAuthProviderUpdate) SetStatus(u useroauthprovider.Status) *UserOAuthProviderUpdate {
	uopu.mutation.SetStatus(u)
	return uopu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uopu *UserOAuthProviderUpdate) SetNillableStatus(u *useroauthprovider.Status) *UserOAuthProviderUpdate {
	if u != nil {
		uopu.SetStatus(*u)
	}
	return uopu
}

// SetProvider sets the "provider" field.
func (uopu *UserOAuthProviderUpdate) SetProvider(u useroauthprovider.Provider) *UserOAuthProviderUpdate {
	uopu.mutation.SetProvider(u)
	return uopu
}

// SetExpiry sets the "expiry" field.
func (uopu *UserOAuthProviderUpdate) SetExpiry(t time.Time) *UserOAuthProviderUpdate {
	uopu.mutation.SetExpiry(t)
	return uopu
}

// SetAccessToken sets the "access_token" field.
func (uopu *UserOAuthProviderUpdate) SetAccessToken(s string) *UserOAuthProviderUpdate {
	uopu.mutation.SetAccessToken(s)
	return uopu
}

// SetRefreshToken sets the "refresh_token" field.
func (uopu *UserOAuthProviderUpdate) SetRefreshToken(s string) *UserOAuthProviderUpdate {
	uopu.mutation.SetRefreshToken(s)
	return uopu
}

// SetUser sets the "user" edge to the User entity.
func (uopu *UserOAuthProviderUpdate) SetUser(u *User) *UserOAuthProviderUpdate {
	return uopu.SetUserID(u.ID)
}

// Mutation returns the UserOAuthProviderMutation object of the builder.
func (uopu *UserOAuthProviderUpdate) Mutation() *UserOAuthProviderMutation {
	return uopu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (uopu *UserOAuthProviderUpdate) ClearUser() *UserOAuthProviderUpdate {
	uopu.mutation.ClearUser()
	return uopu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uopu *UserOAuthProviderUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	uopu.defaults()
	if len(uopu.hooks) == 0 {
		if err = uopu.check(); err != nil {
			return 0, err
		}
		affected, err = uopu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserOAuthProviderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uopu.check(); err != nil {
				return 0, err
			}
			uopu.mutation = mutation
			affected, err = uopu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uopu.hooks) - 1; i >= 0; i-- {
			if uopu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uopu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uopu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uopu *UserOAuthProviderUpdate) SaveX(ctx context.Context) int {
	affected, err := uopu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uopu *UserOAuthProviderUpdate) Exec(ctx context.Context) error {
	_, err := uopu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uopu *UserOAuthProviderUpdate) ExecX(ctx context.Context) {
	if err := uopu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uopu *UserOAuthProviderUpdate) defaults() {
	if _, ok := uopu.mutation.UpdatedAt(); !ok {
		v := useroauthprovider.UpdateDefaultUpdatedAt()
		uopu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uopu *UserOAuthProviderUpdate) check() error {
	if v, ok := uopu.mutation.Status(); ok {
		if err := useroauthprovider.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "UserOAuthProvider.status": %w`, err)}
		}
	}
	if v, ok := uopu.mutation.Provider(); ok {
		if err := useroauthprovider.ProviderValidator(v); err != nil {
			return &ValidationError{Name: "provider", err: fmt.Errorf(`ent: validator failed for field "UserOAuthProvider.provider": %w`, err)}
		}
	}
	if _, ok := uopu.mutation.UserID(); uopu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserOAuthProvider.user"`)
	}
	return nil
}

func (uopu *UserOAuthProviderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   useroauthprovider.Table,
			Columns: useroauthprovider.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: useroauthprovider.FieldID,
			},
		},
	}
	if ps := uopu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uopu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: useroauthprovider.FieldCreatedAt,
		})
	}
	if value, ok := uopu.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: useroauthprovider.FieldCreatedBy,
		})
	}
	if value, ok := uopu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: useroauthprovider.FieldUpdatedAt,
		})
	}
	if value, ok := uopu.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: useroauthprovider.FieldUpdatedBy,
		})
	}
	if value, ok := uopu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: useroauthprovider.FieldStatus,
		})
	}
	if value, ok := uopu.mutation.Provider(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: useroauthprovider.FieldProvider,
		})
	}
	if value, ok := uopu.mutation.Expiry(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: useroauthprovider.FieldExpiry,
		})
	}
	if value, ok := uopu.mutation.AccessToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: useroauthprovider.FieldAccessToken,
		})
	}
	if value, ok := uopu.mutation.RefreshToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: useroauthprovider.FieldRefreshToken,
		})
	}
	if uopu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   useroauthprovider.UserTable,
			Columns: []string{useroauthprovider.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uopu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   useroauthprovider.UserTable,
			Columns: []string{useroauthprovider.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uopu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{useroauthprovider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserOAuthProviderUpdateOne is the builder for updating a single UserOAuthProvider entity.
type UserOAuthProviderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserOAuthProviderMutation
}

// SetUserID sets the "user_id" field.
func (uopuo *UserOAuthProviderUpdateOne) SetUserID(s string) *UserOAuthProviderUpdateOne {
	uopuo.mutation.SetUserID(s)
	return uopuo
}

// SetCreatedAt sets the "created_at" field.
func (uopuo *UserOAuthProviderUpdateOne) SetCreatedAt(t time.Time) *UserOAuthProviderUpdateOne {
	uopuo.mutation.SetCreatedAt(t)
	return uopuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uopuo *UserOAuthProviderUpdateOne) SetNillableCreatedAt(t *time.Time) *UserOAuthProviderUpdateOne {
	if t != nil {
		uopuo.SetCreatedAt(*t)
	}
	return uopuo
}

// SetCreatedBy sets the "created_by" field.
func (uopuo *UserOAuthProviderUpdateOne) SetCreatedBy(s string) *UserOAuthProviderUpdateOne {
	uopuo.mutation.SetCreatedBy(s)
	return uopuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uopuo *UserOAuthProviderUpdateOne) SetUpdatedAt(t time.Time) *UserOAuthProviderUpdateOne {
	uopuo.mutation.SetUpdatedAt(t)
	return uopuo
}

// SetUpdatedBy sets the "updated_by" field.
func (uopuo *UserOAuthProviderUpdateOne) SetUpdatedBy(s string) *UserOAuthProviderUpdateOne {
	uopuo.mutation.SetUpdatedBy(s)
	return uopuo
}

// SetStatus sets the "status" field.
func (uopuo *UserOAuthProviderUpdateOne) SetStatus(u useroauthprovider.Status) *UserOAuthProviderUpdateOne {
	uopuo.mutation.SetStatus(u)
	return uopuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uopuo *UserOAuthProviderUpdateOne) SetNillableStatus(u *useroauthprovider.Status) *UserOAuthProviderUpdateOne {
	if u != nil {
		uopuo.SetStatus(*u)
	}
	return uopuo
}

// SetProvider sets the "provider" field.
func (uopuo *UserOAuthProviderUpdateOne) SetProvider(u useroauthprovider.Provider) *UserOAuthProviderUpdateOne {
	uopuo.mutation.SetProvider(u)
	return uopuo
}

// SetExpiry sets the "expiry" field.
func (uopuo *UserOAuthProviderUpdateOne) SetExpiry(t time.Time) *UserOAuthProviderUpdateOne {
	uopuo.mutation.SetExpiry(t)
	return uopuo
}

// SetAccessToken sets the "access_token" field.
func (uopuo *UserOAuthProviderUpdateOne) SetAccessToken(s string) *UserOAuthProviderUpdateOne {
	uopuo.mutation.SetAccessToken(s)
	return uopuo
}

// SetRefreshToken sets the "refresh_token" field.
func (uopuo *UserOAuthProviderUpdateOne) SetRefreshToken(s string) *UserOAuthProviderUpdateOne {
	uopuo.mutation.SetRefreshToken(s)
	return uopuo
}

// SetUser sets the "user" edge to the User entity.
func (uopuo *UserOAuthProviderUpdateOne) SetUser(u *User) *UserOAuthProviderUpdateOne {
	return uopuo.SetUserID(u.ID)
}

// Mutation returns the UserOAuthProviderMutation object of the builder.
func (uopuo *UserOAuthProviderUpdateOne) Mutation() *UserOAuthProviderMutation {
	return uopuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (uopuo *UserOAuthProviderUpdateOne) ClearUser() *UserOAuthProviderUpdateOne {
	uopuo.mutation.ClearUser()
	return uopuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uopuo *UserOAuthProviderUpdateOne) Select(field string, fields ...string) *UserOAuthProviderUpdateOne {
	uopuo.fields = append([]string{field}, fields...)
	return uopuo
}

// Save executes the query and returns the updated UserOAuthProvider entity.
func (uopuo *UserOAuthProviderUpdateOne) Save(ctx context.Context) (*UserOAuthProvider, error) {
	var (
		err  error
		node *UserOAuthProvider
	)
	uopuo.defaults()
	if len(uopuo.hooks) == 0 {
		if err = uopuo.check(); err != nil {
			return nil, err
		}
		node, err = uopuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserOAuthProviderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uopuo.check(); err != nil {
				return nil, err
			}
			uopuo.mutation = mutation
			node, err = uopuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uopuo.hooks) - 1; i >= 0; i-- {
			if uopuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uopuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uopuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uopuo *UserOAuthProviderUpdateOne) SaveX(ctx context.Context) *UserOAuthProvider {
	node, err := uopuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uopuo *UserOAuthProviderUpdateOne) Exec(ctx context.Context) error {
	_, err := uopuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uopuo *UserOAuthProviderUpdateOne) ExecX(ctx context.Context) {
	if err := uopuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uopuo *UserOAuthProviderUpdateOne) defaults() {
	if _, ok := uopuo.mutation.UpdatedAt(); !ok {
		v := useroauthprovider.UpdateDefaultUpdatedAt()
		uopuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uopuo *UserOAuthProviderUpdateOne) check() error {
	if v, ok := uopuo.mutation.Status(); ok {
		if err := useroauthprovider.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "UserOAuthProvider.status": %w`, err)}
		}
	}
	if v, ok := uopuo.mutation.Provider(); ok {
		if err := useroauthprovider.ProviderValidator(v); err != nil {
			return &ValidationError{Name: "provider", err: fmt.Errorf(`ent: validator failed for field "UserOAuthProvider.provider": %w`, err)}
		}
	}
	if _, ok := uopuo.mutation.UserID(); uopuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserOAuthProvider.user"`)
	}
	return nil
}

func (uopuo *UserOAuthProviderUpdateOne) sqlSave(ctx context.Context) (_node *UserOAuthProvider, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   useroauthprovider.Table,
			Columns: useroauthprovider.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: useroauthprovider.FieldID,
			},
		},
	}
	id, ok := uopuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserOAuthProvider.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uopuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, useroauthprovider.FieldID)
		for _, f := range fields {
			if !useroauthprovider.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != useroauthprovider.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uopuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uopuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: useroauthprovider.FieldCreatedAt,
		})
	}
	if value, ok := uopuo.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: useroauthprovider.FieldCreatedBy,
		})
	}
	if value, ok := uopuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: useroauthprovider.FieldUpdatedAt,
		})
	}
	if value, ok := uopuo.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: useroauthprovider.FieldUpdatedBy,
		})
	}
	if value, ok := uopuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: useroauthprovider.FieldStatus,
		})
	}
	if value, ok := uopuo.mutation.Provider(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: useroauthprovider.FieldProvider,
		})
	}
	if value, ok := uopuo.mutation.Expiry(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: useroauthprovider.FieldExpiry,
		})
	}
	if value, ok := uopuo.mutation.AccessToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: useroauthprovider.FieldAccessToken,
		})
	}
	if value, ok := uopuo.mutation.RefreshToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: useroauthprovider.FieldRefreshToken,
		})
	}
	if uopuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   useroauthprovider.UserTable,
			Columns: []string{useroauthprovider.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uopuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   useroauthprovider.UserTable,
			Columns: []string{useroauthprovider.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserOAuthProvider{config: uopuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uopuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{useroauthprovider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

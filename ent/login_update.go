// Code generated by ent, DO NOT EDIT.

package ent

import (
	"chatapp/backend/ent/login"
	"chatapp/backend/ent/predicate"
	"chatapp/backend/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LoginUpdate is the builder for updating Login entities.
type LoginUpdate struct {
	config
	hooks    []Hook
	mutation *LoginMutation
}

// Where appends a list predicates to the LoginUpdate builder.
func (lu *LoginUpdate) Where(ps ...predicate.Login) *LoginUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetUsername sets the "username" field.
func (lu *LoginUpdate) SetUsername(s string) *LoginUpdate {
	lu.mutation.SetUsername(s)
	return lu
}

// SetEmail sets the "email" field.
func (lu *LoginUpdate) SetEmail(s string) *LoginUpdate {
	lu.mutation.SetEmail(s)
	return lu
}

// SetUID sets the "uid" field.
func (lu *LoginUpdate) SetUID(s string) *LoginUpdate {
	lu.mutation.SetUID(s)
	return lu
}

// SetCreatedAt sets the "created_at" field.
func (lu *LoginUpdate) SetCreatedAt(t time.Time) *LoginUpdate {
	lu.mutation.SetCreatedAt(t)
	return lu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lu *LoginUpdate) SetNillableCreatedAt(t *time.Time) *LoginUpdate {
	if t != nil {
		lu.SetCreatedAt(*t)
	}
	return lu
}

// SetStatus sets the "status" field.
func (lu *LoginUpdate) SetStatus(l login.Status) *LoginUpdate {
	lu.mutation.SetStatus(l)
	return lu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (lu *LoginUpdate) SetUserID(id int) *LoginUpdate {
	lu.mutation.SetUserID(id)
	return lu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (lu *LoginUpdate) SetNillableUserID(id *int) *LoginUpdate {
	if id != nil {
		lu = lu.SetUserID(*id)
	}
	return lu
}

// SetUser sets the "user" edge to the User entity.
func (lu *LoginUpdate) SetUser(u *User) *LoginUpdate {
	return lu.SetUserID(u.ID)
}

// Mutation returns the LoginMutation object of the builder.
func (lu *LoginUpdate) Mutation() *LoginMutation {
	return lu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (lu *LoginUpdate) ClearUser() *LoginUpdate {
	lu.mutation.ClearUser()
	return lu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LoginUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lu.hooks) == 0 {
		if err = lu.check(); err != nil {
			return 0, err
		}
		affected, err = lu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LoginMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lu.check(); err != nil {
				return 0, err
			}
			lu.mutation = mutation
			affected, err = lu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lu.hooks) - 1; i >= 0; i-- {
			if lu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LoginUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LoginUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LoginUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lu *LoginUpdate) check() error {
	if v, ok := lu.mutation.Status(); ok {
		if err := login.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Login.status": %w`, err)}
		}
	}
	return nil
}

func (lu *LoginUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   login.Table,
			Columns: login.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: login.FieldID,
			},
		},
	}
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: login.FieldUsername,
		})
	}
	if value, ok := lu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: login.FieldEmail,
		})
	}
	if value, ok := lu.mutation.UID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: login.FieldUID,
		})
	}
	if value, ok := lu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: login.FieldCreatedAt,
		})
	}
	if value, ok := lu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: login.FieldStatus,
		})
	}
	if lu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   login.UserTable,
			Columns: []string{login.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   login.UserTable,
			Columns: []string{login.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{login.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// LoginUpdateOne is the builder for updating a single Login entity.
type LoginUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LoginMutation
}

// SetUsername sets the "username" field.
func (luo *LoginUpdateOne) SetUsername(s string) *LoginUpdateOne {
	luo.mutation.SetUsername(s)
	return luo
}

// SetEmail sets the "email" field.
func (luo *LoginUpdateOne) SetEmail(s string) *LoginUpdateOne {
	luo.mutation.SetEmail(s)
	return luo
}

// SetUID sets the "uid" field.
func (luo *LoginUpdateOne) SetUID(s string) *LoginUpdateOne {
	luo.mutation.SetUID(s)
	return luo
}

// SetCreatedAt sets the "created_at" field.
func (luo *LoginUpdateOne) SetCreatedAt(t time.Time) *LoginUpdateOne {
	luo.mutation.SetCreatedAt(t)
	return luo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (luo *LoginUpdateOne) SetNillableCreatedAt(t *time.Time) *LoginUpdateOne {
	if t != nil {
		luo.SetCreatedAt(*t)
	}
	return luo
}

// SetStatus sets the "status" field.
func (luo *LoginUpdateOne) SetStatus(l login.Status) *LoginUpdateOne {
	luo.mutation.SetStatus(l)
	return luo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (luo *LoginUpdateOne) SetUserID(id int) *LoginUpdateOne {
	luo.mutation.SetUserID(id)
	return luo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (luo *LoginUpdateOne) SetNillableUserID(id *int) *LoginUpdateOne {
	if id != nil {
		luo = luo.SetUserID(*id)
	}
	return luo
}

// SetUser sets the "user" edge to the User entity.
func (luo *LoginUpdateOne) SetUser(u *User) *LoginUpdateOne {
	return luo.SetUserID(u.ID)
}

// Mutation returns the LoginMutation object of the builder.
func (luo *LoginUpdateOne) Mutation() *LoginMutation {
	return luo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (luo *LoginUpdateOne) ClearUser() *LoginUpdateOne {
	luo.mutation.ClearUser()
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LoginUpdateOne) Select(field string, fields ...string) *LoginUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Login entity.
func (luo *LoginUpdateOne) Save(ctx context.Context) (*Login, error) {
	var (
		err  error
		node *Login
	)
	if len(luo.hooks) == 0 {
		if err = luo.check(); err != nil {
			return nil, err
		}
		node, err = luo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LoginMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = luo.check(); err != nil {
				return nil, err
			}
			luo.mutation = mutation
			node, err = luo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(luo.hooks) - 1; i >= 0; i-- {
			if luo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = luo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, luo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Login)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from LoginMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LoginUpdateOne) SaveX(ctx context.Context) *Login {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LoginUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LoginUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luo *LoginUpdateOne) check() error {
	if v, ok := luo.mutation.Status(); ok {
		if err := login.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Login.status": %w`, err)}
		}
	}
	return nil
}

func (luo *LoginUpdateOne) sqlSave(ctx context.Context) (_node *Login, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   login.Table,
			Columns: login.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: login.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Login.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, login.FieldID)
		for _, f := range fields {
			if !login.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != login.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: login.FieldUsername,
		})
	}
	if value, ok := luo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: login.FieldEmail,
		})
	}
	if value, ok := luo.mutation.UID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: login.FieldUID,
		})
	}
	if value, ok := luo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: login.FieldCreatedAt,
		})
	}
	if value, ok := luo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: login.FieldStatus,
		})
	}
	if luo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   login.UserTable,
			Columns: []string{login.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   login.UserTable,
			Columns: []string{login.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Login{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{login.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

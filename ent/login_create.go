// Code generated by ent, DO NOT EDIT.

package ent

import (
	"chatapp/backend/ent/login"
	"chatapp/backend/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LoginCreate is the builder for creating a Login entity.
type LoginCreate struct {
	config
	mutation *LoginMutation
	hooks    []Hook
}

// SetUsername sets the "username" field.
func (lc *LoginCreate) SetUsername(s string) *LoginCreate {
	lc.mutation.SetUsername(s)
	return lc
}

// SetEmail sets the "email" field.
func (lc *LoginCreate) SetEmail(s string) *LoginCreate {
	lc.mutation.SetEmail(s)
	return lc
}

// SetUUID sets the "uuid" field.
func (lc *LoginCreate) SetUUID(s string) *LoginCreate {
	lc.mutation.SetUUID(s)
	return lc
}

// SetCreatedAt sets the "created_at" field.
func (lc *LoginCreate) SetCreatedAt(t time.Time) *LoginCreate {
	lc.mutation.SetCreatedAt(t)
	return lc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lc *LoginCreate) SetNillableCreatedAt(t *time.Time) *LoginCreate {
	if t != nil {
		lc.SetCreatedAt(*t)
	}
	return lc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (lc *LoginCreate) SetUserID(id int) *LoginCreate {
	lc.mutation.SetUserID(id)
	return lc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (lc *LoginCreate) SetNillableUserID(id *int) *LoginCreate {
	if id != nil {
		lc = lc.SetUserID(*id)
	}
	return lc
}

// SetUser sets the "user" edge to the User entity.
func (lc *LoginCreate) SetUser(u *User) *LoginCreate {
	return lc.SetUserID(u.ID)
}

// Mutation returns the LoginMutation object of the builder.
func (lc *LoginCreate) Mutation() *LoginMutation {
	return lc.mutation
}

// Save creates the Login in the database.
func (lc *LoginCreate) Save(ctx context.Context) (*Login, error) {
	var (
		err  error
		node *Login
	)
	lc.defaults()
	if len(lc.hooks) == 0 {
		if err = lc.check(); err != nil {
			return nil, err
		}
		node, err = lc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LoginMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lc.check(); err != nil {
				return nil, err
			}
			lc.mutation = mutation
			if node, err = lc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(lc.hooks) - 1; i >= 0; i-- {
			if lc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, lc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (lc *LoginCreate) SaveX(ctx context.Context) *Login {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LoginCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LoginCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LoginCreate) defaults() {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		v := login.DefaultCreatedAt()
		lc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LoginCreate) check() error {
	if _, ok := lc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "Login.username"`)}
	}
	if _, ok := lc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Login.email"`)}
	}
	if _, ok := lc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "Login.uuid"`)}
	}
	if _, ok := lc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Login.created_at"`)}
	}
	return nil
}

func (lc *LoginCreate) sqlSave(ctx context.Context) (*Login, error) {
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (lc *LoginCreate) createSpec() (*Login, *sqlgraph.CreateSpec) {
	var (
		_node = &Login{config: lc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: login.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: login.FieldID,
			},
		}
	)
	if value, ok := lc.mutation.Username(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: login.FieldUsername,
		})
		_node.Username = value
	}
	if value, ok := lc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: login.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := lc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: login.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := lc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: login.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := lc.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LoginCreateBulk is the builder for creating many Login entities in bulk.
type LoginCreateBulk struct {
	config
	builders []*LoginCreate
}

// Save creates the Login entities in the database.
func (lcb *LoginCreateBulk) Save(ctx context.Context) ([]*Login, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Login, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LoginMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LoginCreateBulk) SaveX(ctx context.Context) []*Login {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LoginCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LoginCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
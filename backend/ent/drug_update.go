// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/wipha26/app/ent/dispense"
	"github.com/wipha26/app/ent/drug"
	"github.com/wipha26/app/ent/predicate"
)

// DrugUpdate is the builder for updating Drug entities.
type DrugUpdate struct {
	config
	hooks      []Hook
	mutation   *DrugMutation
	predicates []predicate.Drug
}

// Where adds a new predicate for the builder.
func (du *DrugUpdate) Where(ps ...predicate.Drug) *DrugUpdate {
	du.predicates = append(du.predicates, ps...)
	return du
}

// SetName sets the name field.
func (du *DrugUpdate) SetName(s string) *DrugUpdate {
	du.mutation.SetName(s)
	return du
}

// AddDispenseIDs adds the dispenses edge to Dispense by ids.
func (du *DrugUpdate) AddDispenseIDs(ids ...int) *DrugUpdate {
	du.mutation.AddDispenseIDs(ids...)
	return du
}

// AddDispenses adds the dispenses edges to Dispense.
func (du *DrugUpdate) AddDispenses(d ...*Dispense) *DrugUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.AddDispenseIDs(ids...)
}

// Mutation returns the DrugMutation object of the builder.
func (du *DrugUpdate) Mutation() *DrugMutation {
	return du.mutation
}

// RemoveDispenseIDs removes the dispenses edge to Dispense by ids.
func (du *DrugUpdate) RemoveDispenseIDs(ids ...int) *DrugUpdate {
	du.mutation.RemoveDispenseIDs(ids...)
	return du
}

// RemoveDispenses removes dispenses edges to Dispense.
func (du *DrugUpdate) RemoveDispenses(d ...*Dispense) *DrugUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.RemoveDispenseIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (du *DrugUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := du.mutation.Name(); ok {
		if err := drug.NameValidator(v); err != nil {
			return 0, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DrugMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DrugUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DrugUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DrugUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DrugUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   drug.Table,
			Columns: drug.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: drug.FieldID,
			},
		},
	}
	if ps := du.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: drug.FieldName,
		})
	}
	if nodes := du.mutation.RemovedDispensesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   drug.DispensesTable,
			Columns: []string{drug.DispensesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dispense.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.DispensesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   drug.DispensesTable,
			Columns: []string{drug.DispensesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dispense.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{drug.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DrugUpdateOne is the builder for updating a single Drug entity.
type DrugUpdateOne struct {
	config
	hooks    []Hook
	mutation *DrugMutation
}

// SetName sets the name field.
func (duo *DrugUpdateOne) SetName(s string) *DrugUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// AddDispenseIDs adds the dispenses edge to Dispense by ids.
func (duo *DrugUpdateOne) AddDispenseIDs(ids ...int) *DrugUpdateOne {
	duo.mutation.AddDispenseIDs(ids...)
	return duo
}

// AddDispenses adds the dispenses edges to Dispense.
func (duo *DrugUpdateOne) AddDispenses(d ...*Dispense) *DrugUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.AddDispenseIDs(ids...)
}

// Mutation returns the DrugMutation object of the builder.
func (duo *DrugUpdateOne) Mutation() *DrugMutation {
	return duo.mutation
}

// RemoveDispenseIDs removes the dispenses edge to Dispense by ids.
func (duo *DrugUpdateOne) RemoveDispenseIDs(ids ...int) *DrugUpdateOne {
	duo.mutation.RemoveDispenseIDs(ids...)
	return duo
}

// RemoveDispenses removes dispenses edges to Dispense.
func (duo *DrugUpdateOne) RemoveDispenses(d ...*Dispense) *DrugUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.RemoveDispenseIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (duo *DrugUpdateOne) Save(ctx context.Context) (*Drug, error) {
	if v, ok := duo.mutation.Name(); ok {
		if err := drug.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}

	var (
		err  error
		node *Drug
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DrugMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DrugUpdateOne) SaveX(ctx context.Context) *Drug {
	d, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// Exec executes the query on the entity.
func (duo *DrugUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DrugUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DrugUpdateOne) sqlSave(ctx context.Context) (d *Drug, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   drug.Table,
			Columns: drug.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: drug.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Drug.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: drug.FieldName,
		})
	}
	if nodes := duo.mutation.RemovedDispensesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   drug.DispensesTable,
			Columns: []string{drug.DispensesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dispense.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.DispensesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   drug.DispensesTable,
			Columns: []string{drug.DispensesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dispense.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	d = &Drug{config: duo.config}
	_spec.Assign = d.assignValues
	_spec.ScanValues = d.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{drug.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return d, nil
}

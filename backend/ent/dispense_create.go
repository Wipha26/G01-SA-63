// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/wipha26/app/ent/dispense"
	"github.com/wipha26/app/ent/drug"
	"github.com/wipha26/app/ent/patient"
	"github.com/wipha26/app/ent/user"
)

// DispenseCreate is the builder for creating a Dispense entity.
type DispenseCreate struct {
	config
	mutation *DispenseMutation
	hooks    []Hook
}

// SetNote sets the note field.
func (dc *DispenseCreate) SetNote(s string) *DispenseCreate {
	dc.mutation.SetNote(s)
	return dc
}

// SetDrugID sets the drug edge to Drug by id.
func (dc *DispenseCreate) SetDrugID(id int) *DispenseCreate {
	dc.mutation.SetDrugID(id)
	return dc
}

// SetNillableDrugID sets the drug edge to Drug by id if the given value is not nil.
func (dc *DispenseCreate) SetNillableDrugID(id *int) *DispenseCreate {
	if id != nil {
		dc = dc.SetDrugID(*id)
	}
	return dc
}

// SetDrug sets the drug edge to Drug.
func (dc *DispenseCreate) SetDrug(d *Drug) *DispenseCreate {
	return dc.SetDrugID(d.ID)
}

// SetPatientID sets the patient edge to Patient by id.
func (dc *DispenseCreate) SetPatientID(id int) *DispenseCreate {
	dc.mutation.SetPatientID(id)
	return dc
}

// SetNillablePatientID sets the patient edge to Patient by id if the given value is not nil.
func (dc *DispenseCreate) SetNillablePatientID(id *int) *DispenseCreate {
	if id != nil {
		dc = dc.SetPatientID(*id)
	}
	return dc
}

// SetPatient sets the patient edge to Patient.
func (dc *DispenseCreate) SetPatient(p *Patient) *DispenseCreate {
	return dc.SetPatientID(p.ID)
}

// SetUserID sets the user edge to User by id.
func (dc *DispenseCreate) SetUserID(id int) *DispenseCreate {
	dc.mutation.SetUserID(id)
	return dc
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (dc *DispenseCreate) SetNillableUserID(id *int) *DispenseCreate {
	if id != nil {
		dc = dc.SetUserID(*id)
	}
	return dc
}

// SetUser sets the user edge to User.
func (dc *DispenseCreate) SetUser(u *User) *DispenseCreate {
	return dc.SetUserID(u.ID)
}

// Mutation returns the DispenseMutation object of the builder.
func (dc *DispenseCreate) Mutation() *DispenseMutation {
	return dc.mutation
}

// Save creates the Dispense in the database.
func (dc *DispenseCreate) Save(ctx context.Context) (*Dispense, error) {
	if _, ok := dc.mutation.Note(); !ok {
		return nil, &ValidationError{Name: "note", err: errors.New("ent: missing required field \"note\"")}
	}
	if v, ok := dc.mutation.Note(); ok {
		if err := dispense.NoteValidator(v); err != nil {
			return nil, &ValidationError{Name: "note", err: fmt.Errorf("ent: validator failed for field \"note\": %w", err)}
		}
	}
	var (
		err  error
		node *Dispense
	)
	if len(dc.hooks) == 0 {
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DispenseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dc.mutation = mutation
			node, err = dc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DispenseCreate) SaveX(ctx context.Context) *Dispense {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dc *DispenseCreate) sqlSave(ctx context.Context) (*Dispense, error) {
	d, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	d.ID = int(id)
	return d, nil
}

func (dc *DispenseCreate) createSpec() (*Dispense, *sqlgraph.CreateSpec) {
	var (
		d     = &Dispense{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: dispense.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dispense.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.Note(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dispense.FieldNote,
		})
		d.Note = value
	}
	if nodes := dc.mutation.DrugIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dispense.DrugTable,
			Columns: []string{dispense.DrugColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: drug.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dispense.PatientTable,
			Columns: []string{dispense.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dispense.UserTable,
			Columns: []string{dispense.UserColumn},
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
	return d, _spec
}

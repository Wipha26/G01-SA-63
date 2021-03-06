// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/wipha26/app/ent/patient"
)

// Patient is the model entity for the Patient schema.
type Patient struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Firstname holds the value of the "firstname" field.
	Firstname string `json:"firstname,omitempty"`
	// Lastname holds the value of the "lastname" field.
	Lastname string `json:"lastname,omitempty"`
	// Cardid holds the value of the "cardid" field.
	Cardid string `json:"cardid,omitempty"`
	// Allergic holds the value of the "allergic" field.
	Allergic string `json:"allergic,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Brithday holds the value of the "brithday" field.
	Brithday time.Time `json:"brithday,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PatientQuery when eager-loading is set.
	Edges PatientEdges `json:"edges"`
}

// PatientEdges holds the relations/edges for other nodes in the graph.
type PatientEdges struct {
	// Dispenses holds the value of the dispenses edge.
	Dispenses []*Dispense
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DispensesOrErr returns the Dispenses value or an error if the edge
// was not loaded in eager-loading.
func (e PatientEdges) DispensesOrErr() ([]*Dispense, error) {
	if e.loadedTypes[0] {
		return e.Dispenses, nil
	}
	return nil, &NotLoadedError{edge: "dispenses"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Patient) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // firstname
		&sql.NullString{}, // lastname
		&sql.NullString{}, // cardid
		&sql.NullString{}, // allergic
		&sql.NullInt64{},  // age
		&sql.NullTime{},   // brithday
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Patient fields.
func (pa *Patient) assignValues(values ...interface{}) error {
	if m, n := len(values), len(patient.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pa.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field firstname", values[0])
	} else if value.Valid {
		pa.Firstname = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field lastname", values[1])
	} else if value.Valid {
		pa.Lastname = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field cardid", values[2])
	} else if value.Valid {
		pa.Cardid = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field allergic", values[3])
	} else if value.Valid {
		pa.Allergic = value.String
	}
	if value, ok := values[4].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field age", values[4])
	} else if value.Valid {
		pa.Age = int(value.Int64)
	}
	if value, ok := values[5].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field brithday", values[5])
	} else if value.Valid {
		pa.Brithday = value.Time
	}
	return nil
}

// QueryDispenses queries the dispenses edge of the Patient.
func (pa *Patient) QueryDispenses() *DispenseQuery {
	return (&PatientClient{config: pa.config}).QueryDispenses(pa)
}

// Update returns a builder for updating this Patient.
// Note that, you need to call Patient.Unwrap() before calling this method, if this Patient
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Patient) Update() *PatientUpdateOne {
	return (&PatientClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pa *Patient) Unwrap() *Patient {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Patient is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Patient) String() string {
	var builder strings.Builder
	builder.WriteString("Patient(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteString(", firstname=")
	builder.WriteString(pa.Firstname)
	builder.WriteString(", lastname=")
	builder.WriteString(pa.Lastname)
	builder.WriteString(", cardid=")
	builder.WriteString(pa.Cardid)
	builder.WriteString(", allergic=")
	builder.WriteString(pa.Allergic)
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", pa.Age))
	builder.WriteString(", brithday=")
	builder.WriteString(pa.Brithday.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Patients is a parsable slice of Patient.
type Patients []*Patient

func (pa Patients) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}

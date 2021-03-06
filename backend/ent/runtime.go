// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/wipha26/app/ent/dispense"
	"github.com/wipha26/app/ent/drug"
	"github.com/wipha26/app/ent/patient"
	"github.com/wipha26/app/ent/schema"
	"github.com/wipha26/app/ent/user"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	dispenseFields := schema.Dispense{}.Fields()
	_ = dispenseFields
	// dispenseDescNote is the schema descriptor for note field.
	dispenseDescNote := dispenseFields[0].Descriptor()
	// dispense.NoteValidator is a validator for the "note" field. It is called by the builders before save.
	dispense.NoteValidator = dispenseDescNote.Validators[0].(func(string) error)
	drugFields := schema.Drug{}.Fields()
	_ = drugFields
	// drugDescName is the schema descriptor for name field.
	drugDescName := drugFields[0].Descriptor()
	// drug.NameValidator is a validator for the "name" field. It is called by the builders before save.
	drug.NameValidator = drugDescName.Validators[0].(func(string) error)
	patientFields := schema.Patient{}.Fields()
	_ = patientFields
	// patientDescFirstname is the schema descriptor for firstname field.
	patientDescFirstname := patientFields[0].Descriptor()
	// patient.FirstnameValidator is a validator for the "firstname" field. It is called by the builders before save.
	patient.FirstnameValidator = patientDescFirstname.Validators[0].(func(string) error)
	// patientDescLastname is the schema descriptor for lastname field.
	patientDescLastname := patientFields[1].Descriptor()
	// patient.LastnameValidator is a validator for the "lastname" field. It is called by the builders before save.
	patient.LastnameValidator = patientDescLastname.Validators[0].(func(string) error)
	// patientDescCardid is the schema descriptor for cardid field.
	patientDescCardid := patientFields[2].Descriptor()
	// patient.CardidValidator is a validator for the "cardid" field. It is called by the builders before save.
	patient.CardidValidator = patientDescCardid.Validators[0].(func(string) error)
	// patientDescAllergic is the schema descriptor for allergic field.
	patientDescAllergic := patientFields[3].Descriptor()
	// patient.AllergicValidator is a validator for the "allergic" field. It is called by the builders before save.
	patient.AllergicValidator = patientDescAllergic.Validators[0].(func(string) error)
	// patientDescAge is the schema descriptor for age field.
	patientDescAge := patientFields[4].Descriptor()
	// patient.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	patient.AgeValidator = patientDescAge.Validators[0].(func(int) error)
	// patientDescBrithday is the schema descriptor for brithday field.
	patientDescBrithday := patientFields[5].Descriptor()
	// patient.DefaultBrithday holds the default value on creation for the brithday field.
	patient.DefaultBrithday = patientDescBrithday.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
}

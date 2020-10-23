// Code generated by entc, DO NOT EDIT.

package dispense

const (
	// Label holds the string label denoting the dispense type in the database.
	Label = "dispense"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNote holds the string denoting the note field in the database.
	FieldNote = "note"

	// EdgeDrug holds the string denoting the drug edge name in mutations.
	EdgeDrug = "drug"
	// EdgePatient holds the string denoting the patient edge name in mutations.
	EdgePatient = "patient"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"

	// Table holds the table name of the dispense in the database.
	Table = "dispenses"
	// DrugTable is the table the holds the drug relation/edge.
	DrugTable = "dispenses"
	// DrugInverseTable is the table name for the Drug entity.
	// It exists in this package in order to avoid circular dependency with the "drug" package.
	DrugInverseTable = "drugs"
	// DrugColumn is the table column denoting the drug relation/edge.
	DrugColumn = "drug_dispenses"
	// PatientTable is the table the holds the patient relation/edge.
	PatientTable = "dispenses"
	// PatientInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	PatientInverseTable = "patients"
	// PatientColumn is the table column denoting the patient relation/edge.
	PatientColumn = "patient_dispenses"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "dispenses"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_dispenses"
)

// Columns holds all SQL columns for dispense fields.
var Columns = []string{
	FieldID,
	FieldNote,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Dispense type.
var ForeignKeys = []string{
	"drug_dispenses",
	"patient_dispenses",
	"user_dispenses",
}

var (
	// NoteValidator is a validator for the "note" field. It is called by the builders before save.
	NoteValidator func(string) error
)
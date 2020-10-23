package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Dispense holds the schema definition for the Dispense entity.
type Dispense struct {
	ent.Schema
}

// Fields of the Dispense.
func (Dispense) Fields() []ent.Field {
	return []ent.Field{
        field.String("note").NotEmpty(),   
    }

}

// Edges of the Dispense.
func (Dispense) Edges() []ent.Edge {
	return []ent.Edge{
        edge.From("drug",Drug.Type).Ref("dispenses").Unique(),
        edge.From("patient", Patient.Type).Ref("dispenses").Unique(),
        edge.From("user", User.Type).Ref("dispenses").Unique(),
        
        
    }

}

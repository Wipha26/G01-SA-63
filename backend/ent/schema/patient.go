package schema

import (
	"time"
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
 )
 

// Patient holds the schema definition for the Patient entity.
type Patient struct {
    ent.Schema
}

// Fields of the Patient.
func (Patient) Fields() []ent.Field {
    return []ent.Field{
        field.String("firstname").NotEmpty(),
        field.String("lastname").NotEmpty(),
        field.String("cardid").NotEmpty(),
        field.String("allergic").NotEmpty(),
        field.Int("age").Positive(),
        field.Time("brithday").Default(time.Now().Local).Immutable(),
    }
}

// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("dispenses",Dispense.Type),
    }
}

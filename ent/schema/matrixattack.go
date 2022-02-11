package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// MatrixAttack holds the schema definition for the MatrixAttack entity.
type MatrixAttack struct {
	ent.Schema
}

// Fields of the MatrixAttack.
func (MatrixAttack) Fields() []ent.Field {
	return []ent.Field{
		field.String("Id").Unique().Default("unknown"),
		field.String("VendorName").Default("unknown"),
		field.String("NameMatrix").Default("unknown"),
		field.String("VersionMatrix").Default("unknown"),
		field.Time("CreateDate").Default(time.Now),
		field.Time("ModifyDate").Default(time.Now),
	}
}

// Edges of the MatrixAttack.
func (MatrixAttack) Edges() []ent.Edge {
	return nil
}

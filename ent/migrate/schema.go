// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// MatrixAttacksColumns holds the columns for the "matrix_attacks" table.
	MatrixAttacksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "id_matrix", Type: field.TypeString, Unique: true},
		{Name: "vendor_name", Type: field.TypeString, Default: "unknown"},
		{Name: "name_matrix", Type: field.TypeString, Default: "unknown"},
		{Name: "version_matrix", Type: field.TypeString, Default: "unknown"},
		{Name: "create_date", Type: field.TypeTime},
		{Name: "modify_date", Type: field.TypeTime},
	}
	// MatrixAttacksTable holds the schema information for the "matrix_attacks" table.
	MatrixAttacksTable = &schema.Table{
		Name:       "matrix_attacks",
		Columns:    MatrixAttacksColumns,
		PrimaryKey: []*schema.Column{MatrixAttacksColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MatrixAttacksTable,
	}
)

func init() {
}

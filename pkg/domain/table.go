package domain

import (
	"github.com/comsma/knead/pkg/util"
)

type Table struct {
	Name    string
	Columns []Column
}

type Column struct {
	Name             string
	GoType           GoType
	Options          ColumnOptions
	FieldDescription string
}

type ColumnOptions struct {
	DataType          string
	DataTypePrecision string
	DefaultValue      string
	IsNullable        bool
	IsPrimaryKey      bool
	IsIdentity        bool
	IsUnique          bool
}

func (c *Column) GetNameCamelCase() string {
	return util.ToPascalCase(c.Name)
}

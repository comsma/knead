package bun

import (
	"github.com/comsma/knead/pkg/domain"
	"github.com/comsma/knead/pkg/util"
	"github.com/dave/jennifer/jen"
	"io"
)

type Generator struct {
}

func (g Generator) WriteFile(w io.Writer, table *domain.Table) error {
	f := jen.NewFile("gen")

	f.ImportName("github.com/uptrace/bun", "bun")
	f.ImportName("time", "time")

	var fields []jen.Code

	fields = append(fields, newTableAliasField(table))
	for _, col := range table.Columns {
		fields = append(fields, newColumnField(&col))
	}

	f.Type().Id(util.ToPascalCase(table.Name)).Struct(fields...)

	return f.Render(w)
}

func newTableAliasField(table *domain.Table) jen.Code {
	return jen.Qual("github.com/uptrace/bun", "BaseModel").Tag(map[string]string{"bun": "table:" + table.Name})
}
func newColumnField(col *domain.Column) jen.Code {
	f := jen.Id(util.ToPascalCase(col.Name))
	switch col.GoType {
	case domain.GoTypeString:
		f.String()
	case domain.GoTypeTime:
		f.Qual("time", "Time")
	case domain.GoTypeInt64:
		f.Int64()
	case domain.GoTypeInt32:
		f.Int32()
	case domain.GoTypeInt16:
		f.Int16()
	case domain.GoTypeFloat64:
		f.Float64()
	case domain.GoTypeBool:
		f.Bool()
	default:
		f.String()
	}
	f.Tag(map[string]string{"bun": newBunTagForColumn(col)})

	if col.FieldDescription != "" {
		f.Comment(col.FieldDescription)
	}
	return f
}

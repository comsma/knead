package bun

import (
	"github.com/comsma/knead/pkg/domain"
	"github.com/comsma/knead/pkg/orm/bun/tag"
)

func newBunTagForColumn(col *domain.Column) string {
	t := tag.NewTag().ColumnName(col.Name).Type(col.Options.DataTypePrecision)

	if col.Options.IsPrimaryKey {
		t.PrimaryKey()
	} else {
		if col.Options.IsUnique {
			t.Unique()
		}
	}

	if col.Options.IsIdentity {
		t.Identity().AutoIncrement()
	} else {
		if col.Options.DefaultValue != "" {
			t.Default(col.Options.DefaultValue)
		} else {

			if col.Options.IsNullable {
				t.NullZero()
			}
		}
	}

	return t.Build()
}

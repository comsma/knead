package mssql

import (
	"github.com/comsma/knead/pkg/domain"
	"github.com/comsma/knead/pkg/util"
	"log"
)

func (d *Database) GetTableInfo(name string) (*domain.Table, error) {

	rows, _ := d.db.Query(tableInfoQuery, name)
	defer rows.Close()

	var records []*tableInfoRecord
	for rows.Next() {
		var rec tableInfoRecord
		err := rows.Scan(&rec.ColumnName, &rec.DataType, &rec.DataTypePrecision, &rec.DefaultValue, &rec.IsNullable, &rec.IsIdentity, &rec.IsPrimaryKey, &rec.IsUnique, &rec.FieldDescription)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		records = append(records, &rec)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	table := tableInfoRecsToDomain(name, records)

	return &table, nil
}

func tableInfoRecsToDomain(name string, rec []*tableInfoRecord) domain.Table {
	t := domain.Table{
		Name: name,
	}

	t.Columns = make([]domain.Column, len(rec))
	for i, r := range rec {
		c := domain.Column{
			Name:             r.ColumnName,
			FieldDescription: util.NullableString(r.FieldDescription),
			Options: domain.ColumnOptions{
				DefaultValue:      util.NullableString(r.DefaultValue),
				DataTypePrecision: r.DataTypePrecision,
				DataType:          r.DataType,
				IsIdentity:        r.IsIdentity,
				IsNullable:        r.IsNullable,
				IsPrimaryKey:      util.NullableBool(r.IsPrimaryKey),
				IsUnique:          util.NullableBool(r.IsUnique),
			},
		}

		c.GoType = getGoType(c.Options.DataType)
		t.Columns[i] = c
	}

	return t
}

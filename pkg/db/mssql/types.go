package mssql

import "github.com/comsma/knead/pkg/domain"

func getGoType(sqlType string) (typ domain.GoType) {
	switch sqlType {
	case "nvarchar", "char", "varchar", "text", "smalldatetime":
		return domain.GoTypeString
	case "datetime":
		return domain.GoTypeTime
	case "timestamp":
		return domain.GoTypeByteArray
	case "bigint":
		return domain.GoTypeInt64
	case "int":
		return domain.GoTypeInt32
	case "smallint", "tinyint":
		return domain.GoTypeInt16
	case "decimal", "float":
		return domain.GoTypeFloat64
	case "bit":
		return domain.GoTypeBool
	default:
		return domain.GoTypeUnspecified
	}
}

type tableNameQueryResult struct {
	TableName   string
	RowCount    int
	TableSchema string
}

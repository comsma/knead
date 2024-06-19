package tag

import (
	"fmt"
	"strings"
)

const (
	altTag           = "alt"
	pkTag            = "pk"
	autoIncrementTag = "autoincrement"
	typeTag          = "type"
	defaultTag       = "default"
	notNullTag       = "notnull"
	uniqueTag        = "unique"
	nullZeroTag      = "nullzero"
	scanOnlyTag      = "scanonly"
	arrayTag         = "array"
	jsonUseNumberTag = "json_use_number"
	msgPackTag       = "msgpack"
	softDeleteTag    = "soft_delete"
	identity         = "identity"
)

type Tag struct {
	columnName string

	altName       *string
	primaryKey    bool
	autoIncrement bool
	sqlType       *string
	defaultValue  *string
	notNull       bool
	unique        bool
	uniqueGroup   *string
	nullZero      bool
	scanOnly      bool
	pgArray       bool
	jsonUseNumber bool
	msgPack       bool
	softDelete    bool
	identity      bool
}

func NewTag() *Tag {
	return &Tag{}
}

// Build returns a struct tag
func (t *Tag) Build() string {
	tag := []string{t.columnName}

	if t.altName != nil {
		tag = append(tag, fmt.Sprintf("%s:%s", altTag, *t.altName))

	}

	if t.sqlType != nil {
		tag = append(tag, fmt.Sprintf("%s:%s", typeTag, *t.sqlType))
	}

	if t.defaultValue != nil {
		tag = append(tag, fmt.Sprintf("%s:%s", defaultTag, *t.defaultValue))
	}
	if t.autoIncrement {
		tag = append(tag, autoIncrementTag)
	}

	if t.notNull {
		tag = append(tag, notNullTag)
	}

	if t.unique {
		if t.uniqueGroup != nil {
			tag = append(tag, fmt.Sprintf("%s:%s", uniqueTag, *t.uniqueGroup))
		} else {
			tag = append(tag, uniqueTag)
		}
	}

	if t.nullZero {
		tag = append(tag, nullZeroTag)
	}

	if t.scanOnly {
		tag = append(tag, scanOnlyTag)
	}

	if t.pgArray {
		tag = append(tag, arrayTag)
	}

	if t.jsonUseNumber {
		tag = append(tag, jsonUseNumberTag)
	}

	if t.msgPack {
		tag = append(tag, msgPackTag)
	}

	if t.softDelete {
		tag = append(tag, softDeleteTag)
	}

	if t.identity {
		tag = append(tag, identity)
	}

	if t.primaryKey {
		tag = append(tag, pkTag)

	}

	return strings.Join(tag, ",")
}

// ColumnName overrides default column name.
//
//	BunTag: "name"
func (t *Tag) ColumnName(name string) *Tag {
	t.columnName = name
	return t
}

// Alt sets alternative column name. Useful during migrations.
//
//	BunTag: "alt:name"
func (t *Tag) Alt(name string) *Tag {
	t.altName = &name
	return t
}

// PrimaryKey marks column as a primary key and apply notnull option. Multiple/composite primary keys are supported
//
//	BunTag: "pk"
func (t *Tag) PrimaryKey() *Tag {
	t.primaryKey = true
	return t
}

// AutoIncrement marks column as a serial in PostgreSQL, autoincrement in MySQL, and identity in MSSQL. Also applies nullzero option.
//
//	BunTag: "autoincrement
func (t *Tag) AutoIncrement() *Tag {
	t.autoIncrement = true
	return t
}

// Type overrides default SQL type.
//
//	BunTag: "type:sqlType"
func (t *Tag) Type(sqlType string) *Tag {
	t.sqlType = &sqlType
	return t
}

// Default tells CreateTable to set DEFAULT expression.
//
//	BunTag: "default:exp()"
func (t *Tag) Default(exp string) *Tag {
	t.defaultValue = &exp
	return t
}

// NotNull tells CreateTable to add NOT NULL constraint.
//
//	BunTag: "notnull"
func (t *Tag) NotNull() *Tag {
	t.notNull = true
	return t
}

// Unique tells CreateTable to add an unique constraint.
//
//	BunTag: "unique"
func (t *Tag) Unique() *Tag {
	t.unique = true
	return t
}

// UniqueGroup adds unique constraint for a group of columns.
//
//	BunTag: "unique:group"
func (t *Tag) UniqueGroup(group string) *Tag {
	t.uniqueGroup = &group
	return t
}

// NullZero marshals Go zero values as SQL NULL or DEFAULT (when supported).
//
//	BunTag: "nullzero"
func (t *Tag) NullZero() *Tag {
	t.nullZero = true
	return t
}

// ScanOnly tells bun to only use this field to scan query results and ignore in SELECT/INSERT/UPDATE/DELETE.
//
//	BunTag: "scanonly"
func (t *Tag) ScanOnly() *Tag {
	t.scanOnly = true
	return t
}

// PgArray tells bun to use a PostgreSQL array.
//
//	BunTag: "array"
func (t *Tag) PgArray() *Tag {
	t.pgArray = true
	return t
}

// JsonUseNumber uses json.Decoder.UseNumber to decode JSON.
//
//	BunTag: "json_use_number"
func (t *Tag) JsonUseNumber() *Tag {
	t.jsonUseNumber = true
	return t
}

// MsgPack uses MessagePack to encode/decode data.
//
//	BunTag: "msgpack"
func (t *Tag) MsgPack() *Tag {
	t.msgPack = true
	return t
}

// SoftDelete uses soft deletes on the model.
//
//	BunTag: "soft_delete
func (t *Tag) SoftDelete() *Tag {
	t.softDelete = true
	return t
}

func (t *Tag) Identity() *Tag {
	t.identity = true
	return t
}

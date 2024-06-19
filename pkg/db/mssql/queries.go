package mssql

const tableStatsQuery = `
SELECT obj.name AS [table_name], SUM(s.rows) AS [record_count], SCHEMA_NAME(obj.schema_id) AS [schema_name]
FROM sys.objects obj
INNER JOIN sys.partitions s ON obj.object_id = s.object_id
WHERE obj.type = 'U'
GROUP BY obj.schema_id, obj.name
ORDER BY [record_count]`

type tableInfoRecord struct {
	ColumnName        string
	DataType          string
	DataTypePrecision string
	DefaultValue      *string
	IsNullable        bool
	IsIdentity        bool
	IsPrimaryKey      *bool
	IsUnique          *bool
	FieldDescription  *string
}

const tableInfoQuery = `
	SELECT TableColumns.name            AS [column_name],
       ColumnDataType.name          AS [data_type],
        CASE
            WHEN ColumnDataType.[name] IN ('varchar', 'char') THEN  FORMATMESSAGE('%s(%s)', ColumnDataType.[name], IIF(TableColumns.max_length = -1, 'max', CAST(TableColumns.max_length AS VARCHAR(25)) ) )
            WHEN ColumnDataType.[name] IN ('nvarchar','nchar') THEN FORMATMESSAGE('%s(%d)', ColumnDataType.[name], IIF(TableColumns.max_length = -1, 'max', CAST(TableColumns.max_length / 2 AS VARCHAR(25))) )
            WHEN ColumnDataType.[name] IN ('decimal', 'numeric') THEN FORMATMESSAGE('%s(%d,%d)', ColumnDataType.[name], TableColumns.[precision], TableColumns.[scale])
            WHEN ColumnDataType.[name] IN ('datetime2') THEN ColumnDataType.[name] + '(' + CAST(TableColumns.[scale] AS VARCHAR(25)) + ')'
            ELSE ColumnDataType.[name]
        END                         AS [data_type_precision],
        ColumnDefaults.[definition] AS [default_value],
        TableColumns.[is_nullable],
        TableColumns.[is_identity],
        ColumnPrimaryKey.is_primary_key,
        ColumnUniqueKey.is_unique_constraint AS [is_unique],
        ColumnProperties.[value]    AS [field_description]

FROM sys.tables AS Tables

    INNER JOIN sys.columns AS TableColumns
         ON Tables.[object_id] = TableColumns.[object_id]

    LEFT JOIN sys.extended_properties AS ColumnProperties
        ON ColumnProperties.major_id = TableColumns.[object_id]
        AND ColumnProperties.minor_id = TableColumns.[column_id]
    INNER JOIN sys.types AS ColumnDataType
        ON TableColumns.user_type_id = ColumnDataType.user_type_id
    LEFT JOIN sys.default_constraints ColumnDefaults ON ColumnDefaults.parent_object_id = Tables.[object_id]
        AND ColumnDefaults.[parent_column_id] = TableColumns.[column_id]

    LEFT JOIN (
        SELECT SysIndexColumns.[object_id], SysIndexColumns.[column_id], SysIndexes.[is_primary_key], SysIndexes.[is_unique_constraint]
        FROM sys.index_columns SysIndexColumns
        INNER JOIN sys.indexes AS SysIndexes ON SysIndexes.[index_id] = SysIndexColumns.[index_id] AND SysIndexes.[object_id] = SysIndexColumns.[object_id] AND SysIndexes.[is_unique_constraint] = 0 AND SysIndexes.[is_primary_key] = 1 ) AS ColumnPrimaryKey
        ON ColumnPrimaryKey.object_id = TableColumns.[object_id] AND ColumnPrimaryKey.column_id = TableColumns.[column_id]

    LEFT JOIN (
            SELECT SysIndexColumns.[object_id], SysIndexColumns.[column_id], SysIndexes.[is_primary_key], SysIndexes.[is_unique_constraint]
            FROM sys.index_columns SysIndexColumns
            INNER JOIN sys.indexes AS SysIndexes ON SysIndexes.[index_id] = SysIndexColumns.[index_id] AND SysIndexes.[object_id] = SysIndexColumns.[object_id] AND SysIndexes.[is_unique_constraint] = 1 AND SysIndexes.[is_primary_key] = 0 ) AS ColumnUniqueKey
            ON ColumnUniqueKey.object_id = TableColumns.[object_id] AND ColumnUniqueKey.column_id = TableColumns.[column_id]

WHERE Tables.name = (?)
ORDER BY TableColumns.[column_id]`

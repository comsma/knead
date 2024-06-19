package mssql

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func (d *Database) Analyze() {
	rows, err := d.db.Query(tableStatsQuery)
	if err != nil {
		fmt.Println(err)
		return
	}
	var tableNodes []*yaml.Node

	for rows.Next() {
		var row = &tableNameQueryResult{}
		rows.Scan(&row.TableName, &row.RowCount, &row.TableSchema)
		tableNodes = append(tableNodes, &yaml.Node{
			Kind:        yaml.ScalarNode,
			Tag:         "!!str",
			Value:       row.TableName,
			LineComment: fmt.Sprintf("RecordCount:%d, Schema:%s", row.RowCount, row.TableSchema),
		})

	}
	var n = &yaml.Node{
		Kind:   yaml.DocumentNode,
		Line:   1,
		Column: 1,

		Content: []*yaml.Node{{
			Kind: yaml.MappingNode,
			Tag:  "!!map",
			Content: []*yaml.Node{
				{
					Kind:  yaml.ScalarNode,
					Value: "tables",
					Tag:   "!!str",
				}, {
					Kind:    yaml.SequenceNode,
					Tag:     "!!seq",
					Line:    1,
					Column:  1,
					Content: tableNodes,
				}}},
		},
	}

	y, err := yaml.Marshal(n)
	os.WriteFile("test.yml", y, 0644)

}

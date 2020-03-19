package query

import "t-storage/core/column"

type ColumnList []column.QueryColumn

func (cl *ColumnList)CreateTableColumn(tableName string, columnName string, alias string) *column.TableColumn {
	return &column.TableColumn{TableName:tableName, ColumnName:columnName, Alias:alias}
}
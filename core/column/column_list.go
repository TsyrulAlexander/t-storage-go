package column

type ColumnList []QueryColumn

func (cl *ColumnList)CreateTableColumn(tableName string, columnName string) *TableColumn {
	return &TableColumn{TableName: tableName, ColumnName:columnName}
}
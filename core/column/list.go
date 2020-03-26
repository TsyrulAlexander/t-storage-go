package column

type List []QueryColumn

func (cl *List)CreateTableColumn(tableName string, columnName string) *TableColumn {
	return &TableColumn{TableName: tableName, ColumnName:columnName}
}
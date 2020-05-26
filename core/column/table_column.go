package column

type TableColumn struct {
	TableName string
	ColumnName string
	Alias string
}

func CreateTableColumn(tableColumn string, columnName string) *TableColumn {

}

func (tc *TableColumn) GetAlias() string {
	if tc.Alias == "" {
		return tc.TableName + tc.ColumnName
	}
	return tc.Alias
}

package column

type TableColumn struct {
	TableName string
	ColumnName string
	Alias string
}

func (tc *TableColumn) GetAlias() string {
	if tc.Alias == "" {
		return tc.TableName + tc.ColumnName
	}
	return tc.Alias
}

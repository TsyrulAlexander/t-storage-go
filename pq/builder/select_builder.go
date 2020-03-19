package builder

type SelectBuilder struct {
	ColumnBuilder
	JoinBuilder
	ConditionBuilder
}
func (b *SelectBuilder) GetTableNameSql(tableName string) string {
	return GetTableFormat(tableName)
}
func (b *SelectBuilder) GetColumnNameSql(columnName string) string {
	return GetColumnFormat(columnName)
}
func (b *SelectBuilder) GetSelectCommandSql() string {
	return "SELECT"
}
func (b *SelectBuilder) GetFromCommandSql() string {
	return "FROM"
}
func (b *SelectBuilder) GetWhereCommandSql() string {
	return "WHERE"
}


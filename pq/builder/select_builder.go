package builder

type SelectBuilder struct {
	ColumnBuilder
	JoinBuilder
	ConditionBuilder
}
func (b *SelectBuilder) GetTableNameSql(tableName string) string {
	return GetTableNameWithFormat(tableName)
}
func (b *SelectBuilder) GetColumnNameSql(columnName string) string {
	return GetColumnNameWithFormat(columnName)
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


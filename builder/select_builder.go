package builder

type SelectBuilder interface {
	ColumnBuilder
	JoinBuilder
	ConditionBuilder
	GetTableNameSql(tableName string) string
	GetColumnNameSql(columnName string) string
	GetSelectCommandSql() string
	GetFromCommandSql() string
	GetWhereCommandSql() string
}

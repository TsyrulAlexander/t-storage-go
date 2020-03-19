package builder

import (
	"t-storage/core/column"
)

type ColumnBuilder struct {}

const aliasOperator = "AS"

func (b *ColumnBuilder) GetQueryColumnSql(c *column.QueryColumn) string {
	switch t := (*c).(type) {
	case *column.TableColumn:
		return getTableColumnSql(t)
	default:
		panic("not implemented")
	}
}

func (b *ColumnBuilder) GetQueryColumnAliasFormat(alias string) string {
	return "%v " + " AS \"" + alias + "\""
}

func (b *ColumnBuilder) GetColumnSeparatorSql() string {
	return ", "
}

func getTableColumnSql(c *column.TableColumn) string {
	return GetTableFormat(c.TableName) + "." + GetColumnFormat(c.ColumnName)
}


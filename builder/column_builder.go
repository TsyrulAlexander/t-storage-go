package builder

import "t-storage/core/column"

type ColumnBuilder interface {
	GetQueryColumnSql(column *column.QueryColumn) string
	GetQueryColumnAliasFormat(alias string) string
	GetColumnSeparatorSql() string
}
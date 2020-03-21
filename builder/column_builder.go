package builder

import "github.com/tsyrul-alexander/go-query-builder/core/column"

type ColumnBuilder interface {
	GetQueryColumnSql(column *column.QueryColumn) string
	GetQueryColumnAliasFormat(alias string) string
	GetColumnSeparatorSql() string
}
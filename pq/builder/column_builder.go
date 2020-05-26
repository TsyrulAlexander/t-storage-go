package builder

import (
	"github.com/tsyrul-alexander/go-query-builder/builder"
	"github.com/tsyrul-alexander/go-query-builder/core/column"
	"github.com/tsyrul-alexander/go-query-builder/core/column/function"
)

type ColumnBuilder struct {
	FunctionBuilder builder.FunctionBuilder
}

const aliasOperator = "AS"

func (b *ColumnBuilder) GetQueryColumnSql(c column.QueryColumn) string {
	switch t := (c).(type) {
	case *column.TableColumn:
		return b.getTableColumnSql(t)
	case *function.FunctionColumn:
		return b.getFunctionColumnSql(t)
	default:
		panic("not implemented")
	}
}

func (b *ColumnBuilder) getFunctionColumnSql(fc *function.FunctionColumn) string {
	return b.FunctionBuilder.GetFunctionSql(fc.Function)
}

func (b *ColumnBuilder) GetQueryColumnAliasFormat(alias string) string {
	return "%v " + aliasOperator + " \"" + alias + "\""
}

func (b *ColumnBuilder) GetColumnSeparatorSql() string {
	return ", "
}

func (b *ColumnBuilder) getTableColumnSql(c *column.TableColumn) string {
	return GetTableNameWithFormat(c.TableName) + "." + GetColumnNameWithFormat(c.ColumnName)
}


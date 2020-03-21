package builder

import (
	"github.com/tsyrul-alexander/go-query-builder/builder"
	"github.com/tsyrul-alexander/go-query-builder/core/parameter"
)

type InsertBuilder struct {
	builder.ParameterBuilder
}

const separate string = ", "

func (b *InsertBuilder) GetInsertSql(tableName string, columnValues *map[string]parameter.QueryParameter) string  {
	var c, v = b.getColumnNamesAndValuesSql(columnValues)
	return b.getInsertCommandSql(tableName, c) + " " + b.getValuesCommandSql(v)
}

func (b *InsertBuilder) getInsertCommandSql(tableName string, columnsSql string) string {
	return "INSERT INTO " + GetTableNameWithFormat(tableName) + "(" + columnsSql + ")"
}

func (b *InsertBuilder) getValuesCommandSql(valuesSql string) string {
	return "VALUES(" + valuesSql + ")"
}

func (b *InsertBuilder) getColumnNamesAndValuesSql(columnValues *map[string]parameter.QueryParameter) (columns string, values string) {
	var columnsSql = ""
	var valuesSql = ""
	var columnValuesCount = len(*columnValues)
	for key, value := range *columnValues {
		columnsSql += b.getColumnNameSql(key)
		valuesSql += b.getColumnValueSql(&value)
		columnValuesCount--
		if columnValuesCount != 0 {
			columnsSql += separate
			valuesSql += separate
		}
	}
	return columnsSql, valuesSql
}

func (b *InsertBuilder) getColumnNameSql(columnName string) string {
	return GetColumnNameWithFormat(columnName)
}

func (b *InsertBuilder) getColumnValueSql(p *parameter.QueryParameter) string {
	return b.GetParameterSql(p)
}

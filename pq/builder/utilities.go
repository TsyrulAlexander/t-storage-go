package builder

import (
	"github.com/tsyrul-alexander/go-query-builder/core/column"
	"github.com/tsyrul-alexander/go-query-builder/core/condition"
	"github.com/tsyrul-alexander/go-query-builder/core/join"
	"github.com/tsyrul-alexander/go-query-builder/query"
)

const (
	DefaultSelectRowCount = 10
	ParameterSeparator = ","
)

func CreateSelect(tableName string) *query.Select {
	var functionBuilder = &FunctionBuilder{}
	var columnBuilder = &ColumnBuilder{FunctionBuilder: functionBuilder}
	var parameterBuilder = &ParameterBuilder{ParameterSeparator:ParameterSeparator}
	functionBuilder.QueryColumnBuilderFunc = columnBuilder.GetQueryColumnSql
	functionBuilder.QueryParameterBuilderFunc = parameterBuilder.GetParameterSql
	var conditionBuilder = &ConditionBuilder{ColumnBuilder:columnBuilder, ParameterBuilder: parameterBuilder}
	var joinBuilder = &JoinBuilder{ConditionBuilder: conditionBuilder}
	var selectBuilder = &SelectBuilder{ColumnBuilder:columnBuilder, ConditionBuilder: conditionBuilder, JoinBuilder: joinBuilder}
	return &query.Select{
		TableName:  tableName,
		Builder:    selectBuilder,
		Columns:    &column.List{},
		Joins:      &join.List{},
		Conditions: &condition.List{},
		RowCount:   DefaultSelectRowCount,
	}
}

func CreateInsert(tableName string, columnValues *column.ValueList) *query.Insert {
	var parameterBuilder = &ParameterBuilder{ParameterSeparator:ParameterSeparator}
	var insertBuilder = &InsertBuilder{ParameterBuilder: parameterBuilder}
	return &query.Insert{
		TableName: tableName,
		Builder: insertBuilder,
		ColumnValues: columnValues,
	}
}

func GetTableNameWithFormat(tableName string) string {
	return "\"" + tableName + "\""
}

func GetColumnNameWithFormat(columnName string) string {
	return "\"" + columnName + "\""
}
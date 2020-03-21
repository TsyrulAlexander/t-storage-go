package builder

import (
	"t-storage/core/column"
	"t-storage/core/condition"
	"t-storage/core/join"
	"t-storage/query"
)

const defaultSelectRowCount = 10

func CreateSelect(tableName string) *query.Select {
	var columnBuilder = ColumnBuilder{}
	var parameterBuilder = &ParameterBuilder{}
	var conditionBuilder = ConditionBuilder{ColumnBuilder:&columnBuilder, ParameterBuilder: parameterBuilder}
	var joinBuilder = JoinBuilder{ConditionBuilder: &conditionBuilder}
	var selectBuilder = &SelectBuilder{ColumnBuilder:columnBuilder, ConditionBuilder: conditionBuilder, JoinBuilder: joinBuilder}
	return &query.Select{
		TableName:  tableName,
		Builder:    selectBuilder,
		Columns:    &column.ColumnList{},
		Joins:      &join.JoinList{},
		Conditions: &condition.ConditionList{},
		RowCount:   defaultSelectRowCount,
	}
}

func CreateInsert(tableName string, columnValues *column.ColumnValueList) *query.Insert {
	var parameterBuilder = &ParameterBuilder{}
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
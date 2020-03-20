package builder

import (
	"t-storage/query"
)

func CreateSelect(tableName string) *query.Select {
	var columnBuilder = ColumnBuilder{}
	var parameterBuilder = &ParameterBuilder{}
	var conditionBuilder = ConditionBuilder{ColumnBuilder:&columnBuilder, ParameterBuilder: parameterBuilder}
	var joinBuilder = JoinBuilder{ConditionBuilder: &conditionBuilder}
	var selectBuilder = &SelectBuilder{ColumnBuilder:columnBuilder, ConditionBuilder: conditionBuilder, JoinBuilder: joinBuilder}
	return &query.Select{
		TableName: tableName,
		SelectBuilder: selectBuilder,
		Columns: &query.ColumnList{},
		Joins: &query.JoinList{},
		Conditions: &query.ConditionList{},
	}
}

func CreateInsert(tableName string, columnValues *query.ColumnValueList) *query.Insert {
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
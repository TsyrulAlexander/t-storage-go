package builder

import (
	"t-storage/query"
)

func CreateSelect(tableName string) *query.Select {
	var columnBuilder = ColumnBuilder{}
	var conditionBuilder = ConditionBuilder{ColumnBuilder:&columnBuilder}
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

func GetTableFormat(tableName string) string {
	return "\"" + tableName + "\""
}

func GetColumnFormat(columnName string) string {
	return "\"" + columnName + "\""
}
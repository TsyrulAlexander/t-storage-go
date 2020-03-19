package query

import (
	"t-storage/core/column"
	"t-storage/core/condition"
)

const (
	newLine = "\n"
	emptyString = ""
)

func (s *Select)AddTableColumn(tableName string, columnName string) *column.TableColumn {
	var c = s.Columns.CreateTableColumn(tableName, columnName, "")
	var cl = append(*s.Columns, c)
	s.Columns = &cl
	return c
}

func (s *Select)AddColumnValueCondition(comparisonType condition.ComparisonType, tableName string,
		columnName string, value interface{}) condition.QueryCondition {
	var c = s.Columns.CreateTableColumn(tableName, columnName, "")
	var cond = s.Conditions.CreateColumnValueCondition(comparisonType, c, value)
	var cl = append(*s.Conditions, cond)
	s.Conditions = &cl
	return cond
}

func (s *Select) AddConditionGroup(logicalOperation condition.LogicalOperation,
		conditions ...condition.QueryCondition) *condition.GroupQueryCondition {
	return &condition.GroupQueryCondition{
		LogicalOperation:logicalOperation,
		QueryConditions: conditions[:],
	}
}



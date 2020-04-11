package query

import (
	"github.com/tsyrul-alexander/go-query-builder/core/column"
	"github.com/tsyrul-alexander/go-query-builder/core/condition"
	"github.com/tsyrul-alexander/go-query-builder/core/join"
)

func (s *Select)AddTableColumn(tableName string, columnName string) *column.TableColumn {
	var c = s.Columns.CreateTableColumn(tableName, columnName)
	var cl = append(*s.Columns, c)
	s.Columns = &cl
	return c
}

func (s *Select)AddColumnValueCondition(comparisonType condition.ComparisonType, tableName string,
		columnName string, value interface{}) condition.QueryCondition {
	var c = s.Columns.CreateTableColumn(tableName, columnName)
	var cond = s.Conditions.CreateColumnValueCondition(comparisonType, c, value)
	var cl = append(*s.Conditions, cond)
	s.Conditions = &cl
	return cond
}

func (s *Select) AddCondition(condition condition.QueryCondition) condition.QueryCondition {
	var c = append(*s.Conditions, condition)
	s.Conditions = &c
	return condition
}

func (s *Select) AddConditionGroup(logicalOperation condition.LogicalOperation,
		conditions ...condition.QueryCondition) *condition.GroupQueryCondition {
	var cg = &condition.GroupQueryCondition{
		LogicalOperation:logicalOperation,
		QueryConditions: conditions[:],
	}
	var c = append(*s.Conditions, cg)
	s.Conditions = &c
	return cg
}

func (s *Select) AddLeftJoin(joinTableName string, joinTableColumnName string, mainTableName string, mainTableColumnName string) *join.TableJoin {
	var j = s.Joins.CreateLeftJoin(joinTableName, joinTableColumnName, mainTableName, mainTableColumnName)
	var jl = append(*s.Joins, *j)
	s.Joins = &jl
	return j
}



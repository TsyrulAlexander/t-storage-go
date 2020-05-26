package query

import (
	"github.com/tsyrul-alexander/go-query-builder/core/column"
	columnFunction "github.com/tsyrul-alexander/go-query-builder/core/column/function"
	"github.com/tsyrul-alexander/go-query-builder/core/condition"
	"github.com/tsyrul-alexander/go-query-builder/core/function"
	"github.com/tsyrul-alexander/go-query-builder/core/function/value"
	"github.com/tsyrul-alexander/go-query-builder/core/join"
	"github.com/tsyrul-alexander/xz-data-api/model/data/culture"
)

const LczTablePrefix string = "Lcz"

func (s *Select) AddLocalizeColumn(tableName string, columnName string, cultureId culture.CultureId) column.QueryColumn {
	var lczTableName = tableName + LczTablePrefix
	s.AddLczJoin(tableName, lczTableName, cultureId)
	var col = s.Columns.CreateTableColumn(tableName, columnName)
	var lczColumn = s.Columns.CreateTableColumn(lczTableName, columnName)
	var coalesceFunction = function.CreateCoalesceFunction(value.CreateColumnValue(lczColumn), value.CreateColumnValue(col))
	var queryFunc = columnFunction.CreateFunctionColumn(coalesceFunction, tableName + columnName)
	var cl = append(*s.Columns, queryFunc)
	s.Columns = &cl
	return queryFunc
}

func (s *Select) AddLczJoin(tableName string, lczTableName string, cultureId culture.CultureId) {
	var lczJoin = &join.TableJoin{
		JoinTableName: lczTableName,
		Type:          join.LeftJoin,
		Conditions:    &condition.GroupQueryCondition{
			LogicalOperation: condition.And,
			QueryConditions:  []condition.QueryCondition{
				&condition.BinaryQueryCondition{
					ComparisonType: condition.ComparisonTypeEqual,
					LeftCondition:  &condition.ColumnQueryCondition{
						QueryColumn: column.CreateTableColumn(lczTableName, "RecordId"),
					},
					RightCondition: &condition.ColumnQueryCondition{
						QueryColumn: column.CreateTableColumn(tableName, "Id"),
					},
				},
				s.Conditions.CreateColumnValueCondition(condition.ComparisonTypeEqual,
					column.CreateTableColumn(lczTableName, "CultureId"), cultureId),
			},
		},
	}
	//	s.Joins.CreateLeftJoin(lczTableName, "RecordId", tableName, "Id")
	if !s.Joins.GetIfExistsJoin(lczJoin) {
		s.AddJoin(lczJoin)
	}
}

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
	s.AddJoin(j)
	return j
}

func (s *Select) AddJoin(join *join.TableJoin) *join.TableJoin {
	var jl = append(*s.Joins, *join)
	s.Joins = &jl
	return join
}



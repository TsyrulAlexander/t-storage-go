package join

import (
	"github.com/tsyrul-alexander/go-query-builder/core/column"
	"github.com/tsyrul-alexander/go-query-builder/core/condition"
)

type List []TableJoin

func (jl *List) CreateLeftJoin(joinTableName string, joinTableColumnName string, mainTableName string, mainTableColumnName string) *TableJoin {
	return CreateLeftJoin(joinTableName, joinTableColumnName, mainTableName,mainTableColumnName)
}

func (jl *List) GetIfExistsJoin(j *TableJoin) bool {
	var alias = j.GetAlias()
	for _, join := range *jl {
		if join != *j && join.GetAlias() == alias {
			return true
		}
	}
	return false;
}

func CreateLeftJoin(joinTableName string, joinTableColumnName string, mainTableName string, mainTableColumnName string) *TableJoin {
	return &TableJoin{
		JoinTableName: joinTableName,
		Type:          LeftJoin,
		Conditions:    &condition.BinaryQueryCondition{
			ComparisonType: condition.ComparisonTypeEqual,
			LeftCondition:  &condition.ColumnQueryCondition{
				QueryColumn: &column.TableColumn{
					TableName:  joinTableName,
					ColumnName: joinTableColumnName,
				},
			},
			RightCondition: &condition.ColumnQueryCondition{
				QueryColumn: &column.TableColumn{
					TableName:  mainTableName,
					ColumnName: mainTableColumnName,
				},
			},
		},
	}
}
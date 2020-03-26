package join

import (
	"github.com/tsyrul-alexander/go-query-builder/core/column"
	"github.com/tsyrul-alexander/go-query-builder/core/condition"
)

type List []TableJoin

func (j *List) CreateLeftJoin(joinTableName string, joinTableColumnName string, mainTableName string, mainTableColumnName string) *TableJoin {
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

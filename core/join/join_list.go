package join

import (
	"t-storage/core/column"
	"t-storage/core/condition"
)

type JoinList []TableJoin

func (j *JoinList) CreateLeftJoin(joinTableName string, joinTableColumnName string, mainTableName string, mainTableColumnName string) *TableJoin {
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

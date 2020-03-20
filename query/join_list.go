package query

import (
	"t-storage/core/column"
	"t-storage/core/condition"
	"t-storage/core/join"
)

type JoinList []join.TableJoin

func (j *JoinList) CreateLeftJoin(joinTableName string, joinTableColumnName string, mainTableName string, mainTableColumnName string) *join.TableJoin {
	return &join.TableJoin{
		JoinTableName: joinTableName,
		Type:          join.LeftJoin,
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

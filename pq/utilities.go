package pq

import (
	"t-storage/query"
)

func CreateSelect() *query.Select {
	var columnBuilder = builder.ColumnBuilder{}
	var conditionBuilder = builder.ConditionBuilder{ColumnBuilder:&columnBuilder}
	var joinBuilder = JoinBuilder{ConditionBuilder: &conditionBuilder}
	var selectBuilder = &SelectBuilder{ColumnBuilder:columnBuilder, ConditionBuilder: conditionBuilder, JoinBuilder: joinBuilder}
	return &query.Select{SelectBuilder: selectBuilder}
}

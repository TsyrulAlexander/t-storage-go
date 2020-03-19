package query

import (
	"t-storage/core/column"
	"t-storage/core/condition"
	"t-storage/core/parameter"
)

type ConditionList []condition.QueryCondition

func (cl *ConditionList) CreateColumnValueCondition(comparisonType condition.ComparisonType, c column.QueryColumn,
		value interface{}) *condition.BinaryQueryCondition {
	var v = getParameterValue(value)
	return &condition.BinaryQueryCondition{
		ComparisonType: comparisonType,
		LeftCondition: &condition.ColumnQueryCondition{
			QueryColumn: &c,
		},
		RightCondition: &condition.ParameterQueryCondition{
			Value: &v,
		},
	}
}

func getParameterValue(value interface{}) parameter.QueryParameter {
	switch t := value.(type) {
	case string:
		return &parameter.StringParameter{Value:t}
	default:
		panic("not implemented")
	}
}

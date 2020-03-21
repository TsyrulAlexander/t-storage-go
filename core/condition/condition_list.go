package condition

import (
	"github.com/tsyrul-alexander/go-query-builder/core/column"
	"github.com/tsyrul-alexander/go-query-builder/core/parameter"
)

type ConditionList []QueryCondition

func (cl *ConditionList) CreateColumnValueCondition(comparisonType ComparisonType, c column.QueryColumn,
		value interface{}) *BinaryQueryCondition {
	var v = getParameterValue(value)
	return &BinaryQueryCondition{
		ComparisonType: comparisonType,
		LeftCondition: &ColumnQueryCondition{
			QueryColumn: c,
		},
		RightCondition: &ParameterQueryCondition{
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

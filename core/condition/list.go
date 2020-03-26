package condition

import (
	"github.com/google/uuid"
	"github.com/tsyrul-alexander/go-query-builder/core/column"
	"github.com/tsyrul-alexander/go-query-builder/core/parameter"
	"reflect"
)

type List []QueryCondition

func (cl *List) CreateColumnValueCondition(comparisonType ComparisonType, c column.QueryColumn,
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
	var rv = reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.String:
		return parameter.CreateStringParameter(rv.String())
	case reflect.Int:
		return parameter.CreateIntParameter(int(rv.Int()))
	case reflect.Array, reflect.Slice:
			if rv.Type() == reflect.TypeOf(uuid.UUID{}) {
				return parameter.CreateGuidParameter(rv.Interface().(uuid.UUID))
			}
			var parameters []parameter.QueryParameter
			for i := 0; i < rv.Len(); i++ {
				parameters = append(parameters, getParameterValue(rv.Index(i).Interface()))
			}
			return parameter.CreateParameters(parameters...)
	default:
			panic("not implemented")
	}
}

package value

import (
	"github.com/tsyrul-alexander/go-query-builder/core/column"
	"github.com/tsyrul-alexander/go-query-builder/core/parameter"
)

type Value struct {
	QueryColumn column.QueryColumn
	Parameter   parameter.QueryParameter
}

func CreateParameterValue(qp parameter.QueryParameter) *Value {
	return &Value{Parameter:qp}
}

func CreateColumnValue(qc column.QueryColumn) *Value {
	return &Value{QueryColumn:qc}
}

package condition

import (
	"github.com/tsyrul-alexander/go-query-builder/core/parameter"
)
type ParameterQueryCondition struct {
	Value *parameter.QueryParameter
}

func (c *ParameterQueryCondition) getType() Type  {
	return Parameter
}

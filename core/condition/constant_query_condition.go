package condition

import (
	"t-storage/core/parameter"
)
type ParameterQueryCondition struct {
	Value *parameter.QueryParameter
}

func (c *ParameterQueryCondition) getType() Type  {
	return Parameter
}

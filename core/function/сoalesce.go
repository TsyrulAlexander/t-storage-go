package function

import "github.com/tsyrul-alexander/go-query-builder/core/function/value"

type Coalesce struct {
	Values []*value.Value
}

func CreateCoalesceFunction(values ...*value.Value) *Coalesce {
	return &Coalesce{Values:values}
}

func (c *Coalesce) EmptyFunctionFunc() {

}

package function

import "github.com/tsyrul-alexander/go-query-builder/core/function"

type FunctionColumn struct {
	Function function.Function
	Alias string
}

func CreateFunctionColumn(f function.Function, alias string) *FunctionColumn {
	return &FunctionColumn{Function:f, Alias:alias}
}

func (fc *FunctionColumn) GetAlias() string {
	return fc.Alias
}

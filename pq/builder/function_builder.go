package builder

import (
	"github.com/tsyrul-alexander/go-query-builder/core/column"
	"github.com/tsyrul-alexander/go-query-builder/core/function"
	"github.com/tsyrul-alexander/go-query-builder/core/function/value"
	"github.com/tsyrul-alexander/go-query-builder/core/parameter"
)

type FunctionBuilder struct {
	QueryColumnBuilderFunc func(column column.QueryColumn) string
	QueryParameterBuilderFunc func (p parameter.QueryParameter) string
}

func (fb *FunctionBuilder) GetFunctionSql(f function.Function) string  {
	switch ft := f.(type) {
	case *function.Coalesce:
		return fb.GetCoalesceSql(ft)
	default:
		panic("not implemented")
	}
}

func (fb *FunctionBuilder) getFunctionValueSql(v *value.Value) string  {
	if v.QueryColumn != nil {
		return fb.QueryColumnBuilderFunc(v.QueryColumn)
	} else if v.Parameter != nil {
		return fb.QueryParameterBuilderFunc(v.Parameter)
	} else {
		panic("not implemented")
	}
}

func (fb *FunctionBuilder) GetCoalesceSql(c *function.Coalesce) string {
	var sql = "COALESCE " + startArgumentSql()
	for i, v := range c.Values {
		sql += fb.getFunctionValueSql(v)
		if i != (len(c.Values) - 1) {
			sql += separateArgumentSql()
		}
	}
	sql += endArgumentSql()
	return sql
}

func startArgumentSql() string {
	return "("
}

func endArgumentSql() string {
	return ")"
}

func separateArgumentSql() string {
	return ","
}

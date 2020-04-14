package builder

import "github.com/tsyrul-alexander/go-query-builder/core/parameter"

type ParameterBuilder struct {
	ParameterSeparator string
}

func (b *ParameterBuilder) GetParameterSql(p parameter.QueryParameter) string {
	switch t:= p.(type) {
	case *parameter.StringParameter, *parameter.GuidParameter:
		return "'" + t.GetValueSql() + "'"
	case *parameter.IntParameter:
		return t.GetValueSql()
	case *parameter.NullParameter:
		return "NULL"
	case *parameter.ArrayParameter:
		return b.getArrayParameterSql(t)
	default:
		panic("not implemented")
	}
}

func (b *ParameterBuilder)getArrayParameterSql(p *parameter.ArrayParameter) string {
	var testSql = ""
	for _, p := range p.Parameters {
		testSql += b.GetParameterSql(p)
	}
	return testSql
}
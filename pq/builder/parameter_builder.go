package builder

import "github.com/tsyrul-alexander/go-query-builder/core/parameter"

type ParameterBuilder struct {
}

func (b *ParameterBuilder) GetParameterSql(p *parameter.QueryParameter) string {
	switch t:= (*p).(type) {
	case *parameter.StringParameter, *parameter.GuidParameter:
		return "'" + t.GetValueSql() + "'"
	case *parameter.IntParameter:
		return t.GetValueSql()
	default:
		panic("not implemented")
	}
}

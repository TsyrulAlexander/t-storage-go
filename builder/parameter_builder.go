package builder

import "github.com/tsyrul-alexander/go-query-builder/core/parameter"

type ParameterBuilder interface {
	GetParameterSql(p *parameter.QueryParameter) string
}
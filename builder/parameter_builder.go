package builder

import "t-storage/core/parameter"

type ParameterBuilder interface {
	GetParameterSql(p *parameter.QueryParameter) string
}
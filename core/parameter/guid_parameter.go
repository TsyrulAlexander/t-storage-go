package parameter

import "github.com/google/uuid"

type GuidParameter struct {
	Value uuid.UUID
}

func (p *GuidParameter)GetValueSql() string {
	return p.Value.String()
}

func CreateGuidParameter(value uuid.UUID) *GuidParameter {
	return &GuidParameter{Value:value}
}

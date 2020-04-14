package parameter

import "github.com/google/uuid"

type GuidParameter struct {
	Value uuid.UUID
}

func (p *GuidParameter)GetValueSql() string {
	return p.Value.String()
}

func CreateGuidParameter(value uuid.UUID) QueryParameter {
	if value == uuid.Nil {
		return CreateNullParameter()
	}
	return &GuidParameter{Value:value}
}

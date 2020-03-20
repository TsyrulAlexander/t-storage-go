package parameter

type StringParameter struct {
	Value string
}

func (p *StringParameter)GetValueSql() string {
	return p.Value
}

func CreateStringParameter(value string) *StringParameter {
	return &StringParameter{Value:value}
}

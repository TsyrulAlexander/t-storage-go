package parameter

type StringParameter struct {
	Value string
}

func (p *StringParameter)GetValueSql() string {
	return p.Value
}

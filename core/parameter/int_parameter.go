package parameter

import "strconv"

type IntParameter struct {
	Value int
}

func (p *IntParameter)GetValueSql() string {
	return strconv.Itoa(p.Value)
}

func CreateIntParameter(value int) QueryParameter {
	return &IntParameter{Value:value}
}

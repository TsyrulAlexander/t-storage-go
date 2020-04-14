package parameter

type NullParameter struct {

}

func (p *NullParameter)GetValueSql() string {
	return ""
}

func CreateNullParameter() *NullParameter {
	return &NullParameter{}
}


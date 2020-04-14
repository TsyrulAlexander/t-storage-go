package join

import "github.com/tsyrul-alexander/go-query-builder/core/condition"

type TableJoin struct {
	JoinTableName string
	Alias         string
	Type          Type
	Conditions    condition.QueryCondition
}

func (j *TableJoin)GetAlias() string {
	if j.Alias == "" {
		return j.JoinTableName
	}
	return j.Alias
}

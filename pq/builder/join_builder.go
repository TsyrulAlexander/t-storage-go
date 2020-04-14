package builder

import (
	"github.com/tsyrul-alexander/go-query-builder/builder"
	"github.com/tsyrul-alexander/go-query-builder/core/join"
)

type JoinBuilder struct {
	builder.ConditionBuilder
}

func (b *JoinBuilder) GetJoinSql(j *join.TableJoin) string  {
	return getJoinTypeSql(j.Type) + " JOIN  \"" + j.JoinTableName + "\" AS \"" + j.GetAlias() + "\" ON " + b.GetQueryConditionSql(&j.Conditions)
}

func getJoinTypeSql(t join.Type) string {
	switch t {
	case join.LeftJoin:
		return "LEFT"
	case join.InnerJoin:
		return "INNER"
	default:
		panic("not implemented")
	}
}

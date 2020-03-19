package builder

import (
	"t-storage/builder"
	"t-storage/core/join"
)

type JoinBuilder struct {
	builder.ConditionBuilder
}

func (b *JoinBuilder) GetJoinSql(j *join.TableJoin) string  {
	return getJoinTypeSql(j.Type) + " JOIN  \"" + j.JoinTableName + "\" ON " + b.GetQueryConditionSql(&j.Conditions) + "\n"
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

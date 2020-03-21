package join

import "github.com/tsyrul-alexander/go-query-builder/core/condition"

type TableJoin struct {
	JoinTableName string
	Type          Type
	Conditions    condition.QueryCondition
}

package join

import "t-storage/core/condition"

type TableJoin struct {
	JoinTableName string
	Type          Type
	Conditions    condition.QueryCondition
}

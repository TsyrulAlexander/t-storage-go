package builder

import "t-storage/core/condition"

type ConditionBuilder interface {
	GetQueryConditionSql(c *condition.QueryCondition) string
	GetLogicalOperationSql(lo condition.LogicalOperation) string
}

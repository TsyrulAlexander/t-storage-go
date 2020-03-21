package builder

import "github.com/tsyrul-alexander/go-query-builder/core/condition"

type ConditionBuilder interface {
	GetQueryConditionSql(c *condition.QueryCondition) string
	GetLogicalOperationSql(lo condition.LogicalOperation) string
}

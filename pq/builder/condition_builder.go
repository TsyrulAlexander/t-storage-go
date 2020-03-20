package builder

import (
	"t-storage/builder"
	"t-storage/core/condition"
)

type ConditionBuilder struct {
	builder.ParameterBuilder
	builder.ColumnBuilder
}

func (b *ConditionBuilder) GetQueryConditionSql(c *condition.QueryCondition) string {
	switch t := (*c).(type) {
	case *condition.GroupQueryCondition:
		return b.getGroupQueryConditionSql(t)
	case *condition.BinaryQueryCondition:
		return b.getBinaryQueryConditionSql(t)
	case *condition.ColumnQueryCondition:
		return b.getColumnQueryConditionSql(t)
	case *condition.ParameterQueryCondition:
		return b.getParameterQueryConditionSql(t)
	default:
		panic("not implemented")
	}
}
func (b *ConditionBuilder) GetLogicalOperationSql(lo condition.LogicalOperation) string {
	switch lo {
	case condition.And:
		return "AND"
	case condition.Or:
		return "OR"
	default:
		panic("not implemented")
	}
}

func (b *ConditionBuilder) getParameterQueryConditionSql(c *condition.ParameterQueryCondition) string {
	return b.GetParameterSql(c.Value)
}

func (b *ConditionBuilder) getColumnQueryConditionSql(c *condition.ColumnQueryCondition) string {
	return b.ColumnBuilder.GetQueryColumnSql(&c.QueryColumn)
}

func (b *ConditionBuilder) getGroupQueryConditionSql(c *condition.GroupQueryCondition) string {
	var logicalOperatorSql = b.GetLogicalOperationSql(c.LogicalOperation)
	var sqlText = "("
	for index, cond := range c.QueryConditions {
		sqlText += b.GetQueryConditionSql(&cond)
		if index == (len(c.QueryConditions) - 1) {
			sqlText += logicalOperatorSql
		}
	}
	sqlText += ")"
	return sqlText
}

func (b *ConditionBuilder) getBinaryQueryConditionSql(c *condition.BinaryQueryCondition) string {
	var leftExpressionSql = b.GetQueryConditionSql(&c.LeftCondition)
	var rightExpressionSql = b.GetQueryConditionSql(&c.RightCondition)
	return b.getComparisonTypeSql(c.ComparisonType, leftExpressionSql, rightExpressionSql)
}

func (b *ConditionBuilder) getComparisonTypeSql(c condition.ComparisonType, leftExpression string, rightExpression string) string {
	switch c {
	case condition.ComparisonTypeEqual:
		return leftExpression + " = " + rightExpression
	case condition.ComparisonTypeNotEqual:
		return leftExpression + " != " + rightExpression
	default:
		panic("not implemented")
	}
}


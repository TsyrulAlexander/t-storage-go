package condition

import "github.com/tsyrul-alexander/go-query-builder/core/column"

type ColumnQueryCondition struct {
	QueryColumn column.QueryColumn
}

func (c *ColumnQueryCondition) getType() Type  {
	return Column
}
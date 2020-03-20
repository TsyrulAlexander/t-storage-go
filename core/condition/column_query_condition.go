package condition

import "t-storage/core/column"

type ColumnQueryCondition struct {
	QueryColumn column.QueryColumn
}

func (c *ColumnQueryCondition) getType() Type  {
	return Column
}
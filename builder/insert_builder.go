package builder

import "t-storage/core/parameter"

type InsertBuilder interface {
	GetInsertSql(tableName string, columnValues *map[string]parameter.QueryParameter) string
}

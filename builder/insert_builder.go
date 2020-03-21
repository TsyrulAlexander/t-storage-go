package builder

import "github.com/tsyrul-alexander/go-query-builder/core/parameter"

type InsertBuilder interface {
	GetInsertSql(tableName string, columnValues *map[string]parameter.QueryParameter) string
}

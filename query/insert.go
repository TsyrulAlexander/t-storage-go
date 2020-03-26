package query

import (
	"database/sql"
	"github.com/tsyrul-alexander/go-query-builder/builder"
	"github.com/tsyrul-alexander/go-query-builder/core/column"
	"github.com/tsyrul-alexander/go-query-builder/core/parameter"
)

type Insert struct {
	TableName string
	ColumnValues *column.ValueList
	Builder builder.InsertBuilder
}

func (i *Insert) GetSqlText() string {
	var columnValues = i.ColumnValues
	return i.Builder.GetInsertSql(i.TableName, (*map[string]parameter.QueryParameter)(columnValues))
}

func (i *Insert) Execute(db *sql.DB) (sql.Result, error) {
	var sqlText = i.GetSqlText()
	return db.Exec(sqlText)
}
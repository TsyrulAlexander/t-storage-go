package query

import (
	"database/sql"
	"t-storage/builder"
	"t-storage/core/column"
	"t-storage/core/parameter"
)

type Insert struct {
	TableName string
	ColumnValues *column.ColumnValueList
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
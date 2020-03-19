package query

import "t-storage/builder"

type Insert struct {
	TableName string
	Builder builder.InsertBuilder
}

func (i *Insert) GetSqlText() string {
	return "todo"
}
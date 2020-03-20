package query

import "github.com/google/uuid"

type Row map[string]interface{}

func (r *Row)GetUuidValue(columnName string) uuid.UUID {
	var value = (*r)[columnName]
	return (value).(uuid.UUID)
}
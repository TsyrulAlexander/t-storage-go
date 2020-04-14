package query

import (
	"github.com/google/uuid"
)

type Row map[string]interface{}

func (r *Row)GetUuidValue(columnName string) uuid.UUID {
	var value, exist = (*r)[columnName]
	if !exist || value == nil {
		return uuid.Nil
	}
	var v = value.([]uint8)
	var guid, _ = uuid.ParseBytes(v)
	return guid
}

func (r *Row)GetStringValue(columnName string) string {
	var value, exist = (*r)[columnName]
	if !exist || value == nil {
		return ""
	}
	return value.(string)
}

func (r *Row)GetIntValue(columnName string) int {
	var value, exist = (*r)[columnName]
	if !exist || value == nil {
		return 0
	}
	return int(value.(int64))
}
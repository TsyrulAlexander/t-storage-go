package query

import (
	"database/sql"
	"fmt"
	"t-storage/builder"
	"t-storage/core/condition"
)

type Select struct {
	TableName string
	Columns *ColumnList
	Joins *JoinList
	Conditions *ConditionList
	builder.SelectBuilder
}

func (s *Select) Execute(db *sql.DB) (*[]Row, error) {
	var sqlText = s.GetSqlText()
	var sqlRows, err = db.Query(sqlText)
	if err != nil {
		return nil, err
	}
	var rows, rowsError = getRows(sqlRows, s.Columns)
	if rowsError != nil {
		return  nil, err
	}
	return rows, err
}

func (s *Select) GetSqlText() string {
	return s.getSelectSql() + s.getFromSql() + s.getJoinsSql() + s.getWhereSql()
}

func getRows(rs *sql.Rows, columns *ColumnList) (*[]Row, error) {
	var rows []Row
	for rs.Next() {
		var row, err = getRow(rs, columns)
		if err != nil {
			return nil, err
		}
		rows = append(rows, *row)
	}
	return &rows, nil
}

func getRow(rs *sql.Rows, columns *ColumnList) (*Row, error) {
	var scanArray, row = getScanRow(columns)
	if err := rs.Scan(scanArray...); err != nil {
		return nil, err
	}
	return row, nil
}

func getScanRow(columns *ColumnList) ([]interface{}, *Row) {
	var scanArray = make([]interface{}, len(*columns))
	var row = make(Row, len(*columns))
	for i, v := range *columns {
		var columnAlias = v.GetAlias()
		scanArray[i] = row[columnAlias]
	}
	return scanArray, &row
}

func (s *Select) getColumnsSql() string {
	var sqlText = emptyString
	var columnSeparator = s.GetColumnSeparatorSql()
	var columnCount = len(*s.Columns)
	for i, c := range *s.Columns {
		sqlText += fmt.Sprintf(s.GetQueryColumnAliasFormat(c.GetAlias()), s.GetQueryColumnSql(&c))
		if i != (columnCount - 1) {
			sqlText += columnSeparator
		}
	}
	return sqlText
}

func (s *Select) getSelectSql() string {
	return s.GetSelectCommandSql() + " " + s.getColumnsSql() + newLine
}

func (s *Select) getFromSql() string {
	return s.GetFromCommandSql() + " " + s.GetTableNameSql(s.TableName) + newLine
}

func (s *Select) getJoinsSql() string {
	var sqlText = emptyString
	for _, j := range *s.Joins {
		sqlText += s.GetJoinSql(&j) + newLine
	}
	return sqlText
}

func (s *Select) getWhereSql() string {
	if len(*s.Conditions) == 0 {
		return emptyString
	}
	return s.GetWhereCommandSql() + " " + s.getConditionsSql() + newLine
}

func (s *Select) getConditionsSql() string {
	var sqlText = emptyString
	var conditionCount = len(*s.Conditions)
	for i, c := range *s.Conditions {
		sqlText += s.GetQueryConditionSql(&c)
		if i != (conditionCount - 1) {
			sqlText += s.GetLogicalOperationSql(condition.And)
		}
	}
	return sqlText
}
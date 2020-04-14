package query

import (
	"database/sql"
	"errors"
	"github.com/tsyrul-alexander/go-query-builder/core/column"
	"github.com/tsyrul-alexander/go-query-builder/core/condition"
	"github.com/tsyrul-alexander/go-query-builder/core/join"
	"strings"
)

type Select struct {
	TableName  string
	Columns    *column.List
	Joins      *join.List
	Conditions *condition.List
	RowCount   int
	RowOffset  int
	Builder    SelectBuilder
}

func (s *Select) Execute(db *sql.DB) (*[]Row, error) {
	if vErr := s.ValidateSelect(); vErr != nil {
		return nil, vErr
	}
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
	var sb = &strings.Builder{}
	s.Builder.AlterBuildSql(s, sb)
	s.Builder.SetSelectSql(s, sb)
	s.Builder.SetFromSql(s, sb)
	s.Builder.SetJoinSql(s, sb)
	s.Builder.SetWhereSql(s, sb)
	s.Builder.BeforeBuildSql(s, sb)
	return sb.String()
}

func (s *Select) ValidateSelect() error {
	return s.ValidateJoins()
}

func (s *Select) ValidateJoins() error {
	if s.Joins != nil {
		for _, j := range *s.Joins {
			if 	s.Joins.GetIfExistsJoin(&j) {
				return errors.New("joining the " + j.GetAlias() + " exists")
			}
		}
	}
	return nil
}

func getRows(rs *sql.Rows, columns *column.List) (*[]Row, error) {
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

func getRow(rs *sql.Rows, columns *column.List) (*Row, error) {
	var scanArray, valueArray = getScanRow(len(*columns))
	if err := rs.Scan(scanArray...); err != nil {
		return nil, err
	}
	return getRowValue(columns, valueArray), nil
}

func getScanRow(count int) (scan []interface{}, value []interface{}) {
	var scanArray = make([]interface{}, count)
	var valueArray = make([]interface{}, count)
	for i := 0; i < count; i++ {
		scanArray[i] = &valueArray[i]
	}
	return scanArray, valueArray
}

func getRowValue(columns *column.List, valueArray []interface{}) *Row {
	var r = Row{}
	for i, c := range *columns {
		var columnAlias = c.GetAlias()
		r[columnAlias] = valueArray[i]
	}
	return &r
}
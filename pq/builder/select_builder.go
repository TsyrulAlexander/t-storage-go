package builder

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/tsyrul-alexander/go-query-builder/core/condition"
	"github.com/tsyrul-alexander/go-query-builder/query"
)
const (
	newLine = "\n"
	spaceStr = " "
)
type SelectBuilder struct {
	*ColumnBuilder
	*JoinBuilder
	*ConditionBuilder
}
func (b *SelectBuilder) AlterBuildSql(_ *query.Select, _ *strings.Builder) {

}
func (b *SelectBuilder) BeforeBuildSql(s *query.Select, sb *strings.Builder) {
	b.SetLimitSql(s, sb)
}
func (b *SelectBuilder) SetSelectSql(s *query.Select, sb *strings.Builder) {
	sb.WriteString(b.getSelectCommandSql())
	sb.WriteString(spaceStr)
	b.setColumnsSql(s, sb)
	sb.WriteString(newLine)
}
func (b *SelectBuilder) SetFromSql(s *query.Select, sb *strings.Builder) {
	sb.WriteString(b.getFromCommandSql())
	sb.WriteString(spaceStr)
	sb.WriteString(GetTableNameWithFormat(s.TableName))
	sb.WriteString(newLine)
}
func (b *SelectBuilder) SetJoinSql(s *query.Select, sb *strings.Builder) {
	for _, j := range *s.Joins {
		sb.WriteString(s.Builder.GetJoinSql(&j))
		sb.WriteString(newLine)
	}
}
func (b *SelectBuilder) SetWhereSql(s *query.Select, sb *strings.Builder) {
	if len(*s.Conditions) == 0 {
		return
	}
	sb.WriteString(b.getWhereCommandSql())
	sb.WriteString(spaceStr)
	b.setConditionsSql(s, sb)
	sb.WriteString(newLine)
}
func (b *SelectBuilder) SetLimitSql(s *query.Select, sb *strings.Builder) {
	if s.RowCount == 0 {
		return
	}
	sb.WriteString(b.getLimitCommandSql())
	sb.WriteString(spaceStr)
	sb.WriteString(strconv.Itoa(s.RowCount))
	sb.WriteString(newLine)
}
func (b *SelectBuilder) getColumnNameSql(columnName string) string {
	return GetColumnNameWithFormat(columnName)
}
func (b *SelectBuilder) setColumnsSql(s *query.Select, sb *strings.Builder) {
	var columnSeparator = b.GetColumnSeparatorSql()
	var columnCount = len(*s.Columns)
	for i, c := range *s.Columns {
		sb.WriteString(fmt.Sprintf(b.GetQueryColumnAliasFormat(c.GetAlias()), b.GetQueryColumnSql(&c)))
		if i != (columnCount - 1) {
			sb.WriteString(columnSeparator)
		}
	}
}
func (b *SelectBuilder) setConditionsSql(s *query.Select, sb *strings.Builder) {
	var conditionCount = len(*s.Conditions)
	for i, c := range *s.Conditions {
		sb.WriteString(s.Builder.GetQueryConditionSql(&c))
		if i != (conditionCount - 1) {
			sb.WriteString(s.Builder.GetLogicalOperationSql(condition.And))
		}
	}
}
func (b *SelectBuilder) getSelectCommandSql() string {
	return "SELECT"
}
func (b *SelectBuilder) getFromCommandSql() string {
	return "FROM"
}
func (b *SelectBuilder) getWhereCommandSql() string {
	return "WHERE"
}
func (b *SelectBuilder) getLimitCommandSql() string {
	return "LIMIT"
}


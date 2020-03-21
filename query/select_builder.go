package query

import (
	"strings"
	"github.com/tsyrul-alexander/go-query-builder/builder"
)

type SelectBuilder interface {
	builder.ColumnBuilder
	builder.JoinBuilder
	builder.ConditionBuilder
	AlterBuildSql(s *Select, sb *strings.Builder)
	SetSelectSql(s *Select, sb *strings.Builder)
	SetFromSql(s *Select, sb *strings.Builder)
	SetJoinSql(s *Select, sb *strings.Builder)
	SetWhereSql(s *Select, sb *strings.Builder)
	BeforeBuildSql(s *Select, sb *strings.Builder)
}

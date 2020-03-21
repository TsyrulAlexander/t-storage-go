package builder

import "github.com/tsyrul-alexander/go-query-builder/core/join"

type JoinBuilder interface {
	GetJoinSql(join *join.TableJoin) string
}

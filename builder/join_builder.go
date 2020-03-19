package builder

import "t-storage/core/join"

type JoinBuilder interface {
	GetJoinSql(join *join.TableJoin) string
}

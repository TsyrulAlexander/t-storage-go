package builder

import (
	"github.com/tsyrul-alexander/go-query-builder/core/function"
)

type FunctionBuilder interface {
	GetFunctionSql(f function.Function) string
}

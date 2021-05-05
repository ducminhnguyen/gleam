package executor

import (
	"github.com/ducminhnguyen/gleam/flow"
	"github.com/ducminhnguyen/gleam/sql/expression"
)

type Executor interface {
	Exec() *flow.Dataset
	Schema() expression.Schema
}

package executor

import (
	"github.com/ducminhnguyen/gleam/flow"
	"github.com/ducminhnguyen/gleam/sql/model"
)

type TableColumn struct {
	ColumnName string
	ColumnType byte
}

type TableSource struct {
	Dataset   *flow.Dataset
	TableInfo *model.TableInfo
}

var (
	Tables = make(map[string]*TableSource)
)

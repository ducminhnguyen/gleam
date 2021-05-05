package main

import (
	. "github.com/ducminhnguyen/gleam/flow"
	"github.com/ducminhnguyen/gleam/gio"
	"github.com/ducminhnguyen/gleam/plugins/file"
)

func main() {
	gio.Init()
	f := New("Read parquet file")
	a := f.Read(file.Parquet("a.parquet", 3)).Select("select", Field(1, 2, 3, 4, 5))
	a.Printlnf("%v, %v, %v, %v, %v").Run()

}

package io_test

import (
	"github.com/huzhouv/gotools/io"
	"testing"
)

func TestExistsDir(t *testing.T) {
	lg := testool.Wrap(t)

	lg.Case("give a existing dir")
	f := "/Users/sam/workspace/mine/gotools/io/"
	lg.Require(io.Dir.Exists(f), "should exist")

	lg.Case("give a existing file, but not a director")
	f = "/Users/sam/workspace/mine/gotools/io/file_test.go"
	lg.Require(!io.Dir.Exists(f), "should not exist")
}

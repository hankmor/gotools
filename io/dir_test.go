package io_test

import (
	"github.com/huzhouv/gotools/io"
	"github.com/huzhouv/gotools/tester"
	"testing"
)

func TestExistsDir(t *testing.T) {
	lg := tester.Wrap(t)

	lg.Case("give a existing dir")
	f := "/Users/sam/workspace/mine/gotools/io/"
	lg.Require(io.ExistsDir(f), "should exist")

	lg.Case("give a existing file, but not a director")
	f = "/Users/sam/workspace/mine/gotools/io/file_test.go"
	lg.Require(!io.ExistsDir(f), "should not exist")
}

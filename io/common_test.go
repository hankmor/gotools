package io_test

import (
	"github.com/huzhouv/gotools/io"
	"github.com/huzhouv/gotools/tester"
	"testing"
)

func TestExists(t *testing.T) {
	lg := tester.Wrap(t)

	lg.Case("give a existing dir")
	f := "/Users/sam/workspace/mine/gotools/io/"
	lg.Require(io.Exists(f), "should exist")

	lg.Case("give a existing file")
	f = "/Users/sam/workspace/mine/gotools/io/file_test.go"
	lg.Require(io.Exists(f), "should exist")
}

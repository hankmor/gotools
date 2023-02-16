package io_test

import (
	"github.com/huzhouv/gotools/io"
	"github.com/huzhouv/gotools/testool"
	"testing"
)

func TestExistsFile(t *testing.T) {
	lg := testool.Wrap(t)

	lg.Case("give a existing file")
	f := "/Users/sam/workspace/mine/gotools/io/file_test.go"
	lg.Require(io.File.Exists(f), "should exist")

	lg.Case("give a existing dir, but is not a file")
	f = "/Users/sam/workspace/mine/gotools/io/"
	lg.Require(!io.File.Exists(f), "should not exist")
}

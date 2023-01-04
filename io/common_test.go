package io_test

import (
	"gotools/io"
	"gotools/testool"
	"testing"
)

func TestExists(t *testing.T) {
	lg := testool.Wrap(t)

	lg.Case("give a existing dir")
	f := "/Users/sam/workspace/mine/gotools/io/"
	lg.Require(io.Exists(f), "should exist")

	lg.Case("give a existing file")
	f = "/Users/sam/workspace/mine/gotools/io/file_test.go"
	lg.Require(io.Exists(f), "should exist")
}

func TestIsDir(t *testing.T) {
	lg := testool.Wrap(t)

	lg.Case("give a existing dir")
	f := "/Users/sam/workspace/mine/gotools/io/"
	lg.Require(io.IsDir(f), "is dir")

	lg.Case("give a existing file")
	f = "/Users/sam/workspace/mine/gotools/io/file_test.go"
	lg.Require(!io.IsDir(f), "is not a dir")
}

func TestIsRegularFile(t *testing.T) {
	lg := testool.Wrap(t)

	lg.Case("give a existing dir")
	f := "/Users/sam/workspace/mine/gotools/io/"
	lg.Require(!io.IsRegularFile(f), "is not a regular file")

	lg.Case("give a existing regular file")
	f = "/Users/sam/workspace/mine/gotools/io/file_test.go"
	lg.Require(io.IsRegularFile(f), "is a regular file")

	lg.Case("give a soft symlink file")
	f = "/etc"
	lg.Require(!io.IsRegularFile(f), "is not a regular file")
}

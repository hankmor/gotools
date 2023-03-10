package io_test

import (
	"fmt"
	"github.com/hankmor/gotools/io"
	"github.com/hankmor/gotools/testool"
	"os"
	"path/filepath"
	"testing"
)

func TestPathExists(t *testing.T) {
	lg := testool.Wrap(t)

	lg.Case("give an exists path")
	path := "/Users/sam/workspace/mine/gotools"
	lg.Require(io.Path.PathExists(path), "given path should exist")

	lg.Case("give an none exists path")
	path = "/Users/haha"
	lg.Require(!io.Path.PathExists(path), "given path should not exist")
}

func TestExecPath(t *testing.T) {
	execpath, err := os.Executable() // 获得程序路径
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(execpath)
	fmt.Println(dir)

	s, _ := os.Getwd()
	println(s)

	println(io.Path.ExecPath())
	println(io.Path.CurrentPath())
	println(io.Path.ProjectPath())
}

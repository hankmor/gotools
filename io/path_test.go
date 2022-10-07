package io_test

import (
	"fmt"
	"github.com/huzhouv/gotools/io"
	"github.com/huzhouv/gotools/tester"
	"os"
	"path/filepath"
	"testing"
)

func TestPathExists(t *testing.T) {
	lg := tester.Wrap(t)

	lg.Case("give an exists path")
	path := "/Users/sam/workspace/mine/gotools"
	lg.Require(io.PathExists(path), "given path should exist")

	lg.Case("give an none exists path")
	path = "/Users/haha"
	lg.Require(!io.PathExists(path), "given path should not exist")
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

	println(io.ExecPath())
	println(io.CurrentPath())
	println(io.ProjectPath())
}

package io

import "os"

var Dir = &dirs{}

type dirs struct {
}

func (d *dirs) Exists(dir string) bool {
	b := Exists(dir)
	if b {
		f, _ := os.Stat(dir)
		return f.IsDir()
	}
	return b
}

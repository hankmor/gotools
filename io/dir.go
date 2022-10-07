package io

import "os"

func ExistsDir(dir string) bool {
	b := Exists(dir)
	if b {
		f, _ := os.Stat(dir)
		return f.IsDir()
	}
	return b
}

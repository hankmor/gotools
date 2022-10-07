package io

import "os"

func ExistsFile(file string) bool {
	b := Exists(file)
	if b {
		f, _ := os.Stat(file)
		return !f.IsDir()
	}
	return b
}

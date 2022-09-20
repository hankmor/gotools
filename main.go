package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	h := sha256.New()
	h.Write([]byte("hello world"))
	fmt.Printf("%x\n", h.Sum(nil))

	h1 := sha256.New()
	h1.Write([]byte(""))
	fmt.Printf("%x\n", h1.Sum(nil))
}

package random

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRandomAlphabetic(t *testing.T) {
	for i := 0; i < 10; i++ {
		n := rand.Intn(32) + 1
		s := Alphabetic(n)
		if n != len(s) {
			t.Fatalf("test failed")
		}
		fmt.Printf("%d : %s\n", i, s)
	}
}

func TestRandomNumber(t *testing.T) {
	for i := 0; i < 10; i++ {
		n := rand.Intn(32) + 1
		s := Numeric(n)
		if n != len(s) {
			t.Fatalf("test failed")
		}
		fmt.Printf("%d : %s\n", i, s)
	}
}

func TestRandomAlphanumeric(t *testing.T) {
	for i := 0; i < 10; i++ {
		n := rand.Intn(32) + 1
		s := Alphanumeric(n)
		if n != len(s) {
			t.Fatalf("test failed")
		}
		fmt.Printf("%d : %s\n", i, s)
	}
}

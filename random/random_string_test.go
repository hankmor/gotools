package random

import (
	"gotools/tester"
	"math/rand"
	"testing"
	"time"
)

func TestRandomAlphabetic(t *testing.T) {
	tl := tester.Wrap(t)
	tl.Case("loop 10 times to generate random string")
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		n := rand.Intn(32) + 1
		s := Alphabetic(n)
		tl.Require(n == len(s), "length of generated string should be %d", n)
	}
}

func TestRandomNumber(t *testing.T) {
	tl := tester.Wrap(t)
	tl.Case("loop 10 times to generate random number as string")
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		n := rand.Intn(32) + 1
		s := Numeric(n)
		tl.Require(n == len(s), "length of generated string should be %d", n)
	}
}

func TestRandomAlphanumeric(t *testing.T) {
	tl := tester.Wrap(t)
	tl.Case("loop 10 times to generate random Alphanumeric")
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		n := rand.Intn(32) + 1
		s := Alphanumeric(n)
		tl.Require(n == len(s), "length of generated string should be %d", n)
	}
}

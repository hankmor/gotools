package random

import (
	"math/rand"
	"strings"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const lettersLen = len(letters)

const numbers = "0123456789"
const numbersLen = len(numbers)

const alphanumeric = letters + numbers
const alphanumericLen = len(alphanumeric)

func Alphabetic(len int) string {
	if len <= 0 {
		panic("illegal argument: length need > 0")
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	builder := strings.Builder{}

	for i := 0; i < len; i++ {
		i := r.Intn(lettersLen)
		builder.WriteByte(letters[i])
	}
	return builder.String()
}

func Numeric(len int) string {
	if len <= 0 {
		panic("illegal argument: length need > 0")
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	builder := strings.Builder{}

	for i := 0; i < len; i++ {
		i := r.Intn(numbersLen)
		builder.WriteByte(numbers[i])
	}
	return builder.String()
}

func Alphanumeric(len int) string {
	if len <= 0 {
		panic("illegal argument: length need > 0")
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	builder := strings.Builder{}

	for i := 0; i < len; i++ {
		i := r.Intn(alphanumericLen)
		builder.WriteByte(alphanumeric[i])
	}
	return builder.String()
}
